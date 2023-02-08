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

// Package v1alpha1 contains API Schema definitions for the vsphere v1alpha1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/Mirantis/mcc-api/v2/pkg/apis/public/vsphere
// +k8s:defaulter-gen=TypeMeta
// +groupName=vsphere.cluster.sigs.k8s.io
package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/public/vsphere"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: vsphere.GroupName, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	ClusterSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereClusterProviderSpec",
	}
	ClusterStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereClusterProviderStatus",
	}
	MachineSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereMachineProviderSpec",
	}
	MachineStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "VsphereMachineProviderStatus",
	}
)

// Resource is required by pkg/client/listers/...
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func EncodeClusterSpec(spec *VsphereClusterProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(ClusterSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

func EncodeClusterStatus(status *VsphereClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	status.SetGroupVersionKind(ClusterStatusGVK)

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

func EncodeMachineSpec(spec *VsphereMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(MachineSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}

func EncodeMachineStatus(status *VsphereMachineProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	status.SetGroupVersionKind(MachineStatusGVK)

	return &runtime.RawExtension{
		Object: status,
	}, nil
}

func UpdateMachineStatus(machine *clusterv1.Machine, status *VsphereMachineProviderStatus, vmName string, client crclient.Client) error {
	status.SetGroupVersionKind(MachineStatusGVK)
	encodedStatus, err := EncodeMachineStatus(status)
	if err != nil {
		return errors.Errorf("failed to encode machine status for the machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	machine.Status.ProviderStatus = encodedStatus
	machine.Status.InstanceName = vmName
	if err := client.Status().Update(context.Background(), machine); err != nil {
		return errors.Errorf("failed to update status for the machine %v/%v: %v",
			machine.Name, machine.Namespace, err)
	}
	return nil
}
