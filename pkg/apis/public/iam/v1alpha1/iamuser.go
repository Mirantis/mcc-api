package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMUser represents the synced user from KeyCloak
// +kubebuilder:resource:scope=Cluster
type IAMUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ExternalID string `json:"externalID"`
	// +optional
	DisplayName string `json:"displayName,omitempty"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IAMUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&IAMUser{},
		&IAMUserList{},
	)
}
