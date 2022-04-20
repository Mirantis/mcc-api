/*
Copyright 2021 The Mirantis Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"fmt"
	"reflect"
)

const (
	ControlPlaneTag = "kubernetes.io/role:master"
	WorkerTag       = "kubernetes.io/role:node"
)

// Tags defines a map of tags.
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

// ResourceLifecycle configures the lifecycle of a resource
type ResourceLifecycle string

const (
	TagPrefix           = "kaas.mirantis.com"
	KubernetesTagPrefix = "kubernetes.io"

	// ResourceLifecycleOwned is the value we use when tagging resources to indicate
	// that the resource is considered owned and managed by the cluster,
	// and in particular that the lifecycle is tied to the lifecycle of the cluster.
	ResourceLifecycleOwned = ResourceLifecycle("owned")

	// ResourceLifecycleShared is the value we use when tagging resources to indicate
	// that the resource is shared between multiple clusters, and should not be destroyed
	// if the cluster is destroyed.
	ResourceLifecycleShared = ResourceLifecycle("shared")

	// APIServerRole describes the value for the apiserver role
	APIServerRole = "apiserver"

	// NodeOutboundRole describes the value for the node outbound LB role
	NodeOutboundRole = "nodeOutbound"

	// KubernetesServicesRole describes the value for the kubernetes services LB role
	KubernetesServicesRole = "kubernetesServices"

	// ControlPlaneOutboundRole describes the value for the control plane outbound LB role
	ControlPlaneOutboundRole = "controlPlaneOutbound"

	// BastionRole describes the value for the bastion role
	BastionRole = "bastion"

	// CommonRole describes the value for the common role
	CommonRole = "common"

	// VMTagsLastAppliedAnnotation is the key for the machine object annotation
	// which tracks the AdditionalTags in the Machine Provider Config.
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	// for annotation formatting rules.
	VMTagsLastAppliedAnnotation = "sigs.k8s.io/cluster-api-provider-azure-last-applied-tags-vm"
)

var (
	// Use underscore '_' as tags separator because Azure API does not allow both '\' and '/' characters

	// ClusterIDTagKey is the key for resources associated with cluster by unique identifier
	ClusterIDTagKey = fmt.Sprintf("%s_clusterID", TagPrefix)

	// NamespaceTagKey is the key for cluster namespace
	NamespaceTagKey = fmt.Sprintf("%s_namespace", TagPrefix)

	// ClusterNameTagKey is the key for cluster name
	ClusterNameTagKey = fmt.Sprintf("%s_cluster", TagPrefix)

	// MachineRoleTagKey is the key for kubernetes machine role
	MachineRoleTagKey = fmt.Sprintf("%s_role", KubernetesTagPrefix)
)

// BuildParams is used to build tags around an azure resource.
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
