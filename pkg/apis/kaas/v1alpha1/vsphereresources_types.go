package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MachineFolder struct {
	NamePathPair `json:",inline"`
}
type NamePathPair struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// VsphereResourcesList contains a list of VsphereResources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type VsphereResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VsphereResources `json:"items"`
}
type DatastoreFolder struct {
	NamePathPair `json:",inline"`
}
type Network struct {
	NamePathPair `json:",inline"`
	Type         string `json:"type"`
}
type MachineTemplate struct {
	NamePathPair `json:",inline"`
	MccTemplate  string `json:"mccTemplate"`
}

// VsphereResources is the Schema for the vsphereresources API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type VsphereResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status VsphereResourcesStatus `json:"status,omitempty"`
}

// VsphereResourcesStatus defines the observed state of VsphereResources
type VsphereResourcesStatus struct {
	MCCUser           MCCUser           `json:"mccUser,omitempty"`
	CloudProviderUser CloudProviderUser `json:"cloudProviderUser,omitempty"`
}
type ResourcePool struct {
	NamePathPair `json:",inline"`
}
type MCCUser struct {
	Datastores       []Datastore       `json:"datastores,omitempty"`
	DatastoreFolders []DatastoreFolder `json:"datastoreFolders,omitempty"`
	Networks         []Network         `json:"networks,omitempty"`
	ResourcePools    []ResourcePool    `json:"resourcePools,omitempty"`
	MachineFolders   []MachineFolder   `json:"machineFolders,omitempty"`
	MachineTemplates []MachineTemplate `json:"machineTemplates,omitempty"`
}
type Datastore struct {
	NamePathPair `json:",inline"`
	ISOFilePaths []string `json:"isoFilePaths,omitempty"`
}
type CloudProviderUser struct {
	Datastores []Datastore `json:"datastores,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&VsphereResources{}, &VsphereResourcesList{})
}
