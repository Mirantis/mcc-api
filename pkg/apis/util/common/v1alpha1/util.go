package util

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"

	publicapi "github.com/Mirantis/mcc-api/pkg/apis/public"
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

type ClusterSpecGetter interface {
	GetClusterSpecMixin() *v1alpha1.ClusterSpecMixin
	GetNewClusterStatus() runtime.Object
}

type ClusterStatusGetter interface {
	GetClusterStatusMixin() *v1alpha1.ClusterStatusMixin
}

type MachineSpecGetter interface {
	GetMachineSpecMixin() *v1alpha1.MachineSpecMixin
	GetNewMachineStatus() runtime.Object
}

type MachineStatusGetter interface {
	GetMachineStatusMixin() *v1alpha1.MachineStatusMixin
}

func decodeExtension(ext *runtime.RawExtension) (runtime.Object, error) {
	if ext.Object != nil {
		return ext.Object, nil
	}
	s := json.NewSerializer(&json.SimpleMetaFactory{}, publicapi.Scheme, publicapi.Scheme, false)
	obj, _, err := s.Decode(ext.Raw, nil, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse RawExtension value")
	}
	ext.Object = obj
	ext.Raw = nil

	return obj, nil
}

func setObjGVK(obj runtime.Object) error {
	gvks, _, err := publicapi.Scheme.ObjectKinds(obj)
	if err != nil {
		return errors.Wrapf(err, "failed to get GVK for object %v", obj)
	}
	if len(gvks) > 1 {
		return errors.Errorf("got more than one GVK for object %v", obj)
	}
	obj.GetObjectKind().SetGroupVersionKind(gvks[0])
	return nil
}

func GetClusterSpecObj(cluster *clusterv1.Cluster) (runtime.Object, error) {
	if cluster.Spec.ProviderSpec.Value == nil {
		return nil, errors.New("no providerSpec in Cluster object")
	}
	obj, err := decodeExtension(cluster.Spec.ProviderSpec.Value)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetClusterSpec(cluster *clusterv1.Cluster) (*v1alpha1.ClusterSpecMixin, error) {
	obj, err := GetClusterSpecObj(cluster)
	if err != nil {
		return nil, err
	}
	casted, ok := obj.(ClusterSpecGetter)
	if !ok {
		return nil, errors.Errorf("decoded object of type %T doesn't implement GetClusterSpecMixin func", obj)
	}
	return casted.GetClusterSpecMixin(), nil
}

func GetClusterStatusObj(cluster *clusterv1.Cluster) (runtime.Object, error) {
	var obj runtime.Object
	if cluster.Status.ProviderStatus == nil {
		if cluster.Spec.ProviderSpec.Value == nil {
			return nil, errors.New("no providerSpec in Cluster object")
		}
		specobj, err := decodeExtension(cluster.Spec.ProviderSpec.Value)
		if err != nil {
			return nil, err
		}
		casted, ok := specobj.(ClusterSpecGetter)
		if !ok {
			return nil, errors.Errorf("decoded object of type %T doesn't implement GetClusterSpecMixin func", obj)
		}
		obj = casted.GetNewClusterStatus()
		err = setObjGVK(obj)
		if err != nil {
			return nil, err
		}
		cluster.Status.ProviderStatus = &runtime.RawExtension{
			Object: obj,
		}
	} else {
		var err error
		obj, err = decodeExtension(cluster.Status.ProviderStatus)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse providerStatus value")
		}
	}
	return obj, nil
}

func GetClusterStatus(cluster *clusterv1.Cluster) (*v1alpha1.ClusterStatusMixin, error) {
	obj, err := GetClusterStatusObj(cluster)
	if err != nil {
		return nil, err
	}
	casted, ok := obj.(ClusterStatusGetter)
	if !ok {
		return nil, errors.Errorf("decoded object of type %T doesn't implement GetClusterStatusMixin func", obj)
	}
	return casted.GetClusterStatusMixin(), nil
}

func DecodeMachineSpecObj(spec *clusterv1.MachineSpec) (runtime.Object, error) {
	if spec.ProviderSpec.Value == nil {
		return nil, errors.New("no providerSpec given")
	}
	obj, err := decodeExtension(spec.ProviderSpec.Value)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetMachineSpecObj(machine *clusterv1.Machine) (runtime.Object, error) {
	return DecodeMachineSpecObj(&machine.Spec)
}

func DecodeMachineSpec(spec *clusterv1.MachineSpec) (*v1alpha1.MachineSpecMixin, error) {
	obj, err := DecodeMachineSpecObj(spec)
	if err != nil {
		return nil, err
	}
	casted, ok := obj.(MachineSpecGetter)
	if !ok {
		return nil, errors.Errorf("decoded object of type %T doesn't implement GetMachineSpecMixin func", obj)
	}
	return casted.GetMachineSpecMixin(), nil
}

func GetMachineSpec(machine *clusterv1.Machine) (*v1alpha1.MachineSpecMixin, error) {
	return DecodeMachineSpec(&machine.Spec)
}

func GetMachineStatusObj(machine *clusterv1.Machine) (runtime.Object, error) {
	var obj runtime.Object
	if machine.Status.ProviderStatus == nil {
		if machine.Spec.ProviderSpec.Value == nil {
			return nil, errors.New("no providerSpec in Machine object")
		}
		specobj, err := decodeExtension(machine.Spec.ProviderSpec.Value)
		if err != nil {
			return nil, err
		}
		casted, ok := specobj.(MachineSpecGetter)
		if !ok {
			return nil, errors.Errorf("decoded object of type %T doesn't implement GetMachineSpecMixin func", obj)
		}
		obj = casted.GetNewMachineStatus()
		err = setObjGVK(obj)
		if err != nil {
			return nil, err
		}
		machine.Status.ProviderStatus = &runtime.RawExtension{
			Object: obj,
		}
	} else {
		var err error
		obj, err = decodeExtension(machine.Status.ProviderStatus)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse providerStatus value")
		}
	}
	return obj, nil
}

func GetMachineStatus(machine *clusterv1.Machine) (*v1alpha1.MachineStatusMixin, error) {
	obj, err := GetMachineStatusObj(machine)
	if err != nil {
		return nil, err
	}
	casted, ok := obj.(MachineStatusGetter)
	if !ok {
		return nil, errors.Errorf("decoded object of type %T doesn't implement GetMachineStatusMixin func", obj)
	}
	return casted.GetMachineStatusMixin(), nil
}

func GetCurrentRelease(cluster *clusterv1.Cluster) (string, error) {
	clusterStatus, err := GetClusterStatus(cluster)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cluster %s/%s status", cluster.Namespace, cluster.Name)
	}
	if clusterStatus.ReleaseRefs != nil {
		if clusterStatus.ReleaseRefs.Previous.Name != "" {
			return clusterStatus.ReleaseRefs.Previous.Name, nil
		}
		if clusterStatus.ReleaseRefs.Current.Name != "" {
			return clusterStatus.ReleaseRefs.Current.Name, nil
		}
	}
	clusterSpec, err := GetClusterSpec(cluster)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cluster %s/%s spec", cluster.Namespace, cluster.Name)
	}
	return clusterSpec.Release, nil
}
