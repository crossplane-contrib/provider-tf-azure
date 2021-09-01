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

type SpringCloudAppCosmosdbAssociationObservation struct {
}

type SpringCloudAppCosmosdbAssociationParameters struct {
	ApiType string `json:"apiType" tf:"api_type"`

	CosmosdbAccessKey string `json:"cosmosdbAccessKey" tf:"cosmosdb_access_key"`

	CosmosdbAccountId string `json:"cosmosdbAccountId" tf:"cosmosdb_account_id"`

	CosmosdbCassandraKeyspaceName *string `json:"cosmosdbCassandraKeyspaceName,omitempty" tf:"cosmosdb_cassandra_keyspace_name"`

	CosmosdbGremlinDatabaseName *string `json:"cosmosdbGremlinDatabaseName,omitempty" tf:"cosmosdb_gremlin_database_name"`

	CosmosdbGremlinGraphName *string `json:"cosmosdbGremlinGraphName,omitempty" tf:"cosmosdb_gremlin_graph_name"`

	CosmosdbMongoDatabaseName *string `json:"cosmosdbMongoDatabaseName,omitempty" tf:"cosmosdb_mongo_database_name"`

	CosmosdbSqlDatabaseName *string `json:"cosmosdbSqlDatabaseName,omitempty" tf:"cosmosdb_sql_database_name"`

	Name string `json:"name" tf:"name"`

	SpringCloudAppId string `json:"springCloudAppId" tf:"spring_cloud_app_id"`
}

// SpringCloudAppCosmosdbAssociationSpec defines the desired state of SpringCloudAppCosmosdbAssociation
type SpringCloudAppCosmosdbAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SpringCloudAppCosmosdbAssociationParameters `json:"forProvider"`
}

// SpringCloudAppCosmosdbAssociationStatus defines the observed state of SpringCloudAppCosmosdbAssociation.
type SpringCloudAppCosmosdbAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SpringCloudAppCosmosdbAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppCosmosdbAssociation is the Schema for the SpringCloudAppCosmosdbAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAppCosmosdbAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudAppCosmosdbAssociationSpec   `json:"spec"`
	Status            SpringCloudAppCosmosdbAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppCosmosdbAssociationList contains a list of SpringCloudAppCosmosdbAssociations
type SpringCloudAppCosmosdbAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAppCosmosdbAssociation `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAppCosmosdbAssociationKind             = "SpringCloudAppCosmosdbAssociation"
	SpringCloudAppCosmosdbAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: SpringCloudAppCosmosdbAssociationKind}.String()
	SpringCloudAppCosmosdbAssociationKindAPIVersion   = SpringCloudAppCosmosdbAssociationKind + "." + GroupVersion.String()
	SpringCloudAppCosmosdbAssociationGroupVersionKind = GroupVersion.WithKind(SpringCloudAppCosmosdbAssociationKind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAppCosmosdbAssociation{}, &SpringCloudAppCosmosdbAssociationList{})
}
