package v1alpha1

import (
	"fmt"
	"reflect"
)

// Tags defines a map of tags.
// +gocode:public-api=true
type Tags map[string]string

// Equals returns true if the tags are equal.
func (t Tags) Equals(other Tags) bool {
	return reflect.DeepEqual(t, other)
}

// Difference returns the difference between this map of tags and the other map of tags.
// Items are considered equals if key and value are equals.
func (t Tags) Difference(other Tags) Tags {
	res := make(Tags, len(t))

	for key, value := range t {
		if otherValue, ok := other[key]; ok && value == otherValue {
			continue
		}
		res[key] = value
	}

	return res
}

// Merge merges in tags from other. If a tag already exists, it is replaced by the tag in other.
func (t Tags) Merge(other Tags) {
	for k, v := range other {
		t[k] = v
	}
}

const (
	// +gocode:public-api=true
	TagPrefix = "kaas.mirantis.com"
	// +gocode:public-api=true
	KubernetesTagPrefix = "kubernetes.io"

	// ResourceLifecycleOwned is the value we use when tagging resources to indicate
	// that the resource is considered owned and managed by the cluster,
	// and in particular that the lifecycle is tied to the lifecycle of the cluster.
	// +gocode:public-api=true
	ResourceLifecycleOwned = ResourceLifecycle("owned")

	// ResourceLifecycleShared is the value we use when tagging resources to indicate
	// that the resource is shared between multiple clusters, and should not be destroyed
	// if the cluster is destroyed.
	// +gocode:public-api=true
	ResourceLifecycleShared = ResourceLifecycle("shared")

	// APIServerRole describes the value for the apiserver role
	// +gocode:public-api=true
	APIServerRole = "apiserver"

	// NodeOutboundRole describes the value for the node outbound LB role
	// +gocode:public-api=true
	NodeOutboundRole = "nodeOutbound"

	// KubernetesServicesRole describes the value for the kubernetes services LB role
	// +gocode:public-api=true
	KubernetesServicesRole = "kubernetesServices"

	// ControlPlaneOutboundRole describes the value for the control plane outbound LB role
	// +gocode:public-api=true
	ControlPlaneOutboundRole = "controlPlaneOutbound"

	// BastionRole describes the value for the bastion role
	// +gocode:public-api=true
	BastionRole = "bastion"

	// CommonRole describes the value for the common role
	// +gocode:public-api=true
	CommonRole = "common"

	// VMTagsLastAppliedAnnotation is the key for the machine object annotation
	// which tracks the AdditionalTags in the Machine Provider Config.
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	// for annotation formatting rules.
	// +gocode:public-api=true
	VMTagsLastAppliedAnnotation = "sigs.k8s.io/cluster-api-provider-azure-last-applied-tags-vm"
)

var (

	// ClusterIDTagKey is the key for resources associated with cluster by unique identifier
	// +gocode:public-api=true
	ClusterIDTagKey = fmt.Sprintf("%s_clusterID", TagPrefix)

	// NamespaceTagKey is the key for cluster namespace
	// +gocode:public-api=true
	NamespaceTagKey = fmt.Sprintf("%s_namespace", TagPrefix)

	// ClusterNameTagKey is the key for cluster name
	// +gocode:public-api=true
	ClusterNameTagKey = fmt.Sprintf("%s_cluster", TagPrefix)

	// MachineRoleTagKey is the key for kubernetes machine role
	// +gocode:public-api=true
	MachineRoleTagKey = fmt.Sprintf("%s_role", KubernetesTagPrefix)
)

// ResourceLifecycle configures the lifecycle of a resource
// +gocode:public-api=true
type ResourceLifecycle string

const (
	// +gocode:public-api=true
	ControlPlaneTag = "kubernetes.io/role:master"
	// +gocode:public-api=true
	WorkerTag = "kubernetes.io/role:node"
)

// BuildParams is used to build tags around an azure resource.
// +gocode:public-api=true
type BuildParams struct {
	// ClusterName is the cluster associated with the resource.
	ClusterName string

	// Namespace is the namespace of the cluster, it's applied as the tag "Namespace" on Azure.
	Namespace string

	// Role is the role associated to the resource.
	// +optional
	Role *string

	// Any additional tags to be added to the resource.
	// +optional
	Additional Tags
}

// Build builds tags including the cluster tag and returns them in map form.
// +gocode:public-api=true
func Build(params BuildParams) Tags {
	tags := make(Tags)
	for k, v := range params.Additional {
		tags[k] = v
	}

	if params.ClusterName != "" {
		tags[ClusterNameTagKey] = params.ClusterName
	}

	if params.Namespace != "" {
		tags[NamespaceTagKey] = params.Namespace
	}

	return tags
}
