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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Mirantis/mcc-api/pkg/apis/public/equinix/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalClusterProviderSpec is the schema for the equinixmetalclusterproviderspec API
// +k8s:openapi-gen=true
type EquinixMetalClusterProviderSpec struct {
	metav1.TypeMeta               `json:",inline"`
	metav1.ObjectMeta             `json:"metadata,omitempty"`
	kaasv1alpha1.ClusterSpecMixin `json:",inline"`

	// Facility represents the Packet facility for this cluster
	Facility string `json:"facility"`
	// Ceph contains Ceph configuration options for Cluster
	Ceph CephClusterConfig `json:"ceph,omitempty"`
	// Elastic IPs to be requested for k8s services LBs
	ServiceLbIPQuantity int `json:"serviceLbIPQuantity,omitempty"`
	// BGP section contains BGP configuration
	BGP v1alpha1.BGPConfig `json:"bgp,omitempty"`
	// Network section contains network configuration for the cluster
	Network Network `json:"network,omitempty"`
	// ProjectSSHKeys is the list of Equinix Project SSH Key names to be attached to the machines.
	// Project SSH Keys should be already created in Equinix console.
	// Is it needed in order to have access to Out-Of-Band console to debug provisioning failures
	// because MCC does not add any Project SSH Key to machine by default.
	// Details: https://metal.equinix.com/developers/docs/resilience-recovery/serial-over-ssh/
	ProjectSSHKeys []string `json:"projectSSHKeys,omitempty"`
}

type CephClusterConfig struct {
	// ManualConfiguration is an option that indicates that
	// Ceph node roles should be configured by user
	ManualConfiguration bool `json:"manualConfiguration,omitempty"`
}

type Network struct {
	// CIDR is the address of network in CIDR notation to be allocated to machines
	CIDR string `json:"cidr,omitempty"`
	// IncludeRanges defines IP addresses for cluster machines
	// Must be in the CIDR range
	IncludeRanges []string `json:"includeRanges,omitempty"`
	// ExcludeRanges defines IP addresses to exclude from the allocation to cluster machines
	// Must be in the CIDR range
	ExcludeRanges []string `json:"excludeRanges,omitempty"`
	// Gateway is the default router on the relevant network
	Gateway string `json:"gateway,omitempty"`
	// Nameservers defines an external DNS servers accessible from the relevant network.
	Nameservers []string `json:"nameservers,omitempty"`
	// DHCPRanges defines the list of IP ranges to be used in DHCP service
	DHCPRanges []string `json:"dhcpRanges,omitempty"`
	// VlanID is an identifier of VLAN to attach to cluster nodes
	VlanID string `json:"vlanId,omitempty"`
	// AdditionalVlans defines VLANs to be attached to cluster machines
	AdditionalVlans []Vlan `json:"additionalVlans,omitempty"`
	// LoadBalancerHost is the IP address of the load balancer host
	LoadBalancerHost string `json:"loadBalancerHost,omitempty"`
	// MetallbRanges is the array of address ranges for metallb usage
	MetallbRanges []string `json:"metallbRanges,omitempty"`
}

type Vlan struct {
	// ID specifies VLAN ID
	ID string `json:"id"`
	// CIDR is the address of network in CIDR notation to be allocated to machines
	CIDR string `json:"cidr,omitempty"`
	// IncludeRanges defines IP addresses for cluster machines
	// Must be in the CIDR range
	IncludeRanges []string `json:"includeRanges,omitempty"`
	// ExcludeRanges defines IP addresses to exclude from the allocation to cluster machines
	// Must be in the CIDR range
	ExcludeRanges []string `json:"excludeRanges,omitempty"`
}

type IPAMConfig struct {
	// CIDR is the address of network in CIDR notation to be allocated to machines
	CIDR string `json:"cidr,omitempty"`
	// IncludeRanges defines IP addresses for cluster machines
	// Must be in the CIDR range
	IncludeRanges []string `json:"includeRanges,omitempty"`
	// ExcludeRanges defines IP addresses to exclude from the allocation to cluster machines
	// Must be in the CIDR range
	ExcludeRanges []string `json:"excludeRanges,omitempty"`
	// Gateway is the default router on the relevant network
	Gateway string `json:"gateway,omitempty"`
	// Nameservers defines an external DNS servers accessible from the relevant network.
	Nameservers []string `json:"nameservers,omitempty"`
}

func (s *EquinixMetalClusterProviderSpec) GetClusterSpecMixin() *kaasv1alpha1.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}

func (*EquinixMetalClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &EquinixMetalClusterProviderStatus{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderSpec{})
}
