package v1alpha1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

// +gocode:public-api=true
type ClusterStateKey struct {
	StateType LCMClusterStateType
	Arg       string
	Cluster   string
}

func (key ClusterStateKey) String() string {
	if key.Arg == "" {
		return string(key.StateType)
	}
	return fmt.Sprintf("%s/%s", key.StateType, key.Arg)
}

// GetClusterStateName Prepare name regarding to rfc1123
// Whole name size is up to 63 characters
// cordon-drain -> cd-<machine-name> /  swarm-drain: sd-<machine-name>
// *without cluster name, due to a machine-name already has cluster name as prefix
func (key ClusterStateKey) GetClusterStateName() string {
	prefix := ""
	suffix := ""
	switch key.StateType {
	case LCMClusterStateTypeCordonDrain:
		prefix = "cd-"
		suffix = strings.ToLower(key.Arg)
	case LCMClusterStateTypeSwarmDrain:
		prefix = "sd-"
		suffix = strings.ToLower(key.Arg)
	case LCMClusterStateTypeSwarmRemove:
		prefix = "sr-"
		suffix = strings.ToLower(key.Arg)
	default:

		argStr := ""
		if key.Arg != "" {
			argStr = "-" + strings.ToLower(key.Arg)
		}
		return fmt.Sprintf("%s%s-%s-", key.StateType, argStr, key.Cluster)
	}

	resultStrLength := len(prefix) + len(suffix)
	if resultStrLength > NameLength {

		suffix = suffix[resultStrLength-NameLength:]
	}
	return prefix + suffix
}

const (
	// LCMClusterStateTypeHelmDeployed denotes the deployment of
	// helm-controller on the child cluster
	// +gocode:public-api=true
	LCMClusterStateTypeHelmDeployed LCMClusterStateType = "helm-deployed"
	// LCMClusterStateTypeCordonDrain denotes cordon&drain of a node,
	// if it's registered with child apiserver. If it's not registered,
	// the value is just set to "true" immediately.
	// +gocode:public-api=true
	LCMClusterStateTypeCordonDrain LCMClusterStateType = "cordon-drain"
	// LCMClusterStateTypeSwarmDrain denotes cordon&drain of a Swarm node
	// +gocode:public-api=true
	LCMClusterStateTypeSwarmDrain LCMClusterStateType = "swarm-drain"
	// LCMClusterStateTypeSwarmRemove denotes removal of a node from the Swarm
	// +gocode:public-api=true
	LCMClusterStateTypeSwarmRemove LCMClusterStateType = "swarm-remove"
	// LCMClusterStateValueTrue denotes a true value for LCMClusterState
	// +gocode:public-api=true
	LCMClusterStateValueTrue = "true"
	// LCMClusterStateValueFalse denotes a false value for LCMClusterState
	// +gocode:public-api=true
	LCMClusterStateValueFalse = "false"
	// LCMClusterStateReasonAnnotation denotes the reason the LCMClusterState
	// was created
	// +gocode:public-api=true
	LCMClusterStateReasonAnnotation = "lcm.mirantis.com/cluster-state-reason"
	// LCMCLusterStateReasonMaintenance denotes maintenance cause of
	// LCMClusterState creation
	// +gocode:public-api=true
	LCMCLusterStateReasonMaintenance = "maintenance"
	// LCMClusterStateReasonUpgrade denotes upgrade cause of LCMClusterState
	// creation
	// +gocode:public-api=true
	LCMClusterStateReasonUpgrade = "upgrade"
	// LCMClusterStateReasonDelete denotes deletion Node cause of LCMClusterState
	// creation
	// +gocode:public-api=true
	LCMClusterStateReasonDelete = "delete"
	// NameLength denotes the maximum length of LCMClusterState object name
	// +gocode:public-api=true
	NameLength = int(63)
)

// LCMCLusterStateType denotes a type of the child cluster state
// +gocode:public-api=true
type LCMClusterStateType string

// LCMClusterStateList contains a list of LCMClusterState objects
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMClusterStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LCMClusterState `json:"items"`
}

// LCMClusterState is the Schema for the lcmclusterstate API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="clusterName",type="string",JSONPath=".spec.clusterName",description="Cluster Name",priority=0
// +kubebuilder:printcolumn:name="type",type="string",JSONPath=".spec.type",description="Type",priority=0
// +kubebuilder:printcolumn:name="arg",type="string",JSONPath=".spec.arg",description="Arg",priority=1
// +kubebuilder:printcolumn:name="value",type="string",JSONPath=".spec.value",description="Arg",priority=1
// +kubebuilder:printcolumn:name="actualValue",type="string",JSONPath=".status.value",description="Actual Value",priority=2
// +kubebuilder:printcolumn:name="attempt",type="integer",JSONPath=".status.attempt",description="Attempt",priority=2
// +kubebuilder:printcolumn:name="message",type="string",JSONPath=".status.message",description="Message",priority=2
// +gocode:public-api=true
type LCMClusterState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LCMClusterStateSpec   `json:"spec,omitempty"`
	Status LCMClusterStateStatus `json:"status,omitempty"`
}

// LCMClusterStateSpec defines a desired state of a child cluster
type LCMClusterStateSpec struct {
	// The name of the LCMCluster object.
	ClusterName string `json:"clusterName,omitempty"`
	// Type specifies the type of this child cluster state object.
	Type LCMClusterStateType `json:"type"`
	// Arg is an argument specific to this type of the child cluster state,
	// for instance, a node name. The Arg should not change once the
	// object is created.
	Arg string `json:"arg"`
	// Value is the target value for this child cluster state,
	// describing, for example, whether helm controller should
	// be deployed ("true") or not ("false").
	Value string `json:"value"`
}

// LCMClusterStateStatus defines the desired state of a child cluster
type LCMClusterStateStatus struct {
	// Value describes the currently applied value for this LCMClusterState
	Value string `json:"value,omitempty"`
	// Message is a message describing an error, if any
	Message string `json:"message,omitempty"`
	// Attempt number
	Attempt int `json:"attempt"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&LCMClusterState{}, &LCMClusterStateList{})
}
