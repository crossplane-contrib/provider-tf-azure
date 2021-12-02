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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FallbackRouteObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FallbackRouteParameters_2 struct {

	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	EndpointNames []*string `json:"endpointNames" tf:"endpoint_names,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// FallbackRouteSpec defines the desired state of FallbackRoute
type FallbackRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FallbackRouteParameters_2 `json:"forProvider"`
}

// FallbackRouteStatus defines the observed state of FallbackRoute.
type FallbackRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FallbackRouteObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FallbackRoute is the Schema for the FallbackRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FallbackRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FallbackRouteSpec   `json:"spec"`
	Status            FallbackRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FallbackRouteList contains a list of FallbackRoutes
type FallbackRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FallbackRoute `json:"items"`
}

// Repository type metadata.
var (
	FallbackRoute_Kind             = "FallbackRoute"
	FallbackRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FallbackRoute_Kind}.String()
	FallbackRoute_KindAPIVersion   = FallbackRoute_Kind + "." + CRDGroupVersion.String()
	FallbackRoute_GroupVersionKind = CRDGroupVersion.WithKind(FallbackRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&FallbackRoute{}, &FallbackRouteList{})
}
