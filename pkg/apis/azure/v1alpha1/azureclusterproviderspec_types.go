package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// AzureClusterProviderSpec is the schema for the azureclusterproviderspec API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type AzureClusterProviderSpec struct {
	metav1.TypeMeta               `json:",inline"`
	metav1.ObjectMeta             `json:"metadata,omitempty"`
	kaasv1alpha1.ClusterSpecMixin `json:",inline"`

	// NetworkSpec encapsulates all things related to Azure network.
	NetworkSpec NetworkSpec `json:"networkSpec,omitempty"`

	// BastionSpec encapsulates all things related to the Bastions in the cluster.
	// +optional
	BastionSpec BastionSpec `json:"bastionSpec,omitempty"`

	// Location defines Azure region where the cluster will be deployed
	Location string `json:"location"`
}

func (s *AzureClusterProviderSpec) GetClusterSpecMixin() *kaasv1alpha1.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}
func (*AzureClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &AzureClusterProviderStatus{}
}

// BastionSpec specifies how the Bastion feature should be set up for the cluster.
// +gocode:public-api=true
type BastionSpec struct {
	// +optional
	AzureBastion *AzureBastion `json:"azureBastion,omitempty"`
}

// AzureBastion specifies how the Azure Bastion cloud component should be configured.
type AzureBastion struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Subnet SubnetSpec `json:"subnet,omitempty"`
	// +optional
	PublicIP PublicIPSpec `json:"publicIP,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AzureClusterProviderSpec{})
}
