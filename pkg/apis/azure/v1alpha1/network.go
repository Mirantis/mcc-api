package v1alpha1

import (
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
)

const (
	// Internal is the value for the Azure load balancer internal type.
	// +gocode:public-api=true
	Internal = LBType("Internal")
	// Public is the value for the Azure load balancer public type.
	// +gocode:public-api=true
	Public = LBType("Public")
)

// PublicIPSpec defines the inputs to create an Azure public IP address.
type PublicIPSpec struct {
	Name string `json:"name"`
	// +optional
	DNSName string `json:"dnsName,omitempty" sensitive:"true"`
}

// VnetSpec configures an Azure virtual network.
type VnetSpec struct {
	// ResourceGroup is the name of the resource group of the existing virtual network
	// or the resource group where a managed virtual network should be created.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// ID is the identifier of the virtual network this provider should use to create resources.
	ID string `json:"id,omitempty" sensitive:"true"`

	// Name defines a name for the virtual network resource.
	Name string `json:"name" sensitive:"true"`

	// CIDRBlocks defines the virtual network's address space, specified as one or more address prefixes in CIDR notation.
	// +optional
	CIDRBlocks []string `json:"cidrBlocks,omitempty" sensitive:"true"`

	// Tags is a collection of tags describing the resource.
	// +optional
	Tags Tags `json:"tags,omitempty"`
}

// SecurityGroup defines an Azure security group.
// +gocode:public-api=true
type SecurityGroup struct {
	ID            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	SecurityRules SecurityRules     `json:"securityRules,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

// AddressRecord specifies a DNS record mapping a hostname to an IPV4 or IPv6 address.
// +gocode:public-api=true
type AddressRecord struct {
	Hostname string
	IP       string
}

// NetworkSpec specifies what the Azure networking resources should look like.
// +gocode:public-api=true
type NetworkSpec struct {
	// Vnet is the configuration for the Azure virtual network.
	// +optional
	Vnet VnetSpec `json:"vnet,omitempty"`

	// Subnets is the configuration for the control-plane subnet and the node subnet.
	// +optional
	Subnets Subnets `json:"subnets,omitempty"`

	// APIServerLB is the configuration for the control-plane load balancer.
	// +optional
	APIServerLB LoadBalancerSpec `json:"apiServerLB,omitempty"`

	// NodeOutboundLB is the configuration for the node outbound load balancer.
	// +optional
	NodeOutboundLB *LoadBalancerSpec `json:"nodeOutboundLB,omitempty"`

	// PublicServicesLB is the configuration for the load balancer of kubernetes public services.
	// +optional
	PublicServicesLB *LoadBalancerSpec `json:"publicServicesLB,omitempty"`

	// PrivateDNSZoneName defines the zone name for the Azure Private DNS.
	// +optional
	PrivateDNSZoneName string `json:"privateDNSZoneName,omitempty"`
}

// GetNodeSubnet returns the cluster nodes subnet.
func (n *NetworkSpec) GetNodeSubnet() (SubnetSpec, error) {
	for _, sn := range n.Subnets {
		if sn.Role == SubnetNode {
			return sn, nil
		}
	}
	return SubnetSpec{}, errors.Errorf("no subnet found with role %s", SubnetNode)
}

// UpdateNodeSubnet updates the cluster node subnet.
func (n *NetworkSpec) UpdateNodeSubnet(subnet SubnetSpec) {
	for i, sn := range n.Subnets {
		if sn.Role == SubnetNode {
			n.Subnets[i] = subnet
		}
	}
}

// GetBastionSubnet returns the Bastion subnet.
func (n *NetworkSpec) GetBastionSubnet() (SubnetSpec, error) {
	for _, sn := range n.Subnets {
		if sn.Role == SubnetBastion {
			return sn, nil
		}
	}
	return SubnetSpec{}, errors.Errorf("no subnet found with role %s", SubnetBastion)
}

// UpdateBastionSubnet updates the Bastion subnet.
func (n *NetworkSpec) UpdateBastionSubnet(subnet SubnetSpec) {
	for i, sn := range n.Subnets {
		if sn.Role == SubnetBastion {
			n.Subnets[i] = subnet
		}
	}
}

// RouteTable defines an Azure route table.
// +gocode:public-api=true
type RouteTable struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// SecurityRule defines an Azure security rule for security groups.
// +gocode:public-api=true
type SecurityRule struct {
	// Name is a unique name within the network security group.
	Name string `json:"name"`
	// A description for this rule. Restricted to 140 chars.
	Description string `json:"description"`
	// Protocol specifies the protocol type. "Tcp", "Udp", "Icmp", or "*".
	// +kubebuilder:validation:Enum=Tcp;Udp;Icmp;*
	Protocol SecurityGroupProtocol `json:"protocol"`
	// Direction indicates whether the rule applies to inbound, or outbound traffic. "Inbound" or "Outbound".
	// +kubebuilder:validation:Enum=Inbound;Outbound
	Direction SecurityRuleDirection `json:"direction"`
	// Priority is a number between 100 and 4096. Each rule should have a unique value for priority. Rules are processed in priority order, with lower numbers processed before higher numbers. Once traffic matches a rule, processing stops.
	Priority int32 `json:"priority,omitempty"`
	// SourcePorts specifies source port or range. Integer or range between 0 and 65535. Asterix '*' can also be used to match all ports.
	SourcePorts *string `json:"sourcePorts,omitempty"`
	// DestinationPorts specifies the destination port or range. Integer or range between 0 and 65535. Asterix '*' can also be used to match all ports.
	DestinationPorts *string `json:"destinationPorts,omitempty"`
	// Source specifies the CIDR or source IP range. Asterix '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	Source *string `json:"source,omitempty"`
	// Destination is the destination address prefix. CIDR or destination IP range. Asterix '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	Destination *string `json:"destination,omitempty"`
}

var (
	// SecurityGroupNode defines a Kubernetes workload node role
	// +gocode:public-api=true
	SecurityGroupNode = SecurityGroupRole("node")

	// SecurityGroupControlPlane defines a Kubernetes control plane node role
	// +gocode:public-api=true
	SecurityGroupControlPlane = SecurityGroupRole("controlplane")
)

// LoadBalancerSpec defines an Azure load balancer.
type LoadBalancerSpec struct {
	ID          string       `json:"id,omitempty" sensitive:"true"`
	Name        string       `json:"name,omitempty" sensitive:"true"`
	SKU         SKU          `json:"sku,omitempty"`
	FrontendIPs []FrontendIP `json:"frontendIPs,omitempty"`
	Type        LBType       `json:"type,omitempty"`
	// FrontendIPsCount specifies the number of frontend IP addresses for the load balancer.
	FrontendIPsCount *int32 `json:"frontendIPsCount,omitempty"`
	// IdleTimeoutInMinutes specifies the timeout for the TCP idle connection.
	IdleTimeoutInMinutes *int32 `json:"idleTimeoutInMinutes,omitempty"`
}

// Subnets is a slice of Subnet.
// +gocode:public-api=true
type Subnets []SubnetSpec

// +gocode:public-api=true
type SubnetSpec struct {
	// Role defines the subnet role (eg. Node, ControlPlane)
	Role SubnetRole `json:"role,omitempty"`

	// ID defines a unique identifier to reference this resource.
	// +optional
	ID string `json:"id,omitempty" sensitive:"true"`

	// Name defines a name for the subnet resource.
	Name string `json:"name"`

	// CIDRBlocks defines the subnet's address space, specified as one or more address prefixes in CIDR notation.
	// +optional
	CIDRBlocks []string `json:"cidrBlocks,omitempty" sensitive:"true"`

	// SecurityGroup defines the NSG (network security group) that should be attached to this subnet.
	// +optional
	SecurityGroup SecurityGroup `json:"securityGroup,omitempty"`

	// RouteTable defines the route table that should be attached to this subnet.
	// +optional
	RouteTable RouteTable `json:"routeTable,omitempty"`
}

// SecurityRules is a slice of Azure security rules for security groups.
// +gocode:public-api=true
type SecurityRules []SecurityRule

// SKU defines an Azure load balancer SKU.
// +gocode:public-api=true
type SKU string

const (
	// SKUStandard is the value for the Azure load balancer Standard SKU.
	// +gocode:public-api=true
	SKUStandard = SKU("Standard")
)

// FrontendIP defines a load balancer frontend IP configuration.
// +gocode:public-api=true
type FrontendIP struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// +optional
	PrivateIPAddress string `json:"privateIP,omitempty" sensitive:"true"`
	// +optional
	PublicIP *PublicIPSpec `json:"publicIP,omitempty" sensitive:"true"`
}

// LBType defines an Azure load balancer Type.
// +gocode:public-api=true
type LBType string

// SecurityGroupRole defines the unique role of a security group.
// +gocode:public-api=true
type SecurityGroupRole string
