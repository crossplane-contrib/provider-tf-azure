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

type ManagedDiskObservation struct {
}

type ManagedDiskParameters struct {
	DiskId string `json:"diskId" tf:"disk_id"`

	StagingStorageAccountId string `json:"stagingStorageAccountId" tf:"staging_storage_account_id"`

	TargetDiskEncryptionSetId *string `json:"targetDiskEncryptionSetId,omitempty" tf:"target_disk_encryption_set_id"`

	TargetDiskType string `json:"targetDiskType" tf:"target_disk_type"`

	TargetReplicaDiskType string `json:"targetReplicaDiskType" tf:"target_replica_disk_type"`

	TargetResourceGroupId string `json:"targetResourceGroupId" tf:"target_resource_group_id"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {
	RecoveryPublicIpAddressId *string `json:"recoveryPublicIpAddressId,omitempty" tf:"recovery_public_ip_address_id"`

	SourceNetworkInterfaceId *string `json:"sourceNetworkInterfaceId,omitempty" tf:"source_network_interface_id"`

	TargetStaticIp *string `json:"targetStaticIp,omitempty" tf:"target_static_ip"`

	TargetSubnetName *string `json:"targetSubnetName,omitempty" tf:"target_subnet_name"`
}

type SiteRecoveryReplicatedVmObservation struct {
}

type SiteRecoveryReplicatedVmParameters struct {
	ManagedDisk []ManagedDiskParameters `json:"managedDisk,omitempty" tf:"managed_disk"`

	Name string `json:"name" tf:"name"`

	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface"`

	RecoveryReplicationPolicyId string `json:"recoveryReplicationPolicyId" tf:"recovery_replication_policy_id"`

	RecoveryVaultName string `json:"recoveryVaultName" tf:"recovery_vault_name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SourceRecoveryFabricName string `json:"sourceRecoveryFabricName" tf:"source_recovery_fabric_name"`

	SourceRecoveryProtectionContainerName string `json:"sourceRecoveryProtectionContainerName" tf:"source_recovery_protection_container_name"`

	SourceVmId string `json:"sourceVmId" tf:"source_vm_id"`

	TargetAvailabilitySetId *string `json:"targetAvailabilitySetId,omitempty" tf:"target_availability_set_id"`

	TargetNetworkId *string `json:"targetNetworkId,omitempty" tf:"target_network_id"`

	TargetRecoveryFabricId string `json:"targetRecoveryFabricId" tf:"target_recovery_fabric_id"`

	TargetRecoveryProtectionContainerId string `json:"targetRecoveryProtectionContainerId" tf:"target_recovery_protection_container_id"`

	TargetResourceGroupId string `json:"targetResourceGroupId" tf:"target_resource_group_id"`
}

// SiteRecoveryReplicatedVmSpec defines the desired state of SiteRecoveryReplicatedVm
type SiteRecoveryReplicatedVmSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SiteRecoveryReplicatedVmParameters `json:"forProvider"`
}

// SiteRecoveryReplicatedVmStatus defines the observed state of SiteRecoveryReplicatedVm.
type SiteRecoveryReplicatedVmStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SiteRecoveryReplicatedVmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryReplicatedVm is the Schema for the SiteRecoveryReplicatedVms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryReplicatedVm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryReplicatedVmSpec   `json:"spec"`
	Status            SiteRecoveryReplicatedVmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryReplicatedVmList contains a list of SiteRecoveryReplicatedVms
type SiteRecoveryReplicatedVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryReplicatedVm `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryReplicatedVmKind             = "SiteRecoveryReplicatedVm"
	SiteRecoveryReplicatedVmGroupKind        = schema.GroupKind{Group: Group, Kind: SiteRecoveryReplicatedVmKind}.String()
	SiteRecoveryReplicatedVmKindAPIVersion   = SiteRecoveryReplicatedVmKind + "." + GroupVersion.String()
	SiteRecoveryReplicatedVmGroupVersionKind = GroupVersion.WithKind(SiteRecoveryReplicatedVmKind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryReplicatedVm{}, &SiteRecoveryReplicatedVmList{})
}
