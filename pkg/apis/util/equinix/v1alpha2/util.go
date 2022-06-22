package util

import (
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/apis/public/equinix/v1alpha2"
	util "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

var _ = util.ClusterSpecGetter(&v1alpha2.EquinixMetalClusterProviderSpec{})
var _ = util.ClusterStatusGetter(&v1alpha2.EquinixMetalClusterProviderStatus{})
var _ = util.MachineSpecGetter(&v1alpha2.EquinixMetalMachineProviderSpec{})
var _ = util.MachineStatusGetter(&v1alpha2.EquinixMetalMachineProviderStatus{})

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

func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha2.EquinixMetalMachineProviderSpec, error) {
	return DecodeMachineSpec(&machine.Spec)
}

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
