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

type RoleAssignmentObservation struct {
	PrincipalType string `json:"principalType" tf:"principal_type"`
}

type RoleAssignmentParameters struct {
	Condition *string `json:"condition,omitempty" tf:"condition"`

	ConditionVersion *string `json:"conditionVersion,omitempty" tf:"condition_version"`

	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty" tf:"delegated_managed_identity_resource_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name *string `json:"name,omitempty" tf:"name"`

	PrincipalId string `json:"principalId" tf:"principal_id"`

	RoleDefinitionId *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id"`

	RoleDefinitionName *string `json:"roleDefinitionName,omitempty" tf:"role_definition_name"`

	Scope string `json:"scope" tf:"scope"`

	SkipServicePrincipalAadCheck *bool `json:"skipServicePrincipalAadCheck,omitempty" tf:"skip_service_principal_aad_check"`
}

// RoleAssignmentSpec defines the desired state of RoleAssignment
type RoleAssignmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RoleAssignmentParameters `json:"forProvider"`
}

// RoleAssignmentStatus defines the observed state of RoleAssignment.
type RoleAssignmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignment is the Schema for the RoleAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentSpec   `json:"spec"`
	Status            RoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentList contains a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignmentKind             = "RoleAssignment"
	RoleAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: RoleAssignmentKind}.String()
	RoleAssignmentKindAPIVersion   = RoleAssignmentKind + "." + GroupVersion.String()
	RoleAssignmentGroupVersionKind = GroupVersion.WithKind(RoleAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
