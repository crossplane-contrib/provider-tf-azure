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

type BotChannelSlackObservation struct {
}

type BotChannelSlackParameters struct {
	BotName string `json:"botName" tf:"bot_name"`

	ClientId string `json:"clientId" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	LandingPageUrl *string `json:"landingPageUrl,omitempty" tf:"landing_page_url"`

	Location string `json:"location" tf:"location"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	VerificationToken string `json:"verificationToken" tf:"verification_token"`
}

// BotChannelSlackSpec defines the desired state of BotChannelSlack
type BotChannelSlackSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BotChannelSlackParameters `json:"forProvider"`
}

// BotChannelSlackStatus defines the observed state of BotChannelSlack.
type BotChannelSlackStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BotChannelSlackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSlack is the Schema for the BotChannelSlacks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelSlack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelSlackSpec   `json:"spec"`
	Status            BotChannelSlackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSlackList contains a list of BotChannelSlacks
type BotChannelSlackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelSlack `json:"items"`
}

// Repository type metadata.
var (
	BotChannelSlackKind             = "BotChannelSlack"
	BotChannelSlackGroupKind        = schema.GroupKind{Group: Group, Kind: BotChannelSlackKind}.String()
	BotChannelSlackKindAPIVersion   = BotChannelSlackKind + "." + GroupVersion.String()
	BotChannelSlackGroupVersionKind = GroupVersion.WithKind(BotChannelSlackKind)
)

func init() {
	SchemeBuilder.Register(&BotChannelSlack{}, &BotChannelSlackList{})
}
