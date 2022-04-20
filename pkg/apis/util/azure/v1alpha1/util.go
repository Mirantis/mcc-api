package util

import (
	"github.com/pkg/errors"

	"github.com/Mirantis/mcc-api/pkg/apis/public/azure/v1alpha1"
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	util "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
)

var _ = util.ClusterSpecGetter(&v1alpha1.AzureClusterProviderSpec{})
var _ = util.ClusterStatusGetter(&v1alpha1.AzureClusterProviderStatus{})
var _ = util.MachineSpecGetter(&v1alpha1.AzureMachineProviderSpec{})
var _ = util.MachineStatusGetter(&v1alpha1.AzureMachineProviderStatus{})

func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha1.AzureClusterProviderSpec, error) {
	obj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.AzureClusterProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha1.AzureClusterProviderStatus, error) {
	obj, err := util.GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.AzureClusterProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha1.AzureMachineProviderSpec, error) {
	obj, err := util.GetMachineSpecObj(machine)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.AzureMachineProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha1.AzureMachineProviderStatus, error) {
	obj, err := util.GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.AzureMachineProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}
