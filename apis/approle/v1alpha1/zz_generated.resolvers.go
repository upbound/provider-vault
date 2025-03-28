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
	v1alpha1 "github.com/upbound/provider-vault/v2/apis/auth/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AuthBackendLogin.
func (mg *AuthBackendLogin) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      resource.ExtractParamPath("role_id", false),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &AuthBackendRoleList{},
			Managed: &AuthBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleID),
		Extract:      resource.ExtractParamPath("role_id", false),
		Reference:    mg.Spec.InitProvider.RoleIDRef,
		Selector:     mg.Spec.InitProvider.RoleIDSelector,
		To: reference.To{
			List:    &AuthBackendRoleList{},
			Managed: &AuthBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleID")
	}
	mg.Spec.InitProvider.RoleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AuthBackendRole.
func (mg *AuthBackendRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AuthBackendRoleSecretID.
func (mg *AuthBackendRoleSecretID) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleName),
		Extract:      resource.ExtractParamPath("role_name", false),
		Reference:    mg.Spec.ForProvider.RoleNameRef,
		Selector:     mg.Spec.ForProvider.RoleNameSelector,
		To: reference.To{
			List:    &AuthBackendRoleList{},
			Managed: &AuthBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleName")
	}
	mg.Spec.ForProvider.RoleName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleName),
		Extract:      resource.ExtractParamPath("role_name", false),
		Reference:    mg.Spec.InitProvider.RoleNameRef,
		Selector:     mg.Spec.InitProvider.RoleNameSelector,
		To: reference.To{
			List:    &AuthBackendRoleList{},
			Managed: &AuthBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleName")
	}
	mg.Spec.InitProvider.RoleName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleNameRef = rsp.ResolvedReference

	return nil
}
