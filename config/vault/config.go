package vault

import "github.com/crossplane/upjet/v2/pkg/config"

// ConfigureNamespace configures the namespace resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_namespace", func(r *config.Resource) {
		r.Kind = "VaultNamespace"
		r.ShortGroup = "vault"
	})
}
