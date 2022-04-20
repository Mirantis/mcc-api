package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Certificate is the Schema for the certificate API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

func (c *Certificate) ExtractData(kubeClient client.Client) ([]byte, []byte, error) {
	keyBytes, err := c.Spec.PrivateKey.Extract(kubeClient, c.Namespace)
	if err != nil {
		return nil, nil, err
	}
	return c.Spec.Certificate, keyBytes, nil
}

type CertificateSpec struct {
	Certificate []byte       `json:"certificate"`
	PrivateKey  *SecretValue `json:"privateKey"`
}

type CertificateStatus struct {
	ExpirationDateVerified *bool `json:"expirationDateVerified"`
	PrivateKeyVerified     *bool `json:"privateKeyVerified"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateList contains a list of Certificate
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Certificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
