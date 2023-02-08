package v1alpha1

import (
	machinev1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// +gocode:public-api=true
	MachinePoolDeletePolicyNever MachinePoolDeletePolicy = "never"
	// +gocode:public-api=true
	MachinePoolDeletePolicyRandom MachinePoolDeletePolicy = "random"
	// +gocode:public-api=true
	MachinePoolDeletePolicyOldest MachinePoolDeletePolicy = "oldest"
	// +gocode:public-api=true
	MachinePoolDeletePolicyNewest MachinePoolDeletePolicy = "newest"
)

// MachinePoolStatus defines the observed state of MachinePool
type MachinePoolStatus struct {
	// Replicas is the most recently observed number of replicas.
	Replicas int `json:"replicas,omitempty"`

	// The number of ready replicas for this MachinePool.
	ReadyReplicas int `json:"readyReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed MachinePool.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// MachinePool ensures that a specified number of machines replicas are running at any given time.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=mp
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.spec.labelSelector
// +gocode:public-api=true
type MachinePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachinePoolSpec   `json:"spec,omitempty"`
	Status MachinePoolStatus `json:"status,omitempty"`
}

// +gocode:public-api=true
type MachinePoolDeletePolicy string

// MachinePoolList contains a list of MachinePool
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type MachinePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachinePool `json:"items"`
}

// MachinePoolSpec defines the desired state of MachinePool
type MachinePoolSpec struct {
	// Replicas is the number of desired replicas.
	// Defaults to 0.
	// +optional
	Replicas int `json:"replicas,omitempty"`

	// DeletePolicy defines the policy used to identify nodes to delete when downscaling.
	// Defaults to "never". For now the only valid value is "never".
	// +kubebuilder:default:=never
	// +kubebuilder:validation:Enum=never
	DeletePolicy MachinePoolDeletePolicy `json:"deletePolicy,omitempty"`

	// MachineSpec is the object that describes the machine that will be created if
	// insufficient replicas are detected.
	// +optional
	MachineSpec machinev1.MachineSpec `json:"machineSpec,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&MachinePool{}, &MachinePoolList{})
}
