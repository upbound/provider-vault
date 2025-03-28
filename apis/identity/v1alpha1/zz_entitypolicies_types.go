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

type EntityPoliciesInitParameters struct {

	// Entity ID to assign policies to.
	// ID of the entity.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/identity/v1alpha1.Entity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// Reference to a Entity in identity to populate entityId.
	// +kubebuilder:validation:Optional
	EntityIDRef *v1.Reference `json:"entityIdRef,omitempty" tf:"-"`

	// Selector for a Entity in identity to populate entityId.
	// +kubebuilder:validation:Optional
	EntityIDSelector *v1.Selector `json:"entityIdSelector,omitempty" tf:"-"`

	// Defaults to true.
	// Should the resource manage policies exclusively
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of policies to assign to the entity
	// Policies to be tied to the entity.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type EntityPoliciesObservation struct {

	// Entity ID to assign policies to.
	// ID of the entity.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The name of the entity that are assigned the policies.
	// Name of the entity.
	EntityName *string `json:"entityName,omitempty" tf:"entity_name,omitempty"`

	// Defaults to true.
	// Should the resource manage policies exclusively
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of policies to assign to the entity
	// Policies to be tied to the entity.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type EntityPoliciesParameters struct {

	// Entity ID to assign policies to.
	// ID of the entity.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/identity/v1alpha1.Entity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// Reference to a Entity in identity to populate entityId.
	// +kubebuilder:validation:Optional
	EntityIDRef *v1.Reference `json:"entityIdRef,omitempty" tf:"-"`

	// Selector for a Entity in identity to populate entityId.
	// +kubebuilder:validation:Optional
	EntityIDSelector *v1.Selector `json:"entityIdSelector,omitempty" tf:"-"`

	// Defaults to true.
	// Should the resource manage policies exclusively
	// +kubebuilder:validation:Optional
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of policies to assign to the entity
	// Policies to be tied to the entity.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// EntityPoliciesSpec defines the desired state of EntityPolicies
type EntityPoliciesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntityPoliciesParameters `json:"forProvider"`
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
	InitProvider EntityPoliciesInitParameters `json:"initProvider,omitempty"`
}

// EntityPoliciesStatus defines the observed state of EntityPolicies.
type EntityPoliciesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntityPoliciesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EntityPolicies is the Schema for the EntityPoliciess API. Manages policies for an Identity Entity for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type EntityPolicies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policies) || (has(self.initProvider) && has(self.initProvider.policies))",message="spec.forProvider.policies is a required parameter"
	Spec   EntityPoliciesSpec   `json:"spec"`
	Status EntityPoliciesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntityPoliciesList contains a list of EntityPoliciess
type EntityPoliciesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EntityPolicies `json:"items"`
}

// Repository type metadata.
var (
	EntityPolicies_Kind             = "EntityPolicies"
	EntityPolicies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EntityPolicies_Kind}.String()
	EntityPolicies_KindAPIVersion   = EntityPolicies_Kind + "." + CRDGroupVersion.String()
	EntityPolicies_GroupVersionKind = CRDGroupVersion.WithKind(EntityPolicies_Kind)
)

func init() {
	SchemeBuilder.Register(&EntityPolicies{}, &EntityPoliciesList{})
}
