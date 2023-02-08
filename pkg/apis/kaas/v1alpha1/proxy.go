package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProxyList contains a list of Proxy
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Proxy `json:"items"`
}
type ProxySpec struct {
	// httpProxy will be used as the proxy URL for HTTP and HTTPS requests unless overridden by httpsProxy or noProxy.
	HTTPProxy string `json:"httpProxy,omitempty"`
	// httpProxyPasswordSecret will be populated automatically if the httpProxy URL contains a password.
	HTTPProxyPassword *SecretValue `json:"httpProxyPassword,omitempty"`
	// httpsProxywill be used as the proxy URL for HTTPS requests unless overridden by noProxy.
	HTTPSProxy string `json:"httpsProxy,omitempty"`
	// httpsProxyPasswordSecret will be populated automatically if the httpsProxy URL contains a password.
	HTTPSProxyPassword *SecretValue `json:"httpsProxyPassword,omitempty"`
	// noProxy specifies URLs that should be excluded from proxying as a comma-separated list of domain names.
	NoProxy string `json:"noProxy,omitempty"`
	// CACertificate is the custom trusted proxy CA certificate that is being used for MITM proxies.
	CACertificate []byte `json:"caCertificate,omitempty"`
}

// Proxy is the Schema for the proxy API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
// +gocode:public-api=true
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProxySpec `json:"spec"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}
