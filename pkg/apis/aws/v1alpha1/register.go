package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: "aws.kaas.mirantis.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// EncodeMachineStatus marshals the machine status
// +gocode:public-api=true
func EncodeMachineStatus(status *AWSMachineProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

// EncodeMachineSpec marshals the machine provider spec.
// +gocode:public-api=true
func EncodeMachineSpec(spec *AWSMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// EncodeClusterStatus marshals the cluster status.
// +gocode:public-api=true
func EncodeClusterStatus(status *AWSClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

// EncodeClusterSpec marshals the cluster provider spec.
// +gocode:public-api=true
func EncodeClusterSpec(spec *AWSClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}
