package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// ProvisioningState describes the provisioning state of an Azure resource.
// +gocode:public-api=true
type ProvisioningState string

const (
	// VMIdentityNone ...
	// +gocode:public-api=true
	VMIdentityNone VMIdentity = "None"
	// VMIdentitySystemAssigned ...
	// +gocode:public-api=true
	VMIdentitySystemAssigned VMIdentity = "SystemAssigned"
	// VMIdentityUserAssigned ...
	// +gocode:public-api=true
	VMIdentityUserAssigned VMIdentity = "UserAssigned"
)

// VMIdentity defines the identity of the virtual machine, if configured.
// +kubebuilder:validation:Enum=None;SystemAssigned;UserAssigned
// +gocode:public-api=true
type VMIdentity string

const (
	// Creating ...
	// +gocode:public-api=true
	Creating ProvisioningState = "Creating"
	// Deleting ...
	// +gocode:public-api=true
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	// +gocode:public-api=true
	Failed ProvisioningState = "Failed"
	// Migrating ...
	// +gocode:public-api=true
	Migrating ProvisioningState = "Migrating"
	// Succeeded ...
	// +gocode:public-api=true
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	// +gocode:public-api=true
	Updating ProvisioningState = "Updating"
	// Deleted represents a deleted VM
	// NOTE: This state is specific to capz, and does not have corresponding mapping in Azure API (https://docs.microsoft.com/en-us/azure/virtual-machines/states-billing#provisioning-states)
	// +gocode:public-api=true
	Deleted ProvisioningState = "Deleted"
)

// IdentityType represents different types of identities.
// +kubebuilder:validation:Enum=ServicePrincipal;UserAssignedMSI
// +gocode:public-api=true
type IdentityType string

// VM describes an Azure virtual machine.
// +gocode:public-api=true
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

// UserAssignedIdentity defines the user-assigned identities provided
// by the user to be assigned to Azure resources.
// +gocode:public-api=true
type UserAssignedIdentity struct {
	// ProviderID is the identification ID of the user-assigned Identity, the format of an identity is:
	// 'azure:///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'
	ProviderID string `json:"providerID"`
}
