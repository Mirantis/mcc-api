/*
Copyright 2019 The Kubernetes Authors.

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

	kaas "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalMachineProviderStatus is the Schema for the baremetalmachineproviderstatuses API
// +k8s:openapi-gen=true
type BareMetalMachineProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.MachineStatusMixin `json:",inline"`
}

func (s *BareMetalMachineProviderStatus) GetMachineStatusMixin() *kaas.MachineStatusMixin {
	return &s.MachineStatusMixin
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalMachineProviderStatusList contains a list of BareMetalMachineProviderStatus
// +k8s:openapi-gen=true
type BareMetalMachineProviderStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BareMetalMachineProviderStatus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BareMetalMachineProviderStatus{}, &BareMetalMachineProviderStatusList{})
}
