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

type ServiceFabricMeshSecretValueObservation struct {
}

type ServiceFabricMeshSecretValueParameters struct {
	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ServiceFabricMeshSecretId string `json:"serviceFabricMeshSecretId" tf:"service_fabric_mesh_secret_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Value string `json:"value" tf:"value"`
}

// ServiceFabricMeshSecretValueSpec defines the desired state of ServiceFabricMeshSecretValue
type ServiceFabricMeshSecretValueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceFabricMeshSecretValueParameters `json:"forProvider"`
}

// ServiceFabricMeshSecretValueStatus defines the observed state of ServiceFabricMeshSecretValue.
type ServiceFabricMeshSecretValueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceFabricMeshSecretValueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceFabricMeshSecretValue is the Schema for the ServiceFabricMeshSecretValues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServiceFabricMeshSecretValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceFabricMeshSecretValueSpec   `json:"spec"`
	Status            ServiceFabricMeshSecretValueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceFabricMeshSecretValueList contains a list of ServiceFabricMeshSecretValues
type ServiceFabricMeshSecretValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceFabricMeshSecretValue `json:"items"`
}

// Repository type metadata.
var (
	ServiceFabricMeshSecretValueKind             = "ServiceFabricMeshSecretValue"
	ServiceFabricMeshSecretValueGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceFabricMeshSecretValueKind}.String()
	ServiceFabricMeshSecretValueKindAPIVersion   = ServiceFabricMeshSecretValueKind + "." + GroupVersion.String()
	ServiceFabricMeshSecretValueGroupVersionKind = GroupVersion.WithKind(ServiceFabricMeshSecretValueKind)
)

func init() {
	SchemeBuilder.Register(&ServiceFabricMeshSecretValue{}, &ServiceFabricMeshSecretValueList{})
}
