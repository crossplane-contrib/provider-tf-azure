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

type VirtualHubBgpConnectionObservation struct {
}

type VirtualHubBgpConnectionParameters struct {
	Name string `json:"name" tf:"name"`

	PeerAsn int64 `json:"peerAsn" tf:"peer_asn"`

	PeerIp string `json:"peerIp" tf:"peer_ip"`

	VirtualHubId string `json:"virtualHubId" tf:"virtual_hub_id"`
}

// VirtualHubBgpConnectionSpec defines the desired state of VirtualHubBgpConnection
type VirtualHubBgpConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualHubBgpConnectionParameters `json:"forProvider"`
}

// VirtualHubBgpConnectionStatus defines the observed state of VirtualHubBgpConnection.
type VirtualHubBgpConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualHubBgpConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubBgpConnection is the Schema for the VirtualHubBgpConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubBgpConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubBgpConnectionSpec   `json:"spec"`
	Status            VirtualHubBgpConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubBgpConnectionList contains a list of VirtualHubBgpConnections
type VirtualHubBgpConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubBgpConnection `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubBgpConnectionKind             = "VirtualHubBgpConnection"
	VirtualHubBgpConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualHubBgpConnectionKind}.String()
	VirtualHubBgpConnectionKindAPIVersion   = VirtualHubBgpConnectionKind + "." + GroupVersion.String()
	VirtualHubBgpConnectionGroupVersionKind = GroupVersion.WithKind(VirtualHubBgpConnectionKind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubBgpConnection{}, &VirtualHubBgpConnectionList{})
}
