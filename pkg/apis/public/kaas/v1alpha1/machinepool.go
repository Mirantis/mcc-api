package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	machinev1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/cluster/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachinePool ensures that a specified number of machines replicas are running at any given time.
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=mp
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.spec.labelSelector
type MachinePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachinePoolSpec   `json:"spec,omitempty"`
	Status MachinePoolStatus `json:"status,omitempty"`
}

type MachineDeletePolicy string

const (
	MachineDeletePolicyNever  MachineDeletePolicy = "never"
	MachineDeletePolicyRandom MachineDeletePolicy = "random"
	MachineDeletePolicyOldest MachineDeletePolicy = "oldest"
	MachineDeletePolicyNewest MachineDeletePolicy = "newest"
)

// MachinePoolSpec defines the desired state of MachinePool
type MachinePoolSpec struct {
	// Replicas is the number of desired replicas.
	// Defaults to 1.
	// +optional
	// +kubebuilder:default:=1
	Replicas int `json:"replicas,omitempty"`

	// DeletePolicy defines the policy used to identify nodes to delete when downscaling.
	// Defaults to "never". For now the only valid value is "never".
	// +kubebuilder:default:=never
	// +kubebuilder:validation:Enum=never
	DeletePolicy MachineDeletePolicy `json:"deletePolicy,omitempty"`

	// MachineSpec is the object that describes the machine that will be created if
	// insufficient replicas are detected.
	// +optional
	MachineSpec machinev1.MachineSpec `json:"machineSpec,omitempty"`
}

// MachinePoolStatus defines the observed state of MachinePool
type MachinePoolStatus struct {
	// Replicas is the most recently observed number of replicas.
	Replicas int `json:"replicas,omitempty"`

	// The number of ready replicas for this MachinePool.
	ReadyReplicas int `json:"readyReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed MachinePool.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachinePoolList contains a list of MachinePool
type MachinePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachinePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MachinePool{}, &MachinePoolList{})
}
