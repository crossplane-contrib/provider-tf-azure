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
func (in *CognitiveAccount) DeepCopyInto(out *CognitiveAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitiveAccount.
func (in *CognitiveAccount) DeepCopy() *CognitiveAccount {
	if in == nil {
		return nil
	}
	out := new(CognitiveAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitiveAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitiveAccountList) DeepCopyInto(out *CognitiveAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CognitiveAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitiveAccountList.
func (in *CognitiveAccountList) DeepCopy() *CognitiveAccountList {
	if in == nil {
		return nil
	}
	out := new(CognitiveAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitiveAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitiveAccountObservation) DeepCopyInto(out *CognitiveAccountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitiveAccountObservation.
func (in *CognitiveAccountObservation) DeepCopy() *CognitiveAccountObservation {
	if in == nil {
		return nil
	}
	out := new(CognitiveAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitiveAccountParameters) DeepCopyInto(out *CognitiveAccountParameters) {
	*out = *in
	if in.CustomSubdomainName != nil {
		in, out := &in.CustomSubdomainName, &out.CustomSubdomainName
		*out = new(string)
		**out = **in
	}
	if in.Fqdns != nil {
		in, out := &in.Fqdns, &out.Fqdns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MetricsAdvisorAadClientId != nil {
		in, out := &in.MetricsAdvisorAadClientId, &out.MetricsAdvisorAadClientId
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAadTenantId != nil {
		in, out := &in.MetricsAdvisorAadTenantId, &out.MetricsAdvisorAadTenantId
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorSuperUserName != nil {
		in, out := &in.MetricsAdvisorSuperUserName, &out.MetricsAdvisorSuperUserName
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorWebsiteName != nil {
		in, out := &in.MetricsAdvisorWebsiteName, &out.MetricsAdvisorWebsiteName
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutboundNetworkAccessRestrited != nil {
		in, out := &in.OutboundNetworkAccessRestrited, &out.OutboundNetworkAccessRestrited
		*out = new(bool)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.QnaRuntimeEndpoint != nil {
		in, out := &in.QnaRuntimeEndpoint, &out.QnaRuntimeEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitiveAccountParameters.
func (in *CognitiveAccountParameters) DeepCopy() *CognitiveAccountParameters {
	if in == nil {
		return nil
	}
	out := new(CognitiveAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitiveAccountSpec) DeepCopyInto(out *CognitiveAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitiveAccountSpec.
func (in *CognitiveAccountSpec) DeepCopy() *CognitiveAccountSpec {
	if in == nil {
		return nil
	}
	out := new(CognitiveAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitiveAccountStatus) DeepCopyInto(out *CognitiveAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitiveAccountStatus.
func (in *CognitiveAccountStatus) DeepCopy() *CognitiveAccountStatus {
	if in == nil {
		return nil
	}
	out := new(CognitiveAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsObservation) DeepCopyInto(out *NetworkAclsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsObservation.
func (in *NetworkAclsObservation) DeepCopy() *NetworkAclsObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsParameters) DeepCopyInto(out *NetworkAclsParameters) {
	*out = *in
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualNetworkSubnetIds != nil {
		in, out := &in.VirtualNetworkSubnetIds, &out.VirtualNetworkSubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsParameters.
func (in *NetworkAclsParameters) DeepCopy() *NetworkAclsParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObservation) DeepCopyInto(out *StorageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObservation.
func (in *StorageObservation) DeepCopy() *StorageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageParameters) DeepCopyInto(out *StorageParameters) {
	*out = *in
	if in.IdentityClientId != nil {
		in, out := &in.IdentityClientId, &out.IdentityClientId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageParameters.
func (in *StorageParameters) DeepCopy() *StorageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRulesObservation) DeepCopyInto(out *VirtualNetworkRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRulesObservation.
func (in *VirtualNetworkRulesObservation) DeepCopy() *VirtualNetworkRulesObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRulesParameters) DeepCopyInto(out *VirtualNetworkRulesParameters) {
	*out = *in
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRulesParameters.
func (in *VirtualNetworkRulesParameters) DeepCopy() *VirtualNetworkRulesParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRulesParameters)
	in.DeepCopyInto(out)
	return out
}
