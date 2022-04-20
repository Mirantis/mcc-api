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

	kaasIpam "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam"
	"github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam/cloudinit"
	l2tmplTypes "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam/l2template/types"
	kiTypes "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam/types"
	k8sutilTypes "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

// IpamHostSpec defines the desired state of IpamHost
type IpamHostSpec struct {
	NicMacMap          kaasIpam.NicMacMap           `json:"nicMACmap,omitempty"`
	AllocationStatus   string                       `json:"allocationStatus,omitempty"` // empty if IpamHost should be created in non-Invalid state
	L2TemplateSelector *kaasIpam.L2TemplateSelector `json:"l2TemplateSelector,omitempty"`
	NpTemplateOverride l2tmplTypes.NpTemplate       `json:"npTemplateOverride,omitempty"`
	IfMappingOverride  l2tmplTypes.IfMapping        `json:"ifMappingOverride,omitempty"`
}

// IpamHostStatus defines the observed state of IpamHost
type IpamHostStatus struct {
	NicMacMap           kaasIpam.NicMacMap                 `json:"nicMACmap,omitempty"`
	NetconfigV2         *cloudinit.UserDataNetworkV2       `json:"netconfigV2"`
	NetconfigFiles      cloudinit.NetconfigFiles           `json:"netconfigFiles"`
	NetconfigFilesState string                             `json:"netconfigFilesState"`
	ServiceMap          kiTypes.Services2IfAddressPlan     `json:"serviceMap,omitempty"`
	OSmetadataNetwork   *cloudinit.OSmetadataNetworkConfig `json:"osMetadataNetwork,omitempty"`
	AllocationResult    string                             `json:"ipAllocationResult,omitempty"`
	ErrorMessage        string                             `json:"errorMessage,omitempty"`
	L2TemplateRef       string                             `json:"l2TemplateRef,omitempty"`
	L2RenderResult      string                             `json:"l2RenderResult,omitempty"`
	Warnings            k8sutilTypes.WarningsList          `json:"Warnings,omitempty"`
	ObjCreated          string                             `json:"objCreated,omitempty"`
	ObjUpdated          string                             `json:"objUpdated,omitempty"`
	ObjStatusUpdated    string                             `json:"objStatusUpdated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ipamhosts,scope=Namespaced
// IpamHost is the Schema for the ipamhosts API
type IpamHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IpamHostSpec   `json:"spec,omitempty"`
	Status IpamHostStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// IpamHostList contains a list of IpamHost
type IpamHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpamHost `json:"items"`
}

var (
	IpamHostKind = "IpamHost"
	IpamHostGVK  = schema.GroupVersionKind{
		Kind:    IpamHostKind,
		Group:   ResourceGroup,
		Version: Version,
	}
)

func init() {
	SchemeBuilder.Register(&IpamHost{}, &IpamHostList{})
}
