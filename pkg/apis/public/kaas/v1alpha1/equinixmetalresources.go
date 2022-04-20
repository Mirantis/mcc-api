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
	"github.com/packethost/packngo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EquinixMetalResourcesStatus defines the observed state of EquinixMetalResources
type EquinixMetalResourcesStatus struct {
	EquinixMetalFacilities map[string]EquinixMetalFacility `json:"facilities"`
	OperatingSystems       []OperatingSystem               `json:"operatingSystems"`
}

type EquinixMetalFacility struct {
	MachineTypes map[string]MachineType `json:"machineTypes,omitempty"`
}

type MachineType struct {
	Name     string            `json:"name"`
	VCPUs    []*packngo.Cpus   `json:"vCPUs"`
	Memory   string            `json:"memory"`
	Drives   []*packngo.Drives `json:"drives"`
	Price    string            `json:"price"`
	Capacity string            `json:"capacity"`
}

type OperatingSystem struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Distro  string `json:"distro"`
	Version string `json:"version"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalResources is the Schema for the equinixmetalresources API
// +k8s:openapi-gen=true
// +kubebuilder:resource
type EquinixMetalResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status EquinixMetalResourcesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalResourcesList contains a list of EquinixMetalResources
type EquinixMetalResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EquinixMetalResources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EquinixMetalResources{}, &EquinixMetalResourcesList{})
}
