package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// GroupName is the group name used in this package
// +gocode:public-api=true
const GroupName = "storage.kaas.mirantis.com"

var (
	// SchemeGroupVersion is group version used to register storage objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

	// SchemeBuilder initializes a scheme builder
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
	// AddToScheme is a global function that registers this API group & version to a scheme
	// +gocode:public-api=true
	AddToScheme = SchemeBuilder.AddToScheme
)
