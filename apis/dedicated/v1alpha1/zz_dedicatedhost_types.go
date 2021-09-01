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

type DedicatedHostObservation struct {
}

type DedicatedHostParameters struct {
	AutoReplaceOnFailure *bool `json:"autoReplaceOnFailure,omitempty" tf:"auto_replace_on_failure"`

	DedicatedHostGroupId string `json:"dedicatedHostGroupId" tf:"dedicated_host_group_id"`

	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	PlatformFaultDomain int64 `json:"platformFaultDomain" tf:"platform_fault_domain"`

	SkuName string `json:"skuName" tf:"sku_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// DedicatedHostSpec defines the desired state of DedicatedHost
type DedicatedHostSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DedicatedHostParameters `json:"forProvider"`
}

// DedicatedHostStatus defines the observed state of DedicatedHost.
type DedicatedHostStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DedicatedHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHost is the Schema for the DedicatedHosts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DedicatedHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHostSpec   `json:"spec"`
	Status            DedicatedHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHostList contains a list of DedicatedHosts
type DedicatedHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedHost `json:"items"`
}

// Repository type metadata.
var (
	DedicatedHostKind             = "DedicatedHost"
	DedicatedHostGroupKind        = schema.GroupKind{Group: Group, Kind: DedicatedHostKind}.String()
	DedicatedHostKindAPIVersion   = DedicatedHostKind + "." + GroupVersion.String()
	DedicatedHostGroupVersionKind = GroupVersion.WithKind(DedicatedHostKind)
)

func init() {
	SchemeBuilder.Register(&DedicatedHost{}, &DedicatedHostList{})
}
