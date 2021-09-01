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

type BackendObservation struct {
}

type BackendParameters struct {
	Address string `json:"address" tf:"address"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	HostHeader string `json:"hostHeader" tf:"host_header"`

	HttpPort int64 `json:"httpPort" tf:"http_port"`

	HttpsPort int64 `json:"httpsPort" tf:"https_port"`

	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type BackendPoolHealthProbeObservation struct {
	Id string `json:"id" tf:"id"`
}

type BackendPoolHealthProbeParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds"`

	Name string `json:"name" tf:"name"`

	Path *string `json:"path,omitempty" tf:"path"`

	ProbeMethod *string `json:"probeMethod,omitempty" tf:"probe_method"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
}

type BackendPoolLoadBalancingObservation struct {
	Id string `json:"id" tf:"id"`
}

type BackendPoolLoadBalancingParameters struct {
	AdditionalLatencyMilliseconds *int64 `json:"additionalLatencyMilliseconds,omitempty" tf:"additional_latency_milliseconds"`

	Name string `json:"name" tf:"name"`

	SampleSize *int64 `json:"sampleSize,omitempty" tf:"sample_size"`

	SuccessfulSamplesRequired *int64 `json:"successfulSamplesRequired,omitempty" tf:"successful_samples_required"`
}

type BackendPoolObservation struct {
	Id string `json:"id" tf:"id"`
}

type BackendPoolParameters struct {
	Backend []BackendParameters `json:"backend" tf:"backend"`

	HealthProbeName string `json:"healthProbeName" tf:"health_probe_name"`

	LoadBalancingName string `json:"loadBalancingName" tf:"load_balancing_name"`

	Name string `json:"name" tf:"name"`
}

type ExplicitResourceOrderObservation struct {
	BackendPoolHealthProbeIds []string `json:"backendPoolHealthProbeIds" tf:"backend_pool_health_probe_ids"`

	BackendPoolIds []string `json:"backendPoolIds" tf:"backend_pool_ids"`

	BackendPoolLoadBalancingIds []string `json:"backendPoolLoadBalancingIds" tf:"backend_pool_load_balancing_ids"`

	FrontendEndpointIds []string `json:"frontendEndpointIds" tf:"frontend_endpoint_ids"`

	RoutingRuleIds []string `json:"routingRuleIds" tf:"routing_rule_ids"`
}

type ExplicitResourceOrderParameters struct {
}

type ForwardingConfigurationObservation struct {
}

type ForwardingConfigurationParameters struct {
	BackendPoolName string `json:"backendPoolName" tf:"backend_pool_name"`

	CacheDuration *string `json:"cacheDuration,omitempty" tf:"cache_duration"`

	CacheEnabled *bool `json:"cacheEnabled,omitempty" tf:"cache_enabled"`

	CacheQueryParameterStripDirective *string `json:"cacheQueryParameterStripDirective,omitempty" tf:"cache_query_parameter_strip_directive"`

	CacheQueryParameters []string `json:"cacheQueryParameters,omitempty" tf:"cache_query_parameters"`

	CacheUseDynamicCompression *bool `json:"cacheUseDynamicCompression,omitempty" tf:"cache_use_dynamic_compression"`

	CustomForwardingPath *string `json:"customForwardingPath,omitempty" tf:"custom_forwarding_path"`

	ForwardingProtocol *string `json:"forwardingProtocol,omitempty" tf:"forwarding_protocol"`
}

type FrontdoorObservation struct {
	BackendPoolHealthProbes map[string]string `json:"backendPoolHealthProbes" tf:"backend_pool_health_probes"`

	BackendPoolLoadBalancingSettings map[string]string `json:"backendPoolLoadBalancingSettings" tf:"backend_pool_load_balancing_settings"`

	BackendPools map[string]string `json:"backendPools" tf:"backend_pools"`

	Cname string `json:"cname" tf:"cname"`

	ExplicitResourceOrder []ExplicitResourceOrderObservation `json:"explicitResourceOrder" tf:"explicit_resource_order"`

	FrontendEndpoints map[string]string `json:"frontendEndpoints" tf:"frontend_endpoints"`

	HeaderFrontdoorId string `json:"headerFrontdoorId" tf:"header_frontdoor_id"`

	RoutingRules map[string]string `json:"routingRules" tf:"routing_rules"`
}

type FrontdoorParameters struct {
	BackendPool []BackendPoolParameters `json:"backendPool" tf:"backend_pool"`

	BackendPoolHealthProbe []BackendPoolHealthProbeParameters `json:"backendPoolHealthProbe" tf:"backend_pool_health_probe"`

	BackendPoolLoadBalancing []BackendPoolLoadBalancingParameters `json:"backendPoolLoadBalancing" tf:"backend_pool_load_balancing"`

	BackendPoolsSendReceiveTimeoutSeconds *int64 `json:"backendPoolsSendReceiveTimeoutSeconds,omitempty" tf:"backend_pools_send_receive_timeout_seconds"`

	EnforceBackendPoolsCertificateNameCheck bool `json:"enforceBackendPoolsCertificateNameCheck" tf:"enforce_backend_pools_certificate_name_check"`

	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name"`

	FrontendEndpoint []FrontendEndpointParameters `json:"frontendEndpoint" tf:"frontend_endpoint"`

	LoadBalancerEnabled *bool `json:"loadBalancerEnabled,omitempty" tf:"load_balancer_enabled"`

	Location *string `json:"location,omitempty" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	RoutingRule []RoutingRuleParameters `json:"routingRule" tf:"routing_rule"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type FrontendEndpointObservation struct {
	Id string `json:"id" tf:"id"`
}

type FrontendEndpointParameters struct {
	HostName string `json:"hostName" tf:"host_name"`

	Name string `json:"name" tf:"name"`

	SessionAffinityEnabled *bool `json:"sessionAffinityEnabled,omitempty" tf:"session_affinity_enabled"`

	SessionAffinityTtlSeconds *int64 `json:"sessionAffinityTtlSeconds,omitempty" tf:"session_affinity_ttl_seconds"`

	WebApplicationFirewallPolicyLinkId *string `json:"webApplicationFirewallPolicyLinkId,omitempty" tf:"web_application_firewall_policy_link_id"`
}

type RedirectConfigurationObservation struct {
}

type RedirectConfigurationParameters struct {
	CustomFragment *string `json:"customFragment,omitempty" tf:"custom_fragment"`

	CustomHost *string `json:"customHost,omitempty" tf:"custom_host"`

	CustomPath *string `json:"customPath,omitempty" tf:"custom_path"`

	CustomQueryString *string `json:"customQueryString,omitempty" tf:"custom_query_string"`

	RedirectProtocol string `json:"redirectProtocol" tf:"redirect_protocol"`

	RedirectType string `json:"redirectType" tf:"redirect_type"`
}

type RoutingRuleObservation struct {
	Id string `json:"id" tf:"id"`
}

type RoutingRuleParameters struct {
	AcceptedProtocols []string `json:"acceptedProtocols" tf:"accepted_protocols"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	ForwardingConfiguration []ForwardingConfigurationParameters `json:"forwardingConfiguration,omitempty" tf:"forwarding_configuration"`

	FrontendEndpoints []string `json:"frontendEndpoints" tf:"frontend_endpoints"`

	Name string `json:"name" tf:"name"`

	PatternsToMatch []string `json:"patternsToMatch" tf:"patterns_to_match"`

	RedirectConfiguration []RedirectConfigurationParameters `json:"redirectConfiguration,omitempty" tf:"redirect_configuration"`
}

// FrontdoorSpec defines the desired state of Frontdoor
type FrontdoorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FrontdoorParameters `json:"forProvider"`
}

// FrontdoorStatus defines the observed state of Frontdoor.
type FrontdoorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FrontdoorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Frontdoor is the Schema for the Frontdoors API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Frontdoor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorSpec   `json:"spec"`
	Status            FrontdoorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorList contains a list of Frontdoors
type FrontdoorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Frontdoor `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorKind             = "Frontdoor"
	FrontdoorGroupKind        = schema.GroupKind{Group: Group, Kind: FrontdoorKind}.String()
	FrontdoorKindAPIVersion   = FrontdoorKind + "." + GroupVersion.String()
	FrontdoorGroupVersionKind = GroupVersion.WithKind(FrontdoorKind)
)

func init() {
	SchemeBuilder.Register(&Frontdoor{}, &FrontdoorList{})
}
