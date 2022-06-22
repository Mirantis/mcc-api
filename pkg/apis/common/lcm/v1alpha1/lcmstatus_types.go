/*
Copyright 2022 The Mirantis Inc.

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
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StageName string

const (
	AgentUpgraded        StageName = "Agent upgraded"
	PreparePhaseDone     StageName = "Prepare phase done"
	DeployPhaseDone      StageName = "Deploy phase done"
	ReconfigurePhaseDone StageName = "Reconfigure phase done"
	RebootDone           StageName = "Reboot done"

	ControlPlaneMKEUpgraded StageName = "Control plane nodes MKE upgraded"
	WorkersMKEUpgraded      StageName = "Worker nodes MKE upgraded"

	NodeMaintenanceRequested  StageName = "Node maintenance requested"
	NodeWorkloadLocksInactive StageName = "Node workload locks inactive"
	KubernetesDrained         StageName = "Kubernetes drained"
	SwarmDrained              StageName = "Swarm drained"
	SwarmUncordoned           StageName = "Swarm uncordoned"
	KubernetesUncordoned      StageName = "Kubernetes uncordoned"
	NodeMaintenanceRemoved    StageName = "Node maintenance request removed"

	ClusterMaintenaceRequested   StageName = "Requested cluster maintenance"
	ClusterWorkloadLocksInactive StageName = "Cluster workload locks are inactive"
	ClusterMaintenanceRemoved    StageName = "Cluster maintenance request removed"
)

var (
	StageNameByPhase = map[LCMItemPhase]StageName{
		LCMPreparePhase:     PreparePhaseDone,
		LCMDeployPhase:      DeployPhaseDone,
		LCMReconfigurePhase: ReconfigurePhaseDone,
	}
)

type Stage struct {
	Name      StageName   `json:"name"`
	Timestamp metav1.Time `json:"timestamp,omitempty"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
}

func (s *Stage) SetStatus(err error) {
	if s.Success {
		// If this stage is completed already, it's a noop
		return
	}
	s.Timestamp = metav1.Now()
	if err == nil {
		s.Success = true
		s.Message = ""
	} else {
		s.Success = false
		s.Message = err.Error()
	}
}

func (s *Stage) Copy(updated Stage) {
	s.Timestamp = updated.Timestamp
	s.Success = updated.Success
	s.Message = updated.Message
}

type StageList []Stage

func (sl StageList) SetStatus(name StageName, err error) {
	var stage *Stage
	for i := range sl {
		if sl[i].Name == name {
			stage = &sl[i]
		}
	}
	if stage == nil {
		// This means programmatic error, required stage has not been added to ClusterDeploymentStatus
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMClusterUpgradeStatus is the Schema for the lcmclusterupgradestatuses API
// +k8s:openapi-gen=true
type LCMClusterUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// initial MCC release
	FromRelease string `json:"fromRelease"`
	// target MCC release
	ToRelease string    `json:"toRelease"`
	Stages    StageList `json:"stages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMClusterUpgradeStatusList contains a list of LCMClusterUpgradeStatus
type LCMClusterUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMClusterUpgradeStatus `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMMachineAgentUpgradeStatus is the Schema for the LCMMachineAgentUpgradeStatuses API
// +k8s:openapi-gen=true
type LCMMachineAgentUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Stages StageList `json:"stages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMMachineAgentUpgradeStatusList contains a list of LCMMachineAgentUpgradeStatus
type LCMMachineAgentUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMMachineAgentUpgradeStatus `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMMachineUpgradeDrainStatus is the Schema for the LCMMachineUpgradeDrainStatus API
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMMachineUpgradeDrainStatusList contains a list of LCMMachineUpgradeDrainStatus
type LCMMachineUpgradeDrainStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMMachineUpgradeDrainStatus `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMClusterMKEUpgradeStatus is the Schema for the LCMMachineMKEUpgradeStatuses API
// +k8s:openapi-gen=true
type LCMClusterMKEUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Stages StageList `json:"stages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LCMClusterMKEUpgradeStatusList contains a list of LCMClusterMKEUpgradeStatus
type LCMClusterMKEUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LCMClusterMKEUpgradeStatus `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&LCMClusterUpgradeStatus{}, &LCMClusterUpgradeStatusList{},
		&LCMMachineAgentUpgradeStatus{}, &LCMMachineAgentUpgradeStatusList{},
		&LCMClusterMKEUpgradeStatus{}, &LCMClusterMKEUpgradeStatusList{},
		&LCMMachineUpgradeDrainStatus{}, &LCMMachineUpgradeDrainStatusList{},
	)
}
