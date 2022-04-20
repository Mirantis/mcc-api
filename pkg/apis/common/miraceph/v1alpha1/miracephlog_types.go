package v1alpha1

import (
	bktv1alpha1 "github.com/kube-object-storage/lib-bucket-provisioner/pkg/apis/objectbucket.io/v1alpha1"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:resource:path=miracephlogs,scope=Namespaced
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MiraCephLog aggregates all Ceph cluster statuses
type MiraCephLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// MiraCephClusterStatus represents overall Ceph cluster status info
	MiraCephClusterStatus *MiraCephClusterStatus `json:"fullClusterStatus,omitempty"`
	// Status represents ceph-status controller health state
	Status MiraCephLogStatus `json:"status,omitempty"`
}

// MiraCephClusterStatus represents overall Ceph cluster status info
type MiraCephClusterStatus struct {
	// CephClusterStatus contains common ceph cluster status information
	CephClusterStatus cephv1.ClusterStatus `json:"clusterStatus"`
	// CephDaemonsStatus contains Ceph daemon's overall information
	CephDaemonsStatus map[string]CephDaemonsStatus `json:"daemonsStatus,omitempty"`
	// RookOperatorStatus contains Rook operator status
	RookOperatorStatus string `json:"operatorStatus,omitempty"`
	// BlockStorageStatus contains Block storage status information
	BlockStorageStatus BlockStorageStatus `json:"blockStorageStatus,omitempty"`
	// ObjectStorageStatus contains Object storage status information
	ObjectStorageStatus *ObjectStorageStatus `json:"objectStorageStatus,omitempty"`
	// SharedFilesystemStatus contains shared filesystem status information
	SharedFilesystemStatus *SharedFilesystemStatus `json:"sharedFilesystemStatus,omitempty"`
	// CephDetails contains additional Ceph cluster information, such as disk usage, device mapping for osds
	CephDetails CephDetails `json:"cephDetails,omitempty"`
	// CephCSIPluginsStatus contains Ceph CSI plugins status
	CephCSIPluginDaemonsStatus map[string]CephDaemonsStatus `json:"cephCSIPluginDaemonsStatus,omitempty"`
}

// CephDaemonsStatus contains status info for the defined Ceph Daemon
type CephDaemonsStatus struct {
	// Issues represents found Ceph daemon issues, otherwise it is empty
	Issues string `json:"issues"`
	// Status contains human-readable information about expected and current
	// number of Ceph Daemons running
	Status string `json:"status"`
}

// BlockStorageStatus contains Block storage status information
type BlockStorageStatus struct {
	// PoolsStatus represents a key-value mapping with Ceph Pool's name
	// as a key and it's status as a value
	PoolsStatus map[string]MiraPoolStatus `json:"poolsStatus,omitempty"`
	// CephClientsStatus represents  a key-value mapping of described
	// in spec (or default OpenStack) Ceph Clients in Ceph Cluster
	CephClientsStatus map[string]MiraCephClientStatus `json:"clientsStatus,omitempty"`
}

// MiraPoolStatus represents Ceph Pool's status
type MiraPoolStatus struct {
	// Present is a flag whether pool is present on env or not
	Present bool `json:"present"`
	// Status contains info about current phase and mirroring/snapshots as well
	Status *cephv1.CephBlockPoolStatus `json:"status,omitempty"`
}

// MiraCephClientStatus represents Ceph Client's status
type MiraCephClientStatus struct {
	// Present is a flag whether client is present on env or not
	Present bool `json:"present"`
	// Phase is a current client status phase
	Phase cephv1.ConditionType `json:"status,omitempty"`
}

// ObjectStorageStatus contains Object storage status information
type ObjectStorageStatus struct {
	// ObjectStoreStatus represents Ceph Object storage main status info
	ObjectStoreStatus *cephv1.ObjectStoreStatus `json:"objectStoreStatus"`
	// ObjectStoreUsersStatus represents status info for defined Ceph Object storage users
	ObjectStoreUsersStatus map[string]MiraObjectStoreUserStatus `json:"objectStoreUsers,omitempty"`
	// ObjectStoreBucketsStatus represents status info for defined Ceph Object storage buckets
	ObjectStoreBucketsStatus map[string]MiraObjectStoreBucketStatus `json:"objectStoreBuckets,omitempty"`
	// ObjectStorePublicEndpoint represents external endpoint to access object storage
	ObjectStorePublicEndpoint string `json:"objectStorePublicEndpoint,omitempty"`
}

// MiraObjectStoreUserStatus represents Ceph Object storage user status info
type MiraObjectStoreUserStatus struct {
	// Present is a flag whether RGW user is present on env or not
	Present bool `json:"present"`
	// Phase is a current RGW user status phase
	Phase string `json:"phase,omitempty"`
}

// MiraObjectStoreBucketStatus represents Ceph Object storage bucket status info
type MiraObjectStoreBucketStatus struct {
	// Present is a flag whether RGW bucket is present on env or not
	Present bool `json:"present"`
	// Phase is a current RGW bucket status phase
	Phase bktv1alpha1.ObjectBucketClaimStatusPhase `json:"phase,omitempty"`
}

// SharedFilesystemStatus contains shared filesystem status information
type SharedFilesystemStatus struct {
	// CephFsStatus contains status information for CephFs
	CephFsStatus map[string]MiraCephFsStatus `json:"cephFsStatus,omitempty"`
}

type MiraCephFsStatus struct {
	// Present is a flag whether CephFS is present on env or not
	Present bool `json:"present"`
	// Phase is a current CephFs status phase
	Phase cephv1.ConditionType `json:"status,omitempty"`
	// ActiveClient represents current active clients, using that CephFs
	ActiveClients *int `json:"activeClients,omitempty"`
}

// CephDetails contains additional Ceph cluster information, such as disk usage, device mapping for osds
type CephDetails struct {
	// DiskUsageDetails contains verbose info about usage/capacity cluster per class/pools
	DiskUsageDetails DiskUsageDetails `json:"diskUsage,omitempty"`
	// DeviceMapping contains information on which node runs each osds and what disk it's using
	DeviceMapping map[string]DeviceMapping `json:"deviceMapping,omitempty"`
	// deprecated, CephDeviceMapping contains information on which node runs each osds and what disk it's using
	CephDeviceMapping map[string]OsdDeviceMapping `json:"cephDeviceMapping,omitempty"`
}

type DiskUsageDetails struct {
	ClassesDetail map[string]ClassDiskUsageStats `json:"deviceClass,omitempty"`
	PoolsDetail   map[string]PoolDiskUsageStats  `json:"pools,omitempty"`
}

type ClassDiskUsageStats struct {
	UsedBytes      uint64 `json:"bytesUsed,omitempty"`
	AvailableBytes uint64 `json:"bytesAvailable,omitempty"`
	TotalBytes     uint64 `json:"bytesTotal,omitempty"`
}

type PoolDiskUsageStats struct {
	UsedBytes           uint64 `json:"bytesUsed,omitempty"`
	UsedBytesPercentage string `json:"usedPercentage,omitempty"`
	AvailableBytes      uint64 `json:"bytesAvailable,omitempty"`
	TotalBytes          uint64 `json:"bytesTotal,omitempty"`
}

type OsdDetails struct {
	DeviceName string `json:"device,omitempty"`
	Status     string `json:"status,omitempty"`
}

type DeviceMapping map[string]OsdDetails

// OsdDeviceMapping deprecated
type OsdDeviceMapping map[string]string

type MiraLogState string

const (
	LogStateReady  MiraLogState = "Ready"
	LogStateFailed MiraLogState = "Failed"
)

// MiraCephLogStatus defines the observed state of MiraCephLog
type MiraCephLogStatus struct {
	// State represents the state for overall status
	State MiraLogState `json:"state"`
	// LastCheck is a last time when cluster was verified
	LastCheck string `json:"lastCheck"`
	// LastUpdate is a last time when MiraCephLog was updated
	LastUpdate string `json:"lastUpdate,omitempty"`
	// Messages is a list with any possible error/warning messages
	Messages []string `json:"messages,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MiraCephLogList represents a list of MiraCephLog objects
type MiraCephLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items represents a list of MiraCephLog objects
	Items []MiraCephLog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MiraCephLog{}, &MiraCephLogList{})
}
