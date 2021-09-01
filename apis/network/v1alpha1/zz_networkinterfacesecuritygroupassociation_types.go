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

type NetworkInterfaceSecurityGroupAssociationObservation struct {
}

type NetworkInterfaceSecurityGroupAssociationParameters struct {
	NetworkInterfaceId string `json:"networkInterfaceId" tf:"network_interface_id"`

	NetworkSecurityGroupId string `json:"networkSecurityGroupId" tf:"network_security_group_id"`
}

// NetworkInterfaceSecurityGroupAssociationSpec defines the desired state of NetworkInterfaceSecurityGroupAssociation
type NetworkInterfaceSecurityGroupAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NetworkInterfaceSecurityGroupAssociationParameters `json:"forProvider"`
}

// NetworkInterfaceSecurityGroupAssociationStatus defines the observed state of NetworkInterfaceSecurityGroupAssociation.
type NetworkInterfaceSecurityGroupAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NetworkInterfaceSecurityGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSecurityGroupAssociation is the Schema for the NetworkInterfaceSecurityGroupAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSecurityGroupAssociationSpec   `json:"spec"`
	Status            NetworkInterfaceSecurityGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSecurityGroupAssociationList contains a list of NetworkInterfaceSecurityGroupAssociations
type NetworkInterfaceSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceSecurityGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceSecurityGroupAssociationKind             = "NetworkInterfaceSecurityGroupAssociation"
	NetworkInterfaceSecurityGroupAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkInterfaceSecurityGroupAssociationKind}.String()
	NetworkInterfaceSecurityGroupAssociationKindAPIVersion   = NetworkInterfaceSecurityGroupAssociationKind + "." + GroupVersion.String()
	NetworkInterfaceSecurityGroupAssociationGroupVersionKind = GroupVersion.WithKind(NetworkInterfaceSecurityGroupAssociationKind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceSecurityGroupAssociation{}, &NetworkInterfaceSecurityGroupAssociationList{})
}
