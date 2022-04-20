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

// VsphereResourcesStatus defines the observed state of VsphereResources
type VsphereResourcesStatus struct {
	MCCUser           MCCUser           `json:"mccUser,omitempty"`
	CloudProviderUser CloudProviderUser `json:"cloudProviderUser,omitempty"`
}

type MCCUser struct {
	Datastores       []Datastore       `json:"datastores,omitempty"`
	DatastoreFolders []DatastoreFolder `json:"datastoreFolders,omitempty"`
	Networks         []Network         `json:"networks,omitempty"`
	ResourcePools    []ResourcePool    `json:"resourcePools,omitempty"`
	MachineFolders   []MachineFolder   `json:"machineFolders,omitempty"`
	MachineTemplates []MachineTemplate `json:"machineTemplates,omitempty"`
}

type CloudProviderUser struct {
	Datastores []Datastore `json:"datastores,omitempty"`
}

type NamePathPair struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Datastore struct {
	NamePathPair `json:",inline"`
}

type DatastoreFolder struct {
	NamePathPair `json:",inline"`
}

type Network struct {
	NamePathPair `json:",inline"`
	Type         string `json:"type"`
}

type ResourcePool struct {
	NamePathPair `json:",inline"`
}

type MachineFolder struct {
	NamePathPair `json:",inline"`
}

type MachineTemplate struct {
	NamePathPair `json:",inline"`
	MccTemplate  string `json:"mccTemplate"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VsphereResources is the Schema for the vsphereresources API
// +k8s:openapi-gen=true
type VsphereResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status VsphereResourcesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VsphereResourcesList contains a list of VsphereResources
type VsphereResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VsphereResources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VsphereResources{}, &VsphereResourcesList{})
}
