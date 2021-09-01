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
func (in *BackupContainerStorageAccount) DeepCopyInto(out *BackupContainerStorageAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupContainerStorageAccount.
func (in *BackupContainerStorageAccount) DeepCopy() *BackupContainerStorageAccount {
	if in == nil {
		return nil
	}
	out := new(BackupContainerStorageAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupContainerStorageAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupContainerStorageAccountList) DeepCopyInto(out *BackupContainerStorageAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupContainerStorageAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupContainerStorageAccountList.
func (in *BackupContainerStorageAccountList) DeepCopy() *BackupContainerStorageAccountList {
	if in == nil {
		return nil
	}
	out := new(BackupContainerStorageAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupContainerStorageAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupContainerStorageAccountObservation) DeepCopyInto(out *BackupContainerStorageAccountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupContainerStorageAccountObservation.
func (in *BackupContainerStorageAccountObservation) DeepCopy() *BackupContainerStorageAccountObservation {
	if in == nil {
		return nil
	}
	out := new(BackupContainerStorageAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupContainerStorageAccountParameters) DeepCopyInto(out *BackupContainerStorageAccountParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupContainerStorageAccountParameters.
func (in *BackupContainerStorageAccountParameters) DeepCopy() *BackupContainerStorageAccountParameters {
	if in == nil {
		return nil
	}
	out := new(BackupContainerStorageAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupContainerStorageAccountSpec) DeepCopyInto(out *BackupContainerStorageAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupContainerStorageAccountSpec.
func (in *BackupContainerStorageAccountSpec) DeepCopy() *BackupContainerStorageAccountSpec {
	if in == nil {
		return nil
	}
	out := new(BackupContainerStorageAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupContainerStorageAccountStatus) DeepCopyInto(out *BackupContainerStorageAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupContainerStorageAccountStatus.
func (in *BackupContainerStorageAccountStatus) DeepCopy() *BackupContainerStorageAccountStatus {
	if in == nil {
		return nil
	}
	out := new(BackupContainerStorageAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupObservation) DeepCopyInto(out *BackupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupObservation.
func (in *BackupObservation) DeepCopy() *BackupObservation {
	if in == nil {
		return nil
	}
	out := new(BackupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupParameters) DeepCopyInto(out *BackupParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupParameters.
func (in *BackupParameters) DeepCopy() *BackupParameters {
	if in == nil {
		return nil
	}
	out := new(BackupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyFileShare) DeepCopyInto(out *BackupPolicyFileShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyFileShare.
func (in *BackupPolicyFileShare) DeepCopy() *BackupPolicyFileShare {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyFileShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicyFileShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyFileShareList) DeepCopyInto(out *BackupPolicyFileShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupPolicyFileShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyFileShareList.
func (in *BackupPolicyFileShareList) DeepCopy() *BackupPolicyFileShareList {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyFileShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicyFileShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyFileShareObservation) DeepCopyInto(out *BackupPolicyFileShareObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyFileShareObservation.
func (in *BackupPolicyFileShareObservation) DeepCopy() *BackupPolicyFileShareObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyFileShareObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyFileShareParameters) DeepCopyInto(out *BackupPolicyFileShareParameters) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = make([]BackupParameters, len(*in))
		copy(*out, *in)
	}
	if in.RetentionDaily != nil {
		in, out := &in.RetentionDaily, &out.RetentionDaily
		*out = make([]RetentionDailyParameters, len(*in))
		copy(*out, *in)
	}
	if in.RetentionMonthly != nil {
		in, out := &in.RetentionMonthly, &out.RetentionMonthly
		*out = make([]RetentionMonthlyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetentionWeekly != nil {
		in, out := &in.RetentionWeekly, &out.RetentionWeekly
		*out = make([]RetentionWeeklyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetentionYearly != nil {
		in, out := &in.RetentionYearly, &out.RetentionYearly
		*out = make([]RetentionYearlyParameters, len(*in))
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
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyFileShareParameters.
func (in *BackupPolicyFileShareParameters) DeepCopy() *BackupPolicyFileShareParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyFileShareParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyFileShareSpec) DeepCopyInto(out *BackupPolicyFileShareSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyFileShareSpec.
func (in *BackupPolicyFileShareSpec) DeepCopy() *BackupPolicyFileShareSpec {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyFileShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyFileShareStatus) DeepCopyInto(out *BackupPolicyFileShareStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyFileShareStatus.
func (in *BackupPolicyFileShareStatus) DeepCopy() *BackupPolicyFileShareStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyFileShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVm) DeepCopyInto(out *BackupPolicyVm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVm.
func (in *BackupPolicyVm) DeepCopy() *BackupPolicyVm {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicyVm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmBackupObservation) DeepCopyInto(out *BackupPolicyVmBackupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmBackupObservation.
func (in *BackupPolicyVmBackupObservation) DeepCopy() *BackupPolicyVmBackupObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmBackupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmBackupParameters) DeepCopyInto(out *BackupPolicyVmBackupParameters) {
	*out = *in
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmBackupParameters.
func (in *BackupPolicyVmBackupParameters) DeepCopy() *BackupPolicyVmBackupParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmBackupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmList) DeepCopyInto(out *BackupPolicyVmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupPolicyVm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmList.
func (in *BackupPolicyVmList) DeepCopy() *BackupPolicyVmList {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicyVmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmObservation) DeepCopyInto(out *BackupPolicyVmObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmObservation.
func (in *BackupPolicyVmObservation) DeepCopy() *BackupPolicyVmObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmParameters) DeepCopyInto(out *BackupPolicyVmParameters) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = make([]BackupPolicyVmBackupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstantRestoreRetentionDays != nil {
		in, out := &in.InstantRestoreRetentionDays, &out.InstantRestoreRetentionDays
		*out = new(int64)
		**out = **in
	}
	if in.RetentionDaily != nil {
		in, out := &in.RetentionDaily, &out.RetentionDaily
		*out = make([]BackupPolicyVmRetentionDailyParameters, len(*in))
		copy(*out, *in)
	}
	if in.RetentionMonthly != nil {
		in, out := &in.RetentionMonthly, &out.RetentionMonthly
		*out = make([]BackupPolicyVmRetentionMonthlyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetentionWeekly != nil {
		in, out := &in.RetentionWeekly, &out.RetentionWeekly
		*out = make([]BackupPolicyVmRetentionWeeklyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetentionYearly != nil {
		in, out := &in.RetentionYearly, &out.RetentionYearly
		*out = make([]BackupPolicyVmRetentionYearlyParameters, len(*in))
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
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmParameters.
func (in *BackupPolicyVmParameters) DeepCopy() *BackupPolicyVmParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionDailyObservation) DeepCopyInto(out *BackupPolicyVmRetentionDailyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionDailyObservation.
func (in *BackupPolicyVmRetentionDailyObservation) DeepCopy() *BackupPolicyVmRetentionDailyObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionDailyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionDailyParameters) DeepCopyInto(out *BackupPolicyVmRetentionDailyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionDailyParameters.
func (in *BackupPolicyVmRetentionDailyParameters) DeepCopy() *BackupPolicyVmRetentionDailyParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionDailyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionMonthlyObservation) DeepCopyInto(out *BackupPolicyVmRetentionMonthlyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionMonthlyObservation.
func (in *BackupPolicyVmRetentionMonthlyObservation) DeepCopy() *BackupPolicyVmRetentionMonthlyObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionMonthlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionMonthlyParameters) DeepCopyInto(out *BackupPolicyVmRetentionMonthlyParameters) {
	*out = *in
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionMonthlyParameters.
func (in *BackupPolicyVmRetentionMonthlyParameters) DeepCopy() *BackupPolicyVmRetentionMonthlyParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionMonthlyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionWeeklyObservation) DeepCopyInto(out *BackupPolicyVmRetentionWeeklyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionWeeklyObservation.
func (in *BackupPolicyVmRetentionWeeklyObservation) DeepCopy() *BackupPolicyVmRetentionWeeklyObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionWeeklyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionWeeklyParameters) DeepCopyInto(out *BackupPolicyVmRetentionWeeklyParameters) {
	*out = *in
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionWeeklyParameters.
func (in *BackupPolicyVmRetentionWeeklyParameters) DeepCopy() *BackupPolicyVmRetentionWeeklyParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionWeeklyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionYearlyObservation) DeepCopyInto(out *BackupPolicyVmRetentionYearlyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionYearlyObservation.
func (in *BackupPolicyVmRetentionYearlyObservation) DeepCopy() *BackupPolicyVmRetentionYearlyObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionYearlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmRetentionYearlyParameters) DeepCopyInto(out *BackupPolicyVmRetentionYearlyParameters) {
	*out = *in
	if in.Months != nil {
		in, out := &in.Months, &out.Months
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmRetentionYearlyParameters.
func (in *BackupPolicyVmRetentionYearlyParameters) DeepCopy() *BackupPolicyVmRetentionYearlyParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmRetentionYearlyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmSpec) DeepCopyInto(out *BackupPolicyVmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmSpec.
func (in *BackupPolicyVmSpec) DeepCopy() *BackupPolicyVmSpec {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyVmStatus) DeepCopyInto(out *BackupPolicyVmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyVmStatus.
func (in *BackupPolicyVmStatus) DeepCopy() *BackupPolicyVmStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyVmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedFileShare) DeepCopyInto(out *BackupProtectedFileShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedFileShare.
func (in *BackupProtectedFileShare) DeepCopy() *BackupProtectedFileShare {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedFileShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupProtectedFileShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedFileShareList) DeepCopyInto(out *BackupProtectedFileShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupProtectedFileShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedFileShareList.
func (in *BackupProtectedFileShareList) DeepCopy() *BackupProtectedFileShareList {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedFileShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupProtectedFileShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedFileShareObservation) DeepCopyInto(out *BackupProtectedFileShareObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedFileShareObservation.
func (in *BackupProtectedFileShareObservation) DeepCopy() *BackupProtectedFileShareObservation {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedFileShareObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedFileShareParameters) DeepCopyInto(out *BackupProtectedFileShareParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedFileShareParameters.
func (in *BackupProtectedFileShareParameters) DeepCopy() *BackupProtectedFileShareParameters {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedFileShareParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedFileShareSpec) DeepCopyInto(out *BackupProtectedFileShareSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedFileShareSpec.
func (in *BackupProtectedFileShareSpec) DeepCopy() *BackupProtectedFileShareSpec {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedFileShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedFileShareStatus) DeepCopyInto(out *BackupProtectedFileShareStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedFileShareStatus.
func (in *BackupProtectedFileShareStatus) DeepCopy() *BackupProtectedFileShareStatus {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedFileShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedVm) DeepCopyInto(out *BackupProtectedVm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedVm.
func (in *BackupProtectedVm) DeepCopy() *BackupProtectedVm {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedVm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupProtectedVm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedVmList) DeepCopyInto(out *BackupProtectedVmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupProtectedVm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedVmList.
func (in *BackupProtectedVmList) DeepCopy() *BackupProtectedVmList {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedVmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupProtectedVmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedVmObservation) DeepCopyInto(out *BackupProtectedVmObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedVmObservation.
func (in *BackupProtectedVmObservation) DeepCopy() *BackupProtectedVmObservation {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedVmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedVmParameters) DeepCopyInto(out *BackupProtectedVmParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedVmParameters.
func (in *BackupProtectedVmParameters) DeepCopy() *BackupProtectedVmParameters {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedVmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedVmSpec) DeepCopyInto(out *BackupProtectedVmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedVmSpec.
func (in *BackupProtectedVmSpec) DeepCopy() *BackupProtectedVmSpec {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedVmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProtectedVmStatus) DeepCopyInto(out *BackupProtectedVmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProtectedVmStatus.
func (in *BackupProtectedVmStatus) DeepCopy() *BackupProtectedVmStatus {
	if in == nil {
		return nil
	}
	out := new(BackupProtectedVmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionDailyObservation) DeepCopyInto(out *RetentionDailyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionDailyObservation.
func (in *RetentionDailyObservation) DeepCopy() *RetentionDailyObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionDailyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionDailyParameters) DeepCopyInto(out *RetentionDailyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionDailyParameters.
func (in *RetentionDailyParameters) DeepCopy() *RetentionDailyParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionDailyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionMonthlyObservation) DeepCopyInto(out *RetentionMonthlyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionMonthlyObservation.
func (in *RetentionMonthlyObservation) DeepCopy() *RetentionMonthlyObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionMonthlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionMonthlyParameters) DeepCopyInto(out *RetentionMonthlyParameters) {
	*out = *in
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionMonthlyParameters.
func (in *RetentionMonthlyParameters) DeepCopy() *RetentionMonthlyParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionMonthlyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionWeeklyObservation) DeepCopyInto(out *RetentionWeeklyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionWeeklyObservation.
func (in *RetentionWeeklyObservation) DeepCopy() *RetentionWeeklyObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionWeeklyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionWeeklyParameters) DeepCopyInto(out *RetentionWeeklyParameters) {
	*out = *in
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionWeeklyParameters.
func (in *RetentionWeeklyParameters) DeepCopy() *RetentionWeeklyParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionWeeklyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionYearlyObservation) DeepCopyInto(out *RetentionYearlyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionYearlyObservation.
func (in *RetentionYearlyObservation) DeepCopy() *RetentionYearlyObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionYearlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionYearlyParameters) DeepCopyInto(out *RetentionYearlyParameters) {
	*out = *in
	if in.Months != nil {
		in, out := &in.Months, &out.Months
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionYearlyParameters.
func (in *RetentionYearlyParameters) DeepCopy() *RetentionYearlyParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionYearlyParameters)
	in.DeepCopyInto(out)
	return out
}
