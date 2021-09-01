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

type ActiveDirectoryDomainServiceReplicaSetObservation struct {
	DomainControllerIpAddresses []string `json:"domainControllerIpAddresses" tf:"domain_controller_ip_addresses"`

	ExternalAccessIpAddress string `json:"externalAccessIpAddress" tf:"external_access_ip_address"`

	ServiceStatus string `json:"serviceStatus" tf:"service_status"`
}

type ActiveDirectoryDomainServiceReplicaSetParameters struct {
	DomainServiceId string `json:"domainServiceId" tf:"domain_service_id"`

	Location string `json:"location" tf:"location"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

// ActiveDirectoryDomainServiceReplicaSetSpec defines the desired state of ActiveDirectoryDomainServiceReplicaSet
type ActiveDirectoryDomainServiceReplicaSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ActiveDirectoryDomainServiceReplicaSetParameters `json:"forProvider"`
}

// ActiveDirectoryDomainServiceReplicaSetStatus defines the observed state of ActiveDirectoryDomainServiceReplicaSet.
type ActiveDirectoryDomainServiceReplicaSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ActiveDirectoryDomainServiceReplicaSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryDomainServiceReplicaSet is the Schema for the ActiveDirectoryDomainServiceReplicaSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ActiveDirectoryDomainServiceReplicaSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActiveDirectoryDomainServiceReplicaSetSpec   `json:"spec"`
	Status            ActiveDirectoryDomainServiceReplicaSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryDomainServiceReplicaSetList contains a list of ActiveDirectoryDomainServiceReplicaSets
type ActiveDirectoryDomainServiceReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryDomainServiceReplicaSet `json:"items"`
}

// Repository type metadata.
var (
	ActiveDirectoryDomainServiceReplicaSetKind             = "ActiveDirectoryDomainServiceReplicaSet"
	ActiveDirectoryDomainServiceReplicaSetGroupKind        = schema.GroupKind{Group: Group, Kind: ActiveDirectoryDomainServiceReplicaSetKind}.String()
	ActiveDirectoryDomainServiceReplicaSetKindAPIVersion   = ActiveDirectoryDomainServiceReplicaSetKind + "." + GroupVersion.String()
	ActiveDirectoryDomainServiceReplicaSetGroupVersionKind = GroupVersion.WithKind(ActiveDirectoryDomainServiceReplicaSetKind)
)

func init() {
	SchemeBuilder.Register(&ActiveDirectoryDomainServiceReplicaSet{}, &ActiveDirectoryDomainServiceReplicaSetList{})
}
