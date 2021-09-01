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

type RedisLinkedServerObservation struct {
	Name string `json:"name" tf:"name"`
}

type RedisLinkedServerParameters struct {
	LinkedRedisCacheId string `json:"linkedRedisCacheId" tf:"linked_redis_cache_id"`

	LinkedRedisCacheLocation string `json:"linkedRedisCacheLocation" tf:"linked_redis_cache_location"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ServerRole string `json:"serverRole" tf:"server_role"`

	TargetRedisCacheName string `json:"targetRedisCacheName" tf:"target_redis_cache_name"`
}

// RedisLinkedServerSpec defines the desired state of RedisLinkedServer
type RedisLinkedServerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RedisLinkedServerParameters `json:"forProvider"`
}

// RedisLinkedServerStatus defines the observed state of RedisLinkedServer.
type RedisLinkedServerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RedisLinkedServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisLinkedServer is the Schema for the RedisLinkedServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServerSpec   `json:"spec"`
	Status            RedisLinkedServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisLinkedServerList contains a list of RedisLinkedServers
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

// Repository type metadata.
var (
	RedisLinkedServerKind             = "RedisLinkedServer"
	RedisLinkedServerGroupKind        = schema.GroupKind{Group: Group, Kind: RedisLinkedServerKind}.String()
	RedisLinkedServerKindAPIVersion   = RedisLinkedServerKind + "." + GroupVersion.String()
	RedisLinkedServerGroupVersionKind = GroupVersion.WithKind(RedisLinkedServerKind)
)

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}
