package v1alpha1

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ClusterStatus defines the observed state of Cluster
// +gocode:public-api=true
type ClusterStatus struct {

	// If set, indicates that there is a problem reconciling the
	// state, and will be set to a token value suitable for
	// programmatic interpretation.
	// +optional
	ErrorReason common.ClusterStatusError `json:"errorReason,omitempty"`

	// If set, indicates that there is a problem reconciling the
	// state, and will be set to a descriptive error message.
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Provider-specific status.
	// It is recommended that providers maintain their
	// own versioned API types that should be
	// serialized/deserialized from this field.
	// +optional
	// +kubebuilder:validation:XPreserveUnknownFields
	ProviderStatus *runtime.RawExtension `json:"providerStatus,omitempty"`
}

// ClusterList contains a list of Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// +gocode:public-api=true
const ClusterFinalizer = "cluster.cluster.k8s.io"

// Cluster is the Schema for the clusters API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=cl
// +kubebuilder:subresource:status
// +gocode:public-api=true
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

func (o *Cluster) Validate() field.ErrorList {
	errors := field.ErrorList{}

	if len(o.Spec.ClusterNetwork.Pods.CIDRBlocks) == 0 {
		errors = append(errors, field.Invalid(
			field.NewPath("Spec", "ClusterNetwork", "Pods"),
			o.Spec.ClusterNetwork.Pods,
			"invalid cluster configuration: missing Cluster.Spec.ClusterNetwork.Pods"))
	}
	if len(o.Spec.ClusterNetwork.Services.CIDRBlocks) == 0 {
		errors = append(errors, field.Invalid(
			field.NewPath("Spec", "ClusterNetwork", "Services"),
			o.Spec.ClusterNetwork.Services,
			"invalid cluster configuration: missing Cluster.Spec.ClusterNetwork.Services"))
	}
	return errors
}

// ClusterSpec defines the desired state of Cluster
// +gocode:public-api=true
type ClusterSpec struct {
	// Cluster network configuration
	ClusterNetwork *ClusterNetworkingConfig `json:"clusterNetwork,omitempty"`

	// Provider-specific serialized configuration to use during
	// cluster creation. It is recommended that providers maintain
	// their own versioned API types that should be
	// serialized/deserialized from this field.
	// +optional
	ProviderSpec ProviderSpec `json:"providerSpec,omitempty"`
}

// ClusterNetworkingConfig specifies the different networking
// parameters for a cluster.
type ClusterNetworkingConfig struct {
	// The network ranges from which service VIPs are allocated.
	Services NetworkRanges `json:"services,omitempty"`

	// The network ranges from which Pod networks are allocated.
	Pods NetworkRanges `json:"pods,omitempty"`
}

// NetworkRanges represents ranges of network addresses.
type NetworkRanges struct {
	CIDRBlocks []string `json:"cidrBlocks,omitempty" sensitive:"true"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
