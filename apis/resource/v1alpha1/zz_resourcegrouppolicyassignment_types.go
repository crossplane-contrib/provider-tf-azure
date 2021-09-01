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
	Type *string `json:"type,omitempty" tf:"type"`
}

type ResourceGroupPolicyAssignmentObservation struct {
}

type ResourceGroupPolicyAssignmentParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	Enforce *bool `json:"enforce,omitempty" tf:"enforce"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location *string `json:"location,omitempty" tf:"location"`

	Metadata *string `json:"metadata,omitempty" tf:"metadata"`

	Name string `json:"name" tf:"name"`

	NotScopes []string `json:"notScopes,omitempty" tf:"not_scopes"`

	Parameters *string `json:"parameters,omitempty" tf:"parameters"`

	PolicyDefinitionId string `json:"policyDefinitionId" tf:"policy_definition_id"`

	ResourceGroupId string `json:"resourceGroupId" tf:"resource_group_id"`
}

// ResourceGroupPolicyAssignmentSpec defines the desired state of ResourceGroupPolicyAssignment
type ResourceGroupPolicyAssignmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResourceGroupPolicyAssignmentParameters `json:"forProvider"`
}

// ResourceGroupPolicyAssignmentStatus defines the observed state of ResourceGroupPolicyAssignment.
type ResourceGroupPolicyAssignmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResourceGroupPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignment is the Schema for the ResourceGroupPolicyAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroupPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupPolicyAssignmentSpec   `json:"spec"`
	Status            ResourceGroupPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignmentList contains a list of ResourceGroupPolicyAssignments
type ResourceGroupPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupPolicyAssignmentKind             = "ResourceGroupPolicyAssignment"
	ResourceGroupPolicyAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: ResourceGroupPolicyAssignmentKind}.String()
	ResourceGroupPolicyAssignmentKindAPIVersion   = ResourceGroupPolicyAssignmentKind + "." + GroupVersion.String()
	ResourceGroupPolicyAssignmentGroupVersionKind = GroupVersion.WithKind(ResourceGroupPolicyAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupPolicyAssignment{}, &ResourceGroupPolicyAssignmentList{})
}
