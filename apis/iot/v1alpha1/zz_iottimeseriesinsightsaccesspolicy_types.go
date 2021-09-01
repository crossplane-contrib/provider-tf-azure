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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type IotTimeSeriesInsightsAccessPolicyObservation struct {
}

type IotTimeSeriesInsightsAccessPolicyParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	PrincipalObjectId string `json:"principalObjectId" tf:"principal_object_id"`

	Roles []string `json:"roles" tf:"roles"`

	TimeSeriesInsightsEnvironmentId string `json:"timeSeriesInsightsEnvironmentId" tf:"time_series_insights_environment_id"`
}

// IotTimeSeriesInsightsAccessPolicySpec defines the desired state of IotTimeSeriesInsightsAccessPolicy
type IotTimeSeriesInsightsAccessPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IotTimeSeriesInsightsAccessPolicyParameters `json:"forProvider"`
}

// IotTimeSeriesInsightsAccessPolicyStatus defines the observed state of IotTimeSeriesInsightsAccessPolicy.
type IotTimeSeriesInsightsAccessPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IotTimeSeriesInsightsAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IotTimeSeriesInsightsAccessPolicy is the Schema for the IotTimeSeriesInsightsAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IotTimeSeriesInsightsAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotTimeSeriesInsightsAccessPolicySpec   `json:"spec"`
	Status            IotTimeSeriesInsightsAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotTimeSeriesInsightsAccessPolicyList contains a list of IotTimeSeriesInsightsAccessPolicys
type IotTimeSeriesInsightsAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotTimeSeriesInsightsAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	IotTimeSeriesInsightsAccessPolicyKind             = "IotTimeSeriesInsightsAccessPolicy"
	IotTimeSeriesInsightsAccessPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: IotTimeSeriesInsightsAccessPolicyKind}.String()
	IotTimeSeriesInsightsAccessPolicyKindAPIVersion   = IotTimeSeriesInsightsAccessPolicyKind + "." + GroupVersion.String()
	IotTimeSeriesInsightsAccessPolicyGroupVersionKind = GroupVersion.WithKind(IotTimeSeriesInsightsAccessPolicyKind)
)

func init() {
	SchemeBuilder.Register(&IotTimeSeriesInsightsAccessPolicy{}, &IotTimeSeriesInsightsAccessPolicyList{})
}
