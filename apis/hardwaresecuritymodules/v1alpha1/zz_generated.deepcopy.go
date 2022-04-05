//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedHardwareSecurityModule) DeepCopyInto(out *DedicatedHardwareSecurityModule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedHardwareSecurityModule.
func (in *DedicatedHardwareSecurityModule) DeepCopy() *DedicatedHardwareSecurityModule {
	if in == nil {
		return nil
	}
	out := new(DedicatedHardwareSecurityModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedHardwareSecurityModule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedHardwareSecurityModuleList) DeepCopyInto(out *DedicatedHardwareSecurityModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DedicatedHardwareSecurityModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedHardwareSecurityModuleList.
func (in *DedicatedHardwareSecurityModuleList) DeepCopy() *DedicatedHardwareSecurityModuleList {
	if in == nil {
		return nil
	}
	out := new(DedicatedHardwareSecurityModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedHardwareSecurityModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedHardwareSecurityModuleObservation) DeepCopyInto(out *DedicatedHardwareSecurityModuleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedHardwareSecurityModuleObservation.
func (in *DedicatedHardwareSecurityModuleObservation) DeepCopy() *DedicatedHardwareSecurityModuleObservation {
	if in == nil {
		return nil
	}
	out := new(DedicatedHardwareSecurityModuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedHardwareSecurityModuleParameters) DeepCopyInto(out *DedicatedHardwareSecurityModuleParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkProfile != nil {
		in, out := &in.NetworkProfile, &out.NetworkProfile
		*out = make([]NetworkProfileParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.StampID != nil {
		in, out := &in.StampID, &out.StampID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedHardwareSecurityModuleParameters.
func (in *DedicatedHardwareSecurityModuleParameters) DeepCopy() *DedicatedHardwareSecurityModuleParameters {
	if in == nil {
		return nil
	}
	out := new(DedicatedHardwareSecurityModuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedHardwareSecurityModuleSpec) DeepCopyInto(out *DedicatedHardwareSecurityModuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedHardwareSecurityModuleSpec.
func (in *DedicatedHardwareSecurityModuleSpec) DeepCopy() *DedicatedHardwareSecurityModuleSpec {
	if in == nil {
		return nil
	}
	out := new(DedicatedHardwareSecurityModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedHardwareSecurityModuleStatus) DeepCopyInto(out *DedicatedHardwareSecurityModuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedHardwareSecurityModuleStatus.
func (in *DedicatedHardwareSecurityModuleStatus) DeepCopy() *DedicatedHardwareSecurityModuleStatus {
	if in == nil {
		return nil
	}
	out := new(DedicatedHardwareSecurityModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProfileObservation) DeepCopyInto(out *NetworkProfileObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProfileObservation.
func (in *NetworkProfileObservation) DeepCopy() *NetworkProfileObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProfileParameters) DeepCopyInto(out *NetworkProfileParameters) {
	*out = *in
	if in.NetworkInterfacePrivateIPAddresses != nil {
		in, out := &in.NetworkInterfacePrivateIPAddresses, &out.NetworkInterfacePrivateIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProfileParameters.
func (in *NetworkProfileParameters) DeepCopy() *NetworkProfileParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkProfileParameters)
	in.DeepCopyInto(out)
	return out
}
