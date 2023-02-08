package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSSet is the Schema for the DNSSet API
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type DNSSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DNSSpec `json:"spec"`
}
type DNSSpec struct {
	Endpoints []Endpoint `json:"endpoints"`
}
type Endpoint struct {
	// DNSNames groups hostnames of the DNS records.
	DNSNames []string `json:"dnsNames"`
	// Targets specifies the address DNS records point to.
	Targets []string `json:"targets"`
}

// DNSSetList contains a list of DNSSet
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type DNSSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DNSSet `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&DNSSet{}, &DNSSetList{})
}
