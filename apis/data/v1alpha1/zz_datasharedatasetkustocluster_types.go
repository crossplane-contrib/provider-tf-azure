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

type DataShareDatasetKustoClusterObservation struct {
	DisplayName string `json:"displayName" tf:"display_name"`

	KustoClusterLocation string `json:"kustoClusterLocation" tf:"kusto_cluster_location"`
}

type DataShareDatasetKustoClusterParameters struct {
	KustoClusterId string `json:"kustoClusterId" tf:"kusto_cluster_id"`

	Name string `json:"name" tf:"name"`

	ShareId string `json:"shareId" tf:"share_id"`
}

// DataShareDatasetKustoClusterSpec defines the desired state of DataShareDatasetKustoCluster
type DataShareDatasetKustoClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataShareDatasetKustoClusterParameters `json:"forProvider"`
}

// DataShareDatasetKustoClusterStatus defines the observed state of DataShareDatasetKustoCluster.
type DataShareDatasetKustoClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataShareDatasetKustoClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareDatasetKustoCluster is the Schema for the DataShareDatasetKustoClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataShareDatasetKustoCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataShareDatasetKustoClusterSpec   `json:"spec"`
	Status            DataShareDatasetKustoClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareDatasetKustoClusterList contains a list of DataShareDatasetKustoClusters
type DataShareDatasetKustoClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataShareDatasetKustoCluster `json:"items"`
}

// Repository type metadata.
var (
	DataShareDatasetKustoClusterKind             = "DataShareDatasetKustoCluster"
	DataShareDatasetKustoClusterGroupKind        = schema.GroupKind{Group: Group, Kind: DataShareDatasetKustoClusterKind}.String()
	DataShareDatasetKustoClusterKindAPIVersion   = DataShareDatasetKustoClusterKind + "." + GroupVersion.String()
	DataShareDatasetKustoClusterGroupVersionKind = GroupVersion.WithKind(DataShareDatasetKustoClusterKind)
)

func init() {
	SchemeBuilder.Register(&DataShareDatasetKustoCluster{}, &DataShareDatasetKustoClusterList{})
}
