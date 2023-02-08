package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PublicKeySpec defines key in use
type PublicKeySpec struct {
	PublicKey string `json:"publicKey"`
}

// PublicKeyList contains a list of PublicKey
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type PublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []PublicKey `json:"items"`
}

// PublicKey represents the ssh public key.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource
// +gocode:public-api=true
type PublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PublicKeySpec `json:"spec"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&PublicKey{}, &PublicKeyList{})
}
