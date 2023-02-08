package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// HelmBundleSpec describes a set of Helm releases and repositories
// that they use.
type HelmBundleSpec struct {
	// The list of repositories used by this bundle.
	// +optional
	Repositories []HelmRepositorySpec `json:"repositories,omitempty"`
	// The list of releases in this bundle.
	Releases []HelmReleaseSpec `json:"releases"`
	// Release version of controller that is allowed to process this bundle.
	// +optional
	Release string `json:"release,omitempty"`
}

// HelmRepositorySpec describes a repository used by a HelmBundle.
type HelmRepositorySpec struct {
	// The name of the repository.
	Name string `json:"name"`
	// The URL of the repository.
	URL string `json:"url"`
	// Authentication secret name.
	// +optional
	AuthSecret string `json:"authSecret,omitempty"`
}

// HelmReleaseSpec specifies a Helm release.
type HelmReleaseSpec struct {
	// Chart name of the form "repo/name" or just "name".
	Chart string `json:"chart,omitempty"`
	// Chart version.
	Version string `json:"version,omitempty"`
	// Direct URL to the chart's tarball
	ChartURL string `json:"chartURL,omitempty"`
	// The name of this release.
	Name string `json:"name"`
	// The namespace of this release.
	Namespace string `json:"namespace"`
	// The labels to apply to the release.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// Values used for parametrization of the chart.
	// +optional
	// +kubebuilder:validation:XPreserveUnknownFields
	Values runtime.RawExtension `json:"values,omitempty"`
	// Deprecated: HelmV3 specifies whether Helm v3 should be used.
	// This field is deprecated as Helm v3 is always used.
	// +optional
	HelmV3 bool `json:"helmV3,omitempty"`
	// List of releases that must be installed and ready in order for this release to be installed
	// +optional
	Requires []string `json:"requires,omitempty"`
}

// HelmReleaseStatus describes a status of a release.
// +gocode:public-api=true
type HelmReleaseStatus struct {
	// Namespace of the release.
	// +optional
	Namespace string `json:"namespace"`
	// Revision of the release.
	Revision int `json:"revision,omitempty"`
	// Chart name of the form "repo/name" or just "name".
	Chart string `json:"chart"`
	// Chart version.
	Version string `json:"version"`
	// The hash of the parametrization values.
	Hash string `json:"hash"`
	// True if the release was installed successfully.
	// deprecated, use Status
	Success bool `json:"success"`
	// Status of the release
	Status string `json:"status"`
	// Deprecated: HelmV3 indicates whether the release was installed with Helm v3.
	// +optional
	HelmV3 bool `json:"helmV3"`
	// Notes contain the rendered notes for this release.
	// +nullable
	Notes string `json:"notes"`
	// Error message in case if the release installation failed.
	// +optional
	Message string `json:"message,omitempty"`
	// Attempt is a number of installs
	Attempt int `json:"attempt,omitempty"`
	// LastUpdateTime is a time of last try to install
	LastUpdateTime metav1.Time `json:"finishedAt,omitempty"`
	// Reflects readiness of the release
	// +optional
	Ready bool `json:"ready"`
}

// HelmBundleStatus describes a status of a HelmBundle.
type HelmBundleStatus struct {
	// The mapping of the names of the releases to their statuses.
	// +optional
	ReleaseStatuses map[string]HelmReleaseStatus `json:"releaseStatuses,omitempty"`
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// Release version of controller that processed this bundle.
	// +optional
	Release string `json:"release,omitempty"`
}

// HelmBundle specifies several Helm releases to be installed
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type HelmBundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmBundleSpec   `json:"spec,omitempty"`
	Status HelmBundleStatus `json:"status,omitempty"`
}

// HelmBundleList contains a list of HelmBundle objects
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type HelmBundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmBundle `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&HelmBundle{}, &HelmBundleList{})
}
