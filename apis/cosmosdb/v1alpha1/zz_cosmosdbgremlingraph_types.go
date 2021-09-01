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

type CompositeIndexObservation struct {
}

type CompositeIndexParameters struct {
	Index []IndexParameters `json:"index" tf:"index"`
}

type ConflictResolutionPolicyObservation struct {
}

type ConflictResolutionPolicyParameters struct {
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path"`

	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure"`

	Mode string `json:"mode" tf:"mode"`
}

type CosmosdbGremlinGraphAutoscaleSettingsObservation struct {
}

type CosmosdbGremlinGraphAutoscaleSettingsParameters struct {
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput"`
}

type CosmosdbGremlinGraphObservation struct {
}

type CosmosdbGremlinGraphParameters struct {
	AccountName string `json:"accountName" tf:"account_name"`

	AutoscaleSettings []CosmosdbGremlinGraphAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings"`

	ConflictResolutionPolicy []ConflictResolutionPolicyParameters `json:"conflictResolutionPolicy,omitempty" tf:"conflict_resolution_policy"`

	DatabaseName string `json:"databaseName" tf:"database_name"`

	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`

	IndexPolicy []IndexPolicyParameters `json:"indexPolicy,omitempty" tf:"index_policy"`

	Name string `json:"name" tf:"name"`

	PartitionKeyPath string `json:"partitionKeyPath" tf:"partition_key_path"`

	PartitionKeyVersion *int64 `json:"partitionKeyVersion,omitempty" tf:"partition_key_version"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	UniqueKey []UniqueKeyParameters `json:"uniqueKey,omitempty" tf:"unique_key"`
}

type IndexObservation struct {
}

type IndexParameters struct {
	Order string `json:"order" tf:"order"`

	Path string `json:"path" tf:"path"`
}

type IndexPolicyObservation struct {
}

type IndexPolicyParameters struct {
	Automatic *bool `json:"automatic,omitempty" tf:"automatic"`

	CompositeIndex []CompositeIndexParameters `json:"compositeIndex,omitempty" tf:"composite_index"`

	ExcludedPaths []string `json:"excludedPaths,omitempty" tf:"excluded_paths"`

	IncludedPaths []string `json:"includedPaths,omitempty" tf:"included_paths"`

	IndexingMode string `json:"indexingMode" tf:"indexing_mode"`

	SpatialIndex []SpatialIndexParameters `json:"spatialIndex,omitempty" tf:"spatial_index"`
}

type SpatialIndexObservation struct {
	Types []string `json:"types" tf:"types"`
}

type SpatialIndexParameters struct {
	Path string `json:"path" tf:"path"`
}

type UniqueKeyObservation struct {
}

type UniqueKeyParameters struct {
	Paths []string `json:"paths" tf:"paths"`
}

// CosmosdbGremlinGraphSpec defines the desired state of CosmosdbGremlinGraph
type CosmosdbGremlinGraphSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CosmosdbGremlinGraphParameters `json:"forProvider"`
}

// CosmosdbGremlinGraphStatus defines the observed state of CosmosdbGremlinGraph.
type CosmosdbGremlinGraphStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CosmosdbGremlinGraphObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbGremlinGraph is the Schema for the CosmosdbGremlinGraphs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CosmosdbGremlinGraph struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbGremlinGraphSpec   `json:"spec"`
	Status            CosmosdbGremlinGraphStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbGremlinGraphList contains a list of CosmosdbGremlinGraphs
type CosmosdbGremlinGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosdbGremlinGraph `json:"items"`
}

// Repository type metadata.
var (
	CosmosdbGremlinGraphKind             = "CosmosdbGremlinGraph"
	CosmosdbGremlinGraphGroupKind        = schema.GroupKind{Group: Group, Kind: CosmosdbGremlinGraphKind}.String()
	CosmosdbGremlinGraphKindAPIVersion   = CosmosdbGremlinGraphKind + "." + GroupVersion.String()
	CosmosdbGremlinGraphGroupVersionKind = GroupVersion.WithKind(CosmosdbGremlinGraphKind)
)

func init() {
	SchemeBuilder.Register(&CosmosdbGremlinGraph{}, &CosmosdbGremlinGraphList{})
}
