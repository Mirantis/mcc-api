package v1alpha1

// OSDisk defines the operating system disk for a VM.
//
// WARNING: this requires any updates to ManagedDisk to be manually converted. This is due to the odd issue with
// conversion-gen where the warning message generated uses a relative directory import rather than the fully
// qualified import when generating outside of the GOPATH.
// +gocode:public-api=true
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

// ManagedDiskParameters defines the parameters of a managed disk.
type ManagedDiskParameters struct {
	// +optional
	StorageAccountType string `json:"storageAccountType,omitempty"`
	// +optional
	DiskEncryptionSet *DiskEncryptionSetParameters `json:"diskEncryptionSet,omitempty"`
}
