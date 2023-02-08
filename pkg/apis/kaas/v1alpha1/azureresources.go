package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VMSize struct {
	Name                     string `json:"name"`
	CPU                      int    `json:"cpu"`
	RAM                      int    `json:"ram"`
	CachedDiskSize           int    `json:"cachedDiskSize"`
	EphemeralOSDiskSupported bool   `json:"ephemeralOSDiskSupported"`
}

// AzureResourcesList contains a list of AzureResources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type AzureResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureResources `json:"items"`
}

// AzureResourcesStatus defines the observed state of AzureResources
type AzureResourcesStatus struct {
	AzureLocations map[string]AzureLocation `json:"locations"`
}
type AzureLocation struct {
	VMSizes []VMSize `json:"vmSizes,omitempty"`
}

// AzureResources is the Schema for the azureresources API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type AzureResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AzureResourcesStatus `json:"status,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AzureResources{}, &AzureResourcesList{})
}
