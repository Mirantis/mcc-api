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

	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
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
	BGP BGPConfig `json:"bgp,omitempty"`
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

type BGPConfig struct {
	// Type is BGP deployment type, can be local or global
	Type string `json:"type,omitempty"`
	// MyASN is the value of BGP ASN
	MyASN int `json:"myAsn,omitempty"`
	// BIRDPeers contains the list of BGP peers to be used in bird
	BIRDPeers []BGPPeer `json:"birdPeers,omitempty"`
	// MetalLBPeers contains the list of BGP peers to be used in metallb
	MetalLBPeers []BGPPeer `json:"metallbPeers,omitempty"`
}

type BGPPeer struct {
	PeerAs  int      `json:"peerAs"`
	PeerIPs []string `json:"peerIps"`
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
