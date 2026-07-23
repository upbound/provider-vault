/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vault_ad_secret_backend":                            config.IdentifierFromProvider,
	"vault_ad_secret_backend_library":                    config.IdentifierFromProvider,
	"vault_ad_secret_role":                               config.IdentifierFromProvider,
	"vault_alicloud_auth_backend_role":                   config.IdentifierFromProvider,
	"vault_approle_auth_backend_login":                   config.IdentifierFromProvider,
	"vault_approle_auth_backend_role":                    config.IdentifierFromProvider,
	"vault_approle_auth_backend_role_secret_id":          config.IdentifierFromProvider,
	"vault_audit":                                        config.IdentifierFromProvider,
	"vault_audit_request_header":                         config.IdentifierFromProvider,
	"vault_auth_backend":                                 config.IdentifierFromProvider,
	"vault_aws_auth_backend_cert":                        config.IdentifierFromProvider,
	"vault_aws_auth_backend_client":                      config.IdentifierFromProvider,
	"vault_aws_auth_backend_config_identity":             config.IdentifierFromProvider,
	"vault_aws_auth_backend_identity_whitelist":          config.IdentifierFromProvider,
	"vault_aws_auth_backend_login":                       config.IdentifierFromProvider,
	"vault_aws_auth_backend_role":                        config.IdentifierFromProvider,
	"vault_aws_auth_backend_role_tag":                    config.IdentifierFromProvider,
	"vault_aws_auth_backend_roletag_blacklist":           config.IdentifierFromProvider,
	"vault_aws_auth_backend_sts_role":                    config.IdentifierFromProvider,
	"vault_aws_secret_backend":                           config.IdentifierFromProvider,
	"vault_aws_secret_backend_role":                      config.IdentifierFromProvider,
	"vault_azure_auth_backend_config":                    config.IdentifierFromProvider,
	"vault_azure_auth_backend_role":                      config.IdentifierFromProvider,
	"vault_azure_secret_backend":                         config.IdentifierFromProvider,
	"vault_azure_secret_backend_role":                    config.IdentifierFromProvider,
	"vault_cert_auth_backend_role":                       config.IdentifierFromProvider,
	"vault_consul_secret_backend":                        config.IdentifierFromProvider,
	"vault_consul_secret_backend_role":                   config.IdentifierFromProvider,
	"vault_database_secret_backend_connection":           config.IdentifierFromProvider,
	"vault_database_secret_backend_role":                 config.IdentifierFromProvider,
	"vault_database_secret_backend_static_role":          config.IdentifierFromProvider,
	"vault_database_secrets_mount":                       config.IdentifierFromProvider,
	"vault_egp_policy":                                   config.IdentifierFromProvider,
	"vault_gcp_auth_backend":                             config.IdentifierFromProvider,
	"vault_gcp_auth_backend_role":                        config.IdentifierFromProvider,
	"vault_gcp_secret_backend":                           config.IdentifierFromProvider,
	"vault_gcp_secret_impersonated_account":              config.IdentifierFromProvider,
	"vault_gcp_secret_roleset":                           config.IdentifierFromProvider,
	"vault_gcp_secret_static_account":                    config.IdentifierFromProvider,
	"vault_generic_endpoint":                             config.IdentifierFromProvider,
	"vault_generic_secret":                               config.IdentifierFromProvider,
	"vault_github_auth_backend":                          config.IdentifierFromProvider,
	"vault_github_team":                                  config.IdentifierFromProvider,
	"vault_github_user":                                  config.IdentifierFromProvider,
	"vault_identity_entity":                              config.IdentifierFromProvider,
	"vault_identity_entity_alias":                        config.IdentifierFromProvider,
	"vault_identity_entity_policies":                     config.IdentifierFromProvider,
	"vault_identity_group":                               config.IdentifierFromProvider,
	"vault_identity_group_alias":                         config.IdentifierFromProvider,
	"vault_identity_group_member_entity_ids":             config.IdentifierFromProvider,
	"vault_identity_group_member_group_ids":              config.IdentifierFromProvider,
	"vault_identity_group_policies":                      config.IdentifierFromProvider,
	"vault_identity_mfa_duo":                             config.IdentifierFromProvider,
	"vault_identity_mfa_login_enforcement":               config.IdentifierFromProvider,
	"vault_identity_mfa_okta":                            config.IdentifierFromProvider,
	"vault_identity_mfa_pingid":                          config.IdentifierFromProvider,
	"vault_identity_mfa_totp":                            config.IdentifierFromProvider,
	"vault_identity_oidc":                                config.IdentifierFromProvider,
	"vault_identity_oidc_assignment":                     config.IdentifierFromProvider,
	"vault_identity_oidc_client":                         config.IdentifierFromProvider,
	"vault_identity_oidc_key":                            config.IdentifierFromProvider,
	"vault_identity_oidc_key_allowed_client_id":          config.IdentifierFromProvider,
	"vault_identity_oidc_provider":                       config.IdentifierFromProvider,
	"vault_identity_oidc_role":                           config.IdentifierFromProvider,
	"vault_identity_oidc_scope":                          config.IdentifierFromProvider,
	"vault_jwt_auth_backend":                             config.IdentifierFromProvider,
	"vault_jwt_auth_backend_role":                        config.IdentifierFromProvider,
	"vault_kmip_secret_backend":                          config.IdentifierFromProvider,
	"vault_kmip_secret_role":                             config.IdentifierFromProvider,
	"vault_kmip_secret_scope":                            config.IdentifierFromProvider,
	"vault_kubernetes_auth_backend_config":               config.IdentifierFromProvider,
	"vault_kubernetes_auth_backend_role":                 config.IdentifierFromProvider,
	"vault_kubernetes_secret_backend":                    config.IdentifierFromProvider,
	"vault_kubernetes_secret_backend_role":               config.IdentifierFromProvider,
	"vault_kv_secret":                                    config.IdentifierFromProvider,
	"vault_kv_secret_backend_v2":                         config.IdentifierFromProvider,
	"vault_kv_secret_v2":                                 config.IdentifierFromProvider,
	"vault_ldap_auth_backend":                            config.IdentifierFromProvider,
	"vault_ldap_auth_backend_group":                      config.IdentifierFromProvider,
	"vault_ldap_auth_backend_user":                       config.IdentifierFromProvider,
	"vault_managed_keys":                                 config.IdentifierFromProvider,
	"vault_mfa_duo":                                      config.IdentifierFromProvider,
	"vault_mfa_okta":                                     config.IdentifierFromProvider,
	"vault_mfa_pingid":                                   config.IdentifierFromProvider,
	"vault_mfa_totp":                                     config.IdentifierFromProvider,
	"vault_mongodbatlas_secret_backend":                  config.IdentifierFromProvider,
	"vault_mongodbatlas_secret_role":                     config.IdentifierFromProvider,
	"vault_mount":                                        config.IdentifierFromProvider,
	"vault_namespace":                                    config.IdentifierFromProvider,
	"vault_nomad_secret_backend":                         config.IdentifierFromProvider,
	"vault_nomad_secret_role":                            config.IdentifierFromProvider,
	"vault_okta_auth_backend":                            config.IdentifierFromProvider,
	"vault_okta_auth_backend_group":                      config.IdentifierFromProvider,
	"vault_okta_auth_backend_user":                       config.IdentifierFromProvider,
	"vault_pki_secret_backend_cert":                      config.IdentifierFromProvider,
	"vault_pki_secret_backend_config_ca":                 config.IdentifierFromProvider,
	"vault_pki_secret_backend_config_urls":               config.IdentifierFromProvider,
	"vault_pki_secret_backend_crl_config":                config.IdentifierFromProvider,
	"vault_pki_secret_backend_intermediate_cert_request": config.IdentifierFromProvider,
	"vault_pki_secret_backend_intermediate_set_signed":   config.IdentifierFromProvider,
	"vault_pki_secret_backend_role":                      config.IdentifierFromProvider,
	"vault_pki_secret_backend_root_cert":                 config.IdentifierFromProvider,
	"vault_pki_secret_backend_root_sign_intermediate":    config.IdentifierFromProvider,
	"vault_pki_secret_backend_sign":                      config.IdentifierFromProvider,
	"vault_policy":                                       config.IdentifierFromProvider,
	"vault_quota_lease_count":                            config.IdentifierFromProvider,
	"vault_quota_rate_limit":                             config.IdentifierFromProvider,
	"vault_rabbitmq_secret_backend":                      config.IdentifierFromProvider,
	"vault_rabbitmq_secret_backend_role":                 config.IdentifierFromProvider,
	"vault_raft_autopilot":                               config.IdentifierFromProvider,
	"vault_raft_snapshot_agent_config":                   config.IdentifierFromProvider,
	"vault_rgp_policy":                                   config.IdentifierFromProvider,
	"vault_ssh_secret_backend_ca":                        config.IdentifierFromProvider,
	"vault_ssh_secret_backend_role":                      config.IdentifierFromProvider,
	"vault_terraform_cloud_secret_backend":               config.IdentifierFromProvider,
	"vault_terraform_cloud_secret_creds":                 config.IdentifierFromProvider,
	"vault_terraform_cloud_secret_role":                  config.IdentifierFromProvider,
	"vault_token":                                        config.IdentifierFromProvider,
	"vault_token_auth_backend_role":                      config.IdentifierFromProvider,
	"vault_transform_alphabet":                           config.IdentifierFromProvider,
	"vault_transform_role":                               config.IdentifierFromProvider,
	"vault_transform_template":                           config.IdentifierFromProvider,
	"vault_transform_transformation":                     config.IdentifierFromProvider,
	"vault_transit_secret_backend_cache_config":          config.IdentifierFromProvider,
	"vault_transit_secret_backend_key":                   config.IdentifierFromProvider,
}

var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"vault_password_policy": vaultPasswordPolicy(),
}

// vaultPasswordPolicy returns the external name configuration for
// vault_password_policy TF resource.
// Ideally, this resource should have been NameAsIdentifier,
// however, keeping it unnormalized to preserve backward compatibility.
func vaultPasswordPolicy() config.ExternalName { //nolint:gocyclo
	e := config.IdentifierFromProvider
	e.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
		if externalName != "" {
			base["name"] = externalName
		}

	}
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		if externalName == "" {
			return parameters["name"].(string), nil
		}
		return externalName, nil
	}
	// NOTE: Upstream vault client impl surfaces 404s as nil struct + nil error
	// ref: https://github.com/hashicorp/vault/blob/99d122d00c97ba40c6c8965f02edcb0064030c78/api/logical.go#L154
	// Upstream TF provider Read impl always surfaces this as an error-severity diagnostics
	// ref: https://github.com/hashicorp/terraform-provider-vault/blob/0f882c6450f1405418e6349e35354f2209a3af5c/internal/vault/sys/password_policy.go#L175-L178
	e.IsNotFoundDiagnosticFn = func(diags []*tfprotov6.Diagnostic) bool {
		for _, diag := range diags {
			if diag.Severity == tfprotov6.DiagnosticSeverityError &&
				diag.Summary == "Unable to Read Resource from Vault" &&
				strings.HasSuffix(diag.Detail, "Vault response was nil") {
				return true
			}
		}
		return false
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		// If an external name is configured for multiple architectures,
		// Terraform Plugin Framework takes precedence over Terraform
		// Plugin SDKv2
		e, configured := TerraformPluginFrameworkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = ExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.ExternalName = e
	}
}
