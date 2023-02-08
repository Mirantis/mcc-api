package v1alpha1

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"time"
)

// RouteTable defines an AWS routing table.
// +gocode:public-api=true
type RouteTable struct {
	ID string `json:"id"`
}

// SecurityGroupProtocol defines the protocol type for a security group rule.
// +gocode:public-api=true
type SecurityGroupProtocol string

var (
	// ClassicELBSchemeInternetFacing defines an internet-facing, publicly
	// accessible AWS Classic ELB scheme
	// +gocode:public-api=true
	ClassicELBSchemeInternetFacing = ClassicELBScheme("Internet-facing")

	// ClassicELBSchemeInternal defines an internal-only facing
	// load balancer internal to an ELB.
	// +gocode:public-api=true
	ClassicELBSchemeInternal = ClassicELBScheme("internal")
)
var (
	// ClassicELBProtocolTCP defines the ELB API string representing the TCP protocol
	// +gocode:public-api=true
	ClassicELBProtocolTCP = ClassicELBProtocol("TCP")

	// ClassicELBProtocolSSL defines the ELB API string representing the TLS protocol
	// +gocode:public-api=true
	ClassicELBProtocolSSL = ClassicELBProtocol("SSL")

	// ClassicELBProtocolHTTP defines the ELB API string representing the HTTP protocol at L7
	// +gocode:public-api=true
	ClassicELBProtocolHTTP = ClassicELBProtocol("HTTP")

	// ClassicELBProtocolHTTPS defines the ELB API string representing the HTTP protocol at L7
	// +gocode:public-api=true
	ClassicELBProtocolHTTPS = ClassicELBProtocol("HTTPS")
)

// Subnets is a slice of Subnet.
// +gocode:public-api=true
type Subnets []*SubnetSpec

// ToMap returns a map from id to subnet.
func (s Subnets) ToMap() map[string]*SubnetSpec {
	res := make(map[string]*SubnetSpec)
	for _, x := range s {
		res[x.ID] = x
	}
	return res
}

// FindByID returns a single subnet matching the given id or nil.
func (s Subnets) FindByID(id string) *SubnetSpec {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}

	return nil
}

// FilterPrivate returns a slice containing all subnets marked as private.
func (s Subnets) FilterPrivate() Subnets {
	var res Subnets

	for _, x := range s {
		if !x.IsPublic {
			res = append(res, x)
		}
	}

	return res
}

// FilterPublic returns a slice containing all subnets marked as public.
func (s Subnets) FilterPublic() Subnets {
	var res Subnets

	for _, x := range s {
		if x.IsPublic {
			res = append(res, x)
		}
	}

	return res
}

// SecurityGroupRole defines the unique role of a security group.
// +gocode:public-api=true
type SecurityGroupRole string

var (
	// InstanceStatePending is the string representing an instance in a pending state
	// +gocode:public-api=true
	InstanceStatePending = InstanceState("pending")

	// InstanceStateRunning is the string representing an instance in a pending state
	// +gocode:public-api=true
	InstanceStateRunning = InstanceState("running")

	// InstanceStateShuttingDown is the string representing an instance shutting down
	// +gocode:public-api=true
	InstanceStateShuttingDown = InstanceState("shutting-down")

	// InstanceStateTerminated is the string representing an instance that has been terminated
	// +gocode:public-api=true
	InstanceStateTerminated = InstanceState("terminated")

	// InstanceStateStopping is the string representing an instance
	// that is in the process of being stopped and can be restarted
	// +gocode:public-api=true
	InstanceStateStopping = InstanceState("stopping")

	// InstanceStateStopped is the string representing an instance
	// that has been stopped and can be restarted
	// +gocode:public-api=true
	InstanceStateStopped = InstanceState("stopped")
)

// AWSResourceReference is a reference to a specific AWS resource by ID, ARN, or filters.
// Only one of ID, ARN or Filters may be specified. Specifying more than one will result in
// a validation error.
// +gocode:public-api=true
type AWSResourceReference struct {
	// ID of resource
	ID *string `json:"id" sensitive:"true"`
}

// NetworkELB defines an AWS network load balancer.
// +gocode:public-api=true
type NetworkELB struct {
	// The name of the load balancer. It must be unique within the set of load balancers
	// defined in the region. It also serves as identifier.
	Name string `json:"name,omitempty"`

	// ARN uniquely identify AWS resources.
	Arn string `json:"arn,omitempty"`

	// DNSName is the dns name of the load balancer.
	DNSName string `json:"dnsName,omitempty"`

	// Scheme is the load balancer scheme, either internet-facing or private.
	Scheme NetworkELBScheme `json:"scheme,omitempty"`

	// SubnetIDs is an array of subnets in the VPC attached to the load balancer.
	SubnetIDs []string `json:"subnetIds,omitempty"`

	// Listeners is an array of classic elb listeners associated with the load balancer. There must be at least one.
	Listeners []*NetworkELBListener `json:"listeners,omitempty"`

	// TargetGroups is an array of network elb target groups to route requests to one or more registered targets.
	TargetGroups []*NetworkELBTargetGroup `json:"targetGroups,omitempty"`

	// Attributes defines extra attributes associated with the load balancer.
	Attributes []*NetworkELBAttribute `json:"attributes,omitempty"`

	// Tags is a map of tags associated with the load balancer.
	Tags map[string]string `json:"tags,omitempty"`

	// State of the load balancer.
	State string `json:"state,omitempty"`
}

var (
	// TargetTypeInstance defines that target is specified by instance ID
	// +gocode:public-api=true
	TargetTypeInstance = TargetType("instance")

	// TargetTypeIP defines that target is specified by IP address
	// +gocode:public-api=true
	TargetTypeIP = TargetType("ip")
)
var (
	// SecurityGroupBastion defines an SSH bastion role
	// +gocode:public-api=true
	SecurityGroupBastion = SecurityGroupRole("bastion")

	// SecurityGroupNode defines a Kubernetes workload node role
	// +gocode:public-api=true
	SecurityGroupNode = SecurityGroupRole("node")

	// SecurityGroupControlPlane defines a Kubernetes control plane node role
	// +gocode:public-api=true
	SecurityGroupControlPlane = SecurityGroupRole("controlplane")
)
var (
	// SecurityGroupProtocolAll is a wildcard for all IP protocols
	// +gocode:public-api=true
	SecurityGroupProtocolAll = SecurityGroupProtocol("-1")

	// SecurityGroupProtocolIPinIP represents the IP in IP protocol in ingress rules
	// +gocode:public-api=true
	SecurityGroupProtocolIPinIP = SecurityGroupProtocol("4")

	// SecurityGroupProtocolTCP represents the TCP protocol in ingress rules
	// +gocode:public-api=true
	SecurityGroupProtocolTCP = SecurityGroupProtocol("tcp")

	// SecurityGroupProtocolUDP represents the UDP protocol in ingress rules
	// +gocode:public-api=true
	SecurityGroupProtocolUDP = SecurityGroupProtocol("udp")

	// SecurityGroupProtocolICMP represents the ICMP protocol in ingress rules
	// +gocode:public-api=true
	SecurityGroupProtocolICMP = SecurityGroupProtocol("icmp")

	// SecurityGroupProtocolICMPv6 represents the ICMPv6 protocol in ingress rules
	// +gocode:public-api=true
	SecurityGroupProtocolICMPv6 = SecurityGroupProtocol("58")
)

// AWSMachineProviderCondition is a condition in a AWSMachineProviderStatus
// +gocode:public-api=true
type AWSMachineProviderCondition struct {
	// Type is the type of the condition.
	Type AWSMachineProviderConditionType `json:"type"`
	// Status is the status of the condition.
	Status corev1.ConditionStatus `json:"status"`
	// LastProbeTime is the last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime"`
	// LastTransitionTime is the last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// Reason is a unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason"`
	// Message is a human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message"`
}

// Network encapsulates AWS networking resources.
// +gocode:public-api=true
type Network struct {
	// SecurityGroups is a map from the role/kind of the security group to its unique name, if any.
	SecurityGroups map[SecurityGroupRole]*SecurityGroup `json:"securityGroups,omitempty"`

	// APIServerELB is the Kubernetes api server classic load balancer.
	APIServerELB ClassicELB `json:"apiServerElb,omitempty"`

	// APIServerNetworkELB is the Kubernetes api server network load balancer.
	APIServerNetworkELB NetworkELB `json:"apiServerNetworkElb,omitempty"`

	// Applied subnets configuration.
	// +optional
	Subnets Subnets `json:"subnets,omitempty"`
}

// ClassicELBAttributes defines extra attributes associated with a classic load balancer.
type ClassicELBAttributes struct {
	// IdleTimeout is time that the connection is allowed to be idle (no data
	// has been sent over the connection) before it is closed by the load balancer.
	IdleTimeout time.Duration `json:"idleTimeout,omitempty"`
}

// NetworkELBAttributes defines extra attributes associated with a network load balancer.
type NetworkELBAttribute struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// NetworkELBScheme defines the scheme of a network load balancer.
// +gocode:public-api=true
type NetworkELBScheme string

var (
	// NetworkELBSchemeInternetFacing defines an internet-facing, publicly
	// accessible AWS Classic ELB scheme
	// +gocode:public-api=true
	NetworkELBSchemeInternetFacing = NetworkELBScheme("internet-facing")

	// NetworkELBSchemeInternal defines an internal-only facing
	// load balancer internal to an ELB.
	// +gocode:public-api=true
	NetworkELBSchemeInternal = NetworkELBScheme("internal")
)

// NetworkELBTarget defines a target registered with the specified target group.
// +gocode:public-api=true
type NetworkELBTarget struct {
	AvailabilityZone string `json:"availabilityZone"`
	ID               string `json:"id"`
	Port             int64  `json:"port"`
}

// SecurityGroup defines an AWS security group.
// +gocode:public-api=true
type SecurityGroup struct {
	// ID is a unique identifier.
	ID string `json:"id"`

	// Name is the security group name.
	Name string `json:"name"`

	// IngressRules is the inbound rules associated with the security group.
	IngressRules IngressRules `json:"ingressRule"`

	// Tags is a map of tags associated with the security group.
	Tags map[string]string `json:"tags,omitempty"`
}

// String returns a string representation of the security group.
func (s *SecurityGroup) String() string {
	return fmt.Sprintf("id=%s/name=%s", s.ID, s.Name)
}

// ClassicELBScheme defines the scheme of a classic load balancer.
// +gocode:public-api=true
type ClassicELBScheme string

var (
	// NetworkELBProtocolTCP defines the ELB API string representing the TCP protocol
	// +gocode:public-api=true
	NetworkELBProtocolTCP = NetworkELBProtocol("TCP")

	// NetworkELBProtocolUDP defines the ELB API string representing the UDP protocol
	// +gocode:public-api=true
	NetworkELBProtocolUDP = NetworkELBProtocol("UDP")

	// NetworkELBProtocolTLS defines the ELB API string representing the TLS protocol
	// +gocode:public-api=true
	NetworkELBProtocolTLS = NetworkELBProtocol("TLS")
)

// NetworkELBListener defines an AWS network load balancer listener.
type NetworkELBListener struct {
	Protocol NetworkELBProtocol `json:"protocol"`
	Port     int64              `json:"port"`
}

// NetworkELBTargetGroup defines a target group that used to route requests to one or more registered targets.
type NetworkELBTargetGroup struct {
	Arn                        string             `json:"arn"`
	Name                       string             `json:"name"`
	HealthCheckEnabled         bool               `json:"healthCheckEnabled"`
	HealthCheckIntervalSeconds int64              `json:"healthCheckIntervalSeconds"`
	HealthCheckPort            string             `json:"healthCheckPort"`
	HealthCheckProtocol        NetworkELBProtocol `json:"healthCheckProtocol"`
	Protocol                   NetworkELBProtocol `json:"protocol"`
	Port                       int64              `json:"port"`
	HealthyThreshold           int64              `json:"healthyThreshold"`
	UnhealthyThreshold         int64              `json:"unhealthyThreshold"`
	TargetType                 TargetType         `json:"targetType"`
	VpcID                      string             `json:"vcpId"`
}

// InstanceState describes the state of an AWS instance.
// +gocode:public-api=true
type InstanceState string

// Instance describes an AWS instance.
// +gocode:public-api=true
type Instance struct {
	ID string `json:"id"`

	// The current state of the instance.
	State InstanceState `json:"instanceState,omitempty"`

	// The instance type.
	Type string `json:"type,omitempty"`

	// The ID of the subnet of the instance.
	SubnetID string `json:"subnetId,omitempty"`

	// The ID of the AMI used to launch the instance.
	ImageID string `json:"imageId,omitempty"`

	// SecurityGroupIDs are one or more security group IDs this instance belongs to.
	SecurityGroupIDs []string `json:"securityGroupIds,omitempty"`

	// UserData is the raw data script passed to the instance which is run upon bootstrap.
	// This field must not be base64 encoded and should only be used when running a new instance.
	UserData *string `json:"userData,omitempty"`

	// The name of the IAM instance profile associated with the instance, if applicable.
	IAMProfile string `json:"iamProfile,omitempty"`

	// The private IPv4 address assigned to the instance.
	PrivateIP string `json:"privateIp,omitempty"`

	// The public IPv4 address assigned to the instance, if applicable.
	PublicIP string `json:"publicIp,omitempty"`

	// Specifies whether enhanced networking with ENA is enabled.
	ENASupport *bool `json:"enaSupport,omitempty"`

	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EBSOptimized *bool `json:"ebsOptimized,omitempty"`

	// Specifies size (in Gi) of the root storage device
	RootDeviceSize int64 `json:"rootDeviceSize,omitempty"`

	// RootDeviceType specifies volume type of the root storage device
	// +kubebuilder:validation:Enum=standard;io1;io2;gp2;sc1;st1;gp3
	RootDeviceType string `json:"rootDeviceType,omitempty"`

	// The tags associated with the instance.
	Tags map[string]string `json:"tags,omitempty"`

	// The availability zone of the instance.
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// LCM Agent is already installed
	LCMManaged bool `json:"lcmManaged,omitempty"`
}

// TargetType defines what targets are specified by.
// +gocode:public-api=true
type TargetType string

// IngressRule defines an AWS ingress rule for security groups.
// +gocode:public-api=true
type IngressRule struct {
	Description string                `json:"description"`
	Protocol    SecurityGroupProtocol `json:"protocol"`
	FromPort    int64                 `json:"fromPort"`
	ToPort      int64                 `json:"toPort"`

	// List of CIDR blocks to allow access from. Cannot be specified with SourceSecurityGroupID.
	CidrBlocks []string `json:"cidrBlocks"`

	// The security group id to allow access from. Cannot be specified with CidrBlocks.
	SourceSecurityGroupIDs []string `json:"sourceSecurityGroupIds"`
}

// String returns a string representation of the ingress rule.
func (i *IngressRule) String() string {
	return fmt.Sprintf("protocol=%s/range=[%d-%d]/description=%s", i.Protocol, i.FromPort, i.ToPort, i.Description)
}

// Equals returns true if two IngressRule are equal
func (i *IngressRule) Equals(o *IngressRule) bool {
	if len(i.CidrBlocks) != len(o.CidrBlocks) {
		return false
	}

	sort.Strings(i.CidrBlocks)
	sort.Strings(o.CidrBlocks)

	for i, v := range i.CidrBlocks {
		if v != o.CidrBlocks[i] {
			return false
		}
	}

	if len(i.SourceSecurityGroupIDs) != len(o.SourceSecurityGroupIDs) {
		return false
	}

	sort.Strings(i.SourceSecurityGroupIDs)
	sort.Strings(o.SourceSecurityGroupIDs)

	for i, v := range i.SourceSecurityGroupIDs {
		if v != o.SourceSecurityGroupIDs[i] {
			return false
		}
	}

	return i.Description == o.Description &&
		i.FromPort == o.FromPort &&
		i.ToPort == o.ToPort &&
		i.Protocol == o.Protocol
}

// IngressRules is a slice of AWS ingress rules for security groups.
type IngressRules []*IngressRule

// Difference returns the difference between this slice and the other slice.
func (i IngressRules) Difference(o IngressRules) IngressRules {
	var out IngressRules

	for _, x := range i {
		found := false
		for _, y := range o {
			if x.Equals(y) {
				found = true
				break
			}
		}

		if !found {
			out = append(out, x)
		}
	}

	return out
}

// Valid conditions for an AWS machine instance
const (
	// MachineCreated indicates whether the machine has been created or not. If not,
	// it should include a reason and message for the failure.
	// +gocode:public-api=true
	MachineCreated AWSMachineProviderConditionType = "MachineCreated"
)

// ClassicELB defines an AWS classic load balancer.
// +gocode:public-api=true
type ClassicELB struct {
	// The name of the load balancer. It must be unique within the set of load balancers
	// defined in the region. It also serves as identifier.
	Name string `json:"name,omitempty"`

	// DNSName is the dns name of the load balancer.
	DNSName string `json:"dnsName,omitempty"`

	// Scheme is the load balancer scheme, either internet-facing or private.
	Scheme ClassicELBScheme `json:"scheme,omitempty"`

	// SubnetIDs is an array of subnets in the VPC attached to the load balancer.
	SubnetIDs []string `json:"subnetIds,omitempty"`

	// SecurityGroupIDs is an array of security groups assigned to the load balancer.
	SecurityGroupIDs []string `json:"securityGroupIds,omitempty"`

	// Listeners is an array of classic elb listeners associated with the load balancer. There must be at least one.
	Listeners []*ClassicELBListener `json:"listeners,omitempty"`

	// HealthCheck is the classic elb health check associated with the load balancer.
	HealthCheck *ClassicELBHealthCheck `json:"healthChecks,omitempty"`

	// Attributes defines extra attributes associated with the load balancer.
	Attributes ClassicELBAttributes `json:"attributes,omitempty"`

	// Tags is a map of tags associated with the load balancer.
	Tags map[string]string `json:"tags,omitempty"`
}

// ClassicELBHealthCheck defines an AWS classic load balancer health check.
type ClassicELBHealthCheck struct {
	Target             string        `json:"target"`
	Interval           time.Duration `json:"interval"`
	Timeout            time.Duration `json:"timeout"`
	HealthyThreshold   int64         `json:"healthyThreshold"`
	UnhealthyThreshold int64         `json:"unhealthyThreshold"`
}

// NetworkELBProtocol defines listener protocols for a network load balancer.
// +gocode:public-api=true
type NetworkELBProtocol string

// AWSMachineProviderConditionType is a valid value for AWSMachineProviderCondition.Type
// +gocode:public-api=true
type AWSMachineProviderConditionType string

// ClassicELBProtocol defines listener protocols for a classic load balancer.
// +gocode:public-api=true
type ClassicELBProtocol string

// ClassicELBListener defines an AWS classic load balancer listener.
type ClassicELBListener struct {
	Protocol         ClassicELBProtocol `json:"protocol"`
	Port             int64              `json:"port"`
	InstanceProtocol ClassicELBProtocol `json:"instanceProtocol"`
	InstancePort     int64              `json:"instancePort"`
}
