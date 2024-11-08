/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthBackendInitParameters struct {

	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	// The API endpoint to use. Useful if you are running GitHub Enterprise or an API-compatible authentication server.
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	// Specifies the description of the mount. This overrides the current stored value, if any.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization configured users must be part of.
	// The organization users must be part of.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	// The ID of the organization users must be part of. Vault will attempt to fetch and set this value if it is not provided (vault-1.10+)
	OrganizationID *float64 `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Path where the auth backend is mounted. Defaults to auth/github
	// if not specified.
	// Path where the auth backend is mounted
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// If set, will encode an
	// explicit max TTL
	// onto the token in number of seconds. This is a hard cap even if token_ttl and
	// token_max_ttl would otherwise allow a renewal.
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	// Generated Token's Policies
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token that should be generated. Can be service,
	// batch, or default to use the mount's tuned default (which unless changed will be
	// service tokens). For token store roles, there are two additional possibilities:
	// default-service and default-batch which specify the type to return unless the client
	// requests a different type at generation time.
	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`

	// Extra configuration block. Structure is documented below.
	Tune *TuneInitParameters `json:"tune,omitempty" tf:"tune,omitempty"`
}

type AuthBackendObservation struct {

	// The mount accessor related to the auth mount. It is useful for integration with Identity Secrets Engine.
	// The mount accessor related to the auth mount.
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	// The API endpoint to use. Useful if you are running GitHub Enterprise or an API-compatible authentication server.
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	// Specifies the description of the mount. This overrides the current stored value, if any.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization configured users must be part of.
	// The organization users must be part of.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	// The ID of the organization users must be part of. Vault will attempt to fetch and set this value if it is not provided (vault-1.10+)
	OrganizationID *float64 `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Path where the auth backend is mounted. Defaults to auth/github
	// if not specified.
	// Path where the auth backend is mounted
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// If set, will encode an
	// explicit max TTL
	// onto the token in number of seconds. This is a hard cap even if token_ttl and
	// token_max_ttl would otherwise allow a renewal.
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	// Generated Token's Policies
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token that should be generated. Can be service,
	// batch, or default to use the mount's tuned default (which unless changed will be
	// service tokens). For token store roles, there are two additional possibilities:
	// default-service and default-batch which specify the type to return unless the client
	// requests a different type at generation time.
	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`

	// Extra configuration block. Structure is documented below.
	Tune *TuneObservation `json:"tune,omitempty" tf:"tune,omitempty"`
}

type AuthBackendParameters struct {

	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	// The API endpoint to use. Useful if you are running GitHub Enterprise or an API-compatible authentication server.
	// +kubebuilder:validation:Optional
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	// Specifies the description of the mount. This overrides the current stored value, if any.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization configured users must be part of.
	// The organization users must be part of.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	// The ID of the organization users must be part of. Vault will attempt to fetch and set this value if it is not provided (vault-1.10+)
	// +kubebuilder:validation:Optional
	OrganizationID *float64 `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Path where the auth backend is mounted. Defaults to auth/github
	// if not specified.
	// Path where the auth backend is mounted
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// If set, will encode an
	// explicit max TTL
	// onto the token in number of seconds. This is a hard cap even if token_ttl and
	// token_max_ttl would otherwise allow a renewal.
	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token that should be generated. Can be service,
	// batch, or default to use the mount's tuned default (which unless changed will be
	// service tokens). For token store roles, there are two additional possibilities:
	// default-service and default-batch which specify the type to return unless the client
	// requests a different type at generation time.
	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`

	// Extra configuration block. Structure is documented below.
	// +kubebuilder:validation:Optional
	Tune *TuneParameters `json:"tune,omitempty" tf:"tune,omitempty"`
}

type TuneInitParameters struct {

	// List of headers to whitelist and allowing
	// a plugin to include them in the response.
	AllowedResponseHeaders []*string `json:"allowedResponseHeaders,omitempty" tf:"allowed_response_headers"`

	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the request data object.
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys"`

	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the response data object.
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys"`

	// Specifies the default time-to-live.
	// If set, this overrides the global default.
	// Must be a valid duration string
	DefaultLeaseTTL *string `json:"defaultLeaseTtl,omitempty" tf:"default_lease_ttl"`

	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	ListingVisibility *string `json:"listingVisibility,omitempty" tf:"listing_visibility"`

	// Specifies the maximum time-to-live.
	// If set, this overrides the global default.
	// Must be a valid duration string
	MaxLeaseTTL *string `json:"maxLeaseTtl,omitempty" tf:"max_lease_ttl"`

	// List of headers to whitelist and
	// pass from the request to the backend.
	PassthroughRequestHeaders []*string `json:"passthroughRequestHeaders,omitempty" tf:"passthrough_request_headers"`

	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType *string `json:"tokenType,omitempty" tf:"token_type"`
}

type TuneObservation struct {

	// List of headers to whitelist and allowing
	// a plugin to include them in the response.
	AllowedResponseHeaders []*string `json:"allowedResponseHeaders,omitempty" tf:"allowed_response_headers,omitempty"`

	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the request data object.
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys,omitempty"`

	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the response data object.
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys,omitempty"`

	// Specifies the default time-to-live.
	// If set, this overrides the global default.
	// Must be a valid duration string
	DefaultLeaseTTL *string `json:"defaultLeaseTtl,omitempty" tf:"default_lease_ttl,omitempty"`

	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	ListingVisibility *string `json:"listingVisibility,omitempty" tf:"listing_visibility,omitempty"`

	// Specifies the maximum time-to-live.
	// If set, this overrides the global default.
	// Must be a valid duration string
	MaxLeaseTTL *string `json:"maxLeaseTtl,omitempty" tf:"max_lease_ttl,omitempty"`

	// List of headers to whitelist and
	// pass from the request to the backend.
	PassthroughRequestHeaders []*string `json:"passthroughRequestHeaders,omitempty" tf:"passthrough_request_headers,omitempty"`

	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type TuneParameters struct {

	// List of headers to whitelist and allowing
	// a plugin to include them in the response.
	// +kubebuilder:validation:Optional
	AllowedResponseHeaders []*string `json:"allowedResponseHeaders,omitempty" tf:"allowed_response_headers"`

	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the request data object.
	// +kubebuilder:validation:Optional
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys"`

	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the response data object.
	// +kubebuilder:validation:Optional
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys"`

	// Specifies the default time-to-live.
	// If set, this overrides the global default.
	// Must be a valid duration string
	// +kubebuilder:validation:Optional
	DefaultLeaseTTL *string `json:"defaultLeaseTtl,omitempty" tf:"default_lease_ttl"`

	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	// +kubebuilder:validation:Optional
	ListingVisibility *string `json:"listingVisibility,omitempty" tf:"listing_visibility"`

	// Specifies the maximum time-to-live.
	// If set, this overrides the global default.
	// Must be a valid duration string
	// +kubebuilder:validation:Optional
	MaxLeaseTTL *string `json:"maxLeaseTtl,omitempty" tf:"max_lease_ttl"`

	// List of headers to whitelist and
	// pass from the request to the backend.
	// +kubebuilder:validation:Optional
	PassthroughRequestHeaders []*string `json:"passthroughRequestHeaders,omitempty" tf:"passthrough_request_headers"`

	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type"`
}

// AuthBackendSpec defines the desired state of AuthBackend
type AuthBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AuthBackendInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendStatus defines the observed state of AuthBackend.
type AuthBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackend is the Schema for the AuthBackends API. Manages GitHub Auth mounts in Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organization) || (has(self.initProvider) && has(self.initProvider.organization))",message="spec.forProvider.organization is a required parameter"
	Spec   AuthBackendSpec   `json:"spec"`
	Status AuthBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendList contains a list of AuthBackends
type AuthBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackend `json:"items"`
}

// Repository type metadata.
var (
	AuthBackend_Kind             = "AuthBackend"
	AuthBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackend_Kind}.String()
	AuthBackend_KindAPIVersion   = AuthBackend_Kind + "." + CRDGroupVersion.String()
	AuthBackend_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackend{}, &AuthBackendList{})
}
