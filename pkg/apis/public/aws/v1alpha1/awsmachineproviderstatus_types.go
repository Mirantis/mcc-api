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

	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field.
// It containsk AWS-specific status information.
// +k8s:openapi-gen=true
type AWSMachineProviderStatus struct {
	metav1.TypeMeta                 `json:",inline"`
	metav1.ObjectMeta               `json:"metadata,omitempty"`
	kaasv1alpha1.MachineStatusMixin `json:",inline"`
	// TODO: remove InstanceID, InstanceState and use the ones from MachineStatusMixin
	// InstanceID is the instance ID of the machine created in AWS
	// +optional
	InstanceID *string `json:"instanceID,omitempty"`

	// InstanceState is the state of the AWS instance for this machine
	// +optional
	InstanceState *InstanceState `json:"instanceState,omitempty"`
}

func (s *AWSMachineProviderStatus) GetMachineStatusMixin() *kaasv1alpha1.MachineStatusMixin {
	return &s.MachineStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&AWSMachineProviderStatus{})
}
