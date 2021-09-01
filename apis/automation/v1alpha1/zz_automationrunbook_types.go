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

type AutomationRunbookObservation struct {
}

type AutomationRunbookParameters struct {
	AutomationAccountName string `json:"automationAccountName" tf:"automation_account_name"`

	Content *string `json:"content,omitempty" tf:"content"`

	Description *string `json:"description,omitempty" tf:"description"`

	JobSchedule []JobScheduleParameters `json:"jobSchedule,omitempty" tf:"job_schedule"`

	Location string `json:"location" tf:"location"`

	LogProgress bool `json:"logProgress" tf:"log_progress"`

	LogVerbose bool `json:"logVerbose" tf:"log_verbose"`

	Name string `json:"name" tf:"name"`

	PublishContentLink []PublishContentLinkParameters `json:"publishContentLink,omitempty" tf:"publish_content_link"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	RunbookType string `json:"runbookType" tf:"runbook_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type JobScheduleObservation struct {
	JobScheduleId string `json:"jobScheduleId" tf:"job_schedule_id"`
}

type JobScheduleParameters struct {
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	RunOn *string `json:"runOn,omitempty" tf:"run_on"`

	ScheduleName string `json:"scheduleName" tf:"schedule_name"`
}

type PublishContentLinkHashObservation struct {
}

type PublishContentLinkHashParameters struct {
	Algorithm string `json:"algorithm" tf:"algorithm"`

	Value string `json:"value" tf:"value"`
}

type PublishContentLinkObservation struct {
}

type PublishContentLinkParameters struct {
	Hash []PublishContentLinkHashParameters `json:"hash,omitempty" tf:"hash"`

	Uri string `json:"uri" tf:"uri"`

	Version *string `json:"version,omitempty" tf:"version"`
}

// AutomationRunbookSpec defines the desired state of AutomationRunbook
type AutomationRunbookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AutomationRunbookParameters `json:"forProvider"`
}

// AutomationRunbookStatus defines the observed state of AutomationRunbook.
type AutomationRunbookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AutomationRunbookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutomationRunbook is the Schema for the AutomationRunbooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AutomationRunbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationRunbookSpec   `json:"spec"`
	Status            AutomationRunbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutomationRunbookList contains a list of AutomationRunbooks
type AutomationRunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutomationRunbook `json:"items"`
}

// Repository type metadata.
var (
	AutomationRunbookKind             = "AutomationRunbook"
	AutomationRunbookGroupKind        = schema.GroupKind{Group: Group, Kind: AutomationRunbookKind}.String()
	AutomationRunbookKindAPIVersion   = AutomationRunbookKind + "." + GroupVersion.String()
	AutomationRunbookGroupVersionKind = GroupVersion.WithKind(AutomationRunbookKind)
)

func init() {
	SchemeBuilder.Register(&AutomationRunbook{}, &AutomationRunbookList{})
}
