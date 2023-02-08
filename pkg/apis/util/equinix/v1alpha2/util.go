package util

import (
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha2"
	util "github.com/Mirantis/mcc-api/v2/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
)

var (
	// +gocode:public-api=true
	_ = util.ClusterSpecGetter(&v1alpha2.EquinixMetalClusterProviderSpec{})
	// +gocode:public-api=true
	_ = util.ClusterStatusGetter(&v1alpha2.EquinixMetalClusterProviderStatus{})
	// +gocode:public-api=true
	_ = util.MachineSpecGetter(&v1alpha2.EquinixMetalMachineProviderSpec{})
	// +gocode:public-api=true
	_ = util.MachineStatusGetter(&v1alpha2.EquinixMetalMachineProviderStatus{})
)

// +gocode:public-api=true
func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha2.EquinixMetalClusterProviderSpec, error) {
	obj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha2.EquinixMetalClusterProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

// +gocode:public-api=true
func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha2.EquinixMetalClusterProviderStatus, error) {
	obj, err := util.GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha2.EquinixMetalClusterProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

// +gocode:public-api=true
func DecodeMachineSpec(machineSpec *clusterv1.MachineSpec) (*v1alpha2.EquinixMetalMachineProviderSpec, error) {
	obj, err := util.DecodeMachineSpecObj(machineSpec)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha2.EquinixMetalMachineProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

// +gocode:public-api=true
func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha2.EquinixMetalMachineProviderSpec, error) {
	return DecodeMachineSpec(&machine.Spec)
}

// +gocode:public-api=true
func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha2.EquinixMetalMachineProviderStatus, error) {
	obj, err := util.GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha2.EquinixMetalMachineProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}
