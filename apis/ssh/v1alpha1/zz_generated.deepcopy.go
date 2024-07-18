//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedUserKeyConfigInitParameters) DeepCopyInto(out *AllowedUserKeyConfigInitParameters) {
	*out = *in
	if in.Lengths != nil {
		in, out := &in.Lengths, &out.Lengths
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedUserKeyConfigInitParameters.
func (in *AllowedUserKeyConfigInitParameters) DeepCopy() *AllowedUserKeyConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedUserKeyConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedUserKeyConfigObservation) DeepCopyInto(out *AllowedUserKeyConfigObservation) {
	*out = *in
	if in.Lengths != nil {
		in, out := &in.Lengths, &out.Lengths
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedUserKeyConfigObservation.
func (in *AllowedUserKeyConfigObservation) DeepCopy() *AllowedUserKeyConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AllowedUserKeyConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedUserKeyConfigParameters) DeepCopyInto(out *AllowedUserKeyConfigParameters) {
	*out = *in
	if in.Lengths != nil {
		in, out := &in.Lengths, &out.Lengths
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedUserKeyConfigParameters.
func (in *AllowedUserKeyConfigParameters) DeepCopy() *AllowedUserKeyConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedUserKeyConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCA) DeepCopyInto(out *SecretBackendCA) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCA.
func (in *SecretBackendCA) DeepCopy() *SecretBackendCA {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendCA) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCAInitParameters) DeepCopyInto(out *SecretBackendCAInitParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.GenerateSigningKey != nil {
		in, out := &in.GenerateSigningKey, &out.GenerateSigningKey
		*out = new(bool)
		**out = **in
	}
	if in.KeyBits != nil {
		in, out := &in.KeyBits, &out.KeyBits
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCAInitParameters.
func (in *SecretBackendCAInitParameters) DeepCopy() *SecretBackendCAInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCAInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCAList) DeepCopyInto(out *SecretBackendCAList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackendCA, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCAList.
func (in *SecretBackendCAList) DeepCopy() *SecretBackendCAList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCAList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendCAList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCAObservation) DeepCopyInto(out *SecretBackendCAObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.GenerateSigningKey != nil {
		in, out := &in.GenerateSigningKey, &out.GenerateSigningKey
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyBits != nil {
		in, out := &in.KeyBits, &out.KeyBits
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCAObservation.
func (in *SecretBackendCAObservation) DeepCopy() *SecretBackendCAObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCAObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCAParameters) DeepCopyInto(out *SecretBackendCAParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.GenerateSigningKey != nil {
		in, out := &in.GenerateSigningKey, &out.GenerateSigningKey
		*out = new(bool)
		**out = **in
	}
	if in.KeyBits != nil {
		in, out := &in.KeyBits, &out.KeyBits
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCAParameters.
func (in *SecretBackendCAParameters) DeepCopy() *SecretBackendCAParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCAParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCASpec) DeepCopyInto(out *SecretBackendCASpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCASpec.
func (in *SecretBackendCASpec) DeepCopy() *SecretBackendCASpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendCAStatus) DeepCopyInto(out *SecretBackendCAStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendCAStatus.
func (in *SecretBackendCAStatus) DeepCopy() *SecretBackendCAStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendCAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRole) DeepCopyInto(out *SecretBackendRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRole.
func (in *SecretBackendRole) DeepCopy() *SecretBackendRole {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleInitParameters) DeepCopyInto(out *SecretBackendRoleInitParameters) {
	*out = *in
	if in.AlgorithmSigner != nil {
		in, out := &in.AlgorithmSigner, &out.AlgorithmSigner
		*out = new(string)
		**out = **in
	}
	if in.AllowBareDomains != nil {
		in, out := &in.AllowBareDomains, &out.AllowBareDomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowHostCertificates != nil {
		in, out := &in.AllowHostCertificates, &out.AllowHostCertificates
		*out = new(bool)
		**out = **in
	}
	if in.AllowSubdomains != nil {
		in, out := &in.AllowSubdomains, &out.AllowSubdomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowUserCertificates != nil {
		in, out := &in.AllowUserCertificates, &out.AllowUserCertificates
		*out = new(bool)
		**out = **in
	}
	if in.AllowUserKeyIds != nil {
		in, out := &in.AllowUserKeyIds, &out.AllowUserKeyIds
		*out = new(bool)
		**out = **in
	}
	if in.AllowedCriticalOptions != nil {
		in, out := &in.AllowedCriticalOptions, &out.AllowedCriticalOptions
		*out = new(string)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = new(string)
		**out = **in
	}
	if in.AllowedDomainsTemplate != nil {
		in, out := &in.AllowedDomainsTemplate, &out.AllowedDomainsTemplate
		*out = new(bool)
		**out = **in
	}
	if in.AllowedExtensions != nil {
		in, out := &in.AllowedExtensions, &out.AllowedExtensions
		*out = new(string)
		**out = **in
	}
	if in.AllowedUserKeyConfig != nil {
		in, out := &in.AllowedUserKeyConfig, &out.AllowedUserKeyConfig
		*out = make([]AllowedUserKeyConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedUsers != nil {
		in, out := &in.AllowedUsers, &out.AllowedUsers
		*out = new(string)
		**out = **in
	}
	if in.AllowedUsersTemplate != nil {
		in, out := &in.AllowedUsersTemplate, &out.AllowedUsersTemplate
		*out = new(bool)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = new(string)
		**out = **in
	}
	if in.DefaultCriticalOptions != nil {
		in, out := &in.DefaultCriticalOptions, &out.DefaultCriticalOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultExtensions != nil {
		in, out := &in.DefaultExtensions, &out.DefaultExtensions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultUser != nil {
		in, out := &in.DefaultUser, &out.DefaultUser
		*out = new(string)
		**out = **in
	}
	if in.DefaultUserTemplate != nil {
		in, out := &in.DefaultUserTemplate, &out.DefaultUserTemplate
		*out = new(bool)
		**out = **in
	}
	if in.KeyIDFormat != nil {
		in, out := &in.KeyIDFormat, &out.KeyIDFormat
		*out = new(string)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.NotBeforeDuration != nil {
		in, out := &in.NotBeforeDuration, &out.NotBeforeDuration
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleInitParameters.
func (in *SecretBackendRoleInitParameters) DeepCopy() *SecretBackendRoleInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleList) DeepCopyInto(out *SecretBackendRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackendRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleList.
func (in *SecretBackendRoleList) DeepCopy() *SecretBackendRoleList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleObservation) DeepCopyInto(out *SecretBackendRoleObservation) {
	*out = *in
	if in.AlgorithmSigner != nil {
		in, out := &in.AlgorithmSigner, &out.AlgorithmSigner
		*out = new(string)
		**out = **in
	}
	if in.AllowBareDomains != nil {
		in, out := &in.AllowBareDomains, &out.AllowBareDomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowHostCertificates != nil {
		in, out := &in.AllowHostCertificates, &out.AllowHostCertificates
		*out = new(bool)
		**out = **in
	}
	if in.AllowSubdomains != nil {
		in, out := &in.AllowSubdomains, &out.AllowSubdomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowUserCertificates != nil {
		in, out := &in.AllowUserCertificates, &out.AllowUserCertificates
		*out = new(bool)
		**out = **in
	}
	if in.AllowUserKeyIds != nil {
		in, out := &in.AllowUserKeyIds, &out.AllowUserKeyIds
		*out = new(bool)
		**out = **in
	}
	if in.AllowedCriticalOptions != nil {
		in, out := &in.AllowedCriticalOptions, &out.AllowedCriticalOptions
		*out = new(string)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = new(string)
		**out = **in
	}
	if in.AllowedDomainsTemplate != nil {
		in, out := &in.AllowedDomainsTemplate, &out.AllowedDomainsTemplate
		*out = new(bool)
		**out = **in
	}
	if in.AllowedExtensions != nil {
		in, out := &in.AllowedExtensions, &out.AllowedExtensions
		*out = new(string)
		**out = **in
	}
	if in.AllowedUserKeyConfig != nil {
		in, out := &in.AllowedUserKeyConfig, &out.AllowedUserKeyConfig
		*out = make([]AllowedUserKeyConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedUsers != nil {
		in, out := &in.AllowedUsers, &out.AllowedUsers
		*out = new(string)
		**out = **in
	}
	if in.AllowedUsersTemplate != nil {
		in, out := &in.AllowedUsersTemplate, &out.AllowedUsersTemplate
		*out = new(bool)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = new(string)
		**out = **in
	}
	if in.DefaultCriticalOptions != nil {
		in, out := &in.DefaultCriticalOptions, &out.DefaultCriticalOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultExtensions != nil {
		in, out := &in.DefaultExtensions, &out.DefaultExtensions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultUser != nil {
		in, out := &in.DefaultUser, &out.DefaultUser
		*out = new(string)
		**out = **in
	}
	if in.DefaultUserTemplate != nil {
		in, out := &in.DefaultUserTemplate, &out.DefaultUserTemplate
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyIDFormat != nil {
		in, out := &in.KeyIDFormat, &out.KeyIDFormat
		*out = new(string)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.NotBeforeDuration != nil {
		in, out := &in.NotBeforeDuration, &out.NotBeforeDuration
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleObservation.
func (in *SecretBackendRoleObservation) DeepCopy() *SecretBackendRoleObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleParameters) DeepCopyInto(out *SecretBackendRoleParameters) {
	*out = *in
	if in.AlgorithmSigner != nil {
		in, out := &in.AlgorithmSigner, &out.AlgorithmSigner
		*out = new(string)
		**out = **in
	}
	if in.AllowBareDomains != nil {
		in, out := &in.AllowBareDomains, &out.AllowBareDomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowHostCertificates != nil {
		in, out := &in.AllowHostCertificates, &out.AllowHostCertificates
		*out = new(bool)
		**out = **in
	}
	if in.AllowSubdomains != nil {
		in, out := &in.AllowSubdomains, &out.AllowSubdomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowUserCertificates != nil {
		in, out := &in.AllowUserCertificates, &out.AllowUserCertificates
		*out = new(bool)
		**out = **in
	}
	if in.AllowUserKeyIds != nil {
		in, out := &in.AllowUserKeyIds, &out.AllowUserKeyIds
		*out = new(bool)
		**out = **in
	}
	if in.AllowedCriticalOptions != nil {
		in, out := &in.AllowedCriticalOptions, &out.AllowedCriticalOptions
		*out = new(string)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = new(string)
		**out = **in
	}
	if in.AllowedDomainsTemplate != nil {
		in, out := &in.AllowedDomainsTemplate, &out.AllowedDomainsTemplate
		*out = new(bool)
		**out = **in
	}
	if in.AllowedExtensions != nil {
		in, out := &in.AllowedExtensions, &out.AllowedExtensions
		*out = new(string)
		**out = **in
	}
	if in.AllowedUserKeyConfig != nil {
		in, out := &in.AllowedUserKeyConfig, &out.AllowedUserKeyConfig
		*out = make([]AllowedUserKeyConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedUsers != nil {
		in, out := &in.AllowedUsers, &out.AllowedUsers
		*out = new(string)
		**out = **in
	}
	if in.AllowedUsersTemplate != nil {
		in, out := &in.AllowedUsersTemplate, &out.AllowedUsersTemplate
		*out = new(bool)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = new(string)
		**out = **in
	}
	if in.DefaultCriticalOptions != nil {
		in, out := &in.DefaultCriticalOptions, &out.DefaultCriticalOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultExtensions != nil {
		in, out := &in.DefaultExtensions, &out.DefaultExtensions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultUser != nil {
		in, out := &in.DefaultUser, &out.DefaultUser
		*out = new(string)
		**out = **in
	}
	if in.DefaultUserTemplate != nil {
		in, out := &in.DefaultUserTemplate, &out.DefaultUserTemplate
		*out = new(bool)
		**out = **in
	}
	if in.KeyIDFormat != nil {
		in, out := &in.KeyIDFormat, &out.KeyIDFormat
		*out = new(string)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.NotBeforeDuration != nil {
		in, out := &in.NotBeforeDuration, &out.NotBeforeDuration
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleParameters.
func (in *SecretBackendRoleParameters) DeepCopy() *SecretBackendRoleParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleSpec) DeepCopyInto(out *SecretBackendRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleSpec.
func (in *SecretBackendRoleSpec) DeepCopy() *SecretBackendRoleSpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleStatus) DeepCopyInto(out *SecretBackendRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleStatus.
func (in *SecretBackendRoleStatus) DeepCopy() *SecretBackendRoleStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleStatus)
	in.DeepCopyInto(out)
	return out
}
