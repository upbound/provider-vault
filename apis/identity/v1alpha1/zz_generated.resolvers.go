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
	v1alpha1 "github.com/upbound/provider-vault/apis/auth/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EntityPolicies.
func (mg *EntityPolicies) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EntityID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.EntityIDRef,
		Selector:     mg.Spec.ForProvider.EntityIDSelector,
		To: reference.To{
			List:    &EntityList{},
			Managed: &Entity{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EntityID")
	}
	mg.Spec.ForProvider.EntityID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EntityIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EntityID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.EntityIDRef,
		Selector:     mg.Spec.InitProvider.EntityIDSelector,
		To: reference.To{
			List:    &EntityList{},
			Managed: &Entity{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EntityID")
	}
	mg.Spec.InitProvider.EntityID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EntityIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupAlias.
func (mg *GroupAlias) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CanonicalID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CanonicalIDRef,
		Selector:     mg.Spec.ForProvider.CanonicalIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CanonicalID")
	}
	mg.Spec.ForProvider.CanonicalID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CanonicalIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.ForProvider.MountAccessorRef,
		Selector:     mg.Spec.ForProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MountAccessor")
	}
	mg.Spec.ForProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MountAccessorRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CanonicalID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.CanonicalIDRef,
		Selector:     mg.Spec.InitProvider.CanonicalIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CanonicalID")
	}
	mg.Spec.InitProvider.CanonicalID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CanonicalIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.InitProvider.MountAccessorRef,
		Selector:     mg.Spec.InitProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MountAccessor")
	}
	mg.Spec.InitProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MountAccessorRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupMemberEntityIds.
func (mg *GroupMemberEntityIds) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.GroupIDRef,
		Selector:     mg.Spec.InitProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupID")
	}
	mg.Spec.InitProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupMemberGroupIds.
func (mg *GroupMemberGroupIds) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.GroupIDRef,
		Selector:     mg.Spec.InitProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupID")
	}
	mg.Spec.InitProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupPolicies.
func (mg *GroupPolicies) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.GroupIDRef,
		Selector:     mg.Spec.InitProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupID")
	}
	mg.Spec.InitProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OidcKeyAllowedClientID.
func (mg *OidcKeyAllowedClientID) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AllowedClientID),
		Extract:      resource.ExtractParamPath("client_id", false),
		Reference:    mg.Spec.ForProvider.AllowedClientIDRef,
		Selector:     mg.Spec.ForProvider.AllowedClientIDSelector,
		To: reference.To{
			List:    &OidcRoleList{},
			Managed: &OidcRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AllowedClientID")
	}
	mg.Spec.ForProvider.AllowedClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AllowedClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.KeyNameRef,
		Selector:     mg.Spec.ForProvider.KeyNameSelector,
		To: reference.To{
			List:    &OidcKeyList{},
			Managed: &OidcKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyName")
	}
	mg.Spec.ForProvider.KeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AllowedClientID),
		Extract:      resource.ExtractParamPath("client_id", false),
		Reference:    mg.Spec.InitProvider.AllowedClientIDRef,
		Selector:     mg.Spec.InitProvider.AllowedClientIDSelector,
		To: reference.To{
			List:    &OidcRoleList{},
			Managed: &OidcRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AllowedClientID")
	}
	mg.Spec.InitProvider.AllowedClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AllowedClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.KeyNameRef,
		Selector:     mg.Spec.InitProvider.KeyNameSelector,
		To: reference.To{
			List:    &OidcKeyList{},
			Managed: &OidcKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyName")
	}
	mg.Spec.InitProvider.KeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OidcRole.
func (mg *OidcRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Key),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.KeyRef,
		Selector:     mg.Spec.ForProvider.KeySelector,
		To: reference.To{
			List:    &OidcKeyList{},
			Managed: &OidcKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Key")
	}
	mg.Spec.ForProvider.Key = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Key),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.KeyRef,
		Selector:     mg.Spec.InitProvider.KeySelector,
		To: reference.To{
			List:    &OidcKeyList{},
			Managed: &OidcKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Key")
	}
	mg.Spec.InitProvider.Key = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyRef = rsp.ResolvedReference

	return nil
}