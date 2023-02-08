package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AzureClusterProviderStatus is the schema for the azureclusterproviderstatus API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type AzureClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Network encapsulates all things related to Azure network.
	Network NetworkSpec `json:"network,omitempty"`

	kaasv1alpha1.ClusterStatusMixin `json:",inline"`
}

func (s *AzureClusterProviderStatus) GetClusterStatusMixin() *kaasv1alpha1.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AzureClusterProviderStatus{})
}
