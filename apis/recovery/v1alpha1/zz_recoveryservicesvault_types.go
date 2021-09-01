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

type RecoveryServicesVaultObservation struct {
}

type RecoveryServicesVaultParameters struct {
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Sku string `json:"sku" tf:"sku"`

	SoftDeleteEnabled *bool `json:"softDeleteEnabled,omitempty" tf:"soft_delete_enabled"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// RecoveryServicesVaultSpec defines the desired state of RecoveryServicesVault
type RecoveryServicesVaultSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RecoveryServicesVaultParameters `json:"forProvider"`
}

// RecoveryServicesVaultStatus defines the observed state of RecoveryServicesVault.
type RecoveryServicesVaultStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RecoveryServicesVaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryServicesVault is the Schema for the RecoveryServicesVaults API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RecoveryServicesVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryServicesVaultSpec   `json:"spec"`
	Status            RecoveryServicesVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryServicesVaultList contains a list of RecoveryServicesVaults
type RecoveryServicesVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecoveryServicesVault `json:"items"`
}

// Repository type metadata.
var (
	RecoveryServicesVaultKind             = "RecoveryServicesVault"
	RecoveryServicesVaultGroupKind        = schema.GroupKind{Group: Group, Kind: RecoveryServicesVaultKind}.String()
	RecoveryServicesVaultKindAPIVersion   = RecoveryServicesVaultKind + "." + GroupVersion.String()
	RecoveryServicesVaultGroupVersionKind = GroupVersion.WithKind(RecoveryServicesVaultKind)
)

func init() {
	SchemeBuilder.Register(&RecoveryServicesVault{}, &RecoveryServicesVaultList{})
}
