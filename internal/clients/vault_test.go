/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/test"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/vault/api"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	tfvaultschema "github.com/hashicorp/terraform-provider-vault/schema"
	tfvault "github.com/hashicorp/terraform-provider-vault/vault"

	namespacedv1beta1 "github.com/upbound/provider-vault/v4/apis/namespaced/v1beta1"
)

// testAddress is the Vault address these tests hand to the ProviderConfig. It
// is never dialed: TestSetProviderConfiguration only inspects the map that is
// built from it.
const testAddress = "http://vault:8200"

// trueAsString is what the Terraform provider compares add_address_to_env
// against, and what these tests put in a boolean environment variable.
const trueAsString = "true"

func ptr[T any](v T) *T { return &v }

// TestSetProviderConfiguration pins which ProviderConfig fields reach the
// Terraform provider. Writing a key at all suppresses that field's fallback
// chain in the Terraform provider, which is its environment variable and then
// its documented default, so a field the user left unset must leave the key
// absent rather than write its zero value.
func TestSetProviderConfiguration(t *testing.T) {
	cases := map[string]struct {
		spec namespacedv1beta1.ProviderConfigSpec
		want terraform.ProviderConfiguration
	}{
		"unset booleans stay absent": {
			spec: namespacedv1beta1.ProviderConfigSpec{Address: testAddress},
			want: terraform.ProviderConfiguration{
				keyAddress:            testAddress,
				keyMaxLeaseTTLSeconds: 0,
				keyMaxRetries:         0,
				keyMaxRetriesCcc:      0,
			},
		},
		"explicit false is written, not treated as unset": {
			spec: namespacedv1beta1.ProviderConfigSpec{
				Address:               testAddress,
				SkipTLSVerify:         ptr(false),
				SkipChildToken:        ptr(false),
				SkipGetVaultVersion:   ptr(false),
				SetNamespaceFromToken: ptr(false),
			},
			want: terraform.ProviderConfiguration{
				keyAddress:               testAddress,
				keyMaxLeaseTTLSeconds:    0,
				keyMaxRetries:            0,
				keyMaxRetriesCcc:         0,
				keySkipTLSVerify:         false,
				keySkipChildToken:        false,
				keySkipGetVaultVersion:   false,
				keySetNamespaceFromToken: false,
			},
		},
		"explicit true is written": {
			spec: namespacedv1beta1.ProviderConfigSpec{
				Address:               testAddress,
				SkipTLSVerify:         ptr(true),
				SkipChildToken:        ptr(true),
				SkipGetVaultVersion:   ptr(true),
				SetNamespaceFromToken: ptr(true),
			},
			want: terraform.ProviderConfiguration{
				keyAddress:               testAddress,
				keyMaxLeaseTTLSeconds:    0,
				keyMaxRetries:            0,
				keyMaxRetriesCcc:         0,
				keySkipTLSVerify:         true,
				keySkipChildToken:        true,
				keySkipGetVaultVersion:   true,
				keySetNamespaceFromToken: true,
			},
		},
		// The Terraform provider declares add_address_to_env as a string and
		// compares it against "true", so a Go bool arrives as "1" or "0" and
		// never matches. See TestAddAddressToEnvNeedsAString.
		"addAddressToEnv is written as a string": {
			spec: namespacedv1beta1.ProviderConfigSpec{
				Address:         testAddress,
				AddAddressToEnv: ptr(true),
			},
			want: terraform.ProviderConfiguration{
				keyAddress:            testAddress,
				keyMaxLeaseTTLSeconds: 0,
				keyMaxRetries:         0,
				keyMaxRetriesCcc:      0,
				keyAddAddressToEnv:    trueAsString,
			},
		},
		// The Terraform provider declares headers as a list of blocks, and
		// the SDK panics on anything but maps and slices. See
		// TestHeadersReachTheProvider.
		"headers are written as a single element list": {
			spec: namespacedv1beta1.ProviderConfigSpec{
				Address: testAddress,
				Headers: namespacedv1beta1.ProviderHeaders{Name: "X-Vault-Request", Value: trueAsString},
			},
			want: terraform.ProviderConfiguration{
				keyAddress:            testAddress,
				keyMaxLeaseTTLSeconds: 0,
				keyMaxRetries:         0,
				keyMaxRetriesCcc:      0,
				keyHeaders:            []any{map[string]any{"name": "X-Vault-Request", "value": trueAsString}},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ps := terraform.Setup{}
			setProviderConfiguration(&tc.spec, &ps)
			if diff := cmp.Diff(tc.want, ps.Configuration); diff != "" {
				t.Errorf("setProviderConfiguration() -want +got:\n%s", diff)
			}
		})
	}
}

// TestAddAddressToEnvNeedsAString is why setProviderConfiguration formats
// add_address_to_env rather than passing the bool through. The Terraform
// provider declares that field as schema.TypeString and reads it as
// d.Get(...).(string) == "true", and SDKv2 coerces a Go bool to "1" or "0", so
// a bool can never switch it on.
func TestAddAddressToEnvNeedsAString(t *testing.T) {
	cases := map[string]struct {
		value any
		want  string
	}{
		"bool true becomes 1, which never matches": {value: true, want: "1"},
		"bool false becomes 0":                     {value: false, want: "0"},
		"string true survives":                     {value: trueAsString, want: trueAsString},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var got any
			p := schema.Provider{
				Schema: map[string]*schema.Schema{
					keyAddAddressToEnv: {Type: schema.TypeString, Optional: true},
				},
				ConfigureContextFunc: func(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
					got = d.Get(keyAddAddressToEnv)
					return nil, nil
				},
			}
			if d := p.Configure(context.Background(), &tfsdk.ResourceConfig{
				Config: map[string]any{keyAddAddressToEnv: tc.value},
			}); d.HasError() {
				t.Fatalf("Configure() diagnostics = %v", d)
			}
			if got != tc.want {
				t.Errorf("d.Get(%q) = %#v, want %#v", keyAddAddressToEnv, got, tc.want)
			}
		})
	}
}

// TestConfigureNoForkVaultClient exercises the path a running provider takes:
// a ProviderConfig spec becomes ps.Configuration, which schema.Provider.Configure
// consumes without a terraform plan/apply cycle. That is where boolean provider
// config fields were being dropped, so it needs a real Vault to be meaningful.
//
// Set VAULT_ADDR and VAULT_TOKEN to run it, for example:
//
//	docker run -d -p 8200:8200 -e VAULT_DEV_ROOT_TOKEN_ID=root hashicorp/vault
//	VAULT_ADDR=http://127.0.0.1:8200 VAULT_TOKEN=root go test ./internal/clients/...
func TestConfigureNoForkVaultClient(t *testing.T) {
	addr, token := os.Getenv("VAULT_ADDR"), os.Getenv("VAULT_TOKEN")
	if addr == "" || token == "" {
		t.Skip("VAULT_ADDR and VAULT_TOKEN must be set to run this test")
	}

	cases := map[string]struct {
		spec namespacedv1beta1.ProviderConfigSpec
		// env sets TERRAFORM_VAULT_SKIP_CHILD_TOKEN for the subtest when
		// non-empty. Subtests must not run in parallel because of it.
		env            string
		wantChildToken bool
	}{
		"skipChildToken unset: an ephemeral child token is minted": {
			spec:           namespacedv1beta1.ProviderConfigSpec{Address: addr},
			wantChildToken: true,
		},
		"skipChildToken true: the supplied token is used as-is": {
			spec:           namespacedv1beta1.ProviderConfigSpec{Address: addr, SkipChildToken: ptr(true)},
			wantChildToken: false,
		},
		"skipChildToken false: the field wins over the environment": {
			spec:           namespacedv1beta1.ProviderConfigSpec{Address: addr, SkipChildToken: ptr(false)},
			env:            trueAsString,
			wantChildToken: true,
		},
		// Only reachable because an unset field leaves the key absent. A
		// written false would shadow the environment variable entirely.
		"skipChildToken unset: the environment variable applies": {
			spec:           namespacedv1beta1.ProviderConfigSpec{Address: addr},
			env:            trueAsString,
			wantChildToken: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.env != "" {
				t.Setenv("TERRAFORM_VAULT_SKIP_CHILD_TOKEN", tc.env)
			}

			ps := terraform.Setup{}
			setProviderConfiguration(&tc.spec, &ps)
			ps.Configuration[keyToken] = token

			// non-pointer, matching how TerraformSetupBuilder calls it
			p := *tfvaultschema.NewProvider(tfvault.Provider()).SchemaProvider()
			if err := configureNoForkVaultClient(context.Background(), &ps, p); err != nil {
				t.Fatalf("configureNoForkVaultClient() error = %v", err)
			}

			effective := effectiveToken(t, ps)
			if gotChild := effective != token; gotChild != tc.wantChildToken {
				t.Errorf("child token minted = %v, want %v", gotChild, tc.wantChildToken)
			}
		})
	}
}

// effectiveToken reads the token the configured provider actually ended up
// using. ps.Meta holds the Terraform provider's *ProviderMeta, which lives in
// an internal package we cannot name, so match on its method set instead.
func effectiveToken(t *testing.T, ps terraform.Setup) string {
	t.Helper()

	meta, ok := ps.Meta.(interface {
		GetClient() (*api.Client, error)
	})
	if !ok {
		t.Fatalf("ps.Meta does not expose GetClient(), got %T", ps.Meta)
	}

	client, err := meta.GetClient()
	if err != nil {
		t.Fatalf("GetClient() error = %v", err)
	}
	return client.Token()
}

// TestHeadersReachTheProvider feeds the headers block through the Terraform
// provider's own schema for it, which is a list of {name, value} blocks. The
// SDK walks the raw config with reflection and panics on a Go struct, which is
// what setProviderConfiguration used to write.
func TestHeadersReachTheProvider(t *testing.T) {
	var got any
	p := schema.Provider{
		Schema: map[string]*schema.Schema{keyHeaders: tfvault.Provider().Schema[keyHeaders]},
		ConfigureContextFunc: func(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			got = d.Get(keyHeaders)
			return nil, nil
		},
	}

	ps := terraform.Setup{}
	setProviderConfiguration(&namespacedv1beta1.ProviderConfigSpec{
		Address: testAddress,
		Headers: namespacedv1beta1.ProviderHeaders{Name: "X-Vault-Request", Value: trueAsString},
	}, &ps)

	if d := p.Configure(context.Background(), &tfsdk.ResourceConfig{
		Config: map[string]any{keyHeaders: ps.Configuration[keyHeaders]},
	}); d.HasError() {
		t.Fatalf("Configure() diagnostics = %v", d)
	}
	want := []any{map[string]any{"name": "X-Vault-Request", "value": trueAsString}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("d.Get(%q) -want +got:\n%s", keyHeaders, diff)
	}
}

// TestCredentialKeysMatchProviderSchema pins the keys forwarded from the
// credentials JSON against the Terraform provider's schema, so a renamed or
// mistyped key cannot be dropped silently, and every auth_login block the
// Terraform provider knows can be used.
func TestCredentialKeysMatchProviderSchema(t *testing.T) {
	s := tfvault.Provider().Schema
	for _, key := range slices.Concat(credsKeys, authKeys) {
		if _, ok := s[key]; !ok {
			t.Errorf("%q is forwarded but is not a Terraform provider field", key)
		}
	}
	for key := range s {
		if strings.HasPrefix(key, keyAuthLogin) && !slices.Contains(authKeys, key) {
			t.Errorf("Terraform provider auth field %q is not forwarded", key)
		}
	}
}

// TestCommonCredentialsAuthForwardsAuthBlocks checks that a credentials
// Secret with a token and an auth method block reaches ps.Configuration under
// the Terraform provider's keys, with the block wrapped as a list.
func TestCommonCredentialsAuthForwardsAuthBlocks(t *testing.T) {
	creds := []byte(`{"token_name":"tn","auth_login_userpass":{"username":"u","password":"p"},"auth_login_token_file":{"filename":"/tmp/t"}}`)
	kube := &test.MockClient{
		MockGet: test.NewMockGetFn(nil, func(obj client.Object) error {
			if s, ok := obj.(*corev1.Secret); ok {
				s.Data = map[string][]byte{"creds": creds}
			}
			return nil
		}),
	}
	spec := &namespacedv1beta1.ProviderConfigSpec{
		Address: testAddress,
		Credentials: namespacedv1beta1.ProviderCredentials{
			Source: xpv2.CredentialsSourceSecret,
			CommonCredentialSelectors: xpv2.CommonCredentialSelectors{
				SecretRef: &xpv2.SecretKeySelector{
					SecretReference: xpv2.SecretReference{Name: "vault-creds", Namespace: "crossplane-system"},
					Key:             "creds",
				},
			},
		},
	}

	ps := terraform.Setup{Configuration: terraform.ProviderConfiguration{}}
	if err := commonCredentialsAuth(context.Background(), kube, spec, &ps); err != nil {
		t.Fatalf("commonCredentialsAuth() error = %v", err)
	}
	want := terraform.ProviderConfiguration{
		keyTokenName:          "tn",
		keyAuthLoginUserpass:  []any{map[string]any{"username": "u", "password": "p"}},
		keyAuthLoginTokenFile: []any{map[string]any{"filename": "/tmp/t"}},
	}
	if diff := cmp.Diff(want, ps.Configuration); diff != "" {
		t.Errorf("commonCredentialsAuth() -want +got:\n%s", diff)
	}
}
