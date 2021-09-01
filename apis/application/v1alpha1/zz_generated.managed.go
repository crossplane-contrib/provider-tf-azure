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

// GetCondition of this ApplicationGateway.
func (mg *ApplicationGateway) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationGateway.
func (mg *ApplicationGateway) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationGateway.
func (mg *ApplicationGateway) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationGateway.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationGateway) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationGateway.
func (mg *ApplicationGateway) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationGateway.
func (mg *ApplicationGateway) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationGateway.
func (mg *ApplicationGateway) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationGateway.
func (mg *ApplicationGateway) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationGateway.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationGateway) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationGateway.
func (mg *ApplicationGateway) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ApplicationInsights.
func (mg *ApplicationInsights) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationInsights.
func (mg *ApplicationInsights) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationInsights.
func (mg *ApplicationInsights) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationInsights.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationInsights) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationInsights.
func (mg *ApplicationInsights) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationInsights.
func (mg *ApplicationInsights) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationInsights.
func (mg *ApplicationInsights) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationInsights.
func (mg *ApplicationInsights) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationInsights.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationInsights) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationInsights.
func (mg *ApplicationInsights) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationInsightsAnalyticsItem.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationInsightsAnalyticsItem) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationInsightsAnalyticsItem.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationInsightsAnalyticsItem) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationInsightsAnalyticsItem.
func (mg *ApplicationInsightsAnalyticsItem) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationInsightsApiKey.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationInsightsApiKey) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationInsightsApiKey.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationInsightsApiKey) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationInsightsApiKey.
func (mg *ApplicationInsightsApiKey) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationInsightsSmartDetectionRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationInsightsSmartDetectionRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationInsightsSmartDetectionRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationInsightsSmartDetectionRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationInsightsSmartDetectionRule.
func (mg *ApplicationInsightsSmartDetectionRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationInsightsWebTest.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationInsightsWebTest) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationInsightsWebTest.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationInsightsWebTest) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationInsightsWebTest.
func (mg *ApplicationInsightsWebTest) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ApplicationSecurityGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ApplicationSecurityGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ApplicationSecurityGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ApplicationSecurityGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ApplicationSecurityGroup.
func (mg *ApplicationSecurityGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
