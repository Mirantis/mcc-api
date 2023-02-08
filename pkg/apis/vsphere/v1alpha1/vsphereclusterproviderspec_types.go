package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterNetwork struct {
	// ipamEnabled must be set to true if relevant vSphere network has no dhcp. Otherwise it could be false
	// if it is set to true ipam will handle IPs allocation for VMs instead of dhcp.
	IpamEnabled bool `json:"ipamEnabled"`

	// CIDR is the address of network in CIDR notation. To be be allocated to VMs in case when ipamEnabled is
	// set to true
	CIDR string `json:"cidr,omitempty" sensitive:"true"`

	// This pool range includes addresses that will be allocated to VMs in the current cluster.
	// The size of this range limits the number of hosts that can be deployed in the current cluster.
	// Must be in the CIDR range
	IncludeRanges []string `json:"includeRanges,omitempty" sensitive:"true"`

	// This pool range excludes addresses from beein allocated to VMs in current cluster
	// Must be in the CIDR range
	ExcludeRanges []string `json:"excludeRanges,omitempty" sensitive:"true"`

	// The default gateway on the relevant vSphere network
	Gateway string `json:"gateway,omitempty" sensitive:"true"`

	// An external DNS servers accessible from the relevant vSphere network.
	Nameservers []string `json:"nameservers,omitempty" sensitive:"true"`
}
type VsphereResourcesConfig struct {
	// DatastoreName is a relevant vSphere datastore name in specified datacenter
	CloudProviderDatastore string `json:"cloudProviderDatastore" sensitive:"true"`

	// DatastoreName is a relevant vSphere datastore name in specified datacenter
	ClusterAPIDatastore string `json:"clusterApiDatastore,omitempty" sensitive:"true"`

	// ClusterAPIDatastoreFolder is a absolute vSphere datastore cluster or datastore folder path
	// in specified datacenter
	ClusterAPIDatastoreFolder string `json:"clusterApiDatastoreFolder,omitempty" sensitive:"true"`

	// UseLocalDatastores defines whether to use local or shared datastores. If true datastore will be selected
	// automatically from ClusterAPIDatastoreFolder
	UseLocalDatastores bool `json:"useLocalDatastores,omitempty"`

	// MachineFolderPath is a relevant vSphere machine folder full path, where machines shall reside
	MachineFolderPath string `json:"machineFolderPath" sensitive:"true"`

	// NetworkPath is a relevant vSphere network full path in specified datacenter
	NetworkPath string `json:"networkPath" sensitive:"true"`

	// ResourcePoolPath is a relevant vSphere resource pool full path in specified datacenter
	ResourcePoolPath string `json:"resourcePoolPath" sensitive:"true"`

	// SCSIControllerType defines SCSI controller to be used
	SCSIControllerType string `json:"scsiControllerType,omitempty" sensitive:"true"`
}

// KeyPair is how operators can supply custom keypairs for kubeadm to use.
// +gocode:public-api=true
type KeyPair struct {
	// base64 encoded cert and key
	Cert []byte `json:"cert"`
	Key  []byte `json:"key"`
}

// HasCertAndKey returns whether a keypair contains cert and key of non-zero length.
func (kp KeyPair) HasCertAndKey() bool {
	return len(kp.Cert) > 0 && len(kp.Key) > 0
}

// VsphereClusterProviderSpec is the schema for the vsphereclusterproviderspec API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type VsphereClusterProviderSpec struct {
	metav1.TypeMeta               `json:",inline"`
	metav1.ObjectMeta             `json:"metadata,omitempty"`
	kaasv1alpha1.ClusterSpecMixin `json:",inline"`

	VsphereConfig  VsphereResourcesConfig `json:"vsphere"`
	ClusterNetwork ClusterNetwork         `json:"clusterNetwork"`

	LoadBalancerHost string `json:"loadBalancerHost" sensitive:"true"`
}

func (s *VsphereClusterProviderSpec) GetClusterSpecMixin() *kaasv1alpha1.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}
func (*VsphereClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &VsphereClusterProviderStatus{}
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&VsphereClusterProviderSpec{})
}
