/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/upbound/provider-vault/apis/ad/v1alpha1"
	v1alpha1alicloud "github.com/upbound/provider-vault/apis/alicloud/v1alpha1"
	v1alpha1approle "github.com/upbound/provider-vault/apis/approle/v1alpha1"
	v1alpha1audit "github.com/upbound/provider-vault/apis/audit/v1alpha1"
	v1alpha1auth "github.com/upbound/provider-vault/apis/auth/v1alpha1"
	v1alpha1aws "github.com/upbound/provider-vault/apis/aws/v1alpha1"
	v1alpha1azure "github.com/upbound/provider-vault/apis/azure/v1alpha1"
	v1alpha1cert "github.com/upbound/provider-vault/apis/cert/v1alpha1"
	v1alpha1config "github.com/upbound/provider-vault/apis/config/v1alpha1"
	v1alpha1consul "github.com/upbound/provider-vault/apis/consul/v1alpha1"
	v1alpha1database "github.com/upbound/provider-vault/apis/database/v1alpha1"
	v1alpha1egp "github.com/upbound/provider-vault/apis/egp/v1alpha1"
	v1alpha1gcp "github.com/upbound/provider-vault/apis/gcp/v1alpha1"
	v1alpha1generic "github.com/upbound/provider-vault/apis/generic/v1alpha1"
	v1alpha1github "github.com/upbound/provider-vault/apis/github/v1alpha1"
	v1alpha1identity "github.com/upbound/provider-vault/apis/identity/v1alpha1"
	v1alpha1jwt "github.com/upbound/provider-vault/apis/jwt/v1alpha1"
	v1alpha1kmip "github.com/upbound/provider-vault/apis/kmip/v1alpha1"
	v1alpha1kubernetes "github.com/upbound/provider-vault/apis/kubernetes/v1alpha1"
	v1alpha1kv "github.com/upbound/provider-vault/apis/kv/v1alpha1"
	v1alpha1ldap "github.com/upbound/provider-vault/apis/ldap/v1alpha1"
	v1alpha1managed "github.com/upbound/provider-vault/apis/managed/v1alpha1"
	v1alpha1mfa "github.com/upbound/provider-vault/apis/mfa/v1alpha1"
	v1alpha1mongodbatlas "github.com/upbound/provider-vault/apis/mongodbatlas/v1alpha1"
	v1alpha1nomad "github.com/upbound/provider-vault/apis/nomad/v1alpha1"
	v1alpha1okta "github.com/upbound/provider-vault/apis/okta/v1alpha1"
	v1alpha1password "github.com/upbound/provider-vault/apis/password/v1alpha1"
	v1alpha1pki "github.com/upbound/provider-vault/apis/pki/v1alpha1"
	v1alpha1quota "github.com/upbound/provider-vault/apis/quota/v1alpha1"
	v1alpha1rabbitmq "github.com/upbound/provider-vault/apis/rabbitmq/v1alpha1"
	v1alpha1raft "github.com/upbound/provider-vault/apis/raft/v1alpha1"
	v1alpha1rgp "github.com/upbound/provider-vault/apis/rgp/v1alpha1"
	v1alpha1ssh "github.com/upbound/provider-vault/apis/ssh/v1alpha1"
	v1alpha1terraform "github.com/upbound/provider-vault/apis/terraform/v1alpha1"
	v1alpha1token "github.com/upbound/provider-vault/apis/token/v1alpha1"
	v1alpha1transform "github.com/upbound/provider-vault/apis/transform/v1alpha1"
	v1alpha1transit "github.com/upbound/provider-vault/apis/transit/v1alpha1"
	v1alpha1apis "github.com/upbound/provider-vault/apis/v1alpha1"
	v1beta1 "github.com/upbound/provider-vault/apis/v1beta1"
	v1alpha1vault "github.com/upbound/provider-vault/apis/vault/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1alicloud.SchemeBuilder.AddToScheme,
		v1alpha1approle.SchemeBuilder.AddToScheme,
		v1alpha1audit.SchemeBuilder.AddToScheme,
		v1alpha1auth.SchemeBuilder.AddToScheme,
		v1alpha1aws.SchemeBuilder.AddToScheme,
		v1alpha1azure.SchemeBuilder.AddToScheme,
		v1alpha1cert.SchemeBuilder.AddToScheme,
		v1alpha1config.SchemeBuilder.AddToScheme,
		v1alpha1consul.SchemeBuilder.AddToScheme,
		v1alpha1database.SchemeBuilder.AddToScheme,
		v1alpha1egp.SchemeBuilder.AddToScheme,
		v1alpha1gcp.SchemeBuilder.AddToScheme,
		v1alpha1generic.SchemeBuilder.AddToScheme,
		v1alpha1github.SchemeBuilder.AddToScheme,
		v1alpha1identity.SchemeBuilder.AddToScheme,
		v1alpha1jwt.SchemeBuilder.AddToScheme,
		v1alpha1kmip.SchemeBuilder.AddToScheme,
		v1alpha1kubernetes.SchemeBuilder.AddToScheme,
		v1alpha1kv.SchemeBuilder.AddToScheme,
		v1alpha1ldap.SchemeBuilder.AddToScheme,
		v1alpha1managed.SchemeBuilder.AddToScheme,
		v1alpha1mfa.SchemeBuilder.AddToScheme,
		v1alpha1mongodbatlas.SchemeBuilder.AddToScheme,
		v1alpha1nomad.SchemeBuilder.AddToScheme,
		v1alpha1okta.SchemeBuilder.AddToScheme,
		v1alpha1password.SchemeBuilder.AddToScheme,
		v1alpha1pki.SchemeBuilder.AddToScheme,
		v1alpha1quota.SchemeBuilder.AddToScheme,
		v1alpha1rabbitmq.SchemeBuilder.AddToScheme,
		v1alpha1raft.SchemeBuilder.AddToScheme,
		v1alpha1rgp.SchemeBuilder.AddToScheme,
		v1alpha1ssh.SchemeBuilder.AddToScheme,
		v1alpha1terraform.SchemeBuilder.AddToScheme,
		v1alpha1token.SchemeBuilder.AddToScheme,
		v1alpha1transform.SchemeBuilder.AddToScheme,
		v1alpha1transit.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1vault.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
