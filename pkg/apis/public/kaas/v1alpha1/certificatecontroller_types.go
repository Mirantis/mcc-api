package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateConfiguration describes an application to request a certificate for
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="application",type="string",JSONPath=".spec.application.name",description="Application Name",priority=0
// +kubebuilder:printcolumn:name="notBefore",type="string",JSONPath=".status.notBefore",description="NotBefore",priority=0
// +kubebuilder:printcolumn:name="notAfter",type="string",JSONPath=".status.notAfter",description="NotAfter",priority=0
// +kubebuilder:printcolumn:name="renewalTime",type="string",JSONPath=".status.renewalTime",description="RenewalTime",priority=0
// +kubebuilder:printcolumn:name="hostname",type="string",JSONPath=".status.hostname",description="Hostname",priority=0
type CertificateConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateConfigurationSpec   `json:"spec"`
	Status CertificateConfigurationStatus `json:"status,omitempty"`
}

type CertificateConfigurationSpec struct {
	// SecretName defines a secret to put a certificate and its private key in
	SecretName string `json:"secretName"`
	// ServiceName defines a kubernetes service to get an address of an application from
	// Must be mutually exclusive with Hostname
	ServiceName string `json:"serviceName,omitempty"`
	// Hostname defines a single host name to protect
	// Must be mutually exclusive with ServiceName
	Hostname string `json:"hostname,omitempty"`
	// Application field provides additional information about a managed application
	Application Application `json:"application"`
}

type Application struct {
	Name         string   `json:"name"`
	Deployments  []string `json:"deployments,omitempty"`
	Statefulsets []string `json:"statefulsets,omitempty"`
}

type CertificateConfigurationStatus struct {
	// Revision is a number that defines a generation of certificates
	Revision int `json:"revision,omitempty"`
	// CertificateHash is a hash of the last applied certificates
	CertificateHash string `json:"certificateHash,omitempty"`
	// Certificate contains a bundle of generated for application certificates
	Certificate string `json:"certificate,omitempty"`
	// NotBefore defines the start of the validity period for the certificate
	NotBefore metav1.Time `json:"notBefore,omitempty"`
	// NotAfter defines the end of the validity period for the certificate
	NotAfter metav1.Time `json:"notAfter,omitempty"`
	// RenewalTime defines the timing controller attempts to request a new certificate
	RenewalTime metav1.Time `json:"renewalTime,omitempty"`
	// Hostname defines the server name protected by the certificate
	Hostname string `json:"hostname,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateConfigurationList contains a list of CertificateConfiguration
type CertificateConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []CertificateConfiguration `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCCCertificateRequest contains certificate signing request
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type MCCCertificateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MCCCertificateRequestSpec   `json:"spec"`
	Status MCCCertificateRequestStatus `json:"status,omitempty"`
}

type MCCCertificateRequestSpec struct {
	// Application refers to an application request is created for
	Application string `json:"application"`
	// Revision is a number that defines a generation of certificates
	Revision int `json:"revision"`
	// Request contains PEM format CSR
	Request []byte `json:"request"`
}

type MCCCertificateRequestStatus struct {
	// Certificate contains a bundle of generated certificates
	Certificate []byte `json:"certificate,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCCCertificateRequestList contains a list of MCCCertificateRequest
type MCCCertificateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MCCCertificateRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateConfiguration{}, &CertificateConfigurationList{},
		&MCCCertificateRequest{}, &MCCCertificateRequestList{})
}
