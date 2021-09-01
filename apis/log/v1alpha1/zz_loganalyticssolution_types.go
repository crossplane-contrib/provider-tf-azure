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

type LogAnalyticsSolutionObservation struct {
}

type LogAnalyticsSolutionParameters struct {
	Location string `json:"location" tf:"location"`

	Plan []PlanParameters `json:"plan" tf:"plan"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SolutionName string `json:"solutionName" tf:"solution_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	WorkspaceName string `json:"workspaceName" tf:"workspace_name"`

	WorkspaceResourceId string `json:"workspaceResourceId" tf:"workspace_resource_id"`
}

type PlanObservation struct {
	Name string `json:"name" tf:"name"`
}

type PlanParameters struct {
	Product string `json:"product" tf:"product"`

	PromotionCode *string `json:"promotionCode,omitempty" tf:"promotion_code"`

	Publisher string `json:"publisher" tf:"publisher"`
}

// LogAnalyticsSolutionSpec defines the desired state of LogAnalyticsSolution
type LogAnalyticsSolutionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LogAnalyticsSolutionParameters `json:"forProvider"`
}

// LogAnalyticsSolutionStatus defines the observed state of LogAnalyticsSolution.
type LogAnalyticsSolutionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LogAnalyticsSolutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSolution is the Schema for the LogAnalyticsSolutions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsSolution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsSolutionSpec   `json:"spec"`
	Status            LogAnalyticsSolutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSolutionList contains a list of LogAnalyticsSolutions
type LogAnalyticsSolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsSolution `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsSolutionKind             = "LogAnalyticsSolution"
	LogAnalyticsSolutionGroupKind        = schema.GroupKind{Group: Group, Kind: LogAnalyticsSolutionKind}.String()
	LogAnalyticsSolutionKindAPIVersion   = LogAnalyticsSolutionKind + "." + GroupVersion.String()
	LogAnalyticsSolutionGroupVersionKind = GroupVersion.WithKind(LogAnalyticsSolutionKind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsSolution{}, &LogAnalyticsSolutionList{})
}
