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

type IdentityObservation struct {
	PrincipalId string `json:"principalId" tf:"principal_id"`

	TenantId string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {
	Type string `json:"type" tf:"type"`
}

type PostgresqlServerObservation struct {
	Fqdn string `json:"fqdn" tf:"fqdn"`
}

type PostgresqlServerParameters struct {
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login"`

	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty" tf:"administrator_login_password"`

	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled"`

	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`

	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`

	CreationSourceServerId *string `json:"creationSourceServerId,omitempty" tf:"creation_source_server_id"`

	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time"`

	SkuName string `json:"skuName" tf:"sku_name"`

	SslEnforcement *string `json:"sslEnforcement,omitempty" tf:"ssl_enforcement"`

	SslEnforcementEnabled *bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled"`

	SslMinimalTlsVersionEnforced *string `json:"sslMinimalTlsVersionEnforced,omitempty" tf:"ssl_minimal_tls_version_enforced"`

	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`

	StorageProfile []StorageProfileParameters `json:"storageProfile,omitempty" tf:"storage_profile"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	ThreatDetectionPolicy []ThreatDetectionPolicyParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy"`

	Version string `json:"version" tf:"version"`
}

type StorageProfileObservation struct {
}

type StorageProfileParameters struct {
	AutoGrow *string `json:"autoGrow,omitempty" tf:"auto_grow"`

	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`

	GeoRedundantBackup *string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup"`

	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`
}

type ThreatDetectionPolicyObservation struct {
}

type ThreatDetectionPolicyParameters struct {
	DisabledAlerts []string `json:"disabledAlerts,omitempty" tf:"disabled_alerts"`

	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty" tf:"email_account_admins"`

	EmailAddresses []string `json:"emailAddresses,omitempty" tf:"email_addresses"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days"`

	StorageAccountAccessKey *string `json:"storageAccountAccessKey,omitempty" tf:"storage_account_access_key"`

	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint"`
}

// PostgresqlServerSpec defines the desired state of PostgresqlServer
type PostgresqlServerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostgresqlServerParameters `json:"forProvider"`
}

// PostgresqlServerStatus defines the observed state of PostgresqlServer.
type PostgresqlServerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostgresqlServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServer is the Schema for the PostgresqlServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlServerSpec   `json:"spec"`
	Status            PostgresqlServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServerList contains a list of PostgresqlServers
type PostgresqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlServer `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlServerKind             = "PostgresqlServer"
	PostgresqlServerGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlServerKind}.String()
	PostgresqlServerKindAPIVersion   = PostgresqlServerKind + "." + GroupVersion.String()
	PostgresqlServerGroupVersionKind = GroupVersion.WithKind(PostgresqlServerKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlServer{}, &PostgresqlServerList{})
}
