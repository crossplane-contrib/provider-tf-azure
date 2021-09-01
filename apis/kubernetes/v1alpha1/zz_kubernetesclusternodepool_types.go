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

type KubernetesClusterNodePoolKubeletConfigObservation struct {
}

type KubernetesClusterNodePoolKubeletConfigParameters struct {
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls,omitempty" tf:"allowed_unsafe_sysctls"`

	ContainerLogMaxLine *int64 `json:"containerLogMaxLine,omitempty" tf:"container_log_max_line"`

	ContainerLogMaxSizeMb *int64 `json:"containerLogMaxSizeMb,omitempty" tf:"container_log_max_size_mb"`

	CpuCfsQuotaEnabled *bool `json:"cpuCfsQuotaEnabled,omitempty" tf:"cpu_cfs_quota_enabled"`

	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty" tf:"cpu_cfs_quota_period"`

	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty" tf:"cpu_manager_policy"`

	ImageGcHighThreshold *int64 `json:"imageGcHighThreshold,omitempty" tf:"image_gc_high_threshold"`

	ImageGcLowThreshold *int64 `json:"imageGcLowThreshold,omitempty" tf:"image_gc_low_threshold"`

	PodMaxPid *int64 `json:"podMaxPid,omitempty" tf:"pod_max_pid"`

	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty" tf:"topology_manager_policy"`
}

type KubernetesClusterNodePoolLinuxOsConfigObservation struct {
}

type KubernetesClusterNodePoolLinuxOsConfigParameters struct {
	SwapFileSizeMb *int64 `json:"swapFileSizeMb,omitempty" tf:"swap_file_size_mb"`

	SysctlConfig []LinuxOsConfigSysctlConfigParameters `json:"sysctlConfig,omitempty" tf:"sysctl_config"`

	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty" tf:"transparent_huge_page_defrag"`

	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty" tf:"transparent_huge_page_enabled"`
}

type KubernetesClusterNodePoolObservation struct {
}

type KubernetesClusterNodePoolParameters struct {
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`

	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling"`

	EnableHostEncryption *bool `json:"enableHostEncryption,omitempty" tf:"enable_host_encryption"`

	EnableNodePublicIp *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip"`

	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`

	FipsEnabled *bool `json:"fipsEnabled,omitempty" tf:"fips_enabled"`

	KubeletConfig []KubernetesClusterNodePoolKubeletConfigParameters `json:"kubeletConfig,omitempty" tf:"kubelet_config"`

	KubeletDiskType *string `json:"kubeletDiskType,omitempty" tf:"kubelet_disk_type"`

	KubernetesClusterId string `json:"kubernetesClusterId" tf:"kubernetes_cluster_id"`

	LinuxOsConfig []KubernetesClusterNodePoolLinuxOsConfigParameters `json:"linuxOsConfig,omitempty" tf:"linux_os_config"`

	MaxCount *int64 `json:"maxCount,omitempty" tf:"max_count"`

	MaxPods *int64 `json:"maxPods,omitempty" tf:"max_pods"`

	MinCount *int64 `json:"minCount,omitempty" tf:"min_count"`

	Mode *string `json:"mode,omitempty" tf:"mode"`

	Name string `json:"name" tf:"name"`

	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`

	NodeLabels map[string]string `json:"nodeLabels,omitempty" tf:"node_labels"`

	NodePublicIpPrefixId *string `json:"nodePublicIpPrefixId,omitempty" tf:"node_public_ip_prefix_id"`

	NodeTaints []string `json:"nodeTaints,omitempty" tf:"node_taints"`

	OrchestratorVersion *string `json:"orchestratorVersion,omitempty" tf:"orchestrator_version"`

	OsDiskSizeGb *int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb"`

	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type"`

	OsType *string `json:"osType,omitempty" tf:"os_type"`

	PodSubnetId *string `json:"podSubnetId,omitempty" tf:"pod_subnet_id"`

	Priority *string `json:"priority,omitempty" tf:"priority"`

	ProximityPlacementGroupId *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id"`

	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty" tf:"spot_max_price"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled"`

	UpgradeSettings []KubernetesClusterNodePoolUpgradeSettingsParameters `json:"upgradeSettings,omitempty" tf:"upgrade_settings"`

	VmSize string `json:"vmSize" tf:"vm_size"`

	VnetSubnetId *string `json:"vnetSubnetId,omitempty" tf:"vnet_subnet_id"`
}

type KubernetesClusterNodePoolUpgradeSettingsObservation struct {
}

type KubernetesClusterNodePoolUpgradeSettingsParameters struct {
	MaxSurge string `json:"maxSurge" tf:"max_surge"`
}

type LinuxOsConfigSysctlConfigObservation struct {
}

type LinuxOsConfigSysctlConfigParameters struct {
	FsAioMaxNr *int64 `json:"fsAioMaxNr,omitempty" tf:"fs_aio_max_nr"`

	FsFileMax *int64 `json:"fsFileMax,omitempty" tf:"fs_file_max"`

	FsInotifyMaxUserWatches *int64 `json:"fsInotifyMaxUserWatches,omitempty" tf:"fs_inotify_max_user_watches"`

	FsNrOpen *int64 `json:"fsNrOpen,omitempty" tf:"fs_nr_open"`

	KernelThreadsMax *int64 `json:"kernelThreadsMax,omitempty" tf:"kernel_threads_max"`

	NetCoreNetdevMaxBacklog *int64 `json:"netCoreNetdevMaxBacklog,omitempty" tf:"net_core_netdev_max_backlog"`

	NetCoreOptmemMax *int64 `json:"netCoreOptmemMax,omitempty" tf:"net_core_optmem_max"`

	NetCoreRmemDefault *int64 `json:"netCoreRmemDefault,omitempty" tf:"net_core_rmem_default"`

	NetCoreRmemMax *int64 `json:"netCoreRmemMax,omitempty" tf:"net_core_rmem_max"`

	NetCoreSomaxconn *int64 `json:"netCoreSomaxconn,omitempty" tf:"net_core_somaxconn"`

	NetCoreWmemDefault *int64 `json:"netCoreWmemDefault,omitempty" tf:"net_core_wmem_default"`

	NetCoreWmemMax *int64 `json:"netCoreWmemMax,omitempty" tf:"net_core_wmem_max"`

	NetIpv4IpLocalPortRangeMax *int64 `json:"netIpv4IpLocalPortRangeMax,omitempty" tf:"net_ipv4_ip_local_port_range_max"`

	NetIpv4IpLocalPortRangeMin *int64 `json:"netIpv4IpLocalPortRangeMin,omitempty" tf:"net_ipv4_ip_local_port_range_min"`

	NetIpv4NeighDefaultGcThresh1 *int64 `json:"netIpv4NeighDefaultGcThresh1,omitempty" tf:"net_ipv4_neigh_default_gc_thresh1"`

	NetIpv4NeighDefaultGcThresh2 *int64 `json:"netIpv4NeighDefaultGcThresh2,omitempty" tf:"net_ipv4_neigh_default_gc_thresh2"`

	NetIpv4NeighDefaultGcThresh3 *int64 `json:"netIpv4NeighDefaultGcThresh3,omitempty" tf:"net_ipv4_neigh_default_gc_thresh3"`

	NetIpv4TcpFinTimeout *int64 `json:"netIpv4TcpFinTimeout,omitempty" tf:"net_ipv4_tcp_fin_timeout"`

	NetIpv4TcpKeepaliveIntvl *int64 `json:"netIpv4TcpKeepaliveIntvl,omitempty" tf:"net_ipv4_tcp_keepalive_intvl"`

	NetIpv4TcpKeepaliveProbes *int64 `json:"netIpv4TcpKeepaliveProbes,omitempty" tf:"net_ipv4_tcp_keepalive_probes"`

	NetIpv4TcpKeepaliveTime *int64 `json:"netIpv4TcpKeepaliveTime,omitempty" tf:"net_ipv4_tcp_keepalive_time"`

	NetIpv4TcpMaxSynBacklog *int64 `json:"netIpv4TcpMaxSynBacklog,omitempty" tf:"net_ipv4_tcp_max_syn_backlog"`

	NetIpv4TcpMaxTwBuckets *int64 `json:"netIpv4TcpMaxTwBuckets,omitempty" tf:"net_ipv4_tcp_max_tw_buckets"`

	NetIpv4TcpTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty" tf:"net_ipv4_tcp_tw_reuse"`

	NetNetfilterNfConntrackBuckets *int64 `json:"netNetfilterNfConntrackBuckets,omitempty" tf:"net_netfilter_nf_conntrack_buckets"`

	NetNetfilterNfConntrackMax *int64 `json:"netNetfilterNfConntrackMax,omitempty" tf:"net_netfilter_nf_conntrack_max"`

	VmMaxMapCount *int64 `json:"vmMaxMapCount,omitempty" tf:"vm_max_map_count"`

	VmSwappiness *int64 `json:"vmSwappiness,omitempty" tf:"vm_swappiness"`

	VmVfsCachePressure *int64 `json:"vmVfsCachePressure,omitempty" tf:"vm_vfs_cache_pressure"`
}

// KubernetesClusterNodePoolSpec defines the desired state of KubernetesClusterNodePool
type KubernetesClusterNodePoolSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KubernetesClusterNodePoolParameters `json:"forProvider"`
}

// KubernetesClusterNodePoolStatus defines the observed state of KubernetesClusterNodePool.
type KubernetesClusterNodePoolStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KubernetesClusterNodePoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesClusterNodePool is the Schema for the KubernetesClusterNodePools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KubernetesClusterNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterNodePoolSpec   `json:"spec"`
	Status            KubernetesClusterNodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesClusterNodePoolList contains a list of KubernetesClusterNodePools
type KubernetesClusterNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesClusterNodePool `json:"items"`
}

// Repository type metadata.
var (
	KubernetesClusterNodePoolKind             = "KubernetesClusterNodePool"
	KubernetesClusterNodePoolGroupKind        = schema.GroupKind{Group: Group, Kind: KubernetesClusterNodePoolKind}.String()
	KubernetesClusterNodePoolKindAPIVersion   = KubernetesClusterNodePoolKind + "." + GroupVersion.String()
	KubernetesClusterNodePoolGroupVersionKind = GroupVersion.WithKind(KubernetesClusterNodePoolKind)
)

func init() {
	SchemeBuilder.Register(&KubernetesClusterNodePool{}, &KubernetesClusterNodePoolList{})
}
