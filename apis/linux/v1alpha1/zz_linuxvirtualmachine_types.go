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

type AdditionalCapabilitiesObservation struct {
}

type AdditionalCapabilitiesParameters struct {
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled"`
}

type AdminSshKeyObservation struct {
}

type AdminSshKeyParameters struct {
	PublicKey string `json:"publicKey" tf:"public_key"`

	Username string `json:"username" tf:"username"`
}

type BootDiagnosticsObservation struct {
}

type BootDiagnosticsParameters struct {
	StorageAccountUri *string `json:"storageAccountUri,omitempty" tf:"storage_account_uri"`
}

type CertificateObservation struct {
}

type CertificateParameters struct {
	Url string `json:"url" tf:"url"`
}

type DiffDiskSettingsObservation struct {
}

type DiffDiskSettingsParameters struct {
	Option string `json:"option" tf:"option"`
}

type IdentityObservation struct {
	PrincipalId string `json:"principalId" tf:"principal_id"`

	TenantId string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	Type string `json:"type" tf:"type"`
}

type LinuxVirtualMachineObservation struct {
	PrivateIpAddress string `json:"privateIpAddress" tf:"private_ip_address"`

	PrivateIpAddresses []string `json:"privateIpAddresses" tf:"private_ip_addresses"`

	PublicIpAddress string `json:"publicIpAddress" tf:"public_ip_address"`

	PublicIpAddresses []string `json:"publicIpAddresses" tf:"public_ip_addresses"`

	VirtualMachineId string `json:"virtualMachineId" tf:"virtual_machine_id"`
}

type LinuxVirtualMachineParameters struct {
	AdditionalCapabilities []AdditionalCapabilitiesParameters `json:"additionalCapabilities,omitempty" tf:"additional_capabilities"`

	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password"`

	AdminSshKey []AdminSshKeyParameters `json:"adminSshKey,omitempty" tf:"admin_ssh_key"`

	AdminUsername string `json:"adminUsername" tf:"admin_username"`

	AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty" tf:"allow_extension_operations"`

	AvailabilitySetId *string `json:"availabilitySetId,omitempty" tf:"availability_set_id"`

	BootDiagnostics []BootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics"`

	ComputerName *string `json:"computerName,omitempty" tf:"computer_name"`

	CustomData *string `json:"customData,omitempty" tf:"custom_data"`

	DedicatedHostId *string `json:"dedicatedHostId,omitempty" tf:"dedicated_host_id"`

	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication"`

	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled"`

	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`

	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`

	Location string `json:"location" tf:"location"`

	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price"`

	Name string `json:"name" tf:"name"`

	NetworkInterfaceIds []string `json:"networkInterfaceIds" tf:"network_interface_ids"`

	OsDisk []OsDiskParameters `json:"osDisk" tf:"os_disk"`

	Plan []PlanParameters `json:"plan,omitempty" tf:"plan"`

	PlatformFaultDomain *int64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain"`

	Priority *string `json:"priority,omitempty" tf:"priority"`

	ProvisionVmAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent"`

	ProximityPlacementGroupId *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Secret []SecretParameters `json:"secret,omitempty" tf:"secret"`

	Size string `json:"size" tf:"size"`

	SourceImageId *string `json:"sourceImageId,omitempty" tf:"source_image_id"`

	SourceImageReference []SourceImageReferenceParameters `json:"sourceImageReference,omitempty" tf:"source_image_reference"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VirtualMachineScaleSetId *string `json:"virtualMachineScaleSetId,omitempty" tf:"virtual_machine_scale_set_id"`

	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type OsDiskObservation struct {
}

type OsDiskParameters struct {
	Caching string `json:"caching" tf:"caching"`

	DiffDiskSettings []DiffDiskSettingsParameters `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings"`

	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id"`

	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`

	Name *string `json:"name,omitempty" tf:"name"`

	StorageAccountType string `json:"storageAccountType" tf:"storage_account_type"`

	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled"`
}

type PlanObservation struct {
}

type PlanParameters struct {
	Name string `json:"name" tf:"name"`

	Product string `json:"product" tf:"product"`

	Publisher string `json:"publisher" tf:"publisher"`
}

type SecretObservation struct {
}

type SecretParameters struct {
	Certificate []CertificateParameters `json:"certificate" tf:"certificate"`

	KeyVaultId string `json:"keyVaultId" tf:"key_vault_id"`
}

type SourceImageReferenceObservation struct {
}

type SourceImageReferenceParameters struct {
	Offer string `json:"offer" tf:"offer"`

	Publisher string `json:"publisher" tf:"publisher"`

	Sku string `json:"sku" tf:"sku"`

	Version string `json:"version" tf:"version"`
}

// LinuxVirtualMachineSpec defines the desired state of LinuxVirtualMachine
type LinuxVirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LinuxVirtualMachineParameters `json:"forProvider"`
}

// LinuxVirtualMachineStatus defines the observed state of LinuxVirtualMachine.
type LinuxVirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LinuxVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinuxVirtualMachine is the Schema for the LinuxVirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinuxVirtualMachineSpec   `json:"spec"`
	Status            LinuxVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinuxVirtualMachineList contains a list of LinuxVirtualMachines
type LinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinuxVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	LinuxVirtualMachineKind             = "LinuxVirtualMachine"
	LinuxVirtualMachineGroupKind        = schema.GroupKind{Group: Group, Kind: LinuxVirtualMachineKind}.String()
	LinuxVirtualMachineKindAPIVersion   = LinuxVirtualMachineKind + "." + GroupVersion.String()
	LinuxVirtualMachineGroupVersionKind = GroupVersion.WithKind(LinuxVirtualMachineKind)
)

func init() {
	SchemeBuilder.Register(&LinuxVirtualMachine{}, &LinuxVirtualMachineList{})
}
