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

type SentinelDataConnectorOffice365Observation struct {
}

type SentinelDataConnectorOffice365Parameters struct {
	ExchangeEnabled *bool `json:"exchangeEnabled,omitempty" tf:"exchange_enabled"`

	LogAnalyticsWorkspaceId string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	Name string `json:"name" tf:"name"`

	SharepointEnabled *bool `json:"sharepointEnabled,omitempty" tf:"sharepoint_enabled"`

	TeamsEnabled *bool `json:"teamsEnabled,omitempty" tf:"teams_enabled"`

	TenantId *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

// SentinelDataConnectorOffice365Spec defines the desired state of SentinelDataConnectorOffice365
type SentinelDataConnectorOffice365Spec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SentinelDataConnectorOffice365Parameters `json:"forProvider"`
}

// SentinelDataConnectorOffice365Status defines the observed state of SentinelDataConnectorOffice365.
type SentinelDataConnectorOffice365Status struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SentinelDataConnectorOffice365Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorOffice365 is the Schema for the SentinelDataConnectorOffice365s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelDataConnectorOffice365 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelDataConnectorOffice365Spec   `json:"spec"`
	Status            SentinelDataConnectorOffice365Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorOffice365List contains a list of SentinelDataConnectorOffice365s
type SentinelDataConnectorOffice365List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelDataConnectorOffice365 `json:"items"`
}

// Repository type metadata.
var (
	SentinelDataConnectorOffice365Kind             = "SentinelDataConnectorOffice365"
	SentinelDataConnectorOffice365GroupKind        = schema.GroupKind{Group: Group, Kind: SentinelDataConnectorOffice365Kind}.String()
	SentinelDataConnectorOffice365KindAPIVersion   = SentinelDataConnectorOffice365Kind + "." + GroupVersion.String()
	SentinelDataConnectorOffice365GroupVersionKind = GroupVersion.WithKind(SentinelDataConnectorOffice365Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelDataConnectorOffice365{}, &SentinelDataConnectorOffice365List{})
}
