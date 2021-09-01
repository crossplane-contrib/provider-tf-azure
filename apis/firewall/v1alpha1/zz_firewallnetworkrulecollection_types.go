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

type FirewallNetworkRuleCollectionObservation struct {
}

type FirewallNetworkRuleCollectionParameters struct {
	Action string `json:"action" tf:"action"`

	AzureFirewallName string `json:"azureFirewallName" tf:"azure_firewall_name"`

	Name string `json:"name" tf:"name"`

	Priority int64 `json:"priority" tf:"priority"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Rule []FirewallNetworkRuleCollectionRuleParameters `json:"rule" tf:"rule"`
}

type FirewallNetworkRuleCollectionRuleObservation struct {
}

type FirewallNetworkRuleCollectionRuleParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DestinationAddresses []string `json:"destinationAddresses,omitempty" tf:"destination_addresses"`

	DestinationFqdns []string `json:"destinationFqdns,omitempty" tf:"destination_fqdns"`

	DestinationIpGroups []string `json:"destinationIpGroups,omitempty" tf:"destination_ip_groups"`

	DestinationPorts []string `json:"destinationPorts" tf:"destination_ports"`

	Name string `json:"name" tf:"name"`

	Protocols []string `json:"protocols" tf:"protocols"`

	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses"`

	SourceIpGroups []string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups"`
}

// FirewallNetworkRuleCollectionSpec defines the desired state of FirewallNetworkRuleCollection
type FirewallNetworkRuleCollectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FirewallNetworkRuleCollectionParameters `json:"forProvider"`
}

// FirewallNetworkRuleCollectionStatus defines the observed state of FirewallNetworkRuleCollection.
type FirewallNetworkRuleCollectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FirewallNetworkRuleCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNetworkRuleCollection is the Schema for the FirewallNetworkRuleCollections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FirewallNetworkRuleCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNetworkRuleCollectionSpec   `json:"spec"`
	Status            FirewallNetworkRuleCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNetworkRuleCollectionList contains a list of FirewallNetworkRuleCollections
type FirewallNetworkRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallNetworkRuleCollection `json:"items"`
}

// Repository type metadata.
var (
	FirewallNetworkRuleCollectionKind             = "FirewallNetworkRuleCollection"
	FirewallNetworkRuleCollectionGroupKind        = schema.GroupKind{Group: Group, Kind: FirewallNetworkRuleCollectionKind}.String()
	FirewallNetworkRuleCollectionKindAPIVersion   = FirewallNetworkRuleCollectionKind + "." + GroupVersion.String()
	FirewallNetworkRuleCollectionGroupVersionKind = GroupVersion.WithKind(FirewallNetworkRuleCollectionKind)
)

func init() {
	SchemeBuilder.Register(&FirewallNetworkRuleCollection{}, &FirewallNetworkRuleCollectionList{})
}
