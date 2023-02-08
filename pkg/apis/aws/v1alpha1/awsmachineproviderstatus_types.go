package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field.
// It containsk AWS-specific status information.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type AWSMachineProviderStatus struct {
	metav1.TypeMeta                 `json:",inline"`
	metav1.ObjectMeta               `json:"metadata,omitempty"`
	kaasv1alpha1.MachineStatusMixin `json:",inline"`
	// TODO: remove InstanceID, InstanceState and use the ones from MachineStatusMixin
	// InstanceID is the instance ID of the machine created in AWS
	// +optional
	InstanceID *string `json:"instanceID,omitempty"`

	// InstanceState is the state of the AWS instance for this machine
	// +optional
	InstanceState *InstanceState `json:"instanceState,omitempty"`
}

func (s *AWSMachineProviderStatus) GetMachineStatusMixin() *kaasv1alpha1.MachineStatusMixin {
	return &s.MachineStatusMixin
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AWSMachineProviderStatus{})
}
