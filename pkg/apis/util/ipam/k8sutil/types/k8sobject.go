package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// K8sObject -- Minimalistic interface, described k8s object
// +gocode:public-api=true
type K8sObject interface {
	GetFinalizers() []string
	GetAnnotations() map[string]string
	GetClusterName() string
	GetLabels() map[string]string
	GetNamespace() string
	GetName() string
	GetUID() types.UID
	GetOwnerReferences() []metav1.OwnerReference
	SetAnnotations(map[string]string)
	SetClusterName(string)
	SetFinalizers([]string)
	SetLabels(map[string]string)
	SetOwnerReferences([]metav1.OwnerReference)
}
