package v1alpha1

import (
	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BareMetalMachineProviderStatusList contains a list of BareMetalMachineProviderStatus
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BareMetalMachineProviderStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BareMetalMachineProviderStatus `json:"items"`
}

// BareMetalMachineProviderStatus is the Schema for the baremetalmachineproviderstatuses API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BareMetalMachineProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.MachineStatusMixin `json:",inline"`
}

func (s *BareMetalMachineProviderStatus) GetMachineStatusMixin() *kaas.MachineStatusMixin {
	return &s.MachineStatusMixin
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&BareMetalMachineProviderStatus{}, &BareMetalMachineProviderStatusList{})
}
