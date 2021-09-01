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

type BotChannelDirectlineObservation struct {
}

type BotChannelDirectlineParameters struct {
	BotName string `json:"botName" tf:"bot_name"`

	Location string `json:"location" tf:"location"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Site []SiteParameters `json:"site" tf:"site"`
}

type SiteObservation struct {
	Id string `json:"id" tf:"id"`

	Key string `json:"key" tf:"key"`

	Key2 string `json:"key2" tf:"key2"`
}

type SiteParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EnhancedAuthenticationEnabled *bool `json:"enhancedAuthenticationEnabled,omitempty" tf:"enhanced_authentication_enabled"`

	Name string `json:"name" tf:"name"`

	TrustedOrigins []string `json:"trustedOrigins,omitempty" tf:"trusted_origins"`

	V1Allowed *bool `json:"v1Allowed,omitempty" tf:"v1_allowed"`

	V3Allowed *bool `json:"v3Allowed,omitempty" tf:"v3_allowed"`
}

// BotChannelDirectlineSpec defines the desired state of BotChannelDirectline
type BotChannelDirectlineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BotChannelDirectlineParameters `json:"forProvider"`
}

// BotChannelDirectlineStatus defines the observed state of BotChannelDirectline.
type BotChannelDirectlineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BotChannelDirectlineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelDirectline is the Schema for the BotChannelDirectlines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelDirectline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelDirectlineSpec   `json:"spec"`
	Status            BotChannelDirectlineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelDirectlineList contains a list of BotChannelDirectlines
type BotChannelDirectlineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelDirectline `json:"items"`
}

// Repository type metadata.
var (
	BotChannelDirectlineKind             = "BotChannelDirectline"
	BotChannelDirectlineGroupKind        = schema.GroupKind{Group: Group, Kind: BotChannelDirectlineKind}.String()
	BotChannelDirectlineKindAPIVersion   = BotChannelDirectlineKind + "." + GroupVersion.String()
	BotChannelDirectlineGroupVersionKind = GroupVersion.WithKind(BotChannelDirectlineKind)
)

func init() {
	SchemeBuilder.Register(&BotChannelDirectline{}, &BotChannelDirectlineList{})
}
