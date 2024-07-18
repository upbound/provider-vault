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

type CloudSecretBackendInitParameters struct {

	// The default is
	// https://app.0.0.1:8500".
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The unique location this backend should be mounted at. Must not begin or end with a /
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	BasePath *string `json:"basePath,omitempty" tf:"base_path,omitempty"`

	// The default TTL for credentials issued by this backend.
	// Default lease duration for secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	// Maximum possible lease duration for secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type CloudSecretBackendObservation struct {

	// The default is
	// https://app.0.0.1:8500".
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The unique location this backend should be mounted at. Must not begin or end with a /
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	BasePath *string `json:"basePath,omitempty" tf:"base_path,omitempty"`

	// The default TTL for credentials issued by this backend.
	// Default lease duration for secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	// Maximum possible lease duration for secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type CloudSecretBackendParameters struct {

	// The default is
	// https://app.0.0.1:8500".
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The unique location this backend should be mounted at. Must not begin or end with a /
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// +kubebuilder:validation:Optional
	BasePath *string `json:"basePath,omitempty" tf:"base_path,omitempty"`

	// The default TTL for credentials issued by this backend.
	// Default lease duration for secrets in seconds
	// +kubebuilder:validation:Optional
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	// Maximum possible lease duration for secrets in seconds
	// +kubebuilder:validation:Optional
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

// CloudSecretBackendSpec defines the desired state of CloudSecretBackend
type CloudSecretBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudSecretBackendParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CloudSecretBackendInitParameters `json:"initProvider,omitempty"`
}

// CloudSecretBackendStatus defines the observed state of CloudSecretBackend.
type CloudSecretBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudSecretBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudSecretBackend is the Schema for the CloudSecretBackends API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type CloudSecretBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudSecretBackendSpec   `json:"spec"`
	Status            CloudSecretBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudSecretBackendList contains a list of CloudSecretBackends
type CloudSecretBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSecretBackend `json:"items"`
}

// Repository type metadata.
var (
	CloudSecretBackend_Kind             = "CloudSecretBackend"
	CloudSecretBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudSecretBackend_Kind}.String()
	CloudSecretBackend_KindAPIVersion   = CloudSecretBackend_Kind + "." + CRDGroupVersion.String()
	CloudSecretBackend_GroupVersionKind = CRDGroupVersion.WithKind(CloudSecretBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudSecretBackend{}, &CloudSecretBackendList{})
}
