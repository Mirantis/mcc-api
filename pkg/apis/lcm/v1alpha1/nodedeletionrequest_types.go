package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type NodeDeletionRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeDeletionRequest `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=nodedeletionrequests,scope=Cluster
type NodeDeletionRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeDeletionRequestSpec `json:"spec,omitempty"`
}
type NodeDeletionRequestSpec struct {
	NodeName string `json:"nodeName"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&NodeDeletionRequest{}, &NodeDeletionRequestList{})
}
