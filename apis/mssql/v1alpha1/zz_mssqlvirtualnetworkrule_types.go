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

type MssqlVirtualNetworkRuleObservation struct {
}

type MssqlVirtualNetworkRuleParameters struct {
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint"`

	Name string `json:"name" tf:"name"`

	ServerId string `json:"serverId" tf:"server_id"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

// MssqlVirtualNetworkRuleSpec defines the desired state of MssqlVirtualNetworkRule
type MssqlVirtualNetworkRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MssqlVirtualNetworkRuleParameters `json:"forProvider"`
}

// MssqlVirtualNetworkRuleStatus defines the observed state of MssqlVirtualNetworkRule.
type MssqlVirtualNetworkRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MssqlVirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MssqlVirtualNetworkRule is the Schema for the MssqlVirtualNetworkRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MssqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlVirtualNetworkRuleSpec   `json:"spec"`
	Status            MssqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MssqlVirtualNetworkRuleList contains a list of MssqlVirtualNetworkRules
type MssqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MssqlVirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	MssqlVirtualNetworkRuleKind             = "MssqlVirtualNetworkRule"
	MssqlVirtualNetworkRuleGroupKind        = schema.GroupKind{Group: Group, Kind: MssqlVirtualNetworkRuleKind}.String()
	MssqlVirtualNetworkRuleKindAPIVersion   = MssqlVirtualNetworkRuleKind + "." + GroupVersion.String()
	MssqlVirtualNetworkRuleGroupVersionKind = GroupVersion.WithKind(MssqlVirtualNetworkRuleKind)
)

func init() {
	SchemeBuilder.Register(&MssqlVirtualNetworkRule{}, &MssqlVirtualNetworkRuleList{})
}
