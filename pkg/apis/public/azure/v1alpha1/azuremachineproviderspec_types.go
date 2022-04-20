/*
Copyright 2021 The Mirantis Authors.

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

// AzureMachineProviderSpec is the schema for the azuremachineproviderspec API
// +k8s:openapi-gen=true
type AzureMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineSpecMixin `json:",inline"`

	VMSize string `json:"vmSize"`

	// Image is used to provide details of an image to use during VM creation.
	// If image details are omitted the image will default the Azure Marketplace "capi" offer,
	// which is based on Ubuntu.
	// +kubebuilder:validation:nullable
	// +optional
	Image *Image `json:"image,omitempty"`

	// OSDisk specifies the parameters for the operating system disk of the machine
	OSDisk OSDisk `json:"osDisk"`
}

func (s *AzureMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}

func (*AzureMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &AzureMachineProviderStatus{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&AzureMachineProviderSpec{})
}
