// Package database customizes the Upjet configuration for the Vault database
// secret backend resources.
package database

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// passwordEngines lists the database engine blocks of
// vault_database_secret_backend_connection that carry a root "password"
// attribute. Source of truth: GetConnectionDetailsMapping() in
// apis/cluster/database/v1alpha1/zz_secretbackendconnection_terraformed.go
// (every "<engine>[*].password" -> "<engine>[*].passwordSecretRef" entry).
// Keys are the Terraform (snake_case) attribute names, which is what the
// flatmap InstanceDiff uses ("<engine>.0.password", MaxItems:1 blocks).
var passwordEngines = []string{
	"cassandra",
	"couchbase",
	"elasticsearch",
	"hana",
	"influxdb",
	"mongodb",
	"mssql",
	"mysql",
	"mysql_aurora",
	"mysql_legacy",
	"mysql_rds",
	"oracle",
	"postgresql",
	"redis",
	"redis_elasticache",
	"redshift",
	"snowflake",
}

// Configure makes the root "password" of
// vault_database_secret_backend_connection create-only: it is sent to Vault on
// the initial Create, but never re-sent on Update/reconcile/cold-start. Vault's
// database/config/<name> read API never returns the password, so Upjet's diff
// perpetually sees the password (resolved from passwordSecretRef) as a change
// and re-applies it, clobbering a rotate-root'd root password (28P01). Dropping
// the password attribute from the Update diff is the Crossplane-native
// equivalent of the Terraform `lifecycle.ignore_changes = [<engine>[0].password]`
// already relied upon. Fixes https://github.com/upbound/provider-vault/issues/38.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_database_secret_backend_connection", func(r *config.Resource) {
		r.TerraformCustomDiff = passwordCreateOnlyDiff
	})
}

// passwordCreateOnlyDiff drops every engine "<engine>.0.password" attribute
// from an Update diff while leaving the Create diff (empty state ID) intact.
func passwordCreateOnlyDiff(diff *terraform.InstanceDiff, state *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	// Nothing to adjust on a nil/empty diff or a delete.
	if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
		return diff, nil
	}
	// An empty state ID means this is a Create — let the password through so
	// the connection is provisioned with it.
	if state == nil || state.ID == "" {
		return diff, nil
	}
	// Update/reconcile: drop any engine password attribute so the
	// (now-rotated) password in Vault is not overwritten.
	for _, e := range passwordEngines {
		delete(diff.Attributes, e+".0.password")
	}
	return diff, nil
}
