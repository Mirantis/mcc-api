package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MachineOpts struct {
	Name          string            `json:"Name"`
	UUID          string            `json:"UUID"`
	IPMIAddress   string            `json:"IPMIAddress"`
	IPMIPort      uint16            `json:"IPMIPort"`
	IPMIUsername  string            `json:"IPMIUsername"`
	IPMIPassword  string            `json:"IPMIPassword"`
	DeployKernel  string            `json:"DeployKernel"`
	DeployRAMDisk string            `json:"DeployRAMDisk"`
	RootGB        int               `json:"RootGB"`
	ImageSource   string            `json:"ImageSource"`
	ImageChecksum string            `json:"ImageChecksum"`
	MACAddress    string            `json:"MACAddress"`
	SSHKeys       map[string]string `json:"SSHKeys"`
	IPAddress     string            `json:"IPAddress"`
	IPGateway     string            `json:"IPGateway"`
}

type IronicClusterProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserDataSecret *corev1.SecretReference `json:"userDataSecret,omitempty"`

	MachineOpts MachineOpts `json:"machineOpts,inline"`
}
