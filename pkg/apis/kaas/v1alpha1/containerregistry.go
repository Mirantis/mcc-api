package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerRegistryList contains a list of ContainerRegistry
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ContainerRegistry `json:"items"`
}

// ContainerRegistry is the Schema for the proxy API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
// +gocode:public-api=true
type ContainerRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ContainerRegistrySpec `json:"spec"`
}
type ContainerRegistrySpec struct {
	// Domain name of the container registry
	Domain string `json:"domain"`
	// CACert is a base64 encoded CA certificate
	CACert []byte `json:"CACert"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&ContainerRegistry{}, &ContainerRegistryList{})
}
