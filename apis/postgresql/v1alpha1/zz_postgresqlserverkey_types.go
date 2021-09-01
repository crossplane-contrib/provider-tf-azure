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

type PostgresqlServerKeyObservation struct {
}

type PostgresqlServerKeyParameters struct {
	KeyVaultKeyId string `json:"keyVaultKeyId" tf:"key_vault_key_id"`

	ServerId string `json:"serverId" tf:"server_id"`
}

// PostgresqlServerKeySpec defines the desired state of PostgresqlServerKey
type PostgresqlServerKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostgresqlServerKeyParameters `json:"forProvider"`
}

// PostgresqlServerKeyStatus defines the observed state of PostgresqlServerKey.
type PostgresqlServerKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostgresqlServerKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServerKey is the Schema for the PostgresqlServerKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlServerKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlServerKeySpec   `json:"spec"`
	Status            PostgresqlServerKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServerKeyList contains a list of PostgresqlServerKeys
type PostgresqlServerKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlServerKey `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlServerKeyKind             = "PostgresqlServerKey"
	PostgresqlServerKeyGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlServerKeyKind}.String()
	PostgresqlServerKeyKindAPIVersion   = PostgresqlServerKeyKind + "." + GroupVersion.String()
	PostgresqlServerKeyGroupVersionKind = GroupVersion.WithKind(PostgresqlServerKeyKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlServerKey{}, &PostgresqlServerKeyList{})
}
