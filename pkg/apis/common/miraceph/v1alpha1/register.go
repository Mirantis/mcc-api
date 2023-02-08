// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the miraceph v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=lcm.mirantis.com
package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

	"github.com/Mirantis/mcc-api/v2/pkg/errors"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "lcm.mirantis.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	AddToScheme = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func UpdateMiracephStatus(miraceph *MiraCeph, status MiraCephStatus, client client.Client) error {
	miraceph.Status = status
	if err := client.Status().Update(context.TODO(), miraceph); err != nil {
		return errors.Errorf("failed to update status for the miraceph %v/%v: %v",
			miraceph.Namespace, miraceph.Name, err)
	}
	return nil
}

func UpdateMiracephlogStatus(miracephlog *MiraCephLog, status MiraCephLogStatus, client client.Client) error {
	miracephlog.Status = status
	if err := client.Status().Update(context.TODO(), miracephlog); err != nil {
		return errors.Errorf("failed to update status for the miracephlog %v/%v: %v",
			miracephlog.Namespace, miracephlog.Name, err)
	}
	return nil
}

func UpdateCephOsdRemoveRequestStatus(cephRequest *CephOsdRemoveRequest, status *CephOsdRemoveRequestStatus, client client.Client) error {
	cephRequest.Status = status
	if err := client.Status().Update(context.TODO(), cephRequest); err != nil {
		return errors.Errorf("failed to update status for the cephOsdRemoveRequest %v/%v: %v",
			cephRequest.Namespace, cephRequest.Name, err)
	}
	return nil
}

func UpdateCephPerfTestRequestStatus(perfTestRequest *CephPerfTestRequest, status *CephPerfTestRequestStatus, client client.Client) error {
	perfTestRequest.Status = status
	if err := client.Status().Update(context.TODO(), perfTestRequest); err != nil {
		return errors.Errorf("failed to update status for the cephPerfTestRequest %v/%v: %v",
			perfTestRequest.Namespace, perfTestRequest.Name, err)
	}
	return nil
}
