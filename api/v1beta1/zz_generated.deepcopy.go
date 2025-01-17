//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Advertise) DeepCopyInto(out *Advertise) {
	*out = *in
	in.Allowed.DeepCopyInto(&out.Allowed)
	if in.PrefixesWithLocalPref != nil {
		in, out := &in.PrefixesWithLocalPref, &out.PrefixesWithLocalPref
		*out = make([]LocalPrefPrefixes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrefixesWithCommunity != nil {
		in, out := &in.PrefixesWithCommunity, &out.PrefixesWithCommunity
		*out = make([]CommunityPrefixes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Advertise.
func (in *Advertise) DeepCopy() *Advertise {
	if in == nil {
		return nil
	}
	out := new(Advertise)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedPrefixes) DeepCopyInto(out *AllowedPrefixes) {
	*out = *in
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedPrefixes.
func (in *AllowedPrefixes) DeepCopy() *AllowedPrefixes {
	if in == nil {
		return nil
	}
	out := new(AllowedPrefixes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BFDProfile) DeepCopyInto(out *BFDProfile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BFDProfile.
func (in *BFDProfile) DeepCopy() *BFDProfile {
	if in == nil {
		return nil
	}
	out := new(BFDProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfig) DeepCopyInto(out *BGPConfig) {
	*out = *in
	if in.Routers != nil {
		in, out := &in.Routers, &out.Routers
		*out = make([]Router, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BFDProfiles != nil {
		in, out := &in.BFDProfiles, &out.BFDProfiles
		*out = make([]BFDProfile, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfig.
func (in *BGPConfig) DeepCopy() *BGPConfig {
	if in == nil {
		return nil
	}
	out := new(BGPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommunityPrefixes) DeepCopyInto(out *CommunityPrefixes) {
	*out = *in
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommunityPrefixes.
func (in *CommunityPrefixes) DeepCopy() *CommunityPrefixes {
	if in == nil {
		return nil
	}
	out := new(CommunityPrefixes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRRConfiguration) DeepCopyInto(out *FRRConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRRConfiguration.
func (in *FRRConfiguration) DeepCopy() *FRRConfiguration {
	if in == nil {
		return nil
	}
	out := new(FRRConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FRRConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRRConfigurationList) DeepCopyInto(out *FRRConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FRRConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRRConfigurationList.
func (in *FRRConfigurationList) DeepCopy() *FRRConfigurationList {
	if in == nil {
		return nil
	}
	out := new(FRRConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FRRConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRRConfigurationSpec) DeepCopyInto(out *FRRConfigurationSpec) {
	*out = *in
	in.BGP.DeepCopyInto(&out.BGP)
	in.Raw.DeepCopyInto(&out.Raw)
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRRConfigurationSpec.
func (in *FRRConfigurationSpec) DeepCopy() *FRRConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(FRRConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRRConfigurationStatus) DeepCopyInto(out *FRRConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRRConfigurationStatus.
func (in *FRRConfigurationStatus) DeepCopy() *FRRConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(FRRConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalPrefPrefixes) DeepCopyInto(out *LocalPrefPrefixes) {
	*out = *in
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalPrefPrefixes.
func (in *LocalPrefPrefixes) DeepCopy() *LocalPrefPrefixes {
	if in == nil {
		return nil
	}
	out := new(LocalPrefPrefixes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Neighbor) DeepCopyInto(out *Neighbor) {
	*out = *in
	out.PasswordSecret = in.PasswordSecret
	out.HoldTime = in.HoldTime
	out.KeepaliveTime = in.KeepaliveTime
	in.ToAdvertise.DeepCopyInto(&out.ToAdvertise)
	in.ToReceive.DeepCopyInto(&out.ToReceive)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Neighbor.
func (in *Neighbor) DeepCopy() *Neighbor {
	if in == nil {
		return nil
	}
	out := new(Neighbor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawConfig) DeepCopyInto(out *RawConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawConfig.
func (in *RawConfig) DeepCopy() *RawConfig {
	if in == nil {
		return nil
	}
	out := new(RawConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Receive) DeepCopyInto(out *Receive) {
	*out = *in
	in.Allowed.DeepCopyInto(&out.Allowed)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Receive.
func (in *Receive) DeepCopy() *Receive {
	if in == nil {
		return nil
	}
	out := new(Receive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Router) DeepCopyInto(out *Router) {
	*out = *in
	if in.Neighbors != nil {
		in, out := &in.Neighbors, &out.Neighbors
		*out = make([]Neighbor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Router.
func (in *Router) DeepCopy() *Router {
	if in == nil {
		return nil
	}
	out := new(Router)
	in.DeepCopyInto(out)
	return out
}
