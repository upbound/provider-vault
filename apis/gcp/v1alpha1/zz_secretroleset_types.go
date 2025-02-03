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

type BindingInitParameters struct {

	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different formats.
	// Resource name
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// List of GCP IAM roles for the resource.
	// List of roles to apply to the resource
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type BindingObservation struct {

	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different formats.
	// Resource name
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// List of GCP IAM roles for the resource.
	// List of roles to apply to the resource
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type BindingParameters struct {

	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different formats.
	// Resource name
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource" tf:"resource,omitempty"`

	// List of GCP IAM roles for the resource.
	// List of roles to apply to the resource
	// +kubebuilder:validation:Optional
	// +listType=set
	Roles []*string `json:"roles" tf:"roles,omitempty"`
}

type SecretRolesetInitParameters struct {

	// Path where the GCP Secrets Engine is mounted
	// Path where the GCP secrets engine is mounted.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/gcp/v1alpha1.SecretBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a SecretBackend in gcp to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a SecretBackend in gcp to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Binding []BindingInitParameters `json:"binding,omitempty" tf:"binding,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the GCP project that this roleset's service account will belong to.
	// Name of the GCP project that this roleset's service account will belong to.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the Roleset to create
	// Name of the RoleSet to create
	Roleset *string `json:"roleset,omitempty" tf:"roleset,omitempty"`

	// Type of secret generated for this role set. Accepted values: access_token, service_account_key. Defaults to access_token.
	// Type of secret generated for this role set. Defaults to `access_token`. Accepted values: `access_token`, `service_account_key`
	SecretType *string `json:"secretType,omitempty" tf:"secret_type,omitempty"`

	// List of OAuth scopes to assign to access_token secrets generated under this role set (access_token role sets only).
	// List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only)
	// +listType=set
	TokenScopes []*string `json:"tokenScopes,omitempty" tf:"token_scopes,omitempty"`
}

type SecretRolesetObservation struct {

	// Path where the GCP Secrets Engine is mounted
	// Path where the GCP secrets engine is mounted.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Binding []BindingObservation `json:"binding,omitempty" tf:"binding,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the GCP project that this roleset's service account will belong to.
	// Name of the GCP project that this roleset's service account will belong to.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the Roleset to create
	// Name of the RoleSet to create
	Roleset *string `json:"roleset,omitempty" tf:"roleset,omitempty"`

	// Type of secret generated for this role set. Accepted values: access_token, service_account_key. Defaults to access_token.
	// Type of secret generated for this role set. Defaults to `access_token`. Accepted values: `access_token`, `service_account_key`
	SecretType *string `json:"secretType,omitempty" tf:"secret_type,omitempty"`

	// Email of the service account created by Vault for this Roleset.
	// Email of the service account created by Vault for this Roleset
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// List of OAuth scopes to assign to access_token secrets generated under this role set (access_token role sets only).
	// List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only)
	// +listType=set
	TokenScopes []*string `json:"tokenScopes,omitempty" tf:"token_scopes,omitempty"`
}

type SecretRolesetParameters struct {

	// Path where the GCP Secrets Engine is mounted
	// Path where the GCP secrets engine is mounted.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/gcp/v1alpha1.SecretBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a SecretBackend in gcp to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a SecretBackend in gcp to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	// +kubebuilder:validation:Optional
	Binding []BindingParameters `json:"binding,omitempty" tf:"binding,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the GCP project that this roleset's service account will belong to.
	// Name of the GCP project that this roleset's service account will belong to.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the Roleset to create
	// Name of the RoleSet to create
	// +kubebuilder:validation:Optional
	Roleset *string `json:"roleset,omitempty" tf:"roleset,omitempty"`

	// Type of secret generated for this role set. Accepted values: access_token, service_account_key. Defaults to access_token.
	// Type of secret generated for this role set. Defaults to `access_token`. Accepted values: `access_token`, `service_account_key`
	// +kubebuilder:validation:Optional
	SecretType *string `json:"secretType,omitempty" tf:"secret_type,omitempty"`

	// List of OAuth scopes to assign to access_token secrets generated under this role set (access_token role sets only).
	// List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only)
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenScopes []*string `json:"tokenScopes,omitempty" tf:"token_scopes,omitempty"`
}

// SecretRolesetSpec defines the desired state of SecretRoleset
type SecretRolesetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretRolesetParameters `json:"forProvider"`
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
	InitProvider SecretRolesetInitParameters `json:"initProvider,omitempty"`
}

// SecretRolesetStatus defines the observed state of SecretRoleset.
type SecretRolesetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretRolesetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretRoleset is the Schema for the SecretRolesets API. Creates a Roleset for the GCP Secret Backend for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretRoleset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.binding) || (has(self.initProvider) && has(self.initProvider.binding))",message="spec.forProvider.binding is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.project) || (has(self.initProvider) && has(self.initProvider.project))",message="spec.forProvider.project is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleset) || (has(self.initProvider) && has(self.initProvider.roleset))",message="spec.forProvider.roleset is a required parameter"
	Spec   SecretRolesetSpec   `json:"spec"`
	Status SecretRolesetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRolesetList contains a list of SecretRolesets
type SecretRolesetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretRoleset `json:"items"`
}

// Repository type metadata.
var (
	SecretRoleset_Kind             = "SecretRoleset"
	SecretRoleset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretRoleset_Kind}.String()
	SecretRoleset_KindAPIVersion   = SecretRoleset_Kind + "." + CRDGroupVersion.String()
	SecretRoleset_GroupVersionKind = CRDGroupVersion.WithKind(SecretRoleset_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretRoleset{}, &SecretRolesetList{})
}
