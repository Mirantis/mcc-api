package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSResourcesStatus defines the observed state of OpenStackResources
type AWSResourcesStatus struct {
	AWSRegions map[string]AWSRegion `json:"regions"`
}
type AWSRegion struct {
	InstanceTypes map[string]InstanceType `json:"instanceTypes"`
	AMIs          map[string]AMI          `json:"AMIs"`
}
type InstanceType struct {
	VCPUs   string `json:"vCPUs"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
}
type AMI struct {
	Name               string `json:"name"`
	Architecture       string `json:"architecture"`
	Platform           string `json:"platform"`
	VirtualizationType string `json:"virtualizationType"`
	OwnerID            string `json:"ownerID"`
	Public             bool   `json:"public"`
}

// AWSResources is the Schema for the awsresources API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type AWSResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AWSResourcesStatus `json:"status,omitempty"`
}

// AWSResourcesList contains a list of AWSResources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type AWSResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSResources `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AWSResources{}, &AWSResourcesList{})
}
