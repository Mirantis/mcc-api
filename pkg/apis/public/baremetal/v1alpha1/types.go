package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	kaas "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+resource:path=project

// BaremetalClusterProviderSpec is the providerSpec for metal3 in the cluster object
// +k8s:openapi-gen=true
type BaremetalClusterProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.ClusterSpecMixin `json:",inline"`

	// A token that is used for both validation of the practically of the API server from a
	// joining node's point of view and as an authentication method for the node in the
	// bootstrap phase of "kubeadm join"
	// TODO Temporary workaround. Should be removed eventually.
	BootstrapToken string `json:"bootstrapToken,omitempty"`

	// When true, dedicated MetalLB address pools will be used for services in management/regional cluster;
	// baremetal operator services will use dedicated metallb address pool associated with PXE network.
	DedicatedMetallbPools bool `json:"dedicatedMetallbPools"`

	LoadBalancerHost string `json:"loadBalancerHost"`
}

// The naming conventions for bridge interfaces to the network types
// that are defined for managed Kubernetes clusters on bare metal.
const (
	// Connects LCM agents running on the hosts to the Container Cloud LCM API.
	// The LCM API is provided by the regional or management cluster.
	KubernetesLCMBridgeName = "k8s-lcm"
	// Kubernetes Workloads (pods) network serves internal connections
	// between K8s pods running on different nodes. Kubernetes pods subnet
	// provides IP addresses used by Calico to connect pods between nodes.
	KubernetesPodsBridgeName = "k8s-pods"
)

func (s *BaremetalClusterProviderSpec) GetClusterSpecMixin() *kaas.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}

func (*BaremetalClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &BaremetalClusterProviderStatus{}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+resource:path=project

// BaremetalClusterProviderStatus contains the status fields
// relevant to Baremetal in the cluster object.
// +k8s:openapi-gen=true
type BaremetalClusterProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.ClusterStatusMixin `json:",inline"`
	ErrorMessage            string `json:"errorMessage,omitempty"`
}

func (s *BaremetalClusterProviderStatus) GetClusterStatusMixin() *kaas.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

func init() {
	SchemeBuilder.Register(&BaremetalClusterProviderSpec{})
	SchemeBuilder.Register(&BaremetalClusterProviderStatus{})
}
