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

type SecretBackendSignInitParameters struct {

	// List of alternative names
	// List of alternative names.
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

	// If set to true, certs will be renewed if the expiration is within min_seconds_remaining. Default false
	// If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// CN of certificate to create
	// CN of intermediate to create.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The CSR
	// The CSR.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// Flag to exclude CN from SANs
	// Flag to exclude CN from SANs.
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The format of data
	// The format of data.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// List of alternative IPs
	// List of alternative IPs.
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// Specifies the default issuer of this request. Can
	// be the value default, a name, or an issuer ID. Use ACLs to prevent access to
	// the /pki/issuer/:issuer_ref/{issue,sign}/:name paths to prevent users
	// overriding the role's issuer_ref value.
	// Specifies the default issuer of this request.
	IssuerRef *string `json:"issuerRef,omitempty" tf:"issuer_ref,omitempty"`

	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	// Generate a new certificate when the expiration is within this number of seconds
	MinSecondsRemaining *float64 `json:"minSecondsRemaining,omitempty" tf:"min_seconds_remaining,omitempty"`

	// Name of the role to create the certificate against
	// Name of the role to create the certificate against.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of other SANs
	// List of other SANs.
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// Time to live
	// Time to live.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// List of alternative URIs
	// List of alternative URIs.
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

type SecretBackendSignObservation struct {

	// List of alternative names
	// List of alternative names.
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

	// If set to true, certs will be renewed if the expiration is within min_seconds_remaining. Default false
	// If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The CA chain
	// The CA chain.
	CAChain []*string `json:"caChain,omitempty" tf:"ca_chain,omitempty"`

	// The certificate
	// The certicate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// CN of certificate to create
	// CN of intermediate to create.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The CSR
	// The CSR.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// Flag to exclude CN from SANs
	// Flag to exclude CN from SANs.
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The expiration date of the certificate in unix epoch format
	// The certificate expiration as a Unix-style timestamp.
	Expiration *float64 `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The format of data
	// The format of data.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of alternative IPs
	// List of alternative IPs.
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// Specifies the default issuer of this request. Can
	// be the value default, a name, or an issuer ID. Use ACLs to prevent access to
	// the /pki/issuer/:issuer_ref/{issue,sign}/:name paths to prevent users
	// overriding the role's issuer_ref value.
	// Specifies the default issuer of this request.
	IssuerRef *string `json:"issuerRef,omitempty" tf:"issuer_ref,omitempty"`

	// The issuing CA
	// The issuing CA.
	IssuingCA *string `json:"issuingCa,omitempty" tf:"issuing_ca,omitempty"`

	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	// Generate a new certificate when the expiration is within this number of seconds
	MinSecondsRemaining *float64 `json:"minSecondsRemaining,omitempty" tf:"min_seconds_remaining,omitempty"`

	// Name of the role to create the certificate against
	// Name of the role to create the certificate against.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of other SANs
	// List of other SANs.
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// true if the current time (during refresh) is after the start of the early renewal window declared by min_seconds_remaining, and false otherwise; if auto_renew is set to true then the provider will plan to replace the certificate once renewal is pending.
	// Initially false, and then set to true during refresh once the expiration is less than min_seconds_remaining in the future.
	RenewPending *bool `json:"renewPending,omitempty" tf:"renew_pending,omitempty"`

	// The certificate's serial number, hex formatted.
	// The certificate's serial number, hex formatted.
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// Time to live
	// Time to live.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// List of alternative URIs
	// List of alternative URIs.
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

type SecretBackendSignParameters struct {

	// List of alternative names
	// List of alternative names.
	// +kubebuilder:validation:Optional
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

	// If set to true, certs will be renewed if the expiration is within min_seconds_remaining. Default false
	// If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining
	// +kubebuilder:validation:Optional
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// CN of certificate to create
	// CN of intermediate to create.
	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The CSR
	// The CSR.
	// +kubebuilder:validation:Optional
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// Flag to exclude CN from SANs
	// Flag to exclude CN from SANs.
	// +kubebuilder:validation:Optional
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The format of data
	// The format of data.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// List of alternative IPs
	// List of alternative IPs.
	// +kubebuilder:validation:Optional
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// Specifies the default issuer of this request. Can
	// be the value default, a name, or an issuer ID. Use ACLs to prevent access to
	// the /pki/issuer/:issuer_ref/{issue,sign}/:name paths to prevent users
	// overriding the role's issuer_ref value.
	// Specifies the default issuer of this request.
	// +kubebuilder:validation:Optional
	IssuerRef *string `json:"issuerRef,omitempty" tf:"issuer_ref,omitempty"`

	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	// Generate a new certificate when the expiration is within this number of seconds
	// +kubebuilder:validation:Optional
	MinSecondsRemaining *float64 `json:"minSecondsRemaining,omitempty" tf:"min_seconds_remaining,omitempty"`

	// Name of the role to create the certificate against
	// Name of the role to create the certificate against.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of other SANs
	// List of other SANs.
	// +kubebuilder:validation:Optional
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// Time to live
	// Time to live.
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// List of alternative URIs
	// List of alternative URIs.
	// +kubebuilder:validation:Optional
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

// SecretBackendSignSpec defines the desired state of SecretBackendSign
type SecretBackendSignSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendSignParameters `json:"forProvider"`
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
	InitProvider SecretBackendSignInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendSignStatus defines the observed state of SecretBackendSign.
type SecretBackendSignStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendSignObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendSign is the Schema for the SecretBackendSigns API. Sign a new certificate based on the CSR by the PKI.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendSign struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backend) || has(self.initProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.commonName) || has(self.initProvider.commonName)",message="commonName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.csr) || has(self.initProvider.csr)",message="csr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SecretBackendSignSpec   `json:"spec"`
	Status SecretBackendSignStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendSignList contains a list of SecretBackendSigns
type SecretBackendSignList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendSign `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendSign_Kind             = "SecretBackendSign"
	SecretBackendSign_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendSign_Kind}.String()
	SecretBackendSign_KindAPIVersion   = SecretBackendSign_Kind + "." + CRDGroupVersion.String()
	SecretBackendSign_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendSign_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendSign{}, &SecretBackendSignList{})
}
