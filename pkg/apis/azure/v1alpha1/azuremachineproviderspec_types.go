package v1alpha1

import (
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// AzureMachineProviderSpec is the schema for the azuremachineproviderspec API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type AzureMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineSpecMixin `json:",inline"`

	VMSize string `json:"vmSize"`

	// Image is used to provide details of an image to use during VM creation.
	// If image details are omitted the image will default the Azure Marketplace "capi" offer,
	// which is based on Ubuntu.
	// +kubebuilder:validation:nullable
	// +optional
	Image *Image `json:"image,omitempty"`

	// OSDisk specifies the parameters for the operating system disk of the machine
	OSDisk OSDisk `json:"osDisk"`
}

func (s *AzureMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}
func (*AzureMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &AzureMachineProviderStatus{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&AzureMachineProviderSpec{})
}
