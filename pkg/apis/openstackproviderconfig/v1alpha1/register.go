package v1alpha1

import (
	"context"
	"fmt"
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	"sigs.k8s.io/yaml"
)

// +gocode:public-api=true
const GroupName = "openstackproviderconfig"

var (
	// SchemeGroupVersion is group version used to register these objects
	// +gocode:public-api=true
	SchemeGroupVersion = schema.GroupVersion{Group: "openstackproviderconfig.k8s.io", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// +gocode:public-api=true
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// +gocode:public-api=true
	ClusterSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackClusterProviderSpec",
	}
	// +gocode:public-api=true
	ClusterStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackClusterProviderStatus",
	}
	// +gocode:public-api=true
	MachineSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackMachineProviderSpec",
	}
	// +gocode:public-api=true
	MachineStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackMachineProviderStatus",
	}
)

// ClusterConfigFromProviderSpec unmarshals a provider config into an OpenStack Cluster type
// +gocode:public-api=true
func ClusterSpecFromProviderSpec(providerSpec clusterv1.ProviderSpec) (*OpenstackClusterProviderSpec, error) {
	if providerSpec.Value == nil {
		return nil, errors.New("no such providerSpec found in manifest")
	}

	var config OpenstackClusterProviderSpec
	if err := yaml.Unmarshal(providerSpec.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// ClusterStatusFromProviderStatus unmarshals a provider status into an OpenStack Cluster Status type
// +gocode:public-api=true
func ClusterStatusFromProviderStatus(extension *runtime.RawExtension) (*OpenstackClusterProviderStatus, error) {
	if extension == nil {
		return &OpenstackClusterProviderStatus{}, nil
	}

	status := new(OpenstackClusterProviderStatus)
	if err := yaml.Unmarshal(extension.Raw, status); err != nil {
		return nil, err
	}

	return status, nil
}

// This is the same as ClusterSpecFromProviderSpec but we
// expect there to be a specific Spec type for Machines soon
// +gocode:public-api=true
func MachineSpecFromProviderSpec(providerSpec clusterv1.ProviderSpec) (*OpenstackMachineProviderSpec, error) {
	if providerSpec.Value == nil {
		return nil, errors.New("no such providerSpec found in manifest")
	}

	var config OpenstackMachineProviderSpec
	if err := yaml.Unmarshal(providerSpec.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// MachineStatusFromProviderStatus unmarshals a provider status into an OpenStack Machine Status type
// +gocode:public-api=true
func MachineStatusFromProviderStatus(extension *runtime.RawExtension) (*OpenstackMachineProviderStatus, error) {
	if extension == nil {
		return &OpenstackMachineProviderStatus{}, nil
	}

	status := new(OpenstackMachineProviderStatus)
	if err := yaml.Unmarshal(extension.Raw, status); err != nil {
		return nil, err
	}

	return status, nil
}

// +gocode:public-api=true
func EncodeMachineSpec(spec *OpenstackMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(MachineSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

// +gocode:public-api=true
func UpdateMachineSpec(ctx context.Context, machine *clusterv1.Machine, spec *OpenstackMachineProviderSpec, client crclient.Client) error {
	encodedSpec, err := EncodeMachineSpec(spec)
	if err != nil {
		return fmt.Errorf("failed to encode machine spec for machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	machine.Spec.ProviderSpec.Value = encodedSpec
	if err := client.Update(ctx, machine); err != nil {
		return fmt.Errorf("failed to update spec for machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	return nil
}

// +gocode:public-api=true
func EncodeClusterStatus(status *OpenstackClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	status.SetGroupVersionKind(ClusterStatusGVK)

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

// +gocode:public-api=true
func EncodeClusterSpec(spec *OpenstackClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(ClusterSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}
