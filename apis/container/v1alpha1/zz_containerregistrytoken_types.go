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

type ContainerRegistryTokenObservation struct {
}

type ContainerRegistryTokenParameters struct {
	ContainerRegistryName string `json:"containerRegistryName" tf:"container_registry_name"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ScopeMapId string `json:"scopeMapId" tf:"scope_map_id"`
}

// ContainerRegistryTokenSpec defines the desired state of ContainerRegistryToken
type ContainerRegistryTokenSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ContainerRegistryTokenParameters `json:"forProvider"`
}

// ContainerRegistryTokenStatus defines the observed state of ContainerRegistryToken.
type ContainerRegistryTokenStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ContainerRegistryTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRegistryToken is the Schema for the ContainerRegistryTokens API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ContainerRegistryToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerRegistryTokenSpec   `json:"spec"`
	Status            ContainerRegistryTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRegistryTokenList contains a list of ContainerRegistryTokens
type ContainerRegistryTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerRegistryToken `json:"items"`
}

// Repository type metadata.
var (
	ContainerRegistryTokenKind             = "ContainerRegistryToken"
	ContainerRegistryTokenGroupKind        = schema.GroupKind{Group: Group, Kind: ContainerRegistryTokenKind}.String()
	ContainerRegistryTokenKindAPIVersion   = ContainerRegistryTokenKind + "." + GroupVersion.String()
	ContainerRegistryTokenGroupVersionKind = GroupVersion.WithKind(ContainerRegistryTokenKind)
)

func init() {
	SchemeBuilder.Register(&ContainerRegistryToken{}, &ContainerRegistryTokenList{})
}
