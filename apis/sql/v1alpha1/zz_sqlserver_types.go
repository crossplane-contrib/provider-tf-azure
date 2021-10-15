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

type ExtendedAuditingPolicyObservation struct {
}

type ExtendedAuditingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SqlServerObservation struct {
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty" tf:"fully_qualified_domain_name,omitempty"`
}

type SqlServerParameters struct {

	// +kubebuilder:validation:Required
	AdministratorLogin *string `json:"administratorLogin" tf:"administrator_login,omitempty"`

	// +kubebuilder:validation:Required
	AdministratorLoginPasswordSecretRef v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionPolicy *string `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ExtendedAuditingPolicy []ExtendedAuditingPolicyParameters `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThreatDetectionPolicy []ThreatDetectionPolicyParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ThreatDetectionPolicyObservation struct {
}

type ThreatDetectionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// SqlServerSpec defines the desired state of SqlServer
type SqlServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlServerParameters `json:"forProvider"`
}

// SqlServerStatus defines the observed state of SqlServer.
type SqlServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlServer is the Schema for the SqlServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlServerSpec   `json:"spec"`
	Status            SqlServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlServerList contains a list of SqlServers
type SqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlServer `json:"items"`
}

// Repository type metadata.
var (
	SqlServerKind             = "SqlServer"
	SqlServerGroupKind        = schema.GroupKind{Group: Group, Kind: SqlServerKind}.String()
	SqlServerKindAPIVersion   = SqlServerKind + "." + GroupVersion.String()
	SqlServerGroupVersionKind = GroupVersion.WithKind(SqlServerKind)
)

func init() {
	SchemeBuilder.Register(&SqlServer{}, &SqlServerList{})
}
