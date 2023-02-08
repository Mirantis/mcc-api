package v1alpha1

import (
	"fmt"
	bmhv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/external/metal3-io/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// +gocode:public-api=true
	DeviceTypeSSD = "ssd"
	// +gocode:public-api=true
	DeviceTypeHDD = "hdd"
	// +gocode:public-api=true
	DeviceTypeNVME = "nvme"
)

// +gocode:public-api=true
type LvmVolumeType string

const (
	// +gocode:public-api=true
	SoftRaidLevelRaid0 SoftRaidLevel = "raid0"
	// +gocode:public-api=true
	SoftRaidLevelRaid1 SoftRaidLevel = "raid1"
	// +gocode:public-api=true
	SoftRaidLevelRaid10 SoftRaidLevel = "raid10"

	// +gocode:public-api=true
	SoftRaidMetadata10 SoftRaidMetadataVersion = "1.0"
	// +gocode:public-api=true
	SoftRaidMetadata11 SoftRaidMetadataVersion = "1.1"
	// +gocode:public-api=true
	SoftRaidMetadata12 SoftRaidMetadataVersion = "1.2"
)

type DeviceFilter struct {
	ByName string `json:"byName,omitempty"`
	WorkBy string `json:"workBy,omitempty"`
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	MinSizeGiB float32           `json:"minSizeGiB,omitempty"`
	MinSize    resource.Quantity `json:"minSize,omitempty"`
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	MaxSizeGiB float32           `json:"maxSizeGiB,omitempty"`
	MaxSize    resource.Quantity `json:"maxSize,omitempty"`
	Type       string            `json:"type,omitempty"`
	Wipe       bool              `json:"wipe"`
	UseFor     []DeviceUseFor    `json:"useFor,omitempty"`
}

func (dev *DeviceFilter) GetMaxSizeBytes() int64 {
	if dev.MaxSize.IsZero() {
		return int64(dev.MaxSizeGiB * (1 << 30))
	}
	return dev.MaxSize.Value()
}
func (dev *DeviceFilter) GetMinSizeBytes() int64 {
	if dev.MinSize.IsZero() {
		return int64(dev.MinSizeGiB * (1 << 30))
	}
	return dev.MinSize.Value()
}

type FileSystem struct {
	MountPoint     string `json:"mountPoint,omitempty"`
	MountOpts      string `json:"mountOpts,omitempty"`
	Opts           string `json:"opts,omitempty"`
	Partition      string `json:"partition,omitempty"`
	SoftRaidDevice string `json:"softRaidDevice,omitempty"`
	LogicalVolume  string `json:"logicalVolume,omitempty"`
	FileSystem     string `json:"fileSystem"`
}

func (fs *FileSystem) Validate(partitionsMap map[string]DiskPartition, logicalVolumeMap map[string]LV, srDeviceMap map[string]SoftRaidDevice) error {

	if fs.Partition == "" && fs.LogicalVolume == "" && fs.SoftRaidDevice == "" {
		return errors.Errorf("One of partition or logicalVolume or softRaidDevice must be set for filesystem %+v", fs)
	}
	if fs.Partition != "" && fs.LogicalVolume != "" {
		return errors.Errorf("both partition and lv are set in filesystem %+v", fs)
	}
	if fs.Partition != "" && fs.SoftRaidDevice != "" {
		return errors.Errorf("both partition and soft raid (MD) device are set in filesystem %+v", fs)
	}
	if fs.LogicalVolume != "" && fs.SoftRaidDevice != "" {
		return errors.Errorf("both lv and soft raid (MD) device are set in filesystem %+v", fs)
	}

	if fs.Partition != "" {
		if _, ok := partitionsMap[fs.Partition]; !ok {
			return errors.Errorf("unknown partition %s", fs.Partition)
		}
	}
	if fs.LogicalVolume != "" {
		if _, ok := logicalVolumeMap[fs.LogicalVolume]; !ok {
			return errors.Errorf("unknown logical volume %s", fs.LogicalVolume)
		}
	}
	if fs.SoftRaidDevice != "" {
		if _, ok := srDeviceMap[fs.SoftRaidDevice]; !ok {
			return errors.Errorf("unknown soft raid (MD) device %s", fs.SoftRaidDevice)
		}
	}
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +kubebuilder:resource:shortName=bmhprofile;bmhp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Region",type=string,JSONPath=`.metadata.labels.kaas\.mirantis\.com/region`
// +kubebuilder:printcolumn:name="Default",type=string,JSONPath=`.metadata.labels.kaas\.mirantis\.com/defaultBMHProfile`
// +gocode:public-api=true
type BareMetalHostProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BareMetalHostProfileSpec   `json:"spec,omitempty"`
	Status BareMetalHostProfileStatus `json:"status,omitempty"`
}
type Storage struct {
	bmhv1alpha1.Storage

	WorkBy string `json:"work_by,omitempty"`
	ByName string `json:"by_name,omitempty"`

	PartNumber uint32 `json:"part_number,omitempty"`
}
type SoftRaidDevice struct {
	Name string `json:"name"`
	// +kubebuilder:default="raid1"
	// +kubebuilder:validation:Enum=raid0;raid1;raid10
	Level SoftRaidLevel `json:"level,omitempty"`
	// +kubebuilder:default="1.0"
	// +kubebuilder:validation:Enum="1.0";"1.1";"1.2"
	MetadataVersion SoftRaidMetadataVersion `json:"metadata,omitempty"`
	Devices         []PVFilter              `json:"devices"`
}

const (
	// +gocode:public-api=true
	LvmLinear LvmVolumeType = "linear"
	// +gocode:public-api=true
	LvmRaid1 LvmVolumeType = "raid1"
)

type GrubConfig struct {
	ToDevices          []string `json:"toDevices,omitempty"`
	ToDeviceFailOver   string   `json:"toDeviceFailOver,omitempty"`
	DefaultGrubOptions []string `json:"defaultGrubOptions,omitempty"`
}
type DiskPartition struct {
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	SizeGiB float32           `json:"sizeGiB,omitempty"`
	Size    resource.Quantity `json:"size,omitempty"`
	Name    string            `json:"name,omitempty"`
	Flags   []string          `json:"partflags,omitempty"`
	Wipe    bool              `json:"wipe,omitempty"`

	Device Storage `json:"-"`
}

func (s *DiskPartition) GetSizeBytes() int64 {
	if s.Size.IsZero() {
		return int64(s.SizeGiB * (1 << 30))
	}
	return s.Size.Value()
}

// Kernel parameters that should be set using sysctl and using kernel module
// configuration files. For kernel boot parameters that are set using grub
// configuration file please refer to GrubConfig.DefaultGrubOptions.
type KernelParameters struct {
	Sysctl  map[string]string `json:"sysctl,omitempty"`  // sysctl variables dictionary (name/value pairs)
	Modules []KernelModule    `json:"modules,omitempty"` // list of kernel modules to be configured
}
type Device struct {
	Device     DeviceFilter    `json:"device,omitempty"`
	Partitions []DiskPartition `json:"partitions,omitempty"`
}

// PVFilter used to find partitions for both lvm and md
// If SoftRaidDevice is set, this particular device will be used in lvm volume group
type PVFilter struct {
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	MinSizeGiB float32           `json:"minSizeGiB,omitempty"`
	MinSize    resource.Quantity `json:"minSize,omitempty"`
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	MaxSizeGiB     float32           `json:"maxSizeGiB,omitempty"`
	MaxSize        resource.Quantity `json:"maxSize,omitempty"`
	Type           string            `json:"type,omitempty"`
	MaxDevices     int               `json:"maxDevices,omitempty"`
	Partition      string            `json:"partition,omitempty"`
	SoftRaidDevice string            `json:"softRaidDevice,omitempty"`
}

func (dev *PVFilter) GetMaxSizeBytes() int64 {
	if dev.MaxSize.IsZero() {
		return int64(dev.MaxSizeGiB * (1 << 30))
	}
	return dev.MaxSize.Value()
}
func (dev *PVFilter) GetMinSizeBytes() int64 {
	if dev.MinSize.IsZero() {
		return int64(dev.MinSizeGiB * (1 << 30))
	}
	return dev.MinSize.Value()
}
func (dev *PVFilter) Validate(partitionsMap map[string]DiskPartition) error {
	if dev.Partition == "" {
		return nil
	}
	if dev.GetMinSizeBytes() != 0 {
		return errors.Errorf("'minSize' or 'minSizeGiB' option cannot be used alongside with 'partition': %+v", dev)
	}
	if dev.MaxDevices != 0 {
		return errors.Errorf("'maxDevices' option cannot be used alongside with 'partition': %+v", dev)
	}
	if dev.Type != "" {
		return errors.Errorf("'type' option cannot be used alongside with 'partition': %+v", dev)
	}
	if _, ok := partitionsMap[dev.Partition]; !ok {
		return errors.Errorf("unknown partition %s", dev.Partition)
	}
	return nil
}

// +gocode:public-api=true
type (
	// +gocode:public-api=true
	SoftRaidLevel string
	// +gocode:public-api=true
	SoftRaidMetadataVersion string
)
type BareMetalHostProfileSpec struct {
	PreDeployScript  string           `json:"preDeployScript,omitempty"`
	PostDeployScript string           `json:"postDeployScript,omitempty"`
	GrubConfig       GrubConfig       `json:"grubConfig,omitempty"`
	Devices          []Device         `json:"devices,omitempty"`
	SoftRaidDevices  []SoftRaidDevice `json:"softRaidDevices,omitempty"`
	VolumeGroups     []VG             `json:"volumeGroups,omitempty"`
	LogicalVolumes   []LV             `json:"logicalVolumes,omitempty"`
	FileSystems      []FileSystem     `json:"fileSystems,omitempty"`
	RootFSURL        string           `json:"rootFSURL,omitempty"`
	KernelParameters KernelParameters `json:"kernelParameters,omitempty"`
}

func (s *BareMetalHostProfileSpec) Validate() error {
	var configDriveFound bool
	if s == nil {
		return fmt.Errorf("invalid BMH profile spec, undefined object")
	}

	partitionsMap := map[string]DiskPartition{}
	for _, device := range s.Devices {
		for i, partition := range device.Partitions {
			if _, ok := partitionsMap[partition.Name]; ok {
				return errors.Errorf("duplicate partition name %s", partition.Name)
			}
			partitionsMap[partition.Name] = partition
			if partition.Name == "config-2" {
				configDriveFound = true
			}
			if partition.GetSizeBytes() == 0 && len(device.Partitions) > i+1 {

				return errors.Errorf(
					"zero partition size is allowed for the last partition only. dev: %+v partition[%d]: %+v",
					device, i, partition)
			}
		}

	}

	if !configDriveFound {
		return errors.New("At least one config-2 partition required")
	}

	volumeGroupMap := map[string]VG{}
	for _, vg := range s.VolumeGroups {
		if _, ok := volumeGroupMap[vg.Name]; ok {
			return errors.Errorf("duplicate volume group name %s", vg.Name)
		}
		volumeGroupMap[vg.Name] = vg
		for _, dev := range vg.Devices {
			err := dev.Validate(partitionsMap)
			if err != nil {
				return errors.Errorf("%s for vg: %s", err, vg.Name)
			}
		}
	}

	logicalVolumeMap := map[string]LV{}
	for _, lv := range s.LogicalVolumes {
		if _, ok := logicalVolumeMap[lv.Name]; ok {
			return errors.Errorf("duplicate lv name %s", lv.Name)
		}
		logicalVolumeMap[lv.Name] = lv
		if _, ok := volumeGroupMap[lv.VG]; !ok {
			return errors.Errorf("unknown volume group %s", lv.VG)
		}
	}

	srDeviceMap := map[string]SoftRaidDevice{}
	for _, srd := range s.SoftRaidDevices {
		if _, ok := srDeviceMap[srd.Name]; ok {
			return errors.Errorf("duplicate soft raid (MD) device name %s", srd.Name)
		}
		srDeviceMap[srd.Name] = srd
		for _, dev := range srd.Devices {
			err := dev.Validate(partitionsMap)
			if err != nil {
				return errors.Errorf("%s for soft raid device (MD): %s", err, srd.Name)
			}
		}
	}

	for _, fs := range s.FileSystems {
		err := fs.Validate(partitionsMap, logicalVolumeMap, srDeviceMap)
		if err != nil {
			return err
		}
		if fs.SoftRaidDevice != "" {
			for _, vg := range s.VolumeGroups {
				for _, dev := range vg.Devices {
					if dev.SoftRaidDevice == fs.SoftRaidDevice {
						return errors.Errorf("same soft raid device is used by FS and PV: %s", fs.SoftRaidDevice)
					}
				}
			}
		}
	}

	return nil
}

// Kernel module configuration.
type KernelModule struct {
	Filename string `json:"filename"` // absolute path to the module configuration file
	Content  string `json:"content"`  // module configuration that is formatted in order to be written to the configuration file
}
type LV struct {
	Name string `json:"name"`
	VG   string `json:"vg"`
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	SizeGiB float32           `json:"sizeGiB,omitempty"`
	Size    resource.Quantity `json:"size,omitempty"`
	// +kubebuilder:default=linear
	// +kubebuilder:validation:Enum=linear;raid1
	Type LvmVolumeType `json:"type,omitempty"`
}

func (lv *LV) GetSizeBytes() int64 {
	if lv.Size.IsZero() {
		return int64(lv.SizeGiB * (1 << 30))
	}
	return lv.Size.Value()
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +gocode:public-api=true
type BareMetalHostProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BareMetalHostProfile `json:"items"`
}
type BareMetalHostProfileStatus struct{}
type VG struct {
	Name string `json:"name"`
	// +kubebuilder:validation:Format=float
	// +kubebuilder:validation:Type=number
	MinSizeGiB float32           `json:"minSizeGiB,omitempty"`
	MinSize    resource.Quantity `json:"minSize,omitempty"`
	Devices    []PVFilter        `json:"devices"`
}

// +kubebuilder:validation:Enum=ceph
type DeviceUseFor string

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&BareMetalHostProfile{}, &BareMetalHostProfileList{})
}
