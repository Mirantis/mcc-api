package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const UnsupportedClustersName = "unsupported-clusters"

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=unsupportedclusters

// UnsupportedClusters represents clusters with cluster releases unsupported by coming kaas release
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Cluster,path=unsupportedclusters
type UnsupportedClusters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//Clusters is a list of unsupported clusters that should be upgraded to unblock release upgrade
	Clusters []ClusterReference `json:"clusters,omitempty"`
}

type ClusterReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	//SinceReleaseVersion field indicates a release version the cluster is not supported starting from
	SinceReleaseVersion string `json:"sinceReleaseVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UnsupportedClustersList contains a list of UnsupportedClusters
type UnsupportedClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UnsupportedClusters `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UnsupportedClusters{}, &UnsupportedClustersList{})
}
