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

type AdditionalLocationObservation struct {
	GatewayRegionalUrl string `json:"gatewayRegionalUrl" tf:"gateway_regional_url"`

	PrivateIpAddresses []string `json:"privateIpAddresses" tf:"private_ip_addresses"`

	PublicIpAddresses []string `json:"publicIpAddresses" tf:"public_ip_addresses"`
}

type AdditionalLocationParameters struct {
	Location string `json:"location" tf:"location"`

	VirtualNetworkConfiguration []VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration"`
}

type ApiManagementObservation struct {
	DeveloperPortalUrl string `json:"developerPortalUrl" tf:"developer_portal_url"`

	GatewayRegionalUrl string `json:"gatewayRegionalUrl" tf:"gateway_regional_url"`

	GatewayUrl string `json:"gatewayUrl" tf:"gateway_url"`

	ManagementApiUrl string `json:"managementApiUrl" tf:"management_api_url"`

	PortalUrl string `json:"portalUrl" tf:"portal_url"`

	PrivateIpAddresses []string `json:"privateIpAddresses" tf:"private_ip_addresses"`

	PublicIpAddresses []string `json:"publicIpAddresses" tf:"public_ip_addresses"`

	ScmUrl string `json:"scmUrl" tf:"scm_url"`
}

type ApiManagementParameters struct {
	AdditionalLocation []AdditionalLocationParameters `json:"additionalLocation,omitempty" tf:"additional_location"`

	Certificate []CertificateParameters `json:"certificate,omitempty" tf:"certificate"`

	ClientCertificateEnabled *bool `json:"clientCertificateEnabled,omitempty" tf:"client_certificate_enabled"`

	GatewayDisabled *bool `json:"gatewayDisabled,omitempty" tf:"gateway_disabled"`

	HostnameConfiguration []HostnameConfigurationParameters `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location string `json:"location" tf:"location"`

	MinApiVersion *string `json:"minApiVersion,omitempty" tf:"min_api_version"`

	Name string `json:"name" tf:"name"`

	NotificationSenderEmail *string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email"`

	Policy []PolicyParameters `json:"policy,omitempty" tf:"policy"`

	Protocols []ProtocolsParameters `json:"protocols,omitempty" tf:"protocols"`

	PublisherEmail string `json:"publisherEmail" tf:"publisher_email"`

	PublisherName string `json:"publisherName" tf:"publisher_name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Security []SecurityParameters `json:"security,omitempty" tf:"security"`

	SignIn []SignInParameters `json:"signIn,omitempty" tf:"sign_in"`

	SignUp []SignUpParameters `json:"signUp,omitempty" tf:"sign_up"`

	SkuName string `json:"skuName" tf:"sku_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TenantAccess []TenantAccessParameters `json:"tenantAccess,omitempty" tf:"tenant_access"`

	VirtualNetworkConfiguration []ApiManagementVirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration"`

	VirtualNetworkType *string `json:"virtualNetworkType,omitempty" tf:"virtual_network_type"`

	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type ApiManagementVirtualNetworkConfigurationObservation struct {
}

type ApiManagementVirtualNetworkConfigurationParameters struct {
	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

type CertificateObservation struct {
	Expiry string `json:"expiry" tf:"expiry"`

	Subject string `json:"subject" tf:"subject"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type CertificateParameters struct {
	CertificatePassword *string `json:"certificatePassword,omitempty" tf:"certificate_password"`

	EncodedCertificate string `json:"encodedCertificate" tf:"encoded_certificate"`

	StoreName string `json:"storeName" tf:"store_name"`
}

type DeveloperPortalObservation struct {
	Expiry string `json:"expiry" tf:"expiry"`

	Subject string `json:"subject" tf:"subject"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type DeveloperPortalParameters struct {
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	CertificatePassword *string `json:"certificatePassword,omitempty" tf:"certificate_password"`

	HostName string `json:"hostName" tf:"host_name"`

	KeyVaultId *string `json:"keyVaultId,omitempty" tf:"key_vault_id"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`

	SslKeyvaultIdentityClientId *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id"`
}

type HostnameConfigurationObservation struct {
}

type HostnameConfigurationParameters struct {
	DeveloperPortal []DeveloperPortalParameters `json:"developerPortal,omitempty" tf:"developer_portal"`

	Management []ManagementParameters `json:"management,omitempty" tf:"management"`

	Portal []PortalParameters `json:"portal,omitempty" tf:"portal"`

	Proxy []ProxyParameters `json:"proxy,omitempty" tf:"proxy"`

	Scm []ScmParameters `json:"scm,omitempty" tf:"scm"`
}

type IdentityObservation struct {
	PrincipalId string `json:"principalId" tf:"principal_id"`

	TenantId string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type ManagementObservation struct {
	Expiry string `json:"expiry" tf:"expiry"`

	Subject string `json:"subject" tf:"subject"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type ManagementParameters struct {
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	CertificatePassword *string `json:"certificatePassword,omitempty" tf:"certificate_password"`

	HostName string `json:"hostName" tf:"host_name"`

	KeyVaultId *string `json:"keyVaultId,omitempty" tf:"key_vault_id"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`

	SslKeyvaultIdentityClientId *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id"`
}

type PolicyObservation struct {
}

type PolicyParameters struct {
	XmlContent *string `json:"xmlContent,omitempty" tf:"xml_content"`

	XmlLink *string `json:"xmlLink,omitempty" tf:"xml_link"`
}

type PortalObservation struct {
	Expiry string `json:"expiry" tf:"expiry"`

	Subject string `json:"subject" tf:"subject"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type PortalParameters struct {
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	CertificatePassword *string `json:"certificatePassword,omitempty" tf:"certificate_password"`

	HostName string `json:"hostName" tf:"host_name"`

	KeyVaultId *string `json:"keyVaultId,omitempty" tf:"key_vault_id"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`

	SslKeyvaultIdentityClientId *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id"`
}

type ProtocolsObservation struct {
}

type ProtocolsParameters struct {
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2"`
}

type ProxyObservation struct {
	Expiry string `json:"expiry" tf:"expiry"`

	Subject string `json:"subject" tf:"subject"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type ProxyParameters struct {
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	CertificatePassword *string `json:"certificatePassword,omitempty" tf:"certificate_password"`

	DefaultSslBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding"`

	HostName string `json:"hostName" tf:"host_name"`

	KeyVaultId *string `json:"keyVaultId,omitempty" tf:"key_vault_id"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`

	SslKeyvaultIdentityClientId *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id"`
}

type ScmObservation struct {
	Expiry string `json:"expiry" tf:"expiry"`

	Subject string `json:"subject" tf:"subject"`

	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type ScmParameters struct {
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	CertificatePassword *string `json:"certificatePassword,omitempty" tf:"certificate_password"`

	HostName string `json:"hostName" tf:"host_name"`

	KeyVaultId *string `json:"keyVaultId,omitempty" tf:"key_vault_id"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`

	SslKeyvaultIdentityClientId *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id"`
}

type SecurityObservation struct {
}

type SecurityParameters struct {
	EnableBackendSsl30 *bool `json:"enableBackendSsl30,omitempty" tf:"enable_backend_ssl30"`

	EnableBackendTls10 *bool `json:"enableBackendTls10,omitempty" tf:"enable_backend_tls10"`

	EnableBackendTls11 *bool `json:"enableBackendTls11,omitempty" tf:"enable_backend_tls11"`

	EnableFrontendSsl30 *bool `json:"enableFrontendSsl30,omitempty" tf:"enable_frontend_ssl30"`

	EnableFrontendTls10 *bool `json:"enableFrontendTls10,omitempty" tf:"enable_frontend_tls10"`

	EnableFrontendTls11 *bool `json:"enableFrontendTls11,omitempty" tf:"enable_frontend_tls11"`

	EnableTripleDesCiphers *bool `json:"enableTripleDesCiphers,omitempty" tf:"enable_triple_des_ciphers"`

	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled"`

	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled"`

	TlsEcdheRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled"`

	TlsEcdheRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled"`

	TlsRsaWithAes128CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes128CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha256_ciphers_enabled"`

	TlsRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha_ciphers_enabled"`

	TlsRsaWithAes128GcmSha256CiphersEnabled *bool `json:"tlsRsaWithAes128GcmSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_gcm_sha256_ciphers_enabled"`

	TlsRsaWithAes256CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes256CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha256_ciphers_enabled"`

	TlsRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha_ciphers_enabled"`

	TripleDesCiphersEnabled *bool `json:"tripleDesCiphersEnabled,omitempty" tf:"triple_des_ciphers_enabled"`
}

type SignInObservation struct {
}

type SignInParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type SignUpObservation struct {
}

type SignUpParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`

	TermsOfService []TermsOfServiceParameters `json:"termsOfService" tf:"terms_of_service"`
}

type TenantAccessObservation struct {
	PrimaryKey string `json:"primaryKey" tf:"primary_key"`

	SecondaryKey string `json:"secondaryKey" tf:"secondary_key"`

	TenantId string `json:"tenantId" tf:"tenant_id"`
}

type TenantAccessParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type TermsOfServiceObservation struct {
}

type TermsOfServiceParameters struct {
	ConsentRequired bool `json:"consentRequired" tf:"consent_required"`

	Enabled bool `json:"enabled" tf:"enabled"`

	Text *string `json:"text,omitempty" tf:"text"`
}

type VirtualNetworkConfigurationObservation struct {
}

type VirtualNetworkConfigurationParameters struct {
	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

// ApiManagementSpec defines the desired state of ApiManagement
type ApiManagementSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementParameters `json:"forProvider"`
}

// ApiManagementStatus defines the observed state of ApiManagement.
type ApiManagementStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagement is the Schema for the ApiManagements API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSpec   `json:"spec"`
	Status            ApiManagementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementList contains a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagement `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementKind             = "ApiManagement"
	ApiManagementGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementKind}.String()
	ApiManagementKindAPIVersion   = ApiManagementKind + "." + GroupVersion.String()
	ApiManagementGroupVersionKind = GroupVersion.WithKind(ApiManagementKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagement{}, &ApiManagementList{})
}
