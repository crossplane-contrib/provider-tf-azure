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

type EndpointObservation struct {
}

type EndpointParameters struct {
	BatchFrequencyInSeconds *int64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds"`

	ConnectionString string `json:"connectionString" tf:"connection_string"`

	ContainerName *string `json:"containerName,omitempty" tf:"container_name"`

	Encoding *string `json:"encoding,omitempty" tf:"encoding"`

	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format"`

	MaxChunkSizeInBytes *int64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	Type string `json:"type" tf:"type"`
}

type EnrichmentObservation struct {
}

type EnrichmentParameters struct {
	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`

	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type FallbackRouteObservation struct {
}

type FallbackRouteParameters struct {
	Condition *string `json:"condition,omitempty" tf:"condition"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EndpointNames []string `json:"endpointNames,omitempty" tf:"endpoint_names"`

	Source *string `json:"source,omitempty" tf:"source"`
}

type FileUploadObservation struct {
}

type FileUploadParameters struct {
	ConnectionString string `json:"connectionString" tf:"connection_string"`

	ContainerName string `json:"containerName" tf:"container_name"`

	DefaultTtl *string `json:"defaultTtl,omitempty" tf:"default_ttl"`

	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration"`

	MaxDeliveryCount *int64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count"`

	Notifications *bool `json:"notifications,omitempty" tf:"notifications"`

	SasTtl *string `json:"sasTtl,omitempty" tf:"sas_ttl"`
}

type IothubObservation struct {
	EventHubEventsEndpoint string `json:"eventHubEventsEndpoint" tf:"event_hub_events_endpoint"`

	EventHubEventsPath string `json:"eventHubEventsPath" tf:"event_hub_events_path"`

	EventHubOperationsEndpoint string `json:"eventHubOperationsEndpoint" tf:"event_hub_operations_endpoint"`

	EventHubOperationsPath string `json:"eventHubOperationsPath" tf:"event_hub_operations_path"`

	Hostname string `json:"hostname" tf:"hostname"`

	SharedAccessPolicy []SharedAccessPolicyObservation `json:"sharedAccessPolicy" tf:"shared_access_policy"`

	Type string `json:"type" tf:"type"`
}

type IothubParameters struct {
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint"`

	Enrichment []EnrichmentParameters `json:"enrichment,omitempty" tf:"enrichment"`

	EventHubPartitionCount *int64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count"`

	EventHubRetentionInDays *int64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days"`

	FallbackRoute []FallbackRouteParameters `json:"fallbackRoute,omitempty" tf:"fallback_route"`

	FileUpload []FileUploadParameters `json:"fileUpload,omitempty" tf:"file_upload"`

	IpFilterRule []IpFilterRuleParameters `json:"ipFilterRule,omitempty" tf:"ip_filter_rule"`

	Location string `json:"location" tf:"location"`

	MinTlsVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`

	Name string `json:"name" tf:"name"`

	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Route []RouteParameters `json:"route,omitempty" tf:"route"`

	Sku []SkuParameters `json:"sku" tf:"sku"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type IpFilterRuleObservation struct {
}

type IpFilterRuleParameters struct {
	Action string `json:"action" tf:"action"`

	IpMask string `json:"ipMask" tf:"ip_mask"`

	Name string `json:"name" tf:"name"`
}

type RouteObservation struct {
}

type RouteParameters struct {
	Condition *string `json:"condition,omitempty" tf:"condition"`

	Enabled bool `json:"enabled" tf:"enabled"`

	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`

	Name string `json:"name" tf:"name"`

	Source string `json:"source" tf:"source"`
}

type SharedAccessPolicyObservation struct {
	KeyName string `json:"keyName" tf:"key_name"`

	Permissions string `json:"permissions" tf:"permissions"`

	PrimaryKey string `json:"primaryKey" tf:"primary_key"`

	SecondaryKey string `json:"secondaryKey" tf:"secondary_key"`
}

type SharedAccessPolicyParameters struct {
}

type SkuObservation struct {
}

type SkuParameters struct {
	Capacity int64 `json:"capacity" tf:"capacity"`

	Name string `json:"name" tf:"name"`
}

// IothubSpec defines the desired state of Iothub
type IothubSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IothubParameters `json:"forProvider"`
}

// IothubStatus defines the observed state of Iothub.
type IothubStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IothubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Iothub is the Schema for the Iothubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Iothub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSpec   `json:"spec"`
	Status            IothubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubList contains a list of Iothubs
type IothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Iothub `json:"items"`
}

// Repository type metadata.
var (
	IothubKind             = "Iothub"
	IothubGroupKind        = schema.GroupKind{Group: Group, Kind: IothubKind}.String()
	IothubKindAPIVersion   = IothubKind + "." + GroupVersion.String()
	IothubGroupVersionKind = GroupVersion.WithKind(IothubKind)
)

func init() {
	SchemeBuilder.Register(&Iothub{}, &IothubList{})
}
