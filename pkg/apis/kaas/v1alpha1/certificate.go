package v1alpha1

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TLSConfigList contains a list of TLSConfig
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type TLSConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TLSConfig `json:"items"`
}

// TLSConfig is the Schema for the TLS API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
type TLSConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TLSConfigSpec `json:"spec"`
}

func (tls *TLSConfig) ExtractData(ctx context.Context, kubeClient client.Client) ([]byte, []byte, error) {
	keyBytes, err := tls.Spec.PrivateKey.Extract(ctx, kubeClient, tls.Namespace)
	if err != nil {
		return nil, nil, err
	}
	return tls.Spec.ServerCertificate, keyBytes, nil
}

// CertificateList contains a list of Certificate
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Certificate `json:"items"`
}
type TLSConfigSpec struct {
	// ServerName is a hostname of a server.
	ServerName string `json:"serverName"`
	// ServerCertificate is used to authenticate server's identity to a client.
	// A valid certificate bundle can be passed. The server certificate must be on the top of the chain.
	ServerCertificate []byte `json:"serverCertificate"`
	// PrivateKey is a key for a server. The key must correspond to the public key used in the server certificate.
	PrivateKey *SecretValue `json:"privateKey"`
	// CACertificate is the certificate that issued the server certificate.
	// If a CA certificate is unavailable, the top-most intermediate certificate should be used instead.
	CACertificate []byte `json:"caCertificate,omitempty"`
}

// Certificate is the Schema for the certificate API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
// +gocode:public-api=true
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

func (c *Certificate) ExtractData(ctx context.Context, kubeClient client.Client) ([]byte, []byte, error) {
	keyBytes, err := c.Spec.PrivateKey.Extract(ctx, kubeClient, c.Namespace)
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

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{}, &TLSConfig{}, &TLSConfigList{})
}
