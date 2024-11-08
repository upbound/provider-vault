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

type GroupMemberGroupIdsInitParameters struct {

	// Defaults to true.
	// If set to true, allows the resource to manage member group ids
	// exclusively. Beware of race conditions when disabling exclusive management
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// Group ID to assign member entities to.
	// ID of the group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/identity/v1alpha1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// List of member groups that belong to the group
	// Group IDs to be assigned as group members.
	// +listType=set
	MemberGroupIds []*string `json:"memberGroupIds,omitempty" tf:"member_group_ids,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type GroupMemberGroupIdsObservation struct {

	// Defaults to true.
	// If set to true, allows the resource to manage member group ids
	// exclusively. Beware of race conditions when disabling exclusive management
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// Group ID to assign member entities to.
	// ID of the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of member groups that belong to the group
	// Group IDs to be assigned as group members.
	// +listType=set
	MemberGroupIds []*string `json:"memberGroupIds,omitempty" tf:"member_group_ids,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type GroupMemberGroupIdsParameters struct {

	// Defaults to true.
	// If set to true, allows the resource to manage member group ids
	// exclusively. Beware of race conditions when disabling exclusive management
	// +kubebuilder:validation:Optional
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// Group ID to assign member entities to.
	// ID of the group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/identity/v1alpha1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// List of member groups that belong to the group
	// Group IDs to be assigned as group members.
	// +kubebuilder:validation:Optional
	// +listType=set
	MemberGroupIds []*string `json:"memberGroupIds,omitempty" tf:"member_group_ids,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// GroupMemberGroupIdsSpec defines the desired state of GroupMemberGroupIds
type GroupMemberGroupIdsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMemberGroupIdsParameters `json:"forProvider"`
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
	InitProvider GroupMemberGroupIdsInitParameters `json:"initProvider,omitempty"`
}

// GroupMemberGroupIdsStatus defines the observed state of GroupMemberGroupIds.
type GroupMemberGroupIdsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMemberGroupIdsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GroupMemberGroupIds is the Schema for the GroupMemberGroupIdss API. Manages member groups for an Identity Group for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type GroupMemberGroupIds struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupMemberGroupIdsSpec   `json:"spec"`
	Status            GroupMemberGroupIdsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMemberGroupIdsList contains a list of GroupMemberGroupIdss
type GroupMemberGroupIdsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMemberGroupIds `json:"items"`
}

// Repository type metadata.
var (
	GroupMemberGroupIds_Kind             = "GroupMemberGroupIds"
	GroupMemberGroupIds_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupMemberGroupIds_Kind}.String()
	GroupMemberGroupIds_KindAPIVersion   = GroupMemberGroupIds_Kind + "." + CRDGroupVersion.String()
	GroupMemberGroupIds_GroupVersionKind = CRDGroupVersion.WithKind(GroupMemberGroupIds_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupMemberGroupIds{}, &GroupMemberGroupIdsList{})
}
