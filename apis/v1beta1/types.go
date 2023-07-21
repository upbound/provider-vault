/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Required origin URL of the Vault server.
	// This is a URL with a scheme, a hostname
	// and a port but with no path.
	Address string `json:"address"`

	// If true the environment variable
	// VAULT_ADDR in the Terraform process environment
	// will be set to the value of the address argument
	// from this provider. By default, this is false.
	// +optional
	AddAddressToEnv bool `json:"add_address_to_env,omitempty"`

	// Set this to true to disable verification
	// of the Vault server's TLS certificate. This is
	// strongly discouraged except in prototype or
	// development environments, since it exposes the
	// possibility that Terraform can be tricked into
	// writing secrets to a server controlled by an intruder.
	// +optional
	SkipTLSVerify bool `json:"skip_tls_verify,omitempty"`

	// Name to use as the SNI host when connecting
	// via TLS.
	// +optional
	TLSServerName string `json:"tls_server_name,omitempty"`

	// Set this to true to disable creation of an
	// intermediate ephemeral Vault token for Terraform to use.
	// Enabling this is strongly discouraged since it increases
	// the potential for a renewable Vault token being exposed
	// in clear text. Only change this setting when the provided
	// token cannot be permitted to create child tokens and there
	// is no risk of exposure from the output of Terraform.
	// +optional
	SkipChildToken bool `json:"skip_child_token,omitempty"`

	// Used as the duration for the intermediate Vault
	// token Terraform issues itself, which in turn limits the
	// duration of secret leases issued by Vault. Defaults to
	// 20 minutes.
	// +optional
	MaxLeaseTTLSeconds int `json:"max_lease_ttl_seconds,omitempty"`

	// Used as the maximum number of retries when a
	// 5xx error code is encountered. Defaults to 2 retries.
	// +optional
	MaxRetries int `json:"max_retries,omitempty"`

	// Maximum number of retries for Client Controlled
	// Consistency related operations. Defaults to 10 retries.
	// +optional
	MaxRetriesCcc int `json:"max_retries_ccc,omitempty"`

	// Set the namespace to use.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Skip the dynamic fetching of the Vault server
	// version. Set to true when the /sys/seal-status API
	// endpoint is not available.
	// +optional
	SkipGetVaultVersion bool `json:"skip_get_vault_version,omitempty"`

	// Override the target Vault server semantic
	// version. Normally the version is dynamically set
	// from the /sys/seal-status API endpoint. In the case
	// where this endpoint is not available an override can
	// be specified here.
	// +optional
	VaultVersionOverride string `json:"vault_version_override,omitempty"`

	// A configuration block, described below,
	// that provides headers to be sent along with all
	// requests to the Vault server. This block can be
	// specified multiple times.
	// +optional
	Headers ProviderHeaders `json:"headers,omitempty"`

	// Credentials required to authenticate to this provider.
	// There are many options to authenticate. They include
	// - token - (Optional) Vault token that will be used
	// by Terraform to authenticate. May be set via the
	// VAULT_TOKEN environment variable. If none is otherwise
	// supplied, Terraform will attempt to read it from
	// ~/.vault-token (where the vault command stores its
	// current token). Terraform will issue itself a new token
	// that is a child of the one given, with a short TTL to
	// limit the exposure of any requested secrets, unless
	// skip_child_token is set to true (see below). Note
	// that the given token must have the update capability
	// on the auth/token/create path in Vault in order to create
	// child tokens. A token is required for the provider. A
	// token can explicitly set via token argument, alternatively
	// a token can be dynamically set via an auth_login* block.
	// +optional
	Credentials ProviderCredentials `json:"credentials"`
}

// ProviderHeaders optional.
type ProviderHeaders struct {
	// Required header name
	Name string `json:"name"` //nolint:unused
	// Required header value
	Value string `json:"value"` //nolint:unused
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a Vault provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,vault}
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,vault}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
