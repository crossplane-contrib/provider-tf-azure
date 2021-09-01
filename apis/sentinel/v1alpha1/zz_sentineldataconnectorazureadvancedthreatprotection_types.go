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

type SentinelDataConnectorAzureAdvancedThreatProtectionObservation struct {
}

type SentinelDataConnectorAzureAdvancedThreatProtectionParameters struct {
	LogAnalyticsWorkspaceId string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	Name string `json:"name" tf:"name"`

	TenantId *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

// SentinelDataConnectorAzureAdvancedThreatProtectionSpec defines the desired state of SentinelDataConnectorAzureAdvancedThreatProtection
type SentinelDataConnectorAzureAdvancedThreatProtectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SentinelDataConnectorAzureAdvancedThreatProtectionParameters `json:"forProvider"`
}

// SentinelDataConnectorAzureAdvancedThreatProtectionStatus defines the observed state of SentinelDataConnectorAzureAdvancedThreatProtection.
type SentinelDataConnectorAzureAdvancedThreatProtectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SentinelDataConnectorAzureAdvancedThreatProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorAzureAdvancedThreatProtection is the Schema for the SentinelDataConnectorAzureAdvancedThreatProtections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelDataConnectorAzureAdvancedThreatProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelDataConnectorAzureAdvancedThreatProtectionSpec   `json:"spec"`
	Status            SentinelDataConnectorAzureAdvancedThreatProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorAzureAdvancedThreatProtectionList contains a list of SentinelDataConnectorAzureAdvancedThreatProtections
type SentinelDataConnectorAzureAdvancedThreatProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelDataConnectorAzureAdvancedThreatProtection `json:"items"`
}

// Repository type metadata.
var (
	SentinelDataConnectorAzureAdvancedThreatProtectionKind             = "SentinelDataConnectorAzureAdvancedThreatProtection"
	SentinelDataConnectorAzureAdvancedThreatProtectionGroupKind        = schema.GroupKind{Group: Group, Kind: SentinelDataConnectorAzureAdvancedThreatProtectionKind}.String()
	SentinelDataConnectorAzureAdvancedThreatProtectionKindAPIVersion   = SentinelDataConnectorAzureAdvancedThreatProtectionKind + "." + GroupVersion.String()
	SentinelDataConnectorAzureAdvancedThreatProtectionGroupVersionKind = GroupVersion.WithKind(SentinelDataConnectorAzureAdvancedThreatProtectionKind)
)

func init() {
	SchemeBuilder.Register(&SentinelDataConnectorAzureAdvancedThreatProtection{}, &SentinelDataConnectorAzureAdvancedThreatProtectionList{})
}
