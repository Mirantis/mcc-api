package v1alpha2

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type CephClusterConfig struct {
	// Disabled flag indicates that Ceph is disabled for that cluster and
	// KaaSCephCluster will not be deployed. If true, applies for Equinix Metal v2
	// management and regional clusters only.
	Disabled bool `json:"disabled,omitempty"`
	// ManualConfiguration is an option that indicates that
	// Ceph node roles should be configured by user
	ManualConfiguration bool `json:"manualConfiguration,omitempty"`
}
type Network struct {
	// Custom field indicates that the networking configuration for each machine
	// will be configured manually by the user via Subnet and L2Template resources.
	// If set to true, the equinix provider will not generate default Subnets and
	// L2Templates for the cluster.
	// CIDR, IncludeRanges, ExcludeRanges, Gateway and Nameservers may be empty.
	// VlanID and AdditionalVlans may also be empty, but if set they will be attached
	// to each machine in the cluster if the machines' Network.VLANs are not specified.
	Custom bool `json:"custom,omitempty"`
	// CIDR is the address of network in CIDR notation to be allocated to machines
	CIDR string `json:"cidr,omitempty" sensitive:"true"`
	// IncludeRanges defines IP addresses for cluster machines
	// Must be in the CIDR range
	IncludeRanges []string `json:"includeRanges,omitempty" sensitive:"true"`
	// ExcludeRanges defines IP addresses to exclude from the allocation to cluster machines
	// Must be in the CIDR range
	ExcludeRanges []string `json:"excludeRanges,omitempty" sensitive:"true"`
	// Gateway is the default router on the relevant network
	Gateway string `json:"gateway,omitempty" sensitive:"true"`
	// Nameservers defines an external DNS servers accessible from the relevant network.
	Nameservers []string `json:"nameservers,omitempty" sensitive:"true"`
	// DHCPRanges defines the list of IP ranges to be used in DHCP service
	DHCPRanges []string `json:"dhcpRanges,omitempty" sensitive:"true"`
	// VlanID is an identifier of VLAN to attach to cluster nodes
	VlanID string `json:"vlanId,omitempty" sensitive:"true"`
	// AdditionalVlans defines VLANs to be attached to cluster machines
	AdditionalVlans []Vlan `json:"additionalVlans,omitempty" sensitive:"true"`
	// LoadBalancerHost is the IP address of the load balancer host
	LoadBalancerHost string `json:"loadBalancerHost,omitempty" sensitive:"true"`
	// MetallbRanges is the array of address ranges for MetalLB usage
	MetallbRanges []string `json:"metallbRanges,omitempty" sensitive:"true"`
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

// +gocode:public-api=true
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

// EquinixMetalClusterProviderSpec is the schema for the equinixmetalclusterproviderspec API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
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

func (s *EquinixMetalClusterProviderSpec) GetClusterSpecMixin() *kaasv1alpha1.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}
func (*EquinixMetalClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &EquinixMetalClusterProviderStatus{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderSpec{})
}
