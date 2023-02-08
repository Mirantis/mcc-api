package v1alpha1

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/lcm/v1alpha1"
)

// +gocode:public-api=true
type MachineStatusMixin struct {
	// The private IPv4 address assigned to the instance.
	// +optional
	PrivateIP             string                      `json:"privateIp,omitempty"`
	PublicIP              string                      `json:"publicIp,omitempty"`
	Status                v1alpha1.LCMMachineState    `json:"status,omitempty"`
	Hardware              MachineHardware             `json:"hardware,omitempty"`
	ProviderInstanceState InstanceState               `json:"providerInstanceState,omitempty"`
	Maintenance           bool                        `json:"maintenance,omitempty"`
	UpgradeIndex          int                         `json:"upgradeIndex,omitempty"`
	Reboot                *MachineReboot              `json:"reboot,omitempty"`
	Delete                bool                        `json:"delete,omitempty"`
	PrepareDeletionPhase  MachinePrepareDeletionPhase `json:"prepareDeletionPhase,omitempty"`
	ConditionsSummary
}

const (
	// +gocode:public-api=true
	KubeletCondition ConditionType = "Kubelet"
	// +gocode:public-api=true
	LCMCondition ConditionType = "LCM"
	// +gocode:public-api=true
	ProviderInstanceCondition ConditionType = "ProviderInstance"
	// +gocode:public-api=true
	SwarmCondition ConditionType = "Swarm"
	// +gocode:public-api=true
	MaintenanceCondition ConditionType = "Maintenance"
	// +gocode:public-api=true
	GracefulRebootCondition ConditionType = "GracefulReboot"
	// +gocode:public-api=true
	RebootCondition ConditionType = "Reboot"
	// +gocode:public-api=true
	PrepareDeletionCondition ConditionType = "PrepareDeletion"
)

type MachineReboot struct {
	// Indicates that the OS is waiting for a reboot
	Required bool `json:"required"`
	// Information about the reasons why a reboot is required
	Reason string `json:"reason"`
}

const (
	// +gocode:public-api=true
	MachinePrepareDeletionPhaseStarted MachinePrepareDeletionPhase = "started"
	// +gocode:public-api=true
	MachinePrepareDeletionPhaseCompleted MachinePrepareDeletionPhase = "completed"
	// +gocode:public-api=true
	MachinePrepareDeletionPhaseAborting MachinePrepareDeletionPhase = "aborting"
	// +gocode:public-api=true
	MachinePrepareDeletionPhaseFailed MachinePrepareDeletionPhase = "failed"
)

// +gocode:public-api=true
type MachineSpecMixin struct {
	NodeLabels []NodeLabel `json:"nodeLabels,omitempty"`
	// The name of the RHELLicense object
	RHELLicense string `json:"rhelLicense,omitempty"`

	// Distribution represents ID of `Distribution` object inside list of
	// allowed distributions on per-cluster release basis.
	// +optional
	Distribution string `json:"distribution,omitempty"`
	// Maintenance defines if machine should switch into maintenance state
	Maintenance bool `json:"maintenance,omitempty"`
	// UpgradeIndex is a positive value that
	// defines the order in which machines are upgraded
	// (machines with lower indexes upgraded first)
	// +optional
	UpgradeIndex *int `json:"upgradeIndex,omitempty"`

	// Delete defines if machine should be deleted.
	Delete bool `json:"delete,omitempty"`
	// DeletionPolicy defines the policy used to describe machine deletion workflow.
	// Defaults to "unsafe" - delete machine without cordon & drain workload.
	// "gracefull" - proceed cordon & drain node before deleting machine.
	DeletionPolicy MachineDeletionPolicy `json:"deletionPolicy,omitempty"`
}
type (
	// +gocode:public-api=true
	MachineDeletionPolicy string
	// +gocode:public-api=true
	MachinePrepareDeletionPhase string
)

// +gocode:public-api=true
type MachineHardware struct {
	CPU int `json:"cpu,omitempty"`
	// RAM in GB
	RAM     int               `json:"ram,omitempty"`
	Storage []*MachineStorage `json:"storage,omitempty"`
}

// +gocode:public-api=true
type PublicKeyRef struct {
	Name string `json:"name,omitempty"`
}
type MachineStorage struct {
	Name string `json:"name,omitempty"`
	// Size in GiB
	Size int `json:"size,omitempty"`
	// Type is a drive storage type: ssd, hdd, nvme
	// +kubebuilder:validation:Enum=hdd;ssd;nvme
	Type string `json:"type,omitempty"`
	// ByID is a disk ID, e.g. /dev/disk/by-id/nvme-SAMSUNG_MFOOA1234BAR-00042_S42XXX1K710069
	ByID string `json:"byID,omitempty"`
	// ByPath is a disk path, e.g. /dev/disk/by-path/pci-0000:00:07.0
	ByPath string `json:"byPath,omitempty"`
	// IsBoot reflects if the operating system was booted from that drive
	IsBoot bool `json:"isBoot,omitempty"`
	// IsLVP reflects if the LVP partition was created on that drive
	IsLVP bool `json:"isLVP,omitempty"`
	// IsCeph reflects that this drive will be allocated for Ceph storage
	IsCeph bool `json:"isCeph,omitempty"`
}

const (
	// +gocode:public-api=true
	MachineDeletionPolicyGraceful MachineDeletionPolicy = "graceful"
	// +gocode:public-api=true
	MachineDeletionPolicyUnsafe MachineDeletionPolicy = "unsafe"
	// +gocode:public-api=true
	MachineDeletionPolicyForced MachineDeletionPolicy = "forced"
	// +gocode:public-api=true
	MachineDeletionPolicyDefault MachineDeletionPolicy = MachineDeletionPolicyUnsafe
)

type NodeLabel struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

// +gocode:public-api=true
type InstanceState struct {
	ID    string `json:"id"`
	State string `json:"state"`
	Ready bool   `json:"ready"`
}
