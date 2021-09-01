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

type AutomationVariableDatetimeObservation struct {
}

type AutomationVariableDatetimeParameters struct {
	AutomationAccountName string `json:"automationAccountName" tf:"automation_account_name"`

	Description *string `json:"description,omitempty" tf:"description"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Value *string `json:"value,omitempty" tf:"value"`
}

// AutomationVariableDatetimeSpec defines the desired state of AutomationVariableDatetime
type AutomationVariableDatetimeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AutomationVariableDatetimeParameters `json:"forProvider"`
}

// AutomationVariableDatetimeStatus defines the observed state of AutomationVariableDatetime.
type AutomationVariableDatetimeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AutomationVariableDatetimeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutomationVariableDatetime is the Schema for the AutomationVariableDatetimes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AutomationVariableDatetime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationVariableDatetimeSpec   `json:"spec"`
	Status            AutomationVariableDatetimeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutomationVariableDatetimeList contains a list of AutomationVariableDatetimes
type AutomationVariableDatetimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutomationVariableDatetime `json:"items"`
}

// Repository type metadata.
var (
	AutomationVariableDatetimeKind             = "AutomationVariableDatetime"
	AutomationVariableDatetimeGroupKind        = schema.GroupKind{Group: Group, Kind: AutomationVariableDatetimeKind}.String()
	AutomationVariableDatetimeKindAPIVersion   = AutomationVariableDatetimeKind + "." + GroupVersion.String()
	AutomationVariableDatetimeGroupVersionKind = GroupVersion.WithKind(AutomationVariableDatetimeKind)
)

func init() {
	SchemeBuilder.Register(&AutomationVariableDatetime{}, &AutomationVariableDatetimeList{})
}
