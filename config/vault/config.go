package vault

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure applies the hand-written resource configurations for the
// resources in the root vault group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_namespace", func(r *config.Resource) {
		r.Kind = "VaultNamespace"
		r.ShortGroup = "vault"
	})

	p.AddResourceConfigurator("vault_raft_snapshot_agent_config", func(r *config.Resource) {
		// The upstream Terraform schema declares the cloud storage
		// credentials as plain strings. They are secrets, so mark them
		// sensitive to generate *SecretRef fields instead of storing
		// them in clear text in the spec.
		for _, key := range []string{
			"aws_secret_access_key",
			"aws_session_token",
			"azure_account_key",
			"google_service_account_key",
		} {
			r.TerraformResource.Schema[key].Sensitive = true
		}
	})
}
