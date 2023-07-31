package kubernetesauthbackendconfig

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_kubernetes_auth_backend_config", func(r *config.Resource) {
		r.TerraformResource.Schema["kubernetes_ca_cert"].Sensitive = true
	})
}
