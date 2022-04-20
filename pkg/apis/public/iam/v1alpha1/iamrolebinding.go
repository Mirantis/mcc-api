package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

type IAMNestedName struct {
	Name string `json:"name"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
type IAMGlobalRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserRoleReference `json:",inline"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IAMGlobalRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMGlobalRoleBinding `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource
type IAMClusterRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserRoleReference `json:",inline"`
	// Cluster is in the same namespace as IAMClusterRoleBinding
	Cluster IAMNestedName `json:"cluster"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource

type IAMClusterRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMClusterRoleBinding `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource
type IAMRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserRoleReference `json:",inline"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type IAMRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMRoleBinding `json:"items"`
}

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
