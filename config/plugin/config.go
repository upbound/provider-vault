/*
Copyright 2026 Upbound Inc.
*/

package plugin

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the plugin group resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_plugin", func(r *config.Resource) {
		r.ShortGroup = "plugin"
	})
}
