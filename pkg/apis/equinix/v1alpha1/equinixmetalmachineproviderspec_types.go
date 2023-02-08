package v1alpha1

import (
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
	Ceph CephMachineConfig `json:"ceph,omitempty"`

	OS          string `json:"OS"`
	MachineType string `json:"machineType"`

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
	Tags Tags `json:"tags,omitempty"`
}

func (s *EquinixMetalMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}
func (*EquinixMetalMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &EquinixMetalMachineProviderStatus{}
}

type CephMachineConfig struct {
	// ManagerMonitor is an option that reflects that Manager and Monitor
	// roles should be enabled for that Node in CephCluster spec
	ManagerMonitor bool `json:"managerMonitor,omitempty"`
	// Storage is an option that reflects that Storage
	// role should be enabled for that Node in CephCluster spec
	Storage bool `json:"storage,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&EquinixMetalMachineProviderSpec{})
}
