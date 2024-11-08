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

type AuthBackendUserInitParameters struct {

	// Groups within the Okta auth backend to associate with this user
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path to the Okta auth backend
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Policies to associate with this user
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Name of the user within Okta
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthBackendUserObservation struct {

	// Groups within the Okta auth backend to associate with this user
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path to the Okta auth backend
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Policies to associate with this user
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Name of the user within Okta
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthBackendUserParameters struct {

	// Groups within the Okta auth backend to associate with this user
	// +kubebuilder:validation:Optional
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path to the Okta auth backend
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Policies to associate with this user
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Name of the user within Okta
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// AuthBackendUserSpec defines the desired state of AuthBackendUser
type AuthBackendUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendUserParameters `json:"forProvider"`
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
	InitProvider AuthBackendUserInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendUserStatus defines the observed state of AuthBackendUser.
type AuthBackendUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendUser is the Schema for the AuthBackendUsers API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   AuthBackendUserSpec   `json:"spec"`
	Status AuthBackendUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendUserList contains a list of AuthBackendUsers
type AuthBackendUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendUser `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendUser_Kind             = "AuthBackendUser"
	AuthBackendUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendUser_Kind}.String()
	AuthBackendUser_KindAPIVersion   = AuthBackendUser_Kind + "." + CRDGroupVersion.String()
	AuthBackendUser_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendUser_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendUser{}, &AuthBackendUserList{})
}
