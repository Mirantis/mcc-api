/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonerrors "github.com/Mirantis/mcc-api/v2/pkg/apis/public/cluster/common"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/public/equinix/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalMachineProviderStatus is the schema for the equinixmetalmachineproviderstatus API
// +k8s:openapi-gen=true
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

const (
	NonCompliantStorageError commonerrors.MachineStatusError = "NonCompliantStorageError"

	InspectionError          commonerrors.MachineStatusError = "InspectionError"
	ProvisioningError        commonerrors.MachineStatusError = "ProvisioningError"
	MachineInaccessibleError commonerrors.MachineStatusError = "MachineInaccessibleError"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&EquinixMetalMachineProviderStatus{})
}
