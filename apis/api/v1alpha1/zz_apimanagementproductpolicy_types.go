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

type ApiManagementProductPolicyObservation struct {
}

type ApiManagementProductPolicyParameters struct {
	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`

	ProductId string `json:"productId" tf:"product_id"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	XmlContent *string `json:"xmlContent,omitempty" tf:"xml_content"`

	XmlLink *string `json:"xmlLink,omitempty" tf:"xml_link"`
}

// ApiManagementProductPolicySpec defines the desired state of ApiManagementProductPolicy
type ApiManagementProductPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementProductPolicyParameters `json:"forProvider"`
}

// ApiManagementProductPolicyStatus defines the observed state of ApiManagementProductPolicy.
type ApiManagementProductPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementProductPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementProductPolicy is the Schema for the ApiManagementProductPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementProductPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementProductPolicySpec   `json:"spec"`
	Status            ApiManagementProductPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementProductPolicyList contains a list of ApiManagementProductPolicys
type ApiManagementProductPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementProductPolicy `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementProductPolicyKind             = "ApiManagementProductPolicy"
	ApiManagementProductPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementProductPolicyKind}.String()
	ApiManagementProductPolicyKindAPIVersion   = ApiManagementProductPolicyKind + "." + GroupVersion.String()
	ApiManagementProductPolicyGroupVersionKind = GroupVersion.WithKind(ApiManagementProductPolicyKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementProductPolicy{}, &ApiManagementProductPolicyList{})
}
