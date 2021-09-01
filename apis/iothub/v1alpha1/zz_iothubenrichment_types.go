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

type IothubEnrichmentObservation struct {
}

type IothubEnrichmentParameters struct {
	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`

	IothubName string `json:"iothubName" tf:"iothub_name"`

	Key string `json:"key" tf:"key"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Value string `json:"value" tf:"value"`
}

// IothubEnrichmentSpec defines the desired state of IothubEnrichment
type IothubEnrichmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IothubEnrichmentParameters `json:"forProvider"`
}

// IothubEnrichmentStatus defines the observed state of IothubEnrichment.
type IothubEnrichmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IothubEnrichmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IothubEnrichment is the Schema for the IothubEnrichments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IothubEnrichment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubEnrichmentSpec   `json:"spec"`
	Status            IothubEnrichmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubEnrichmentList contains a list of IothubEnrichments
type IothubEnrichmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IothubEnrichment `json:"items"`
}

// Repository type metadata.
var (
	IothubEnrichmentKind             = "IothubEnrichment"
	IothubEnrichmentGroupKind        = schema.GroupKind{Group: Group, Kind: IothubEnrichmentKind}.String()
	IothubEnrichmentKindAPIVersion   = IothubEnrichmentKind + "." + GroupVersion.String()
	IothubEnrichmentGroupVersionKind = GroupVersion.WithKind(IothubEnrichmentKind)
)

func init() {
	SchemeBuilder.Register(&IothubEnrichment{}, &IothubEnrichmentList{})
}
