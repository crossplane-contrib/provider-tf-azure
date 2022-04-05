/*
Copyright 2022 The Crossplane Authors.

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

// GetCondition of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupContainerStorageAccount.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupContainerStorageAccount) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupContainerStorageAccount.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupContainerStorageAccount) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupContainerStorageAccount.
func (mg *BackupContainerStorageAccount) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupPolicyFileShare.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupPolicyFileShare) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupPolicyFileShare.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupPolicyFileShare) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupPolicyFileShare.
func (mg *BackupPolicyFileShare) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupPolicyVM.
func (mg *BackupPolicyVM) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupPolicyVM.
func (mg *BackupPolicyVM) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupPolicyVM.
func (mg *BackupPolicyVM) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupPolicyVM.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupPolicyVM) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupPolicyVM.
func (mg *BackupPolicyVM) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupPolicyVM.
func (mg *BackupPolicyVM) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupPolicyVM.
func (mg *BackupPolicyVM) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupPolicyVM.
func (mg *BackupPolicyVM) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupPolicyVM.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupPolicyVM) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupPolicyVM.
func (mg *BackupPolicyVM) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupProtectedFileShare.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupProtectedFileShare) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupProtectedFileShare.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupProtectedFileShare) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupProtectedFileShare.
func (mg *BackupProtectedFileShare) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupProtectedVM.
func (mg *BackupProtectedVM) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupProtectedVM.
func (mg *BackupProtectedVM) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupProtectedVM.
func (mg *BackupProtectedVM) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupProtectedVM.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupProtectedVM) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupProtectedVM.
func (mg *BackupProtectedVM) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupProtectedVM.
func (mg *BackupProtectedVM) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupProtectedVM.
func (mg *BackupProtectedVM) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupProtectedVM.
func (mg *BackupProtectedVM) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupProtectedVM.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupProtectedVM) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupProtectedVM.
func (mg *BackupProtectedVM) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SiteRecoveryFabric.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SiteRecoveryFabric) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SiteRecoveryFabric.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SiteRecoveryFabric) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this SiteRecoveryFabric.
func (mg *SiteRecoveryFabric) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SiteRecoveryNetworkMapping.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SiteRecoveryNetworkMapping) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SiteRecoveryNetworkMapping.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SiteRecoveryNetworkMapping) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this SiteRecoveryNetworkMapping.
func (mg *SiteRecoveryNetworkMapping) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SiteRecoveryProtectionContainer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SiteRecoveryProtectionContainer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SiteRecoveryProtectionContainer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SiteRecoveryProtectionContainer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this SiteRecoveryProtectionContainer.
func (mg *SiteRecoveryProtectionContainer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SiteRecoveryProtectionContainerMapping.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SiteRecoveryProtectionContainerMapping) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SiteRecoveryProtectionContainerMapping.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SiteRecoveryProtectionContainerMapping) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this SiteRecoveryProtectionContainerMapping.
func (mg *SiteRecoveryProtectionContainerMapping) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SiteRecoveryReplicatedVM.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SiteRecoveryReplicatedVM) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SiteRecoveryReplicatedVM.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SiteRecoveryReplicatedVM) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this SiteRecoveryReplicatedVM.
func (mg *SiteRecoveryReplicatedVM) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SiteRecoveryReplicationPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SiteRecoveryReplicationPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SiteRecoveryReplicationPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SiteRecoveryReplicationPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this SiteRecoveryReplicationPolicy.
func (mg *SiteRecoveryReplicationPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Vault.
func (mg *Vault) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Vault.
func (mg *Vault) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Vault.
func (mg *Vault) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Vault.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Vault) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Vault.
func (mg *Vault) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Vault.
func (mg *Vault) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Vault.
func (mg *Vault) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Vault.
func (mg *Vault) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Vault.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Vault) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Vault.
func (mg *Vault) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
