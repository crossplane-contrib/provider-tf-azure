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

type DataShareObservation struct {
}

type DataShareParameters struct {
	AccountId string `json:"accountId" tf:"account_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Kind string `json:"kind" tf:"kind"`

	Name string `json:"name" tf:"name"`

	SnapshotSchedule []SnapshotScheduleParameters `json:"snapshotSchedule,omitempty" tf:"snapshot_schedule"`

	Terms *string `json:"terms,omitempty" tf:"terms"`
}

type SnapshotScheduleObservation struct {
}

type SnapshotScheduleParameters struct {
	Name string `json:"name" tf:"name"`

	Recurrence string `json:"recurrence" tf:"recurrence"`

	StartTime string `json:"startTime" tf:"start_time"`
}

// DataShareSpec defines the desired state of DataShare
type DataShareSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataShareParameters `json:"forProvider"`
}

// DataShareStatus defines the observed state of DataShare.
type DataShareStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataShare is the Schema for the DataShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataShareSpec   `json:"spec"`
	Status            DataShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareList contains a list of DataShares
type DataShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataShare `json:"items"`
}

// Repository type metadata.
var (
	DataShareKind             = "DataShare"
	DataShareGroupKind        = schema.GroupKind{Group: Group, Kind: DataShareKind}.String()
	DataShareKindAPIVersion   = DataShareKind + "." + GroupVersion.String()
	DataShareGroupVersionKind = GroupVersion.WithKind(DataShareKind)
)

func init() {
	SchemeBuilder.Register(&DataShare{}, &DataShareList{})
}
