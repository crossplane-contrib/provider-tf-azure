/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SQLDatabaseAutoscaleSettingsObservation struct {
}

type SQLDatabaseAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type SQLDatabaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLDatabaseParameters struct {

	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []SQLDatabaseAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

// SQLDatabaseSpec defines the desired state of SQLDatabase
type SQLDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLDatabaseParameters `json:"forProvider"`
}

// SQLDatabaseStatus defines the observed state of SQLDatabase.
type SQLDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDatabase is the Schema for the SQLDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLDatabaseSpec   `json:"spec"`
	Status            SQLDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDatabaseList contains a list of SQLDatabases
type SQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLDatabase `json:"items"`
}

// Repository type metadata.
var (
	SQLDatabase_Kind             = "SQLDatabase"
	SQLDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLDatabase_Kind}.String()
	SQLDatabase_KindAPIVersion   = SQLDatabase_Kind + "." + CRDGroupVersion.String()
	SQLDatabase_GroupVersionKind = CRDGroupVersion.WithKind(SQLDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLDatabase{}, &SQLDatabaseList{})
}
