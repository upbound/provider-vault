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

type AuthBackendStsRoleObservation struct {

	// AWS account ID to be associated with STS role.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// AWS ARN for STS role to be assumed when interacting with the account specified.
	StsRole *string `json:"stsRole,omitempty" tf:"sts_role,omitempty"`
}

type AuthBackendStsRoleParameters struct {

	// AWS account ID to be associated with STS role.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Unique name of the auth backend to configure.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// AWS ARN for STS role to be assumed when interacting with the account specified.
	// +kubebuilder:validation:Optional
	StsRole *string `json:"stsRole,omitempty" tf:"sts_role,omitempty"`
}

// AuthBackendStsRoleSpec defines the desired state of AuthBackendStsRole
type AuthBackendStsRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendStsRoleParameters `json:"forProvider"`
}

// AuthBackendStsRoleStatus defines the observed state of AuthBackendStsRole.
type AuthBackendStsRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendStsRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendStsRole is the Schema for the AuthBackendStsRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendStsRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.accountId)",message="accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.stsRole)",message="stsRole is a required parameter"
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