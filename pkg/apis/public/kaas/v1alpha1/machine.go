package v1alpha1

import (
	"github.com/Mirantis/mcc-api/pkg/apis/common/lcm/v1alpha1"
)

const (
	KubeletCondition          ConditionType = "Kubelet"
	LCMCondition              ConditionType = "LCM"
	ProviderInstanceCondition ConditionType = "ProviderInstance"
	SwarmCondition            ConditionType = "Swarm"
	MaintenanceCondition      ConditionType = "Maintenance"
)

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
}

type NodeLabel struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type PublicKeyRef struct {
	Name string `json:"name,omitempty"`
}

type InstanceState struct {
	ID    string `json:"id"`
	State string `json:"state"`
	Ready bool   `json:"ready"`
}

type MachineStatusMixin struct {
	// The private IPv4 address assigned to the instance.
	// +optional
	PrivateIP             string                   `json:"privateIp,omitempty"`
	PublicIP              string                   `json:"publicIp,omitempty"`
	Status                v1alpha1.LCMMachineState `json:"status,omitempty"`
	Hardware              MachineHardware          `json:"hardware,omitempty"`
	ProviderInstanceState InstanceState            `json:"providerInstanceState,omitempty"`
	Maintenance           bool                     `json:"maintenance,omitempty"`
	ConditionsSummary
}

type MachineHardware struct {
	CPU int `json:"cpu,omitempty"`
	// RAM in GB
	RAM     int               `json:"ram,omitempty"`
	Storage []*MachineStorage `json:"storage,omitempty"`
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
