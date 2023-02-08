package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +gocode:public-api=true
type UserRoleReference struct {
	// +optional
	External bool `json:"external"`
	// +optional
	Legacy bool `json:"legacy"`
	// +optional
	LegacyRole string        `json:"legacyRole"`
	Role       IAMNestedName `json:"role"`
	User       IAMNestedName `json:"user"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource
// +gocode:public-api=true
type IAMClusterRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMClusterRoleBinding `json:"items"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type IAMGlobalRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMGlobalRoleBinding `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource
// +gocode:public-api=true
type IAMClusterRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserRoleReference `json:",inline"`
	// Cluster is in the same namespace as IAMClusterRoleBinding
	Cluster IAMNestedName `json:"cluster"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +gocode:public-api=true
type IAMGlobalRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserRoleReference `json:",inline"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type IAMRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMRoleBinding `json:"items"`
}
type IAMNestedName struct {
	Name string `json:"name"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource
// +gocode:public-api=true
type IAMRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserRoleReference `json:",inline"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(
		&IAMGlobalRoleBinding{},
		&IAMGlobalRoleBindingList{},
		&IAMClusterRoleBinding{},
		&IAMClusterRoleBindingList{},
		&IAMRoleBinding{},
		&IAMRoleBindingList{},
	)
}
