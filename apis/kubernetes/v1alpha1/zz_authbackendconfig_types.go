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

type AuthBackendConfigInitParameters struct {

	// Unique name of the kubernetes backend to configure.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/auth/v1alpha1.Backend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault v1.5.4+ or Vault auth kubernetes plugin v0.7.1+
	// Optional disable JWT issuer validation. Allows to skip ISS validation.
	DisableIssValidation *bool `json:"disableIssValidation,omitempty" tf:"disable_iss_validation,omitempty"`

	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault v1.5.4+ or Vault auth kubernetes plugin v0.7.1+
	// Optional disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod.
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	// JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	// Optional JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost *string `json:"kubernetesHost,omitempty" tf:"kubernetes_host,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []*string `json:"pemKeys,omitempty" tf:"pem_keys,omitempty"`

	// A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	// A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwtSecretRef *v1.SecretKeySelector `json:"tokenReviewerJwtSecretRef,omitempty" tf:"-"`
}

type AuthBackendConfigObservation struct {

	// Unique name of the kubernetes backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault v1.5.4+ or Vault auth kubernetes plugin v0.7.1+
	// Optional disable JWT issuer validation. Allows to skip ISS validation.
	DisableIssValidation *bool `json:"disableIssValidation,omitempty" tf:"disable_iss_validation,omitempty"`

	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault v1.5.4+ or Vault auth kubernetes plugin v0.7.1+
	// Optional disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod.
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	// Optional JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost *string `json:"kubernetesHost,omitempty" tf:"kubernetes_host,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []*string `json:"pemKeys,omitempty" tf:"pem_keys,omitempty"`
}

type AuthBackendConfigParameters struct {

	// Unique name of the kubernetes backend to configure.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/auth/v1alpha1.Backend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault v1.5.4+ or Vault auth kubernetes plugin v0.7.1+
	// Optional disable JWT issuer validation. Allows to skip ISS validation.
	// +kubebuilder:validation:Optional
	DisableIssValidation *bool `json:"disableIssValidation,omitempty" tf:"disable_iss_validation,omitempty"`

	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault v1.5.4+ or Vault auth kubernetes plugin v0.7.1+
	// Optional disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod.
	// +kubebuilder:validation:Optional
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	// JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	// Optional JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	// +kubebuilder:validation:Optional
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	// +kubebuilder:validation:Optional
	KubernetesHost *string `json:"kubernetesHost,omitempty" tf:"kubernetes_host,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	// +kubebuilder:validation:Optional
	PemKeys []*string `json:"pemKeys,omitempty" tf:"pem_keys,omitempty"`

	// A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	// A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	// +kubebuilder:validation:Optional
	TokenReviewerJwtSecretRef *v1.SecretKeySelector `json:"tokenReviewerJwtSecretRef,omitempty" tf:"-"`
}

// AuthBackendConfigSpec defines the desired state of AuthBackendConfig
type AuthBackendConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendConfigParameters `json:"forProvider"`
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
	InitProvider AuthBackendConfigInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendConfigStatus defines the observed state of AuthBackendConfig.
type AuthBackendConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendConfig is the Schema for the AuthBackendConfigs API. Manages Kubernetes auth backend configs in Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kubernetesHost) || (has(self.initProvider) && has(self.initProvider.kubernetesHost))",message="spec.forProvider.kubernetesHost is a required parameter"
	Spec   AuthBackendConfigSpec   `json:"spec"`
	Status AuthBackendConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendConfigList contains a list of AuthBackendConfigs
type AuthBackendConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendConfig `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendConfig_Kind             = "AuthBackendConfig"
	AuthBackendConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendConfig_Kind}.String()
	AuthBackendConfig_KindAPIVersion   = AuthBackendConfig_Kind + "." + CRDGroupVersion.String()
	AuthBackendConfig_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendConfig{}, &AuthBackendConfigList{})
}
