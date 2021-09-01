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

type HdinsightStormClusterComponentVersionObservation struct {
}

type HdinsightStormClusterComponentVersionParameters struct {
	Storm string `json:"storm" tf:"storm"`
}

type HdinsightStormClusterGatewayObservation struct {
}

type HdinsightStormClusterGatewayParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

type HdinsightStormClusterMetastoresAmbariObservation struct {
}

type HdinsightStormClusterMetastoresAmbariParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightStormClusterMetastoresHiveObservation struct {
}

type HdinsightStormClusterMetastoresHiveParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightStormClusterMetastoresObservation struct {
}

type HdinsightStormClusterMetastoresOozieObservation struct {
}

type HdinsightStormClusterMetastoresOozieParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	Password string `json:"password" tf:"password"`

	Server string `json:"server" tf:"server"`

	Username string `json:"username" tf:"username"`
}

type HdinsightStormClusterMetastoresParameters struct {
	Ambari []HdinsightStormClusterMetastoresAmbariParameters `json:"ambari,omitempty" tf:"ambari"`

	Hive []HdinsightStormClusterMetastoresHiveParameters `json:"hive,omitempty" tf:"hive"`

	Oozie []HdinsightStormClusterMetastoresOozieParameters `json:"oozie,omitempty" tf:"oozie"`
}

type HdinsightStormClusterMonitorObservation struct {
}

type HdinsightStormClusterMonitorParameters struct {
	LogAnalyticsWorkspaceId string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	PrimaryKey string `json:"primaryKey" tf:"primary_key"`
}

type HdinsightStormClusterObservation struct {
	HttpsEndpoint string `json:"httpsEndpoint" tf:"https_endpoint"`

	SshEndpoint string `json:"sshEndpoint" tf:"ssh_endpoint"`
}

type HdinsightStormClusterParameters struct {
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	ComponentVersion []HdinsightStormClusterComponentVersionParameters `json:"componentVersion" tf:"component_version"`

	Gateway []HdinsightStormClusterGatewayParameters `json:"gateway" tf:"gateway"`

	Location string `json:"location" tf:"location"`

	Metastores []HdinsightStormClusterMetastoresParameters `json:"metastores,omitempty" tf:"metastores"`

	Monitor []HdinsightStormClusterMonitorParameters `json:"monitor,omitempty" tf:"monitor"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Roles []HdinsightStormClusterRolesParameters `json:"roles" tf:"roles"`

	StorageAccount []HdinsightStormClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Tier string `json:"tier" tf:"tier"`

	TlsMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`
}

type HdinsightStormClusterRolesHeadNodeObservation struct {
}

type HdinsightStormClusterRolesHeadNodeParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightStormClusterRolesObservation struct {
}

type HdinsightStormClusterRolesParameters struct {
	HeadNode []HdinsightStormClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node"`

	WorkerNode []HdinsightStormClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node"`

	ZookeeperNode []HdinsightStormClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightStormClusterRolesWorkerNodeObservation struct {
}

type HdinsightStormClusterRolesWorkerNodeParameters struct {
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count"`

	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightStormClusterRolesZookeeperNodeObservation struct {
}

type HdinsightStormClusterRolesZookeeperNodeParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Username string `json:"username" tf:"username"`

	VirtualNetworkId *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`

	VmSize string `json:"vmSize" tf:"vm_size"`
}

type HdinsightStormClusterStorageAccountObservation struct {
}

type HdinsightStormClusterStorageAccountParameters struct {
	IsDefault bool `json:"isDefault" tf:"is_default"`

	StorageAccountKey string `json:"storageAccountKey" tf:"storage_account_key"`

	StorageContainerId string `json:"storageContainerId" tf:"storage_container_id"`
}

// HdinsightStormClusterSpec defines the desired state of HdinsightStormCluster
type HdinsightStormClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HdinsightStormClusterParameters `json:"forProvider"`
}

// HdinsightStormClusterStatus defines the observed state of HdinsightStormCluster.
type HdinsightStormClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HdinsightStormClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightStormCluster is the Schema for the HdinsightStormClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HdinsightStormCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightStormClusterSpec   `json:"spec"`
	Status            HdinsightStormClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightStormClusterList contains a list of HdinsightStormClusters
type HdinsightStormClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HdinsightStormCluster `json:"items"`
}

// Repository type metadata.
var (
	HdinsightStormClusterKind             = "HdinsightStormCluster"
	HdinsightStormClusterGroupKind        = schema.GroupKind{Group: Group, Kind: HdinsightStormClusterKind}.String()
	HdinsightStormClusterKindAPIVersion   = HdinsightStormClusterKind + "." + GroupVersion.String()
	HdinsightStormClusterGroupVersionKind = GroupVersion.WithKind(HdinsightStormClusterKind)
)

func init() {
	SchemeBuilder.Register(&HdinsightStormCluster{}, &HdinsightStormClusterList{})
}
