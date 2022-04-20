package v1alpha1

import (
	v1 "k8s.io/api/core/v1"

	miracephv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/common/miraceph/v1alpha1"
)

// +k8s:deepcopy-gen=true

// CephClusterStatus is a common status of Ceph cluster and MiraCeph object
type CephClusterStatus struct {
	// MiraCephInfo describes status of validation and reconcile of MiraCeph
	// resource placed on a managed cluster
	MiraCephInfo miracephv1alpha1.MiraCephStatus `json:"miraCephInfo"`
	// ShortClusterInfo describes the summary of Ceph component's health. Used
	// in Cluster resource status conditions
	ShortClusterInfo *miracephv1alpha1.MiraCephLogStatus `json:"shortClusterInfo,omitempty"`
	// FullClusterInfo describes the overall status of Ceph Cluster and all it's
	// daemons within healthchecks results
	FullClusterInfo *miracephv1alpha1.MiraCephClusterStatus `json:"fullClusterInfo,omitempty"`
}

// +k8s:deepcopy-gen=true

// CephClusterSpec is the Schema for the KaaS Public API
type CephClusterSpec struct {
	// Version is a version of Ceph itself. Chart's default used if not set
	Version string `json:"version,omitempty"`
	// FailureDomain is a general fault-tolerance level for Ceph.
	// Currently is not used at all
	FailureDomain string `json:"failureDomain,omitempty"`
	// ExtraLogging displays verbose difference of all API changes in controller's logs
	ExtraLogging bool `json:"extraLogging,omitempty"`
	// Network is a section which defines the specific network range(s)
	// for Ceph daemons to communicate with each other and the an external
	// connections
	Network miracephv1alpha1.MiraCephNetworkSpec `json:"network,omitempty"`
	// Mgr contains a list of Ceph Manager modules to enable in Ceph Cluster
	Mgr *miracephv1alpha1.Mgr `json:"mgr,omitempty"`
	// Pools is a list of Ceph RBD Pools configurations
	Pools []miracephv1alpha1.MiraCephPool `json:"pools,omitempty"`
	// Clients is a list of Ceph Clients used for Ceph Cluster connection by
	// consumer services
	Clients []miracephv1alpha1.MiraCephClient `json:"clients,omitempty"`
	// Nodes contains full cluster nodes configuration to use as Ceph Nodes
	Nodes map[string]CephNodeReduced `json:"nodes,omitempty"`
	// NodeGroups contains ceph nodes configuration which could be applied for a group of nodes
	NodeGroups map[string]CephNodeGroup `json:"nodeGroups,omitempty"`
	// ObjectStorage contains full RadosGW Object Storage configurations: RGW itself
	// and RGW multisite feature
	ObjectStorage *miracephv1alpha1.MiraCephObjectStorage `json:"objectStorage,omitempty"`
	// External enables usage of external Ceph Cluster connected to internal
	// Container Cloud cluster instead of local Ceph Cluster
	External miracephv1alpha1.MiraCephExternalClusterSpec `json:"external,omitempty"`
	// Ingress provides ability to configure custom ingress rule for an external
	// access to Ceph Cluster resources, for example, public endpoint
	// for Ceph Object Store access
	Ingress *miracephv1alpha1.MiraIngress `json:"ingress,omitempty"`
	// HyperConverge provides an ability to configure resources requests and limitations
	// for Ceph Daemons. Also provides an ability to spawn those Ceph Daemons on a tainted
	// nodes
	HyperConverge *miracephv1alpha1.MiraCephHyperConverge `json:"hyperconverge,omitempty"`
	// RBDMirror allows to configure RBD mirroring between two Ceph Clusters
	RBDMirror *miracephv1alpha1.MiraRBDMirrorSpec `json:"rbdMirror,omitempty"`
	// DisableOsKeys disables automatic generating of openstack-ceph-keys secret.
	// Valuable only for MOS managed clusters
	DisableOsKeys bool `json:"disableOsSharedKeys,omitempty"`
	// SharedFilesystem enables such system as CephFS
	SharedFilesystem *miracephv1alpha1.MiraCephSharedFilesystem `json:"sharedFilesystem,omitempty"`
	// RookConfig is a key-value mapping which contains ceph config keys with a specified
	// values
	RookConfig map[string]string `json:"rookConfig,omitempty"`
	// Deprecated, all maintenance happens automatically
	Maintenance bool `json:"maintenance,omitempty"`
	// Deprecated, use ObjectStorage.Rgw instead
	RGW *miracephv1alpha1.MiraCephRGW `json:"rgw,omitempty"`
	// Deprecated, create CephOsdRemoveRequest instead
	ManageOsds *bool `json:"manageOsds,omitempty"`
}

// +k8s:deepcopy-gen=true

// CephNodeReduced is the node definition for the KaaS Public API
type CephNodeReduced struct {
	// StorageDevices is a list of the defined node's devices to use as Ceph OSD
	StorageDevices []CephStorageDeviceReduced `json:"storageDevices,omitempty"`
	// Roles is a list of control daemons to spawn on the defined node: Ceph Monitor,
	// Ceph Manager and/or Ceph RadosGW daemons. Possible values are: mon, mgr, rgw
	Roles []string `json:"roles,omitempty"`
	// Crush represents Ceph crush topology rules to apply on
	// the defined node
	Crush map[string]string `json:"crush,omitempty"`
	// Resources represents kubernetes resource requirements for the defined node
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
}

// +k8s:deepcopy-gen=true

// CephNodeGroup is the node groups definition for the KaaS Public API
type CephNodeGroup struct {
	// Spec is the node definition which could be applied on a group of nodes
	Spec CephNodeReduced `json:"spec"`
	// Nodes is a list of MCC Machines names where Spec could be applied
	Nodes []string `json:"nodes,omitempty"`
	// Label is a valid kubernetes label selector which defines the nodes
	// where Spec could be applied
	Label string `json:"label,omitempty"`
}

// +k8s:deepcopy-gen=true

// CephStorageDeviceReduced is the device definition for the KaaS Public API
type CephStorageDeviceReduced struct {
	// Name is a device name, for example 'sdc'
	Name string `json:"name,omitempty"`
	// FullPath is a device by-path definition
	FullPath string `json:"fullPath,omitempty"`
	// Config is a key-value mapping contains specific configurations
	// for the defined device, for example 'deviceClass' or 'metadataDevice'
	Config map[string]string `json:"config"`
}

// +k8s:deepcopy-gen=true

// CephOperationRequestSpec is the definition of operation for Ceph cluster
type CephOperationRequestSpec struct {
	// OsdRemove is a definition for osd removal
	OsdRemove *miracephv1alpha1.CephOsdRemoveRequestSpec `json:"osdRemove,omitempty"`
}

// +k8s:deepcopy-gen=true

// CephOperationRequestStatus is the status of corresponding operation requets
type CephOperationRequestStatus struct {
	// ChildNodesMapping is a key-value mapping which reflects
	// BareMetal Machine names with their corresponding kubernetes
	// node names
	ChildNodesMapping map[string]string `json:"childNodesMapping"`
	// OsdRemoveStatus is a status of a current remove OSD request
	OsdRemoveStatus *miracephv1alpha1.CephOsdRemoveRequestStatus `json:"osdRemoveStatus,omitempty"`
}
