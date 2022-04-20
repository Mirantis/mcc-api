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

	l2tmplTypes "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam/l2template/types"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// L2TemplateList is a list of IP.
type L2TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []L2Template `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=l2templates,scope=Namespaced
// L2Template is the Schema for the l2templates API
type L2Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   L2TemplateSpec   `json:"spec"`
	Status L2TemplateStatus `json:"status,omitempty"`
}

// L2TemplateSpec is the spec for a IP resource.
type L2TemplateSpec struct {
	ClusterRef        string                 `json:"clusterRef,omitempty"`
	NpTemplate        l2tmplTypes.NpTemplate `json:"npTemplate"`
	L3Layout          l2tmplTypes.L3Layout   `json:"l3Layout"`
	IfMapping         l2tmplTypes.IfMapping  `json:"ifMapping,omitempty"`
	AutoIfMappingPrio l2tmplTypes.IfMapping  `json:"autoIfMappingPrio,omitempty"`
}

// These are the valid phases of a L2Template
const (
	L2TemplateNone        = ""
	L2TemplateReady       = "Ready"
	L2TemplateFailed      = "Failed"
	L2TemplateTerminating = "Terminating"
)

// L2TemplateStatus represents the current state of a L2Template resource.
type L2TemplateStatus struct {
	Phase            string `json:"phase"`
	Reason           string `json:"reason,omitempty"`
	ObjCreated       string `json:"objCreated,omitempty"`
	ObjUpdated       string `json:"objUpdated,omitempty"`
	ObjStatusUpdated string `json:"objStatusUpdated,omitempty"`
}

var (
	L2TemplateKind = "L2Template"
	L2TemplateGVK  = schema.GroupVersionKind{
		Kind:    L2TemplateKind,
		Group:   ResourceGroup,
		Version: Version,
	}
)

func init() {
	SchemeBuilder.Register(&L2Template{}, &L2TemplateList{})
}
