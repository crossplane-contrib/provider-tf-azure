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

type ApiManagementOpenidConnectProviderObservation struct {
}

type ApiManagementOpenidConnectProviderParameters struct {
	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`

	ClientId string `json:"clientId" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	Description *string `json:"description,omitempty" tf:"description"`

	DisplayName string `json:"displayName" tf:"display_name"`

	MetadataEndpoint string `json:"metadataEndpoint" tf:"metadata_endpoint"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

// ApiManagementOpenidConnectProviderSpec defines the desired state of ApiManagementOpenidConnectProvider
type ApiManagementOpenidConnectProviderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementOpenidConnectProviderParameters `json:"forProvider"`
}

// ApiManagementOpenidConnectProviderStatus defines the observed state of ApiManagementOpenidConnectProvider.
type ApiManagementOpenidConnectProviderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementOpenidConnectProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementOpenidConnectProvider is the Schema for the ApiManagementOpenidConnectProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementOpenidConnectProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementOpenidConnectProviderSpec   `json:"spec"`
	Status            ApiManagementOpenidConnectProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementOpenidConnectProviderList contains a list of ApiManagementOpenidConnectProviders
type ApiManagementOpenidConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementOpenidConnectProvider `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementOpenidConnectProviderKind             = "ApiManagementOpenidConnectProvider"
	ApiManagementOpenidConnectProviderGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementOpenidConnectProviderKind}.String()
	ApiManagementOpenidConnectProviderKindAPIVersion   = ApiManagementOpenidConnectProviderKind + "." + GroupVersion.String()
	ApiManagementOpenidConnectProviderGroupVersionKind = GroupVersion.WithKind(ApiManagementOpenidConnectProviderKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementOpenidConnectProvider{}, &ApiManagementOpenidConnectProviderList{})
}
