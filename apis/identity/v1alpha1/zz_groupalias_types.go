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

type GroupAliasInitParameters struct {

	// ID of the group to which this is an alias.
	// ID of the group to which this is an alias.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/identity/v1alpha1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	CanonicalID *string `json:"canonicalId,omitempty" tf:"canonical_id,omitempty"`

	// Reference to a Group in identity to populate canonicalId.
	// +kubebuilder:validation:Optional
	CanonicalIDRef *v1.Reference `json:"canonicalIdRef,omitempty" tf:"-"`

	// Selector for a Group in identity to populate canonicalId.
	// +kubebuilder:validation:Optional
	CanonicalIDSelector *v1.Selector `json:"canonicalIdSelector,omitempty" tf:"-"`

	// Mount accessor of the authentication backend to which this alias belongs to.
	// Mount accessor to which this alias belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/auth/v1alpha1.Backend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("accessor",true)
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Reference to a Backend in auth to populate mountAccessor.
	// +kubebuilder:validation:Optional
	MountAccessorRef *v1.Reference `json:"mountAccessorRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate mountAccessor.
	// +kubebuilder:validation:Optional
	MountAccessorSelector *v1.Selector `json:"mountAccessorSelector,omitempty" tf:"-"`

	// Name of the group alias to create.
	// Name of the group alias.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type GroupAliasObservation struct {

	// ID of the group to which this is an alias.
	// ID of the group to which this is an alias.
	CanonicalID *string `json:"canonicalId,omitempty" tf:"canonical_id,omitempty"`

	// The id of the created group alias.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Mount accessor of the authentication backend to which this alias belongs to.
	// Mount accessor to which this alias belongs to.
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Name of the group alias to create.
	// Name of the group alias.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type GroupAliasParameters struct {

	// ID of the group to which this is an alias.
	// ID of the group to which this is an alias.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/identity/v1alpha1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CanonicalID *string `json:"canonicalId,omitempty" tf:"canonical_id,omitempty"`

	// Reference to a Group in identity to populate canonicalId.
	// +kubebuilder:validation:Optional
	CanonicalIDRef *v1.Reference `json:"canonicalIdRef,omitempty" tf:"-"`

	// Selector for a Group in identity to populate canonicalId.
	// +kubebuilder:validation:Optional
	CanonicalIDSelector *v1.Selector `json:"canonicalIdSelector,omitempty" tf:"-"`

	// Mount accessor of the authentication backend to which this alias belongs to.
	// Mount accessor to which this alias belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/auth/v1alpha1.Backend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("accessor",true)
	// +kubebuilder:validation:Optional
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Reference to a Backend in auth to populate mountAccessor.
	// +kubebuilder:validation:Optional
	MountAccessorRef *v1.Reference `json:"mountAccessorRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate mountAccessor.
	// +kubebuilder:validation:Optional
	MountAccessorSelector *v1.Selector `json:"mountAccessorSelector,omitempty" tf:"-"`

	// Name of the group alias to create.
	// Name of the group alias.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// GroupAliasSpec defines the desired state of GroupAlias
type GroupAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupAliasParameters `json:"forProvider"`
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
	InitProvider GroupAliasInitParameters `json:"initProvider,omitempty"`
}

// GroupAliasStatus defines the observed state of GroupAlias.
type GroupAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GroupAlias is the Schema for the GroupAliass API. Creates an Identity Group Alias for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type GroupAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GroupAliasSpec   `json:"spec"`
	Status GroupAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupAliasList contains a list of GroupAliass
type GroupAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupAlias `json:"items"`
}

// Repository type metadata.
var (
	GroupAlias_Kind             = "GroupAlias"
	GroupAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupAlias_Kind}.String()
	GroupAlias_KindAPIVersion   = GroupAlias_Kind + "." + CRDGroupVersion.String()
	GroupAlias_GroupVersionKind = CRDGroupVersion.WithKind(GroupAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupAlias{}, &GroupAliasList{})
}
