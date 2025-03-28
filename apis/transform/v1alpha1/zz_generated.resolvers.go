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
	v1alpha1 "github.com/upbound/provider-vault/apis/vault/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Alphabet.
func (mg *Alphabet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.PathRef,
		Selector:     mg.Spec.ForProvider.PathSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Path")
	}
	mg.Spec.ForProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PathRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.PathRef,
		Selector:     mg.Spec.InitProvider.PathSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Path")
	}
	mg.Spec.InitProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PathRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Role.
func (mg *Role) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.PathRef,
		Selector:     mg.Spec.ForProvider.PathSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Path")
	}
	mg.Spec.ForProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PathRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.PathRef,
		Selector:     mg.Spec.InitProvider.PathSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Path")
	}
	mg.Spec.InitProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PathRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Template.
func (mg *Template) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.PathRef,
		Selector:     mg.Spec.ForProvider.PathSelector,
		To: reference.To{
			List:    &AlphabetList{},
			Managed: &Alphabet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Path")
	}
	mg.Spec.ForProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PathRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.PathRef,
		Selector:     mg.Spec.InitProvider.PathSelector,
		To: reference.To{
			List:    &AlphabetList{},
			Managed: &Alphabet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Path")
	}
	mg.Spec.InitProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PathRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Transformation.
func (mg *Transformation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.PathRef,
		Selector:     mg.Spec.ForProvider.PathSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Path")
	}
	mg.Spec.ForProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PathRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Path),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.PathRef,
		Selector:     mg.Spec.InitProvider.PathSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Path")
	}
	mg.Spec.InitProvider.Path = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PathRef = rsp.ResolvedReference

	return nil
}
