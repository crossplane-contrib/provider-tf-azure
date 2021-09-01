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

type PrivateEndpointObservation struct {
}

type PrivateEndpointParameters struct {
	AllowedRequestTypes []string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types"`

	DeniedRequestTypes []string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types"`

	Id string `json:"id" tf:"id"`
}

type PublicNetworkObservation struct {
}

type PublicNetworkParameters struct {
	AllowedRequestTypes []string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types"`

	DeniedRequestTypes []string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types"`
}

type SignalrServiceNetworkAclObservation struct {
}

type SignalrServiceNetworkAclParameters struct {
	DefaultAction string `json:"defaultAction" tf:"default_action"`

	PrivateEndpoint []PrivateEndpointParameters `json:"privateEndpoint,omitempty" tf:"private_endpoint"`

	PublicNetwork []PublicNetworkParameters `json:"publicNetwork" tf:"public_network"`

	SignalrServiceId string `json:"signalrServiceId" tf:"signalr_service_id"`
}

// SignalrServiceNetworkAclSpec defines the desired state of SignalrServiceNetworkAcl
type SignalrServiceNetworkAclSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SignalrServiceNetworkAclParameters `json:"forProvider"`
}

// SignalrServiceNetworkAclStatus defines the observed state of SignalrServiceNetworkAcl.
type SignalrServiceNetworkAclStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SignalrServiceNetworkAclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SignalrServiceNetworkAcl is the Schema for the SignalrServiceNetworkAcls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SignalrServiceNetworkAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignalrServiceNetworkAclSpec   `json:"spec"`
	Status            SignalrServiceNetworkAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SignalrServiceNetworkAclList contains a list of SignalrServiceNetworkAcls
type SignalrServiceNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignalrServiceNetworkAcl `json:"items"`
}

// Repository type metadata.
var (
	SignalrServiceNetworkAclKind             = "SignalrServiceNetworkAcl"
	SignalrServiceNetworkAclGroupKind        = schema.GroupKind{Group: Group, Kind: SignalrServiceNetworkAclKind}.String()
	SignalrServiceNetworkAclKindAPIVersion   = SignalrServiceNetworkAclKind + "." + GroupVersion.String()
	SignalrServiceNetworkAclGroupVersionKind = GroupVersion.WithKind(SignalrServiceNetworkAclKind)
)

func init() {
	SchemeBuilder.Register(&SignalrServiceNetworkAcl{}, &SignalrServiceNetworkAclList{})
}
