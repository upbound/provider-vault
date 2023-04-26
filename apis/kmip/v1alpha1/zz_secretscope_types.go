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

type SecretScopeObservation struct {

	// Force deletion even if there are managed objects in the scope
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path where KMIP backend is mounted
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Name of the scope
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type SecretScopeParameters struct {

	// Force deletion even if there are managed objects in the scope
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path where KMIP backend is mounted
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Name of the scope
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// SecretScopeSpec defines the desired state of SecretScope
type SecretScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretScopeParameters `json:"forProvider"`
}

// SecretScopeStatus defines the observed state of SecretScope.
type SecretScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretScope is the Schema for the SecretScopes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.path)",message="path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scope)",message="scope is a required parameter"
	Spec   SecretScopeSpec   `json:"spec"`
	Status SecretScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretScopeList contains a list of SecretScopes
type SecretScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretScope `json:"items"`
}

// Repository type metadata.
var (
	SecretScope_Kind             = "SecretScope"
	SecretScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretScope_Kind}.String()
	SecretScope_KindAPIVersion   = SecretScope_Kind + "." + CRDGroupVersion.String()
	SecretScope_GroupVersionKind = CRDGroupVersion.WithKind(SecretScope_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretScope{}, &SecretScopeList{})
}