package v1alpha1

import (
	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// BaremetalClusterProviderStatus contains the status fields
// relevant to Baremetal in the cluster object.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=project
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BaremetalClusterProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.ClusterStatusMixin `json:",inline"`
	ErrorMessage            string `json:"errorMessage,omitempty"`
}

func (s *BaremetalClusterProviderStatus) GetClusterStatusMixin() *kaas.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// The naming conventions for bridge interfaces to the network types
// that are defined for managed Kubernetes clusters on bare metal.
const (
	// Connects LCM agents running on the hosts to the Container Cloud LCM API.
	// The LCM API is provided by the regional or management cluster.
	// +gocode:public-api=true
	KubernetesLCMBridgeName = "k8s-lcm"
	// Kubernetes Workloads (pods) network serves internal connections
	// between K8s pods running on different nodes. Kubernetes pods subnet
	// provides IP addresses used by Calico to connect pods between nodes.
	// +gocode:public-api=true
	KubernetesPodsBridgeName = "k8s-pods"
)

// BaremetalClusterProviderSpec is the providerSpec for metal3 in the cluster object
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=project
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BaremetalClusterProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.ClusterSpecMixin `json:",inline"`

	// A token that is used for both validation of the practically of the API server from a
	// joining node's point of view and as an authentication method for the node in the
	// bootstrap phase of "kubeadm join"
	// TODO Temporary workaround. Should be removed eventually.
	// +optional
	BootstrapToken string `json:"bootstrapToken,omitempty" sensitive:"true"`

	// When true, dedicated MetalLB address pools will be used for services in management/regional cluster;
	// baremetal operator services will use dedicated MetalLB address pool associated with PXE network.
	DedicatedMetallbPools bool `json:"dedicatedMetallbPools"`

	// LB IP address to expose Kubernetes API and swarm API of the cluster.
	// If it's not set (empty), LB IP address will be obtained from dedicated
	// IPaddr object.
	LoadBalancerHost string `json:"loadBalancerHost" sensitive:"true"`

	// When true, the cluster is configured to use BGP to announce
	// external IPs for the k8s cluster services and for the k8s API LB
	// (i.e., "layer 3" mode),
	// Rack and MultiRackCluster objects will be used to configure that.
	// When false (by default), the cluster is configured to use ARP
	// to announce external IPs (i.e., "layer 2" mode),
	// Rack and MultiRackCluster objects will not be used.
	// +optional
	UseBGPAnnouncement bool `json:"useBGPAnnouncement,omitempty"`
}

func (s *BaremetalClusterProviderSpec) GetClusterSpecMixin() *kaas.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}
func (*BaremetalClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &BaremetalClusterProviderStatus{}
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&BaremetalClusterProviderSpec{})
	SchemeBuilder.Register(&BaremetalClusterProviderStatus{})
}
