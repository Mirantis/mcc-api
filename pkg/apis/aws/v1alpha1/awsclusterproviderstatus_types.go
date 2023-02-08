package v1alpha1

import (
	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSClusterProviderStatus contains the status fields
// relevant to AWS in the cluster object.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type AWSClusterProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.ClusterStatusMixin `json:",inline"`

	Network Network  `json:"network,omitempty"`
	Bastion Instance `json:"bastion,omitempty"`
}

func (s *AWSClusterProviderStatus) GetClusterStatusMixin() *kaas.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AWSClusterProviderStatus{})
}
