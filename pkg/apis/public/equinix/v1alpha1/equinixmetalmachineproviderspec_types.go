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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalMachineProviderSpec is the schema for the equinixmetalmachineproviderspec API
// +k8s:openapi-gen=true
type EquinixMetalMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineSpecMixin `json:",inline"`

	// Ceph contains Ceph configuration options for Machine
	Ceph CephMachineConfig `json:"ceph,omitempty"`

	OS          string `json:"OS"`
	MachineType string `json:"machineType"`

	// IPXEUrl can be used to set the pxe boot url when using custom OSes with this provider.
	// Note that OS should also be set to "custom_ipxe" if using this value.
	// +optional
	IPXEUrl string `json:"ipxeURL,omitempty"`

	// HardwareReservationID is the unique device hardware reservation ID or `next-available` to
	// automatically let the EquinixMetal api determine one.
	// +optional
	HardwareReservationID string `json:"hardwareReservationID,omitempty"`

	// ProviderID is the unique identifier as specified by the cloud provider.
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// Tags is an optional set of tags to add to EquinixMetal resources managed by the EquinixMetal provider.
	// +optional
	Tags Tags `json:"tags,omitempty"`
}

type CephMachineConfig struct {
	// ManagerMonitor is an option that reflects that Manager and Monitor
	// roles should be enabled for that Node in CephCluster spec
	ManagerMonitor bool `json:"managerMonitor,omitempty"`
	// Storage is an option that reflects that Storage
	// role should be enabled for that Node in CephCluster spec
	Storage bool `json:"storage,omitempty"`
}

func (s *EquinixMetalMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}

func (*EquinixMetalMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &EquinixMetalMachineProviderStatus{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&EquinixMetalMachineProviderSpec{})
}
