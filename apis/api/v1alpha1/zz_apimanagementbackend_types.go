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

type ApiManagementBackendObservation struct {
}

type ApiManagementBackendParameters struct {
	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`

	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Protocol string `json:"protocol" tf:"protocol"`

	Proxy []ApiManagementBackendProxyParameters `json:"proxy,omitempty" tf:"proxy"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ResourceId *string `json:"resourceId,omitempty" tf:"resource_id"`

	ServiceFabricCluster []ServiceFabricClusterParameters `json:"serviceFabricCluster,omitempty" tf:"service_fabric_cluster"`

	Title *string `json:"title,omitempty" tf:"title"`

	Tls []TlsParameters `json:"tls,omitempty" tf:"tls"`

	Url string `json:"url" tf:"url"`
}

type ApiManagementBackendProxyObservation struct {
}

type ApiManagementBackendProxyParameters struct {
	Password *string `json:"password,omitempty" tf:"password"`

	Url string `json:"url" tf:"url"`

	Username string `json:"username" tf:"username"`
}

type AuthorizationObservation struct {
}

type AuthorizationParameters struct {
	Parameter *string `json:"parameter,omitempty" tf:"parameter"`

	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type CredentialsObservation struct {
}

type CredentialsParameters struct {
	Authorization []AuthorizationParameters `json:"authorization,omitempty" tf:"authorization"`

	Certificate []string `json:"certificate,omitempty" tf:"certificate"`

	Header map[string]string `json:"header,omitempty" tf:"header"`

	Query map[string]string `json:"query,omitempty" tf:"query"`
}

type ServerX509NameObservation struct {
}

type ServerX509NameParameters struct {
	IssuerCertificateThumbprint string `json:"issuerCertificateThumbprint" tf:"issuer_certificate_thumbprint"`

	Name string `json:"name" tf:"name"`
}

type ServiceFabricClusterObservation struct {
}

type ServiceFabricClusterParameters struct {
	ClientCertificateId *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id"`

	ClientCertificateThumbprint *string `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint"`

	ManagementEndpoints []string `json:"managementEndpoints" tf:"management_endpoints"`

	MaxPartitionResolutionRetries int64 `json:"maxPartitionResolutionRetries" tf:"max_partition_resolution_retries"`

	ServerCertificateThumbprints []string `json:"serverCertificateThumbprints,omitempty" tf:"server_certificate_thumbprints"`

	ServerX509Name []ServerX509NameParameters `json:"serverX509Name,omitempty" tf:"server_x509_name"`
}

type TlsObservation struct {
}

type TlsParameters struct {
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty" tf:"validate_certificate_chain"`

	ValidateCertificateName *bool `json:"validateCertificateName,omitempty" tf:"validate_certificate_name"`
}

// ApiManagementBackendSpec defines the desired state of ApiManagementBackend
type ApiManagementBackendSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementBackendParameters `json:"forProvider"`
}

// ApiManagementBackendStatus defines the observed state of ApiManagementBackend.
type ApiManagementBackendStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementBackend is the Schema for the ApiManagementBackends API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementBackendSpec   `json:"spec"`
	Status            ApiManagementBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementBackendList contains a list of ApiManagementBackends
type ApiManagementBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementBackend `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementBackendKind             = "ApiManagementBackend"
	ApiManagementBackendGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementBackendKind}.String()
	ApiManagementBackendKindAPIVersion   = ApiManagementBackendKind + "." + GroupVersion.String()
	ApiManagementBackendGroupVersionKind = GroupVersion.WithKind(ApiManagementBackendKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementBackend{}, &ApiManagementBackendList{})
}
