/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-vault/xpprovider"
	"github.com/hashicorp/vault/api"

	clusterv1beta1 "github.com/upbound/provider-vault/v4/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-vault/v4/apis/namespaced/v1beta1"
)

const (
	// provider config
	// source: https://registry.terraform.io/providers/hashicorp/vault/latest/docs

	// required origin URL of the Vault server
	keyAddress = "address"

	// optional provider arguments
	// remove options that we do not yet want to implement
	keyAddAddressToEnv       = "add_address_to_env"
	keyToken                 = "token"
	keyTokenName             = "token_name"
	keyCaCertFile            = "ca_cert_file"
	keyCaCertDir             = "ca_cert_dir"
	keyAuthLoginUserpass     = "auth_login_userpath"
	keyAuthLoginAWS          = "auth_login_aws"
	keyAuthLoginCert         = "auth_login_cert"
	keyAuthLoginGCP          = "auth_login_gcp"
	keyAuthLoginKerberos     = "auth_login_kerberos"
	keyAuthLoginRadius       = "auth_login_radius"
	keyAuthLoginOCI          = "auth_login_oci"
	keyAuthLoginOIDC         = "auth_login_oidc"
	keyAuthLoginJWT          = "auth_login_jwt"
	keyAuthLoginAzure        = "auth_login_azure"
	keyAuthLogin             = "auth_login"
	keyClientAuth            = "client_auth"
	keySkipTLSVerify         = "skip_tls_verify"
	keyTLSServerName         = "tls_server_name"
	keySkipChildToken        = "skip_child_token"
	keyMaxLeaseTTLSeconds    = "max_lease_ttl_seconds"
	keyMaxRetries            = "max_retries"
	keyMaxRetriesCcc         = "max_retries_ccc"
	keyNamespace             = "namespace"
	keySkipGetVaultVersion   = "skip_get_vault_version"
	keySetNamespaceFromToken = "set_namespace_from_token"
	keyVaultVersionOverride  = "vault_version_override"
	keyHeaders               = "headers"
	keyRole                  = "role"

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

	credentialsSourceKubernetes xpv2.CredentialsSource = "Kubernetes"

	errConfigure = "failed to configure the no-fork Vault client"
)

// tokenRefreshMargin is how long before its expiry a cached token stops being
// reused, so that a reconcile never starts with a token about to expire.
const tokenRefreshMargin = time.Minute

// tokenReuseWithoutTTL bounds the reuse of tokens that report no TTL, such as
// root tokens, so that the lookup still runs regularly.
const tokenReuseWithoutTTL = 20 * time.Minute

var now = time.Now

// A tokenCache remembers the token the Terraform provider ended up with for a
// given provider configuration. The Terraform provider logs in and mints a
// child token every time it is configured, and it is configured on every
// reconcile, so without the cache every reconcile of every resource creates
// one or two Vault leases that are never revoked.
type tokenCache struct {
	mu      sync.Mutex
	entries map[string]cachedToken
}

type cachedToken struct {
	token     string
	expiresAt time.Time
}

var tokens = &tokenCache{entries: map[string]cachedToken{}}

func (c *tokenCache) get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.entries[key]
	if !ok {
		return "", false
	}
	if !now().Add(tokenRefreshMargin).Before(e.expiresAt) {
		delete(c.entries, key)
		return "", false
	}
	return e.token, true
}

func (c *tokenCache) put(key, token string, expiresAt time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, e := range c.entries {
		if !now().Before(e.expiresAt) {
			delete(c.entries, k)
		}
	}
	c.entries[key] = cachedToken{token: token, expiresAt: expiresAt}
}

func (c *tokenCache) drop(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entries, key)
}

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

		setProviderConfiguration(pcSpec, &ps)

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

		return configureWithCachedToken(ctx, ps, *tfProvider)
	}
}

// configureWithCachedToken configures the Terraform provider, reusing the
// token of an earlier login for the same configuration while it is valid. A
// reused token is passed as the token with skip_child_token set and the
// auth_login blocks removed, so the Terraform provider only looks the token
// up instead of logging in and creating a child token again. If that lookup
// fails, for example because the token was revoked, the cached entry is
// dropped and the original configuration logs in again.
func configureWithCachedToken(ctx context.Context, ps terraform.Setup, p schema.Provider) (terraform.Setup, error) {
	key, err := configurationKey(ps.Configuration)
	if err != nil {
		return ps, errors.Wrap(err, errConfigure)
	}

	if token, ok := tokens.get(key); ok {
		cached := ps
		cached.Configuration = withToken(ps.Configuration, token)
		if err := configureNoForkVaultClient(ctx, &cached, p); err == nil {
			return cached, nil
		}
		tokens.drop(key)
	}

	if err := configureNoForkVaultClient(ctx, &ps, p); err != nil {
		return ps, errors.Wrap(err, errConfigure)
	}
	if token, expiresAt, ok := configuredToken(ps.Meta); ok {
		tokens.put(key, token, expiresAt)
	}
	return ps, nil
}

// configurationKey identifies a provider configuration, credentials included,
// so that two ProviderConfigs only share a token when they are identical.
func configurationKey(cfg terraform.ProviderConfiguration) (string, error) {
	data, err := json.Marshal(cfg)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

// withToken returns a copy of the configuration that authenticates with the
// given token instead of any auth_login block, and keeps it as is.
func withToken(cfg terraform.ProviderConfiguration, token string) terraform.ProviderConfiguration {
	out := make(terraform.ProviderConfiguration, len(cfg)+2)
	for k, v := range cfg {
		if strings.HasPrefix(k, keyAuthLogin) {
			continue
		}
		out[k] = v
	}
	out[keyToken] = token
	out[keySkipChildToken] = true
	return out
}

// configuredToken returns the token the configured Terraform provider uses and
// when it expires, read from the token itself so that a child token and a
// login token are handled alike.
func configuredToken(meta any) (string, time.Time, bool) {
	m, ok := meta.(interface {
		GetClient() (*api.Client, error)
	})
	if !ok {
		return "", time.Time{}, false
	}
	c, err := m.GetClient()
	if err != nil || c == nil || c.Token() == "" {
		return "", time.Time{}, false
	}
	info, err := c.Auth().Token().LookupSelf()
	if err != nil || info == nil {
		return "", time.Time{}, false
	}
	ttl, err := info.TokenTTL()
	if err != nil {
		return "", time.Time{}, false
	}
	if ttl <= 0 {
		ttl = tokenReuseWithoutTTL
	}
	return c.Token(), now().Add(ttl), true
}

// setProviderConfiguration translates a ProviderConfig spec into the Terraform
// provider configuration that schema.Provider.Configure consumes.
//
// A field is only written when the user actually set it, unless the Terraform
// provider's own default is the Go zero value. Writing a zero value for an
// unset field is not harmless: it is indistinguishable from the user asking
// for it, so it silently overrides any non-zero default the Terraform provider
// documents.
func setProviderConfiguration(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) {
	ps.Configuration = map[string]any{}

	// Assign mandatory address parameter
	ps.Configuration[keyAddress] = pcSpec.Address

	// Assign optional parameters.
	//
	// The booleans below are pointers so that "unset" stays distinguishable
	// from "set to false". Writing a key at all suppresses the Terraform
	// provider's own fallback chain for that field, which is its VAULT_* or
	// TERRAFORM_VAULT_* environment variable and then its documented default,
	// so an unset field must leave the key absent.
	//
	// add_address_to_env is the odd one out: the Terraform provider declares it
	// as a string and compares it against "true", so a Go bool reaches it as
	// "1" or "0" and never matches. Format it explicitly.
	if pcSpec.AddAddressToEnv != nil {
		ps.Configuration[keyAddAddressToEnv] = strconv.FormatBool(*pcSpec.AddAddressToEnv)
	}
	if pcSpec.SkipTLSVerify != nil {
		ps.Configuration[keySkipTLSVerify] = *pcSpec.SkipTLSVerify
	}
	if len(pcSpec.TLSServerName) > 0 {
		ps.Configuration[keyTLSServerName] = pcSpec.TLSServerName
	}
	if pcSpec.SkipChildToken != nil {
		ps.Configuration[keySkipChildToken] = *pcSpec.SkipChildToken
	}
	ps.Configuration[keyMaxLeaseTTLSeconds] = pcSpec.MaxLeaseTTLSeconds
	ps.Configuration[keyMaxRetries] = pcSpec.MaxRetries
	ps.Configuration[keyMaxRetriesCcc] = pcSpec.MaxRetriesCcc
	if len(pcSpec.Namespace) > 0 {
		ps.Configuration[keyNamespace] = pcSpec.Namespace
	}
	if pcSpec.SkipGetVaultVersion != nil {
		ps.Configuration[keySkipGetVaultVersion] = *pcSpec.SkipGetVaultVersion
	}
	// set_namespace_from_token defaults to true, so writing false for an unset
	// field would silently flip that default.
	if pcSpec.SetNamespaceFromToken != nil {
		ps.Configuration[keySetNamespaceFromToken] = *pcSpec.SetNamespaceFromToken
	}
	if len(pcSpec.VaultVersionOverride) > 0 {
		ps.Configuration[keyVaultVersionOverride] = pcSpec.VaultVersionOverride
	}
	if pcSpec.Headers != (namespacedv1beta1.ProviderHeaders{}) {
		ps.Configuration[keyHeaders] = pcSpec.Headers
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

	mount := "kubernetes"
	if pcSpec.MountPath != nil && *pcSpec.MountPath != "" {
		mount = *pcSpec.MountPath
	}

	ps.Configuration[keyAuthLoginJWT] = []any{
		map[string]string{
			"jwt":   string(jwt),
			"mount": mount,
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
	ps.FrameworkProvider = xpprovider.FrameworkProvider(&p)
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
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
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
