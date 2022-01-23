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

type DatabaseMigrationProjectObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatabaseMigrationProjectParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Required
	SourcePlatform *string `json:"sourcePlatform" tf:"source_platform,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetPlatform *string `json:"targetPlatform" tf:"target_platform,omitempty"`
}

// DatabaseMigrationProjectSpec defines the desired state of DatabaseMigrationProject
type DatabaseMigrationProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseMigrationProjectParameters `json:"forProvider"`
}

// DatabaseMigrationProjectStatus defines the observed state of DatabaseMigrationProject.
type DatabaseMigrationProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseMigrationProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationProject is the Schema for the DatabaseMigrationProjects API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DatabaseMigrationProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseMigrationProjectSpec   `json:"spec"`
	Status            DatabaseMigrationProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationProjectList contains a list of DatabaseMigrationProjects
type DatabaseMigrationProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationProject `json:"items"`
}

// Repository type metadata.
var (
	DatabaseMigrationProject_Kind             = "DatabaseMigrationProject"
	DatabaseMigrationProject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseMigrationProject_Kind}.String()
	DatabaseMigrationProject_KindAPIVersion   = DatabaseMigrationProject_Kind + "." + CRDGroupVersion.String()
	DatabaseMigrationProject_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseMigrationProject_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseMigrationProject{}, &DatabaseMigrationProjectList{})
}
