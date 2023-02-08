package v1alpha1

import (
	kaas "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// BYOMachineProviderSpec is the type that will be embedded in a Machine.Spec.ProviderSpec field
// for an BYO instance.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BYOMachineProviderSpec struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	kaas.MachineSpecMixin `json:",inline"`
}

func (*BYOMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &BYOMachineProviderStatus{}
}
func (s *BYOMachineProviderSpec) GetMachineSpecMixin() *kaas.MachineSpecMixin {
	return &s.MachineSpecMixin
}

// BYOMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field.
// It containsk BYO-specific status information.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type BYOMachineProviderStatus struct {
	metav1.TypeMeta         `json:",inline"`
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	kaas.MachineStatusMixin `json:",inline"`
}

func (s *BYOMachineProviderStatus) GetMachineStatusMixin() *kaas.MachineStatusMixin {
	return &s.MachineStatusMixin
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&BYOMachineProviderSpec{})
	SchemeBuilder.Register(&BYOMachineProviderStatus{})
}
