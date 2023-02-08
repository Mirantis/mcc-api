/*
Copyright 2021 The Mirantis Authors.

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
	"k8s.io/apimachinery/pkg/runtime"

	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureClusterProviderSpec is the schema for the azureclusterproviderspec API
// +k8s:openapi-gen=true
type AzureClusterProviderSpec struct {
	metav1.TypeMeta               `json:",inline"`
	metav1.ObjectMeta             `json:"metadata,omitempty"`
	kaasv1alpha1.ClusterSpecMixin `json:",inline"`

	// NetworkSpec encapsulates all things related to Azure network.
	NetworkSpec NetworkSpec `json:"networkSpec,omitempty"`

	// BastionSpec encapsulates all things related to the Bastions in the cluster.
	// +optional
	BastionSpec BastionSpec `json:"bastionSpec,omitempty"`

	// Location defines Azure region where the cluster will be deployed
	Location string `json:"location"`
}

// BastionSpec specifies how the Bastion feature should be set up for the cluster.
type BastionSpec struct {
	// +optional
	AzureBastion *AzureBastion `json:"azureBastion,omitempty"`
}

// AzureBastion specifies how the Azure Bastion cloud component should be configured.
type AzureBastion struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Subnet SubnetSpec `json:"subnet,omitempty"`
	// +optional
	PublicIP PublicIPSpec `json:"publicIP,omitempty"`
}

func (s *AzureClusterProviderSpec) GetClusterSpecMixin() *kaasv1alpha1.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}

func (*AzureClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &AzureClusterProviderStatus{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&AzureClusterProviderSpec{})
}
