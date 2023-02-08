package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ClusterMaintenanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMaintenanceRequest `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=clusterworkloadlocks,scope=Cluster
// +kubebuilder:subresource:status
// +gocode:public-api=true
type ClusterWorkloadLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterWorkloadLockSpec `json:"spec,omitempty"`
	Status WorkloadLockStatus      `json:"status,omitempty"`
}
type WorkloadLockStatus struct {
	// +kubebuilder:validation:Enum=inactive;active;failed
	State        WorkloadState `json:"state"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	Release      string        `json:"release,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ClusterWorkloadLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterWorkloadLock `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type NodeMaintenanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeMaintenanceRequest `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=nodemaintenancerequests,scope=Cluster
type NodeMaintenanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeMaintenanceRequestSpec `json:"spec,omitempty"`
}

const (
	// Lock object is inactive when workload is turned off and it's safe
	// to put the node offline.
	// +gocode:public-api=true
	WorkloadLockInactive WorkloadState = "inactive"

	// Node is not ready to maintenance when lock state is "active".
	// +gocode:public-api=true
	WorkloadLockActive WorkloadState = "active"

	// Workload controller must set lock to "failed" when it is unable
	// to gracefully shutdown workload. Also fill NodeWorkloadLock.Status.ErrorMessage
	// with corresponding error message.
	// +gocode:public-api=true
	WorkloadLockFailed WorkloadState = "failed"

	// MaintenanceScopeDrain is set when a node is scheduled to be drained.
	// +gocode:public-api=true
	MaintenanceScopeDrain MaintenanceRequestScope = "drain"
	// MaintenanceScopeOS is set when a node is scheduled for maintenance via maintenance api
	// or will be rebooted after upgrade.
	// MaintenanceScopeOS includes MaintenanceScopeDrain.
	// +gocode:public-api=true
	MaintenanceScopeOS MaintenanceRequestScope = "os"
)

type ClusterMaintenanceRequestSpec struct {
	// +kubebuilder:validation:Enum=os;drain
	Scope MaintenanceRequestScope `json:"scope,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type NodeWorkloadLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeWorkloadLock `json:"items"`
}
type NodeMaintenanceRequestSpec struct {
	NodeName string `json:"nodeName"`
	// +kubebuilder:validation:Enum=os;drain
	Scope MaintenanceRequestScope `json:"scope,omitempty"`
}
type NodeWorkloadLockSpec struct {
	ClusterWorkloadLockSpec      `json:",inline"`
	NodeName                     string `json:"nodeName"`
	NodeDeletionRequestSupported bool   `json:"nodeDeletionRequestSupported,omitempty"`
}

// +gocode:public-api=true
type WorkloadState string

// +gocode:public-api=true
type MaintenanceRequestState string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=clustermaintenancerequests,scope=Cluster
type ClusterMaintenanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterMaintenanceRequestSpec `json:"spec,omitempty"`
}
type ClusterWorkloadLockSpec struct {
	ControllerName string `json:"controllerName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=nodeworkloadlocks,scope=Cluster
// +kubebuilder:subresource:status
// +gocode:public-api=true
type NodeWorkloadLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeWorkloadLockSpec `json:"spec,omitempty"`
	Status WorkloadLockStatus   `json:"status,omitempty"`
}

// +gocode:public-api=true
type MaintenanceRequestScope string

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&NodeMaintenanceRequest{}, &NodeMaintenanceRequestList{}, &ClusterMaintenanceRequest{}, &ClusterMaintenanceRequestList{}, &NodeWorkloadLock{}, &NodeWorkloadLockList{}, &ClusterWorkloadLock{}, &ClusterWorkloadLockList{})
}
