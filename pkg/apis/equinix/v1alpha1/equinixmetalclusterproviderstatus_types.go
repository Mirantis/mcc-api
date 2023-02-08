package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ElasticIPReservation struct {
	// purpose of public address block, see ElasticIPPurpose
	Purpose ElasticIPPurpose `json:"purpose"`
	// number of IPs in a block, can be 1, 2, 4, 8, 16
	Quantity int `json:"quantity"`
	// subnet address in CIDR notation, read-only, it's empty when block is not allocated yet
	CIDR string `json:"cidr,omitempty"`
	// Equinix resource ID, read-only, it's empty when block is not allocated yet
	ID string `json:"id"`
	// Equinix elastic IP reservation state, can be "created" or "pending"
	State string `json:"state"`
}

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
	ElasticIPBlocks []ElasticIPReservation `json:"elasticIPBlocks"`

	// MetalLB IP ranges that were allocated for cluster
	MetalLBRanges []string `json:"metalLBRanges,omitempty"`
}

func (s *EquinixMetalClusterProviderStatus) GetClusterStatusMixin() *kaasv1alpha1.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderStatus{})
}
