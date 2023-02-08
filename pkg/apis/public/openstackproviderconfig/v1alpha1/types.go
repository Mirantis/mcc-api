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

	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+resource:path=project

// OpenstackMachineProviderSpec is the type that will be embedded in a Machine.Spec.ProviderSpec field
// for an OpenStack Instance. It is used by the Openstack machine actuator to create a single machine instance.
// +k8s:openapi-gen=true
type OpenstackMachineProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.MachineSpecMixin `json:",inline"`

	// The flavor reference for the flavor for your server instance.
	Flavor string `json:"flavor"`
	// The name of the image to use for your server instance.
	Image string `json:"image"`

	// A networks object. Required parameter when there are multiple networks defined for the tenant.
	// When you do not specify the networks parameter, the server attaches to the only network created for the current tenant.
	Networks []NetworkParam `json:"networks,omitempty"`
	// The floatingIP which will be associated to the machine, only used for master.
	// The floatingIP should have been created and haven't been associated.
	FloatingIP string `json:"floatingIP,omitempty"`

	// The availability zone from which to launch the server.
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// The names of the security groups to assign to the instance
	SecurityGroups []string `json:"securityGroups,omitempty"`

	// The list of additional volumes being created.
	AdditionalVolumes []VolumeSpec `json:"additionalVolumes,omitempty"`

	// Metadata mapping. Allows you to create a map of key value pairs to add to the server instance.
	ServerMetadata map[string]string `json:"serverMetadata,omitempty"`
}

func (s *OpenstackMachineProviderSpec) GetMachineSpecMixin() *kaas.MachineSpecMixin {
	return &s.MachineSpecMixin
}

func (*OpenstackMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &OpenstackMachineProviderStatus{}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OpenstackMachineProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	kaas.MachineStatusMixin `json:",inline"`

	AdditionalVolumes *[]VolumeStatus `json:"additionalVolumes,omitempty"`
}

func (s *OpenstackMachineProviderStatus) GetMachineStatusMixin() *kaas.MachineStatusMixin {
	return &s.MachineStatusMixin
}

type NetworkParam struct {
	// The UUID of the network. Required if you omit the port attribute.
	UUID string `json:"uuid,omitempty"`
	// A fixed IPv4 address for the NIC.
	FixedIP string `json:"fixed_ip,omitempty"`
	// Filters for optional network query
	Filter Filter `json:"filter,omitempty"`
}

type Filter struct {
	Status       string `json:"status,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	AdminStateUp *bool  `json:"admin_state_up,omitempty"`
	TenantID     string `json:"tenant_id,omitempty"`
	ProjectID    string `json:"project_id,omitempty"`
	Shared       *bool  `json:"shared,omitempty"`
	ID           string `json:"id,omitempty"`
	Marker       string `json:"marker,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	SortKey      string `json:"sort_key,omitempty"`
	SortDir      string `json:"sort_dir,omitempty"`
	Tags         string `json:"tags,omitempty"`
	TagsAny      string `json:"tags-any,omitempty"`
	NotTags      string `json:"not-tags,omitempty"`
	NotTagsAny   string `json:"not-tags-any,omitempty"`
}

type RootVolume struct {
	VolumeType string `json:"volumeType"`
	Size       int    `json:"diskSize,omitempty"`
}

type BastionSpec struct {
	Flavor           string            `json:"flavor,omitempty"`
	Image            string            `json:"image,omitempty"`
	AvailabilityZone string            `json:"availabilityZone,omitempty"`
	RedeployAllowed  bool              `json:"redeployAllowed,omitempty"`
	ServerMetadata   map[string]string `json:"serverMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackClusterProviderSpec is the providerSpec for OpenStack in the cluster object
// +k8s:openapi-gen=true
type OpenstackClusterProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.ClusterSpecMixin `json:",inline"`

	// NodeCIDR is the OpenStack Subnet to be created. Cluster actuator will create a
	// network, a subnet with NodeCIDR, and a router connected to this subnet.
	// If you leave this empty, no network will be created.
	NodeCIDR string `json:"nodeCidr,omitempty"`
	// DNSNameservers is the list of nameservers for OpenStack Subnet being created.
	DNSNameservers []string `json:"dnsNameservers,omitempty"`
	// ExternalNetworkID is the ID of an external OpenStack Network. This is necessary
	// to get public internet to the VMs.
	ExternalNetworkID string `json:"externalNetworkId,omitempty"`
	// Bastion host configuration
	// +optional
	Bastion *BastionSpec `json:"bastion,omitempty"`
}

func (s *OpenstackClusterProviderSpec) GetClusterSpecMixin() *kaas.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}

func (*OpenstackClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &OpenstackClusterProviderStatus{}
}

type BastionStatus struct {
	PublicIP string `json:"publicIP,omitempty"`
	// LCM Agent is already installed
	LCMManaged bool `json:"lcmManaged,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackClusterProviderStatus contains the status fields
// relevant to OpenStack in the cluster object.
// +k8s:openapi-gen=true
type OpenstackClusterProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.ClusterStatusMixin `json:",inline"`

	// Network contains all information about the created OpenStack Network.
	// It includes Subnets and Router.
	Network *Network `json:"network,omitempty"`

	// ControlPlaneSecurityGroups contains all the information about the OpenStack
	// Security Group that needs to be applied to control plane nodes.
	// TODO: Maybe instead of two properties, we add a property to the group?
	ControlPlaneSecurityGroup *SecurityGroup `json:"controlPlaneSecurityGroup,omitempty"`

	// GlobalSecurityGroup contains all the information about the OpenStack Security
	// Group that needs to be applied to all nodes, both control plane and worker nodes.
	GlobalSecurityGroup *SecurityGroup `json:"globalSecurityGroup,omitempty"`
	// BastionSecurityGroup contains all the information about the OpenStack Security
	// Group that needs to be applied to bastion node
	BastionSecurityGroup *SecurityGroup `json:"bastionSecurityGroup,omitempty"`
	//Bastion contains all information about bastion node
	Bastion BastionStatus `json:"bastion,omitempty"`
}

func (s *OpenstackClusterProviderStatus) GetClusterStatusMixin() *kaas.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// Network represents basic information about the associated OpenStach Neutron Network
type Network struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	Subnet       *Subnet       `json:"subnet,omitempty"`
	Router       *Router       `json:"router,omitempty"`
	LoadBalancer *LoadBalancer `json:"loadbalancer,omitempty"`
}

// Subnet represents basic information about the associated OpenStack Neutron Subnet
type Subnet struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	CIDR string `json:"cidr"`
}

// Router represents basic information about the associated OpenStack Neutron Router
type Router struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Listener represents basic information about associated OpenStack Listener
type Listener struct {
	ID   string `json:"id"`
	Port int    `json:"port"`
}

// Pool represents basic information about associated OpenStack Pool
type Pool struct {
	ID string `json:"id"`
}

// LoadBalancer represents basic information about the associated OpenStack LoadBalancer
type LoadBalancer struct {
	Name       string              `json:"name"`
	ID         string              `json:"id"`
	FloatingIP string              `json:"floatingIP"`
	Listeners  map[string]Listener `json:"listeners"`
	Pools      map[string]Pool     `json:"pools"`
}

// VolumeSpec represents basic information about the Openstack Volume to be created
type VolumeSpec struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

// VolumeStatus represents basic information about created Openstack Volume
type VolumeStatus struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func init() {
	SchemeBuilder.Register(&OpenstackMachineProviderSpec{})
	SchemeBuilder.Register(&OpenstackMachineProviderStatus{})
	SchemeBuilder.Register(&OpenstackClusterProviderSpec{})
	SchemeBuilder.Register(&OpenstackClusterProviderStatus{})
}
