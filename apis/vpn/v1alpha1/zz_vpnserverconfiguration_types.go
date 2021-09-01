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

type AzureActiveDirectoryAuthenticationObservation struct {
}

type AzureActiveDirectoryAuthenticationParameters struct {
	Audience string `json:"audience" tf:"audience"`

	Issuer string `json:"issuer" tf:"issuer"`

	Tenant string `json:"tenant" tf:"tenant"`
}

type ClientRevokedCertificateObservation struct {
}

type ClientRevokedCertificateParameters struct {
	Name string `json:"name" tf:"name"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type ClientRootCertificateObservation struct {
}

type ClientRootCertificateParameters struct {
	Name string `json:"name" tf:"name"`

	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type RadiusClientRootCertificateObservation struct {
}

type RadiusClientRootCertificateParameters struct {
	Name string `json:"name" tf:"name"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type RadiusObservation struct {
}

type RadiusParameters struct {
	ClientRootCertificate []RadiusClientRootCertificateParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate"`

	Server []ServerParameters `json:"server,omitempty" tf:"server"`

	ServerRootCertificate []ServerRootCertificateParameters `json:"serverRootCertificate" tf:"server_root_certificate"`
}

type RadiusServerClientRootCertificateObservation struct {
}

type RadiusServerClientRootCertificateParameters struct {
	Name string `json:"name" tf:"name"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type RadiusServerObservation struct {
}

type RadiusServerParameters struct {
	Address string `json:"address" tf:"address"`

	ClientRootCertificate []RadiusServerClientRootCertificateParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate"`

	Secret string `json:"secret" tf:"secret"`

	ServerRootCertificate []RadiusServerServerRootCertificateParameters `json:"serverRootCertificate" tf:"server_root_certificate"`
}

type RadiusServerServerRootCertificateObservation struct {
}

type RadiusServerServerRootCertificateParameters struct {
	Name string `json:"name" tf:"name"`

	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type ServerObservation struct {
}

type ServerParameters struct {
	Address string `json:"address" tf:"address"`

	Score int64 `json:"score" tf:"score"`

	Secret string `json:"secret" tf:"secret"`
}

type ServerRootCertificateObservation struct {
}

type ServerRootCertificateParameters struct {
	Name string `json:"name" tf:"name"`

	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type VpnServerConfigurationIpsecPolicyObservation struct {
}

type VpnServerConfigurationIpsecPolicyParameters struct {
	DhGroup string `json:"dhGroup" tf:"dh_group"`

	IkeEncryption string `json:"ikeEncryption" tf:"ike_encryption"`

	IkeIntegrity string `json:"ikeIntegrity" tf:"ike_integrity"`

	IpsecEncryption string `json:"ipsecEncryption" tf:"ipsec_encryption"`

	IpsecIntegrity string `json:"ipsecIntegrity" tf:"ipsec_integrity"`

	PfsGroup string `json:"pfsGroup" tf:"pfs_group"`

	SaDataSizeKilobytes int64 `json:"saDataSizeKilobytes" tf:"sa_data_size_kilobytes"`

	SaLifetimeSeconds int64 `json:"saLifetimeSeconds" tf:"sa_lifetime_seconds"`
}

type VpnServerConfigurationObservation struct {
}

type VpnServerConfigurationParameters struct {
	AzureActiveDirectoryAuthentication []AzureActiveDirectoryAuthenticationParameters `json:"azureActiveDirectoryAuthentication,omitempty" tf:"azure_active_directory_authentication"`

	ClientRevokedCertificate []ClientRevokedCertificateParameters `json:"clientRevokedCertificate,omitempty" tf:"client_revoked_certificate"`

	ClientRootCertificate []ClientRootCertificateParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate"`

	IpsecPolicy []VpnServerConfigurationIpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	Radius []RadiusParameters `json:"radius,omitempty" tf:"radius"`

	RadiusServer []RadiusServerParameters `json:"radiusServer,omitempty" tf:"radius_server"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VpnAuthenticationTypes []string `json:"vpnAuthenticationTypes" tf:"vpn_authentication_types"`

	VpnProtocols []string `json:"vpnProtocols,omitempty" tf:"vpn_protocols"`
}

// VpnServerConfigurationSpec defines the desired state of VpnServerConfiguration
type VpnServerConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpnServerConfigurationParameters `json:"forProvider"`
}

// VpnServerConfigurationStatus defines the observed state of VpnServerConfiguration.
type VpnServerConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpnServerConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpnServerConfiguration is the Schema for the VpnServerConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VpnServerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnServerConfigurationSpec   `json:"spec"`
	Status            VpnServerConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnServerConfigurationList contains a list of VpnServerConfigurations
type VpnServerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnServerConfiguration `json:"items"`
}

// Repository type metadata.
var (
	VpnServerConfigurationKind             = "VpnServerConfiguration"
	VpnServerConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: VpnServerConfigurationKind}.String()
	VpnServerConfigurationKindAPIVersion   = VpnServerConfigurationKind + "." + GroupVersion.String()
	VpnServerConfigurationGroupVersionKind = GroupVersion.WithKind(VpnServerConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&VpnServerConfiguration{}, &VpnServerConfigurationList{})
}
