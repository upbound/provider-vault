/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-vault/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SecretBackendIntermediateSetSigned.
func (mg *SecretBackendIntermediateSetSigned) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Certificate),
		Extract:      common.ExtractCrt(),
		Reference:    mg.Spec.ForProvider.CertificateRef,
		Selector:     mg.Spec.ForProvider.CertificateSelector,
		To: reference.To{
			List:    &SecretBackendRootSignIntermediateList{},
			Managed: &SecretBackendRootSignIntermediate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Certificate")
	}
	mg.Spec.ForProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendRootSignIntermediate.
func (mg *SecretBackendRootSignIntermediate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Csr),
		Extract:      common.ExtractCsr(),
		Reference:    mg.Spec.ForProvider.CsrRef,
		Selector:     mg.Spec.ForProvider.CsrSelector,
		To: reference.To{
			List:    &SecretBackendIntermediateCertRequestList{},
			Managed: &SecretBackendIntermediateCertRequest{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Csr")
	}
	mg.Spec.ForProvider.Csr = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CsrRef = rsp.ResolvedReference

	return nil
}