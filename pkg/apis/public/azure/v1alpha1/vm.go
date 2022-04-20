/*
Copyright 2021 The Mirantis Authors.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// ProvisioningState describes the provisioning state of an Azure resource.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Migrating ...
	Migrating ProvisioningState = "Migrating"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
	// Deleted represents a deleted VM
	// NOTE: This state is specific to capz, and does not have corresponding mapping in Azure API (https://docs.microsoft.com/en-us/azure/virtual-machines/states-billing#provisioning-states)
	Deleted ProvisioningState = "Deleted"
)

// VMIdentity defines the identity of the virtual machine, if configured.
// +kubebuilder:validation:Enum=None;SystemAssigned;UserAssigned
type VMIdentity string

const (
	// VMIdentityNone ...
	VMIdentityNone VMIdentity = "None"
	// VMIdentitySystemAssigned ...
	VMIdentitySystemAssigned VMIdentity = "SystemAssigned"
	// VMIdentityUserAssigned ...
	VMIdentityUserAssigned VMIdentity = "UserAssigned"
)

// UserAssignedIdentity defines the user-assigned identities provided
// by the user to be assigned to Azure resources.
type UserAssignedIdentity struct {
	// ProviderID is the identification ID of the user-assigned Identity, the format of an identity is:
	// 'azure:///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'
	ProviderID string `json:"providerID"`
}

// IdentityType represents different types of identities.
// +kubebuilder:validation:Enum=ServicePrincipal;UserAssignedMSI
type IdentityType string

// VM describes an Azure virtual machine.
type VM struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// Hardware profile
	VMSize string `json:"vmSize,omitempty"`
	// Storage profile
	Image         Image  `json:"image,omitempty"`
	OSDisk        OSDisk `json:"osDisk,omitempty"`
	StartupScript string `json:"startupScript,omitempty"`
	// State - The provisioning state, which only appears in the response.
	State    ProvisioningState `json:"vmState,omitempty"`
	Identity VMIdentity        `json:"identity,omitempty"`
	Tags     Tags              `json:"tags,omitempty"`

	// Addresses contains the addresses associated with the Azure VM.
	Addresses []corev1.NodeAddress `json:"addresses,omitempty"`
}
