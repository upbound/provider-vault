package vault_namespace

import "github.com/upbound/upjet/pkg/config"

// ConfigureNamespace configures the namespace resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_namespace", func(r *config.Resource) {
		r.Kind = "VaultNamespace"
		r.ShortGroup = "vault"
		r.UseAsync = true
		r.ExternalName = config.IdentifierFromProvider
	})
}
