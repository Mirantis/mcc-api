package tags

import (
	"path"
	"reflect"
)

const (
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

	// NameKubernetesClusterPrefix is the tag name we use to differentiate multiple
	// logically independent clusters running in the same AZ.
	// The tag key = NameKubernetesClusterPrefix + clusterID
	// The tag value is an ownership value
	// +gocode:public-api=true
	NameKubernetesClusterPrefix = "kubernetes.io/cluster/"

	// NameKaaSCluster is the human readable cluster name, showed in KaaS
	// +gocode:public-api=true
	NameKaaSCluster = "cluster-name"

	// NameAWSProviderPrefix is the tag prefix we use to differentiate
	// cluster-api-provider-aws owned components from other tooling that
	// uses NameKubernetesClusterPrefix
	// +gocode:public-api=true
	NameAWSProviderPrefix = "sigs.k8s.io/cluster-api-provider-aws/"

	// NameAWSProviderManaged is the tag name we use to differentiate
	// cluster-api-provider-aws owned components from other tooling that
	// uses NameKubernetesClusterPrefix
	// +gocode:public-api=true
	NameAWSProviderManaged = NameAWSProviderPrefix + "managed"

	// NameAWSClusterAPIRole is the tag name we use to mark roles for resources
	// dedicated to this cluster api provider implementation.
	// +gocode:public-api=true
	NameAWSClusterAPIRole = NameAWSProviderPrefix + "role"

	// ValueAPIServerRole describes the value for the apiserver role
	// +gocode:public-api=true
	ValueAPIServerRole = "apiserver"

	// ValueBastionRole describes the value for the bastion role
	// +gocode:public-api=true
	ValueBastionRole = "bastion"

	// ValueCommonRole describes the value for the common role
	// +gocode:public-api=true
	ValueCommonRole = "common"

	// ValuePublicRole describes the value for the public role
	// +gocode:public-api=true
	ValuePublicRole = "public"

	// ValuePrivateRole describes the value for the private role
	// +gocode:public-api=true
	ValuePrivateRole = "private"
)

// ResourceLifecycle configures the lifecycle of a resource
// +gocode:public-api=true
type ResourceLifecycle string

// Map defines a map of tags.
// +gocode:public-api=true
type Map map[string]string

// Equals returns true if the maps are equal.
func (m Map) Equals(other Map) bool {
	return reflect.DeepEqual(m, other)
}

// HasOwned returns true if the tags contains a tag that marks the resource as owned by the cluster.
func (m Map) HasOwned(cluster string) bool {
	value, ok := m[path.Join(NameKubernetesClusterPrefix, cluster)]
	return ok && ResourceLifecycle(value) == ResourceLifecycleOwned
}

// HasManaged returns true if the map contains NameAWSProviderManaged key set to true.
func (m Map) HasManaged() bool {
	value, ok := m[NameAWSProviderManaged]
	return ok && value == "true"
}

// GetRole returns the Cluster API role for the tagged resource
func (m Map) GetRole() string {
	return m[NameAWSClusterAPIRole]
}

// Difference returns the difference between this map and the other map.
// Items are considered equals if key and value are equals.
func (m Map) Difference(other Map) Map {
	res := make(Map, len(m))

	for key, value := range m {
		if otherValue, ok := other[key]; ok && value == otherValue {
			continue
		}
		res[key] = value
	}

	return res
}
