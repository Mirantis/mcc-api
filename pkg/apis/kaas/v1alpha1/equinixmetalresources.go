package v1alpha1

import (
	"github.com/packethost/packngo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MachineType struct {
	Name     string            `json:"name"`
	VCPUs    []*packngo.Cpus   `json:"vCPUs"`
	Memory   string            `json:"memory"`
	Drives   []*packngo.Drives `json:"drives"`
	Capacity string            `json:"capacity"`
}
type EquinixMetalFacility struct {
	MachineTypes map[string]MachineType `json:"machineTypes,omitempty"`
}
type OperatingSystem struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Distro  string `json:"distro"`
	Version string `json:"version"`
}

// EquinixMetalResources is the Schema for the equinixmetalresources API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type EquinixMetalResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status EquinixMetalResourcesStatus `json:"status,omitempty"`
}

// EquinixMetalResourcesList contains a list of EquinixMetalResources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type EquinixMetalResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EquinixMetalResources `json:"items"`
}

// EquinixMetalResourcesStatus defines the observed state of EquinixMetalResources
type EquinixMetalResourcesStatus struct {
	EquinixMetalFacilities map[string]EquinixMetalFacility `json:"facilities"`
	OperatingSystems       []OperatingSystem               `json:"operatingSystems"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalResources{}, &EquinixMetalResourcesList{})
}
