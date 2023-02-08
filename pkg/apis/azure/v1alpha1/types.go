package v1alpha1

// SubnetRole defines the unique role of a subnet.
// +gocode:public-api=true
type SubnetRole string

const (
	// SecurityRuleDirectionInbound defines an ingress security rule.
	// +gocode:public-api=true
	SecurityRuleDirectionInbound = SecurityRuleDirection("Inbound")

	// SecurityRuleDirectionOutbound defines an egress security rule.
	// +gocode:public-api=true
	SecurityRuleDirectionOutbound = SecurityRuleDirection("Outbound")
)
const (
	// InvalidConfigurationMachineError represents that the combination
	// of configuration in the MachineSpec is not supported by this cluster.
	// This is not a transient error, but
	// indicates a state that must be fixed before progress can be made.
	//
	// Example: the ProviderSpec specifies an instance type that doesn't exist.
	// +gocode:public-api=true
	InvalidConfigurationMachineError MachineStatusError = "InvalidConfiguration"

	// UnsupportedChangeMachineError indicates that the MachineSpec has been updated in a way that
	// is not supported for reconciliation on this cluster. The spec may be
	// completely valid from a configuration standpoint, but the controller
	// does not support changing the real world state to match the new
	// spec.
	//
	// Example: the responsible controller is not capable of changing the
	// container runtime from docker to rkt.
	// +gocode:public-api=true
	UnsupportedChangeMachineError MachineStatusError = "UnsupportedChange"

	// InsufficientResourcesMachineError generally refers to exceeding one's quota in a cloud provider,
	// or running out of physical machines in an on-premise environment.
	// +gocode:public-api=true
	InsufficientResourcesMachineError MachineStatusError = "InsufficientResources"

	// CreateMachineError indicates an error while trying to create a Node to match this
	// Machine. This may indicate a transient problem that will be fixed
	// automatically with time, such as a service outage, or a terminal
	// error during creation that doesn't match a more specific
	// MachineStatusError value.
	//
	// Example: timeout trying to connect to GCE.
	// +gocode:public-api=true
	CreateMachineError MachineStatusError = "CreateError"

	// UpdateMachineError indicates an error while trying to update a Node that this
	// Machine represents. This may indicate a transient problem that will be
	// fixed automatically with time, such as a service outage,
	//
	// Example: error updating load balancers.
	// +gocode:public-api=true
	UpdateMachineError MachineStatusError = "UpdateError"

	// DeleteMachineError indicates an error was encountered while trying to delete the Node that this
	// Machine represents. This could be a transient or terminal error, but
	// will only be observable if the provider's Machine controller has
	// added a finalizer to the object to more gracefully handle deletions.
	//
	// Example: cannot resolve EC2 IP address.
	// +gocode:public-api=true
	DeleteMachineError MachineStatusError = "DeleteError"

	// JoinClusterTimeoutMachineError indicates that the machine did not join the cluster
	// as a new node within the expected timeframe after instance
	// creation at the provider succeeded
	//
	// Example use case: A controller that deletes Machines which do
	// not result in a Node joining the cluster within a given timeout
	// and that are managed by a MachineSet.
	// +gocode:public-api=true
	JoinClusterTimeoutMachineError = "JoinClusterTimeoutError"
)
const (
	// ControlPlane machine label
	// +gocode:public-api=true
	ControlPlane string = "control-plane"
	// Node machine label
	// +gocode:public-api=true
	Node string = "node"
	// Bastion label
	// +gocode:public-api=true
	Bastion string = "bastion"
)
const (
	// SubnetNode defines a Kubernetes workload node role
	// +gocode:public-api=true
	SubnetNode = SubnetRole(Node)

	// SubnetBastion defines a Bastion node role
	// +gocode:public-api=true
	SubnetBastion = SubnetRole(Bastion)
)
const (
	// SecurityGroupProtocolAll is a wildcard for all IP protocols.
	// +gocode:public-api=true
	SecurityGroupProtocolAll = SecurityGroupProtocol("*")
	// SecurityGroupProtocolTCP represents the TCP protocol.
	// +gocode:public-api=true
	SecurityGroupProtocolTCP = SecurityGroupProtocol("Tcp")
	// SecurityGroupProtocolUDP represents the UDP protocol.
	// +gocode:public-api=true
	SecurityGroupProtocolUDP = SecurityGroupProtocol("Udp")
	// SecurityGroupProtocolICMP represents the ICMP protocol.
	// +gocode:public-api=true
	SecurityGroupProtocolICMP = SecurityGroupProtocol("Icmp")
)

// SecurityRuleDirection defines the direction type for a security group rule.
// +gocode:public-api=true
type SecurityRuleDirection string

// MachineStatusError defines errors states for Machine objects.
// +gocode:public-api=true
type MachineStatusError string

// SecurityGroupProtocol defines the protocol type for a security group rule.
// +gocode:public-api=true
type SecurityGroupProtocol string
