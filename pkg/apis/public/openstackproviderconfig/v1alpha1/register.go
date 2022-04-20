/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	"sigs.k8s.io/yaml"

	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
)

const GroupName = "openstackproviderconfig"

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "openstackproviderconfig.k8s.io", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	ClusterSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackClusterProviderSpec",
	}
	ClusterStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackClusterProviderStatus",
	}
	MachineSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackMachineProviderSpec",
	}
	MachineStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "OpenstackMachineProviderStatus",
	}
)

// ClusterConfigFromProviderSpec unmarshals a provider config into an OpenStack Cluster type
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

func EncodeMachineSpec(spec *OpenstackMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(MachineSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

func UpdateMachineSpec(machine *clusterv1.Machine, spec *OpenstackMachineProviderSpec, client crclient.Client) error {
	encodedSpec, err := EncodeMachineSpec(spec)
	if err != nil {
		return fmt.Errorf("failed to encode machine spec for machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	machine.Spec.ProviderSpec.Value = encodedSpec
	if err := client.Update(context.Background(), machine); err != nil {
		return fmt.Errorf("failed to update spec for machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	return nil
}

func EncodeClusterStatus(status *OpenstackClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	status.SetGroupVersionKind(ClusterStatusGVK)

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

func EncodeClusterSpec(spec *OpenstackClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(ClusterSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}
