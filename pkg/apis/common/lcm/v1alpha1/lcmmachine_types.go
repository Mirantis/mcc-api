/*
Copyright 2019 The Mirantis Authors.

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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"

	sysinfo "github.com/elastic/go-sysinfo/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// LCMMachineType denotes a type of the machine
type LCMMachineType string

type LCMStateItemRunner string

type LCMItemPhase string

type LCMMachineState string

type LCMParams map[string]string

type CordondrainStatus string

type LCMType string

const (
	// LCMMachineTypeUnassigned denotes a machine that's doesn't
	// have an assigned role yet.
	LCMMachineTypeUnassigned LCMMachineType = "unassigned"
	// LCMMachineTypeControl denotes a machine that belongs
	// to the control plane.
	LCMMachineTypeControl LCMMachineType = "control"
	// LCMMachineTypeControl denotes a machine that's a worker node.
	LCMMachineTypeWorker LCMMachineType = "worker"
	// LCMMachineTypeBastion denotes a machine that's a bastion node
	LCMMachineTypeBastion LCMMachineType = "bastion"

	LCMDownloaderRunner LCMStateItemRunner = "downloader"
	LCMBashRunner       LCMStateItemRunner = "bash"
	LCMAnsibleRunner    LCMStateItemRunner = "ansible"

	LCMPreparePhase     LCMItemPhase = "prepare"
	LCMDeployPhase      LCMItemPhase = "deploy"
	LCMReconfigurePhase LCMItemPhase = "reconfigure"

	LCMStateItemParamControlNodeIP     = "controlNodeIP"
	LCMStateItemParamControlPlaneNodes = "controlPlaneNodes"
	LCMStateItemParamIsDedicatedMaster = "isDedicatedMaster"
	LCMStateItemParamTrueValue         = "true"

	LCMStateItemParamUCPTag = "ucp_tag"

	// LCMMachineIndexAnnotation names an annotation that stores the
	// index used for the ordering of LCMMachines
	LCMMachineIndexAnnotation = "lcm.mirantis.com/index"

	// LCMMachineWasDeployedAnnotation names an annotation that is present
	// if this machine's agent ever started processing StateItems of the
	// "deploy" phase. This means that the corresponding node needs to be
	// cordoned and drained if it's present in the child apiserver.
	LCMMachineWasDeployedAnnotation = "lcm.mirantis.com/was-deployed"

	// LCMMachineWasReadyAnnotation names an annotation that is present
	// if the machine has ever reached Ready state, meaning that it no
	// longer requires kubeadm token for updating
	LCMMachineWasReadyAnnotation = "lcm.mirantis.com/was-ready"

	// The agent didn't set the address yet.
	LCMMachineStateUninitialized LCMMachineState = "Uninitialized"
	// The agent did set the IP address, but there aren't any
	// state items set for this machine.
	LCMMachineStatePending LCMMachineState = "Pending"
	// The agent did return an IP address, and the state items
	// are only set for Prepare phase.
	LCMMachineStatePrepare LCMMachineState = "Prepare"
	// The agent did return an IP address, and the state items are
	// set for the phases up to and including the Deploy phase
	// (also possibly for Reconfigure phase, but this isn't
	// checked), but Prepare and/or Deploy phases didn't finish
	// yet.
	LCMMachineStateDeploy LCMMachineState = "Deploy"
	// The agent did return an IP address, all of the state items
	// for the machine are set and fully processed by the agent, but
	// the desired UCP version doesn't match the actual one.
	LCMMachineStateUpgrade LCMMachineState = "Upgrade"
	// The agent did return an IP address, and the state items are
	// set for the phases up to and including the Reconfigure phase,
	// and the Prepare and Deploy phases are fully processed,
	// but Reconfigure phase isn't processed yet.
	LCMMachineStateReconfigure LCMMachineState = "Reconfigure"
	// The agent did return an IP address, all of the state items
	// for the machine are set and fully processed by the agent.
	LCMMachineStateReady LCMMachineState = "Ready"
	// The agent is going to proceed with a node reboot
	LCMMachineStateReboot LCMMachineState = "Reboot"

	CordondrainStatusCordondraining CordondrainStatus = "cordondraining"
	CordondrainStatusCordondrained  CordondrainStatus = "cordondrained"
	CordondrainStatusUncordoning    CordondrainStatus = "uncordoning"

	LCMTypeMKE LCMType = "ucp"
	LCMTypeBYO LCMType = "byo"
	LCMTypeK0s LCMType = "k0s"
)

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
}

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
}

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
}

// HardwareDetails collects all of the information about hardware
// discovered on the host.
type HardwareDetails struct {
	CPU     CPU     `json:"cpu"`
	Memory  Memory  `json:"memory"`
	Storage Storage `json:"storage"`
	Network Network `json:"network"`
}

type CPU struct {
	// Count cores
	Count int `json:"count"`
	// Threads of all cores
	Threads int `json:"threads"`
}

type Memory struct {
	// Total memory in bytes
	Total int `json:"total"`
}

type Storage struct {
	// List of available disks
	Disks []*Disk `json:"disks"`
}

type Disk struct {
	// Serial number of disk
	SerialNumber string `json:"serialNumber"`
	// Name of disk
	Name string `json:"name"`
	// Size in bytes
	Size int `json:"size"`
	// Type of disk (example: hdd, ssd, fdd, odd)
	Type string `json:"type"`
	// Vendor name (example: ATA)
	Vendor string `json:"vendor"`
	// Model name (example: VBOX_HARDDISK)
	Model string `json:"model"`
	// Databus
	BusPath string `json:"busPath"`
	// Path to disk by path
	ByPath string `json:"byPath"`
	// Path to disk by ID
	ByID string `json:"byID"`
	// Partitions of disk
	// +nullable
	Partitions []*DiskPartition `json:"partitions"`
}

type DiskPartition struct {
	// Name of disk partition
	Name string `json:"name"`
	// Size in bytes
	Size int `json:"size"`
	// Label of disk partition
	// +optional
	Label string `json:"label"`
}

type Network struct {
	//  List of available network interface controllers
	// +nullable
	NICs []*NIC `json:"nics"`
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

type Repository struct {
	// Remote URI
	URI string `json:"uri,omitempty"`
	// Release name
	Release string `json:"release,omitempty"`
	// Section names or components. There can be several section names, separated by spaces
	Section []string `json:"section,omitempty"`
}

type RebootStatus struct {
	// Time when reboot was requested
	RequestedAt metav1.Time `json:"requestedAt,omitempty"`
	// Indicates whether reboot is finished or not
	Completed bool `json:"completed,omitempty"`
	// An error message to be displayed for the operator
	ErrorMessage string `json:"errorMessage,omitempty"`
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

func ControlNodeIPChanged(current, expected LCMStateItem, m *LCMMachine) bool {
	cur := current.Params[LCMStateItemParamControlNodeIP]
	exp := expected.Params[LCMStateItemParamControlNodeIP]
	return cur != exp
}

func UCPTagChanged(current, expected LCMStateItem, m *LCMMachine) bool {
	cur := current.Params[LCMStateItemParamUCPTag]
	exp := expected.Params[LCMStateItemParamUCPTag]
	return cur != exp
}

func (item LCMStateItem) Hash(lcmVersion string) string {
	// TODO: use something with more guarantees of stability than JSON
	if len(item.Params) == 0 {
		item.Params = nil
	} else { //ignore controlNodeIP for hash calculation if agent version requires so
		if AgentGreater187(lcmVersion) { //"" for tests
			newParams := make(map[string]string)
			for k, v := range item.Params {
				if k != LCMStateItemParamControlNodeIP && k != LCMStateItemParamUCPTag {
					newParams[k] = v
				}
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

type LCMComponents struct {
	// UCP version
	UCPVersion string `json:"ucpVersion,omitempty"`
}

type LCMComponentsStatus struct {
	// UCP version
	UCPVersion string `json:"ucpVersion,omitempty"`
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

// IsReady returns true if the agent successfully handled the status
// item.
func (s LCMStateItemStatus) IsReady() bool {
	return s.ExitCode == 0
}

type AgentConfig struct {
	Version     string   `json:"version,omitempty"`
	DownloadURL string   `json:"downloadURL,omitempty"`
	SHA256      string   `json:"sha256,omitempty"`
	Args        []string `json:"args,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMMachine is the Schema for the lcmmachines API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="clusterName",type="string",JSONPath=".spec.clusterName",description="Cluster Name",priority=0
// +kubebuilder:printcolumn:name="type",type="string",JSONPath=".spec.type",description="Type",priority=0
// +kubebuilder:printcolumn:name="state",type="string",JSONPath=".status.state",description="State",priority=0
// +kubebuilder:printcolumn:name="internalIP",type="string",JSONPath=".status.addresses[?(@.type==\"InternalIP\")].address",description="Internal IP",priority=1
// +kubebuilder:printcolumn:name="hostname",type="string",JSONPath=".status.addresses[?(@.type==\"Hostname\")].address",description="Hostname",priority=1
// +kubebuilder:printcolumn:name="agentVersion",type="string",JSONPath=".status.agentVersion",description="Agent Version",priority=1
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
	// bastion
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
		// An important point here is that "Prepare" state
		// means that state items for Deploy/Reconfigure phase
		// aren't set yet. If state items are set for the
		// Deploy phase, the state will not be Prepare even if
		// the Prepare items aren't ready yet.
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

func (m *LCMMachine) IsBYO() bool {
	return m.Spec.LCMType == LCMTypeBYO
}

func (m *LCMMachine) IsMKE() bool {
	return m.Spec.LCMType == LCMTypeMKE
}

func (m *LCMMachine) IsK0s() bool {
	return m.Spec.LCMType == LCMTypeK0s
}

// ProxyStateItemParamsWithKeys gets proxy related Params and keys map from LCMStateItem list
// proxy params are present in all LCMStateItem items in current implementation, so just taking first
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
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMMachineList contains a list of LCMMachine objects
type LCMMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LCMMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LCMMachine{}, &LCMMachineList{})
}
