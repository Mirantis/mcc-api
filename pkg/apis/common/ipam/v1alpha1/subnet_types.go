/*
Copyright © 2020 Mirantis

Inspired by https://github.com/inwinstack/ipam/

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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	CIDR          string   `json:"cidr"`
	IncludeRanges []string `json:"includeRanges,omitempty"`
	ExcludeRanges []string `json:"excludeRanges,omitempty"`
	UseWholeCidr  bool     `json:"useWholeCidr,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	Nameservers   []string `json:"nameservers,omitempty"`
	SubnetPoolRef string   `json:"subnetPoolRef,omitempty"`
}

// These are the valid phases of a Subnet.
// The Phase is a start phrase of status message, ex:
//   OK
//   OK, but overfull
//   OK, but intersect with another subnet
//   ERR, shit happens
//   TERM: will be removed soon

const (
	SubnetNone        = ""
	SubnetActive      = "OK"
	SubnetFailed      = "ERR"
	SubnetTerminating = "TERM"
)

type AllocatedIPs []string

// SubnetStatus defines the observed state of Subnet
type SubnetStatus struct {
	CIDR             string       `json:"cidr"`
	Gateway          string       `json:"gateway,omitempty"`
	Nameservers      []string     `json:"nameservers,omitempty"`
	Ranges           []string     `json:"ranges"`
	StatusMessage    string       `json:"statusMessage"`
	AllocatedIPs     AllocatedIPs `json:"allocatedIPs"`
	Capacity         int          `json:"capacity"`
	Allocatable      int          `json:"allocatable"`
	SubnetPoolRef    string       `json:"subnetPoolRef,omitempty"`
	ObjCreated       string       `json:"objCreated,omitempty"`
	ObjUpdated       string       `json:"objUpdated,omitempty"`
	ObjStatusUpdated string       `json:"objStatusUpdated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=subnets,scope=Namespaced
// Subnet is the Schema for the subnets API
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubnetSpec   `json:"spec,omitempty"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// SubnetList contains a list of Subnet
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

var (
	SubnetKind = "Subnet"
	SubnetGVK  = schema.GroupVersionKind{
		Kind:    SubnetKind,
		Group:   ResourceGroup,
		Version: Version,
	}
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
