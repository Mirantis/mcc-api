package v1alpha2

import (
	commonerrors "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/common"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EquinixMetalMachineProviderStatus is the schema for the equinixmetalmachineproviderstatus API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type EquinixMetalMachineProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineStatusMixin `json:",inline"`

	// Addresses contains the EquinixMetal device associated addresses.
	Addresses []corev1.NodeAddress `json:"addresses,omitempty"`

	// BGP is enabled for node, read-only, result of scope.BGPEnabled()
	BGPEnabled bool `json:"bgpEnabled,omitempty"`

	// IPv4 BGP neighbors of node, read-only, result of scope.GetBGPNeighbors()
	BGPNeighbors []v1alpha1.BGPNeighborPeers `json:"bgpNeighbors,omitempty"`

	//  IPv4 private networks of node, read-only, result of packet.GetDeviceAddresses()
	PrivateNetworks []v1alpha1.PrivateNetwork `json:"privateNetworks,omitempty"`

	// VLANAddresses stores IP addresses assigned to additional VLANs if configured
	VLANAddresses map[string]string `json:"vlanAddresses,omitempty"`
}

func (s *EquinixMetalMachineProviderStatus) GetMachineStatusMixin() *kaasv1alpha1.MachineStatusMixin {
	return &s.MachineStatusMixin
}

const (
	// +gocode:public-api=true
	NonCompliantStorageError commonerrors.MachineStatusError = "NonCompliantStorageError"
	// +gocode:public-api=true
	InspectionError commonerrors.MachineStatusError = "InspectionError"
	// +gocode:public-api=true
	ProvisioningError commonerrors.MachineStatusError = "ProvisioningError"
	// +gocode:public-api=true
	MachineInaccessibleError commonerrors.MachineStatusError = "MachineInaccessibleError"
)

// NetworkStatus provides information about one of a VM's networks.
// +gocode:public-api=true
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalMachineProviderStatus{})
}
