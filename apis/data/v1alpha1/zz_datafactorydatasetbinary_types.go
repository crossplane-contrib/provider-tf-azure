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

type AzureBlobStorageLocationObservation struct {
}

type AzureBlobStorageLocationParameters struct {
	Container string `json:"container" tf:"container"`

	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled"`

	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled"`

	Filename *string `json:"filename,omitempty" tf:"filename"`

	Path *string `json:"path,omitempty" tf:"path"`
}

type CompressionObservation struct {
}

type CompressionParameters struct {
	Level *string `json:"level,omitempty" tf:"level"`

	Type string `json:"type" tf:"type"`
}

type DataFactoryDatasetBinaryObservation struct {
}

type DataFactoryDatasetBinaryParameters struct {
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`

	Annotations []string `json:"annotations,omitempty" tf:"annotations"`

	AzureBlobStorageLocation []AzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location"`

	Compression []CompressionParameters `json:"compression,omitempty" tf:"compression"`

	DataFactoryName string `json:"dataFactoryName" tf:"data_factory_name"`

	Description *string `json:"description,omitempty" tf:"description"`

	Folder *string `json:"folder,omitempty" tf:"folder"`

	HttpServerLocation []HttpServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location"`

	LinkedServiceName string `json:"linkedServiceName" tf:"linked_service_name"`

	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SftpServerLocation []SftpServerLocationParameters `json:"sftpServerLocation,omitempty" tf:"sftp_server_location"`
}

type HttpServerLocationObservation struct {
}

type HttpServerLocationParameters struct {
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled"`

	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled"`

	Filename string `json:"filename" tf:"filename"`

	Path string `json:"path" tf:"path"`

	RelativeUrl string `json:"relativeUrl" tf:"relative_url"`
}

type SftpServerLocationObservation struct {
}

type SftpServerLocationParameters struct {
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled"`

	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled"`

	Filename string `json:"filename" tf:"filename"`

	Path string `json:"path" tf:"path"`
}

// DataFactoryDatasetBinarySpec defines the desired state of DataFactoryDatasetBinary
type DataFactoryDatasetBinarySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataFactoryDatasetBinaryParameters `json:"forProvider"`
}

// DataFactoryDatasetBinaryStatus defines the observed state of DataFactoryDatasetBinary.
type DataFactoryDatasetBinaryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataFactoryDatasetBinaryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetBinary is the Schema for the DataFactoryDatasetBinarys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryDatasetBinary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryDatasetBinarySpec   `json:"spec"`
	Status            DataFactoryDatasetBinaryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetBinaryList contains a list of DataFactoryDatasetBinarys
type DataFactoryDatasetBinaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryDatasetBinary `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryDatasetBinaryKind             = "DataFactoryDatasetBinary"
	DataFactoryDatasetBinaryGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryDatasetBinaryKind}.String()
	DataFactoryDatasetBinaryKindAPIVersion   = DataFactoryDatasetBinaryKind + "." + GroupVersion.String()
	DataFactoryDatasetBinaryGroupVersionKind = GroupVersion.WithKind(DataFactoryDatasetBinaryKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryDatasetBinary{}, &DataFactoryDatasetBinaryList{})
}
