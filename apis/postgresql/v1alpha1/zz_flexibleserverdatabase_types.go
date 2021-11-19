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

type FlexibleServerDatabaseObservation struct {
}

type FlexibleServerDatabaseParameters struct {

	// +kubebuilder:validation:Optional
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// +crossplane:generate:reference:type=FlexibleServer
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// FlexibleServerDatabaseSpec defines the desired state of FlexibleServerDatabase
type FlexibleServerDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerDatabaseParameters `json:"forProvider"`
}

// FlexibleServerDatabaseStatus defines the observed state of FlexibleServerDatabase.
type FlexibleServerDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerDatabase is the Schema for the FlexibleServerDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FlexibleServerDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerDatabaseSpec   `json:"spec"`
	Status            FlexibleServerDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerDatabaseList contains a list of FlexibleServerDatabases
type FlexibleServerDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerDatabase `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerDatabase_Kind             = "FlexibleServerDatabase"
	FlexibleServerDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerDatabase_Kind}.String()
	FlexibleServerDatabase_KindAPIVersion   = FlexibleServerDatabase_Kind + "." + CRDGroupVersion.String()
	FlexibleServerDatabase_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerDatabase{}, &FlexibleServerDatabaseList{})
}
