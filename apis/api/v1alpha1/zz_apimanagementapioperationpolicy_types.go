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

type ApiManagementApiOperationPolicyObservation struct {
}

type ApiManagementApiOperationPolicyParameters struct {
	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`

	ApiName string `json:"apiName" tf:"api_name"`

	OperationId string `json:"operationId" tf:"operation_id"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	XmlContent *string `json:"xmlContent,omitempty" tf:"xml_content"`

	XmlLink *string `json:"xmlLink,omitempty" tf:"xml_link"`
}

// ApiManagementApiOperationPolicySpec defines the desired state of ApiManagementApiOperationPolicy
type ApiManagementApiOperationPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementApiOperationPolicyParameters `json:"forProvider"`
}

// ApiManagementApiOperationPolicyStatus defines the observed state of ApiManagementApiOperationPolicy.
type ApiManagementApiOperationPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementApiOperationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiOperationPolicy is the Schema for the ApiManagementApiOperationPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementApiOperationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiOperationPolicySpec   `json:"spec"`
	Status            ApiManagementApiOperationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiOperationPolicyList contains a list of ApiManagementApiOperationPolicys
type ApiManagementApiOperationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementApiOperationPolicy `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementApiOperationPolicyKind             = "ApiManagementApiOperationPolicy"
	ApiManagementApiOperationPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementApiOperationPolicyKind}.String()
	ApiManagementApiOperationPolicyKindAPIVersion   = ApiManagementApiOperationPolicyKind + "." + GroupVersion.String()
	ApiManagementApiOperationPolicyGroupVersionKind = GroupVersion.WithKind(ApiManagementApiOperationPolicyKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementApiOperationPolicy{}, &ApiManagementApiOperationPolicyList{})
}
