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

type SecretBackendInitParameters struct {

	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Denotes that the resource is used to bootstrap the Consul ACL system.
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap *bool `json:"bootstrap,omitempty" tf:"bootstrap,omitempty"`

	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CACert *string `json:"caCert,omitempty" tf:"ca_cert,omitempty"`

	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCertSecretRef *v1.SecretKeySelector `json:"clientCertSecretRef,omitempty" tf:"-"`

	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKeySecretRef *v1.SecretKeySelector `json:"clientKeySecretRef,omitempty" tf:"-"`

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

	// Specifies if the secret backend is local only.
	// Specifies if the secret backend is local only
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

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

	// The unique location this backend should be mounted at. Must not begin or end with a /. Defaults
	// to consul.
	// Unique name of the Vault Consul mount to configure
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the URL scheme to use. Defaults to http.
	// Specifies the URL scheme to use. Defaults to "http".
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// The Consul management token this backend should use to issue new tokens. This field is required
	// when bootstrap is false.
	// Specifies the Consul token to use when managing or issuing new tokens.
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type SecretBackendObservation struct {

	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Denotes that the resource is used to bootstrap the Consul ACL system.
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap *bool `json:"bootstrap,omitempty" tf:"bootstrap,omitempty"`

	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CACert *string `json:"caCert,omitempty" tf:"ca_cert,omitempty"`

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

	// Specifies if the secret backend is local only.
	// Specifies if the secret backend is local only
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

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

	// The unique location this backend should be mounted at. Must not begin or end with a /. Defaults
	// to consul.
	// Unique name of the Vault Consul mount to configure
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the URL scheme to use. Defaults to http.
	// Specifies the URL scheme to use. Defaults to "http".
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type SecretBackendParameters struct {

	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Denotes that the resource is used to bootstrap the Consul ACL system.
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	// +kubebuilder:validation:Optional
	Bootstrap *bool `json:"bootstrap,omitempty" tf:"bootstrap,omitempty"`

	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	// +kubebuilder:validation:Optional
	CACert *string `json:"caCert,omitempty" tf:"ca_cert,omitempty"`

	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	// +kubebuilder:validation:Optional
	ClientCertSecretRef *v1.SecretKeySelector `json:"clientCertSecretRef,omitempty" tf:"-"`

	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	// +kubebuilder:validation:Optional
	ClientKeySecretRef *v1.SecretKeySelector `json:"clientKeySecretRef,omitempty" tf:"-"`

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

	// Specifies if the secret backend is local only.
	// Specifies if the secret backend is local only
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

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

	// The unique location this backend should be mounted at. Must not begin or end with a /. Defaults
	// to consul.
	// Unique name of the Vault Consul mount to configure
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the URL scheme to use. Defaults to http.
	// Specifies the URL scheme to use. Defaults to "http".
	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// The Consul management token this backend should use to issue new tokens. This field is required
	// when bootstrap is false.
	// Specifies the Consul token to use when managing or issuing new tokens.
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

// SecretBackendSpec defines the desired state of SecretBackend
type SecretBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendParameters `json:"forProvider"`
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
	InitProvider SecretBackendInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendStatus defines the observed state of SecretBackend.
type SecretBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackend is the Schema for the SecretBackends API. Creates a Consul secret backend for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.address) || (has(self.initProvider) && has(self.initProvider.address))",message="spec.forProvider.address is a required parameter"
	Spec   SecretBackendSpec   `json:"spec"`
	Status SecretBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendList contains a list of SecretBackends
type SecretBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackend `json:"items"`
}

// Repository type metadata.
var (
	SecretBackend_Kind             = "SecretBackend"
	SecretBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackend_Kind}.String()
	SecretBackend_KindAPIVersion   = SecretBackend_Kind + "." + CRDGroupVersion.String()
	SecretBackend_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackend{}, &SecretBackendList{})
}
