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

type APIReleaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIReleaseParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

// APIReleaseSpec defines the desired state of APIRelease
type APIReleaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIReleaseParameters `json:"forProvider"`
}

// APIReleaseStatus defines the observed state of APIRelease.
type APIReleaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIReleaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIRelease is the Schema for the APIReleases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type APIRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIReleaseSpec   `json:"spec"`
	Status            APIReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIReleaseList contains a list of APIReleases
type APIReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIRelease `json:"items"`
}

// Repository type metadata.
var (
	APIRelease_Kind             = "APIRelease"
	APIRelease_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIRelease_Kind}.String()
	APIRelease_KindAPIVersion   = APIRelease_Kind + "." + CRDGroupVersion.String()
	APIRelease_GroupVersionKind = CRDGroupVersion.WithKind(APIRelease_Kind)
)

func init() {
	SchemeBuilder.Register(&APIRelease{}, &APIReleaseList{})
}
