package util

import (
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	util "github.com/Mirantis/mcc-api/v2/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
)

var (
	// +gocode:public-api=true
	_ = util.ClusterSpecGetter(&v1alpha1.EquinixMetalClusterProviderSpec{})
	// +gocode:public-api=true
	_ = util.ClusterStatusGetter(&v1alpha1.EquinixMetalClusterProviderStatus{})
	// +gocode:public-api=true
	_ = util.MachineSpecGetter(&v1alpha1.EquinixMetalMachineProviderSpec{})
	// +gocode:public-api=true
	_ = util.MachineStatusGetter(&v1alpha1.EquinixMetalMachineProviderStatus{})
)

// +gocode:public-api=true
func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha1.EquinixMetalClusterProviderSpec, error) {
	obj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.EquinixMetalClusterProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

// +gocode:public-api=true
func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha1.EquinixMetalClusterProviderStatus, error) {
	obj, err := util.GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.EquinixMetalClusterProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

// +gocode:public-api=true
func DecodeMachineSpec(machineSpec *clusterv1.MachineSpec) (*v1alpha1.EquinixMetalMachineProviderSpec, error) {
	obj, err := util.DecodeMachineSpecObj(machineSpec)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.EquinixMetalMachineProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

// +gocode:public-api=true
func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha1.EquinixMetalMachineProviderSpec, error) {
	return DecodeMachineSpec(&machine.Spec)
}

// +gocode:public-api=true
func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha1.EquinixMetalMachineProviderStatus, error) {
	obj, err := util.GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.EquinixMetalMachineProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

// Get my ASN and BGP peers for MetalLB
// +gocode:public-api=true
func GetClusterMetalLBPeers(cluster *clusterv1.Cluster, isManaged bool) (int, []v1alpha1.BGPPeer, error) {
	clusterSpec, err := GetClusterSpec(cluster)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "failed to get spec for cluster %s/%s", cluster.Namespace, cluster.Name)
	}

	myAsn := v1alpha1.LocalBGPMyASN
	if clusterSpec.BGP.MyASN != 0 {
		myAsn = clusterSpec.BGP.MyASN
	}

	peers := v1alpha1.RegionalMetalLBPeers
	if isManaged {
		peers = v1alpha1.ManagedMetalLBPeers
	}
	if len(clusterSpec.BGP.MetalLBPeers) > 0 {
		peers = clusterSpec.BGP.MetalLBPeers
	}
	return myAsn, peers, nil
}

// Get my ASN and BGP peers for bird
// +gocode:public-api=true
func GetClusterBirdPeers(cluster *clusterv1.Cluster, isManaged, isGreaterThan280 bool) (int, []v1alpha1.BGPPeer, error) {
	clusterSpec, err := GetClusterSpec(cluster)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "failed to get spec for cluster %s/%s", cluster.Namespace, cluster.Name)
	}

	myAsn := v1alpha1.LocalBGPMyASN
	if clusterSpec.BGP.MyASN != 0 {
		myAsn = clusterSpec.BGP.MyASN
	}

	peers := v1alpha1.RegionalBIRDPeers
	if isManaged && isGreaterThan280 {
		peers = v1alpha1.ManagedBIRDPeers
	}
	if len(clusterSpec.BGP.MetalLBPeers) > 0 {
		peers = clusterSpec.BGP.BIRDPeers
	}
	return myAsn, peers, nil
}
