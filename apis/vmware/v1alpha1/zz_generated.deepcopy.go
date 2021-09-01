// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitObservation) DeepCopyInto(out *CircuitObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitObservation.
func (in *CircuitObservation) DeepCopy() *CircuitObservation {
	if in == nil {
		return nil
	}
	out := new(CircuitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitParameters) DeepCopyInto(out *CircuitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitParameters.
func (in *CircuitParameters) DeepCopy() *CircuitParameters {
	if in == nil {
		return nil
	}
	out := new(CircuitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementClusterObservation) DeepCopyInto(out *ManagementClusterObservation) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementClusterObservation.
func (in *ManagementClusterObservation) DeepCopy() *ManagementClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ManagementClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementClusterParameters) DeepCopyInto(out *ManagementClusterParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementClusterParameters.
func (in *ManagementClusterParameters) DeepCopy() *ManagementClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ManagementClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareCluster) DeepCopyInto(out *VmwareCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareCluster.
func (in *VmwareCluster) DeepCopy() *VmwareCluster {
	if in == nil {
		return nil
	}
	out := new(VmwareCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmwareCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareClusterList) DeepCopyInto(out *VmwareClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VmwareCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareClusterList.
func (in *VmwareClusterList) DeepCopy() *VmwareClusterList {
	if in == nil {
		return nil
	}
	out := new(VmwareClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmwareClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareClusterObservation) DeepCopyInto(out *VmwareClusterObservation) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareClusterObservation.
func (in *VmwareClusterObservation) DeepCopy() *VmwareClusterObservation {
	if in == nil {
		return nil
	}
	out := new(VmwareClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareClusterParameters) DeepCopyInto(out *VmwareClusterParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareClusterParameters.
func (in *VmwareClusterParameters) DeepCopy() *VmwareClusterParameters {
	if in == nil {
		return nil
	}
	out := new(VmwareClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareClusterSpec) DeepCopyInto(out *VmwareClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareClusterSpec.
func (in *VmwareClusterSpec) DeepCopy() *VmwareClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VmwareClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareClusterStatus) DeepCopyInto(out *VmwareClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareClusterStatus.
func (in *VmwareClusterStatus) DeepCopy() *VmwareClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VmwareClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareExpressRouteAuthorization) DeepCopyInto(out *VmwareExpressRouteAuthorization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareExpressRouteAuthorization.
func (in *VmwareExpressRouteAuthorization) DeepCopy() *VmwareExpressRouteAuthorization {
	if in == nil {
		return nil
	}
	out := new(VmwareExpressRouteAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmwareExpressRouteAuthorization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareExpressRouteAuthorizationList) DeepCopyInto(out *VmwareExpressRouteAuthorizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VmwareExpressRouteAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareExpressRouteAuthorizationList.
func (in *VmwareExpressRouteAuthorizationList) DeepCopy() *VmwareExpressRouteAuthorizationList {
	if in == nil {
		return nil
	}
	out := new(VmwareExpressRouteAuthorizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmwareExpressRouteAuthorizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareExpressRouteAuthorizationObservation) DeepCopyInto(out *VmwareExpressRouteAuthorizationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareExpressRouteAuthorizationObservation.
func (in *VmwareExpressRouteAuthorizationObservation) DeepCopy() *VmwareExpressRouteAuthorizationObservation {
	if in == nil {
		return nil
	}
	out := new(VmwareExpressRouteAuthorizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareExpressRouteAuthorizationParameters) DeepCopyInto(out *VmwareExpressRouteAuthorizationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareExpressRouteAuthorizationParameters.
func (in *VmwareExpressRouteAuthorizationParameters) DeepCopy() *VmwareExpressRouteAuthorizationParameters {
	if in == nil {
		return nil
	}
	out := new(VmwareExpressRouteAuthorizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareExpressRouteAuthorizationSpec) DeepCopyInto(out *VmwareExpressRouteAuthorizationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareExpressRouteAuthorizationSpec.
func (in *VmwareExpressRouteAuthorizationSpec) DeepCopy() *VmwareExpressRouteAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(VmwareExpressRouteAuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareExpressRouteAuthorizationStatus) DeepCopyInto(out *VmwareExpressRouteAuthorizationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareExpressRouteAuthorizationStatus.
func (in *VmwareExpressRouteAuthorizationStatus) DeepCopy() *VmwareExpressRouteAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := new(VmwareExpressRouteAuthorizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwarePrivateCloud) DeepCopyInto(out *VmwarePrivateCloud) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwarePrivateCloud.
func (in *VmwarePrivateCloud) DeepCopy() *VmwarePrivateCloud {
	if in == nil {
		return nil
	}
	out := new(VmwarePrivateCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmwarePrivateCloud) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwarePrivateCloudList) DeepCopyInto(out *VmwarePrivateCloudList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VmwarePrivateCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwarePrivateCloudList.
func (in *VmwarePrivateCloudList) DeepCopy() *VmwarePrivateCloudList {
	if in == nil {
		return nil
	}
	out := new(VmwarePrivateCloudList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmwarePrivateCloudList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwarePrivateCloudObservation) DeepCopyInto(out *VmwarePrivateCloudObservation) {
	*out = *in
	if in.Circuit != nil {
		in, out := &in.Circuit, &out.Circuit
		*out = make([]CircuitObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwarePrivateCloudObservation.
func (in *VmwarePrivateCloudObservation) DeepCopy() *VmwarePrivateCloudObservation {
	if in == nil {
		return nil
	}
	out := new(VmwarePrivateCloudObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwarePrivateCloudParameters) DeepCopyInto(out *VmwarePrivateCloudParameters) {
	*out = *in
	if in.InternetConnectionEnabled != nil {
		in, out := &in.InternetConnectionEnabled, &out.InternetConnectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ManagementCluster != nil {
		in, out := &in.ManagementCluster, &out.ManagementCluster
		*out = make([]ManagementClusterParameters, len(*in))
		copy(*out, *in)
	}
	if in.NsxtPassword != nil {
		in, out := &in.NsxtPassword, &out.NsxtPassword
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VcenterPassword != nil {
		in, out := &in.VcenterPassword, &out.VcenterPassword
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwarePrivateCloudParameters.
func (in *VmwarePrivateCloudParameters) DeepCopy() *VmwarePrivateCloudParameters {
	if in == nil {
		return nil
	}
	out := new(VmwarePrivateCloudParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwarePrivateCloudSpec) DeepCopyInto(out *VmwarePrivateCloudSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwarePrivateCloudSpec.
func (in *VmwarePrivateCloudSpec) DeepCopy() *VmwarePrivateCloudSpec {
	if in == nil {
		return nil
	}
	out := new(VmwarePrivateCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwarePrivateCloudStatus) DeepCopyInto(out *VmwarePrivateCloudStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwarePrivateCloudStatus.
func (in *VmwarePrivateCloudStatus) DeepCopy() *VmwarePrivateCloudStatus {
	if in == nil {
		return nil
	}
	out := new(VmwarePrivateCloudStatus)
	in.DeepCopyInto(out)
	return out
}
