package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: "azure.kaas.mirantis.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// +gocode:public-api=true
	ClusterSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "AzureClusterProviderSpec",
	}
)

// +gocode:public-api=true
func EncodeMachineSpec(spec *AzureMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// +gocode:public-api=true
func EncodeClusterSpec(spec *AzureClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(ClusterSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}
