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

type BotConnectionObservation struct {
}

type BotConnectionParameters struct {
	BotName string `json:"botName" tf:"bot_name"`

	ClientId string `json:"clientId" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Scopes *string `json:"scopes,omitempty" tf:"scopes"`

	ServiceProviderName string `json:"serviceProviderName" tf:"service_provider_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// BotConnectionSpec defines the desired state of BotConnection
type BotConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BotConnectionParameters `json:"forProvider"`
}

// BotConnectionStatus defines the observed state of BotConnection.
type BotConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BotConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotConnection is the Schema for the BotConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotConnectionSpec   `json:"spec"`
	Status            BotConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotConnectionList contains a list of BotConnections
type BotConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotConnection `json:"items"`
}

// Repository type metadata.
var (
	BotConnectionKind             = "BotConnection"
	BotConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: BotConnectionKind}.String()
	BotConnectionKindAPIVersion   = BotConnectionKind + "." + GroupVersion.String()
	BotConnectionGroupVersionKind = GroupVersion.WithKind(BotConnectionKind)
)

func init() {
	SchemeBuilder.Register(&BotConnection{}, &BotConnectionList{})
}
