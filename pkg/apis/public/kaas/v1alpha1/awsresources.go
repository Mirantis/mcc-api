/*
Copyright 2020 The Mirantis Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSResources is the Schema for the awsresources API
// +k8s:openapi-gen=true
// +kubebuilder:resource
type AWSResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AWSResourcesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSResourcesList contains a list of AWSResources
type AWSResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSResources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSResources{}, &AWSResourcesList{})
}
