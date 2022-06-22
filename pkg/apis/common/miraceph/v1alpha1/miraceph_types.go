package v1alpha1

import (
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MiraCephSpec defines the desired configuration of resulting Ceph Cluster
// and all corresponding resources
type MiraCephSpec struct {
	// Network is a section which defines the specific network range(s)
	// for Ceph daemons to communicate with each other and the an external
	// connections
	Network *MiraCephNetworkSpec `json:"network"`
	// RookNamespace stands for a Rook namespace in an observable cluster
	// Equals to 'rook-ceph' by default.
	RookNamespace string `json:"rookNamespace,omitempty"`
	// Version is a version of Ceph itself. Chart's default used if not set
	Version string `json:"version"`
	// ExtraLogging displays verbose difference of all changes in controller's logs.
	// If it is enabled it will log all changes in MiraCeph spec and all corresponding
	// generated resources such as  ingress, labels, secrets and so on
	ExtraLogging bool `json:"extraLogging,omitempty"`
	// Mgr contains a list of Ceph Manager modules to enable in Ceph Cluster
	Mgr *Mgr `json:"mgr,omitempty"`
	// Nodes contains full cluster nodes configuration to use as Ceph Nodes
	Nodes            []*MiraCephNode `json:"nodes"`
	DashboardEnabled bool            `json:"dashboard"`
	// DataDirHostPath is a default hostPath directory where Rook stores all
	// valuable info. Equals to '/var/lib/rook' by default
	DataDirHostPath string `json:"dataDirHostPath,omitempty"`
	// Pools is a list of Ceph RBD Pools configurations
	Pools []*MiraCephPool `json:"pools,omitempty"`
	// Clients is a list of Ceph Clients used for Ceph Cluster connection by
	// consumer services
	Clients []*MiraCephClient `json:"clients,omitempty"`
	// ObjectStorage contains full RadosGW Object Storage configurations: RGW itself
	// and RGW multisite feature
	ObjectStorage *MiraCephObjectStorage `json:"objectStorage,omitempty"`
	// External enables usage of external Ceph Cluster connected to internal
	// Container Cloud cluster instead of local Ceph Cluster
	External MiraCephExternalClusterSpec `json:"external,omitempty"`
	// SharedFilesystem enables such system as CephFS
	SharedFilesystem *MiraCephSharedFilesystem `json:"sharedFilesystem,omitempty"`
	// Ingress provides ability to configure custom ingress rule for an external
	// access to Ceph Cluster resources, for example, public endpoint
	// for Ceph Object Store access
	Ingress *MiraIngress `json:"ingress,omitempty"`
	// HyperConverge provides an ability to configure resources requests and limitations
	// for Ceph Daemons. Also provides an ability to spawn those Ceph Daemons on a tainted
	// nodes
	HyperConverge *MiraCephHyperConverge `json:"hyperconverge,omitempty"`
	// RookConfig is a key-value mapping which contains ceph config keys with a specified
	// values
	RookConfig map[string]string `json:"rookConfig,omitempty"`
	// RBDMirror allows to configure RBD mirroring between two Ceph Clusters
	RBDMirror *MiraRBDMirrorSpec `json:"rbdMirror,omitempty"`
	// DisableOsKeys disables automatic generating of openstack-ceph-keys secret.
	// Valuable only for MOS managed clusters
	DisableOsKeys bool `json:"disableOsSharedKeys,omitempty"`
	// deprecated, use ObjectStore.Rgw section instead
	Rgw *MiraCephRGW `json:"rgw,omitempty"`
	// deprecated, create CephOsdRemoveRequest instead
	ManageOsds *bool `json:"manageOsds,omitempty"`
}

// MiraCephExternalClusterSpec represents external Ceph Cluster
// connected to internal Container Cloud cluster instead of local
// Ceph Cluster
type MiraCephExternalClusterSpec struct {
	// Enable flag allows to enable external Ceph Cluster support
	Enable bool `json:"enable"`
	// FSID represents external Ceph Cluster FSID
	FSID string `json:"fsid"`
	// MonData represents comma-separated list of <monName>=<IP:port> addresses
	// of external Ceph Monitor daemons
	MonData string `json:"monData"`
	// ExternalAdminSecret represents external Ceph Cluster admin keyring
	ExternalAdminSecret string `json:"adminSecret"`
}

type MiraCephSharedFilesystem struct {
	// CephFS to create, for now supported only one filesystem at time.
	CephFS []MiraCephFS `json:"cephFS,omitempty"`
}

type MiraCephFS struct {
	// CephFS name
	Name string `json:"name"`
	// The settings used to create the filesystem metadata pool. Must use replication.
	MetadataPool MiraCephPoolSpec `json:"metadataPool"`
	// The settings to create the filesystem data pools. Must use replication.
	DataPools []MiraCephFsPool `json:"dataPools,omitempty"`
	// When set to ‘true’ the filesystem will remain when the CephFilesystem resource is deleted
	// This is a security measure to avoid loss of data if the CephFilesystem resource is deleted accidentally.
	PreserveFilesystemOnDelete bool `json:"preserveFilesystemOnDelete,omitempty"`
	// Metadata server settings correspond to the MDS daemon settings
	MetadataServer MiraCephMetadataServer `json:"metadataServer"`
	// The settings to create the filesystem data pool. Must use replication.
	// Deprecated, use array definition `dataPools` instead.
	DataPool *MiraCephPoolSpec `json:"dataPool,omitempty"`
}

type MiraCephMetadataServer struct {
	// The number of active MDS instances. As load increases, CephFS will automatically
	// partition the filesystem across the MDS instances. Rook will create double the
	// number of MDS instances as requested by the active count. The extra instances will
	// be in standby mode for failover
	ActiveCount int32 `json:"activeCount"`
	// If true, the extra MDS instances will be in active standby mode and will keep
	// a warm cache of the filesystem metadata for faster failover. The instances will
	// be assigned by CephFS in failover pairs. If false, the extra MDS instances will
	// all be on passive standby mode and will not maintain a warm cache of the metadata.
	ActiveStandby bool `json:"activeStandby,omitempty"`
	// Resources represents kubernetes resource requirements for mds instances
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
}

// Mgr contains a list of Ceph Manager modules to enable in Ceph Cluster
type Mgr struct {
	// Modules is a list of Ceph Manager modules names to enable in Ceph
	Modules []string `json:"modules,omitempty"`
}

type ValidationResult string

const (
	ValidationFailed  ValidationResult = "Failed"
	ValidationSucceed ValidationResult = "Succeed"
)

// MiraValidation reflects validation result for MiraCeph spec
type MiraValidation struct {
	// Result is a spec validation result, which could be Succeed or Failed
	Result ValidationResult `json:"result,omitempty"`
	// Messages represents a list of possible issues or validation messages
	// found during spec validating
	Messages []string `json:"messages,omitempty"`
}

type MiraPhase string

const (
	PhaseCreating    MiraPhase = "Creating"
	PhaseDeploying   MiraPhase = "Deploying"
	PhaseValidation  MiraPhase = "Validation"
	PhaseReady       MiraPhase = "Ready"
	PhaseOnHold      MiraPhase = "OnHold"
	PhaseMaintenance MiraPhase = "Maintenance"
	PhaseDeleting    MiraPhase = "Deleting"
	PhaseFailed      MiraPhase = "Failed"
)

// MiraCephStatus defines the observed state of MiraCeph
type MiraCephStatus struct {
	//+kubebuilder:default=Creating
	// Phase is a current MiraCeph handling phase
	Phase MiraPhase `json:"phase"`
	// Message is a description of a current phase if exists
	Message string `json:"message,omitempty"`
	// Validation reflects validation result for spec
	Validation MiraValidation `json:"validation,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Validation",type=string,JSONPath=`.status.validation.result`,description="Validation status"
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`,description="Deployment phase"
// +kubebuilder:printcolumn:name="Message",type=string,JSONPath=`.status.message`,description="Cluster status message"
// +kubebuilder:resource:path=miracephs,scope=Namespaced
// +kubebuilder:subresource:status
// +genclient

// MiraCeph is the Schema for the miracephs API which contains
// a valid Ceph configuration which is handled by ceph-controller and
// produce all related objects and daemons in Rook (K8S based Ceph)
type MiraCeph struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired configuration of resulting Ceph Cluster
	// and all corresponding resources
	Spec MiraCephSpec `json:"spec,omitempty"`
	// Status represents current status of handling Ceph Cluster configuration
	// by ceph-controller
	Status MiraCephStatus `json:"status,omitempty"`
}

// MiraCephObjectStorage contains full RadosGW Object Storage configurations:
// RGW itself and RGW multisite feature
type MiraCephObjectStorage struct {
	// Rgw represents Ceph RadosGW settings
	Rgw *MiraCephRGW `json:"rgw"`
	// MultiSite represents Ceph RadosGW multisite/multizone feature settings
	MultiSite *MiraCephMultiSite `json:"multiSite,omitempty"`
}

// MiraCephRGW represents Ceph RadosGW settings
type MiraCephRGW struct {
	// Name represents the name of specified object storage
	Name string `json:"name"`
	// Users is a list of user names to create for object storage
	// with radosgw-admin
	Users []string `json:"users,omitempty"`
	// Buckets is a list of initial buckets to create in object storage
	// with radosgw-admin
	Buckets []string `json:"buckets,omitempty"`
	// Replicas is a number of replicas for each Ceph RadosGW instance.
	// Not used in a product currently
	Replicas *int `json:"replicas,omitempty"`
	// MetadataPool represents Ceph Pool's settings which stores RGW metadata.
	// Mutually exclusive with Zone
	MetadataPool MiraCephPoolSpec `json:"metadataPool"`
	// DataPool represents Ceph Pool's settings which stores RGW data.
	// Mutually exclusive with Zone
	DataPool MiraCephPoolSpec `json:"dataPool"`
	// PreservePoolsOnDelete is a flag whether keep RGW metadata/data pools
	// on RGW delete or not
	PreservePoolsOnDelete bool `json:"preservePoolsOnDelete"`
	// Gateway represents Ceph RGW daemons settings
	Gateway MiraGateway `json:"gateway"`
	// SSLCert used for access to RGW Gateway endpoint, if not specified will be generated self-signed
	SSLCert *MiraCert `json:"SSLCert,omitempty"`
	// Zone represents RGW zone if multisite feature enabled
	Zone cephv1.ZoneSpec `json:"zone,omitempty"`
	// HealthCheck represents Ceph RGW daemons healthchecks
	HealthCheck cephv1.BucketHealthCheckSpec `json:"healthCheck,omitempty"`
	// deprecated, use MiraCeph.ingress instead
	Ingress *MiraIngress `json:"ingress,omitempty"`
	// deprecated, use SSLCert instead, now SSL cert for RGW can be specified not only for external endpoint
	ExternalSSL *MiraCert `json:"externalSsl,omitempty"`
}

// MiraCephMultiSite represents Ceph RadosGW multisite/multizone feature settings
type MiraCephMultiSite struct {
	// Realms is a list of Ceph Object storage multisite realms
	Realms []*MiraCephRealm `json:"realms"`
	// ZoneGroups is a list of Ceph Object storage multisite zonegroups
	ZoneGroups []*MiraCephZoneGroup `json:"zoneGroups"`
	// Zones is a list of Ceph Object storage multisite zones
	Zones []*MiraCephZone `json:"zones"`
}

// MiraCephRealm represents RGW multisite realm namespace
type MiraCephRealm struct {
	// Name represents realm's name
	Name string `json:"name"`
	// Pull stands for the Endpoint, the access key and the system key
	// of the system user from the realm being pulled from
	Pull *MiraCephRealmPull `json:"pullEndpoint,omitempty"`
}

// MiraCephRealmPull stands for the Endpoint, the access key and the system key
// of the system user from the realm being pulled from
type MiraCephRealmPull struct {
	// Endpoint represents an endpoint from the master zone in the master zone group
	Endpoint string `json:"endpoint"`
	// AccessKey is an access key of the system user from the realm being pulled from
	AccessKey string `json:"accessKey"`
	// SecretKey is a system key of the system user from the realm being pulled from
	SecretKey string `json:"secretKey"`
}

// MiraCephZoneGroup represents multisite zone group
type MiraCephZoneGroup struct {
	// Name represents zone group's name
	Name string `json:"name"`
	// Realm is a name of the realm for which zone group belongs to
	Realm string `json:"realmName"`
}

// MiraCephZone represents multisite zone
type MiraCephZone struct {
	// Name represents zone's name
	Name string `json:"name"`
	// MetadataPool represents Ceph Pool's setting which contains
	// RGW zone metadata
	MetadataPool MiraCephPoolSpec `json:"metadataPool"`
	// DataPool represents Ceph Pool's setting which contains
	// RGW zone data
	DataPool MiraCephPoolSpec `json:"dataPool"`
	// ZoneGroup is a name of the zone group for which zone belongs to
	ZoneGroup string `json:"zoneGroupName"`
}

// MiraCephHyperConverge represents hyperconverge parameters for Ceph daemons
type MiraCephHyperConverge struct {
	// Resources requirements for ceph daemons, such as: mon, mgr, mds, rgw, osd, osd-hdd, osd-ssd, osd-nvme, prepareosd
	Resources cephv1.ResourceSpec `json:"resources,omitempty"`
	// Tolerations rules for ceph daemons: osd, mon, mgr.
	Tolerations map[string]MiraCephToleration `json:"tolerations,omitempty"`
}

// MiraCephToleration represents kubernetes toleration rules
type MiraCephToleration struct {
	// Rules is a list of kubernetes tolerations defined for some
	// Ceph daemon
	Rules []v1.Toleration `json:"rules"`
}

// MiraIngress provides an ability to configure custom ingress rule for an external
// access to Ceph Cluster resources, for example, public endpoint
// for Ceph Object Store access
type MiraIngress struct {
	// Domain is a public domain used for ingress public endpoint
	Domain string `json:"publicDomain"`

	MiraCert `json:",inline"`

	// CustomIngress represents Extra/Custom Ingress configuration
	CustomIngress *MiraCustomIngress `json:"customIngress,omitempty"`
}

// MiraCustomIngress represents custom Ingress Controller configuration
type MiraCustomIngress struct {
	// ClassName is a name of Ingress Controller class. Default for
	// MOS cloud is 'openstack-ingress-nginx'
	ClassName string `json:"className,omitempty"`
	// Annotations is an extra annotations set to proxy
	Annotations map[string]string `json:"annotations,omitempty"`
}

// MiraCert represents custom certificate settings
type MiraCert struct {
	// Cacert represents CA certificate
	Cacert string `json:"cacert"`
	// TLSCert represents SSL certificate based on the defined Cacert and TLSKey
	TLSCert string `json:"tlsCert"`
	// TLSKey represents SSL secret key used for TLSCert generate
	TLSKey string `json:"tlsKey"`
}

// MiraGateway represents Ceph RGW daemon settings
type MiraGateway struct {
	// Port the rgw service will be listening on (http)
	Port int32 `json:"port"`
	// SecurePort the rgw service will be listening on (https)
	SecurePort int32 `json:"securePort"`
	// Instances is the number of pods in the rgw replicaset.
	// If AllNodes is specified, a daemonset will be created.
	Instances int32 `json:"instances"`
	// AllNodes is a flag whether the rgw pods should be
	// started as a daemonset on all nodes
	AllNodes bool `json:"allNodes"`
	// Resources requirements for RGW instances
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
	// ExternalRgwEndpoints represents external RGW Endpoints
	// to use, when external Ceph cluster is used
	ExternalRgwEndpoints []MiraExternalRgwEndpoint `json:"externalRgwEndpoints,omitempty"`
}

// MiraExternalRgwEndpoint represents external RGW Endpoints
// to use, when external Ceph cluster is used
type MiraExternalRgwEndpoint struct {
	// IP represent external endpoint IP address
	IP string `json:"ip"`
	// Hostname represent name of the host where external endpoint
	// is placed
	Hostname string `json:"hostname,omitempty"`
	// Public is a flag whether external endpoint is public or not
	Public bool `json:"public,omitempty"`
	// FullPath represents full path to the external endpoint. Will be used
	// instead of IP:port if specified
	FullPath string `json:"fullpath,omitempty"`
}

type MiraCephClient struct {
	cephv1.ClientSpec `json:",inline"`
}

// MiraCephNode contains specific node configuration to use it in Ceph Cluster
type MiraCephNode struct {
	// Name represents kubernetes cluster node name
	Name string `json:"name"`
	// Roles is a list of control daemons to spawn on the defined node: Ceph Monitor,
	// Ceph Manager and/or Ceph RadosGW daemons. Possible values are: mon, mgr, rgw
	Roles []string `json:"roles"`
	// Crush represents ceph crush topology rules to apply on
	// the defined node
	Crush map[string]string `json:"crush,omitempty"`
	// Config is a general Ceph Node config to apply for all control/storage daemons
	// on the defined node. Config accepts such keys as deviceClass, metadataDevice,
	// and so no. Full key definitions could be found here:
	// https://github.com/rook/rook/blob/master/Documentation/ceph-cluster-crd.md#osd-configuration-settings
	Config map[string]string `json:"config,omitempty"`
	// NodeGroup is a list of kubernetes node names
	// which allows to specify defined spec to a group of nodes
	// instead of one node defined with Name parameter. Name should be
	// interpreted as a node group name instead of node name if specified
	NodeGroup []string `json:"nodeGroup,omitempty"`
	// NodesByLabel is a valid kubernetes label selector expression
	// which allows to specify defined spec to a group of selected nodes
	// instead of one node defined with Name parameter. Name should be
	// interpreted as a node group name instead of node name if specified
	NodesByLabel string `json:"nodesByLabel,omitempty"`
	// Resources represents kubernetes resource requirements for node or node group
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`

	cephv1.Selection `json:",inline"`
}

// MiraCephFsPool stands for specified CephFS Pool configuration
type MiraCephFsPool struct {
	// Name represents CephFS pool name
	Name string `json:"name"`

	MiraCephPoolSpec `json:",inline"`
}

// MiraCephPool stands for specified Ceph RBD Pool configuration
type MiraCephPool struct {
	// Name represents Ceph RBD pool name
	Name string `json:"name"`
	// UseAsFullName uses Name as a resulting pool name instead of "<Name>-<DeviceClass>"
	UseAsFullName bool `json:"useAsFullName,omitempty"`
	// Role represents pool role. The following values are reserved for
	// MOS managed clusters: vms, images, backup, volumes, volumes-backend
	Role string `json:"role"`
	// Default represents whether Ceph Pool's StorageClass would be default or not
	Default bool `json:"default"`

	MiraCephPoolSpec `json:",inline"`
}

type MiraCephPoolSpec struct {
	// Replicated represents Ceph Pool's replica settings
	Replicated *MiraReplicatedSpec `json:"replicated,omitempty"`
	// FailureDomain represents level of cluster fault-tolerance.
	// Possible values are: osd, host, region or zone if available;
	// technically also any type in the crush map
	FailureDomain string `json:"failureDomain,omitempty"`
	// CrushRoot is the root of the crush hierarchy utilized by the pool
	CrushRoot string `json:"crushRoot,omitempty"`
	// DeviceClass is the device class the OSD should set to (options are: hdd, ssd, or nvme)
	DeviceClass string `json:"deviceClass,omitempty"`
	// ErasureCoded represents Ceph Pool's erasure coding settings
	ErasureCoded *MiraErasureCodedSpec `json:"erasureCoded,omitempty"`
	// Mirroring allows to enable RBD mirroring feature in modes: pool, image
	Mirroring *MiraPoolMirrorSpec `json:"mirroring,omitempty"`
	// AllowVolumeExpansion allows to extend volumes sizes in pool
	AllowVolumeExpansion bool `json:"allowVolumeExpansion,omitempty"`
}

// MiraErasureCodedSpec represents the spec for erasure code in a pool
type MiraErasureCodedSpec struct {
	// CodingChunks is a number of coding chunks per object
	// in an erasure coded storage pool (required for erasure-coded pool type)
	CodingChunks uint `json:"codingChunks"`
	// DataChunks is a number of data chunks per object
	// in an erasure coded storage pool (required for erasure-coded pool type)
	DataChunks uint `json:"dataChunks"`
	// Algorithm represents the algorithm for erasure coding
	Algorithm string `json:"algorithm,omitempty"`
}

// MiraReplicatedSpec represents the spec for replication in a pool
type MiraReplicatedSpec struct {
	// Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)
	Size uint `json:"size"`
	// TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity
	TargetSizeRatio float64 `json:"targetSizeRatio,omitempty"`
}

// MiraPoolMirrorSpec spec represents RBD mirroring
// settings for a specific Ceph RBD Pool
type MiraPoolMirrorSpec struct {
	// Mode - mirroring mode to run
	Mode string `json:"mode"`
}

// MiraRBDMirrorSpec allows to configure RBD mirroring between two Ceph Clusters
type MiraRBDMirrorSpec struct {
	// Count of rbd-mirror daemons to spawn
	Count int `json:"daemonsCount"`

	// Peers is a list of secret's names defined in kubernetes.
	// Currently (Ceph Octopus release) only a single peer is supported
	Peers []MiraRBDSecret `json:"peers,omitempty"`
}

type MiraRBDSecret struct {
	// Site is a name of remote site associated with the token
	Site string `json:"site"`
	// Token represents base64 encoded information about
	// remote cluster; contains fsid,client_id,key,mon_host
	Token string `json:"token"`
	// Pools is a list of Ceph Pools names to mirror
	Pools []string `json:"pools,omitempty"`
}

// MiraCephNetworkSpec is a section which defines the specific network range(s)
// for Ceph daemons to communicate with each other and the an external
// connections
type MiraCephNetworkSpec struct {
	// ClusterNet defines internal network for Ceph Daemons intra-communication
	ClusterNet string `json:"clusterNet"`
	// ClusterNet defines public network for an external access to Ceph Cluster
	PublicNet string `json:"publicNet"`
	// HostNetwork is deprecated field, always true to have persistan mons ips
	HostNetwork bool `json:"hostNetwork,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MiraCephList contains a list of MiraCeph
type MiraCephList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items contains a list of MiraCeph objects
	Items []MiraCeph `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MiraCeph{}, &MiraCephList{})
}
