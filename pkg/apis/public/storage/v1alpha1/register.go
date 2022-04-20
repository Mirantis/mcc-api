// Package v1alpha1 contains API Schema definitions for the storage v1alpha1 API group
// +k8s:deepcopy-gen=package
// +groupName=storage.kaas.mirantis.com
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// GroupName is the group name used in this package
const GroupName = "storage.kaas.mirantis.com"

var (
	// SchemeGroupVersion is group version used to register storage objects
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

	// SchemeBuilder initializes a scheme builder
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)
