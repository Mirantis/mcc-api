package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OperationRequestType string

const (
	OsdRemoveOperation   OperationRequestType = "osdRemove"
	PerfTestingOperation OperationRequestType = "perfTest"
)

// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`,description="Phase"
// +kubebuilder:printcolumn:name="Approve",type=boolean,JSONPath=`.spec.approve`,description="Approve"
// +kubebuilder:resource:path=cephosdremoverequests,scope=Namespaced
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CephOsdRemoveRequest stands for handling requests for removing osds from cluster
type CephOsdRemoveRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   *CephOsdRemoveRequestSpec   `json:"spec,omitempty"`
	Status *CephOsdRemoveRequestStatus `json:"status,omitempty"`
}

// CephOsdRemoveRequestSpec contains approval flag,
// map of nodes with osd to-remove list
// and whether to keep request in queue on fail
type CephOsdRemoveRequestSpec struct {
	// Nodes is a map of nodes, which contains specification how osds
	// should be removed: by devices or osd ids
	// Optional
	Nodes map[string]NodeCleanUpSpec `json:"nodes,omitempty"`
	// Approve is a ceph team emergency break to ask operator to
	// think twice before removing OSD. Could be only manually be
	// enabled by user.
	// Optional
	Approve bool `json:"approve,omitempty"`
	// KeepOnFail is used to keep request in queue when validation
	// or processing phases are failed and do not move to next request
	// until flag/request itself removed or Nodes section is updated
	// Optional
	KeepOnFail bool `json:"keepOnFail,omitempty"`
	// Resolved allows to keep request in history when it is failed and
	// do not block MiraCeph reconciling.
	// Optional
	Resolved bool `json:"resolved,omitempty"`
	// Resume failed request, which has set KeepOnFail and InputWaiting
	// status, to continue osd remove actions
	// Optional
	ResumeFailed bool `json:"resumeFailed,omitempty"`
}

// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1

// NodeCleanUpSpec describes how should be OSD cleaned up on particular node
// Can be set only one field at time.
type NodeCleanUpSpec struct {
	// CompleteCleanUp is a flag for total node cleanup and drop from crush map
	// Node will be cleaned up with all its osd/devices if possible
	// Optional
	CompleteCleanUp bool `json:"completeCleanUp,omitempty"`
	// Same as CompleteCleanUp, but without devices cleanup on host
	// May be useful when host is going to be reprovisioned and
	// no need to spent time for devices clean up
	// Optional
	DropFromClusterOnly bool `json:"dropFromClusterOnly,omitempty"`
	// CleanUpByDevice describes devices or it pathes to cleanup
	// Optional
	// +kubebuilder:validation:MinItems:=1
	CleanUpByDevice []DeviceCleanUpSpec `json:"cleanupByDevice,omitempty"`
	// CleanUpByOsdID is a list of Osds, placed on node to cleanup, can be omitted.
	// Optional
	// +kubebuilder:validation:MinItems:=1
	CleanUpByOsdID []int `json:"cleanupByOsdId,omitempty"`
}

// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1

// DeviceCleanUpSpec is a spec describing dev names or pathes to cleanup
// If disk contain partition of some osd in use it will be untouched
type DeviceCleanUpSpec struct {
	// Name represents physical dev names on a node, used for osd, e.g. 'sdb', 'nvme1e0'
	// +kubebuilder:validation:Pattern:=`^[\w]+$`
	// Optional
	Name string `json:"name,omitempty"`
	// Path is a full dev path (by-path or by-id) on a node,
	// where osd lives, e.g. '/dev/disk/by-path/...' or '/dev/disk/by-id/...'
	// +kubebuilder:validation:Pattern:=`^\/dev\/disk\/by-(path|id)\/.+`
	// Optional
	Path string `json:"path,omitempty"`
}

// HandleRequestPhase is a enum for all supported
// handle request phases
type HandleRequestPhase string

// Phases are moving in next order:
// Pending -> Validating -> ApproveWaiting -> Processing -> Complete
//                   	\-> Failed                      \-> Failed
//
const (
	RequestPhaseApproveWaiting        HandleRequestPhase = "ApproveWaiting"
	RequestPhaseCompleted             HandleRequestPhase = "Completed"
	RequestPhaseCompletedWithWarnings HandleRequestPhase = "CompletedWithWarnings"
	RequestPhaseFailed                HandleRequestPhase = "Failed"
	RequestPhaseInputWaiting          HandleRequestPhase = "InputWaiting"
	RequestPhasePending               HandleRequestPhase = "Pending"
	RequestPhaseProcessing            HandleRequestPhase = "Processing"
	RequestPhaseValidating            HandleRequestPhase = "Validating"
)

// CephOsdRemoveRequestStatus contains status of removing osds process
// and possible info/error messages found on during process
type CephOsdRemoveRequestStatus struct {
	// Phase is a current request phase
	Phase HandleRequestPhase `json:"phase"`
	// RemoveInfo contains map, describing on what is going to be removed
	// in next view: node -> osd ID -> associated devices info,
	// issues found during validation/processing phases
	// and warnings which user should pay attention to
	RemoveInfo *RequestRemoveInfo `json:"removeInfo,omitempty"`
	// Messages is a list of info messages describing what's a reason
	// of moving request to next phase
	Messages []string `json:"messages,omitempty"`
	// Conditions is a history list of changing request itself
	Conditions []CephOsdRemoveRequestCondition `json:"conditions"`
}

type RequestRemoveInfo struct {
	// CleanUpMap is a map of cleanup from host-osdId to device
	// based on this map user will decide whether approve current request or not
	// after that it will contain all remove statuses and errors during remove
	CleanUpMap map[string]HostMapping `json:"cleanUpMap"`
	// Issues found during validation/processing phases, describing occured problem
	Issues []string `json:"issues,omitempty"`
	// Warnings found during validation/processing phases, user attention required
	Warnings []string `json:"warnings,omitempty"`
}

type HostMapping struct {
	// CompleteCleanUp is a flag whether make complete host cleanup from crush map
	CompleteCleanUp bool `json:"completeCleanUp,omitempty"`
	// DropFromClusterOnly is a flag whether make complete host cleanup from
	// crush map, but do not cleanup used devices
	DropFromClusterOnly bool `json:"dropFromClusterOnly,omitempty"`
	// OsdMapping represents a mapping from osdID -> devices, also contains
	// osd remove statuses such as osd remove itself, deployment remove,
	// device clean up job
	OsdMapping map[string]OsdMapping `json:"osdMapping"`
	// NodeIsDown indicates host availability
	NodeIsDown bool `json:"nodeIsDown,omitempty"`
	// VolumesInfoMissed indicates volume info availability for host
	VolumesInfoMissed bool `json:"volumeInfoMissed,omitempty"`
	// HostRemoveStatus represents host remove status, if node marked for complete clean up
	HostRemoveStatus *RemoveStatus `json:"hostRemoveStatus,omitempty"`
}

type OsdMapping struct {
	// DeviceMapping is a mapping device -> device info, with short device info, such
	// as path, class, partition, etc
	DeviceMapping map[string]DeviceInfo `json:"deviceMapping"`
	// RemoveStatus describing current phase and errors if happened
	// for osd,deployment or device clean up
	RemoveStatus *RemoveResult `json:"removeStatus,omitempty"`
}

// DeviceInfo represents short device info which provide all
// needed info for clean up procedure
type DeviceInfo struct {
	// Class is a device class: hdd, ssd
	Class string `json:"deviceClass,omitempty"`
	// Path is a full device path by-path
	Path string `json:"devicePath,omitempty"`
	// ID is a device id, used in by-id
	ID string `json:"deviceID,omitempty"`
	// Partition used for current OsdID on disk
	Partition string `json:"usedPartition,omitempty"`
	// Type is a purpose of device: block or db
	Type string `json:"devicePurpose,omitempty"`
	// ZapDisk is a flag whether to zap disk at all or not
	ZapDisk bool `json:"zapDisk,omitempty"`
	// device availability marker, if not available device clean up job skipped
	NotAvailable bool `json:"notAvailable,omitempty"`
}

// RemoveStatusPhase is a enum for handling remove
// during processing phase
type RemoveStatusPhase string

const (
	RemovePending          RemoveStatusPhase = "Pending"
	RemoveWaitingRebalance RemoveStatusPhase = "Rebalancing"
	RemoveInProgress       RemoveStatusPhase = "Removing"
	RemoveCompleted        RemoveStatusPhase = "Completed"
	RemoveFinished         RemoveStatusPhase = "Removed"
	RemoveFailed           RemoveStatusPhase = "Failed"
	RemoveSkipped          RemoveStatusPhase = "Skipped"
)

// RemoveResult keeps all osd remove relatd statuses in one place
type RemoveResult struct {
	// OsdRemoveStatus represents Ceph OSD remove status itself
	OsdRemoveStatus *RemoveStatus `json:"osdRemoveStatus,omitempty"`
	// DeployRemoveStatus is a deployment status related to Ceph OSD
	DeployRemoveStatus *RemoveStatus `json:"deploymentRemoveStatus,omitempty"`
	// DeviceCleanUpJob represents clean up job for related Ceph OSD and it's associated devices
	DeviceCleanUpJob *RemoveStatus `json:"deviceCleanUpJob,omitempty"`
}

// RemoveStatus handling status description
type RemoveStatus struct {
	// Status is a current remove status
	Status RemoveStatusPhase `json:"status"`
	// Name is an object name for handling, optional
	Name string `json:"name,omitempty"`
	// Error faced during handling
	Error string `json:"errorReason,omitempty"`
}

// CephOsdRemoveRequestCondition contains history of changes/updates
// for request
type CephOsdRemoveRequestCondition struct {
	// Timestamp is a timestamp when this condition appeared
	Timestamp string `json:"timestamp"`
	// Phase is a current request handling phase
	Phase HandleRequestPhase `json:"phase"`
	// Nodes is a mapping of nodes within their Ceph OSDs / devices to clean up
	Nodes map[string]NodeCleanUpSpec `json:"nodes,omitempty"`
	// MiraCephVersion is a version of miraCeph used for that
	// condition in format <generation>-<resourceVersion>
	MiraCephVersion *MiraCephSpecVersion `json:"miraCephVersion,omitempty"`
}

type MiraCephSpecVersion struct {
	// ResourceVersion is a MiraCeph resource version
	ResourceVersion string `json:"miraCephResourceVersion"`
	// Generation is a MiraCeph generation ID
	Generation int64 `json:"miraCephGeneration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CephOsdRemoveRequestList contains a list of CephOsdRemoveRequest objects
type CephOsdRemoveRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items contains a list of CephOsdRemoveRequest objects
	Items []CephOsdRemoveRequest `json:"items"`
}

// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`,description="Phase"
// +kubebuilder:printcolumn:name="Start time",type=string,JSONPath=`.status.lastStartTime`,description="Last job start time"
// +kubebuilder:printcolumn:name="Duration",type=string,JSONPath=`.status.lastDurationTime`,description="Last job duration time"
// +kubebuilder:printcolumn:name="Job status",type=string,JSONPath=`.status.lastJobStatus`,description="Last job status"
// +kubebuilder:printcolumn:name="Schedule",type=string,JSONPath=`.spec.periodic.schedule`,description="Schedule"
// +kubebuilder:resource:path=cephperftestrequests,scope=Namespaced
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CephPerfTestRequest stands for running performance testing operations on cluster
type CephPerfTestRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CephPerfTestRequestSpec    `json:"spec"`
	Status *CephPerfTestRequestStatus `json:"status,omitempty"`
}

// CephPerfTestRequestSpec contains perf test run parameters,
// image name to use to run test and periodic schedule to run
// perf test as cron job
type CephPerfTestRequestSpec struct {
	// Parameters are arguments to the entrypoint.
	// Parameters is a list, which takes the command arguments for
	// a performance test execution.
	Parameters []string `json:"parameters"`
	// Entrypoint command to run performance test in container.
	// If performance image is updated, you may update command as well.
	// By default equals image entrypoint.
	// Optional
	Command []string `json:"command,omitempty"`
	// Image name which will be used for running a performance test.
	// By default used image 'vineethac/fio_image'.
	// Default fio image is recommended as it has support of multitude
	// io engines https://fio.readthedocs.io/en/latest/fio_doc.html#i-o-engine
	// Optional
	Image string `json:"image,omitempty"`
	// Periodic allows to configure perf test run as periodic job
	// Optional
	Periodic *CephPerfTestPeriodicParams `json:"periodic,omitempty"`
}

// CephPerfTestPeriodicParams contains schedule for running perf test
type CephPerfTestPeriodicParams struct {
	// Schedule in base cron format, e.g '* * 30 2 \*'
	// see https://en.wikipedia.org/wiki/Cron
	Schedule string `json:"schedule"`
	// Suspend executing cron job
	// Optional
	Suspended bool `json:"suspended,omitempty"`
	// Number of runs and results to keep in history
	// Default is 5
	// Optional
	RunsToKeep int32 `json:"runsToKeep,omitempty"`
}

// PerfTestStatusPhase is a enum for handling perf test phases
type PerfTestStatusPhase string

const (
	PerfTestPending    PerfTestStatusPhase = "Pending"
	PerfTestInProgress PerfTestStatusPhase = "Running"
	PerfTestFinished   PerfTestStatusPhase = "Finished"
	PerfTestFailed     PerfTestStatusPhase = "Failed"
	PerfTestSuspended  PerfTestStatusPhase = "Suspended"
	PerfTestScheduling PerfTestStatusPhase = "Scheduling"
	PerfTestWaitingRun PerfTestStatusPhase = "WaitingNextRun"
)

type CephPerfTestRequestStatus struct {
	// Current phase of perf test run
	Phase PerfTestStatusPhase `json:"phase"`
	// Last time of perf test run start
	LastStartTime string `json:"lastStartTime,omitempty"`
	// Duration time for perf test run, set only when run successfully completed
	LastDurationTime string `json:"lastDurationTime,omitempty"`
	// Last job status
	LastJobStatus string `json:"lastJobStatus,omitempty"`
	// Any issues, warnings found during processing, executing perf test run
	Messages []string `json:"messages,omitempty"`
	// Perf test results or where to find them
	Results *CephPerfTestResults `json:"results,omitempty"`
	// History of statuses and timings for cron jobs
	StatusHistory []CephPerfTestCronHistory `json:"statusHistory,omitempty"`
}

// CephPerfTestResults stands for showing perf test results
// or where to find them
type CephPerfTestResults struct {
	// StoredOnPvc means that perf test results can be found on
	// pvc with 'pvcName' in 'pvcNamespace'
	StoredOnPvc *PvcStoredResults `json:"storedOnPvc,omitempty"`
}

// PvcStoredResults means that perf test results can be found on
// pvc with 'pvcName' in 'pvcNamespace'
type PvcStoredResults struct {
	PvcName      string `json:"pvcName"`
	PvcNamespace string `json:"pvcNamespace"`
}

type CephPerfTestCronHistory struct {
	StartTime    string   `json:"startTime"`
	JobStatus    string   `json:"jobStatus"`
	DurationTime string   `json:"durationTime,omitempty"`
	Messages     []string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CephPerfTestRequestList contains a list of CephPerfTestRequest objects
type CephPerfTestRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items contains a list of CephPerfTestRequest objects
	Items []CephPerfTestRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CephOsdRemoveRequest{}, &CephPerfTestRequest{}, &CephOsdRemoveRequestList{}, &CephPerfTestRequestList{})
}
