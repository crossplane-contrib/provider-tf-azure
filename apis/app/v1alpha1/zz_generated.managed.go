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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AppConfiguration.
func (mg *AppConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppConfiguration.
func (mg *AppConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppConfiguration.
func (mg *AppConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppConfiguration.
func (mg *AppConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppConfiguration.
func (mg *AppConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppConfiguration.
func (mg *AppConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppConfiguration.
func (mg *AppConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppConfiguration.
func (mg *AppConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppService.
func (mg *AppService) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppService.
func (mg *AppService) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppService.
func (mg *AppService) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppService.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppService) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppService.
func (mg *AppService) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppService.
func (mg *AppService) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppService.
func (mg *AppService) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppService.
func (mg *AppService) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppService.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppService) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppService.
func (mg *AppService) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceActiveSlot.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceActiveSlot) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceActiveSlot.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceActiveSlot) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceActiveSlot.
func (mg *AppServiceActiveSlot) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceCertificate.
func (mg *AppServiceCertificate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceCertificate.
func (mg *AppServiceCertificate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceCertificate.
func (mg *AppServiceCertificate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceCertificate.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceCertificate) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceCertificate.
func (mg *AppServiceCertificate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceCertificate.
func (mg *AppServiceCertificate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceCertificate.
func (mg *AppServiceCertificate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceCertificate.
func (mg *AppServiceCertificate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceCertificate.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceCertificate) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceCertificate.
func (mg *AppServiceCertificate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceCertificateBinding.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceCertificateBinding) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceCertificateBinding.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceCertificateBinding) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceCertificateBinding.
func (mg *AppServiceCertificateBinding) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceCertificateOrder.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceCertificateOrder) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceCertificateOrder.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceCertificateOrder) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceCustomHostnameBinding.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceCustomHostnameBinding) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceCustomHostnameBinding.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceCustomHostnameBinding) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceCustomHostnameBinding.
func (mg *AppServiceCustomHostnameBinding) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceEnvironment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceEnvironment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceEnvironment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceEnvironment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceEnvironment.
func (mg *AppServiceEnvironment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceEnvironmentV3.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceEnvironmentV3) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceEnvironmentV3.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceEnvironmentV3) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceEnvironmentV3.
func (mg *AppServiceEnvironmentV3) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceHybridConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceHybridConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceHybridConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceHybridConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceHybridConnection.
func (mg *AppServiceHybridConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceManagedCertificate.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceManagedCertificate) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceManagedCertificate.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceManagedCertificate) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceManagedCertificate.
func (mg *AppServiceManagedCertificate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServicePlan.
func (mg *AppServicePlan) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServicePlan.
func (mg *AppServicePlan) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServicePlan.
func (mg *AppServicePlan) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServicePlan.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServicePlan) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServicePlan.
func (mg *AppServicePlan) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServicePlan.
func (mg *AppServicePlan) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServicePlan.
func (mg *AppServicePlan) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServicePlan.
func (mg *AppServicePlan) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServicePlan.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServicePlan) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServicePlan.
func (mg *AppServicePlan) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceSlot.
func (mg *AppServiceSlot) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceSlot.
func (mg *AppServiceSlot) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceSlot.
func (mg *AppServiceSlot) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceSlot.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceSlot) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceSlot.
func (mg *AppServiceSlot) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceSlot.
func (mg *AppServiceSlot) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceSlot.
func (mg *AppServiceSlot) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceSlot.
func (mg *AppServiceSlot) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceSlot.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceSlot) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceSlot.
func (mg *AppServiceSlot) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceSlotVirtualNetworkSwiftConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceSlotVirtualNetworkSwiftConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceSlotVirtualNetworkSwiftConnection.
func (mg *AppServiceSlotVirtualNetworkSwiftConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceSourceControlToken.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceSourceControlToken) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceSourceControlToken.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceSourceControlToken) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceSourceControlToken.
func (mg *AppServiceSourceControlToken) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceVirtualNetworkSwiftConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceVirtualNetworkSwiftConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceVirtualNetworkSwiftConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceVirtualNetworkSwiftConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppServiceVirtualNetworkSwiftConnection.
func (mg *AppServiceVirtualNetworkSwiftConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
