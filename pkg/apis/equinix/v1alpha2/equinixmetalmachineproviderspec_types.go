package v1alpha2

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	ipamcommon "github.com/Mirantis/mcc-api/v2/pkg/apis/external/kaas-ipam/pkg/ipam/types"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EquinixMetalMachineProviderSpec is the schema for the equinixmetalmachineproviderspec API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type EquinixMetalMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	kaasv1alpha1.MachineSpecMixin `json:",inline"`

	// Ceph contains Ceph configuration options for Machine
	Ceph v1alpha1.CephMachineConfig `json:"ceph,omitempty"`

	MachineType string `json:"machineType"`

	// BareMetalHostProfile is the name of the BareMetalHostProfile that should be used
	// to provision the machine. BareMetalHostProfile defines how the storage devices
	// and operating system are provisioned and configured. Custom BareMetalHostProfile should
	// be created in the same namespace as the machine. If unset, the default
	// BareMetalHostProfile (namespace: kaas, name: <machineType>) will be chosen.
	// +optional
	BareMetalHostProfile string `json:"bareMetalHostProfile,omitempty"`

	// Network field allows to customize network configuration for each machine.
	// +optional
	Network MachineNetwork `json:"network,omitempty"`

	// IPXEUrl can be used to set the pxe boot url when using custom OSes with this provider.
	// Note that OS should also be set to "custom_ipxe" if using this value.
	// +optional
	IPXEUrl string `json:"ipxeURL,omitempty" sensitive:"true"`

	// HardwareReservationID is the unique device hardware reservation ID or `next-available` to
	// automatically let the EquinixMetal api determine one.
	// +optional
	HardwareReservationID string `json:"hardwareReservationID,omitempty" sensitive:"true"`

	// ProviderID is the unique identifier as specified by the cloud provider.
	// +optional
	ProviderID *string `json:"providerID,omitempty" sensitive:"true"`

	// Tags is an optional set of tags to add to EquinixMetal resources managed by the EquinixMetal provider.
	// +optional
	Tags v1alpha1.Tags `json:"tags,omitempty"`
}

func (s *EquinixMetalMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}
func (*EquinixMetalMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &EquinixMetalMachineProviderStatus{}
}

// MachineNetwork contains the configuration for custom machine network
type MachineNetwork struct {
	// L2TemplateSelector is the selector of the L2Template that should be used
	// to configure networking on the machine. Custom L2Template should
	// be created in the same namespace as the machine.
	// L2TemplateSelector will be passed to IpamHost while creating.
	// +optional
	L2TemplateSelector ipamcommon.L2TemplateSelector `json:"l2TemplateSelector,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalMachineProviderSpec{})
}
