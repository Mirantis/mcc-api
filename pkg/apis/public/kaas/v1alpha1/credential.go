package v1alpha1

import (
	"context"

	v1core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Mirantis/mcc-api/pkg/errors"
)

type SecretValue struct {
	Value  *string         `json:"value,omitempty"`
	Secret *SecretValueRef `json:"secret,omitempty"`
}

type SecretValueRef struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

func (sv SecretValue) Extract(kubeClient client.Client, namespace string) ([]byte, error) {
	if sv.Secret == nil {
		return nil, errors.Errorf("secret is not referenced")
	}
	secret := &v1core.Secret{}
	err := kubeClient.Get(context.Background(), client.ObjectKey{
		Name:      sv.Secret.Name,
		Namespace: namespace,
	}, secret)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get the referenced secret %v/%v",
			namespace, sv.Secret.Name)
	}
	secretValue, ok := secret.Data[sv.Secret.Key]
	if !ok {
		return nil, errors.Errorf("secret %v/%v doesn't have a value under the %v key",
			namespace, sv.Secret.Name, sv.Secret.Key)
	}
	return secretValue, nil
}

type CredentialStatusMixin struct {
	// A boolean flag indicating if Credential is valid or not
	Valid *bool `json:"valid"`
	// A message describing an error, if any
	Message string `json:"message,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenStackCredential is the Schema for the openstackcredentials API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type OpenStackCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenStackCredentialSpec   `json:"spec"`
	Status OpenStackCredentialStatus `json:"status,omitempty"`
}

type OpenStackCredentialSpec struct {
	AuthInfo   *OpenStackAuthInfo `json:"auth,omitempty"`
	AuthType   string             `json:"autType,omitempty"`
	RegionName string             `json:"regionName,omitempty"`

	// CACert is a base64 encoded CA Cert bundle that can be used as part of
	// verifying SSL API requests.
	CACert []byte `json:"CACert,omitempty"`
}

type OpenStackCredentialStatus struct {
	CredentialStatusMixin `json:",inline"`
}

type OpenStackAuthInfo struct {
	// AuthURL is the keystone/identity endpoint URL.
	AuthURL string `json:"authURL,omitempty"`

	// Username is the username of the user.
	Username string `json:"userName,omitempty"`

	// UserID is the unique ID of a user.
	UserID string `json:"userID,omitempty"`

	// Password is the password of the user.
	Password *SecretValue `json:"password,omitempty"`

	// ProjectName is the common/human-readable name of a project.
	// Users can be scoped to a project.
	// ProjectName on its own is not enough to ensure a unique scope. It must
	// also be combined with either a ProjectDomainName or ProjectDomainID.
	// ProjectName cannot be combined with ProjectID in a scope.
	ProjectName string `json:"projectName,omitempty"`

	// ProjectID is the unique ID of a project.
	// It can be used to scope a user to a specific project.
	ProjectID string `json:"projectID,omitempty"`

	// UserDomainName is the name of the domain where a user resides.
	// It is used to identify the source domain of a user.
	UserDomainName string `json:"userDomainName,omitempty"`

	// UserDomainID is the unique ID of the domain where a user resides.
	// It is used to identify the source domain of a user.
	UserDomainID string `json:"userDomainID,omitempty"`

	// ProjectDomainName is the name of the domain where a project resides.
	// It is used to identify the source domain of a project.
	// ProjectDomainName can be used in addition to a ProjectName when scoping
	// a user to a specific project.
	ProjectDomainName string `json:"projectDomainName,omitempty"`

	// ProjectDomainID is the name of the domain where a project resides.
	// It is used to identify the source domain of a project.
	// ProjectDomainID can be used in addition to a ProjectName when scoping
	// a user to a specific project.
	ProjectDomainID string `json:"projectDomainID,omitempty"`

	// DomainName is the name of a domain which can be used to identify the
	// source domain of either a user or a project.
	// If UserDomainName and ProjectDomainName are not specified, then DomainName
	// is used as a default choice.
	// It can also be used be used to specify a domain-only scope.
	DomainName string `json:"domainName,omitempty"`

	// DomainID is the unique ID of a domain which can be used to identify the
	// source domain of eitehr a user or a project.
	// If UserDomainID and ProjectDomainID are not specified, then DomainID is
	// used as a default choice.
	// It can also be used be used to specify a domain-only scope.
	DomainID string `json:"domainID,omitempty"`

	// DefaultDomain is the domain ID to fall back on if no other domain has
	// been specified and a domain is required for scope.
	DefaultDomain string `json:"defaultDomain,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenStackCredentialList contains a list of OpenStackCredential
type OpenStackCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []OpenStackCredential `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSCredential is the Schema for the awscredentials API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type AWSCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSCredentialSpec   `json:"spec"`
	Status AWSCredentialStatus `json:"status,omitempty"`
}

type AWSCredentialSpec struct {
	// AWS Access key ID
	AccessKeyID string `json:"accessKeyID"`
	// AWS Secret Access Key
	SecretAccessKey *SecretValue `json:"secretAccessKey"`
}

type AWSCredentialStatus struct {
	CredentialStatusMixin `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AWSCredentialList contains a list of AWSCredential
type AWSCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AWSCredential `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BYOCredential is the Schema for the byocredentials API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type BYOCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BYOCredentialSpec   `json:"spec"`
	Status BYOCredentialStatus `json:"status,omitempty"`
}

type BYOCredentialSpec struct {
	Docker     DockerEndpoint `json:"docker"`
	KubeConfig *SecretValue   `json:"kubeConfig"`
}

type BYOCredentialStatus struct {
	CredentialStatusMixin `json:",inline"`
}

type DockerEndpoint struct {
	Host       string       `json:"host"`
	CACert     []byte       `json:"caCert"`
	ClientCert []byte       `json:"clientCert"`
	ClientKey  *SecretValue `json:"clientKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AWSCredentialList contains a list of BYOCredential
type BYOCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BYOCredential `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VsphereCredential is the Schema for the vspherecredentials API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type VsphereCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VsphereCredentialSpec   `json:"spec"`
	Status VsphereCredentialStatus `json:"status,omitempty"`
}

type VsphereCredentialSpec struct {
	// VsphereConfig is a configuration of vSphere server
	VsphereConfig VsphereGlobalConfig `json:"vsphere"`

	// CluserApi is a struct containing credentials for operations
	ClusterAPI CredentialSpec `json:"clusterApi"`

	// CloudProvider is a struct containing configuration on defined vcenter server datacenter
	CloudProvider CredentialSpec `json:"cloudProvider"`
}

type VsphereGlobalConfig struct {
	// Server is an IP or FQDN of vSphere server (without port)
	Server string `json:"server"`

	// Port is a port of vSphere server
	Port string `json:"port,omitempty"`

	// Insecure is a flag that disables TLS peer verification
	Insecure bool `json:"insecure"`

	// Datacenter is a datacenter name in vSphere
	Datacenter string `json:"datacenter"`
}

type CredentialSpec struct {
	// Username is the username of the user.
	Username string `json:"username,omitempty"`

	// Password is the password of the user.
	Password *SecretValue `json:"password,omitempty"`
}

type VsphereCredentialStatus struct {
	CredentialStatusMixin `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VsphereCredentialList contains a list of VsphereCredential
type VsphereCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VsphereCredential `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalCredential is the Schema for the equinixmetalcredentials API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type EquinixMetalCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EquinixMetalCredentialSpec   `json:"spec"`
	Status EquinixMetalCredentialStatus `json:"status,omitempty"`
}

type EquinixMetalCredentialSpec struct {
	// ProjectID represents the Packet Project where this cluster will be placed into
	ProjectID string `json:"projectID"`
	// APIToken is a Project API key to access Equinix Metal API
	APIToken *SecretValue `json:"apiToken,omitempty"`
}

type EquinixMetalCredentialStatus struct {
	CredentialStatusMixin `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EquinixMetalCredentialList contains a list of EquinixMetalCredential
type EquinixMetalCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EquinixMetalCredential `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureCredential is the Schema for the azurecredentials API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type AzureCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureCredentialSpec   `json:"spec"`
	Status AzureCredentialStatus `json:"status,omitempty"`
}

type AzureCredentialSpec struct {
	// Environment is an optional field to select Azure cloud: AzureChinaCloud|AzureGermanCloud|AzureUSGovernmentCloud
	// otherwise AzurePublicCloud environment will be used
	// +kubebuilder:validation:Enum=AzurePublicCloud;AzureChinaCloud;AzureGermanCloud;AzureUSGovernmentCloud
	// +kubebuilder:default:=AzurePublicCloud
	Environment string `json:"environment,omitempty"`
	// SubscriptionID is the ID of the Azure subscription
	SubscriptionID string `json:"subscriptionID"`
	// TenantID is the ID of application directory (tenant)
	TenantID string `json:"tenantID"`
	// ClientID is the ID of application (client)
	ClientID string `json:"clientID"`
	// Client secret is an application password
	ClientSecret *SecretValue `json:"clientSecret,omitempty"`
}

type AzureCredentialStatus struct {
	CredentialStatusMixin `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AzureCredentialList contains a list of AzureCredential
type AzureCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AzureCredential `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&OpenStackCredential{}, &OpenStackCredentialList{},
		&AWSCredential{}, &AWSCredentialList{},
		&BYOCredential{}, &BYOCredentialList{},
		&VsphereCredential{}, &VsphereCredentialList{},
		&EquinixMetalCredential{}, &EquinixMetalCredentialList{},
		&AzureCredential{}, &AzureCredentialList{},
	)
}
