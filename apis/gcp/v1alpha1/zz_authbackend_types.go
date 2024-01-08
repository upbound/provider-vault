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

type AuthBackendInitParameters struct {
	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email,omitempty"`

	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Specifies overrides to service endpoints used when making API requests to GCP.
	CustomEndpoint []CustomEndpointInitParameters `json:"customEndpoint,omitempty" tf:"custom_endpoint,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// Specifies if the auth method is local only
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	PrivateKeyID *string `json:"privateKeyId,omitempty" tf:"private_key_id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type AuthBackendObservation struct {

	// The accessor of the auth backend
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email,omitempty"`

	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Specifies overrides to service endpoints used when making API requests to GCP.
	CustomEndpoint []CustomEndpointObservation `json:"customEndpoint,omitempty" tf:"custom_endpoint,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies if the auth method is local only
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	PrivateKeyID *string `json:"privateKeyId,omitempty" tf:"private_key_id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type AuthBackendParameters struct {

	// +kubebuilder:validation:Optional
	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email,omitempty"`

	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	CredentialsSecretRef *v1.SecretKeySelector `json:"credentialsSecretRef,omitempty" tf:"-"`

	// Specifies overrides to service endpoints used when making API requests to GCP.
	// +kubebuilder:validation:Optional
	CustomEndpoint []CustomEndpointParameters `json:"customEndpoint,omitempty" tf:"custom_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// Specifies if the auth method is local only
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateKeyID *string `json:"privateKeyId,omitempty" tf:"private_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type CustomEndpointInitParameters struct {

	// Replaces the service endpoint used in API requests to https://www.googleapis.com.
	API *string `json:"api,omitempty" tf:"api,omitempty"`

	// Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
	Compute *string `json:"compute,omitempty" tf:"compute,omitempty"`

	// Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
	Crm *string `json:"crm,omitempty" tf:"crm,omitempty"`

	// Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
	IAM *string `json:"iam,omitempty" tf:"iam,omitempty"`
}

type CustomEndpointObservation struct {

	// Replaces the service endpoint used in API requests to https://www.googleapis.com.
	API *string `json:"api,omitempty" tf:"api,omitempty"`

	// Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
	Compute *string `json:"compute,omitempty" tf:"compute,omitempty"`

	// Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
	Crm *string `json:"crm,omitempty" tf:"crm,omitempty"`

	// Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
	IAM *string `json:"iam,omitempty" tf:"iam,omitempty"`
}

type CustomEndpointParameters struct {

	// Replaces the service endpoint used in API requests to https://www.googleapis.com.
	// +kubebuilder:validation:Optional
	API *string `json:"api,omitempty" tf:"api,omitempty"`

	// Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
	// +kubebuilder:validation:Optional
	Compute *string `json:"compute,omitempty" tf:"compute,omitempty"`

	// Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
	// +kubebuilder:validation:Optional
	Crm *string `json:"crm,omitempty" tf:"crm,omitempty"`

	// Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
	// +kubebuilder:validation:Optional
	IAM *string `json:"iam,omitempty" tf:"iam,omitempty"`
}

// AuthBackendSpec defines the desired state of AuthBackend
type AuthBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendParameters `json:"forProvider"`
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
	InitProvider AuthBackendInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendStatus defines the observed state of AuthBackend.
type AuthBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackend is the Schema for the AuthBackends API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendSpec   `json:"spec"`
	Status            AuthBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendList contains a list of AuthBackends
type AuthBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackend `json:"items"`
}

// Repository type metadata.
var (
	AuthBackend_Kind             = "AuthBackend"
	AuthBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackend_Kind}.String()
	AuthBackend_KindAPIVersion   = AuthBackend_Kind + "." + CRDGroupVersion.String()
	AuthBackend_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackend{}, &AuthBackendList{})
}
