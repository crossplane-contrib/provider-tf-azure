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

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionObservation) DeepCopyInto(out *ActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionObservation.
func (in *ActionObservation) DeepCopy() *ActionObservation {
	if in == nil {
		return nil
	}
	out := new(ActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionParameters) DeepCopyInto(out *ActionParameters) {
	*out = *in
	if in.ActionGroupID != nil {
		in, out := &in.ActionGroupID, &out.ActionGroupID
		*out = new(string)
		**out = **in
	}
	if in.WebhookProperties != nil {
		in, out := &in.WebhookProperties, &out.WebhookProperties
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionParameters.
func (in *ActionParameters) DeepCopy() *ActionParameters {
	if in == nil {
		return nil
	}
	out := new(ActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation) DeepCopyInto(out *ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation.
func (in *ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation) DeepCopy() *ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters) DeepCopyInto(out *ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters) {
	*out = *in
	if in.ComponentID != nil {
		in, out := &in.ComponentID, &out.ComponentID
		*out = new(string)
		**out = **in
	}
	if in.FailedLocationCount != nil {
		in, out := &in.FailedLocationCount, &out.FailedLocationCount
		*out = new(float64)
		**out = **in
	}
	if in.WebTestID != nil {
		in, out := &in.WebTestID, &out.WebTestID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters.
func (in *ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters) DeepCopy() *ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriteriaObservation) DeepCopyInto(out *CriteriaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaObservation.
func (in *CriteriaObservation) DeepCopy() *CriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(CriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriteriaParameters) DeepCopyInto(out *CriteriaParameters) {
	*out = *in
	if in.Aggregation != nil {
		in, out := &in.Aggregation, &out.Aggregation
		*out = new(string)
		**out = **in
	}
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]DimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricNamespace != nil {
		in, out := &in.MetricNamespace, &out.MetricNamespace
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.SkipMetricValidation != nil {
		in, out := &in.SkipMetricValidation, &out.SkipMetricValidation
		*out = new(bool)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaParameters.
func (in *CriteriaParameters) DeepCopy() *CriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(CriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionObservation) DeepCopyInto(out *DimensionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionObservation.
func (in *DimensionObservation) DeepCopy() *DimensionObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionParameters) DeepCopyInto(out *DimensionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionParameters.
func (in *DimensionParameters) DeepCopy() *DimensionParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicCriteriaDimensionObservation) DeepCopyInto(out *DynamicCriteriaDimensionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicCriteriaDimensionObservation.
func (in *DynamicCriteriaDimensionObservation) DeepCopy() *DynamicCriteriaDimensionObservation {
	if in == nil {
		return nil
	}
	out := new(DynamicCriteriaDimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicCriteriaDimensionParameters) DeepCopyInto(out *DynamicCriteriaDimensionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicCriteriaDimensionParameters.
func (in *DynamicCriteriaDimensionParameters) DeepCopy() *DynamicCriteriaDimensionParameters {
	if in == nil {
		return nil
	}
	out := new(DynamicCriteriaDimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicCriteriaObservation) DeepCopyInto(out *DynamicCriteriaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicCriteriaObservation.
func (in *DynamicCriteriaObservation) DeepCopy() *DynamicCriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(DynamicCriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicCriteriaParameters) DeepCopyInto(out *DynamicCriteriaParameters) {
	*out = *in
	if in.Aggregation != nil {
		in, out := &in.Aggregation, &out.Aggregation
		*out = new(string)
		**out = **in
	}
	if in.AlertSensitivity != nil {
		in, out := &in.AlertSensitivity, &out.AlertSensitivity
		*out = new(string)
		**out = **in
	}
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]DynamicCriteriaDimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvaluationFailureCount != nil {
		in, out := &in.EvaluationFailureCount, &out.EvaluationFailureCount
		*out = new(float64)
		**out = **in
	}
	if in.EvaluationTotalCount != nil {
		in, out := &in.EvaluationTotalCount, &out.EvaluationTotalCount
		*out = new(float64)
		**out = **in
	}
	if in.IgnoreDataBefore != nil {
		in, out := &in.IgnoreDataBefore, &out.IgnoreDataBefore
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricNamespace != nil {
		in, out := &in.MetricNamespace, &out.MetricNamespace
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.SkipMetricValidation != nil {
		in, out := &in.SkipMetricValidation, &out.SkipMetricValidation
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicCriteriaParameters.
func (in *DynamicCriteriaParameters) DeepCopy() *DynamicCriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(DynamicCriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricAlert) DeepCopyInto(out *MonitorMetricAlert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricAlert.
func (in *MonitorMetricAlert) DeepCopy() *MonitorMetricAlert {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorMetricAlert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricAlertList) DeepCopyInto(out *MonitorMetricAlertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitorMetricAlert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricAlertList.
func (in *MonitorMetricAlertList) DeepCopy() *MonitorMetricAlertList {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricAlertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorMetricAlertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricAlertObservation) DeepCopyInto(out *MonitorMetricAlertObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricAlertObservation.
func (in *MonitorMetricAlertObservation) DeepCopy() *MonitorMetricAlertObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricAlertObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricAlertParameters) DeepCopyInto(out *MonitorMetricAlertParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]ActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplicationInsightsWebTestLocationAvailabilityCriteria != nil {
		in, out := &in.ApplicationInsightsWebTestLocationAvailabilityCriteria, &out.ApplicationInsightsWebTestLocationAvailabilityCriteria
		*out = make([]ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoMitigate != nil {
		in, out := &in.AutoMitigate, &out.AutoMitigate
		*out = new(bool)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]CriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DynamicCriteria != nil {
		in, out := &in.DynamicCriteria, &out.DynamicCriteria
		*out = make([]DynamicCriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
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
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(float64)
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
	if in.TargetResourceLocation != nil {
		in, out := &in.TargetResourceLocation, &out.TargetResourceLocation
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceType != nil {
		in, out := &in.TargetResourceType, &out.TargetResourceType
		*out = new(string)
		**out = **in
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricAlertParameters.
func (in *MonitorMetricAlertParameters) DeepCopy() *MonitorMetricAlertParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricAlertParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricAlertSpec) DeepCopyInto(out *MonitorMetricAlertSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricAlertSpec.
func (in *MonitorMetricAlertSpec) DeepCopy() *MonitorMetricAlertSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricAlertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricAlertStatus) DeepCopyInto(out *MonitorMetricAlertStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricAlertStatus.
func (in *MonitorMetricAlertStatus) DeepCopy() *MonitorMetricAlertStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricAlertStatus)
	in.DeepCopyInto(out)
	return out
}
