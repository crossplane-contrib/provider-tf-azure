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

type OrchestratedVirtualMachineScaleSetObservation struct {
	UniqueId string `json:"uniqueId" tf:"unique_id"`
}

type OrchestratedVirtualMachineScaleSetParameters struct {
	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	PlatformFaultDomainCount int64 `json:"platformFaultDomainCount" tf:"platform_fault_domain_count"`

	ProximityPlacementGroupId *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty" tf:"single_placement_group"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Zones []string `json:"zones,omitempty" tf:"zones"`
}

// OrchestratedVirtualMachineScaleSetSpec defines the desired state of OrchestratedVirtualMachineScaleSet
type OrchestratedVirtualMachineScaleSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OrchestratedVirtualMachineScaleSetParameters `json:"forProvider"`
}

// OrchestratedVirtualMachineScaleSetStatus defines the observed state of OrchestratedVirtualMachineScaleSet.
type OrchestratedVirtualMachineScaleSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OrchestratedVirtualMachineScaleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestratedVirtualMachineScaleSet is the Schema for the OrchestratedVirtualMachineScaleSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OrchestratedVirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrchestratedVirtualMachineScaleSetSpec   `json:"spec"`
	Status            OrchestratedVirtualMachineScaleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestratedVirtualMachineScaleSetList contains a list of OrchestratedVirtualMachineScaleSets
type OrchestratedVirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrchestratedVirtualMachineScaleSet `json:"items"`
}

// Repository type metadata.
var (
	OrchestratedVirtualMachineScaleSetKind             = "OrchestratedVirtualMachineScaleSet"
	OrchestratedVirtualMachineScaleSetGroupKind        = schema.GroupKind{Group: Group, Kind: OrchestratedVirtualMachineScaleSetKind}.String()
	OrchestratedVirtualMachineScaleSetKindAPIVersion   = OrchestratedVirtualMachineScaleSetKind + "." + GroupVersion.String()
	OrchestratedVirtualMachineScaleSetGroupVersionKind = GroupVersion.WithKind(OrchestratedVirtualMachineScaleSetKind)
)

func init() {
	SchemeBuilder.Register(&OrchestratedVirtualMachineScaleSet{}, &OrchestratedVirtualMachineScaleSetList{})
}
