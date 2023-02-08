package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ServiceUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceUser `json:"items"`
}
type ServiceUserStatus struct {
	Created bool `json:"created"`
}

// ServiceUser represents an admin user created in KeyCloak
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
type ServiceUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceUserSpec   `json:"spec"`
	Status ServiceUserStatus `json:"status,omitempty"`
}
type ServiceUserSpec struct {
	Password *SecretValue `json:"password,omitempty"`
}
