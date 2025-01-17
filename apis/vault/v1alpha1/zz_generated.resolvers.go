/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VaultNamespace.
func (mg *VaultNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Namespace),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.NamespaceRef,
		Selector:     mg.Spec.ForProvider.NamespaceSelector,
		To: reference.To{
			List:    &VaultNamespaceList{},
			Managed: &VaultNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Namespace")
	}
	mg.Spec.ForProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Namespace),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.NamespaceRef,
		Selector:     mg.Spec.InitProvider.NamespaceSelector,
		To: reference.To{
			List:    &VaultNamespaceList{},
			Managed: &VaultNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Namespace")
	}
	mg.Spec.InitProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NamespaceRef = rsp.ResolvedReference

	return nil
}
