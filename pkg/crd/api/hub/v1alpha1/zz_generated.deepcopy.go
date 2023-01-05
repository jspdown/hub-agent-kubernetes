//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOIDC) DeepCopyInto(out *AccessControlOIDC) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.AuthParams != nil {
		in, out := &in.AuthParams, &out.AuthParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StateCookie != nil {
		in, out := &in.StateCookie, &out.StateCookie
		*out = new(StateCookie)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(Session)
		(*in).DeepCopyInto(*out)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOIDC.
func (in *AccessControlOIDC) DeepCopy() *AccessControlOIDC {
	if in == nil {
		return nil
	}
	out := new(AccessControlOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOIDCGoogle) DeepCopyInto(out *AccessControlOIDCGoogle) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.AuthParams != nil {
		in, out := &in.AuthParams, &out.AuthParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StateCookie != nil {
		in, out := &in.StateCookie, &out.StateCookie
		*out = new(StateCookie)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(Session)
		(*in).DeepCopyInto(*out)
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOIDCGoogle.
func (in *AccessControlOIDCGoogle) DeepCopy() *AccessControlOIDCGoogle {
	if in == nil {
		return nil
	}
	out := new(AccessControlOIDCGoogle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicy) DeepCopyInto(out *AccessControlPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicy.
func (in *AccessControlPolicy) DeepCopy() *AccessControlPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessControlPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyBasicAuth) DeepCopyInto(out *AccessControlPolicyBasicAuth) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyBasicAuth.
func (in *AccessControlPolicyBasicAuth) DeepCopy() *AccessControlPolicyBasicAuth {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyBasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyJWT) DeepCopyInto(out *AccessControlPolicyJWT) {
	*out = *in
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyJWT.
func (in *AccessControlPolicyJWT) DeepCopy() *AccessControlPolicyJWT {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyJWT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyList) DeepCopyInto(out *AccessControlPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessControlPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyList.
func (in *AccessControlPolicyList) DeepCopy() *AccessControlPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessControlPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicySpec) DeepCopyInto(out *AccessControlPolicySpec) {
	*out = *in
	if in.JWT != nil {
		in, out := &in.JWT, &out.JWT
		*out = new(AccessControlPolicyJWT)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(AccessControlPolicyBasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(AccessControlOIDC)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDCGoogle != nil {
		in, out := &in.OIDCGoogle, &out.OIDCGoogle
		*out = new(AccessControlOIDCGoogle)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicySpec.
func (in *AccessControlPolicySpec) DeepCopy() *AccessControlPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyStatus) DeepCopyInto(out *AccessControlPolicyStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyStatus.
func (in *AccessControlPolicyStatus) DeepCopy() *AccessControlPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Catalog) DeepCopyInto(out *Catalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Catalog.
func (in *Catalog) DeepCopy() *Catalog {
	if in == nil {
		return nil
	}
	out := new(Catalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Catalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogList) DeepCopyInto(out *CatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Catalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogList.
func (in *CatalogList) DeepCopy() *CatalogList {
	if in == nil {
		return nil
	}
	out := new(CatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogService) DeepCopyInto(out *CatalogService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogService.
func (in *CatalogService) DeepCopy() *CatalogService {
	if in == nil {
		return nil
	}
	out := new(CatalogService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpec) DeepCopyInto(out *CatalogSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]CatalogService, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpec.
func (in *CatalogSpec) DeepCopy() *CatalogSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogStatus) DeepCopyInto(out *CatalogStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogStatus.
func (in *CatalogStatus) DeepCopy() *CatalogStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngress) DeepCopyInto(out *EdgeIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngress.
func (in *EdgeIngress) DeepCopy() *EdgeIngress {
	if in == nil {
		return nil
	}
	out := new(EdgeIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressACP) DeepCopyInto(out *EdgeIngressACP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressACP.
func (in *EdgeIngressACP) DeepCopy() *EdgeIngressACP {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressACP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressList) DeepCopyInto(out *EdgeIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressList.
func (in *EdgeIngressList) DeepCopy() *EdgeIngressList {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressService) DeepCopyInto(out *EdgeIngressService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressService.
func (in *EdgeIngressService) DeepCopy() *EdgeIngressService {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressSpec) DeepCopyInto(out *EdgeIngressSpec) {
	*out = *in
	out.Service = in.Service
	if in.ACP != nil {
		in, out := &in.ACP, &out.ACP
		*out = new(EdgeIngressACP)
		**out = **in
	}
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressSpec.
func (in *EdgeIngressSpec) DeepCopy() *EdgeIngressSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressStatus) DeepCopyInto(out *EdgeIngressStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressStatus.
func (in *EdgeIngressStatus) DeepCopy() *EdgeIngressStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClass) DeepCopyInto(out *IngressClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClass.
func (in *IngressClass) DeepCopy() *IngressClass {
	if in == nil {
		return nil
	}
	out := new(IngressClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassList) DeepCopyInto(out *IngressClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassList.
func (in *IngressClassList) DeepCopy() *IngressClassList {
	if in == nil {
		return nil
	}
	out := new(IngressClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassSpec) DeepCopyInto(out *IngressClassSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassSpec.
func (in *IngressClassSpec) DeepCopy() *IngressClassSpec {
	if in == nil {
		return nil
	}
	out := new(IngressClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Session) DeepCopyInto(out *Session) {
	*out = *in
	if in.Refresh != nil {
		in, out := &in.Refresh, &out.Refresh
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Session.
func (in *Session) DeepCopy() *Session {
	if in == nil {
		return nil
	}
	out := new(Session)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateCookie) DeepCopyInto(out *StateCookie) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateCookie.
func (in *StateCookie) DeepCopy() *StateCookie {
	if in == nil {
		return nil
	}
	out := new(StateCookie)
	in.DeepCopyInto(out)
	return out
}
