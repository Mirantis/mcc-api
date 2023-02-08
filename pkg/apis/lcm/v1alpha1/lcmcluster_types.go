package v1alpha1

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	errs "github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

type LCMStateItems []LCMStateItem

// LCMClusterSpec defines the desired state of LCMCluster
type LCMClusterSpec struct {
	// Mapping of machine types to the lists of state items.
	MachineTypes map[LCMMachineType]LCMStateItems `json:"machineTypes"`
	// Components which are upgraded/managed separatelly from StateItems
	Components LCMComponents `json:"components,omitempty"`
	// The endpoint that should be used to access the tenant cluster.
	// +optional
	TenantClusterEndpoint LCMClusterEndpoint `json:"externalEndpoint,omitempty"`
	// The endpoint that should be used to access the management cluster.
	ManagementClusterEndpoint LCMClusterEndpoint `json:"managementClusterEndpoint"`
	// HelmControllerHAEnabled indicates if helm-controller is running in HA mode
	HelmControllerHAEnabled bool `json:"helmControllerHAEnabled,omitempty"`
	// Service account name to be used for HelmController
	HelmControllerServiceAccountName string `json:"helmControllerServiceAccountName"`
	// DeploymentSecretName specifies the name of the secret
	// that's used for deployment
	// TODO move it to a CRD
	DeploymentSecretName string `json:"deploymentSecretName,omitempty"`
	// KubeconfigSecretName specifies the name of the secret
	// that's used for tenant cluster's kubeconfig
	// TODO move it to AgentDataSecretName
	KubeconfigSecretName string `json:"kubeconfigSecretName,omitempty"`
	// ProxySecretName specifies the name of the secret
	// that's used for setting proxy params
	ProxySecretName string `json:"proxySecretName,omitempty"`
	// AgentDataSecretName specifies the name of the secret
	// that's used for data sent from agent
	AgentDataSecretName string `json:"agentDataSecretName,omitempty"`
	// MKETLSSecretName specifies the name of the secret
	// that's used for setting TLS configuration for MKE
	MKETLSSecretName string `json:"mkeTLSSecretName,omitempty"`
	// TokenTTL specifies TTL (time to live) value for tokens.
	// Note that the initial token is currently
	// considered valid for this duration, so it shoudn't be less
	// than 24h currently. Defaults to 24 hours.
	// +optional
	TokenTTL *metav1.Duration `json:"tokenTTL,omitempty"`
	// MinTokenRemainingTTL specifies the minimum remaining
	// TTL (time to live) for a token that is to be used
	// to join a new node. Defaults to 1 hour.
	// +optional
	MinTokenRemainingTTL *metav1.Duration `json:"minTokenRemainingTTL,omitempty"`
	// DedicatedControlPlane specifies that the cluster's
	// control nodes must be tainted. Defaults to true.
	// +optional
	DedicatedControlPlane *bool `json:"dedicatedControlPlane,omitempty"`
	// Images specifies the images to use in the child cluster
	// +optional
	Images LCMClusterImages `json:"images,omitempty"`
	// Max number of workers being prepared at the given moment.
	// This is used to limit the network load that can occur
	// when downloading the files to the nodes.
	MaxWorkerPrepareCount int `json:"maxWorkerPrepareCount,omitempty"`
	// Max number of workers to upgrade at the same time.
	// This is used to avoid breaking the workloads for which
	// PodDisruptionBudgets aren't configured. It's implemented
	// as an upper limit on the numebr of machines that are
	// cordoned at the given moment.
	// Defaults to 1.
	// +optional
	MaxWorkerUpgradeCount int `json:"maxWorkerUpgradeCount,omitempty"`
	// NodesRebootRequired flag is true when nodes must be rebooted after upgrade
	NodesRebootRequired bool `json:"nodesRebootRequired,omitempty"`

	// OIDC specifies OpenID configuration
	// +optional
	OIDC *LCMOIDCSettings `json:"oidc,omitempty"`
	// MCC release associated with requested configuration
	Release string `json:"release,omitempty"`
	// Requested version of MCR
	MCRVersion string `json:"mcrVersion,omitempty"`
	// Hash of proxy configuration
	ProxyHash string `json:"proxyHash,omitempty"`

	// Indicates if helm-controller would be installed as separate chart (true)
	// or from lcm-controller chart by lcm-controller (false)
	HelmControllerFromExternalChart bool `json:"helmControllerFromExternalChart,omitempty"`

	// Maintenance flag indicates that the cluster should be switched into maintenance state.
	Maintenance bool `json:"maintenance,omitempty"`

	// LCMType contains the LCM distribution type
	// +kubebuilder:validation:Enum=ucp;byo;k0s
	LCMType LCMType `json:"lcmType,omitempty"`

	// Version of LCM agent machines of this cluster should have
	AgentVersion string `json:"agentVesrion,omitempty"`

	// MKEUpgradeAttempts specifies a custom value for MKE Update procedure retries
	// If not set the default value results total 5 retries
	// +optional
	// +kubebuilder:default=5
	MKEUpgradeAttempts uint64 `json:"mkeUpgradeAttempts,omitempty"`
}

// LCMClusterList contains a list of LCMCluster objects
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LCMCluster `json:"items"`
}

const (
	// +gocode:public-api=true
	ClusterSecretAnnotation = "lcm.mirantis.com/cluster-secret-type"
	// +gocode:public-api=true
	ClusterSecretTypeKubeconfig = "kubeconfig"
	// +gocode:public-api=true
	ClusterSecretTypeAgentData = "agentData"
	// +gocode:public-api=true
	ClusterSecretTypeDeployment = "deployment"
	// +gocode:public-api=true
	ClusterSecretClusterNameAnnotation = "lcm.mirantis.com/cluster-name"
	// default token TTL
	// +gocode:public-api=true
	DefaultTokenTTL = 24 * time.Hour
	// the duration of the token to be considered "reliable"
	// (i.e. so the playbooks will have enough time to process it
	// before it expires)
	// +gocode:public-api=true
	DefaultMinTokenRemainingTokenTTL = time.Hour
	// this annotation set to true on new history objects
	// +gocode:public-api=true
	NewHistoryObjectAnnotation = "lcm.mirantis.com/new-history"
)

// LCMClusterStatus defines the observed state of LCMCluster
type LCMClusterStatus struct {
	// The number of nodes requested
	RequestedNodes int `json:"requestedNodes"`
	// The number of nodes ready
	ReadyNodes int `json:"readyNodes"`
	// Helm controller installed
	HelmControllerDeployed bool `json:"helmControllerDeployed"`
	// Status of components
	Components LCMComponentsStatus `json:"components,omitempty"`
	// Applied OIDC configuration
	OIDC             LCMOIDCSettings `json:"oidc,omitempty"`
	OIDCUpdateTime   *metav1.Time    `json:"oidcUpdateTime,omitempty"`
	OIDCSettingsHash *string         `json:"oidcSettingHash,omitempty"`
	// MCC release associated with processed configuration
	Release string `json:"release,omitempty"`
	// Current version of MCR
	MCRVersion string `json:"mcrVersion,omitempty"`
	// Hash of proxy configuration
	ProxyHash string `json:"proxyHash,omitempty"`
	// Maintenance flag indicates that the cluster is in maintenance state
	Maintenance bool `json:"maintenance,omitempty"`
	// Hash of license data
	MCCLicenseHash string `json:"mccLicenseHash,omitempty"`
	// Hash of MKE TLS data
	MKETLSSettingsHash string `json:"mkeTLSSettingsHash,omitempty"`
	// MKEAuthHash is set, when custom tls config for regional MKE is set
	// and all nodes have a fallback kubeconfig with custom MKE auth data prepared.
	MKEAuthHash string `json:"mkeAuthHash,omitempty"`
	// LCMOperationStuck flag indicates that some LCM operation with the cluster
	// is stuck, and operator needs to take a closer look at its status
	LCMOperationStuck bool `json:"lcmOperationStuck,omitempty"`
}

// LCMCluster is the Schema for the lcmclusters API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="requested",type="integer",JSONPath=".status.requestedNodes",description="Requested Nodes",priority=0
// +kubebuilder:printcolumn:name="ready",type="integer",JSONPath=".status.readyNodes",description="Ready Nodes",priority=0
// +kubebuilder:printcolumn:name="helmCtlDeployed",type="boolean",JSONPath=".status.helmControllerDeployed",description="Helm Ctl Deployed",priority=0
// +gocode:public-api=true
type LCMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LCMClusterSpec   `json:"spec,omitempty"`
	Status LCMClusterStatus `json:"status,omitempty"`
}

// GetTokenTTL returns TokenTTL value for the cluster
// defaulting to DefaultTokenTTL
func (c *LCMCluster) GetTokenTTL() time.Duration {
	if c.Spec.TokenTTL == nil {
		return DefaultTokenTTL
	}
	return c.Spec.TokenTTL.Duration
}

// GetMinTokenRemainingTTL returns MinTokenRemainingTTL value for the cluster
// defaulting to DefaultMinTokenRemainingTokenTTL
func (c *LCMCluster) GetMinTokenRemainingTTL() time.Duration {
	if c.Spec.MinTokenRemainingTTL == nil {
		return DefaultMinTokenRemainingTokenTTL
	}
	return c.Spec.MinTokenRemainingTTL.Duration
}
func (c *LCMCluster) IsDedicatedControlPlane() bool {
	if c.Spec.DedicatedControlPlane == nil {
		return true
	}
	return *c.Spec.DedicatedControlPlane
}
func (c *LCMCluster) IsBYO() bool {
	return c.Spec.LCMType == LCMTypeBYO
}
func (c *LCMCluster) IsMKE() bool {
	return c.Spec.LCMType == LCMTypeMKE
}
func (c *LCMCluster) IsK0s() bool {
	return c.Spec.LCMType == LCMTypeK0s
}

// LCMClusterEndpoint specifies the host and port of a cluster's
// apiserver.
type LCMClusterEndpoint struct {
	// Host name of the cluster's apiserver
	Host string `json:"host,omitempty"`
	// Port number of the cluster's apiserver
	Port int `json:"port,omitempty"`
}

// LCMClusterImages specifies the images to use in the child cluster.
type LCMClusterImages struct {
	// HelmController specifies the image for helm-controller.
	// +optional
	HelmController string `json:"helmController,omitempty"`
	// Tiller specifies the image for tiller.
	// +optional
	Tiller string `json:"tiller,omitempty"`
	// UCPSourceRepo specifies the source repository for UCP images.
	// +optional
	UCPSourceRepo string `json:"ucpSourceRepo,omitempty"`
	// UCPTargetRepo specifies the target repository for UCP images.
	// +optional
	UCPTargetRepo string `json:"ucpTargetRepo,omitempty"`
}

// LCMOIDCSettings specifies settings for OpenID
type LCMOIDCSettings struct {
	// URL specifies OpenID Provider URL
	URL string `json:"url"`
	// CACert specifies CA certificate to use for communication
	CACert string `json:"caCert"`
	// ClientID specifies OpenID client id
	ClientID string `json:"clientID"`
	// AdminRole specifies admin role
	AdminRole string `json:"adminRole"`
	// GroupsClaim specifies groups claims for admin role
	GroupsClaim string `json:"groupsClaim"`
}

// IsClusterSecret returns true if the specified secret refers to a
// kubeconfig or deployment secret
// +gocode:public-api=true
func IsClusterSecret(secret metav1.Object) bool {
	ann := secret.GetAnnotations()
	return ann[ClusterSecretAnnotation] != "" && ann[ClusterSecretClusterNameAnnotation] != ""
}

// ClusterNameFromSecret returns the NamespacedName of the cluster
// for the specified cluster secret.
// +gocode:public-api=true
func ClusterNameFromSecret(secret metav1.Object) types.NamespacedName {
	ann := secret.GetAnnotations()
	return types.NamespacedName{
		Name:      ann[ClusterSecretClusterNameAnnotation],
		Namespace: secret.GetNamespace(),
	}
}

// +gocode:public-api=true
func ProxyHashFromCluster(cluster *LCMCluster) (string, error) {
	proxySettings := ProxyStateItemParams(cluster.Spec.MachineTypes[LCMMachineTypeControl])
	if proxySettings == nil {
		return "", nil
	}
	return ProxyHash(proxySettings)
}

// +gocode:public-api=true
func ProxyHash(settings map[string]string) (string, error) {
	out, err := json.Marshal(settings)
	if err != nil {
		return "", errs.Wrap(err, "failed to marshal proxy settings")
	}
	return fmt.Sprintf("%x", sha256.Sum256(out)), nil
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&LCMCluster{}, &LCMClusterList{})
}
