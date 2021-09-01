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

type CorsObservation struct {
}

type CorsParameters struct {
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
}

type FeaturesObservation struct {
}

type FeaturesParameters struct {
	Flag string `json:"flag" tf:"flag"`

	Value string `json:"value" tf:"value"`
}

type SignalrServiceObservation struct {
	Hostname string `json:"hostname" tf:"hostname"`

	IpAddress string `json:"ipAddress" tf:"ip_address"`

	PrimaryAccessKey string `json:"primaryAccessKey" tf:"primary_access_key"`

	PrimaryConnectionString string `json:"primaryConnectionString" tf:"primary_connection_string"`

	PublicPort int64 `json:"publicPort" tf:"public_port"`

	SecondaryAccessKey string `json:"secondaryAccessKey" tf:"secondary_access_key"`

	SecondaryConnectionString string `json:"secondaryConnectionString" tf:"secondary_connection_string"`

	ServerPort int64 `json:"serverPort" tf:"server_port"`
}

type SignalrServiceParameters struct {
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors"`

	Features []FeaturesParameters `json:"features,omitempty" tf:"features"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Sku []SkuParameters `json:"sku" tf:"sku"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	UpstreamEndpoint []UpstreamEndpointParameters `json:"upstreamEndpoint,omitempty" tf:"upstream_endpoint"`
}

type SkuObservation struct {
}

type SkuParameters struct {
	Capacity int64 `json:"capacity" tf:"capacity"`

	Name string `json:"name" tf:"name"`
}

type UpstreamEndpointObservation struct {
}

type UpstreamEndpointParameters struct {
	CategoryPattern []string `json:"categoryPattern" tf:"category_pattern"`

	EventPattern []string `json:"eventPattern" tf:"event_pattern"`

	HubPattern []string `json:"hubPattern" tf:"hub_pattern"`

	UrlTemplate string `json:"urlTemplate" tf:"url_template"`
}

// SignalrServiceSpec defines the desired state of SignalrService
type SignalrServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SignalrServiceParameters `json:"forProvider"`
}

// SignalrServiceStatus defines the observed state of SignalrService.
type SignalrServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SignalrServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SignalrService is the Schema for the SignalrServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SignalrService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignalrServiceSpec   `json:"spec"`
	Status            SignalrServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SignalrServiceList contains a list of SignalrServices
type SignalrServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignalrService `json:"items"`
}

// Repository type metadata.
var (
	SignalrServiceKind             = "SignalrService"
	SignalrServiceGroupKind        = schema.GroupKind{Group: Group, Kind: SignalrServiceKind}.String()
	SignalrServiceKindAPIVersion   = SignalrServiceKind + "." + GroupVersion.String()
	SignalrServiceGroupVersionKind = GroupVersion.WithKind(SignalrServiceKind)
)

func init() {
	SchemeBuilder.Register(&SignalrService{}, &SignalrServiceList{})
}
