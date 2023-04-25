package raftautopilot

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_raft_autopilot", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vault"
	})
}
