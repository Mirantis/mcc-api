package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +gocode:public-api=true
type MachineOpts struct {
	Name          string
	UUID          string
	IPMIAddress   string
	IPMIPort      uint16
	IPMIUsername  string
	IPMIPassword  string
	DeployKernel  string
	DeployRAMDisk string
	RootGB        int
	ImageSource   string
	ImageChecksum string
	MACAddress    string
	SSHKeys       map[string]string
	IPAddress     string
	IPGateway     string
}

// +gocode:public-api=true
type IronicClusterProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserDataSecret *corev1.SecretReference `json:"userDataSecret,omitempty"`

	MachineOpts MachineOpts `json:"machineOpts,inline"`
}
