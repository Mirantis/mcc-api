package util

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

	"github.com/Mirantis/mcc-api/pkg/apis/public/byo/v1alpha1"
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	util "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
)

var _ = util.ClusterSpecGetter(&v1alpha1.BYOClusterProviderSpec{})
var _ = util.ClusterStatusGetter(&v1alpha1.BYOClusterProviderStatus{})
var _ = util.MachineSpecGetter(&v1alpha1.BYOMachineProviderSpec{})
var _ = util.MachineStatusGetter(&v1alpha1.BYOMachineProviderStatus{})

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "byo.kaas.mirantis.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	ClusterSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "BYOClusterProviderSpec",
	}
	ClusterStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "BYOClusterProviderStatus",
	}
	MachineSpecGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "BYOMachineProviderSpec",
	}
	MachineStatusGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "BYOMachineProviderStatus",
	}
)

func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha1.BYOClusterProviderSpec, error) {
	obj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.BYOClusterProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha1.BYOClusterProviderStatus, error) {
	obj, err := util.GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.BYOClusterProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha1.BYOMachineProviderSpec, error) {
	obj, err := util.GetMachineSpecObj(machine)
	if err != nil {
		return nil, err
	}
	spec, ok := obj.(*v1alpha1.BYOMachineProviderSpec)
	if !ok {
		return nil, errors.Errorf("unexpected spec type: %v", obj)
	}
	return spec, err
}

func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha1.BYOMachineProviderStatus, error) {
	obj, err := util.GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	status, ok := obj.(*v1alpha1.BYOMachineProviderStatus)
	if !ok {
		return nil, errors.Errorf("unexpected status type: %v", obj)
	}
	return status, err
}

func EncodeMachineSpec(spec *v1alpha1.BYOMachineProviderSpec) (*runtime.RawExtension, error) {
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}

	spec.SetGroupVersionKind(MachineSpecGVK)

	return &runtime.RawExtension{
		Object: spec,
	}, nil
}
