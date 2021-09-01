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

type AutoscaleCapacityObservation struct {
}

type AutoscaleCapacityParameters struct {
	MaxInstanceCount int64 `json:"maxInstanceCount" tf:"max_instance_count"`

	MinInstanceCount int64 `json:"minInstanceCount" tf:"min_instance_count"`
}

type AutoscaleRecurrenceScheduleObservation struct {
}

type AutoscaleRecurrenceScheduleParameters struct {
	Days []string `json:"days" tf:"days"`

	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	Time string `json:"time" tf:"time"`
}

type HdinsightInteractiveQueryClusterComponentVersionObservation struct {
}

type HdinsightInteractiveQueryClusterComponentVersionParameters struct {
	InteractiveHive string `json:"interactiveHive" tf:"interactive_hive"`
}

type HdinsightInteractiveQueryClusterGatewayObservation struct {
}

type HdinsightInteractiveQueryClusterGatewayParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresAmbariObservation struct {
}

type HdinsightInteractiveQueryClusterMetastoresAmbariParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresHiveObservation struct {
}

type HdinsightInteractiveQueryClusterMetastoresHiveParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresObservation struct {
}

type HdinsightInteractiveQueryClusterMetastoresOozieObservation struct {
}

type HdinsightInteractiveQueryClusterMetastoresOozieParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresParameters struct {
	Ambari []HdinsightInteractiveQueryClusterMetastoresAmbariParameters `json:"ambari,omitempty" tf:"ambari"`

	Hive []HdinsightInteractiveQueryClusterMetastoresHiveParameters `json:"hive,omitempty" tf:"hive"`

	Oozie []HdinsightInteractiveQueryClusterMetastoresOozieParameters `json:"oozie,omitempty" tf:"oozie"`
}

type HdinsightInteractiveQueryClusterMonitorObservation struct {
}

type HdinsightInteractiveQueryClusterMonitorParameters struct {
	LogAnalyticsWorkspaceId string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	PrimaryKey string `json:"primaryKey" tf:"primary_key"`
}

type HdinsightInteractiveQueryClusterNetworkObservation struct {
}

type HdinsightInteractiveQueryClusterNetworkParameters struct {
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction"`

	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`
}

type HdinsightInteractiveQueryClusterObservation struct {
	HttpsEndpoint string `json:"httpsEndpoint" tf:"https_endpoint"`

	SshEndpoint string `json:"sshEndpoint" tf:"ssh_endpoint"`
}

type HdinsightInteractiveQueryClusterParameters struct {
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	ComponentVersion []HdinsightInteractiveQueryClusterComponentVersionParameters `json:"componentVersion" tf:"component_version"`

	EncryptionInTransitEnabled *bool `json:"encryptionInTransitEnabled,omitempty" tf:"encryption_in_transit_enabled"`

	Gateway []HdinsightInteractiveQueryClusterGatewayParameters `json:"gateway" tf:"gateway"`

	Location string `json:"location" tf:"location"`

	Metastores []HdinsightInteractiveQueryClusterMetastoresParameters `json:"metastores,omitempty" tf:"metastores"`

	Monitor []HdinsightInteractiveQueryClusterMonitorParameters `json:"monitor,omitempty" tf:"monitor"`

	Name string `json:"name" tf:"name"`

	Network []HdinsightInteractiveQueryClusterNetworkParameters `json:"network,omitempty" tf:"network"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Roles []HdinsightInteractiveQueryClusterRolesParameters `json:"roles" tf:"roles"`

	StorageAccount []HdinsightInteractiveQueryClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account"`

	StorageAccountGen2 []HdinsightInteractiveQueryClusterStorageAccountGen2Parameters `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Tier string `json:"tier" tf:"tier"`

	TlsMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`
}

type HdinsightInteractiveQueryClusterRolesHeadNodeObservation struct {
}

type HdinsightInteractiveQueryClusterRolesHeadNodeParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightInteractiveQueryClusterRolesObservation struct {
}

type HdinsightInteractiveQueryClusterRolesParameters struct {
	HeadNode []HdinsightInteractiveQueryClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node"`

	WorkerNode []HdinsightInteractiveQueryClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node"`

	ZookeeperNode []HdinsightInteractiveQueryClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeObservation struct {
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeParameters struct {
	Autoscale []RolesWorkerNodeAutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale"`

	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count"`

	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightInteractiveQueryClusterRolesZookeeperNodeObservation struct {
}

type HdinsightInteractiveQueryClusterRolesZookeeperNodeParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightInteractiveQueryClusterStorageAccountGen2Observation struct {
}

type HdinsightInteractiveQueryClusterStorageAccountGen2Parameters struct {
	FilesystemId string `json:"filesystemId" tf:"filesystem_id"`

	IsDefault bool `json:"isDefault" tf:"is_default"`

	ManagedIdentityResourceId string `json:"managedIdentityResourceId" tf:"managed_identity_resource_id"`

	StorageResourceId string `json:"storageResourceId" tf:"storage_resource_id"`
}

type HdinsightInteractiveQueryClusterStorageAccountObservation struct {
}

type HdinsightInteractiveQueryClusterStorageAccountParameters struct {
	IsDefault bool `json:"isDefault" tf:"is_default"`

	StorageAccountKey string `json:"storageAccountKey" tf:"storage_account_key"`

	StorageContainerId string `json:"storageContainerId" tf:"storage_container_id"`
}

type RolesWorkerNodeAutoscaleObservation struct {
}

type RolesWorkerNodeAutoscaleParameters struct {
	Capacity []AutoscaleCapacityParameters `json:"capacity,omitempty" tf:"capacity"`

	Recurrence []WorkerNodeAutoscaleRecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence"`
}

type WorkerNodeAutoscaleRecurrenceObservation struct {
}

type WorkerNodeAutoscaleRecurrenceParameters struct {
	Schedule []AutoscaleRecurrenceScheduleParameters `json:"schedule" tf:"schedule"`

	Timezone string `json:"timezone" tf:"timezone"`
}

// HdinsightInteractiveQueryClusterSpec defines the desired state of HdinsightInteractiveQueryCluster
type HdinsightInteractiveQueryClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HdinsightInteractiveQueryClusterParameters `json:"forProvider"`
}

// HdinsightInteractiveQueryClusterStatus defines the observed state of HdinsightInteractiveQueryCluster.
type HdinsightInteractiveQueryClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HdinsightInteractiveQueryClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightInteractiveQueryCluster is the Schema for the HdinsightInteractiveQueryClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HdinsightInteractiveQueryCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightInteractiveQueryClusterSpec   `json:"spec"`
	Status            HdinsightInteractiveQueryClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightInteractiveQueryClusterList contains a list of HdinsightInteractiveQueryClusters
type HdinsightInteractiveQueryClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HdinsightInteractiveQueryCluster `json:"items"`
}

// Repository type metadata.
var (
	HdinsightInteractiveQueryClusterKind             = "HdinsightInteractiveQueryCluster"
	HdinsightInteractiveQueryClusterGroupKind        = schema.GroupKind{Group: Group, Kind: HdinsightInteractiveQueryClusterKind}.String()
	HdinsightInteractiveQueryClusterKindAPIVersion   = HdinsightInteractiveQueryClusterKind + "." + GroupVersion.String()
	HdinsightInteractiveQueryClusterGroupVersionKind = GroupVersion.WithKind(HdinsightInteractiveQueryClusterKind)
)

func init() {
	SchemeBuilder.Register(&HdinsightInteractiveQueryCluster{}, &HdinsightInteractiveQueryClusterList{})
}
