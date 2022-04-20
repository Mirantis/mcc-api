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

	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VsphereClusterProviderSpec is the schema for the vsphereclusterproviderspec API
// +k8s:openapi-gen=true
type VsphereClusterProviderSpec struct {
	metav1.TypeMeta               `json:",inline"`
	metav1.ObjectMeta             `json:"metadata,omitempty"`
	kaasv1alpha1.ClusterSpecMixin `json:",inline"`

	VsphereConfig  VsphereResourcesConfig `json:"vsphere"`
	ClusterNetwork ClusterNetwork         `json:"clusterNetwork"`

	LoadBalancerHost string `json:"loadBalancerHost"`
}

type ClusterNetwork struct {
	// ipamEnabled must be set to true if relevant vSphere network has no dhcp. Otherwise it could be false
	// if it is set to true ipam will handle IPs allocation for VMs instead of dhcp.
	IpamEnabled bool `json:"ipamEnabled"`

	// CIDR is the address of network in CIDR notation. To be be allocated to VMs in case when ipamEnabled is
	// set to true
	CIDR string `json:"cidr,omitempty"`

	// This pool range includes addresses that will be allocated to VMs in the current cluster.
	// The size of this range limits the number of hosts that can be deployed in the current cluster.
	// Must be in the CIDR range
	IncludeRanges []string `json:"includeRanges,omitempty"`

	// This pool range excludes addresses from beein allocated to VMs in current cluster
	// Must be in the CIDR range
	ExcludeRanges []string `json:"excludeRanges,omitempty"`

	// The default gateway on the relevant vSphere network
	Gateway string `json:"gateway,omitempty"`

	// An external DNS servers accessible from the relevant vSphere network.
	Nameservers []string `json:"nameservers,omitempty"`
}

type VsphereResourcesConfig struct {
	// DatastoreName is a relevant vSphere datastore name in specified datacenter
	CloudProviderDatastore string `json:"cloudProviderDatastore"`

	// DatastoreName is a relevant vSphere datastore name in specified datacenter
	ClusterAPIDatastore string `json:"clusterApiDatastore,omitempty"`

	// ClusterAPIDatastoreFolder is a absolute vSphere datastore cluster or datastore folder path
	// in specified datacenter
	ClusterAPIDatastoreFolder string `json:"clusterApiDatastoreFolder,omitempty"`

	// UseLocalDatastores defines whether to use local or shared datastores. If true datastore will be selected
	// automatically from ClusterAPIDatastoreFolder
	UseLocalDatastores bool `json:"useLocalDatastores,omitempty"`

	// MachineFolderPath is a relevant vSphere machine folder full path, where machines shall reside
	MachineFolderPath string `json:"machineFolderPath"`

	// NetworkPath is a relevant vSphere network full path in specified datacenter
	NetworkPath string `json:"networkPath"`

	// ResourcePoolPath is a relevant vSphere resource pool full path in specified datacenter
	ResourcePoolPath string `json:"resourcePoolPath"`

	// SCSIControllerType defines SCSI controller to be used
	SCSIControllerType string `json:"scsiControllerType,omitempty"`
}

func (s *VsphereClusterProviderSpec) GetClusterSpecMixin() *kaasv1alpha1.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}

func (*VsphereClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &VsphereClusterProviderStatus{}
}

// KeyPair is how operators can supply custom keypairs for kubeadm to use.
type KeyPair struct {
	// base64 encoded cert and key
	Cert []byte `json:"cert"`
	Key  []byte `json:"key"`
}

// HasCertAndKey returns whether a keypair contains cert and key of non-zero length.
func (kp KeyPair) HasCertAndKey() bool {
	return len(kp.Cert) > 0 && len(kp.Key) > 0
}

func init() {
	SchemeBuilder.Register(&VsphereClusterProviderSpec{})
}
