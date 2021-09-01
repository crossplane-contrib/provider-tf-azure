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

type ApiManagementDiagnosticBackendRequestObservation struct {
}

type ApiManagementDiagnosticBackendRequestParameters struct {
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	DataMasking []BackendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type ApiManagementDiagnosticBackendResponseDataMaskingObservation struct {
}

type ApiManagementDiagnosticBackendResponseDataMaskingParameters struct {
	Headers []BackendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	QueryParams []BackendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type ApiManagementDiagnosticBackendResponseObservation struct {
}

type ApiManagementDiagnosticBackendResponseParameters struct {
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	DataMasking []ApiManagementDiagnosticBackendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type ApiManagementDiagnosticFrontendRequestDataMaskingHeadersObservation struct {
}

type ApiManagementDiagnosticFrontendRequestDataMaskingHeadersParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type ApiManagementDiagnosticFrontendRequestDataMaskingObservation struct {
}

type ApiManagementDiagnosticFrontendRequestDataMaskingParameters struct {
	Headers []ApiManagementDiagnosticFrontendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	QueryParams []ApiManagementDiagnosticFrontendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type ApiManagementDiagnosticFrontendRequestDataMaskingQueryParamsObservation struct {
}

type ApiManagementDiagnosticFrontendRequestDataMaskingQueryParamsParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type ApiManagementDiagnosticFrontendRequestObservation struct {
}

type ApiManagementDiagnosticFrontendRequestParameters struct {
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	DataMasking []ApiManagementDiagnosticFrontendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type ApiManagementDiagnosticFrontendResponseDataMaskingHeadersObservation struct {
}

type ApiManagementDiagnosticFrontendResponseDataMaskingHeadersParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type ApiManagementDiagnosticFrontendResponseDataMaskingObservation struct {
}

type ApiManagementDiagnosticFrontendResponseDataMaskingParameters struct {
	Headers []ApiManagementDiagnosticFrontendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	QueryParams []ApiManagementDiagnosticFrontendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type ApiManagementDiagnosticFrontendResponseDataMaskingQueryParamsObservation struct {
}

type ApiManagementDiagnosticFrontendResponseDataMaskingQueryParamsParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type ApiManagementDiagnosticFrontendResponseObservation struct {
}

type ApiManagementDiagnosticFrontendResponseParameters struct {
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	DataMasking []ApiManagementDiagnosticFrontendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type ApiManagementDiagnosticObservation struct {
}

type ApiManagementDiagnosticParameters struct {
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors"`

	ApiManagementLoggerId string `json:"apiManagementLoggerId" tf:"api_management_logger_id"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`

	BackendRequest []ApiManagementDiagnosticBackendRequestParameters `json:"backendRequest,omitempty" tf:"backend_request"`

	BackendResponse []ApiManagementDiagnosticBackendResponseParameters `json:"backendResponse,omitempty" tf:"backend_response"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	FrontendRequest []ApiManagementDiagnosticFrontendRequestParameters `json:"frontendRequest,omitempty" tf:"frontend_request"`

	FrontendResponse []ApiManagementDiagnosticFrontendResponseParameters `json:"frontendResponse,omitempty" tf:"frontend_response"`

	HttpCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol"`

	Identifier string `json:"identifier" tf:"identifier"`

	LogClientIp *bool `json:"logClientIp,omitempty" tf:"log_client_ip"`

	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage"`

	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity"`
}

type BackendRequestDataMaskingHeadersObservation struct {
}

type BackendRequestDataMaskingHeadersParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type BackendRequestDataMaskingObservation struct {
}

type BackendRequestDataMaskingParameters struct {
	Headers []BackendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	QueryParams []BackendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type BackendRequestDataMaskingQueryParamsObservation struct {
}

type BackendRequestDataMaskingQueryParamsParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type BackendResponseDataMaskingHeadersObservation struct {
}

type BackendResponseDataMaskingHeadersParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

type BackendResponseDataMaskingQueryParamsObservation struct {
}

type BackendResponseDataMaskingQueryParamsParameters struct {
	Mode string `json:"mode" tf:"mode"`

	Value string `json:"value" tf:"value"`
}

// ApiManagementDiagnosticSpec defines the desired state of ApiManagementDiagnostic
type ApiManagementDiagnosticSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementDiagnosticParameters `json:"forProvider"`
}

// ApiManagementDiagnosticStatus defines the observed state of ApiManagementDiagnostic.
type ApiManagementDiagnosticStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementDiagnosticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementDiagnostic is the Schema for the ApiManagementDiagnostics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementDiagnostic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementDiagnosticSpec   `json:"spec"`
	Status            ApiManagementDiagnosticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementDiagnosticList contains a list of ApiManagementDiagnostics
type ApiManagementDiagnosticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementDiagnostic `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementDiagnosticKind             = "ApiManagementDiagnostic"
	ApiManagementDiagnosticGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementDiagnosticKind}.String()
	ApiManagementDiagnosticKindAPIVersion   = ApiManagementDiagnosticKind + "." + GroupVersion.String()
	ApiManagementDiagnosticGroupVersionKind = GroupVersion.WithKind(ApiManagementDiagnosticKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementDiagnostic{}, &ApiManagementDiagnosticList{})
}
