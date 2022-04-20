package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PublicKey represents the ssh public key.
// +kubebuilder:resource
type PublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PublicKeySpec `json:"spec"`
}

// PublicKeySpec defines key in use
type PublicKeySpec struct {
	PublicKey string `json:"publicKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PublicKeyList contains a list of PublicKey
type PublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []PublicKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PublicKey{}, &PublicKeyList{})
}
