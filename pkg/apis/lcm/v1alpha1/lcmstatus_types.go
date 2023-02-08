package v1alpha1

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// +gocode:public-api=true
	AgentUpgraded StageName = "Agent upgraded"
	// +gocode:public-api=true
	PreparePhaseDone StageName = "Prepare phase done"
	// +gocode:public-api=true
	DeployPhaseWait StageName = "Previous machines are deployed"
	// +gocode:public-api=true
	DeployPhaseDone StageName = "Deploy phase done"
	// +gocode:public-api=true
	ReconfigurePhaseDone StageName = "Reconfigure phase done"
	// +gocode:public-api=true
	RebootDone StageName = "Reboot done"
	// +gocode:public-api=true
	UpgradeDone StageName = "Machine Upgraded"

	// +gocode:public-api=true
	ControlPlaneMKEUpgraded StageName = "Control plane nodes MKE upgraded"
	// +gocode:public-api=true
	WorkersMKEUpgraded StageName = "Worker nodes MKE upgraded"

	// +gocode:public-api=true
	NodeMaintenanceRequested StageName = "Node maintenance requested"
	// +gocode:public-api=true
	NodeWorkloadLocksInactive StageName = "Node workload locks inactive"
	// +gocode:public-api=true
	KubernetesDrained StageName = "Kubernetes drained"
	// +gocode:public-api=true
	SwarmDrained StageName = "Swarm drained"
	// +gocode:public-api=true
	SwarmUncordoned StageName = "Swarm uncordoned"
	// +gocode:public-api=true
	KubernetesUncordoned StageName = "Kubernetes uncordoned"
	// +gocode:public-api=true
	NodeMaintenanceRemoved StageName = "Node maintenance request removed"

	// +gocode:public-api=true
	ClusterMaintenaceRequested StageName = "Requested cluster maintenance"
	// +gocode:public-api=true
	ClusterWorkloadLocksInactive StageName = "Cluster workload locks are inactive"
	// +gocode:public-api=true
	ClusterMachinesUpgradeDone StageName = "All machines of the cluster upgraded"
	// +gocode:public-api=true
	ClusterMaintenanceRemoved StageName = "Cluster maintenance request removed"
)

// LCMClusterUpgradeStatus is the Schema for the lcmclusterupgradestatuses API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type LCMClusterUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// initial MCC release
	FromRelease string `json:"fromRelease"`
	// target MCC release
	ToRelease string    `json:"toRelease"`
	Stages    StageList `json:"stages"`
}

// +gocode:public-api=true
type StageName string

// LCMMachineAgentUpgradeStatusList contains a list of LCMMachineAgentUpgradeStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMMachineAgentUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMMachineAgentUpgradeStatus `json:"items"`
}

// LCMMachineUpgradeDrainStatusList contains a list of LCMMachineUpgradeDrainStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMMachineUpgradeDrainStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMMachineUpgradeDrainStatus `json:"items"`
}

const (
	// +gocode:public-api=true
	StageNotStarted StageStatus = "NotStarted"
	// +gocode:public-api=true
	StageInProgress StageStatus = "InProgress"
	// +gocode:public-api=true
	StageSuccessful StageStatus = "Success"
	// +gocode:public-api=true
	StageFailed StageStatus = "Fail"
)

type StageList []Stage

func (sl StageList) SetStatus(name StageName, err error) {
	var stage *Stage
	for i := range sl {
		if sl[i].Name == name {
			stage = &sl[i]
		}
	}
	if stage == nil {

		panic(errors.Errorf("stage %s not found, can not report its state: %v", name, err))
	}
	stage.SetStatus(err)
}
func (sl StageList) Index(name StageName) int {
	for i := range sl {
		if sl[i].Name == name {
			return i
		}
	}
	return -1
}

// LCMMachineAgentDeploymentStatus is the Schema for the LCMMachineAgentDeploymentStatuses API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type LCMMachineAgentDeploymentStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// MCC release
	Release string    `json:"release"`
	Stages  StageList `json:"stages"`
}

// LCMClusterUpgradeStatusList contains a list of LCMClusterUpgradeStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMClusterUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMClusterUpgradeStatus `json:"items"`
}

// +gocode:public-api=true
type StageStatus string
type Stage struct {
	Name      StageName   `json:"name"`
	Timestamp metav1.Time `json:"timestamp,omitempty"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	// +kubebuilder:default=NotStarted
	// +optional
	Status StageStatus `json:"status,omitempty"`
}

func (s *Stage) SetStatus(err error) {
	if s.Success {

		return
	}
	s.Timestamp = metav1.Now()
	if err == nil {
		s.Success = true
		s.Message = ""
		s.Status = StageSuccessful
	} else {
		s.Success = false
		s.Message = err.Error()
		s.Status = StageInProgress
	}
}
func (s *Stage) Copy(updated Stage) {
	s.Timestamp = updated.Timestamp
	s.Success = updated.Success
	s.Status = updated.Status
	s.Message = updated.Message
}

// LCMMachineAgentDeploymentStatusList contains a list of LCMMachineAgentDeploymentStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMMachineAgentDeploymentStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMMachineAgentDeploymentStatus `json:"items"`
}

// LCMClusterMKEUpgradeStatus is the Schema for the LCMMachineMKEUpgradeStatuses API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type LCMClusterMKEUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Stages StageList `json:"stages"`
}

var (
	// +gocode:public-api=true
	StageNameByPhase = map[LCMItemPhase]StageName{
		LCMPreparePhase:     PreparePhaseDone,
		LCMDeployPhase:      DeployPhaseDone,
		LCMReconfigurePhase: ReconfigurePhaseDone,
	}
)

// LCMClusterMKEUpgradeStatusList contains a list of LCMClusterMKEUpgradeStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LCMClusterMKEUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMClusterMKEUpgradeStatus `json:"items"`
}

// LCMMachineAgentUpgradeStatus is the Schema for the LCMMachineAgentUpgradeStatuses API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type LCMMachineAgentUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Stages StageList `json:"stages"`
}

// LCMMachineUpgradeDrainStatus is the Schema for the LCMMachineUpgradeDrainStatus API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type LCMMachineUpgradeDrainStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Stages StageList `json:"stages"`
}

func (s *LCMMachineUpgradeDrainStatus) InsertStage(index int, stageName StageName) {
	if index == len(s.Stages) {
		s.Stages = append(s.Stages, Stage{Name: stageName})
		return
	}
	var newStages StageList
	for i := range s.Stages {
		if i == index {
			newStages = append(newStages, Stage{Name: stageName})
		}
		newStages = append(newStages, s.Stages[i])
	}
	s.Stages = newStages
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(
		&LCMMachineAgentDeploymentStatus{}, &LCMMachineAgentDeploymentStatusList{},
		&LCMClusterUpgradeStatus{}, &LCMClusterUpgradeStatusList{},
		&LCMMachineAgentUpgradeStatus{}, &LCMMachineAgentUpgradeStatusList{},
		&LCMClusterMKEUpgradeStatus{}, &LCMClusterMKEUpgradeStatusList{},
		&LCMMachineUpgradeDrainStatus{}, &LCMMachineUpgradeDrainStatusList{},
	)
}
