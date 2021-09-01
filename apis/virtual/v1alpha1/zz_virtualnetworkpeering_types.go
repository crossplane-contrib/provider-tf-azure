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

type VirtualNetworkPeeringObservation struct {
}

type VirtualNetworkPeeringParameters struct {
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty" tf:"allow_forwarded_traffic"`

	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty" tf:"allow_gateway_transit"`

	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty" tf:"allow_virtual_network_access"`

	Name string `json:"name" tf:"name"`

	RemoteVirtualNetworkId string `json:"remoteVirtualNetworkId" tf:"remote_virtual_network_id"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	UseRemoteGateways *bool `json:"useRemoteGateways,omitempty" tf:"use_remote_gateways"`

	VirtualNetworkName string `json:"virtualNetworkName" tf:"virtual_network_name"`
}

// VirtualNetworkPeeringSpec defines the desired state of VirtualNetworkPeering
type VirtualNetworkPeeringSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualNetworkPeeringParameters `json:"forProvider"`
}

// VirtualNetworkPeeringStatus defines the observed state of VirtualNetworkPeering.
type VirtualNetworkPeeringStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualNetworkPeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkPeering is the Schema for the VirtualNetworkPeerings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkPeeringSpec   `json:"spec"`
	Status            VirtualNetworkPeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkPeeringList contains a list of VirtualNetworkPeerings
type VirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkPeering `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkPeeringKind             = "VirtualNetworkPeering"
	VirtualNetworkPeeringGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualNetworkPeeringKind}.String()
	VirtualNetworkPeeringKindAPIVersion   = VirtualNetworkPeeringKind + "." + GroupVersion.String()
	VirtualNetworkPeeringGroupVersionKind = GroupVersion.WithKind(VirtualNetworkPeeringKind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkPeering{}, &VirtualNetworkPeeringList{})
}
