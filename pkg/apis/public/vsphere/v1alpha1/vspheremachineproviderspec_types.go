/*
Copyright 2022 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VsphereMachineProviderSpec is the schema for the vspheremachineproviderspec API
// +k8s:openapi-gen=true
type VsphereMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineSpecMixin `json:",inline"`

	// This value is set automatically at runtime and should not be set or
	// modified by users.
	// MachineRef is used to lookup the VM.
	// +optional
	MachineRef string `json:"machineRef,omitempty"`

	// Template is the name, inventory path, or instance UUID of the template
	// used to clone new machines.
	Template string `json:"template"`

	// Network is the network configuration for this machine's VM.
	Network NetworkSpec `json:"network"`

	// NumCPUs is the number of virtual processors in a virtual machine.
	// Defaults to the analogue property value in the template from which this
	// machine is cloned.
	// +optional
	NumCPUs int32 `json:"numCPUs,omitempty"`
	// NumCPUs is the number of cores among which to distribute CPUs in this
	// virtual machine.
	// Defaults to the analogue property value in the template from which this
	// machine is cloned.
	// +optional
	NumCoresPerSocket int32 `json:"numCoresPerSocket,omitempty"`
	// MemoryMiB is the size of a virtual machine's memory, in MiB.
	// Defaults to the analogue property value in the template from which this
	// machine is cloned.
	// +optional
	MemoryMiB int64 `json:"memoryMiB,omitempty"`
	// DiskGiB is the size of a virtual machine's disk, in GiB.
	// Defaults to the analogue property value in the template from which this
	// machine is cloned.
	// +optional
	DiskGiB int32 `json:"diskGiB,omitempty"`

	// TrustedCerts is a list of trusted certificates to add to the machine's VM.
	// +optional
	TrustedCerts [][]byte `json:"trustedCerts,omitempty"`

	// NTPServers is a list of NTP servers to use instead of the machine image's
	// default NTP server list.
	// +optional
	NTPServers []string `json:"ntpServers,omitempty"`
}

func (s *VsphereMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}

func (*VsphereMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &VsphereMachineProviderStatus{}
}

// NetworkSpec defines the virtual machine's network configuration.
type NetworkSpec struct {
	// Devices is the list of network devices used by the virtual machine.
	// TODO(akutz) Make sure at least one network matches the
	//             ClusterSpec.CloudProviderConfiguration.Network.Name
	Devices []NetworkDeviceSpec `json:"devices"`

	// Routes is a list of optional, static routes applied to the virtual
	// machine.
	// +optional
	Routes []NetworkRouteSpec `json:"routes,omitempty"`

	// PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API
	// server endpoint on this machine
	// +optional
	PreferredAPIServerCIDR string `json:"preferredAPIServerCidr,omitempty"`
}

// NetworkDeviceSpec defines the network configuration for a virtual machine's
// network device.
type NetworkDeviceSpec struct {
	// DHCP4 is a flag that indicates whether or not to use DHCP for IPv4
	// on this device.
	// If true then IPAddrs should not contain any IPv4 addresses.
	// +optional
	DHCP4 bool `json:"dhcp4,omitempty"`

	// DHCP6 is a flag that indicates whether or not to use DHCP for IPv6
	// on this device.
	// If true then IPAddrs should not contain any IPv6 addresses.
	// +optional
	DHCP6 bool `json:"dhcp6,omitempty"`

	// Gateway4 is the IPv4 gateway used by this device.
	// Required when DHCP4 is false.
	// +optional
	Gateway4 string `json:"gateway4,omitempty"`

	// Gateway4 is the IPv4 gateway used by this device.
	// Required when DHCP6 is false.
	// +optional
	Gateway6 string `json:"gateway6,omitempty"`

	// IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign
	// to this device.
	// Required when DHCP4 and DHCP6 are both false.
	// +optional
	IPAddrs []string `json:"ipAddrs,omitempty"`

	// MTU is the device’s Maximum Transmission Unit size in bytes.
	// +optional
	MTU *int64 `json:"mtu,omitempty"`

	// MACAddr is the MAC address used by this device.
	// It is generally a good idea to omit this field and allow a MAC address
	// to be generated.
	// Please note that this value must use the VMware OUI to work with the
	// in-tree vSphere cloud provider.
	// +optional
	MACAddr string `json:"macAddr,omitempty"`

	// Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS
	// nameservers.
	// Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).
	// +optional
	Nameservers []string `json:"nameservers,omitempty"`

	// Routes is a list of optional, static routes applied to the device.
	// +optional
	Routes []NetworkRouteSpec `json:"routes,omitempty"`

	// SearchDomains is a list of search domains used when resolving IP
	// addresses with DNS.
	// +optional
	SearchDomains []string `json:"searchDomains,omitempty"`
}

// NetworkRouteSpec defines a static network route.
type NetworkRouteSpec struct {
	// To is an IPv4 or IPv6 address.
	To string `json:"to"`
	// Via is an IPv4 or IPv6 address.
	Via string `json:"via"`
	// Metric is the weight/priority of the route.
	Metric int32 `json:"metric"`
}

func init() {
	SchemeBuilder.Register(&VsphereMachineProviderSpec{})
}
