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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
	ElasticIPBlocks []ElasticIPReservation `json:"elasticIPBlocks"`

	// MetalLB IP ranges that were allocated for cluster
	MetalLBRanges []string `json:"metalLBRanges,omitempty"`
}

type ElasticIPReservation struct {
	// purpose of public address block, see ElasticIPPurpose
	Purpose ElasticIPPurpose `json:"purpose"`
	// number of IPs in a block, can be 1, 2, 4, 8, 16
	Quantity int `json:"quantity"`
	// subnet address in CIDR notation, read-only, it's empty when block is not allocated yet
	CIDR string `json:"cidr,omitempty"`
	// Equinix resource ID, read-only, it's empty when block is not allocated yet
	ID string `json:"id"`
	// Equinix elastic IP reservation state, can be "created" or "pending"
	State string `json:"state"`
}

func (s *EquinixMetalClusterProviderStatus) GetClusterStatusMixin() *kaasv1alpha1.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&EquinixMetalClusterProviderStatus{})
}
