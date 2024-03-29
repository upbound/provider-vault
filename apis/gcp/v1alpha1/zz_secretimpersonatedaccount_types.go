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

type SecretImpersonatedAccountInitParameters struct {

	// Path where the GCP Secrets Engine is mounted
	// Path where the GCP secrets engine is mounted.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Name of the Impersonated Account to create
	// Name of the Impersonated Account to create
	ImpersonatedAccount *string `json:"impersonatedAccount,omitempty" tf:"impersonated_account,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Email of the GCP service account to impersonate.
	// Email of the GCP service account.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	// List of OAuth scopes to assign to `access_token` secrets generated under this impersonated account (`access_token` impersonated accounts only)
	TokenScopes []*string `json:"tokenScopes,omitempty" tf:"token_scopes,omitempty"`
}

type SecretImpersonatedAccountObservation struct {

	// Path where the GCP Secrets Engine is mounted
	// Path where the GCP secrets engine is mounted.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Impersonated Account to create
	// Name of the Impersonated Account to create
	ImpersonatedAccount *string `json:"impersonatedAccount,omitempty" tf:"impersonated_account,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Email of the GCP service account to impersonate.
	// Email of the GCP service account.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Project the service account belongs to.
	// Project of the GCP Service Account managed by this impersonated account
	ServiceAccountProject *string `json:"serviceAccountProject,omitempty" tf:"service_account_project,omitempty"`

	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	// List of OAuth scopes to assign to `access_token` secrets generated under this impersonated account (`access_token` impersonated accounts only)
	TokenScopes []*string `json:"tokenScopes,omitempty" tf:"token_scopes,omitempty"`
}

type SecretImpersonatedAccountParameters struct {

	// Path where the GCP Secrets Engine is mounted
	// Path where the GCP secrets engine is mounted.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Name of the Impersonated Account to create
	// Name of the Impersonated Account to create
	// +kubebuilder:validation:Optional
	ImpersonatedAccount *string `json:"impersonatedAccount,omitempty" tf:"impersonated_account,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Email of the GCP service account to impersonate.
	// Email of the GCP service account.
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	// List of OAuth scopes to assign to `access_token` secrets generated under this impersonated account (`access_token` impersonated accounts only)
	// +kubebuilder:validation:Optional
	TokenScopes []*string `json:"tokenScopes,omitempty" tf:"token_scopes,omitempty"`
}

// SecretImpersonatedAccountSpec defines the desired state of SecretImpersonatedAccount
type SecretImpersonatedAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretImpersonatedAccountParameters `json:"forProvider"`
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
	InitProvider SecretImpersonatedAccountInitParameters `json:"initProvider,omitempty"`
}

// SecretImpersonatedAccountStatus defines the observed state of SecretImpersonatedAccount.
type SecretImpersonatedAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretImpersonatedAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretImpersonatedAccount is the Schema for the SecretImpersonatedAccounts API. Creates a Impersonated Account for the GCP Secret Backend for Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretImpersonatedAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backend) || has(self.initProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.impersonatedAccount) || has(self.initProvider.impersonatedAccount)",message="impersonatedAccount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceAccountEmail) || has(self.initProvider.serviceAccountEmail)",message="serviceAccountEmail is a required parameter"
	Spec   SecretImpersonatedAccountSpec   `json:"spec"`
	Status SecretImpersonatedAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretImpersonatedAccountList contains a list of SecretImpersonatedAccounts
type SecretImpersonatedAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretImpersonatedAccount `json:"items"`
}

// Repository type metadata.
var (
	SecretImpersonatedAccount_Kind             = "SecretImpersonatedAccount"
	SecretImpersonatedAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretImpersonatedAccount_Kind}.String()
	SecretImpersonatedAccount_KindAPIVersion   = SecretImpersonatedAccount_Kind + "." + CRDGroupVersion.String()
	SecretImpersonatedAccount_GroupVersionKind = CRDGroupVersion.WithKind(SecretImpersonatedAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretImpersonatedAccount{}, &SecretImpersonatedAccountList{})
}
