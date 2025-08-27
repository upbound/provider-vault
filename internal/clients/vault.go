/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"os"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	clusterv1beta1 "github.com/upbound/provider-vault/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-vault/apis/namespaced/v1beta1"
)

const (
	// provider config
	// source: https://registry.terraform.io/providers/hashicorp/vault/latest/docs

	// required origin URL of the Vault server
	keyAddress = "address"

	// optional provider arguments
	// remove options that we do not yet want to implement
	keyAddAddressToEnv      = "add_address_to_env"
	keyToken                = "token"
	keyTokenName            = "token_name"
	keyCaCertFile           = "ca_cert_file"
	keyCaCertDir            = "ca_cert_dir"
	keyAuthLoginUserpass    = "auth_login_userpath"
	keyAuthLoginAWS         = "auth_login_aws"
	keyAuthLoginCert        = "auth_login_cert"
	keyAuthLoginGCP         = "auth_login_gcp"
	keyAuthLoginKerberos    = "auth_login_kerberos"
	keyAuthLoginRadius      = "auth_login_radius"
	keyAuthLoginOCI         = "auth_login_oci"
	keyAuthLoginOIDC        = "auth_login_oidc"
	keyAuthLoginJWT         = "auth_login_jwt"
	keyAuthLoginAzure       = "auth_login_azure"
	keyAuthLogin            = "auth_login"
	keyClientAuth           = "client_auth"
	keySkipTLSVerify        = "skip_tls_verify"
	keyTLSServerName        = "tls_server_name"
	keySkipChildToken       = "skip_child_token"
	keyMaxLeaseTTLSeconds   = "max_lease_ttl_seconds"
	keyMaxRetries           = "max_retries"
	keyMaxRetriesCcc        = "max_retries_ccc"
	keyNamespace            = "namespace"
	keySkipGetVaultVersion  = "skip_get_vault_version"
	keyVaultVersionOverride = "vault_version_override"
	keyHeaders              = "headers"
	keyRole                 = "role"

	// error messages
	errNoProviderConfig      = "no providerConfigRef provided"
	errGetProviderConfig     = "cannot get referenced ProviderConfig"
	errTrackUsage            = "cannot track ProviderConfig usage"
	errExtractCredentials    = "cannot extract credentials"
	errUnmarshalCredentials  = "cannot unmarshal vault credentials as JSON"
	errNoServiceAccountToken = "no service account token found"
	errNoRole                = "no role provided"

	// Service account token path
	serviceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"

	credentialsSourceKubernetes xpv1.CredentialsSource = "Kubernetes"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
// nolint:gocyclo
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		// set provider configuration
		ps.Configuration = map[string]any{}

		// Assign mandatory address parameter
		ps.Configuration[keyAddress] = pcSpec.Address

		// Assign optional parameters
		ps.Configuration[keyAddAddressToEnv] = pcSpec.AddAddressToEnv
		ps.Configuration[keySkipTLSVerify] = pcSpec.SkipTLSVerify
		if len(pcSpec.TLSServerName) > 0 {
			ps.Configuration[keyTLSServerName] = pcSpec.TLSServerName
		}
		ps.Configuration[keySkipChildToken] = pcSpec.SkipChildToken
		ps.Configuration[keyMaxLeaseTTLSeconds] = pcSpec.MaxLeaseTTLSeconds
		ps.Configuration[keyMaxRetries] = pcSpec.MaxRetries
		ps.Configuration[keyMaxRetriesCcc] = pcSpec.MaxRetriesCcc
		if len(pcSpec.Namespace) > 0 {
			ps.Configuration[keyNamespace] = pcSpec.Namespace
		}
		ps.Configuration[keySkipGetVaultVersion] = pcSpec.SkipGetVaultVersion
		if len(pcSpec.VaultVersionOverride) > 0 {
			ps.Configuration[keyVaultVersionOverride] = pcSpec.VaultVersionOverride
		}
		if pcSpec.Headers != (namespacedv1beta1.ProviderHeaders{}) {
			ps.Configuration[keyHeaders] = pcSpec.Headers
		}

		switch pcSpec.Credentials.Source { //nolint:exhaustive
		case credentialsSourceKubernetes:
			if err := kubernetesAuth(pcSpec, &ps); err != nil {
				return ps, err
			}
		default:
			if err := commonCredentialsAuth(ctx, client, pcSpec, &ps); err != nil {
				return ps, err
			}
		}

		return ps, errors.Wrap(
			configureNoForkVaultClient(ctx, &ps, *tfProvider),
			"failed to configure the no-fork Vault client",
		)
	}
}

func commonCredentialsAuth(ctx context.Context, client client.Client, pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, errExtractCredentials)
	}

	creds := map[string]any{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return errors.Wrap(err, errUnmarshalCredentials)
	}

	// Set credentials in Terraform
	// provider configuration
	credsKeys := [...]string{keyToken, keyTokenName, keyCaCertFile, keyCaCertDir}
	for _, key := range credsKeys {
		if v, ok := creds[key]; ok {
			ps.Configuration[key] = v
		}
	}
	// structured auth methods need to be wrapped in a single element array
	// see: https://registry.terraform.io/providers/hashicorp/vault/latest/docs#vault-authentication-configuration-options
	authKeys := [...]string{
		keyAuthLoginUserpass, keyAuthLoginAWS,
		keyAuthLoginCert, keyAuthLoginGCP, keyAuthLoginKerberos,
		keyAuthLoginRadius, keyAuthLoginOCI, keyAuthLoginOIDC,
		keyAuthLoginJWT, keyAuthLoginAzure, keyAuthLogin, keyClientAuth,
	}
	for _, key := range authKeys {
		if v, ok := creds[key]; ok {
			ps.Configuration[key] = []interface{}{v}
		}
	}

	return nil
}

func kubernetesAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	jwt, err := os.ReadFile(serviceAccountTokenPath)
	if err != nil {
		return errors.Wrap(err, errNoServiceAccountToken)
	}

	if pcSpec.Role == nil || *pcSpec.Role == "" {
		return errors.New(errNoRole)
	}

	ps.Configuration[keyAuthLoginJWT] = []any{
		map[string]string{
			"jwt":   string(jwt),
			"mount": "kubernetes",
			"role":  *pcSpec.Role,
		},
	}

	return nil
}

func configureNoForkVaultClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}

func legacyToModernProviderConfigSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	// TODO(erhan): this is hacky and potentially lossy, generate or manually implement
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveProviderConfigLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveProviderConfigModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed")
	}
}

func resolveProviderConfigLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return legacyToModernProviderConfigSpec(pc)
}

func resolveProviderConfigModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrapf(err, "referenced provider config kind %q is invalid for %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}
	pcObj, ok := pcRuntimeObj.(resource.ProviderConfig)
	if !ok {
		return nil, errors.Errorf("referenced provider config kind %q is not a provider config type %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.CommonCredentialSelectors.SecretRef != nil {
			pcSpec.Credentials.CommonCredentialSelectors.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config kind")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
