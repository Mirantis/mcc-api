/*
Copyright © 2020 Mirantis

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

// SubnetPoolSpec defines the desired state of SubnetPool
type SubnetPoolSpec struct {
	CIDR          string   `json:"cidr"`
	BlockSize     string   `json:"blockSize"`
	GatewayPolicy string   `json:"gatewayPolicy,omitempty"`
	Nameservers   []string `json:"nameservers,omitempty"`
}

// These are the valid phases of a SubnetPool.
// The Phase is a start phrase of status message, ex:
//   OK
//   OK, but overfull
//   OK, but intersect with another pool
//   ERR, shit happens
//   TERM: will be removed soon

const (
	SubnetPoolNone        = ""
	SubnetPoolActive      = "OK"
	SubnetPoolFailed      = "ERR"
	SubnetPoolTerminating = "TERM"
)

type AllocatedSubnets []string

// SubnetPoolStatus defines the observed state of SubnetPool
type SubnetPoolStatus struct {
	StatusMessage    string           `json:"statusMessage"`
	AllocatedSubnets AllocatedSubnets `json:"allocatedSubnets"`
	BlockSize        string           `json:"blockSize"` // block size is unchangebly !!!
	Capacity         int              `json:"capacity"`
	Allocatable      int              `json:"allocatable"`
	ObjCreated       string           `json:"objCreated,omitempty"`
	ObjUpdated       string           `json:"objUpdated,omitempty"`
	ObjStatusUpdated string           `json:"objStatusUpdated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=subnetpools,scope=Namespaced
// SubnetPool is the Schema for the subnetpools API
type SubnetPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubnetPoolSpec   `json:"spec,omitempty"`
	Status SubnetPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// SubnetPoolList contains a list of SubnetPool
type SubnetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetPool `json:"items"`
}

var (
	SubnetPoolKind = "SubnetPool"
	SubnetPoolGVK  = schema.GroupVersionKind{
		Kind:    SubnetPoolKind,
		Group:   ResourceGroup,
		Version: Version,
	}
)

func init() {
	SchemeBuilder.Register(&SubnetPool{}, &SubnetPoolList{})
}
