package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenStackResourcesList contains a list of OpenStackResources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type OpenStackResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenStackResources `json:"items"`
}

// OpenStackResourcesStatus defines the observed state of OpenStackResources
// +gocode:public-api=true
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
type ExternalNetwork struct {
	Name string `json:"Name"`
	ID   string `json:"ID"`
}
type ComputeAZ struct {
	Name string `json:"Name"`
}
type Flavor struct {
	Name      string `json:"Name"`
	ID        string `json:"ID"`
	RAM       string `json:"RAM"`
	Disk      string `json:"Disk"`
	Ephemeral string `json:"Ephemeral"`
	VCPUs     string `json:"VCPUs"`
}

// OpenStackResources is the Schema for the openstackresources API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type OpenStackResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status OpenStackResourcesStatus `json:"status,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&OpenStackResources{}, &OpenStackResourcesList{})
}
