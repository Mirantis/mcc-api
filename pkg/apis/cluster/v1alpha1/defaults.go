package v1alpha1

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// PopulateDefaultsMachineDeployment fills in default field values
// Currently it is called after reading objects, but it could be called in an admission webhook also
// +gocode:public-api=true
func PopulateDefaultsMachineDeployment(d *MachineDeployment) {
	if d.Spec.Replicas == nil {
		d.Spec.Replicas = new(int32)
		*d.Spec.Replicas = 1
	}

	if d.Spec.MinReadySeconds == nil {
		d.Spec.MinReadySeconds = new(int32)
		*d.Spec.MinReadySeconds = 0
	}

	if d.Spec.RevisionHistoryLimit == nil {
		d.Spec.RevisionHistoryLimit = new(int32)
		*d.Spec.RevisionHistoryLimit = 1
	}

	if d.Spec.ProgressDeadlineSeconds == nil {
		d.Spec.ProgressDeadlineSeconds = new(int32)
		*d.Spec.ProgressDeadlineSeconds = 600
	}

	if d.Spec.Strategy == nil {
		d.Spec.Strategy = &MachineDeploymentStrategy{}
	}

	if d.Spec.Strategy.Type == "" {
		d.Spec.Strategy.Type = common.RollingUpdateMachineDeploymentStrategyType
	}

	if d.Spec.Strategy.Type == common.RollingUpdateMachineDeploymentStrategyType {
		if d.Spec.Strategy.RollingUpdate == nil {
			d.Spec.Strategy.RollingUpdate = &MachineRollingUpdateDeployment{}
		}
		if d.Spec.Strategy.RollingUpdate.MaxSurge == nil {
			ios1 := intstr.FromInt(1)
			d.Spec.Strategy.RollingUpdate.MaxSurge = &ios1
		}
		if d.Spec.Strategy.RollingUpdate.MaxUnavailable == nil {
			ios0 := intstr.FromInt(0)
			d.Spec.Strategy.RollingUpdate.MaxUnavailable = &ios0
		}
	}

	if len(d.Namespace) == 0 {
		d.Namespace = metav1.NamespaceDefault
	}
}
