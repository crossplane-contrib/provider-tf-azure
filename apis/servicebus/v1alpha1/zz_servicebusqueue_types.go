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

type ServicebusQueueObservation struct {
}

type ServicebusQueueParameters struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle"`

	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration"`

	DefaultMessageTtl *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl"`

	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window"`

	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations"`

	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express"`

	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning"`

	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to"`

	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to"`

	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration"`

	MaxDeliveryCount *int64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count"`

	MaxSizeInMegabytes *int64 `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes"`

	Name string `json:"name" tf:"name"`

	NamespaceName string `json:"namespaceName" tf:"namespace_name"`

	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection"`

	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Status *string `json:"status,omitempty" tf:"status"`
}

// ServicebusQueueSpec defines the desired state of ServicebusQueue
type ServicebusQueueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicebusQueueParameters `json:"forProvider"`
}

// ServicebusQueueStatus defines the observed state of ServicebusQueue.
type ServicebusQueueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicebusQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusQueue is the Schema for the ServicebusQueues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicebusQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusQueueSpec   `json:"spec"`
	Status            ServicebusQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusQueueList contains a list of ServicebusQueues
type ServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicebusQueue `json:"items"`
}

// Repository type metadata.
var (
	ServicebusQueueKind             = "ServicebusQueue"
	ServicebusQueueGroupKind        = schema.GroupKind{Group: Group, Kind: ServicebusQueueKind}.String()
	ServicebusQueueKindAPIVersion   = ServicebusQueueKind + "." + GroupVersion.String()
	ServicebusQueueGroupVersionKind = GroupVersion.WithKind(ServicebusQueueKind)
)

func init() {
	SchemeBuilder.Register(&ServicebusQueue{}, &ServicebusQueueList{})
}
