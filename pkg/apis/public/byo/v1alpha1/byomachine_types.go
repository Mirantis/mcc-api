/*
Copyright 2020 The Mirantis Inc.

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

	kaas "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BYOMachineProviderSpec is the type that will be embedded in a Machine.Spec.ProviderSpec field
// for an BYO instance.
// +k8s:openapi-gen=true
type BYOMachineProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.MachineSpecMixin `json:",inline"`
}

func (*BYOMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &BYOMachineProviderStatus{}
}

func (s *BYOMachineProviderSpec) GetMachineSpecMixin() *kaas.MachineSpecMixin {
	return &s.MachineSpecMixin
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BYOMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field.
// It containsk BYO-specific status information.
// +k8s:openapi-gen=true
type BYOMachineProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.MachineStatusMixin `json:",inline"`
}

func (s *BYOMachineProviderStatus) GetMachineStatusMixin() *kaas.MachineStatusMixin {
	return &s.MachineStatusMixin
}

func init() {
	SchemeBuilder.Register(&BYOMachineProviderSpec{})
	SchemeBuilder.Register(&BYOMachineProviderStatus{})
}
