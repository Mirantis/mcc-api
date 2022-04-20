package v1alpha1

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"sigs.k8s.io/yaml"
)

func expectValidationFailed(t *testing.T, sp *BareMetalHostProfileSpec, message string, invalidateFunc func(spec *BareMetalHostProfileSpec)) {
	invalidateFunc(sp)
	if err := sp.Validate(); err == nil {
		t.Errorf(message)
	}
}

func TestBareMetalHost_AvailableProfile_validation(t *testing.T) {
	var (
		validBMHProfile BareMetalHostProfileSpec
		storage         Storage
	)

	data, err := ioutil.ReadFile(filepath.Join("testdata", "fixture_bmhprofile.yaml"))
	if err != nil {
		t.Errorf("failed to read fixture: %s", err)
	}
	err = yaml.Unmarshal(data, &validBMHProfile)
	if err != nil {
		t.Errorf("failed to unmarshal fixture: %s", err)
	}
	if !strings.EqualFold(validBMHProfile.Devices[0].Device.Type, DeviceTypeHDD) {
		t.Errorf("Invalid device type, expect HDD: %v", validBMHProfile.Devices[0].Device)
	}
	if !strings.EqualFold(validBMHProfile.Devices[3].Device.Type, DeviceTypeSSD) {
		t.Errorf("Invalid device type, expect SSD: %v", validBMHProfile.Devices[3].Device)
	}
	if !strings.EqualFold(validBMHProfile.Devices[4].Device.Type, DeviceTypeNVME) {
		t.Errorf("Invalid device type, expect NVME: %v", validBMHProfile.Devices[4].Device)
	}
	if validBMHProfile.Devices[0].Partitions[0].Device != storage {
		t.Errorf("Invalid device")
	}
	if validBMHProfile.Devices[0].Device.ByName != "/dev/sda" {
		t.Errorf("Invalid byName")
	}
	if validBMHProfile.VolumeGroups[0].Devices[0].Type != DeviceTypeSSD {
		t.Errorf("Invalid device type for VG ssd")
	}
	err = validBMHProfile.Validate()
	if err != nil {
		t.Errorf("Failed to validate: %s", err)
	}
	if validBMHProfile.RootFSURL != "http://httpd-http/images/tgz-bionic" {
		t.Errorf("Unvalid Root FS URL")
	}
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "duplicate partition name", func(spec *BareMetalHostProfileSpec) {
		spec.Devices[0].Partitions[0].Name = "cloudInit"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "invalid partition size", func(spec *BareMetalHostProfileSpec) {
		spec.Devices[0].Partitions[0].SizeGiB = 0
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "duplicate vg name", func(spec *BareMetalHostProfileSpec) {
		spec.VolumeGroups[0].Name = "hdd"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "ambiguous device filter", func(spec *BareMetalHostProfileSpec) {
		spec.VolumeGroups[0].Devices[0].Partition = "whatever"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "unknown vg partition", func(spec *BareMetalHostProfileSpec) {
		spec.VolumeGroups[1].Devices[0].Partition = "nonexistent"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "duplicate lv name", func(spec *BareMetalHostProfileSpec) {
		spec.LogicalVolumes[0].Name = "docker"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "unknown lv volume group", func(spec *BareMetalHostProfileSpec) {
		spec.LogicalVolumes[0].VG = "nonexistent"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "missing both lv and partition", func(spec *BareMetalHostProfileSpec) {
		spec.FileSystems[0].Partition = ""
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "both lv and partition are set", func(spec *BareMetalHostProfileSpec) {
		spec.FileSystems[0].LogicalVolume = "whatever"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "nonexistent fs partition", func(spec *BareMetalHostProfileSpec) {
		spec.FileSystems[0].Partition = "nonexistent"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "nonexistent md partition", func(spec *BareMetalHostProfileSpec) {
		spec.SoftRaidDevices[0].Devices[0].Partition = "nonexistent"
	})
	expectValidationFailed(t, validBMHProfile.DeepCopy(), "md double use", func(spec *BareMetalHostProfileSpec) {
		spec.FileSystems = append(spec.FileSystems, FileSystem{
			MountPoint:     "/mnt/1",
			SoftRaidDevice: "/dev/md1",
			FileSystem:     "ext4",
		})
	})
}
