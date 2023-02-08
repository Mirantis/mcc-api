package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"

	lcm "github.com/Mirantis/mcc-api/v2/pkg/apis/common/lcm/v1alpha1"
)

type ClusterSpecMixin struct {
	// Release contains name of the ClusterRelease to install on cluster
	// +optional
	Release string `json:"release,omitempty"`
	// HelmReleases is a list of Helm releases in release that should be deployed
	// +optional
	HelmReleases []HelmRelease `json:"helmReleases,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
	// The name of the Credentials object
	Credentials string `json:"credentials"`
	// DedicatedControlPlane specifies that the cluster's
	// control nodes must be tainted. Defaults to true.
	// +optional
	DedicatedControlPlane *bool `json:"dedicatedControlPlane,omitempty"`
	// KaaS defines enabled KaaS pieces on this cluster
	// +optional
	KaaS KaaS `json:"kaas,omitempty"`
	// List of references to objects containing the ssh public key
	// +optional
	PublicKeys []PublicKeyRef `json:"publicKeys,omitempty"`
	// The name of the Proxy object
	// +optional
	Proxy string `json:"proxy,omitempty"`
	// TLS configuration for cluster's applications
	// +optional
	TLS TLS `json:"tls,omitempty"`
	// Max number of workers being prepared at the given moment.
	// This is used to limit the network load that can occur
	// when downloading the files to the nodes.
	MaxWorkerPrepareCount *int `json:"maxWorkerPrepareCount,omitempty"`
	// Max number of workers to upgrade at the same time.
	// This is used to avoid breaking the workloads for which
	// PodDisruptionBudgets aren't configured. It's implemented
	// as an upper limit on the number of machines that are
	// cordoned at the given moment.
	MaxWorkerUpgradeCount *int `json:"maxWorkerUpgradeCount,omitempty"`
	// Maintenance defines if cluster should switch into maintenance state
	// +optional
	Maintenance bool `json:"maintenance,omitempty"`
}

// Represents a Helm Release that will be installed into a cluster.
// Values will be merged with values in Release object
type HelmRelease struct {
	// Name of the release
	Name string `json:"name"`
	// Enabled specifies that the release must be installed
	Enabled *bool `json:"enabled,omitempty"`
	// Release Values overrides
	// +optional
	Values runtime.RawExtension `json:"values,omitempty"`
}

func (helmRelease *HelmRelease) IsEnabled() bool {
	return helmRelease.Enabled == nil || *helmRelease.Enabled
}

// OIDC configuration in the cluster object
type OIDC struct {
	Ready       bool   `json:"ready"`
	IssuerURL   string `json:"issuerUrl"`
	ClientID    string `json:"clientId"`
	GroupsClaim string `json:"groupsClaim"`
	// base64 encoded certificate
	Certificate string `json:"certificate"`
}

type LoadBalancerStatus struct {
	ID     string `json:"id,omitempty"`
	Ready  bool   `json:"ready"`
	Status string `json:"status,omitempty"`
}

type KaaS struct {
	Release              string                `json:"release,omitempty"`
	Management           ManagementClusterSpec `json:"management,omitempty"`
	Regional             []RegionalClusterSpec `json:"regional,omitempty"`
	RegionalHelmReleases []HelmRelease         `json:"regionalHelmReleases,omitempty"`
}

type ManagementClusterSpec struct {
	Enabled      bool          `json:"enabled"`
	HelmReleases []HelmRelease `json:"helmReleases,omitempty"`
	// AutoSyncSalesForceConfig defines if SalesForce config of Management cluster
	// should be propagated to all regional and child clusters.
	AutoSyncSalesForceConfig bool `json:"autoSyncSalesForceConfig,omitempty"`
}

type RegionalClusterSpec struct {
	Provider     string        `json:"provider"`
	HelmReleases []HelmRelease `json:"helmReleases,omitempty"`
}

// ReleaseReferenceLong represents a Release Reference. It has enough information to retrieve Release
// in any namespace and basic information to show to user
type ReleaseReferenceLong struct {
	// Version of the relese
	Version string `json:"version"`
	// Name is unique within a namespace to reference a release resource.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace defines the space within which the release name must be unique.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// RebootRequired indicates that machines will be rebooted as part of the upgrade.
	// +optional
	RebootRequired bool `json:"rebootRequired,omitempty"`
}

type CurrentReleaseReferenceLong struct {
	ReleaseReferenceLong `json:",inline"`
	// UnsupportedSinceKaaSVersion reveals that current cluster's release is not supported
	// +optional
	UnsupportedSinceKaaSVersion string             `json:"unsupportedSinceKaaSVersion,omitempty"`
	LCMType                     lcm.LCMType        `json:"lcmType,omitempty"`
	AllowedNodeLabels           []AllowedNodeLabel `json:"allowedNodeLabels,omitempty"`
}

type ClusterStatusMixin struct {
	Nodes                                *NodesStatus        `json:"nodes,omitempty"`
	ReleaseRefs                          *ReleaseRefs        `json:"releaseRefs,omitempty"`
	LoadBalancerHost                     string              `json:"loadBalancerHost,omitempty"`
	LoadBalancerStatus                   *LoadBalancerStatus `json:"loadBalancerStatus,omitempty"`
	APIServerCertificate                 []byte              `json:"apiServerCertificate"`
	Helm                                 HelmStatus          `json:"helm,omitempty"`
	PersistentVolumesProviderProvisioned bool                `json:"persistentVolumesProviderProvisioned,omitempty"`
	ObservedGeneration                   int64               `json:"observedGeneration,omitempty"`
	NotReadyObjects                      Objects             `json:"notReadyObjects,omitempty"`
	Ceph                                 *CephStatus         `json:"ceph,omitempty"`
	Maintenance                          bool                `json:"maintenance,omitempty"`
	// OIDC issuer configuration.
	OIDC              *OIDC  `json:"oidc,omitempty"`
	OIDCSettingsHash  string `json:"oidcSettingsHash"`
	UCPDashboard      string `json:"ucpDashboard,omitempty"`
	ConditionsSummary `json:",inline"`
	MKE               *MKE         `json:"mke,omitempty"`
	TLS               MCCTLSStatus `json:"tls,omitempty"`
}

// MKE stores MKE-specific information
type MKE struct {
	ClusterID string `json:"clusterID,omitempty"`
}

const (
	StacklightCondition   ConditionType = "StackLight"
	HelmCondition         ConditionType = "Helm"
	KubernetesCondition   ConditionType = "Kubernetes"
	BastionCondition      ConditionType = "Bastion"
	NodeCondition         ConditionType = "Nodes"
	OIDCCondition         ConditionType = "OIDC"
	LoadBalancerCondition ConditionType = "LoadBalancer"
	CephCondition         ConditionType = "Ceph"
)

type NodesStatus struct {
	Requested int `json:"requested"`
	Ready     int `json:"ready"`
}

type ReleaseRefs struct {
	Current   CurrentReleaseReferenceLong `json:"current"`
	Previous  CurrentReleaseReferenceLong `json:"previous,omitempty"`
	Available []ReleaseReferenceLong      `json:"available"`
}

type HelmStatus struct {
	Ready    bool               `json:"ready"`
	Releases HelmReleasesStatus `json:"releases,omitempty"`
}

type CephStatus struct {
	Ready   bool   `json:"ready"`
	Message string `json:"message,omitempty"`
}
