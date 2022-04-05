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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LogAnalyticsDataSourceWindowsPerformanceCounterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogAnalyticsDataSourceWindowsPerformanceCounterParameters struct {

	// +kubebuilder:validation:Required
	CounterName *string `json:"counterName" tf:"counter_name,omitempty"`

	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// +kubebuilder:validation:Required
	IntervalSeconds *float64 `json:"intervalSeconds" tf:"interval_seconds,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ObjectName *string `json:"objectName" tf:"object_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	WorkspaceName *string `json:"workspaceName" tf:"workspace_name,omitempty"`
}

// LogAnalyticsDataSourceWindowsPerformanceCounterSpec defines the desired state of LogAnalyticsDataSourceWindowsPerformanceCounter
type LogAnalyticsDataSourceWindowsPerformanceCounterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsDataSourceWindowsPerformanceCounterParameters `json:"forProvider"`
}

// LogAnalyticsDataSourceWindowsPerformanceCounterStatus defines the observed state of LogAnalyticsDataSourceWindowsPerformanceCounter.
type LogAnalyticsDataSourceWindowsPerformanceCounterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsDataSourceWindowsPerformanceCounterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataSourceWindowsPerformanceCounter is the Schema for the LogAnalyticsDataSourceWindowsPerformanceCounters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LogAnalyticsDataSourceWindowsPerformanceCounter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsDataSourceWindowsPerformanceCounterSpec   `json:"spec"`
	Status            LogAnalyticsDataSourceWindowsPerformanceCounterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataSourceWindowsPerformanceCounterList contains a list of LogAnalyticsDataSourceWindowsPerformanceCounters
type LogAnalyticsDataSourceWindowsPerformanceCounterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsDataSourceWindowsPerformanceCounter `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsDataSourceWindowsPerformanceCounter_Kind             = "LogAnalyticsDataSourceWindowsPerformanceCounter"
	LogAnalyticsDataSourceWindowsPerformanceCounter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsDataSourceWindowsPerformanceCounter_Kind}.String()
	LogAnalyticsDataSourceWindowsPerformanceCounter_KindAPIVersion   = LogAnalyticsDataSourceWindowsPerformanceCounter_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsDataSourceWindowsPerformanceCounter_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsDataSourceWindowsPerformanceCounter_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsDataSourceWindowsPerformanceCounter{}, &LogAnalyticsDataSourceWindowsPerformanceCounterList{})
}
