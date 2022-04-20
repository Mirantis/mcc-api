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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// IPaddrList is a list of IPaddr
type IPaddrList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []IPaddr `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ipaddrs,scope=Namespaced
// IPaddr is the Schema for the ipaddrs API
type IPaddr struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   IPaddrSpec   `json:"spec"`
	Status IPaddrStatus `json:"status,omitempty"`
}

// IPaddrSpec is the spec for a IP resource.
type IPaddrSpec struct {
	SubnetRef     string `json:"subnetRef"`
	WantedAddress string `json:"wantedAddress,omitempty"`
	MAC           string `json:"mac,omitempty"`
}

// IPaddrPhase -- type of phase of resource
type IPaddrPhase string

// These are the valid phases of a IP.
const (
	IPaddrNone        IPaddrPhase = ""
	IPaddrActive      IPaddrPhase = "Active"
	IPaddrFailed      IPaddrPhase = "Failed"
	IPaddrTerminating IPaddrPhase = "Terminating"
)

// IPaddrStatus represents the current state of a IP resource.
type IPaddrStatus struct {
	Phase            IPaddrPhase `json:"phase"`
	Reason           string      `json:"reason,omitempty"`
	Address          string      `json:"address"`
	CIDR             string      `json:"cidr"`
	Gateway          string      `json:"gateway,omitempty"`
	Nameservers      []string    `json:"nameservers,omitempty"`
	MAC              string      `json:"mac"`
	Ports            []int       `json:"ports,omitempty"`
	ObjCreated       string      `json:"objCreated,omitempty"`
	ObjUpdated       string      `json:"objUpdated,omitempty"`
	ObjStatusUpdated string      `json:"objStatusUpdated,omitempty"`
}

var (
	IPaddrKind = "IPaddr"
	IPaddrGVK  = schema.GroupVersionKind{
		Kind:    IPaddrKind,
		Group:   ResourceGroup,
		Version: Version,
	}
)

func init() {
	SchemeBuilder.Register(&IPaddr{}, &IPaddrList{})
}
