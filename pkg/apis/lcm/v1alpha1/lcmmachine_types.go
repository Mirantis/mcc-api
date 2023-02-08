package v1alpha1

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	sysinfo "github.com/elastic/go-sysinfo/types"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
)

const (
	// LCMMachineTypeUnassigned denotes a machine that's doesn't
	// have an assigned role yet.
	// +gocode:public-api=true
	LCMMachineTypeUnassigned LCMMachineType = "unassigned"
	// LCMMachineTypeControl denotes a machine that belongs
	// to the control plane.
	// +gocode:public-api=true
	LCMMachineTypeControl LCMMachineType = "control"
	// LCMMachineTypeControl denotes a machine that's a worker node.
	// +gocode:public-api=true
	LCMMachineTypeWorker LCMMachineType = "worker"
	// LCMMachineTypeBastion denotes a machine that's a bastion node
	// +gocode:public-api=true
	LCMMachineTypeBastion LCMMachineType = "bastion"
	// +gocode:public-api=true
	LCMDownloaderRunner LCMStateItemRunner = "downloader"
	// +gocode:public-api=true
	LCMBashRunner LCMStateItemRunner = "bash"
	// +gocode:public-api=true
	LCMAnsibleRunner LCMStateItemRunner = "ansible"

	// +gocode:public-api=true
	LCMPreparePhase LCMItemPhase = "prepare"
	// +gocode:public-api=true
	LCMDeployPhase LCMItemPhase = "deploy"
	// +gocode:public-api=true
	LCMReconfigurePhase LCMItemPhase = "reconfigure"

	// +gocode:public-api=true
	LCMStateItemParamControlNodeIP = "controlNodeIP"
	// +gocode:public-api=true
	LCMStateItemParamControlPlaneNodes = "controlPlaneNodes"
	// +gocode:public-api=true
	LCMStateItemParamIsDedicatedMaster = "isDedicatedMaster"
	// +gocode:public-api=true
	LCMStateItemParamTrueValue = "true"
	// +gocode:public-api=true
	LCMStateItemParamFalseValue = "false"

	// +gocode:public-api=true
	LCMStateItemParamTargetRelease = "target_release"
	// +gocode:public-api=true
	LCMStateItemParamReleaseNeedsUpgrade = "release_needs_upgrade"

	// +gocode:public-api=true
	LCMStateItemParamUCPTag = "ucp_tag"
	// +gocode:public-api=true
	LCMStateItemParamUCPLB = "ucp_lb"
	// +gocode:public-api=true
	LCMStateItemParamSupplementaryAddressesInSSLKeys = "supplementary_addresses_in_ssl_keys"

	// +gocode:public-api=true
	LCMStateItemParamMirrorsMgmtDisabledTag = "mirrors_mgmt_disabled"

	// LCMClusterSkipMaintenanceAnnotation if set the evacuation is skip and
	// a maintenance won’t be initiated
	LCMClusterSkipMaintenanceAnnotation = "lcm.mirantis.com/skip-maintenance"

	// LCMMachineIndexAnnotation names an annotation that stores the
	// index used for the ordering of LCMMachines
	// +gocode:public-api=true
	LCMMachineIndexAnnotation = "lcm.mirantis.com/index"

	// LCMClusterMirrorsMgmtDisabledAnnotation
	LCMClusterMirrorsMgmtDisabledAnnotation = "lcm.mirantis.com/mirrors-mgmt-disabled"

	// The agent didn't set the address yet.
	// +gocode:public-api=true
	LCMMachineStateUninitialized LCMMachineState = "Uninitialized"
	// The agent did set the IP address, but there aren't any
	// state items set for this machine.
	// +gocode:public-api=true
	LCMMachineStatePending LCMMachineState = "Pending"
	// The agent did return an IP address, and the state items
	// are only set for Prepare phase.
	// +gocode:public-api=true
	LCMMachineStatePrepare LCMMachineState = "Prepare"
	// The agent did return an IP address, and the state items are
	// set for the phases up to and including the Deploy phase
	// (also possibly for Reconfigure phase, but this isn't
	// checked), but Prepare and/or Deploy phases didn't finish
	// yet.
	// +gocode:public-api=true
	LCMMachineStateDeploy LCMMachineState = "Deploy"
	// The agent did return an IP address, all of the state items
	// for the machine are set and fully processed by the agent, but
	// the desired UCP version doesn't match the actual one.
	// +gocode:public-api=true
	LCMMachineStateUpgrade LCMMachineState = "Upgrade"
	// The agent did return an IP address, and the state items are
	// set for the phases up to and including the Reconfigure phase,
	// and the Prepare and Deploy phases are fully processed,
	// but Reconfigure phase isn't processed yet.
	// +gocode:public-api=true
	LCMMachineStateReconfigure LCMMachineState = "Reconfigure"
	// The agent did return an IP address, all of the state items
	// for the machine are set and fully processed by the agent.
	// +gocode:public-api=true
	LCMMachineStateReady LCMMachineState = "Ready"
	// The agent is going to proceed with a node reboot
	// +gocode:public-api=true
	LCMMachineStateReboot LCMMachineState = "Reboot"

	// +gocode:public-api=true
	CordondrainStatusCordondraining CordondrainStatus = "cordondraining"
	// +gocode:public-api=true
	CordondrainStatusCordondrained CordondrainStatus = "cordondrained"
	// +gocode:public-api=true
	CordondrainStatusUncordoning CordondrainStatus = "uncordoning"

	// +gocode:public-api=true
	LCMTypeMKE LCMType = "ucp"
	// +gocode:public-api=true
	LCMTypeBYO LCMType = "byo"
	// +gocode:public-api=true
	LCMTypeK0s LCMType = "k0s"

	// +gocode:public-api=true
	PrepareDeletionPhaseStarted PrepareDeletionPhase = "started"
	// +gocode:public-api=true
	PrepareDeletionPhaseCompleted PrepareDeletionPhase = "completed"
	// +gocode:public-api=true
	PrepareDeletionPhaseAborting PrepareDeletionPhase = "aborting"
	// +gocode:public-api=true
	PrepareDeletionPhaseFailed PrepareDeletionPhase = "failed"

	LCMMachineGracefulRebootAnnotation = "lcm.mirantis.com/graceful-reboot"
)

type DiskPartition struct {
	// Name of disk partition
	Name string `json:"name"`
	// Size in bytes
	Size int `json:"size"`
	// Label of disk partition
	// +optional
	Label string `json:"label"`
	// Path to part of disk by path
	// +optional
	// +nullable
	ByPaths []string `json:"byPaths"`
	// Path to part of disk by ID
	// +optional
	// +nullable
	ByIDs []string `json:"byIDs"`
}

// +gocode:public-api=true
type LCMType string

// LCMMachineStatus defines the observed state of LCMMachine
type LCMMachineStatus struct {
	// StateItemStatuses maps state item names to its statuses.
	// +optional
	StateItemStatuses map[string]LCMStateItemStatus `json:"stateItemStatuses,omitempty"`
	// +optional
	Addresses []v1.NodeAddress `json:"addresses,omitempty"`
	// +optional
	AgentVersion string `json:"agentVersion,omitempty"`
	// +optional
	AgentUpgradeStatus *LCMAgentUpgradeStatus `json:"lcmAgentUpgradeStatus,omitempty"`
	// +optional
	State LCMMachineState `json:"state,omitempty"`
	// Status of components
	Components LCMComponentsStatus `json:"components,omitempty"`
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// RHELLicenseApplied status
	// +optional
	RHELLicenseApplied *bool `json:"rhelLicenseApplied,omitempty"`
	// RHELLicenseError content
	// +optional
	RHELLicenseError string `json:"rhelLicenseError,omitempty"`
	// MCC release associated with processed configuration
	Release   string `json:"release,omitempty"`
	Interface string `json:"interface,omitempty"`
	// Status of the most recent reboot request
	Reboot *RebootStatus `json:"reboot,omitempty"`
	// Current version of MCR
	MCRVersion string `json:"mcrVersion,omitempty"`
	// Host system information
	HostInfo HostInfo `json:"hostInfo,omitempty"`
	// Name of the secret containing a service account token
	TokenSecret string `json:"tokenSecret,omitempty"`
	// MKEAuthHash contains a hash of the last fetched custom mke auth data
	MKEAuthHash string `json:"mkeAuthHash,omitempty"`
	// Maintenance flag indicates that Node is drained, cordoned and switched
	// into maintenance state
	Maintenance bool `json:"maintenance,omitempty"`
	// KubernetesCordonDrain indicates kubelet state in which the machine is
	// with respect to LCMClusterState objects present for it
	KubernetesCordonDrain CordondrainStatus `json:"kubernetesCordonDrain,omitempty"`
	// KubernetesCordonDrainError contains an error message if there was one
	// while the kubernetes machine drain process
	// +optional
	KubernetesCordonDrainError string `json:"kubernetesCordonDrainError,omitempty"`
	// SwarmCordonDrain indicates swarm state in which the machine is
	// with respect to LCMClusterState objects present for it
	SwarmCordonDrain CordondrainStatus `json:"swarmCordonDrain,omitempty"`
	// SwarmCordonDrainError contains an error message if there was one
	// while the swarm machine drain process
	// +optional
	SwarmCordonDrainError string `json:"swarmCordonDrainError,omitempty"`
	// LCMOperationStuck flag indicates that some LCM operation with the machine
	// is stuck, and operator needs to take a closer look at its status
	// +optional
	LCMOperationStuck bool `json:"lcmOperationStuck,omitempty"`
	// PrepareDeletionPhase indicates status of preparation Node for deletion.
	// +optional
	PrepareDeletionPhase PrepareDeletionPhase `json:"prepareDeletionPhase,omitempty"`
	// MKEManagerReady indicates swarm endpoint readiness
	// +optional
	MKEManagerReady *bool `json:"mkeManagerReady,omitempty"`

	// Distribution represents a list of possible Distribution ID values.
	// Computed from HostInfo data.
	// +optional
	Distribution []string `json:"Distribution,omitempty"`
}

// HardwareDetails collects all of the information about hardware
// discovered on the host.
type HardwareDetails struct {
	CPU     CPU     `json:"cpu"`
	Memory  Memory  `json:"memory"`
	Storage Storage `json:"storage"`
	Network Network `json:"network"`
}
type Memory struct {
	// Total memory in bytes
	Total int `json:"total"`
}

// LCMStateItemStatus defines a status for LCMStateItem
type LCMStateItemStatus struct {
	// Hash() value of StateItem.
	Hash string `json:"hash"`
	// Exit code of the command that's used to execute this item
	ExitCode int `json:"exitCode"`
	// An optional message from the runner to be displayed
	// for the operator
	Message string `json:"message,omitempty"`
	// Item-specific JSON object with the result.
	Result runtime.RawExtension `json:"result,omitempty"`
	// Attempt number of re-runs
	Attempt int `json:"attempt"`
	// Start time of running item
	StartedAt metav1.Time `json:"startedAt,omitempty"`
	// Finish time of running item
	FinishedAt metav1.Time `json:"finishedAt,omitempty"`
}

// IsReady returns true if the agent successfully handled the status
// item.
func (s LCMStateItemStatus) IsReady() bool {
	return s.ExitCode == 0
}

type LCMComponents struct {
	// UCP version
	UCPVersion string `json:"ucpVersion,omitempty"`
}
type AgentConfig struct {
	Version     string   `json:"version,omitempty"`
	DownloadURL string   `json:"downloadURL,omitempty"`
	SHA256      string   `json:"sha256,omitempty"`
	Args        []string `json:"args,omitempty"`
}

// LCMMachineList contains a list of LCMMachine objects
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LCMMachine `json:"items"`
}

// LCMStateItem defines a part of a machine state that
// can be represented e.g. by an Ansible playbook and
// inventory variables.
type LCMStateItem struct {
	// Runner to use, like bash, ansible
	Runner LCMStateItemRunner `json:"runner"`
	// The name of the state item
	// (may correspond to playbook / role name)
	Name string `json:"name"`
	// Version is specified as a separate field from Name
	// so as to make this state item occupy the same "slot" in
	// the status. It can be a hash of playbook / role content,
	// etc.
	Version string `json:"version"`
	// Item parameters (e.g. may be used for Ansible inventory).
	// Changing the params causes the item to be re-run.
	// TODO: use runtime.RawExtension for values
	Params map[string]string `json:"params,omitempty"`

	// Phase in which run the item
	// +optional
	Phase LCMItemPhase `json:"phase,omitempty"`
}

func (item LCMStateItem) Hash(lcmVersion string) string {

	if len(item.Params) == 0 {
		item.Params = nil
	} else {
		if AgentGreater187(lcmVersion) {
			newParams := make(map[string]string)
			for k, v := range item.Params {
				if k == LCMStateItemParamControlNodeIP || k == LCMStateItemParamUCPTag {
					continue
				}
				if k == LCMStateItemParamTargetRelease || k == LCMStateItemParamReleaseNeedsUpgrade {
					continue
				}
				if k == LCMStateItemParamUCPLB || k == LCMStateItemParamSupplementaryAddressesInSSLKeys ||
					k == LCMStateItemParamMirrorsMgmtDisabledTag {
					if AgentGreater257(lcmVersion) {
						continue
					}
				}
				newParams[k] = v
			}
			item.Params = newParams
		}
	}

	out, err := json.Marshal(item)
	if err != nil {
		log.Panicf("json.Marshal(): %v", err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(out))
}

// LCMMachineType denotes a type of the machine
// +gocode:public-api=true
type LCMMachineType string
type HostInfo struct {
	// Host boot time
	BootTime *metav1.Time `json:"bootTime,omitempty"`
	// Hardware architecture (e.g. x86_64, arm, ppc, mips)
	Architecture string `json:"architecture,omitempty"`
	// Kernel version
	KernelVersion string `json:"kernelVersion,omitempty"`
	// OS information
	OS *sysinfo.OSInfo `json:"os,omitempty"`
	// System timezone
	Timezone string `json:"timezone,omitempty"`
	// Repository list
	// +nullable
	Repositories []*Repository `json:"repositories,omitempty"`
	// Available new distributive updates
	UpdatesAvailable bool `json:"updatesAvailable,omitempty"`
	// Hardware of host
	// +nullable
	Hardware *HardwareDetails `json:"hardware,omitempty"`
	// OS needs reboot
	RebootRequired bool `json:"rebootRequired,omitempty"`
	// Information about the reasons why a reboot is required
	RebootReason string `json:"rebootReason,omitempty"`
}

// LCMMachineSpec defines the desired state of LCMMachine
type LCMMachineSpec struct {
	// The name of the cluster this machine belongs too.
	// This field will be set by the cluster actuator.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`
	// The type of the machine.
	// This field will be set by the cluster actuator.
	// When LCMMachine object is first created by the agent,
	// this field has "unassigned" type.
	// +kubebuilder:validation:Enum=control;worker;bastion;unassigned
	Type LCMMachineType `json:"type,omitempty"`
	// Components which may be upgraded/managed separatelly  from StateItems
	Components LCMComponents `json:"components,omitempty"`
	// StateItems represent the target state for this machine.
	// +optional
	StateItems []LCMStateItem `json:"stateItems,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
	// StateItemsOverwrites stores machine specific params.
	// +optional
	StateItemsOverwrites map[string]LCMParams `json:"stateItemsOverwrites,omitempty"`
	// DeploymentSecretName specifies the name of the secret
	// that's used for deployment
	// TODO move it to a CRD
	DeploymentSecretName string `json:"deploymentSecretName,omitempty"`
	// KubeconfigSecretName specifies the name of the secret
	// that's used for tenant cluster's kubeconfig
	// TODO rename it to AgentDataSecretName
	KubeconfigSecretName string `json:"kubeconfigSecretName,omitempty"`
	// ProxySecretName specifies the name of the secret
	// that's used for setting proxy params
	ProxySecretName string `json:"proxySecretName,omitempty"`
	// TLSSecretName specifies the name of the secret
	// that's used for applying TLS configuration on a node
	TLSSecretName string `json:"tlsSecretName,omitempty"`
	// AgentConfig stores configuration for a node agent
	// +optional
	AgentConfig AgentConfig `json:"agentConfig,omitempty"`
	// SSH Authorized Keys is a list of user defined ssh keys
	// in the authorized keys format
	// +optional
	SSHAuthorizedKeys []string `json:"sshAuthorizedKeys,omitempty"`
	// RHELLicenseSubscription object reference
	// +optional
	RHELLicenseSubscription string `json:"rhelLicenseSubscription,omitempty"`
	// MCC release associated with requested configuration
	Release string `json:"release,omitempty"`
	// Interface to get InternalIP from
	Interface string `json:"interface,omitempty"`
	// CIDR to get InternalIP from
	CIDR string `json:"cidr,omitempty"`
	// Flag indicating that machine needs to be rebooted after deploy stage is finished
	RebootRequired bool `json:"rebootRequired,omitempty"`
	// Requested version of MCR
	MCRVersion string `json:"mcrVersion,omitempty"`
	// Name of the service account to get auth tokens from
	ServiceAccount string `json:"serviceAccount,omitempty"`
	// Maintenance flag indicates that Node should be drained, cordoned and switched
	// into maintenance state
	Maintenance bool `json:"maintenance,omitempty"`
	// HaltAgent flag indicates that only public keys management part of
	// lcm agent should be left functioning
	HaltAgent bool `json:"haltAgent,omitempty"`
	// LCMType contains the LCM distribution type
	// +kubebuilder:validation:Enum=ucp;byo;k0s
	LCMType LCMType `json:"lcmType,omitempty"`
	// PrepareDeletion flag indicates that LCM has to prepare Node for deleting.
	PrepareDeletion bool `json:"prepareDeletion,omitempty"`

	// Distribution represents ID of `Distribution` object inside list of
	// allowed distributions on per-cluster release basis.
	// It is synced from Machine object.
	// +optional
	Distribution string `json:"distribution,omitempty"`
}
type RebootStatus struct {
	// Time when reboot was requested
	RequestedAt metav1.Time `json:"requestedAt,omitempty"`
	// Indicates whether reboot is finished or not
	Completed bool `json:"completed,omitempty"`
	// An error message to be displayed for the operator
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// LCMAgentUpgradeStatus defines a status of LCM Agent upgrade
type LCMAgentUpgradeStatus struct {
	// Error stores a error message in case of failed upgrade
	Error string `json:"error,omitempty"`
	// Attempt number of re-runs
	Attempt int `json:"attempt"`
	// Start time of running item
	StartedAt metav1.Time `json:"startedAt,omitempty"`
	// Finish time of running item
	FinishedAt metav1.Time `json:"finishedAt,omitempty"`
}
type Network struct {
	//  List of available network interface controllers
	// +nullable
	NICs []*NIC `json:"nics"`
}

// +gocode:public-api=true
type CordondrainStatus string

// +gocode:public-api=true
type PrepareDeletionPhase string
type CPU struct {
	// Count cores
	Count int `json:"count"`
	// Threads of all cores
	Threads int `json:"threads"`
}
type LCMComponentsStatus struct {
	// UCP version
	UCPVersion string `json:"ucpVersion,omitempty"`
}
type Disk struct {
	// Serial number of disk
	// +optional
	SerialNumber string `json:"serialNumber"`
	// Name of disk
	Name string `json:"name"`
	// Size in bytes
	Size int `json:"size"`
	// Type of disk (example: hdd, ssd, fdd, odd)
	Type string `json:"type"`
	// Vendor name (example: ATA)
	// +optional
	Vendor string `json:"vendor"`
	// Model name (example: VBOX_HARDDISK)
	// +optional
	Model string `json:"model"`
	// Databus
	// +optional
	BusPath string `json:"busPath"`
	// Path to disk by path
	// DEPRECATED. Use ByPaths field
	// +optional
	ByPath string `json:"byPath"`
	// Path to disk by path
	// +optional
	// +nullable
	ByPaths []string `json:"byPaths"`
	// Path to disk by ID
	// DEPRECATED. Use ByIDs field
	// +optional
	ByID string `json:"byID"`
	// Path to disk by ID
	// +optional
	// +nullable
	ByIDs []string `json:"byIDs"`
	// Partitions of disk
	// +nullable
	Partitions []*DiskPartition `json:"partitions"`
}

// +gocode:public-api=true
type LCMItemPhase string

// +gocode:public-api=true
type LCMMachineState string
type Repository struct {
	// Remote URI
	URI string `json:"uri,omitempty"`
	// Release name
	Release string `json:"release,omitempty"`
	// Section names or components. There can be several section names, separated by spaces
	Section []string `json:"section,omitempty"`
}

// Network interface controller
type NIC struct {
	// Name of network interface controller
	Name string `json:"name"`
	// Media access control address
	MacAddress string `json:"macAddress"`
	// Is virtual network interface
	IsVirtual bool `json:"isVirtual"`
	// PCI address
	// +nullable
	PCIAddress *string `json:"pciAddress"`
}

// LCMMachine is the Schema for the lcmmachines API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="clusterName",type="string",JSONPath=".spec.clusterName",description="Cluster Name",priority=0
// +kubebuilder:printcolumn:name="type",type="string",JSONPath=".spec.type",description="Type",priority=0
// +kubebuilder:printcolumn:name="state",type="string",JSONPath=".status.state",description="State",priority=0
// +kubebuilder:printcolumn:name="internalIP",type="string",JSONPath=".status.addresses[?(@.type==\"InternalIP\")].address",description="Internal IP",priority=1
// +kubebuilder:printcolumn:name="hostname",type="string",JSONPath=".status.addresses[?(@.type==\"Hostname\")].address",description="Hostname",priority=1
// +kubebuilder:printcolumn:name="agentVersion",type="string",JSONPath=".status.agentVersion",description="Agent Version",priority=1
// +gocode:public-api=true
type LCMMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LCMMachineSpec   `json:"spec,omitempty"`
	Status LCMMachineStatus `json:"status,omitempty"`
}

func (m *LCMMachine) IsNodeInitialized() bool {
	gotIP := false
	for _, addr := range m.Status.Addresses {
		if addr.Type == v1.NodeInternalIP {
			gotIP = true
			break
		}
	}
	if !gotIP {
		return false
	}
	if m.Spec.RHELLicenseSubscription != "" && !BoolValue(m.Status.RHELLicenseApplied) {
		return false
	}
	return true
}

// IsReady returns true if all of the LCMMachine's state items
// were successfully processed by the agent.
// TODO: test it
func (m *LCMMachine) IsReady() bool {
	if len(m.Status.StateItemStatuses) == 0 {
		return false
	}

	if !m.IsNodeInitialized() {
		return false
	}

	for _, item := range m.Spec.StateItems {
		status, found := m.Status.StateItemStatuses[item.Name]
		if !found || item.Hash(m.Spec.AgentConfig.Version) != status.Hash || !status.IsReady() {
			return false
		}
	}

	return true
}

// WasReady returns true if Status.Release is not empty
// It's set by agent/ucp controller (for byo) when machine gets ready
func (m *LCMMachine) WasReady() bool {
	return m.Status.Release != ""
}

// IsPhaseReady returns true if all of the LCMMachine's state items
// from the phase were successfully processed by the agent.
func (m *LCMMachine) IsPhaseReady(phase LCMItemPhase) bool {
	if !m.IsNodeInitialized() {
		return false
	}

	for _, item := range m.Spec.StateItems {
		if item.Phase != phase {
			continue
		}
		status, found := m.Status.StateItemStatuses[item.Name]
		if !found || item.Hash(m.Spec.AgentConfig.Version) != status.Hash || !status.IsReady() {
			return false
		}
	}
	return true
}

// ComputeState returns the machine state based on its status,
// StateItems and other components state
func (m *LCMMachine) ComputeState() LCMMachineState {
	if !m.IsNodeInitialized() {
		return LCMMachineStateUninitialized
	}

	if m.Spec.Type == LCMMachineTypeBastion {
		return LCMMachineStateReady
	}
	if m.IsK0s() {
		state := m.ComputeItemsState()
		if state == LCMMachineStateReady {
			if m.NeedsReboot() {
				return LCMMachineStateReboot
			}
		}
		return state
	}
	current := m.Status.Components.UCPVersion
	expected := m.Spec.Components.UCPVersion

	if m.IsBYO() {
		if expected == "" || current == "" {
			return LCMMachineStatePending
		}
		if current != expected {
			return LCMMachineStateUpgrade
		}
		return LCMMachineStateReady
	}

	if m.IsMKE() {
		state := m.ComputeItemsState()
		if state == LCMMachineStateReady {
			if m.NeedsReboot() {
				return LCMMachineStateReboot
			}
			if current == "" || expected == "" {
				return LCMMachineStateDeploy
			}
			if current != expected {
				return LCMMachineStateUpgrade
			}
		}
		return state
	}
	return LCMMachineStateUninitialized
}
func (m *LCMMachine) MKEReady() bool {
	if m.Spec.Type == LCMMachineTypeBastion {
		return true
	}
	if m.IsBYO() || m.IsMKE() {
		current := m.Status.Components.UCPVersion
		expected := m.Spec.Components.UCPVersion
		return current != "" && expected != "" && current == expected
	}
	return true
}
func (m *LCMMachine) ComputeItemsState() LCMMachineState {
	if !m.IsNodeInitialized() {
		return LCMMachineStateUninitialized
	}
	if len(m.Spec.StateItems) == 0 {
		return LCMMachineStatePending
	}

	havePhases := map[LCMItemPhase]bool{}
	for _, item := range m.Spec.StateItems {
		havePhases[item.Phase] = true
	}

	state := LCMMachineStateReady

	if havePhases[LCMPreparePhase] {

		state = LCMMachineStatePrepare
	}

	if havePhases[LCMDeployPhase] {
		if m.IsPhaseReady(LCMDeployPhase) {
			state = LCMMachineStateReady
		} else {
			state = LCMMachineStateDeploy
		}
	}

	if havePhases[LCMReconfigurePhase] && state != LCMMachineStateDeploy {
		if m.IsPhaseReady(LCMReconfigurePhase) {
			state = LCMMachineStateReady
		} else {
			state = LCMMachineStateReconfigure
		}
	}

	return state
}
func (m *LCMMachine) NeedsReboot() bool {
	return m.Spec.RebootRequired && (m.Status.Reboot == nil || !m.Status.Reboot.Completed)
}
func (m *LCMMachine) GracefulRebootRequired() bool {
	return m.GracefulRebootRequested() && !m.Spec.RebootRequired && (m.Status.Reboot == nil || !m.Status.Reboot.Completed)
}
func (m *LCMMachine) GracefulRebootRequested() bool {
	_, exists := m.Annotations[LCMMachineGracefulRebootAnnotation]
	return exists
}

// ControlPlaneRebootInProgress returns true if reboot is pending, in progress or MKE Manger has not started yet
func (m *LCMMachine) ControlPlaneRebootInProgress() bool {
	return m.NeedsReboot() ||
		m.Status.Reboot != nil && m.Status.Reboot.Completed &&
			(m.Status.MKEManagerReady == nil || !*m.Status.MKEManagerReady)
}
func (m *LCMMachine) IsBYO() bool {
	return m.Spec.LCMType == LCMTypeBYO
}
func (m *LCMMachine) IsMKE() bool {
	return m.Spec.LCMType == LCMTypeMKE
}
func (m *LCMMachine) IsK0s() bool {
	return m.Spec.LCMType == LCMTypeK0s
}
func (m *LCMMachine) SwarmDrainAllowed() bool {

	return m.Spec.Type != LCMMachineTypeControl
}

// +gocode:public-api=true
type LCMParams map[string]string

// +gocode:public-api=true
type LCMStateItemRunner string
type Storage struct {
	// List of available disks
	Disks []*Disk `json:"disks"`
}

// ProxyStateItemParamsWithKeys gets proxy related Params and keys map from LCMStateItem list
// proxy params are present in all LCMStateItem items in current implementation, so just taking first
// +gocode:public-api=true
func ProxyStateItemParamsWithKeys(stateItems []LCMStateItem) map[string]string {
	for _, item := range stateItems {
		return map[string]string{
			"http_proxy":  item.Params["cluster_http_proxy"],
			"HTTP_PROXY":  item.Params["cluster_http_proxy"],
			"https_proxy": item.Params["cluster_https_proxy"],
			"HTTPS_PROXY": item.Params["cluster_https_proxy"],
			"no_proxy":    item.Params["cluster_no_proxy"],
			"NO_PROXY":    item.Params["cluster_no_proxy"],
		}
	}
	return nil
}

// ProxyStateItemParams gets proxy related Params from LCMStateItem list
// proxy params are present in all LCMStateItem items in current implementation, so just taking first
// +gocode:public-api=true
func ProxyStateItemParams(stateItems []LCMStateItem) map[string]string {
	for _, item := range stateItems {
		httpProxy, httpProxyExists := item.Params["cluster_http_proxy"]
		httpsProxy, httpsProxyExists := item.Params["cluster_http_proxy"]
		noProxy, noProxyExists := item.Params["cluster_no_proxy"]
		if httpProxyExists || httpsProxyExists || noProxyExists {
			return map[string]string{
				"http_proxy":  httpProxy,
				"https_proxy": httpsProxy,
				"no_proxy":    noProxy,
			}
		}
	}
	return nil
}

// BoolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
// +gocode:public-api=true
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&LCMMachine{}, &LCMMachineList{})
}
