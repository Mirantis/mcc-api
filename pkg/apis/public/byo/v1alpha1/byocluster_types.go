/*
Copyright 2020 The Mirantis Inc.

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

	kaas "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BYOClusterProviderSpec is the providerConfig for BYO in the cluster
// object
// +k8s:openapi-gen=true
type BYOClusterProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.ClusterSpecMixin `json:",inline"`
}

func (s *BYOClusterProviderSpec) GetClusterSpecMixin() *kaas.ClusterSpecMixin {
	return &s.ClusterSpecMixin
}

func (*BYOClusterProviderSpec) GetNewClusterStatus() runtime.Object {
	return &BYOClusterProviderStatus{}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BYOClusterProviderStatus contains the status fields
// relevant to BYO in the cluster object.
// +k8s:openapi-gen=true
type BYOClusterProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.ClusterStatusMixin `json:",inline"`
}

func (s *BYOClusterProviderStatus) GetClusterStatusMixin() *kaas.ClusterStatusMixin {
	return &s.ClusterStatusMixin
}

func init() {
	SchemeBuilder.Register(&BYOClusterProviderSpec{})
	SchemeBuilder.Register(&BYOClusterProviderStatus{})
}
