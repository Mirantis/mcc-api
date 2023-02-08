package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type CephClusterConfig struct {
	// ManualConfiguration is an option that indicates that
	// Ceph node roles should be configured by user
	ManualConfiguration bool `json:"manualConfiguration,omitempty"`
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
	BGP BGPConfig `json:"bgp,omitempty"`
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

type BGPConfig struct {
	// Type is BGP deployment type, can be local or global
	Type string `json:"type,omitempty" sensitive:"true"`
	// MyASN is the value of BGP ASN
	MyASN int `json:"myAsn,omitempty" sensitive:"true"`
	// BIRDPeers contains the list of BGP peers to be used in bird
	BIRDPeers []BGPPeer `json:"birdPeers,omitempty" sensitive:"true"`
	// MetalLBPeers contains the list of BGP peers to be used in MetalLB
	MetalLBPeers []BGPPeer `json:"metallbPeers,omitempty" sensitive:"true"`
}
type BGPPeer struct {
	PeerAs  int      `json:"peerAs"`
	PeerIPs []string `json:"peerIps"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderSpec{})
}
