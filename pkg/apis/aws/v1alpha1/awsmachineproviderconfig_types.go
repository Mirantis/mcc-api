package v1alpha1

import (
	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// AWSMachineProviderSpec is the type that will be embedded in a Machine.Spec.ProviderSpec field
// for an AWS instance. It is used by the AWS machine actuators to create a single machine instance,
// using the RunInstances call (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html)
// Required parameters such as region that are not specified by this configuration, will be defaulted
// by the actuators.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type AWSMachineProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.MachineSpecMixin `json:",inline"`

	// AMI is the reference to the AMI from which to create the machine instance.
	AMI AWSResourceReference `json:"ami"`

	// ImageLookupOrg is the AWS Organization ID to use for image lookup if AMI is not set.
	ImageLookupOrg string `json:"imageLookupOrg,omitempty"`

	// InstanceType is the type of instance to create. Example: m4.xlarge
	InstanceType string `json:"instanceType,omitempty"`

	// AdditionalTags is the set of tags to add to an instance, in addition to the ones
	// added by default by the actuators. These tags are additive. The actuators will ensure
	// these tags are present, but will not remove any other tags that may exist on the
	// instance.
	// +optional
	AdditionalTags map[string]string `json:"additionalTags,omitempty"`

	// IAMInstanceProfile is a name of an IAM instance profile to assign to the instance
	// +optional
	IAMInstanceProfile string `json:"iamInstanceProfile,omitempty"`

	// PublicIP specifies whether the instance should get a public IP.
	// Precedence for this setting is as follows:
	// 1. This field if set
	// 2. Cluster/flavor setting
	// 3. Subnet default
	// +optional
	PublicIP *bool `json:"publicIP,omitempty" sensitive:"true"`

	// AdditionalSecurityGroups is an array of references to security groups that should be applied to the
	// instance. These security groups would be set in addition to any security groups defined
	// at the cluster level or in the actuators.
	// +optional
	AdditionalSecurityGroups []AWSResourceReference `json:"additionalSecurityGroups,omitempty"`

	// Subnet is a reference to the subnet to use for this instance. If not specified,
	// the cluster subnet will be used.
	// +optional
	Subnet *AWSResourceReference `json:"subnet,omitempty"`

	// RootDeviceSize is the size of the root volume.
	// +optional
	RootDeviceSize int64 `json:"rootDeviceSize,omitempty"`

	// RootDeviceType specifies volume type of the root storage device
	// +optional
	// +kubebuilder:validation:Enum=standard;io1;io2;gp2;sc1;st1;gp3
	RootDeviceType string `json:"rootDeviceType,omitempty"`
}

func (*AWSMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &AWSMachineProviderStatus{}
}
func (s *AWSMachineProviderSpec) GetMachineSpecMixin() *kaas.MachineSpecMixin {
	return &s.MachineSpecMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AWSMachineProviderSpec{})
}
