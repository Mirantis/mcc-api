package v1alpha1

import (
	"context"
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/vsphere"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: vsphere.GroupName, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// +gocode:public-api=true
	ClusterSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereClusterProviderSpec",
	}
	// +gocode:public-api=true
	ClusterStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereClusterProviderStatus",
	}
	// +gocode:public-api=true
	MachineSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereMachineProviderSpec",
	}
	// +gocode:public-api=true
	MachineStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereMachineProviderStatus",
	}
)

// Resource is required by pkg/client/listers/...
// +gocode:public-api=true
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +gocode:public-api=true
func EncodeClusterSpec(spec *VsphereClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(ClusterSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// +gocode:public-api=true
func EncodeClusterStatus(status *VsphereClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	status.SetGroupVersionKind(ClusterStatusGVK)

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

// +gocode:public-api=true
func EncodeMachineSpec(spec *VsphereMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(MachineSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// +gocode:public-api=true
func EncodeMachineStatus(status *VsphereMachineProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	status.SetGroupVersionKind(MachineStatusGVK)

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

// +gocode:public-api=true
func UpdateMachineStatus(ctx context.Context, machine *clusterv1.Machine, status *VsphereMachineProviderStatus, vmName string, client crclient.Client) error {
	status.SetGroupVersionKind(MachineStatusGVK)
	encodedStatus, err := EncodeMachineStatus(status)
	if err != nil {
		return errors.Errorf("failed to encode machine status for the machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	machine.Status.ProviderStatus = encodedStatus
	machine.Status.InstanceName = vmName
	if err := client.Status().Update(ctx, machine); err != nil {
		return errors.Errorf("failed to update status for the machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	return nil
}
