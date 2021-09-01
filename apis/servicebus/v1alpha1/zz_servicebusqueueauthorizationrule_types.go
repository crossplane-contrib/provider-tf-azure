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

type ServicebusQueueAuthorizationRuleObservation struct {
	PrimaryConnectionString string `json:"primaryConnectionString" tf:"primary_connection_string"`

	PrimaryConnectionStringAlias string `json:"primaryConnectionStringAlias" tf:"primary_connection_string_alias"`

	PrimaryKey string `json:"primaryKey" tf:"primary_key"`

	SecondaryConnectionString string `json:"secondaryConnectionString" tf:"secondary_connection_string"`

	SecondaryConnectionStringAlias string `json:"secondaryConnectionStringAlias" tf:"secondary_connection_string_alias"`

	SecondaryKey string `json:"secondaryKey" tf:"secondary_key"`
}

type ServicebusQueueAuthorizationRuleParameters struct {
	Listen *bool `json:"listen,omitempty" tf:"listen"`

	Manage *bool `json:"manage,omitempty" tf:"manage"`

	Name string `json:"name" tf:"name"`

	NamespaceName string `json:"namespaceName" tf:"namespace_name"`

	QueueName string `json:"queueName" tf:"queue_name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Send *bool `json:"send,omitempty" tf:"send"`
}

// ServicebusQueueAuthorizationRuleSpec defines the desired state of ServicebusQueueAuthorizationRule
type ServicebusQueueAuthorizationRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicebusQueueAuthorizationRuleParameters `json:"forProvider"`
}

// ServicebusQueueAuthorizationRuleStatus defines the observed state of ServicebusQueueAuthorizationRule.
type ServicebusQueueAuthorizationRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicebusQueueAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusQueueAuthorizationRule is the Schema for the ServicebusQueueAuthorizationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicebusQueueAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusQueueAuthorizationRuleSpec   `json:"spec"`
	Status            ServicebusQueueAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusQueueAuthorizationRuleList contains a list of ServicebusQueueAuthorizationRules
type ServicebusQueueAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicebusQueueAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	ServicebusQueueAuthorizationRuleKind             = "ServicebusQueueAuthorizationRule"
	ServicebusQueueAuthorizationRuleGroupKind        = schema.GroupKind{Group: Group, Kind: ServicebusQueueAuthorizationRuleKind}.String()
	ServicebusQueueAuthorizationRuleKindAPIVersion   = ServicebusQueueAuthorizationRuleKind + "." + GroupVersion.String()
	ServicebusQueueAuthorizationRuleGroupVersionKind = GroupVersion.WithKind(ServicebusQueueAuthorizationRuleKind)
)

func init() {
	SchemeBuilder.Register(&ServicebusQueueAuthorizationRule{}, &ServicebusQueueAuthorizationRuleList{})
}
