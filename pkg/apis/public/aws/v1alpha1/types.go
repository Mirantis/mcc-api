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
	"fmt"
	"sort"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSResourceReference is a reference to a specific AWS resource by ID, ARN, or filters.
// Only one of ID, ARN or Filters may be specified. Specifying more than one will result in
// a validation error.
type AWSResourceReference struct {
	// ID of resource
	ID *string `json:"id"`
}

// AWSMachineProviderConditionType is a valid value for AWSMachineProviderCondition.Type
type AWSMachineProviderConditionType string

// Valid conditions for an AWS machine instance
const (
	// MachineCreated indicates whether the machine has been created or not. If not,
	// it should include a reason and message for the failure.
	MachineCreated AWSMachineProviderConditionType = "MachineCreated"
)

// AWSMachineProviderCondition is a condition in a AWSMachineProviderStatus
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
type Network struct {
	// SecurityGroups is a map from the role/kind of the security group to its unique name, if any.
	SecurityGroups map[SecurityGroupRole]*SecurityGroup `json:"securityGroups,omitempty"`

	// APIServerELB is the Kubernetes api server classic load balancer.
	APIServerELB ClassicELB `json:"apiServerElb,omitempty"`

	// APIServerNetworkELB is the Kubernetes api server network load balancer.
	APIServerNetworkELB NetworkELB `json:"apiServerNetworkElb,omitempty"`
}

// ClassicELBScheme defines the scheme of a classic load balancer.
type ClassicELBScheme string

var (
	// ClassicELBSchemeInternetFacing defines an internet-facing, publicly
	// accessible AWS Classic ELB scheme
	ClassicELBSchemeInternetFacing = ClassicELBScheme("Internet-facing")

	// ClassicELBSchemeInternal defines an internal-only facing
	// load balancer internal to an ELB.
	ClassicELBSchemeInternal = ClassicELBScheme("internal")
)

// ClassicELBProtocol defines listener protocols for a classic load balancer.
type ClassicELBProtocol string

var (
	// ClassicELBProtocolTCP defines the ELB API string representing the TCP protocol
	ClassicELBProtocolTCP = ClassicELBProtocol("TCP")

	// ClassicELBProtocolSSL defines the ELB API string representing the TLS protocol
	ClassicELBProtocolSSL = ClassicELBProtocol("SSL")

	// ClassicELBProtocolHTTP defines the ELB API string representing the HTTP protocol at L7
	ClassicELBProtocolHTTP = ClassicELBProtocol("HTTP")

	// ClassicELBProtocolHTTPS defines the ELB API string representing the HTTP protocol at L7
	ClassicELBProtocolHTTPS = ClassicELBProtocol("HTTPS")
)

// ClassicELB defines an AWS classic load balancer.
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

// ClassicELBAttributes defines extra attributes associated with a classic load balancer.
type ClassicELBAttributes struct {
	// IdleTimeout is time that the connection is allowed to be idle (no data
	// has been sent over the connection) before it is closed by the load balancer.
	IdleTimeout time.Duration `json:"idleTimeout,omitempty"`
}

// ClassicELBListener defines an AWS classic load balancer listener.
type ClassicELBListener struct {
	Protocol         ClassicELBProtocol `json:"protocol"`
	Port             int64              `json:"port"`
	InstanceProtocol ClassicELBProtocol `json:"instanceProtocol"`
	InstancePort     int64              `json:"instancePort"`
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
type NetworkELBProtocol string

var (
	// NetworkELBProtocolTCP defines the ELB API string representing the TCP protocol
	NetworkELBProtocolTCP = NetworkELBProtocol("TCP")

	// NetworkELBProtocolUDP defines the ELB API string representing the UDP protocol
	NetworkELBProtocolUDP = NetworkELBProtocol("UDP")

	// NetworkELBProtocolTLS defines the ELB API string representing the TLS protocol
	NetworkELBProtocolTLS = NetworkELBProtocol("TLS")
)

// NetworkELBScheme defines the scheme of a network load balancer.
type NetworkELBScheme string

var (
	// NetworkELBSchemeInternetFacing defines an internet-facing, publicly
	// accessible AWS Classic ELB scheme
	NetworkELBSchemeInternetFacing = NetworkELBScheme("internet-facing")

	// NetworkELBSchemeInternal defines an internal-only facing
	// load balancer internal to an ELB.
	NetworkELBSchemeInternal = NetworkELBScheme("internal")
)

// NetworkELB defines an AWS network load balancer.
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

	//TargetGroups is an array of network elb target groups to route requests to one or more registered targets.
	TargetGroups []*NetworkELBTargetGroup `json:"targetGroups,omitempty"`

	// Attributes defines extra attributes associated with the load balancer.
	Attributes []*NetworkELBAttribute `json:"attributes,omitempty"`

	// Tags is a map of tags associated with the load balancer.
	Tags map[string]string `json:"tags,omitempty"`

	// State of the load balancer.
	State string `json:"state,omitempty"`
}

// NetworkELBAttributes defines extra attributes associated with a network load balancer.
type NetworkELBAttribute struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// NetworkELBListener defines an AWS network load balancer listener.
type NetworkELBListener struct {
	Protocol NetworkELBProtocol `json:"protocol"`
	Port     int64              `json:"port"`
}

//TargetType defines what targets are specified by.
type TargetType string

var (
	//TargetTypeInstance defines that target is specified by instance ID
	TargetTypeInstance = TargetType("instance")

	//TargetTypeIP defines that target is specified by IP address
	TargetTypeIP = TargetType("ip")
)

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

// NetworkELBTarget defines a target registered with the specified target group.
type NetworkELBTarget struct {
	AvailabilityZone string `json:"availabilityZone"`
	ID               string `json:"id"`
	Port             int64  `json:"port"`
}

// Subnets is a slice of Subnet.
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

// RouteTable defines an AWS routing table.
type RouteTable struct {
	ID string `json:"id"`
}

// SecurityGroupRole defines the unique role of a security group.
type SecurityGroupRole string

var (
	// SecurityGroupBastion defines an SSH bastion role
	SecurityGroupBastion = SecurityGroupRole("bastion")

	// SecurityGroupNode defines a Kubernetes workload node role
	SecurityGroupNode = SecurityGroupRole("node")

	// SecurityGroupControlPlane defines a Kubernetes control plane node role
	SecurityGroupControlPlane = SecurityGroupRole("controlplane")
)

// SecurityGroup defines an AWS security group.
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

// SecurityGroupProtocol defines the protocol type for a security group rule.
type SecurityGroupProtocol string

var (
	// SecurityGroupProtocolAll is a wildcard for all IP protocols
	SecurityGroupProtocolAll = SecurityGroupProtocol("-1")

	// SecurityGroupProtocolIPinIP represents the IP in IP protocol in ingress rules
	SecurityGroupProtocolIPinIP = SecurityGroupProtocol("4")

	// SecurityGroupProtocolTCP represents the TCP protocol in ingress rules
	SecurityGroupProtocolTCP = SecurityGroupProtocol("tcp")

	// SecurityGroupProtocolUDP represents the UDP protocol in ingress rules
	SecurityGroupProtocolUDP = SecurityGroupProtocol("udp")

	// SecurityGroupProtocolICMP represents the ICMP protocol in ingress rules
	SecurityGroupProtocolICMP = SecurityGroupProtocol("icmp")

	// SecurityGroupProtocolICMPv6 represents the ICMPv6 protocol in ingress rules
	SecurityGroupProtocolICMPv6 = SecurityGroupProtocol("58")
)

// IngressRule defines an AWS ingress rule for security groups.
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

// InstanceState describes the state of an AWS instance.
type InstanceState string

var (
	// InstanceStatePending is the string representing an instance in a pending state
	InstanceStatePending = InstanceState("pending")

	// InstanceStateRunning is the string representing an instance in a pending state
	InstanceStateRunning = InstanceState("running")

	// InstanceStateShuttingDown is the string representing an instance shutting down
	InstanceStateShuttingDown = InstanceState("shutting-down")

	// InstanceStateTerminated is the string representing an instance that has been terminated
	InstanceStateTerminated = InstanceState("terminated")

	// InstanceStateStopping is the string representing an instance
	// that is in the process of being stopped and can be restarted
	InstanceStateStopping = InstanceState("stopping")

	// InstanceStateStopped is the string representing an instance
	// that has been stopped and can be restarted
	InstanceStateStopped = InstanceState("stopped")
)

// Instance describes an AWS instance.
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

	// LCM Agent is already installed
	LCMManaged bool `json:"lcmManaged,omitempty"`
}
