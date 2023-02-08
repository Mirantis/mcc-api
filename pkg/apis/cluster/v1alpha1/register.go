package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects.
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: "cluster.k8s.io", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds registered types to the builder.
	// Required by pkg/client/...
	// TODO(pwittrock): Remove this after removing pkg/client/...
	// +gocode:public-api=true
	AddToScheme = SchemeBuilder.AddToScheme
)

// Required by pkg/client/listers/...
// TODO(pwittrock): Remove this after removing pkg/client/...
// +gocode:public-api=true
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
