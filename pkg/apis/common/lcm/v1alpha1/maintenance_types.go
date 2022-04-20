package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=nodemaintenancerequests,scope=Cluster
type NodeMaintenanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeMaintenanceRequestSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=clustermaintenancerequests,scope=Cluster
type ClusterMaintenanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterMaintenanceRequestSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=nodeworkloadlocks,scope=Cluster
// +kubebuilder:subresource:status
type NodeWorkloadLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeWorkloadLockSpec `json:"spec,omitempty"`
	Status WorkloadLockStatus   `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=clusterworkloadlocks,scope=Cluster
// +kubebuilder:subresource:status
type ClusterWorkloadLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterWorkloadLockSpec `json:"spec,omitempty"`
	Status WorkloadLockStatus      `json:"status,omitempty"`
}

type MaintenanceRequestState string
type WorkloadState string
type MaintenanceRequestScope string

const (
	// Lock object is inactive when workload is turned off and it's safe
	// to put the node offline.
	WorkloadLockInactive WorkloadState = "inactive"

	// Node is not ready to maintenance when lock state is "active".
	WorkloadLockActive WorkloadState = "active"

	// Workload controller must set lock to "failed" when it is unable
	// to gracefully shutdown workload. Also fill NodeWorkloadLock.Status.ErrorMessage
	// with corresponding error message.
	WorkloadLockFailed WorkloadState = "failed"

	// MaintenanceScopeDrain is set when a node is scheduled to be drained.
	MaintenanceScopeDrain MaintenanceRequestScope = "drain"
	// MaintenanceScopeOS is set when a node is scheduled for maintenance via maintenance api
	// or will be rebooted after upgrade.
	// MaintenanceScopeOS includes MaintenanceScopeDrain.
	MaintenanceScopeOS MaintenanceRequestScope = "os"
)

type NodeMaintenanceRequestSpec struct {
	NodeName string `json:"nodeName"`
	// +kubebuilder:validation:Enum=os;drain
	Scope MaintenanceRequestScope `json:"scope,omitempty"`
}

type ClusterMaintenanceRequestSpec struct {
	// +kubebuilder:validation:Enum=os;drain
	Scope MaintenanceRequestScope `json:"scope,omitempty"`
}

type NodeWorkloadLockSpec struct {
	ClusterWorkloadLockSpec `json:",inline"`
	NodeName                string `json:"nodeName"`
}

type WorkloadLockStatus struct {
	// +kubebuilder:validation:Enum=inactive;active;failed
	State        WorkloadState `json:"state"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	Release      string        `json:"release,omitempty"`
}

type ClusterWorkloadLockSpec struct {
	ControllerName string `json:"controllerName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeMaintenanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeMaintenanceRequest `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterMaintenanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMaintenanceRequest `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeWorkloadLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeWorkloadLock `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterWorkloadLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterWorkloadLock `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeMaintenanceRequest{}, &NodeMaintenanceRequestList{}, &ClusterMaintenanceRequest{}, &ClusterMaintenanceRequestList{}, &NodeWorkloadLock{}, &NodeWorkloadLockList{}, &ClusterWorkloadLock{}, &ClusterWorkloadLockList{})
}
