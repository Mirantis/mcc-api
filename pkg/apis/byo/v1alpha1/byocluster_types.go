package v1alpha1

import (
	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// BYOClusterProviderSpec is the providerConfig for BYO in the cluster
// object
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BYOClusterProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.ClusterSpecMixin `json:",inline"`
}

func (s *BYOClusterProviderSpec) GetClusterSpecMixin() *kaas.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}
func (*BYOClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &BYOClusterProviderStatus{}
}

// BYOClusterProviderStatus contains the status fields
// relevant to BYO in the cluster object.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BYOClusterProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.ClusterStatusMixin `json:",inline"`
}

func (s *BYOClusterProviderStatus) GetClusterStatusMixin() *kaas.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&BYOClusterProviderSpec{})
	SchemeBuilder.Register(&BYOClusterProviderStatus{})
}
