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

type MariadbVirtualNetworkRuleObservation struct {
}

type MariadbVirtualNetworkRuleParameters struct {
	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ServerName string `json:"serverName" tf:"server_name"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

// MariadbVirtualNetworkRuleSpec defines the desired state of MariadbVirtualNetworkRule
type MariadbVirtualNetworkRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MariadbVirtualNetworkRuleParameters `json:"forProvider"`
}

// MariadbVirtualNetworkRuleStatus defines the observed state of MariadbVirtualNetworkRule.
type MariadbVirtualNetworkRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MariadbVirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MariadbVirtualNetworkRule is the Schema for the MariadbVirtualNetworkRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MariadbVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MariadbVirtualNetworkRuleSpec   `json:"spec"`
	Status            MariadbVirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MariadbVirtualNetworkRuleList contains a list of MariadbVirtualNetworkRules
type MariadbVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MariadbVirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	MariadbVirtualNetworkRuleKind             = "MariadbVirtualNetworkRule"
	MariadbVirtualNetworkRuleGroupKind        = schema.GroupKind{Group: Group, Kind: MariadbVirtualNetworkRuleKind}.String()
	MariadbVirtualNetworkRuleKindAPIVersion   = MariadbVirtualNetworkRuleKind + "." + GroupVersion.String()
	MariadbVirtualNetworkRuleGroupVersionKind = GroupVersion.WithKind(MariadbVirtualNetworkRuleKind)
)

func init() {
	SchemeBuilder.Register(&MariadbVirtualNetworkRule{}, &MariadbVirtualNetworkRuleList{})
}
