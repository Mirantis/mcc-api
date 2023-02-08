/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Mirantis/mcc-api/v2/pkg/apis/public/equinix/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EquinixMetalClusterProviderStatus is the schema for the equinixmetalclusterproviderstatus API
// +k8s:openapi-gen=true
type EquinixMetalClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.ClusterStatusMixin `json:",inline"`

	// Elastic IP blocks that were requested for cluster
	ElasticIPBlocks []v1alpha1.ElasticIPReservation `json:"elasticIPBlocks"`

	// MetalLB IP ranges that were allocated for cluster
	MetalLBRanges []string `json:"metalLBRanges,omitempty"`

	// VlanStatus describes resources created for managing requested VLANs
	VlanStatus map[string]VlanStatus `json:"vlans"`
}

type VlanStatus struct {
	// SubnetID refers to a subnet created for managing the VLAN
	SubnetID string `json:"subnetID,omitempty"`
}

func (s *EquinixMetalClusterProviderStatus) GetClusterStatusMixin() *kaasv1alpha1.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderStatus{})
}
