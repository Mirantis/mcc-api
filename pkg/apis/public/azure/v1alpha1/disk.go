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

// OSDisk defines the operating system disk for a VM.
//
// WARNING: this requires any updates to ManagedDisk to be manually converted. This is due to the odd issue with
// conversion-gen where the warning message generated uses a relative directory import rather than the fully
// qualified import when generating outside of the GOPATH.
type OSDisk struct {
	OSType string `json:"osType"`
	// DiskSizeGB is the size in GB to assign to the OS disk.
	// Will have a default of 30GB if not provided
	// +optional
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// ManagedDisk specifies the Managed Disk parameters for the OS disk.
	// +optional
	ManagedDisk      *ManagedDiskParameters `json:"managedDisk,omitempty"`
	DiffDiskSettings *DiffDiskSettings      `json:"diffDiskSettings,omitempty"`
	// CachingType specifies the caching requirements.
	// +optional
	// +kubebuilder:validation:Enum=None;ReadOnly;ReadWrite
	CachingType string `json:"cachingType,omitempty"`
}

// ManagedDiskParameters defines the parameters of a managed disk.
type ManagedDiskParameters struct {
	// +optional
	StorageAccountType string `json:"storageAccountType,omitempty"`
	// +optional
	DiskEncryptionSet *DiskEncryptionSetParameters `json:"diskEncryptionSet,omitempty"`
}

// DiskEncryptionSetParameters defines disk encryption options.
type DiskEncryptionSetParameters struct {
	// ID defines resourceID for diskEncryptionSet resource. It must be in the same subscription
	ID string `json:"id,omitempty"`
}

// DiffDiskSettings describe ephemeral disk settings for the os disk.
type DiffDiskSettings struct {
	// Option enables ephemeral OS when set to "Local"
	// See https://docs.microsoft.com/en-us/azure/virtual-machines/ephemeral-os-disks for full details
	// +kubebuilder:validation:Enum=Local
	Option string `json:"option"`
}
