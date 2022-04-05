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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DSCNodeConfigurationObservation struct {
	ConfigurationName *string `json:"configurationName,omitempty" tf:"configuration_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DSCNodeConfigurationParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Required
	ContentEmbedded *string `json:"contentEmbedded" tf:"content_embedded,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// DSCNodeConfigurationSpec defines the desired state of DSCNodeConfiguration
type DSCNodeConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DSCNodeConfigurationParameters `json:"forProvider"`
}

// DSCNodeConfigurationStatus defines the observed state of DSCNodeConfiguration.
type DSCNodeConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DSCNodeConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DSCNodeConfiguration is the Schema for the DSCNodeConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DSCNodeConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DSCNodeConfigurationSpec   `json:"spec"`
	Status            DSCNodeConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DSCNodeConfigurationList contains a list of DSCNodeConfigurations
type DSCNodeConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DSCNodeConfiguration `json:"items"`
}

// Repository type metadata.
var (
	DSCNodeConfiguration_Kind             = "DSCNodeConfiguration"
	DSCNodeConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DSCNodeConfiguration_Kind}.String()
	DSCNodeConfiguration_KindAPIVersion   = DSCNodeConfiguration_Kind + "." + CRDGroupVersion.String()
	DSCNodeConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(DSCNodeConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&DSCNodeConfiguration{}, &DSCNodeConfigurationList{})
}
