/*
Copyright © 2020 Mirantis

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

package config

import (
	"time"
)

var SupportedProviders = []string{
	"baremetal",
	"vsphere",
	"equinixmetal",
}

const (
	TimeInputFormatRFC3339  = "2006-01-02T15:04:05.99999Z07"
	TimeOutputFormatRFC3339 = "2006-01-02T15:04:05.00000Z07"

	RequiredCRDsWaitTime        time.Duration = 10 * time.Minute
	RequiredCRDsWaitDelay       time.Duration = 5 * time.Second
	RequeueLongWaitTime         time.Duration = 3 * time.Minute
	RequeueTemporaryError       time.Duration = 17 * time.Second
	IpamResourceCreationTimeout time.Duration = 37 * time.Second
	IpamResourceUpdateTimeout   time.Duration = 33 * time.Second
	DefaultReconcileTimeout     time.Duration = 1 * time.Minute

	MaxLengthNICname        = 15
	MaxLengthVirtualNICname = MaxLengthNICname
	FakeIfMappingCount      = 128

	UIDlabel                          = "ipam/UID"
	PermanentIDlabel                  = "ipam/PermanentID"
	MacLabel                          = "ipam/MAC"
	IPlabel                           = "ipam/IP"
	IpUidLabel                        = "ipam/IP-UID-" //nolint:golint
	SubnetIDLabel                     = "ipam/SubnetID"
	SubnetPoolIDLabel                 = "ipam/SubnetPoolID"
	DefaultSubnetLabel                = "ipam/DefaultSubnet"
	ProvisionInterfaceKey             = "provision"
	RegionNameLabel                   = "kaas.mirantis.com/region"
	DefaultForRegionLabel             = "ipam/DefaultForRegion"
	DefaultForRegionKey               = "default-for-region"
	NamespaceLabel                    = "ipam/Namespace"
	DeprecatedClusterLabel            = "ipam/Cluster"
	ClusterIDLabel                    = "ipam/ClusterID"
	ClusterRefLabel                   = "cluster.sigs.k8s.io/cluster-name"
	ClusterProviderLabel              = "kaas.mirantis.com/provider"
	DefaultForClusterLabel            = "ipam/DefaultForCluster"
	DefaultForNamespaceLabel          = "ipam/DefaultForNamespace"
	IpamHostIDLabel                   = "ipam/IpamHostID"
	MachineIDAnnotation               = "kaas.mirantis.com/uid"
	MachineIDLabel                    = "ipam/MachineID"
	BMHIDLabel                        = "ipam/BMHostID"
	L2TemplateIDLabel                 = "ipam/L2TemplateID"
	L2TemplatePreInstalledLabel       = "ipam/PreInstalledL2Template"
	L2TemplateFromPreInstalledLabel   = "ipam/GeneratedFromPreInstalled"
	L2TemplatePreInstalledBackupLabel = "ipam/BackupOfPreInstalledL2Template"
	ServiceLabel                      = "ipam/SVC"
	PerServiceLabelPrefix             = ServiceLabel + "-"

	MgmtSubnetServiceName  = "k8s-lcm" // Do not change without sync with core. PrimaryIP depends of it
	MgmtSubnetServiceLabel = PerServiceLabelPrefix + MgmtSubnetServiceName

	AllocationReqAnnotationPrefix = "ipam/AllocationReq-"
	ForcedDeletionAnnotation      = "ipam/ForcedDeletion"
	ForcedUpdateAnnotation        = "ipam/ForcedUpdate"

	CustomFinalizer             = "finalizer.ipam.mirantis.com"
	ForegroundDeletionFinalizer = "foregroundDeletion"
	ReqToReconcileAnnotation    = "ipam/req-to-reconcile"
	DefaultConfigFile           = "/etc/cluster-api-provider/conf"

	SvcLBserviceName    = "LBhost"
	SvcLBserviceLabel   = PerServiceLabelPrefix + SvcLBserviceName
	SvcMtLbServiceName  = "MetalLB"
	SvcMtLbServiceLabel = PerServiceLabelPrefix + SvcMtLbServiceName

	NetconfigFilesStateOK  = "OK"
	NetconfigFilesStateERR = "ERR"

	NetconfigUpdateModeGracePeriod = "MANUAL-GRACEPERIOD"
	NetconfigUpdateModeManual      = "MANUAL"
	NetconfigUpdateModeAuto        = "AUTO"
	NetconfigUpdateModeAutoUnsafe  = "AUTO-UNSAFE"
)

var NetconfigUpdateGracePeriod = 3 * time.Hour
