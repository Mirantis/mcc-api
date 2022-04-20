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
)

// AzureResourcesStatus defines the observed state of AzureResources
type AzureResourcesStatus struct {
	AzureLocations map[string]AzureLocation `json:"locations"`
}

type AzureLocation struct {
	VMSizes []VMSize `json:"vmSizes,omitempty"`
}

type VMSize struct {
	Name                     string `json:"name"`
	CPU                      int    `json:"cpu"`
	RAM                      int    `json:"ram"`
	CachedDiskSize           int    `json:"cachedDiskSize"`
	EphemeralOSDiskSupported bool   `json:"ephemeralOSDiskSupported"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureResources is the Schema for the azureresources API
// +k8s:openapi-gen=true
// +kubebuilder:resource
type AzureResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AzureResourcesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureResourcesList contains a list of AzureResources
type AzureResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureResources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureResources{}, &AzureResourcesList{})
}
