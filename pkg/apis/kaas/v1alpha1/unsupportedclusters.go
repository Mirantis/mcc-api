package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +gocode:public-api=true
const UnsupportedClustersName = "unsupported-clusters"

type ClusterReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	// SinceReleaseVersion field indicates a release version the cluster is not supported starting from
	SinceReleaseVersion string `json:"sinceReleaseVersion"`
}

// UnsupportedClusters represents clusters with cluster releases unsupported by coming kaas release
// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=unsupportedclusters
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Cluster,path=unsupportedclusters
// +gocode:public-api=true
type UnsupportedClusters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Clusters is a list of unsupported clusters that should be upgraded to unblock release upgrade
	Clusters []ClusterReference `json:"clusters,omitempty"`
}

// UnsupportedClustersList contains a list of UnsupportedClusters
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type UnsupportedClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UnsupportedClusters `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&UnsupportedClusters{}, &UnsupportedClustersList{})
}
