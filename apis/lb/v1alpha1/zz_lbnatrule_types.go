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

type LbNatRuleObservation struct {
	BackendIpConfigurationId string `json:"backendIpConfigurationId" tf:"backend_ip_configuration_id"`

	FrontendIpConfigurationId string `json:"frontendIpConfigurationId" tf:"frontend_ip_configuration_id"`
}

type LbNatRuleParameters struct {
	BackendPort int64 `json:"backendPort" tf:"backend_port"`

	EnableFloatingIp *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip"`

	EnableTcpReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset"`

	FrontendIpConfigurationName string `json:"frontendIpConfigurationName" tf:"frontend_ip_configuration_name"`

	FrontendPort int64 `json:"frontendPort" tf:"frontend_port"`

	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes"`

	LoadbalancerId string `json:"loadbalancerId" tf:"loadbalancer_id"`

	Name string `json:"name" tf:"name"`

	Protocol string `json:"protocol" tf:"protocol"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

// LbNatRuleSpec defines the desired state of LbNatRule
type LbNatRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LbNatRuleParameters `json:"forProvider"`
}

// LbNatRuleStatus defines the observed state of LbNatRule.
type LbNatRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LbNatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbNatRule is the Schema for the LbNatRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LbNatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbNatRuleSpec   `json:"spec"`
	Status            LbNatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbNatRuleList contains a list of LbNatRules
type LbNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbNatRule `json:"items"`
}

// Repository type metadata.
var (
	LbNatRuleKind             = "LbNatRule"
	LbNatRuleGroupKind        = schema.GroupKind{Group: Group, Kind: LbNatRuleKind}.String()
	LbNatRuleKindAPIVersion   = LbNatRuleKind + "." + GroupVersion.String()
	LbNatRuleGroupVersionKind = GroupVersion.WithKind(LbNatRuleKind)
)

func init() {
	SchemeBuilder.Register(&LbNatRule{}, &LbNatRuleList{})
}
