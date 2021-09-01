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

type HdinsightSparkClusterComponentVersionObservation struct {
}

type HdinsightSparkClusterComponentVersionParameters struct {
	Spark string `json:"spark" tf:"spark"`
}

type HdinsightSparkClusterGatewayObservation struct {
}

type HdinsightSparkClusterGatewayParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresAmbariObservation struct {
}

type HdinsightSparkClusterMetastoresAmbariParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresHiveObservation struct {
}

type HdinsightSparkClusterMetastoresHiveParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresObservation struct {
}

type HdinsightSparkClusterMetastoresOozieObservation struct {
}

type HdinsightSparkClusterMetastoresOozieParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresParameters struct {
	Ambari []HdinsightSparkClusterMetastoresAmbariParameters `json:"ambari,omitempty" tf:"ambari"`

	Hive []HdinsightSparkClusterMetastoresHiveParameters `json:"hive,omitempty" tf:"hive"`

	Oozie []HdinsightSparkClusterMetastoresOozieParameters `json:"oozie,omitempty" tf:"oozie"`
}

type HdinsightSparkClusterMonitorObservation struct {
}

type HdinsightSparkClusterMonitorParameters struct {
	LogAnalyticsWorkspaceId string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	PrimaryKey string `json:"primaryKey" tf:"primary_key"`
}

type HdinsightSparkClusterNetworkObservation struct {
}

type HdinsightSparkClusterNetworkParameters struct {
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction"`

	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`
}

type HdinsightSparkClusterObservation struct {
	HttpsEndpoint string `json:"httpsEndpoint" tf:"https_endpoint"`

	SshEndpoint string `json:"sshEndpoint" tf:"ssh_endpoint"`
}

type HdinsightSparkClusterParameters struct {
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	ComponentVersion []HdinsightSparkClusterComponentVersionParameters `json:"componentVersion" tf:"component_version"`

	EncryptionInTransitEnabled *bool `json:"encryptionInTransitEnabled,omitempty" tf:"encryption_in_transit_enabled"`

	Gateway []HdinsightSparkClusterGatewayParameters `json:"gateway" tf:"gateway"`

	Location string `json:"location" tf:"location"`

	Metastores []HdinsightSparkClusterMetastoresParameters `json:"metastores,omitempty" tf:"metastores"`

	Monitor []HdinsightSparkClusterMonitorParameters `json:"monitor,omitempty" tf:"monitor"`

	Name string `json:"name" tf:"name"`

	Network []HdinsightSparkClusterNetworkParameters `json:"network,omitempty" tf:"network"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Roles []HdinsightSparkClusterRolesParameters `json:"roles" tf:"roles"`

	StorageAccount []HdinsightSparkClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account"`

	StorageAccountGen2 []HdinsightSparkClusterStorageAccountGen2Parameters `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Tier string `json:"tier" tf:"tier"`

	TlsMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`
}

type HdinsightSparkClusterRolesHeadNodeObservation struct {
}

type HdinsightSparkClusterRolesHeadNodeParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightSparkClusterRolesObservation struct {
}

type HdinsightSparkClusterRolesParameters struct {
	HeadNode []HdinsightSparkClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node"`

	WorkerNode []HdinsightSparkClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node"`

	ZookeeperNode []HdinsightSparkClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightSparkClusterRolesWorkerNodeAutoscaleObservation struct {
}

type HdinsightSparkClusterRolesWorkerNodeAutoscaleParameters struct {
	Capacity []WorkerNodeAutoscaleCapacityParameters `json:"capacity,omitempty" tf:"capacity"`

	Recurrence []RolesWorkerNodeAutoscaleRecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence"`
}

type HdinsightSparkClusterRolesWorkerNodeObservation struct {
}

type HdinsightSparkClusterRolesWorkerNodeParameters struct {
	Autoscale []HdinsightSparkClusterRolesWorkerNodeAutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale"`

	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count"`

	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightSparkClusterRolesZookeeperNodeObservation struct {
}

type HdinsightSparkClusterRolesZookeeperNodeParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightSparkClusterStorageAccountGen2Observation struct {
}

type HdinsightSparkClusterStorageAccountGen2Parameters struct {
	FilesystemId string `json:"filesystemId" tf:"filesystem_id"`

	IsDefault bool `json:"isDefault" tf:"is_default"`

	ManagedIdentityResourceId string `json:"managedIdentityResourceId" tf:"managed_identity_resource_id"`

	StorageResourceId string `json:"storageResourceId" tf:"storage_resource_id"`
}

type HdinsightSparkClusterStorageAccountObservation struct {
}

type HdinsightSparkClusterStorageAccountParameters struct {
	IsDefault bool `json:"isDefault" tf:"is_default"`

	StorageAccountKey string `json:"storageAccountKey" tf:"storage_account_key"`

	StorageContainerId string `json:"storageContainerId" tf:"storage_container_id"`
}

type RolesWorkerNodeAutoscaleRecurrenceObservation struct {
}

type RolesWorkerNodeAutoscaleRecurrenceParameters struct {
	Schedule []WorkerNodeAutoscaleRecurrenceScheduleParameters `json:"schedule" tf:"schedule"`

	Timezone string `json:"timezone" tf:"timezone"`
}

type WorkerNodeAutoscaleCapacityObservation struct {
}

type WorkerNodeAutoscaleCapacityParameters struct {
	MaxInstanceCount int64 `json:"maxInstanceCount" tf:"max_instance_count"`

	MinInstanceCount int64 `json:"minInstanceCount" tf:"min_instance_count"`
}

type WorkerNodeAutoscaleRecurrenceScheduleObservation struct {
}

type WorkerNodeAutoscaleRecurrenceScheduleParameters struct {
	Days []string `json:"days" tf:"days"`

	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	Time string `json:"time" tf:"time"`
}

// HdinsightSparkClusterSpec defines the desired state of HdinsightSparkCluster
type HdinsightSparkClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HdinsightSparkClusterParameters `json:"forProvider"`
}

// HdinsightSparkClusterStatus defines the observed state of HdinsightSparkCluster.
type HdinsightSparkClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HdinsightSparkClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightSparkCluster is the Schema for the HdinsightSparkClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HdinsightSparkCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightSparkClusterSpec   `json:"spec"`
	Status            HdinsightSparkClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightSparkClusterList contains a list of HdinsightSparkClusters
type HdinsightSparkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HdinsightSparkCluster `json:"items"`
}

// Repository type metadata.
var (
	HdinsightSparkClusterKind             = "HdinsightSparkCluster"
	HdinsightSparkClusterGroupKind        = schema.GroupKind{Group: Group, Kind: HdinsightSparkClusterKind}.String()
	HdinsightSparkClusterKindAPIVersion   = HdinsightSparkClusterKind + "." + GroupVersion.String()
	HdinsightSparkClusterGroupVersionKind = GroupVersion.WithKind(HdinsightSparkClusterKind)
)

func init() {
	SchemeBuilder.Register(&HdinsightSparkCluster{}, &HdinsightSparkClusterList{})
}
