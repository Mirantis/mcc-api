package k8sutil

import (
	"context"
	"fmt"
	k8types "github.com/Mirantis/mcc-api/v2/pkg/apis/util/ipam/k8sutil/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// +gocode:public-api=true
	ClusterAPIgroup = "cluster.k8s.io"
	// +gocode:public-api=true
	ClusterAPIversion = "v1alpha1"
	// +gocode:public-api=true
	ClusterAPIkind = "Cluster"
	// +gocode:public-api=true
	ClusterAPIgroupVersion = ClusterAPIgroup + "/" + ClusterAPIversion
)

// GetClusterObj -- returns Cluster Obj as Unstructured
// +gocode:public-api=true
func GetClusterObj(ctx context.Context, cl client.Client, namespace, name string) (*unstructured.Unstructured, error) {
	if cl == nil {
		return nil, fmt.Errorf("%w: k8s client was not given, imposible to fetch Cluster", k8types.ErrorWrongParametr)
	}

	nn := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	cluster := &unstructured.Unstructured{}
	cluster.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   ClusterAPIgroup,
		Version: ClusterAPIversion,
		Kind:    ClusterAPIkind,
	})

	if err := cl.Get(ctx, nn, cluster); err != nil {
		return nil, err
	}
	return cluster, nil
}

// getClusterOwnerReference -- Call k8s API to fetch Cluster object
// and build OwnerReference to it
// +gocode:public-api=true
func getClusterOwnerReference(ctx context.Context, cl client.Client, namespace, name string) (*metav1.OwnerReference, error) {
	cluster, err := GetClusterObj(ctx, cl, namespace, name)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch Cluster '%s/%s': %w", namespace, name, err)
	}
	return &metav1.OwnerReference{
		APIVersion: cluster.GetAPIVersion(),
		Kind:       cluster.GetKind(),
		Name:       cluster.GetName(),
		UID:        cluster.GetUID(),
	}, nil
}
