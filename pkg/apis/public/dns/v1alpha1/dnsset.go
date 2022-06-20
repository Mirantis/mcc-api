package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSSet is the Schema for the DNSSet API
// +k8s:openapi-gen=true
// +kubebuilder:resource
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSSetList contains a list of DNSSet
type DNSSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DNSSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSSet{}, &DNSSetList{})
}
