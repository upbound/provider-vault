package vault

import (
	"github.com/upbound/upjet/pkg/config"
)

// ConfigureNamespace configures the namespace resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_namespace", func(r *config.Resource) {
		r.Kind = "VaultNamespace"
		r.ShortGroup = "vault"
	})
	p.AddResourceConfigurator("vault_identity_group_alias", func(r *config.Resource) {
		r.References = config.References{
			"canonicalIdFromGroupRef": {
				TerraformName: "vault_identity_group",
				Extractor:     "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
		}
	})
}
