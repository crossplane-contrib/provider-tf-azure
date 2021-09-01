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

type DataShareDatasetDataLakeGen1Observation struct {
	DisplayName string `json:"displayName" tf:"display_name"`
}

type DataShareDatasetDataLakeGen1Parameters struct {
	DataLakeStoreId string `json:"dataLakeStoreId" tf:"data_lake_store_id"`

	DataShareId string `json:"dataShareId" tf:"data_share_id"`

	FileName *string `json:"fileName,omitempty" tf:"file_name"`

	FolderPath string `json:"folderPath" tf:"folder_path"`

	Name string `json:"name" tf:"name"`
}

// DataShareDatasetDataLakeGen1Spec defines the desired state of DataShareDatasetDataLakeGen1
type DataShareDatasetDataLakeGen1Spec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataShareDatasetDataLakeGen1Parameters `json:"forProvider"`
}

// DataShareDatasetDataLakeGen1Status defines the observed state of DataShareDatasetDataLakeGen1.
type DataShareDatasetDataLakeGen1Status struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataShareDatasetDataLakeGen1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareDatasetDataLakeGen1 is the Schema for the DataShareDatasetDataLakeGen1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataShareDatasetDataLakeGen1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataShareDatasetDataLakeGen1Spec   `json:"spec"`
	Status            DataShareDatasetDataLakeGen1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareDatasetDataLakeGen1List contains a list of DataShareDatasetDataLakeGen1s
type DataShareDatasetDataLakeGen1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataShareDatasetDataLakeGen1 `json:"items"`
}

// Repository type metadata.
var (
	DataShareDatasetDataLakeGen1Kind             = "DataShareDatasetDataLakeGen1"
	DataShareDatasetDataLakeGen1GroupKind        = schema.GroupKind{Group: Group, Kind: DataShareDatasetDataLakeGen1Kind}.String()
	DataShareDatasetDataLakeGen1KindAPIVersion   = DataShareDatasetDataLakeGen1Kind + "." + GroupVersion.String()
	DataShareDatasetDataLakeGen1GroupVersionKind = GroupVersion.WithKind(DataShareDatasetDataLakeGen1Kind)
)

func init() {
	SchemeBuilder.Register(&DataShareDatasetDataLakeGen1{}, &DataShareDatasetDataLakeGen1List{})
}
