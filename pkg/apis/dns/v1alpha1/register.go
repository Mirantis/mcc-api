package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
	// +gocode:public-api=true
	AddToScheme = SchemeBuilder.AddToScheme
)

// +gocode:public-api=true
const GroupName = "dns.mirantis.com"
