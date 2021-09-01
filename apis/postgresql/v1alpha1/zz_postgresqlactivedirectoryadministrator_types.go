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

type PostgresqlActiveDirectoryAdministratorObservation struct {
}

type PostgresqlActiveDirectoryAdministratorParameters struct {
	Login string `json:"login" tf:"login"`

	ObjectId string `json:"objectId" tf:"object_id"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ServerName string `json:"serverName" tf:"server_name"`

	TenantId string `json:"tenantId" tf:"tenant_id"`
}

// PostgresqlActiveDirectoryAdministratorSpec defines the desired state of PostgresqlActiveDirectoryAdministrator
type PostgresqlActiveDirectoryAdministratorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostgresqlActiveDirectoryAdministratorParameters `json:"forProvider"`
}

// PostgresqlActiveDirectoryAdministratorStatus defines the observed state of PostgresqlActiveDirectoryAdministrator.
type PostgresqlActiveDirectoryAdministratorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostgresqlActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlActiveDirectoryAdministrator is the Schema for the PostgresqlActiveDirectoryAdministrators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlActiveDirectoryAdministratorSpec   `json:"spec"`
	Status            PostgresqlActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlActiveDirectoryAdministratorList contains a list of PostgresqlActiveDirectoryAdministrators
type PostgresqlActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlActiveDirectoryAdministratorKind             = "PostgresqlActiveDirectoryAdministrator"
	PostgresqlActiveDirectoryAdministratorGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlActiveDirectoryAdministratorKind}.String()
	PostgresqlActiveDirectoryAdministratorKindAPIVersion   = PostgresqlActiveDirectoryAdministratorKind + "." + GroupVersion.String()
	PostgresqlActiveDirectoryAdministratorGroupVersionKind = GroupVersion.WithKind(PostgresqlActiveDirectoryAdministratorKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlActiveDirectoryAdministrator{}, &PostgresqlActiveDirectoryAdministratorList{})
}
