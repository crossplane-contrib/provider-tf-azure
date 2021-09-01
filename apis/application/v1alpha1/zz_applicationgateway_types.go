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

type ApplicationGatewayObservation struct {
}

type ApplicationGatewayParameters struct {
	AuthenticationCertificate []AuthenticationCertificateParameters `json:"authenticationCertificate,omitempty" tf:"authentication_certificate"`

	AutoscaleConfiguration []AutoscaleConfigurationParameters `json:"autoscaleConfiguration,omitempty" tf:"autoscale_configuration"`

	BackendAddressPool []BackendAddressPoolParameters `json:"backendAddressPool" tf:"backend_address_pool"`

	BackendHttpSettings []BackendHttpSettingsParameters `json:"backendHttpSettings" tf:"backend_http_settings"`

	CustomErrorConfiguration []CustomErrorConfigurationParameters `json:"customErrorConfiguration,omitempty" tf:"custom_error_configuration"`

	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2"`

	FirewallPolicyId *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id"`

	FrontendIpConfiguration []FrontendIpConfigurationParameters `json:"frontendIpConfiguration" tf:"frontend_ip_configuration"`

	FrontendPort []FrontendPortParameters `json:"frontendPort" tf:"frontend_port"`

	GatewayIpConfiguration []GatewayIpConfigurationParameters `json:"gatewayIpConfiguration" tf:"gateway_ip_configuration"`

	HttpListener []HttpListenerParameters `json:"httpListener" tf:"http_listener"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	Probe []ProbeParameters `json:"probe,omitempty" tf:"probe"`

	RedirectConfiguration []RedirectConfigurationParameters `json:"redirectConfiguration,omitempty" tf:"redirect_configuration"`

	RequestRoutingRule []RequestRoutingRuleParameters `json:"requestRoutingRule" tf:"request_routing_rule"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	RewriteRuleSet []RewriteRuleSetParameters `json:"rewriteRuleSet,omitempty" tf:"rewrite_rule_set"`

	Sku []SkuParameters `json:"sku" tf:"sku"`

	SslCertificate []SslCertificateParameters `json:"sslCertificate,omitempty" tf:"ssl_certificate"`

	SslPolicy []SslPolicyParameters `json:"sslPolicy,omitempty" tf:"ssl_policy"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TrustedRootCertificate []TrustedRootCertificateParameters `json:"trustedRootCertificate,omitempty" tf:"trusted_root_certificate"`

	UrlPathMap []UrlPathMapParameters `json:"urlPathMap,omitempty" tf:"url_path_map"`

	WafConfiguration []WafConfigurationParameters `json:"wafConfiguration,omitempty" tf:"waf_configuration"`

	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type AuthenticationCertificateObservation struct {
	Id string `json:"id" tf:"id"`
}

type AuthenticationCertificateParameters struct {
	Data string `json:"data" tf:"data"`

	Name string `json:"name" tf:"name"`
}

type AutoscaleConfigurationObservation struct {
}

type AutoscaleConfigurationParameters struct {
	MaxCapacity *int64 `json:"maxCapacity,omitempty" tf:"max_capacity"`

	MinCapacity int64 `json:"minCapacity" tf:"min_capacity"`
}

type BackendAddressPoolObservation struct {
	Id string `json:"id" tf:"id"`
}

type BackendAddressPoolParameters struct {
	Fqdns []string `json:"fqdns,omitempty" tf:"fqdns"`

	IpAddresses []string `json:"ipAddresses,omitempty" tf:"ip_addresses"`

	Name string `json:"name" tf:"name"`
}

type BackendHttpSettingsAuthenticationCertificateObservation struct {
	Id string `json:"id" tf:"id"`
}

type BackendHttpSettingsAuthenticationCertificateParameters struct {
	Name string `json:"name" tf:"name"`
}

type BackendHttpSettingsObservation struct {
	Id string `json:"id" tf:"id"`

	ProbeId string `json:"probeId" tf:"probe_id"`
}

type BackendHttpSettingsParameters struct {
	AffinityCookieName *string `json:"affinityCookieName,omitempty" tf:"affinity_cookie_name"`

	AuthenticationCertificate []BackendHttpSettingsAuthenticationCertificateParameters `json:"authenticationCertificate,omitempty" tf:"authentication_certificate"`

	ConnectionDraining []ConnectionDrainingParameters `json:"connectionDraining,omitempty" tf:"connection_draining"`

	CookieBasedAffinity string `json:"cookieBasedAffinity" tf:"cookie_based_affinity"`

	HostName *string `json:"hostName,omitempty" tf:"host_name"`

	Name string `json:"name" tf:"name"`

	Path *string `json:"path,omitempty" tf:"path"`

	PickHostNameFromBackendAddress *bool `json:"pickHostNameFromBackendAddress,omitempty" tf:"pick_host_name_from_backend_address"`

	Port int64 `json:"port" tf:"port"`

	ProbeName *string `json:"probeName,omitempty" tf:"probe_name"`

	Protocol string `json:"protocol" tf:"protocol"`

	RequestTimeout *int64 `json:"requestTimeout,omitempty" tf:"request_timeout"`

	TrustedRootCertificateNames []string `json:"trustedRootCertificateNames,omitempty" tf:"trusted_root_certificate_names"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case"`

	Negate *bool `json:"negate,omitempty" tf:"negate"`

	Pattern string `json:"pattern" tf:"pattern"`

	Variable string `json:"variable" tf:"variable"`
}

type ConnectionDrainingObservation struct {
}

type ConnectionDrainingParameters struct {
	DrainTimeoutSec int64 `json:"drainTimeoutSec" tf:"drain_timeout_sec"`

	Enabled bool `json:"enabled" tf:"enabled"`
}

type CustomErrorConfigurationObservation struct {
	Id string `json:"id" tf:"id"`
}

type CustomErrorConfigurationParameters struct {
	CustomErrorPageUrl string `json:"customErrorPageUrl" tf:"custom_error_page_url"`

	StatusCode string `json:"statusCode" tf:"status_code"`
}

type DisabledRuleGroupObservation struct {
}

type DisabledRuleGroupParameters struct {
	RuleGroupName string `json:"ruleGroupName" tf:"rule_group_name"`

	Rules []int64 `json:"rules,omitempty" tf:"rules"`
}

type ExclusionObservation struct {
}

type ExclusionParameters struct {
	MatchVariable string `json:"matchVariable" tf:"match_variable"`

	Selector *string `json:"selector,omitempty" tf:"selector"`

	SelectorMatchOperator *string `json:"selectorMatchOperator,omitempty" tf:"selector_match_operator"`
}

type FrontendIpConfigurationObservation struct {
	Id string `json:"id" tf:"id"`
}

type FrontendIpConfigurationParameters struct {
	Name string `json:"name" tf:"name"`

	PrivateIpAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address"`

	PrivateIpAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation"`

	PublicIpAddressId *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`
}

type FrontendPortObservation struct {
	Id string `json:"id" tf:"id"`
}

type FrontendPortParameters struct {
	Name string `json:"name" tf:"name"`

	Port int64 `json:"port" tf:"port"`
}

type GatewayIpConfigurationObservation struct {
	Id string `json:"id" tf:"id"`
}

type GatewayIpConfigurationParameters struct {
	Name string `json:"name" tf:"name"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

type HttpListenerCustomErrorConfigurationObservation struct {
	Id string `json:"id" tf:"id"`
}

type HttpListenerCustomErrorConfigurationParameters struct {
	CustomErrorPageUrl string `json:"customErrorPageUrl" tf:"custom_error_page_url"`

	StatusCode string `json:"statusCode" tf:"status_code"`
}

type HttpListenerObservation struct {
	FrontendIpConfigurationId string `json:"frontendIpConfigurationId" tf:"frontend_ip_configuration_id"`

	FrontendPortId string `json:"frontendPortId" tf:"frontend_port_id"`

	Id string `json:"id" tf:"id"`

	SslCertificateId string `json:"sslCertificateId" tf:"ssl_certificate_id"`
}

type HttpListenerParameters struct {
	CustomErrorConfiguration []HttpListenerCustomErrorConfigurationParameters `json:"customErrorConfiguration,omitempty" tf:"custom_error_configuration"`

	FirewallPolicyId *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id"`

	FrontendIpConfigurationName string `json:"frontendIpConfigurationName" tf:"frontend_ip_configuration_name"`

	FrontendPortName string `json:"frontendPortName" tf:"frontend_port_name"`

	HostName *string `json:"hostName,omitempty" tf:"host_name"`

	HostNames []string `json:"hostNames,omitempty" tf:"host_names"`

	Name string `json:"name" tf:"name"`

	Protocol string `json:"protocol" tf:"protocol"`

	RequireSni *bool `json:"requireSni,omitempty" tf:"require_sni"`

	SslCertificateName *string `json:"sslCertificateName,omitempty" tf:"ssl_certificate_name"`
}

type IdentityObservation struct {
}

type IdentityParameters struct {
	IdentityIds []string `json:"identityIds" tf:"identity_ids"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type MatchObservation struct {
}

type MatchParameters struct {
	Body *string `json:"body,omitempty" tf:"body"`

	StatusCode []string `json:"statusCode,omitempty" tf:"status_code"`
}

type PathRuleObservation struct {
	BackendAddressPoolId string `json:"backendAddressPoolId" tf:"backend_address_pool_id"`

	BackendHttpSettingsId string `json:"backendHttpSettingsId" tf:"backend_http_settings_id"`

	Id string `json:"id" tf:"id"`

	RedirectConfigurationId string `json:"redirectConfigurationId" tf:"redirect_configuration_id"`

	RewriteRuleSetId string `json:"rewriteRuleSetId" tf:"rewrite_rule_set_id"`
}

type PathRuleParameters struct {
	BackendAddressPoolName *string `json:"backendAddressPoolName,omitempty" tf:"backend_address_pool_name"`

	BackendHttpSettingsName *string `json:"backendHttpSettingsName,omitempty" tf:"backend_http_settings_name"`

	FirewallPolicyId *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id"`

	Name string `json:"name" tf:"name"`

	Paths []string `json:"paths" tf:"paths"`

	RedirectConfigurationName *string `json:"redirectConfigurationName,omitempty" tf:"redirect_configuration_name"`

	RewriteRuleSetName *string `json:"rewriteRuleSetName,omitempty" tf:"rewrite_rule_set_name"`
}

type ProbeObservation struct {
	Id string `json:"id" tf:"id"`
}

type ProbeParameters struct {
	Host *string `json:"host,omitempty" tf:"host"`

	Interval int64 `json:"interval" tf:"interval"`

	Match []MatchParameters `json:"match,omitempty" tf:"match"`

	MinimumServers *int64 `json:"minimumServers,omitempty" tf:"minimum_servers"`

	Name string `json:"name" tf:"name"`

	Path string `json:"path" tf:"path"`

	PickHostNameFromBackendHttpSettings *bool `json:"pickHostNameFromBackendHttpSettings,omitempty" tf:"pick_host_name_from_backend_http_settings"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	Protocol string `json:"protocol" tf:"protocol"`

	Timeout int64 `json:"timeout" tf:"timeout"`

	UnhealthyThreshold int64 `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type RedirectConfigurationObservation struct {
	Id string `json:"id" tf:"id"`

	TargetListenerId string `json:"targetListenerId" tf:"target_listener_id"`
}

type RedirectConfigurationParameters struct {
	IncludePath *bool `json:"includePath,omitempty" tf:"include_path"`

	IncludeQueryString *bool `json:"includeQueryString,omitempty" tf:"include_query_string"`

	Name string `json:"name" tf:"name"`

	RedirectType string `json:"redirectType" tf:"redirect_type"`

	TargetListenerName *string `json:"targetListenerName,omitempty" tf:"target_listener_name"`

	TargetUrl *string `json:"targetUrl,omitempty" tf:"target_url"`
}

type RequestHeaderConfigurationObservation struct {
}

type RequestHeaderConfigurationParameters struct {
	HeaderName string `json:"headerName" tf:"header_name"`

	HeaderValue string `json:"headerValue" tf:"header_value"`
}

type RequestRoutingRuleObservation struct {
	BackendAddressPoolId string `json:"backendAddressPoolId" tf:"backend_address_pool_id"`

	BackendHttpSettingsId string `json:"backendHttpSettingsId" tf:"backend_http_settings_id"`

	HttpListenerId string `json:"httpListenerId" tf:"http_listener_id"`

	Id string `json:"id" tf:"id"`

	RedirectConfigurationId string `json:"redirectConfigurationId" tf:"redirect_configuration_id"`

	RewriteRuleSetId string `json:"rewriteRuleSetId" tf:"rewrite_rule_set_id"`

	UrlPathMapId string `json:"urlPathMapId" tf:"url_path_map_id"`
}

type RequestRoutingRuleParameters struct {
	BackendAddressPoolName *string `json:"backendAddressPoolName,omitempty" tf:"backend_address_pool_name"`

	BackendHttpSettingsName *string `json:"backendHttpSettingsName,omitempty" tf:"backend_http_settings_name"`

	HttpListenerName string `json:"httpListenerName" tf:"http_listener_name"`

	Name string `json:"name" tf:"name"`

	RedirectConfigurationName *string `json:"redirectConfigurationName,omitempty" tf:"redirect_configuration_name"`

	RewriteRuleSetName *string `json:"rewriteRuleSetName,omitempty" tf:"rewrite_rule_set_name"`

	RuleType string `json:"ruleType" tf:"rule_type"`

	UrlPathMapName *string `json:"urlPathMapName,omitempty" tf:"url_path_map_name"`
}

type ResponseHeaderConfigurationObservation struct {
}

type ResponseHeaderConfigurationParameters struct {
	HeaderName string `json:"headerName" tf:"header_name"`

	HeaderValue string `json:"headerValue" tf:"header_value"`
}

type RewriteRuleObservation struct {
}

type RewriteRuleParameters struct {
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition"`

	Name string `json:"name" tf:"name"`

	RequestHeaderConfiguration []RequestHeaderConfigurationParameters `json:"requestHeaderConfiguration,omitempty" tf:"request_header_configuration"`

	ResponseHeaderConfiguration []ResponseHeaderConfigurationParameters `json:"responseHeaderConfiguration,omitempty" tf:"response_header_configuration"`

	RuleSequence int64 `json:"ruleSequence" tf:"rule_sequence"`

	Url []UrlParameters `json:"url,omitempty" tf:"url"`
}

type RewriteRuleSetObservation struct {
	Id string `json:"id" tf:"id"`
}

type RewriteRuleSetParameters struct {
	Name string `json:"name" tf:"name"`

	RewriteRule []RewriteRuleParameters `json:"rewriteRule,omitempty" tf:"rewrite_rule"`
}

type SkuObservation struct {
}

type SkuParameters struct {
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`

	Name string `json:"name" tf:"name"`

	Tier string `json:"tier" tf:"tier"`
}

type SslCertificateObservation struct {
	Id string `json:"id" tf:"id"`

	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type SslCertificateParameters struct {
	Data *string `json:"data,omitempty" tf:"data"`

	KeyVaultSecretId *string `json:"keyVaultSecretId,omitempty" tf:"key_vault_secret_id"`

	Name string `json:"name" tf:"name"`

	Password *string `json:"password,omitempty" tf:"password"`
}

type SslPolicyObservation struct {
}

type SslPolicyParameters struct {
	CipherSuites []string `json:"cipherSuites,omitempty" tf:"cipher_suites"`

	DisabledProtocols []string `json:"disabledProtocols,omitempty" tf:"disabled_protocols"`

	MinProtocolVersion *string `json:"minProtocolVersion,omitempty" tf:"min_protocol_version"`

	PolicyName *string `json:"policyName,omitempty" tf:"policy_name"`

	PolicyType *string `json:"policyType,omitempty" tf:"policy_type"`
}

type TrustedRootCertificateObservation struct {
	Id string `json:"id" tf:"id"`
}

type TrustedRootCertificateParameters struct {
	Data string `json:"data" tf:"data"`

	Name string `json:"name" tf:"name"`
}

type UrlObservation struct {
}

type UrlParameters struct {
	Path *string `json:"path,omitempty" tf:"path"`

	QueryString *string `json:"queryString,omitempty" tf:"query_string"`

	Reroute *bool `json:"reroute,omitempty" tf:"reroute"`
}

type UrlPathMapObservation struct {
	DefaultBackendAddressPoolId string `json:"defaultBackendAddressPoolId" tf:"default_backend_address_pool_id"`

	DefaultBackendHttpSettingsId string `json:"defaultBackendHttpSettingsId" tf:"default_backend_http_settings_id"`

	DefaultRedirectConfigurationId string `json:"defaultRedirectConfigurationId" tf:"default_redirect_configuration_id"`

	DefaultRewriteRuleSetId string `json:"defaultRewriteRuleSetId" tf:"default_rewrite_rule_set_id"`

	Id string `json:"id" tf:"id"`
}

type UrlPathMapParameters struct {
	DefaultBackendAddressPoolName *string `json:"defaultBackendAddressPoolName,omitempty" tf:"default_backend_address_pool_name"`

	DefaultBackendHttpSettingsName *string `json:"defaultBackendHttpSettingsName,omitempty" tf:"default_backend_http_settings_name"`

	DefaultRedirectConfigurationName *string `json:"defaultRedirectConfigurationName,omitempty" tf:"default_redirect_configuration_name"`

	DefaultRewriteRuleSetName *string `json:"defaultRewriteRuleSetName,omitempty" tf:"default_rewrite_rule_set_name"`

	Name string `json:"name" tf:"name"`

	PathRule []PathRuleParameters `json:"pathRule" tf:"path_rule"`
}

type WafConfigurationObservation struct {
}

type WafConfigurationParameters struct {
	DisabledRuleGroup []DisabledRuleGroupParameters `json:"disabledRuleGroup,omitempty" tf:"disabled_rule_group"`

	Enabled bool `json:"enabled" tf:"enabled"`

	Exclusion []ExclusionParameters `json:"exclusion,omitempty" tf:"exclusion"`

	FileUploadLimitMb *int64 `json:"fileUploadLimitMb,omitempty" tf:"file_upload_limit_mb"`

	FirewallMode string `json:"firewallMode" tf:"firewall_mode"`

	MaxRequestBodySizeKb *int64 `json:"maxRequestBodySizeKb,omitempty" tf:"max_request_body_size_kb"`

	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check"`

	RuleSetType *string `json:"ruleSetType,omitempty" tf:"rule_set_type"`

	RuleSetVersion string `json:"ruleSetVersion" tf:"rule_set_version"`
}

// ApplicationGatewaySpec defines the desired state of ApplicationGateway
type ApplicationGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApplicationGatewayParameters `json:"forProvider"`
}

// ApplicationGatewayStatus defines the observed state of ApplicationGateway.
type ApplicationGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApplicationGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationGateway is the Schema for the ApplicationGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationGatewaySpec   `json:"spec"`
	Status            ApplicationGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationGatewayList contains a list of ApplicationGateways
type ApplicationGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationGateway `json:"items"`
}

// Repository type metadata.
var (
	ApplicationGatewayKind             = "ApplicationGateway"
	ApplicationGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: ApplicationGatewayKind}.String()
	ApplicationGatewayKindAPIVersion   = ApplicationGatewayKind + "." + GroupVersion.String()
	ApplicationGatewayGroupVersionKind = GroupVersion.WithKind(ApplicationGatewayKind)
)

func init() {
	SchemeBuilder.Register(&ApplicationGateway{}, &ApplicationGatewayList{})
}
