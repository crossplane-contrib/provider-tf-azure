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

type DataFactoryCustomDatasetObservation struct {
}

type DataFactoryCustomDatasetParameters struct {
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`

	Annotations []string `json:"annotations,omitempty" tf:"annotations"`

	DataFactoryId string `json:"dataFactoryId" tf:"data_factory_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Folder *string `json:"folder,omitempty" tf:"folder"`

	LinkedService []LinkedServiceParameters `json:"linkedService" tf:"linked_service"`

	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	SchemaJson *string `json:"schemaJson,omitempty" tf:"schema_json"`

	Type string `json:"type" tf:"type"`

	TypePropertiesJson string `json:"typePropertiesJson" tf:"type_properties_json"`
}

type LinkedServiceObservation struct {
}

type LinkedServiceParameters struct {
	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`
}

// DataFactoryCustomDatasetSpec defines the desired state of DataFactoryCustomDataset
type DataFactoryCustomDatasetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataFactoryCustomDatasetParameters `json:"forProvider"`
}

// DataFactoryCustomDatasetStatus defines the observed state of DataFactoryCustomDataset.
type DataFactoryCustomDatasetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataFactoryCustomDatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryCustomDataset is the Schema for the DataFactoryCustomDatasets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryCustomDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryCustomDatasetSpec   `json:"spec"`
	Status            DataFactoryCustomDatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryCustomDatasetList contains a list of DataFactoryCustomDatasets
type DataFactoryCustomDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryCustomDataset `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryCustomDatasetKind             = "DataFactoryCustomDataset"
	DataFactoryCustomDatasetGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryCustomDatasetKind}.String()
	DataFactoryCustomDatasetKindAPIVersion   = DataFactoryCustomDatasetKind + "." + GroupVersion.String()
	DataFactoryCustomDatasetGroupVersionKind = GroupVersion.WithKind(DataFactoryCustomDatasetKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryCustomDataset{}, &DataFactoryCustomDatasetList{})
}
