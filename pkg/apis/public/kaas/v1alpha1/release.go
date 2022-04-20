package v1alpha1

import (
	"strings"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	lcm "github.com/Mirantis/mcc-api/pkg/apis/common/lcm/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ClusterRelease represents one release for LCM
// +kubebuilder:resource:scope=Cluster
type ClusterRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterReleaseSpec `json:"spec"`
}

// ReleaseSpec defines release parameters
type ClusterReleaseSpec struct {
	// Version string for this release.
	Version string `json:"version"`
	// Description for the release
	Description string `json:"description"`
	// ReleaseNotes provide information about changes between releases
	ReleaseNotes []ReleaseNotesVersion `json:"releaseNotes"`
	// Url to fetch lcm-agent binary
	Agent lcm.AgentConfig `json:"agent"`
	// Mapping of machine types to the lists of state items.
	MachineTypes map[lcm.LCMMachineType]lcm.LCMStateItems `json:"machineTypes"`
	// Helm repositories and releases description
	Helm Helm `json:"helm,omitempty"`
	// Images specifies the images to use in the child cluster
	Images lcm.LCMClusterImages `json:"images,omitempty"`
	// Kubernetes LCM type
	LCMType lcm.LCMType `json:"lcmType,omitempty"`
	// The list of allowed node labels to add to the cluster node
	// to run certain components on dedicated nodes
	AllowedNodeLabels []AllowedNodeLabel `json:"allowedNodeLabels,omitempty"`
	// Contains a list of per-provider allowed distributions.
	AllowedDistributions []Distribution `json:"allowedDistributions,omitempty"`
}

// GetDefaultDistribution returns first encountered object in list that has
// default field set to True.
func (cr *ClusterRelease) GetDefaultDistribution() (Distribution, error) {
	for _, el := range cr.Spec.AllowedDistributions {
		if el.Default {
			return el, nil
		}
	}

	return Distribution{}, errors.Errorf("no default distribution defined")
}

// GetDistributionByID returns first encountered object in list that has
// specified ID.
func (cr *ClusterRelease) GetDistributionByID(id string) (Distribution, error) {
	if id == "" {
		return Distribution{}, errors.New("distribution ID can't be empty")
	}

	for _, el := range cr.Spec.AllowedDistributions {
		if strings.EqualFold(id, el.ID) {
			return el, nil
		}
	}

	return Distribution{}, errors.Errorf("no distribution with ID '%s' found", id)
}

type RemoteArtifact struct {
	// URL specifies the location of the original remote artifact file in the
	// vendor infrastructure.
	URL string `json:"url"`

	// Checksum contains the verification sum for the remote artifact file, with
	// optional prefix to the encryption algorithm, separated by a semicolon.
	Checksum string `json:"checksum"`
}

type Distribution struct {
	// ID is a unique identifier of the distro. It is used in the
	// Machine object spec to define the distro for the machine.
	ID string `json:"id"`

	// Version contains the version of the operating system distro.
	Version string `json:"version"`

	// Description is a text description of the distro for UI.
	Description string `json:"description,omitempty"`

	// Default flag, when set to 'true' and there is no other user-defined
	// distribution field in the Machine spec, will be used as default distro.
	Default bool `json:"default,omitempty"`

	// Image is object that defines the
	// provisioning image for the distro.
	Image RemoteArtifact `json:"image"`
}

type AllowedNodeLabel struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
}

// ReleaseNotesVersion contains information about new changes in this version
type ReleaseNotesVersion struct {
	// Version for these notes
	Version string `json:"version"`
	// Notes contains actual notes
	Notes []ReleaseNotesItem `json:"notes"`
}

// ReleaseNotesItem contains one item from release notes (one separate change)
type ReleaseNotesItem struct {
	// Text describes the change
	Text string `json:"text"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterReleaseList contains a list of Release
type ClusterReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ClusterRelease `json:"items"`
}

// Helm repositories and releases description
type Helm struct {
	// List of Helm repositories available in this release
	Repositories []lcm.HelmRepositorySpec `json:"repositories,omitempty"`
	// Liste of Helm releases available in this release
	Releases []HelmReleaseSpec `json:"releases"`
	// Description of helm-controller release
	Controller HelmReleaseSpec `json:"controller,omitempty"`
}

type HelmReleaseSpec struct {
	lcm.HelmReleaseSpec `json:",inline"`
	// Set if this Helm release is required for overall release (e.g. cloud provider)
	// +optional
	Required bool `json:"required,omitempty"`
	// Set if this Helm release requires some default StorageClass to be present in
	// the cluster. It won't be deployed until condition is met
	// +optional
	RequiresPersistentVolumes bool `json:"requiresPersistentVolumes,omitempty"`
	// Set if this Helm release requires local volumes, provided by
	// local-volume-provisioner (lvp should be in requires list)
	LocalVolumes []HelmReleaseLocalVolume `json:"localVolumes,omitempty"`
}

type HelmReleaseLocalVolume struct {
	// Name for local volume dir
	Name string `json:"name"`
	// Volumes mount points
	BindMounts []LocalVolumeBindMount `json:"bindMounts"`
}

type LocalVolumeBindMount struct {
	// Number of volumes per node with defined name
	VolumePerNode int64 `json:"volPerNode"`
	// Prefix path for mount point (`/var/lib/local-volumes` by default)
	Prefix string `json:"prefix,omitempty"`
}

// HelmReleaseSpecList for overrides helmReleaseSpecs in other structs
type HelmReleaseList struct {
	// List of Helm releases available in this release
	Releases []HelmRelease `json:"helmReleases"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// KaaSRelease describes release of KaaS itself
// +kubebuilder:resource:scope=Cluster
type KaaSRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KaaSReleaseSpec `json:"spec"`
}

type KaaSReleaseSpec struct {
	Version                  string                    `json:"version"`
	ClusterRelease           string                    `json:"clusterRelease"`
	HelmRepositories         []lcm.HelmRepositorySpec  `json:"helmRepositories,omitempty"`
	Management               ManagementSpec            `json:"management"`
	Regional                 []RegionalSpec            `json:"regional"`
	RegionalHelmReleases     []HelmReleaseSpec         `json:"regionalHelmReleases,omitempty"`
	SupportedClusterReleases []SupportedClusterRelease `json:"supportedClusterReleases,omitempty"`
	Bootstrap                Bootstrap                 `json:"bootstrap"`
}

type ManagementSpec struct {
	HelmReleases []HelmReleaseSpec `json:"helmReleases"`
}

type RegionalSpec struct {
	Provider     string            `json:"provider"`
	HelmReleases []HelmReleaseSpec `json:"helmReleases"`
}

type SupportedClusterRelease struct {
	Version           string             `json:"version"`
	Name              string             `json:"name"`
	Tag               string             `json:"tag,omitempty"`
	AvailableUpgrades []AvailableUpgrade `json:"availableUpgrades,omitempty"`
	Providers         Providers          `json:"providers,omitempty"`
}

type AvailableUpgrade struct {
	Version        string `json:"version"`
	RebootRequired bool   `json:"rebootRequired,omitempty"`
}

type Providers struct {
	Supported []string `json:"supported,omitempty"`
}

func (p *Providers) Contain(provider string) bool {
	if len(p.Supported) == 0 {
		return true
	}
	for _, s := range p.Supported {
		if s == provider {
			return true
		}
	}
	return false
}

type Bootstrap struct {
	Version      string            `json:"version"`
	HelmReleases []HelmReleaseSpec `json:"helmReleases,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KaaSReleaseList contains a list of KaaSRelease
type KaaSReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KaaSRelease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterRelease{}, &ClusterReleaseList{}, &KaaSRelease{}, &KaaSReleaseList{})
}
