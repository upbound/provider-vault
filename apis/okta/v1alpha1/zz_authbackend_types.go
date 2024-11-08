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

	// The Okta url. Examples: oktapreview.com, okta.com (default)
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa *bool `json:"bypassOktaMfa,omitempty" tf:"bypass_okta_mfa,omitempty"`

	// The description of the auth backend
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	Group []GroupInitParameters `json:"group,omitempty" tf:"group,omitempty"`

	// Maximum duration after which authentication will be expired
	MaxTTL *string `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The Okta organization. This will be the first part of the url https://XXX.okta.com.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// path to mount the backend
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Duration after which authentication will be expired
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The Okta API token. This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled.
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`

	User []UserInitParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type AuthBackendObservation struct {

	// The mount accessor related to the auth mount.
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	// The Okta url. Examples: oktapreview.com, okta.com (default)
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa *bool `json:"bypassOktaMfa,omitempty" tf:"bypass_okta_mfa,omitempty"`

	// The description of the auth backend
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	Group []GroupObservation `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum duration after which authentication will be expired
	MaxTTL *string `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The Okta organization. This will be the first part of the url https://XXX.okta.com.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// path to mount the backend
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Duration after which authentication will be expired
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`

	User []UserObservation `json:"user,omitempty" tf:"user,omitempty"`
}

type AuthBackendParameters struct {

	// The Okta url. Examples: oktapreview.com, okta.com (default)
	// +kubebuilder:validation:Optional
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	// +kubebuilder:validation:Optional
	BypassOktaMfa *bool `json:"bypassOktaMfa,omitempty" tf:"bypass_okta_mfa,omitempty"`

	// The description of the auth backend
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// +kubebuilder:validation:Optional
	Group []GroupParameters `json:"group,omitempty" tf:"group,omitempty"`

	// Maximum duration after which authentication will be expired
	// +kubebuilder:validation:Optional
	MaxTTL *string `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The Okta organization. This will be the first part of the url https://XXX.okta.com.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// path to mount the backend
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Duration after which authentication will be expired
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The Okta API token. This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled.
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`

	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`

	// +kubebuilder:validation:Optional
	User []UserParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type GroupInitParameters struct {
	GroupName *string `json:"groupName,omitempty" tf:"group_name"`

	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies"`
}

type GroupObservation struct {
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type GroupParameters struct {

	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name"`

	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies"`
}

type UserInitParameters struct {

	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups"`

	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies"`

	Username *string `json:"username,omitempty" tf:"username"`
}

type UserObservation struct {

	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// +kubebuilder:validation:Optional
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups"`

	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username"`
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

// AuthBackend is the Schema for the AuthBackends API. <no value>
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
