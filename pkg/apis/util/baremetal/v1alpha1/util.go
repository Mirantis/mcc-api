package util

import (
	"github.com/Mirantis/mcc-api/pkg/apis/public/baremetal/v1alpha1"
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	util "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

var _ = util.ClusterSpecGetter(&v1alpha1.BaremetalClusterProviderSpec{})
var _ = util.ClusterStatusGetter(&v1alpha1.BaremetalClusterProviderStatus{})
var _ = util.MachineSpecGetter(&v1alpha1.BareMetalMachineProviderSpec{})
var _ = util.MachineStatusGetter(&v1alpha1.BareMetalMachineProviderStatus{})

func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha1.BaremetalClusterProviderSpec, error) {
	obj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.BaremetalClusterProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha1.BaremetalClusterProviderStatus, error) {
	obj, err := util.GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.BaremetalClusterProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha1.BareMetalMachineProviderSpec, error) {
	obj, err := util.GetMachineSpecObj(machine)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.BareMetalMachineProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha1.BareMetalMachineProviderStatus, error) {
	obj, err := util.GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.BareMetalMachineProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}
