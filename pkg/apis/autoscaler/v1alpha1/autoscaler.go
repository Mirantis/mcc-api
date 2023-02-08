package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AutoScaler a component that automatically adjusts the size of a Cluster so that all pods have a place to run and
// there are no unneeded nodes.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=as
// +kubebuilder:subresource:status
// +gocode:public-api=true
type AutoScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoScalerSpec   `json:"spec,omitempty"`
	Status AutoScalerStatus `json:"status,omitempty"`
}

// AutoScalerSpec defines the desired state of AutoScaler
type AutoScalerSpec struct {
	Cluster       string         `json:"cluster"`
	ScalingGroups []ScalingGroup `json:"scalingGroups"`
}
type ScalingGroup struct {
	MachinePool string `json:"machinePool"`
	MinSize     int    `json:"minSize,omitempty"`
	MaxSize     int    `json:"maxSize,omitempty"`
}

// AutoScalerStatus defines the observed state of AutoScaler
type AutoScalerStatus struct{}

// AutoScalerList contains a list of AutoScaler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type AutoScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoScaler `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AutoScaler{}, &AutoScalerList{})
}
