package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GracefulRebootRequestSpec struct {
	// Machines is a list of machines that need to be rebooted.
	// An empty list means that all machines in the cluster need to be rebooted.
	Machines []string `json:"machines,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type GracefulRebootRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GracefulRebootRequest `json:"items"`
}

const (
	// +gocode:public-api=true
	LCMMachineGracefulRebootAnnotation = "lcm.mirantis.com/graceful-reboot"
	// +gocode:public-api=true
	LCMMachineGracefulRebootAnnotationValue = "true"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GracefulRebootRequest is the Schema for the proxy API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type GracefulRebootRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GracefulRebootRequestSpec `json:"spec"`
}

func (g *GracefulRebootRequest) Contains(machineName string) bool {
	if len(g.Spec.Machines) == 0 {
		return true
	}
	for _, machine := range g.Spec.Machines {
		if machine == machineName {
			return true
		}
	}
	return false
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&GracefulRebootRequest{}, &GracefulRebootRequestList{})
}
