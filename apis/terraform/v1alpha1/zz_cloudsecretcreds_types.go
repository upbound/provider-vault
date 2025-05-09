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

type CloudSecretCredsInitParameters struct {

	// the path to the provider cloud secret backend to
	// read credentials from, with no leading or trailing /s.
	// provider cloud secret backend to generate tokens from
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/terraform/v1alpha1.CloudSecretBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("backend",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a CloudSecretBackend in terraform to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a CloudSecretBackend in terraform to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/terraform/v1alpha1.CloudSecretRole
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a CloudSecretRole in terraform to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a CloudSecretRole in terraform to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

type CloudSecretCredsObservation struct {

	// the path to the provider cloud secret backend to
	// read credentials from, with no leading or trailing /s.
	// provider cloud secret backend to generate tokens from
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization associated with the token provided
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Name of the role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The team id associated with the token provided.g., settings/teams/team-xxxxxxxxxxxxx)
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// The public identifier for a specific token. It can be used
	// to look up information about a token or to revoke a token
	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`
}

type CloudSecretCredsParameters struct {

	// the path to the provider cloud secret backend to
	// read credentials from, with no leading or trailing /s.
	// provider cloud secret backend to generate tokens from
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/terraform/v1alpha1.CloudSecretBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("backend",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a CloudSecretBackend in terraform to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a CloudSecretBackend in terraform to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/terraform/v1alpha1.CloudSecretRole
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a CloudSecretRole in terraform to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a CloudSecretRole in terraform to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

// CloudSecretCredsSpec defines the desired state of CloudSecretCreds
type CloudSecretCredsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudSecretCredsParameters `json:"forProvider"`
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
	InitProvider CloudSecretCredsInitParameters `json:"initProvider,omitempty"`
}

// CloudSecretCredsStatus defines the observed state of CloudSecretCreds.
type CloudSecretCredsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudSecretCredsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CloudSecretCreds is the Schema for the CloudSecretCredss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type CloudSecretCreds struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudSecretCredsSpec   `json:"spec"`
	Status            CloudSecretCredsStatus `json:"status,omitempty"`
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
