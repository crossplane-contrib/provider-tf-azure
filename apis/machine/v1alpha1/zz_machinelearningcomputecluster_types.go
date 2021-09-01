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

type IdentityObservation struct {
	PrincipalId string `json:"principalId" tf:"principal_id"`

	TenantId string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {
	Type string `json:"type" tf:"type"`
}

type MachineLearningComputeClusterObservation struct {
}

type MachineLearningComputeClusterParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Identity []IdentityParameters `json:"identity" tf:"identity"`

	Location string `json:"location" tf:"location"`

	MachineLearningWorkspaceId string `json:"machineLearningWorkspaceId" tf:"machine_learning_workspace_id"`

	Name string `json:"name" tf:"name"`

	ScaleSettings []ScaleSettingsParameters `json:"scaleSettings" tf:"scale_settings"`

	SubnetResourceId *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VmPriority string `json:"vmPriority" tf:"vm_priority"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type ScaleSettingsObservation struct {
}

type ScaleSettingsParameters struct {
	MaxNodeCount int64 `json:"maxNodeCount" tf:"max_node_count"`

	MinNodeCount int64 `json:"minNodeCount" tf:"min_node_count"`

	ScaleDownNodesAfterIdleDuration string `json:"scaleDownNodesAfterIdleDuration" tf:"scale_down_nodes_after_idle_duration"`
}

// MachineLearningComputeClusterSpec defines the desired state of MachineLearningComputeCluster
type MachineLearningComputeClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MachineLearningComputeClusterParameters `json:"forProvider"`
}

// MachineLearningComputeClusterStatus defines the observed state of MachineLearningComputeCluster.
type MachineLearningComputeClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MachineLearningComputeClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MachineLearningComputeCluster is the Schema for the MachineLearningComputeClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MachineLearningComputeCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MachineLearningComputeClusterSpec   `json:"spec"`
	Status            MachineLearningComputeClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MachineLearningComputeClusterList contains a list of MachineLearningComputeClusters
type MachineLearningComputeClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineLearningComputeCluster `json:"items"`
}

// Repository type metadata.
var (
	MachineLearningComputeClusterKind             = "MachineLearningComputeCluster"
	MachineLearningComputeClusterGroupKind        = schema.GroupKind{Group: Group, Kind: MachineLearningComputeClusterKind}.String()
	MachineLearningComputeClusterKindAPIVersion   = MachineLearningComputeClusterKind + "." + GroupVersion.String()
	MachineLearningComputeClusterGroupVersionKind = GroupVersion.WithKind(MachineLearningComputeClusterKind)
)

func init() {
	SchemeBuilder.Register(&MachineLearningComputeCluster{}, &MachineLearningComputeClusterList{})
}
