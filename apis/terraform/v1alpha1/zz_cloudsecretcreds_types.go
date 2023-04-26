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

type CloudSecretCredsObservation struct {

	// Upbound official provider cloud secret backend to generate tokens from
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// name of the Upbound official provider cloud or enterprise organization
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Name of the role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// g., settings/teams/team-xxxxxxxxxxxxx)
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// id of the Upbound official provider token provided
	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`
}

type CloudSecretCredsParameters struct {

	// Upbound official provider cloud secret backend to generate tokens from
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// CloudSecretCredsSpec defines the desired state of CloudSecretCreds
type CloudSecretCredsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudSecretCredsParameters `json:"forProvider"`
}

// CloudSecretCredsStatus defines the observed state of CloudSecretCreds.
type CloudSecretCredsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudSecretCredsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudSecretCreds is the Schema for the CloudSecretCredss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type CloudSecretCreds struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.role)",message="role is a required parameter"
	Spec   CloudSecretCredsSpec   `json:"spec"`
	Status CloudSecretCredsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudSecretCredsList contains a list of CloudSecretCredss
type CloudSecretCredsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSecretCreds `json:"items"`
}

// Repository type metadata.
var (
	CloudSecretCreds_Kind             = "CloudSecretCreds"
	CloudSecretCreds_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudSecretCreds_Kind}.String()
	CloudSecretCreds_KindAPIVersion   = CloudSecretCreds_Kind + "." + CRDGroupVersion.String()
	CloudSecretCreds_GroupVersionKind = CRDGroupVersion.WithKind(CloudSecretCreds_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudSecretCreds{}, &CloudSecretCredsList{})
}