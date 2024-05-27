package vault

import (
	"github.com/upbound/provider-vault/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// ConfigureNamespace configures the namespace resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_namespace", func(r *config.Resource) {
		r.Kind = "VaultNamespace"
		r.ShortGroup = "vault"
	})
	p.AddResourceConfigurator("vault_identity_group_alias", func(r *config.Resource) {
        r.Kind = "GroupAlias"
        r.ShortGroup = "identity"
        r.References["canonical_id"] = config.Reference{
            Type: "Group",
        }
        r.References["mount_accessor"] = config.Reference{
            Type: "github.com/upbound/provider-vault/apis/jwt/v1alpha1.AuthBackend",
            Extractor: common.AccessorExtractor,
        }
    })
}
