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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BgpSettingsObservation struct {
}

type BgpSettingsParameters struct {

	// +kubebuilder:validation:Optional
	Asn *int64 `json:"asn,omitempty" tf:"asn"`

	// +kubebuilder:validation:Optional
	PeerWeight *int64 `json:"peerWeight,omitempty" tf:"peer_weight"`

	// +kubebuilder:validation:Optional
	PeeringAddress *string `json:"peeringAddress,omitempty" tf:"peering_address"`

	// +kubebuilder:validation:Optional
	PeeringAddresses []PeeringAddressesParameters `json:"peeringAddresses,omitempty" tf:"peering_addresses"`
}

type CustomRouteObservation struct {
}

type CustomRouteParameters struct {

	// +kubebuilder:validation:Optional
	AddressPrefixes []string `json:"addressPrefixes,omitempty" tf:"address_prefixes"`
}

type PeeringAddressesObservation struct {
	DefaultAddresses []string `json:"defaultAddresses,omitempty" tf:"default_addresses"`

	TunnelIPAddresses []string `json:"tunnelIpAddresses,omitempty" tf:"tunnel_ip_addresses"`
}

type PeeringAddressesParameters struct {

	// +kubebuilder:validation:Optional
	ApipaAddresses []string `json:"apipaAddresses,omitempty" tf:"apipa_addresses"`

	// +kubebuilder:validation:Optional
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name"`
}

type RevokedCertificateObservation struct {
}

type RevokedCertificateParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type RootCertificateObservation struct {
}

type RootCertificateParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type VirtualNetworkGatewayIPConfigurationObservation struct {
}

type VirtualNetworkGatewayIPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation"`

	// +kubebuilder:validation:Required
	PublicIPAddressID string `json:"publicIpAddressId" tf:"public_ip_address_id"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

type VirtualNetworkGatewayObservation struct {
}

type VirtualNetworkGatewayParameters struct {

	// +kubebuilder:validation:Optional
	ActiveActive *bool `json:"activeActive,omitempty" tf:"active_active"`

	// +kubebuilder:validation:Optional
	BgpSettings []BgpSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings"`

	// +kubebuilder:validation:Optional
	CustomRoute []CustomRouteParameters `json:"customRoute,omitempty" tf:"custom_route"`

	// +kubebuilder:validation:Optional
	DefaultLocalNetworkGatewayID *string `json:"defaultLocalNetworkGatewayId,omitempty" tf:"default_local_network_gateway_id"`

	// +kubebuilder:validation:Optional
	EnableBgp *bool `json:"enableBgp,omitempty" tf:"enable_bgp"`

	// +kubebuilder:validation:Optional
	Generation *string `json:"generation,omitempty" tf:"generation"`

	// +kubebuilder:validation:Required
	IPConfiguration []VirtualNetworkGatewayIPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PrivateIPAddressEnabled *bool `json:"privateIpAddressEnabled,omitempty" tf:"private_ip_address_enabled"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Sku string `json:"sku" tf:"sku"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	VpnClientConfiguration []VpnClientConfigurationParameters `json:"vpnClientConfiguration,omitempty" tf:"vpn_client_configuration"`

	// +kubebuilder:validation:Optional
	VpnType *string `json:"vpnType,omitempty" tf:"vpn_type"`
}

type VpnClientConfigurationObservation struct {
}

type VpnClientConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AadAudience *string `json:"aadAudience,omitempty" tf:"aad_audience"`

	// +kubebuilder:validation:Optional
	AadIssuer *string `json:"aadIssuer,omitempty" tf:"aad_issuer"`

	// +kubebuilder:validation:Optional
	AadTenant *string `json:"aadTenant,omitempty" tf:"aad_tenant"`

	// +kubebuilder:validation:Required
	AddressSpace []string `json:"addressSpace" tf:"address_space"`

	// +kubebuilder:validation:Optional
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty" tf:"radius_server_address"`

	// +kubebuilder:validation:Optional
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty" tf:"radius_server_secret"`

	// +kubebuilder:validation:Optional
	RevokedCertificate []RevokedCertificateParameters `json:"revokedCertificate,omitempty" tf:"revoked_certificate"`

	// +kubebuilder:validation:Optional
	RootCertificate []RootCertificateParameters `json:"rootCertificate,omitempty" tf:"root_certificate"`

	// +kubebuilder:validation:Optional
	VpnClientProtocols []string `json:"vpnClientProtocols,omitempty" tf:"vpn_client_protocols"`
}

// VirtualNetworkGatewaySpec defines the desired state of VirtualNetworkGateway
type VirtualNetworkGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkGatewayParameters `json:"forProvider"`
}

// VirtualNetworkGatewayStatus defines the observed state of VirtualNetworkGateway.
type VirtualNetworkGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGateway is the Schema for the VirtualNetworkGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewaySpec   `json:"spec"`
	Status            VirtualNetworkGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayList contains a list of VirtualNetworkGateways
type VirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGateway `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkGatewayKind             = "VirtualNetworkGateway"
	VirtualNetworkGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualNetworkGatewayKind}.String()
	VirtualNetworkGatewayKindAPIVersion   = VirtualNetworkGatewayKind + "." + GroupVersion.String()
	VirtualNetworkGatewayGroupVersionKind = GroupVersion.WithKind(VirtualNetworkGatewayKind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkGateway{}, &VirtualNetworkGatewayList{})
}
