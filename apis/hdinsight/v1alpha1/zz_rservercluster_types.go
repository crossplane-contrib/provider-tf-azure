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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RServerClusterGatewayObservation struct {
}

type RServerClusterGatewayParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type RServerClusterObservation struct {
	EdgeSSHEndpoint *string `json:"edgeSshEndpoint,omitempty" tf:"edge_ssh_endpoint,omitempty"`

	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SSHEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint,omitempty"`
}

type RServerClusterParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// +kubebuilder:validation:Required
	Gateway []RServerClusterGatewayParameters `json:"gateway" tf:"gateway,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Roles []RServerClusterRolesParameters `json:"roles" tf:"roles,omitempty"`

	// +kubebuilder:validation:Required
	Rstudio *bool `json:"rstudio" tf:"rstudio,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccount []RServerClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	TLSMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

type RServerClusterRolesEdgeNodeObservation struct {
}

type RServerClusterRolesEdgeNodeParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type RServerClusterRolesHeadNodeObservation struct {
}

type RServerClusterRolesHeadNodeParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type RServerClusterRolesObservation struct {
}

type RServerClusterRolesParameters struct {

	// +kubebuilder:validation:Required
	EdgeNode []RServerClusterRolesEdgeNodeParameters `json:"edgeNode" tf:"edge_node,omitempty"`

	// +kubebuilder:validation:Required
	HeadNode []RServerClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node,omitempty"`

	// +kubebuilder:validation:Required
	WorkerNode []RServerClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node,omitempty"`

	// +kubebuilder:validation:Required
	ZookeeperNode []RServerClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node,omitempty"`
}

type RServerClusterRolesWorkerNodeObservation struct {
}

type RServerClusterRolesWorkerNodeParameters struct {

	// +kubebuilder:validation:Optional
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TargetInstanceCount *int64 `json:"targetInstanceCount" tf:"target_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type RServerClusterRolesZookeeperNodeObservation struct {
}

type RServerClusterRolesZookeeperNodeParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type RServerClusterStorageAccountObservation struct {
}

type RServerClusterStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageContainerID *string `json:"storageContainerId" tf:"storage_container_id,omitempty"`
}

// RServerClusterSpec defines the desired state of RServerCluster
type RServerClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RServerClusterParameters `json:"forProvider"`
}

// RServerClusterStatus defines the observed state of RServerCluster.
type RServerClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RServerClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RServerCluster is the Schema for the RServerClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RServerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RServerClusterSpec   `json:"spec"`
	Status            RServerClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RServerClusterList contains a list of RServerClusters
type RServerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RServerCluster `json:"items"`
}

// Repository type metadata.
var (
	RServerCluster_Kind             = "RServerCluster"
	RServerCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RServerCluster_Kind}.String()
	RServerCluster_KindAPIVersion   = RServerCluster_Kind + "." + CRDGroupVersion.String()
	RServerCluster_GroupVersionKind = CRDGroupVersion.WithKind(RServerCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&RServerCluster{}, &RServerClusterList{})
}
