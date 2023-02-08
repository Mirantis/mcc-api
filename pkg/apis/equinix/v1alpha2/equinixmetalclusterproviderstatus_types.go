package v1alpha2

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EquinixMetalClusterProviderStatus is the schema for the equinixmetalclusterproviderstatus API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type EquinixMetalClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.ClusterStatusMixin `json:",inline"`

	// Elastic IP blocks that were requested for cluster
	ElasticIPBlocks []v1alpha1.ElasticIPReservation `json:"elasticIPBlocks"`

	// MetalLB IP ranges that were allocated for cluster
	MetalLBRanges []string `json:"metalLBRanges,omitempty"`

	// VlanStatus describes resources created for managing requested VLANs
	VlanStatus map[string]VlanStatus `json:"vlans"`
}

func (s *EquinixMetalClusterProviderStatus) GetClusterStatusMixin() *kaasv1alpha1.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

type VlanStatus struct {
	// SubnetID refers to a subnet created for managing the VLAN
	SubnetID string `json:"subnetID,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderStatus{})
}
