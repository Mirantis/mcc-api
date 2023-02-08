package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkStatus provides information about one of a VM's networks.
type NetworkStatus struct {
	// Connected is a flag that indicates whether this network is currently
	// connected to the VM.
	Connected bool `json:"connected,omitempty"`

	// IPAddrs is one or more IP addresses reported by vm-tools.
	// +optional
	IPAddrs []string `json:"ipAddrs,omitempty"`

	// MACAddr is the MAC address of the network device.
	MACAddr string `json:"macAddr"`

	// NetworkName is the name of the network.
	// +optional
	NetworkName string `json:"networkName,omitempty"`
}

// VsphereMachineProviderStatus is the schema for the vspheremachineproviderstatus API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type VsphereMachineProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineStatusMixin `json:",inline"`

	// This value is set automatically at runtime and should not be set or
	// modified by users.
	// MachineRef is used to lookup the VM.
	// +optional
	MachineRef string `json:"machineRef,omitempty"`

	// TaskRef is a managed object reference to a Task related to the machine.
	// This value is set automatically at runtime and should not be set or
	// modified by users.
	// +optional
	TaskRef string `json:"taskRef,omitempty"`

	// Network returns the network status for each of the machine's configured
	// network interfaces.
	// +optional
	Network []NetworkStatus `json:"networkStatus,omitempty"`
}

func (s *VsphereMachineProviderStatus) GetMachineStatusMixin() *kaasv1alpha1.MachineStatusMixin {
	return &s.MachineStatusMixin
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&VsphereMachineProviderStatus{})
}
