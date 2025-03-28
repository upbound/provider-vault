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

type AuthBackendStsRoleInitParameters struct {

	// The AWS account ID to configure the STS role for.
	// AWS account ID to be associated with STS role.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to aws.
	// Unique name of the auth backend to configure.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/auth/v1alpha1.Backend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by account_id.
	// AWS ARN for STS role to be assumed when interacting with the account specified.
	StsRole *string `json:"stsRole,omitempty" tf:"sts_role,omitempty"`
}

type AuthBackendStsRoleObservation struct {

	// The AWS account ID to configure the STS role for.
	// AWS account ID to be associated with STS role.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to aws.
	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by account_id.
	// AWS ARN for STS role to be assumed when interacting with the account specified.
	StsRole *string `json:"stsRole,omitempty" tf:"sts_role,omitempty"`
}

type AuthBackendStsRoleParameters struct {

	// The AWS account ID to configure the STS role for.
	// AWS account ID to be associated with STS role.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to aws.
	// Unique name of the auth backend to configure.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/auth/v1alpha1.Backend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by account_id.
	// AWS ARN for STS role to be assumed when interacting with the account specified.
	// +kubebuilder:validation:Optional
	StsRole *string `json:"stsRole,omitempty" tf:"sts_role,omitempty"`
}

// AuthBackendStsRoleSpec defines the desired state of AuthBackendStsRole
type AuthBackendStsRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendStsRoleParameters `json:"forProvider"`
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
	InitProvider AuthBackendStsRoleInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendStsRoleStatus defines the observed state of AuthBackendStsRole.
type AuthBackendStsRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendStsRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendStsRole is the Schema for the AuthBackendStsRoles API. Configures an STS role in the Vault AWS Auth backend.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendStsRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stsRole) || (has(self.initProvider) && has(self.initProvider.stsRole))",message="spec.forProvider.stsRole is a required parameter"
	Spec   AuthBackendStsRoleSpec   `json:"spec"`
	Status AuthBackendStsRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendStsRoleList contains a list of AuthBackendStsRoles
type AuthBackendStsRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendStsRole `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendStsRole_Kind             = "AuthBackendStsRole"
	AuthBackendStsRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendStsRole_Kind}.String()
	AuthBackendStsRole_KindAPIVersion   = AuthBackendStsRole_Kind + "." + CRDGroupVersion.String()
	AuthBackendStsRole_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendStsRole_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendStsRole{}, &AuthBackendStsRoleList{})
}
