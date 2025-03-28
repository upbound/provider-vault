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

type TeamInitParameters struct {

	// Path where the github auth backend is mounted. Defaults to github
	// if not specified.
	// Auth backend to which team mapping will be configured.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/github/v1alpha1.AuthBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a AuthBackend in github to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a AuthBackend in github to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	// Policies to be assigned to this team.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// GitHub team name in "slugified" format.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type TeamObservation struct {

	// Path where the github auth backend is mounted. Defaults to github
	// if not specified.
	// Auth backend to which team mapping will be configured.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	// Policies to be assigned to this team.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// GitHub team name in "slugified" format.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type TeamParameters struct {

	// Path where the github auth backend is mounted. Defaults to github
	// if not specified.
	// Auth backend to which team mapping will be configured.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/github/v1alpha1.AuthBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a AuthBackend in github to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a AuthBackend in github to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	// Policies to be assigned to this team.
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// GitHub team name in "slugified" format.
	// +kubebuilder:validation:Optional
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

// TeamSpec defines the desired state of Team
type TeamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamParameters `json:"forProvider"`
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
	InitProvider TeamInitParameters `json:"initProvider,omitempty"`
}

// TeamStatus defines the observed state of Team.
type TeamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Team is the Schema for the Teams API. Manages Team mappings for Github Auth backend mounts in Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.team) || (has(self.initProvider) && has(self.initProvider.team))",message="spec.forProvider.team is a required parameter"
	Spec   TeamSpec   `json:"spec"`
	Status TeamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamList contains a list of Teams
type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

// Repository type metadata.
var (
	Team_Kind             = "Team"
	Team_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Team_Kind}.String()
	Team_KindAPIVersion   = Team_Kind + "." + CRDGroupVersion.String()
	Team_GroupVersionKind = CRDGroupVersion.WithKind(Team_Kind)
)

func init() {
	SchemeBuilder.Register(&Team{}, &TeamList{})
}
