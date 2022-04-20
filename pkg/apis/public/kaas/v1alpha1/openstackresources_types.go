/*
Copyright 2019 The Mirantis Authors.

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

// OpenStackResourcesStatus defines the observed state of OpenStackResources
type OpenStackResourcesStatus struct {
	Images           []*Image           `json:"images,omitempty"`
	Flavors          []*Flavor          `json:"flavors,omitempty"`
	ExternalNetworks []*ExternalNetwork `json:"externalNetworks,omitempty"`
	ComputeAZs       []*ComputeAZ       `json:"computeAZ,omitempty"`
}

type Image struct {
	Name string `json:"Name"`
	ID   string `json:"ID"`
}

type Flavor struct {
	Name      string `json:"Name"`
	ID        string `json:"ID"`
	RAM       string `json:"RAM"`
	Disk      string `json:"Disk"`
	Ephemeral string `json:"Ephemeral"`
	VCPUs     string `json:"VCPUs"`
}

type ExternalNetwork struct {
	Name string `json:"Name"`
	ID   string `json:"ID"`
}

type ComputeAZ struct {
	Name string `json:"Name"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenStackResources is the Schema for the openstackresources API
// +k8s:openapi-gen=true
// +kubebuilder:resource
type OpenStackResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status OpenStackResourcesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenStackResourcesList contains a list of OpenStackResources
type OpenStackResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenStackResources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenStackResources{}, &OpenStackResourcesList{})
}
