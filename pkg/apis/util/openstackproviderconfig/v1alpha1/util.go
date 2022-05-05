package util

import (
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/apis/public/openstackproviderconfig/v1alpha1"
	util "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"

	"github.com/Mirantis/mcc-api/pkg/errors"
)

var _ = util.ClusterSpecGetter(&v1alpha1.OpenstackClusterProviderSpec{})
var _ = util.ClusterStatusGetter(&v1alpha1.OpenstackClusterProviderStatus{})
var _ = util.MachineSpecGetter(&v1alpha1.OpenstackMachineProviderSpec{})
var _ = util.MachineStatusGetter(&v1alpha1.OpenstackMachineProviderStatus{})

func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha1.OpenstackClusterProviderSpec, error) {
	obj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.OpenstackClusterProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha1.OpenstackClusterProviderStatus, error) {
	obj, err := util.GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.OpenstackClusterProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

func DecodeMachineSpec(machineSpec *clusterv1.MachineSpec) (*v1alpha1.OpenstackMachineProviderSpec, error) {
	obj, err := util.DecodeMachineSpecObj(machineSpec)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.OpenstackMachineProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha1.OpenstackMachineProviderSpec, error) {
	return DecodeMachineSpec(&machine.Spec)
}

func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha1.OpenstackMachineProviderStatus, error) {
	obj, err := util.GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.OpenstackMachineProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

func MergeMetadata(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
