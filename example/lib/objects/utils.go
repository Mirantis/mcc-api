package objects

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
)

func loadCluster(filename string) (*clusterv1.Cluster, error) {
	var cluster *clusterv1.Cluster

	err := File(filename, func(obj runtime.Object) error {
		if cluster != nil {
			return errors.Errorf("expected only one Cluster object")
		}
		var ok bool
		cluster, ok = obj.(*clusterv1.Cluster)
		if !ok {
			return errors.Errorf("expected Cluster object, got %T", obj)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func LoadMachines(filename string) ([]*clusterv1.Machine, error) {
	var machines []*clusterv1.Machine

	err := File(filename, func(obj runtime.Object) error {
		switch cobj := obj.(type) {
		case *clusterv1.Machine:
			machines = append(machines, cobj)
		case *clusterv1.MachineList:
			for _, machine := range cobj.Items {
				machines = append(machines, machine.DeepCopy())
			}
		default:
			return errors.Errorf("expected Machine or MachineList object, got %T", obj)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return machines, nil
}
