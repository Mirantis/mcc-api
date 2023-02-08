package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: "equinix.kaas.mirantis.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// +gocode:public-api=true
func EncodeClusterSpec(spec *EquinixMetalClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// +gocode:public-api=true
func EncodeClusterStatus(status *EquinixMetalClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

// +gocode:public-api=true
func EncodeMachineSpec(spec *EquinixMetalMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// +gocode:public-api=true
func EncodeMachineStatus(status *EquinixMetalMachineProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: status,
	}, nil
}
