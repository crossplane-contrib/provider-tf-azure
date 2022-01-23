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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LogAnalyticsSavedSearchObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogAnalyticsSavedSearchParameters struct {

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FunctionAlias *string `json:"functionAlias,omitempty" tf:"function_alias,omitempty"`

	// +kubebuilder:validation:Optional
	FunctionParameters []*string `json:"functionParameters,omitempty" tf:"function_parameters,omitempty"`

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LogAnalyticsSavedSearchSpec defines the desired state of LogAnalyticsSavedSearch
type LogAnalyticsSavedSearchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsSavedSearchParameters `json:"forProvider"`
}

// LogAnalyticsSavedSearchStatus defines the observed state of LogAnalyticsSavedSearch.
type LogAnalyticsSavedSearchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsSavedSearchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSavedSearch is the Schema for the LogAnalyticsSavedSearchs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LogAnalyticsSavedSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsSavedSearchSpec   `json:"spec"`
	Status            LogAnalyticsSavedSearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSavedSearchList contains a list of LogAnalyticsSavedSearchs
type LogAnalyticsSavedSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsSavedSearch `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsSavedSearch_Kind             = "LogAnalyticsSavedSearch"
	LogAnalyticsSavedSearch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsSavedSearch_Kind}.String()
	LogAnalyticsSavedSearch_KindAPIVersion   = LogAnalyticsSavedSearch_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsSavedSearch_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsSavedSearch_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsSavedSearch{}, &LogAnalyticsSavedSearchList{})
}
