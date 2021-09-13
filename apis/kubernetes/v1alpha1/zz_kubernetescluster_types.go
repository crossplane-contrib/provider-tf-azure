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

type AciConnectorLinuxObservation struct {
}

type AciConnectorLinuxParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Optional
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name"`
}

type AddonProfileObservation struct {
}

type AddonProfileParameters struct {

	// +kubebuilder:validation:Optional
	AciConnectorLinux []AciConnectorLinuxParameters `json:"aciConnectorLinux,omitempty" tf:"aci_connector_linux"`

	// +kubebuilder:validation:Optional
	AzurePolicy []AzurePolicyParameters `json:"azurePolicy,omitempty" tf:"azure_policy"`

	// +kubebuilder:validation:Optional
	HTTPApplicationRouting []HTTPApplicationRoutingParameters `json:"httpApplicationRouting,omitempty" tf:"http_application_routing"`

	// +kubebuilder:validation:Optional
	IngressApplicationGateway []IngressApplicationGatewayParameters `json:"ingressApplicationGateway,omitempty" tf:"ingress_application_gateway"`

	// +kubebuilder:validation:Optional
	KubeDashboard []KubeDashboardParameters `json:"kubeDashboard,omitempty" tf:"kube_dashboard"`

	// +kubebuilder:validation:Optional
	OmsAgent []OmsAgentParameters `json:"omsAgent,omitempty" tf:"oms_agent"`
}

type AllowedObservation struct {
}

type AllowedParameters struct {

	// +kubebuilder:validation:Required
	Day string `json:"day" tf:"day"`

	// +kubebuilder:validation:Required
	Hours []int64 `json:"hours" tf:"hours"`
}

type AutoScalerProfileObservation struct {
}

type AutoScalerProfileParameters struct {

	// +kubebuilder:validation:Optional
	BalanceSimilarNodeGroups *bool `json:"balanceSimilarNodeGroups,omitempty" tf:"balance_similar_node_groups"`

	// +kubebuilder:validation:Optional
	EmptyBulkDeleteMax *string `json:"emptyBulkDeleteMax,omitempty" tf:"empty_bulk_delete_max"`

	// +kubebuilder:validation:Optional
	Expander *string `json:"expander,omitempty" tf:"expander"`

	// +kubebuilder:validation:Optional
	MaxGracefulTerminationSec *string `json:"maxGracefulTerminationSec,omitempty" tf:"max_graceful_termination_sec"`

	// +kubebuilder:validation:Optional
	MaxNodeProvisioningTime *string `json:"maxNodeProvisioningTime,omitempty" tf:"max_node_provisioning_time"`

	// +kubebuilder:validation:Optional
	MaxUnreadyNodes *int64 `json:"maxUnreadyNodes,omitempty" tf:"max_unready_nodes"`

	// +kubebuilder:validation:Optional
	MaxUnreadyPercentage *float64 `json:"maxUnreadyPercentage,omitempty" tf:"max_unready_percentage"`

	// +kubebuilder:validation:Optional
	NewPodScaleUpDelay *string `json:"newPodScaleUpDelay,omitempty" tf:"new_pod_scale_up_delay"`

	// +kubebuilder:validation:Optional
	ScaleDownDelayAfterAdd *string `json:"scaleDownDelayAfterAdd,omitempty" tf:"scale_down_delay_after_add"`

	// +kubebuilder:validation:Optional
	ScaleDownDelayAfterDelete *string `json:"scaleDownDelayAfterDelete,omitempty" tf:"scale_down_delay_after_delete"`

	// +kubebuilder:validation:Optional
	ScaleDownDelayAfterFailure *string `json:"scaleDownDelayAfterFailure,omitempty" tf:"scale_down_delay_after_failure"`

	// +kubebuilder:validation:Optional
	ScaleDownUnneeded *string `json:"scaleDownUnneeded,omitempty" tf:"scale_down_unneeded"`

	// +kubebuilder:validation:Optional
	ScaleDownUnready *string `json:"scaleDownUnready,omitempty" tf:"scale_down_unready"`

	// +kubebuilder:validation:Optional
	ScaleDownUtilizationThreshold *string `json:"scaleDownUtilizationThreshold,omitempty" tf:"scale_down_utilization_threshold"`

	// +kubebuilder:validation:Optional
	ScanInterval *string `json:"scanInterval,omitempty" tf:"scan_interval"`

	// +kubebuilder:validation:Optional
	SkipNodesWithLocalStorage *bool `json:"skipNodesWithLocalStorage,omitempty" tf:"skip_nodes_with_local_storage"`

	// +kubebuilder:validation:Optional
	SkipNodesWithSystemPods *bool `json:"skipNodesWithSystemPods,omitempty" tf:"skip_nodes_with_system_pods"`
}

type AzureActiveDirectoryObservation struct {
}

type AzureActiveDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	AdminGroupObjectIds []string `json:"adminGroupObjectIds,omitempty" tf:"admin_group_object_ids"`

	// +kubebuilder:validation:Optional
	AzureRbacEnabled *bool `json:"azureRbacEnabled,omitempty" tf:"azure_rbac_enabled"`

	// +kubebuilder:validation:Optional
	ClientAppID *string `json:"clientAppId,omitempty" tf:"client_app_id"`

	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed"`

	// +kubebuilder:validation:Optional
	ServerAppID *string `json:"serverAppId,omitempty" tf:"server_app_id"`

	// +kubebuilder:validation:Optional
	ServerAppSecret *string `json:"serverAppSecret,omitempty" tf:"server_app_secret"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type AzurePolicyObservation struct {
}

type AzurePolicyParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`
}

type DefaultNodePoolObservation struct {
}

type DefaultNodePoolParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`

	// +kubebuilder:validation:Optional
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling"`

	// +kubebuilder:validation:Optional
	EnableHostEncryption *bool `json:"enableHostEncryption,omitempty" tf:"enable_host_encryption"`

	// +kubebuilder:validation:Optional
	EnableNodePublicIP *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip"`

	// +kubebuilder:validation:Optional
	FipsEnabled *bool `json:"fipsEnabled,omitempty" tf:"fips_enabled"`

	// +kubebuilder:validation:Optional
	KubeletConfig []KubeletConfigParameters `json:"kubeletConfig,omitempty" tf:"kubelet_config"`

	// +kubebuilder:validation:Optional
	KubeletDiskType *string `json:"kubeletDiskType,omitempty" tf:"kubelet_disk_type"`

	// +kubebuilder:validation:Optional
	LinuxOsConfig []LinuxOsConfigParameters `json:"linuxOsConfig,omitempty" tf:"linux_os_config"`

	// +kubebuilder:validation:Optional
	MaxCount *int64 `json:"maxCount,omitempty" tf:"max_count"`

	// +kubebuilder:validation:Optional
	MaxPods *int64 `json:"maxPods,omitempty" tf:"max_pods"`

	// +kubebuilder:validation:Optional
	MinCount *int64 `json:"minCount,omitempty" tf:"min_count"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`

	// +kubebuilder:validation:Optional
	NodeLabels map[string]string `json:"nodeLabels,omitempty" tf:"node_labels"`

	// +kubebuilder:validation:Optional
	NodePublicIPPrefixID *string `json:"nodePublicIpPrefixId,omitempty" tf:"node_public_ip_prefix_id"`

	// +kubebuilder:validation:Optional
	NodeTaints []string `json:"nodeTaints,omitempty" tf:"node_taints"`

	// +kubebuilder:validation:Optional
	OnlyCriticalAddonsEnabled *bool `json:"onlyCriticalAddonsEnabled,omitempty" tf:"only_critical_addons_enabled"`

	// +kubebuilder:validation:Optional
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty" tf:"orchestrator_version"`

	// +kubebuilder:validation:Optional
	OsDiskSizeGb *int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb"`

	// +kubebuilder:validation:Optional
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type"`

	// +kubebuilder:validation:Optional
	PodSubnetID *string `json:"podSubnetId,omitempty" tf:"pod_subnet_id"`

	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`

	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled"`

	// +kubebuilder:validation:Optional
	UpgradeSettings []UpgradeSettingsParameters `json:"upgradeSettings,omitempty" tf:"upgrade_settings"`

	// +kubebuilder:validation:Required
	VMSize string `json:"vmSize" tf:"vm_size"`

	// +kubebuilder:validation:Optional
	VnetSubnetID *string `json:"vnetSubnetId,omitempty" tf:"vnet_subnet_id"`
}

type HTTPApplicationRoutingObservation struct {
	HTTPApplicationRoutingZoneName string `json:"httpApplicationRoutingZoneName" tf:"http_application_routing_zone_name"`
}

type HTTPApplicationRoutingParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`
}

type IdentityObservation struct {
	PrincipalID string `json:"principalId" tf:"principal_id"`

	TenantID string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id"`
}

type IngressApplicationGatewayIdentityObservation struct {
	ClientID string `json:"clientId" tf:"client_id"`

	ObjectID string `json:"objectId" tf:"object_id"`

	UserAssignedIdentityID string `json:"userAssignedIdentityId" tf:"user_assigned_identity_id"`
}

type IngressApplicationGatewayIdentityParameters struct {
}

type IngressApplicationGatewayObservation struct {
	EffectiveGatewayID string `json:"effectiveGatewayId" tf:"effective_gateway_id"`

	IngressApplicationGatewayIdentity []IngressApplicationGatewayIdentityObservation `json:"ingressApplicationGatewayIdentity" tf:"ingress_application_gateway_identity"`
}

type IngressApplicationGatewayParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	// +kubebuilder:validation:Optional
	GatewayName *string `json:"gatewayName,omitempty" tf:"gateway_name"`

	// +kubebuilder:validation:Optional
	SubnetCidr *string `json:"subnetCidr,omitempty" tf:"subnet_cidr"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`
}

type KubeAdminConfigObservation struct {
	ClientCertificate string `json:"clientCertificate" tf:"client_certificate"`

	ClientKey string `json:"clientKey" tf:"client_key"`

	ClusterCaCertificate string `json:"clusterCaCertificate" tf:"cluster_ca_certificate"`

	Host string `json:"host" tf:"host"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

type KubeAdminConfigParameters struct {
}

type KubeConfigObservation struct {
	ClientCertificate string `json:"clientCertificate" tf:"client_certificate"`

	ClientKey string `json:"clientKey" tf:"client_key"`

	ClusterCaCertificate string `json:"clusterCaCertificate" tf:"cluster_ca_certificate"`

	Host string `json:"host" tf:"host"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

type KubeConfigParameters struct {
}

type KubeDashboardObservation struct {
}

type KubeDashboardParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`
}

type KubeletConfigObservation struct {
}

type KubeletConfigParameters struct {

	// +kubebuilder:validation:Optional
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls,omitempty" tf:"allowed_unsafe_sysctls"`

	// +kubebuilder:validation:Optional
	CPUCfsQuotaEnabled *bool `json:"cpuCfsQuotaEnabled,omitempty" tf:"cpu_cfs_quota_enabled"`

	// +kubebuilder:validation:Optional
	CPUCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty" tf:"cpu_cfs_quota_period"`

	// +kubebuilder:validation:Optional
	CPUManagerPolicy *string `json:"cpuManagerPolicy,omitempty" tf:"cpu_manager_policy"`

	// +kubebuilder:validation:Optional
	ContainerLogMaxLine *int64 `json:"containerLogMaxLine,omitempty" tf:"container_log_max_line"`

	// +kubebuilder:validation:Optional
	ContainerLogMaxSizeMb *int64 `json:"containerLogMaxSizeMb,omitempty" tf:"container_log_max_size_mb"`

	// +kubebuilder:validation:Optional
	ImageGcHighThreshold *int64 `json:"imageGcHighThreshold,omitempty" tf:"image_gc_high_threshold"`

	// +kubebuilder:validation:Optional
	ImageGcLowThreshold *int64 `json:"imageGcLowThreshold,omitempty" tf:"image_gc_low_threshold"`

	// +kubebuilder:validation:Optional
	PodMaxPid *int64 `json:"podMaxPid,omitempty" tf:"pod_max_pid"`

	// +kubebuilder:validation:Optional
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty" tf:"topology_manager_policy"`
}

type KubeletIdentityObservation struct {
}

type KubeletIdentityParameters struct {

	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id"`

	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id"`

	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id"`
}

type KubernetesClusterObservation struct {
	Fqdn string `json:"fqdn" tf:"fqdn"`

	KubeAdminConfig []KubeAdminConfigObservation `json:"kubeAdminConfig" tf:"kube_admin_config"`

	KubeAdminConfigRaw string `json:"kubeAdminConfigRaw" tf:"kube_admin_config_raw"`

	KubeConfig []KubeConfigObservation `json:"kubeConfig" tf:"kube_config"`

	KubeConfigRaw string `json:"kubeConfigRaw" tf:"kube_config_raw"`

	PrivateFqdn string `json:"privateFqdn" tf:"private_fqdn"`
}

type KubernetesClusterParameters struct {

	// +kubebuilder:validation:Optional
	APIServerAuthorizedIPRanges []string `json:"apiServerAuthorizedIpRanges,omitempty" tf:"api_server_authorized_ip_ranges"`

	// +kubebuilder:validation:Optional
	AddonProfile []AddonProfileParameters `json:"addonProfile,omitempty" tf:"addon_profile"`

	// +kubebuilder:validation:Optional
	AutoScalerProfile []AutoScalerProfileParameters `json:"autoScalerProfile,omitempty" tf:"auto_scaler_profile"`

	// +kubebuilder:validation:Optional
	AutomaticChannelUpgrade *string `json:"automaticChannelUpgrade,omitempty" tf:"automatic_channel_upgrade"`

	// +kubebuilder:validation:Optional
	DNSPrefix *string `json:"dnsPrefix,omitempty" tf:"dns_prefix"`

	// +kubebuilder:validation:Optional
	DNSPrefixPrivateCluster *string `json:"dnsPrefixPrivateCluster,omitempty" tf:"dns_prefix_private_cluster"`

	// +kubebuilder:validation:Required
	DefaultNodePool []DefaultNodePoolParameters `json:"defaultNodePool" tf:"default_node_pool"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id"`

	// +kubebuilder:validation:Optional
	EnablePodSecurityPolicy *bool `json:"enablePodSecurityPolicy,omitempty" tf:"enable_pod_security_policy"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	// +kubebuilder:validation:Optional
	KubeletIdentity []KubeletIdentityParameters `json:"kubeletIdentity,omitempty" tf:"kubelet_identity"`

	// +kubebuilder:validation:Optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version"`

	// +kubebuilder:validation:Optional
	LinuxProfile []LinuxProfileParameters `json:"linuxProfile,omitempty" tf:"linux_profile"`

	// +kubebuilder:validation:Optional
	LocalAccountDisabled *bool `json:"localAccountDisabled,omitempty" tf:"local_account_disabled"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NetworkProfile []NetworkProfileParameters `json:"networkProfile,omitempty" tf:"network_profile"`

	// +kubebuilder:validation:Optional
	NodeResourceGroup *string `json:"nodeResourceGroup,omitempty" tf:"node_resource_group"`

	// +kubebuilder:validation:Optional
	PrivateClusterEnabled *bool `json:"privateClusterEnabled,omitempty" tf:"private_cluster_enabled"`

	// +kubebuilder:validation:Optional
	PrivateClusterPublicFqdnEnabled *bool `json:"privateClusterPublicFqdnEnabled,omitempty" tf:"private_cluster_public_fqdn_enabled"`

	// +kubebuilder:validation:Optional
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id"`

	// +kubebuilder:validation:Optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	RoleBasedAccessControl []RoleBasedAccessControlParameters `json:"roleBasedAccessControl,omitempty" tf:"role_based_access_control"`

	// +kubebuilder:validation:Optional
	ServicePrincipal []ServicePrincipalParameters `json:"servicePrincipal,omitempty" tf:"service_principal"`

	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	WindowsProfile []WindowsProfileParameters `json:"windowsProfile,omitempty" tf:"windows_profile"`
}

type LinuxOsConfigObservation struct {
}

type LinuxOsConfigParameters struct {

	// +kubebuilder:validation:Optional
	SwapFileSizeMb *int64 `json:"swapFileSizeMb,omitempty" tf:"swap_file_size_mb"`

	// +kubebuilder:validation:Optional
	SysctlConfig []SysctlConfigParameters `json:"sysctlConfig,omitempty" tf:"sysctl_config"`

	// +kubebuilder:validation:Optional
	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty" tf:"transparent_huge_page_defrag"`

	// +kubebuilder:validation:Optional
	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty" tf:"transparent_huge_page_enabled"`
}

type LinuxProfileObservation struct {
}

type LinuxProfileParameters struct {

	// +kubebuilder:validation:Required
	AdminUsername string `json:"adminUsername" tf:"admin_username"`

	// +kubebuilder:validation:Required
	SSHKey []SSHKeyParameters `json:"sshKey" tf:"ssh_key"`
}

type LoadBalancerProfileObservation struct {
	EffectiveOutboundIps []string `json:"effectiveOutboundIps" tf:"effective_outbound_ips"`
}

type LoadBalancerProfileParameters struct {

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes"`

	// +kubebuilder:validation:Optional
	ManagedOutboundIPCount *int64 `json:"managedOutboundIpCount,omitempty" tf:"managed_outbound_ip_count"`

	// +kubebuilder:validation:Optional
	OutboundIPAddressIds []string `json:"outboundIpAddressIds,omitempty" tf:"outbound_ip_address_ids"`

	// +kubebuilder:validation:Optional
	OutboundIPPrefixIds []string `json:"outboundIpPrefixIds,omitempty" tf:"outbound_ip_prefix_ids"`

	// +kubebuilder:validation:Optional
	OutboundPortsAllocated *int64 `json:"outboundPortsAllocated,omitempty" tf:"outbound_ports_allocated"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	Allowed []AllowedParameters `json:"allowed,omitempty" tf:"allowed"`

	// +kubebuilder:validation:Optional
	NotAllowed []NotAllowedParameters `json:"notAllowed,omitempty" tf:"not_allowed"`
}

type NetworkProfileObservation struct {
}

type NetworkProfileParameters struct {

	// +kubebuilder:validation:Optional
	DNSServiceIP *string `json:"dnsServiceIp,omitempty" tf:"dns_service_ip"`

	// +kubebuilder:validation:Optional
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty" tf:"docker_bridge_cidr"`

	// +kubebuilder:validation:Optional
	LoadBalancerProfile []LoadBalancerProfileParameters `json:"loadBalancerProfile,omitempty" tf:"load_balancer_profile"`

	// +kubebuilder:validation:Optional
	LoadBalancerSku *string `json:"loadBalancerSku,omitempty" tf:"load_balancer_sku"`

	// +kubebuilder:validation:Optional
	NetworkMode *string `json:"networkMode,omitempty" tf:"network_mode"`

	// +kubebuilder:validation:Required
	NetworkPlugin string `json:"networkPlugin" tf:"network_plugin"`

	// +kubebuilder:validation:Optional
	NetworkPolicy *string `json:"networkPolicy,omitempty" tf:"network_policy"`

	// +kubebuilder:validation:Optional
	OutboundType *string `json:"outboundType,omitempty" tf:"outbound_type"`

	// +kubebuilder:validation:Optional
	PodCidr *string `json:"podCidr,omitempty" tf:"pod_cidr"`

	// +kubebuilder:validation:Optional
	ServiceCidr *string `json:"serviceCidr,omitempty" tf:"service_cidr"`
}

type NotAllowedObservation struct {
}

type NotAllowedParameters struct {

	// +kubebuilder:validation:Required
	End string `json:"end" tf:"end"`

	// +kubebuilder:validation:Required
	Start string `json:"start" tf:"start"`
}

type OmsAgentIdentityObservation struct {
	ClientID string `json:"clientId" tf:"client_id"`

	ObjectID string `json:"objectId" tf:"object_id"`

	UserAssignedIdentityID string `json:"userAssignedIdentityId" tf:"user_assigned_identity_id"`
}

type OmsAgentIdentityParameters struct {
}

type OmsAgentObservation struct {
	OmsAgentIdentity []OmsAgentIdentityObservation `json:"omsAgentIdentity" tf:"oms_agent_identity"`
}

type OmsAgentParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id"`
}

type RoleBasedAccessControlObservation struct {
}

type RoleBasedAccessControlParameters struct {

	// +kubebuilder:validation:Optional
	AzureActiveDirectory []AzureActiveDirectoryParameters `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory"`

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`
}

type SSHKeyObservation struct {
}

type SSHKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyData string `json:"keyData" tf:"key_data"`
}

type ServicePrincipalObservation struct {
}

type ServicePrincipalParameters struct {

	// +kubebuilder:validation:Required
	ClientID string `json:"clientId" tf:"client_id"`

	// +kubebuilder:validation:Required
	ClientSecret string `json:"clientSecret" tf:"client_secret"`
}

type SysctlConfigObservation struct {
}

type SysctlConfigParameters struct {

	// +kubebuilder:validation:Optional
	FsAioMaxNr *int64 `json:"fsAioMaxNr,omitempty" tf:"fs_aio_max_nr"`

	// +kubebuilder:validation:Optional
	FsFileMax *int64 `json:"fsFileMax,omitempty" tf:"fs_file_max"`

	// +kubebuilder:validation:Optional
	FsInotifyMaxUserWatches *int64 `json:"fsInotifyMaxUserWatches,omitempty" tf:"fs_inotify_max_user_watches"`

	// +kubebuilder:validation:Optional
	FsNrOpen *int64 `json:"fsNrOpen,omitempty" tf:"fs_nr_open"`

	// +kubebuilder:validation:Optional
	KernelThreadsMax *int64 `json:"kernelThreadsMax,omitempty" tf:"kernel_threads_max"`

	// +kubebuilder:validation:Optional
	NetCoreNetdevMaxBacklog *int64 `json:"netCoreNetdevMaxBacklog,omitempty" tf:"net_core_netdev_max_backlog"`

	// +kubebuilder:validation:Optional
	NetCoreOptmemMax *int64 `json:"netCoreOptmemMax,omitempty" tf:"net_core_optmem_max"`

	// +kubebuilder:validation:Optional
	NetCoreRmemDefault *int64 `json:"netCoreRmemDefault,omitempty" tf:"net_core_rmem_default"`

	// +kubebuilder:validation:Optional
	NetCoreRmemMax *int64 `json:"netCoreRmemMax,omitempty" tf:"net_core_rmem_max"`

	// +kubebuilder:validation:Optional
	NetCoreSomaxconn *int64 `json:"netCoreSomaxconn,omitempty" tf:"net_core_somaxconn"`

	// +kubebuilder:validation:Optional
	NetCoreWmemDefault *int64 `json:"netCoreWmemDefault,omitempty" tf:"net_core_wmem_default"`

	// +kubebuilder:validation:Optional
	NetCoreWmemMax *int64 `json:"netCoreWmemMax,omitempty" tf:"net_core_wmem_max"`

	// +kubebuilder:validation:Optional
	NetIPv4IPLocalPortRangeMax *int64 `json:"netIpv4IpLocalPortRangeMax,omitempty" tf:"net_ipv4_ip_local_port_range_max"`

	// +kubebuilder:validation:Optional
	NetIPv4IPLocalPortRangeMin *int64 `json:"netIpv4IpLocalPortRangeMin,omitempty" tf:"net_ipv4_ip_local_port_range_min"`

	// +kubebuilder:validation:Optional
	NetIPv4NeighDefaultGcThresh1 *int64 `json:"netIpv4NeighDefaultGcThresh1,omitempty" tf:"net_ipv4_neigh_default_gc_thresh1"`

	// +kubebuilder:validation:Optional
	NetIPv4NeighDefaultGcThresh2 *int64 `json:"netIpv4NeighDefaultGcThresh2,omitempty" tf:"net_ipv4_neigh_default_gc_thresh2"`

	// +kubebuilder:validation:Optional
	NetIPv4NeighDefaultGcThresh3 *int64 `json:"netIpv4NeighDefaultGcThresh3,omitempty" tf:"net_ipv4_neigh_default_gc_thresh3"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPFinTimeout *int64 `json:"netIpv4TcpFinTimeout,omitempty" tf:"net_ipv4_tcp_fin_timeout"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPKeepaliveIntvl *int64 `json:"netIpv4TcpKeepaliveIntvl,omitempty" tf:"net_ipv4_tcp_keepalive_intvl"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPKeepaliveProbes *int64 `json:"netIpv4TcpKeepaliveProbes,omitempty" tf:"net_ipv4_tcp_keepalive_probes"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPKeepaliveTime *int64 `json:"netIpv4TcpKeepaliveTime,omitempty" tf:"net_ipv4_tcp_keepalive_time"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPMaxSynBacklog *int64 `json:"netIpv4TcpMaxSynBacklog,omitempty" tf:"net_ipv4_tcp_max_syn_backlog"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPMaxTwBuckets *int64 `json:"netIpv4TcpMaxTwBuckets,omitempty" tf:"net_ipv4_tcp_max_tw_buckets"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty" tf:"net_ipv4_tcp_tw_reuse"`

	// +kubebuilder:validation:Optional
	NetNetfilterNfConntrackBuckets *int64 `json:"netNetfilterNfConntrackBuckets,omitempty" tf:"net_netfilter_nf_conntrack_buckets"`

	// +kubebuilder:validation:Optional
	NetNetfilterNfConntrackMax *int64 `json:"netNetfilterNfConntrackMax,omitempty" tf:"net_netfilter_nf_conntrack_max"`

	// +kubebuilder:validation:Optional
	VMMaxMapCount *int64 `json:"vmMaxMapCount,omitempty" tf:"vm_max_map_count"`

	// +kubebuilder:validation:Optional
	VMSwappiness *int64 `json:"vmSwappiness,omitempty" tf:"vm_swappiness"`

	// +kubebuilder:validation:Optional
	VMVfsCachePressure *int64 `json:"vmVfsCachePressure,omitempty" tf:"vm_vfs_cache_pressure"`
}

type UpgradeSettingsObservation struct {
}

type UpgradeSettingsParameters struct {

	// +kubebuilder:validation:Required
	MaxSurge string `json:"maxSurge" tf:"max_surge"`
}

type WindowsProfileObservation struct {
}

type WindowsProfileParameters struct {

	// +kubebuilder:validation:Optional
	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password"`

	// +kubebuilder:validation:Required
	AdminUsername string `json:"adminUsername" tf:"admin_username"`

	// +kubebuilder:validation:Optional
	License *string `json:"license,omitempty" tf:"license"`
}

// KubernetesClusterSpec defines the desired state of KubernetesCluster
type KubernetesClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KubernetesClusterParameters `json:"forProvider"`
}

// KubernetesClusterStatus defines the observed state of KubernetesCluster.
type KubernetesClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KubernetesClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesCluster is the Schema for the KubernetesClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesClusterList contains a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesCluster `json:"items"`
}

// Repository type metadata.
var (
	KubernetesClusterKind             = "KubernetesCluster"
	KubernetesClusterGroupKind        = schema.GroupKind{Group: Group, Kind: KubernetesClusterKind}.String()
	KubernetesClusterKindAPIVersion   = KubernetesClusterKind + "." + GroupVersion.String()
	KubernetesClusterGroupVersionKind = GroupVersion.WithKind(KubernetesClusterKind)
)

func init() {
	SchemeBuilder.Register(&KubernetesCluster{}, &KubernetesClusterList{})
}
