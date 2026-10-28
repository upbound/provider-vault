/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/vault/api"

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

// fakeMeta is what a configured Terraform provider exposes to effectiveToken.
type fakeMeta struct{ client *api.Client }

func (m fakeMeta) GetClient() (*api.Client, error) { return m.client, nil }

// newTokenServer answers auth/token/lookup-self for the tokens it knows with
// their TTL in seconds and refuses every other token, like Vault does for a
// revoked one.
func newTokenServer(ttls map[string]int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/auth/token/lookup-self" {
			http.Error(w, r.URL.Path, http.StatusNotFound)
			return
		}
		ttl, ok := ttls[r.Header.Get("X-Vault-Token")]
		if !ok {
			http.Error(w, `{"errors":["permission denied"]}`, http.StatusForbidden)
			return
		}
		fmt.Fprintf(w, `{"data":{"ttl":%d}}`, ttl)
	}))
}

// loginProvider mimics the Terraform provider's Configure: a configuration
// with an auth_login block logs in and gets a fresh token, a configuration
// with a token uses it as is, and either way the token is looked up, which
// fails Configure when Vault refuses it.
func loginProvider(address string, logins *[]string) schema.Provider {
	return schema.Provider{
		Schema: tfvault.Provider().Schema,
		ConfigureContextFunc: func(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			token, _ := d.Get(keyToken).(string)
			if _, ok := d.GetOk(keyAuthLoginJWT); ok {
				token = fmt.Sprintf("login-%d", len(*logins)+1)
				*logins = append(*logins, token)
			}
			cfg := api.DefaultConfig()
			cfg.Address = address
			c, err := api.NewClient(cfg)
			if err != nil {
				return nil, diag.FromErr(err)
			}
			c.SetToken(token)
			if _, err := c.Auth().Token().LookupSelf(); err != nil {
				return nil, diag.FromErr(err)
			}
			return fakeMeta{client: c}, nil
		},
	}
}

const (
	readerRole = "reader"
	writerRole = "writer"
	firstLogin = "login-1"
)

func jwtConfiguration(role string) terraform.ProviderConfiguration {
	return terraform.ProviderConfiguration{
		keyAddress:      testAddress,
		keyAuthLoginJWT: []any{map[string]any{"jwt": "sa-token", "mount": "kubernetes", "role": role}},
	}
}

// TestConfigureWithCachedToken pins the token lifecycle across reconciles:
// one login per configuration while its token is valid, a new login when the
// token nears expiry or is refused, and no sharing between configurations.
func TestConfigureWithCachedToken(t *testing.T) {
	ttls := map[string]int{firstLogin: 1200, "login-2": 1200, "login-3": 0, "login-4": 1200}
	server := newTokenServer(ttls)
	defer server.Close()

	base := time.Date(2026, 10, 27, 12, 0, 0, 0, time.UTC)
	clock := base
	now = func() time.Time { return clock }
	tokens = &tokenCache{entries: map[string]cachedToken{}}
	t.Cleanup(func() {
		now = time.Now
		tokens = &tokenCache{entries: map[string]cachedToken{}}
	})

	var logins []string
	p := loginProvider(server.URL, &logins)
	configure := func(cfg terraform.ProviderConfiguration) terraform.Setup {
		t.Helper()
		got, err := configureWithCachedToken(context.Background(), terraform.Setup{Configuration: cfg}, p)
		if err != nil {
			t.Fatalf("configureWithCachedToken() error = %v", err)
		}
		return got
	}

	first := configure(jwtConfiguration(readerRole))
	second := configure(jwtConfiguration(readerRole))
	if diff := cmp.Diff([]string{firstLogin}, logins); diff != "" {
		t.Errorf("logins after two reconciles of one configuration: -want +got:\n%s", diff)
	}
	if got := effectiveToken(t, second); got != firstLogin {
		t.Errorf("second reconcile uses token %q, want the first login's", got)
	}
	if _, ok := second.Configuration[keyAuthLoginJWT]; ok {
		t.Errorf("second reconcile still carries %s", keyAuthLoginJWT)
	}
	if second.Configuration[keySkipChildToken] != true || second.Configuration[keyToken] != firstLogin {
		t.Errorf("second reconcile configuration = %v, want the cached token with skip_child_token", second.Configuration)
	}
	if first.Configuration[keyToken] != nil {
		t.Errorf("first reconcile configuration gained a token: %v", first.Configuration[keyToken])
	}

	configure(jwtConfiguration(writerRole))
	if diff := cmp.Diff([]string{firstLogin, "login-2"}, logins); diff != "" {
		t.Errorf("logins after a different configuration: -want +got:\n%s", diff)
	}

	clock = base.Add(1200*time.Second - tokenRefreshMargin)
	configure(jwtConfiguration(readerRole))
	if diff := cmp.Diff([]string{firstLogin, "login-2", "login-3"}, logins); diff != "" {
		t.Errorf("logins once the token nears expiry: -want +got:\n%s", diff)
	}

	// login-3 reports no TTL, so it is reused for a bounded time only.
	clock = clock.Add(tokenReuseWithoutTTL - tokenRefreshMargin - time.Second)
	if got := effectiveToken(t, configure(jwtConfiguration(readerRole))); got != "login-3" {
		t.Errorf("token without TTL not reused within the bound, got %q", got)
	}
	delete(ttls, "login-3")
	if got := effectiveToken(t, configure(jwtConfiguration(readerRole))); got != "login-4" {
		t.Errorf("refused token not replaced by a new login, got %q", got)
	}
	if diff := cmp.Diff([]string{firstLogin, "login-2", "login-3", "login-4"}, logins); diff != "" {
		t.Errorf("logins at the end: -want +got:\n%s", diff)
	}
}

// TestConfigureWithCachedTokenAgainstVault counts the tokens that repeated
// configurations create on a real Vault. Set VAULT_ADDR and VAULT_TOKEN to
// run it, see TestConfigureNoForkVaultClient.
func TestConfigureWithCachedTokenAgainstVault(t *testing.T) {
	addr, token := os.Getenv("VAULT_ADDR"), os.Getenv("VAULT_TOKEN")
	if addr == "" || token == "" {
		t.Skip("VAULT_ADDR and VAULT_TOKEN must be set to run this test")
	}
	tokens = &tokenCache{entries: map[string]cachedToken{}}
	t.Cleanup(func() { tokens = &tokenCache{entries: map[string]cachedToken{}} })

	cfg := api.DefaultConfig()
	cfg.Address = addr
	root, err := api.NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	root.SetToken(token)
	accessors := func() int {
		t.Helper()
		s, err := root.Logical().List("auth/token/accessors")
		if err != nil {
			t.Fatalf("cannot list token accessors: %v", err)
		}
		keys, _ := s.Data["keys"].([]any)
		return len(keys)
	}

	const reconciles = 5
	before := accessors()
	effective := make([]string, 0, reconciles)
	for range reconciles {
		ps := terraform.Setup{}
		setProviderConfiguration(&namespacedv1beta1.ProviderConfigSpec{Address: addr}, &ps)
		ps.Configuration[keyToken] = token
		p := *tfvaultschema.NewProvider(tfvault.Provider()).SchemaProvider()
		got, err := configureWithCachedToken(context.Background(), ps, p)
		if err != nil {
			t.Fatalf("configureWithCachedToken() error = %v", err)
		}
		effective = append(effective, effectiveToken(t, got))
	}
	if created := accessors() - before; created != 1 {
		t.Errorf("%d reconciles created %d tokens, want 1 child token", reconciles, created)
	}
	for i, tok := range effective {
		if tok == token {
			t.Errorf("reconcile %d used the configured token itself, want the child token", i)
		}
		if tok != effective[0] {
			t.Errorf("reconcile %d used token %q, want the child token of the first reconcile", i, tok)
		}
	}
}
