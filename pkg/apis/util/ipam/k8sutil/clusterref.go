/*
Copyright © 2020 Mirantis

Inspired by https://github.com/inwinstack/ipam/, https://github.com/inwinstack/blended/

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8sutil

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kiConfig "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

const (
	ClusterAPIgroup        = "cluster.k8s.io"
	ClusterAPIversion      = "v1alpha1"
	ClusterAPIkind         = "Cluster"
	ClusterAPIgroupVersion = ClusterAPIgroup + "/" + ClusterAPIversion
)

//-----------------------------------------------------------------------------

// GetClusterObj -- returns Cluster Obj as Unstructured
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

//-----------------------------------------------------------------------------
// GetClusterProvider -- returns provider for Cluster
func GetClusterProvider(ctx context.Context, cl client.Client, namespace, name string) (string, error) {
	cluster, err := GetClusterObj(ctx, cl, namespace, name)
	if err != nil {
		return "", fmt.Errorf("unable to fetch Cluster '%s/%s': %w", namespace, name, err)
	}
	rv, ok := cluster.GetLabels()[kiConfig.ClusterProviderLabel]
	if !ok {
		return "", fmt.Errorf("%w for Cluster '%s/%s'", k8types.ErrorProviderUndefined, namespace, name)
	}
	return rv, nil
}

//-----------------------------------------------------------------------------
// GetClusterRef -- returns Ref to Cluster in the *metav1.OwnerReference format
func GetClusterRef(ctx context.Context, cl client.Client, m k8types.K8sObject) (*metav1.OwnerReference, error) {
	ows := m.GetOwnerReferences()
	for i := range ows {
		if ows[i].APIVersion == ClusterAPIgroupVersion && ows[i].Kind == ClusterAPIkind {
			return &ows[i], nil
		}
	}
	name, ok := m.GetLabels()[kiConfig.ClusterRefLabel]
	if !ok || name == "" {
		return nil, fmt.Errorf("%w for resource '%s/%s'", k8types.ErrorClusterUndefined, m.GetNamespace(), m.GetName())
	}
	// call k8s API to fetch Cluster and build OwnerRef
	if cl == nil {
		return nil, fmt.Errorf("%w, k8s client not given,unable to build Cluster OwnerReferences for resource '%s/%s'", k8types.ErrorInoperable, m.GetNamespace(), m.GetName())
	}
	rv, err := getClusterOwnerReference(ctx, cl, m.GetNamespace(), name)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

//-----------------------------------------------------------------------------
// GetClusterName -- returns Name of Cluster related to the k8types.K8sObject
func GetClusterName(ctx context.Context, cl client.Client, m k8types.K8sObject) (string, error) {
	name := m.GetLabels()[kiConfig.ClusterRefLabel]
	if name != "" {
		return name, nil
	}

	ows := m.GetOwnerReferences()
	for i := range ows {
		if ows[i].APIVersion == ClusterAPIgroupVersion && ows[i].Kind == ClusterAPIkind {
			return ows[i].Name, nil
		}
	}

	return "", fmt.Errorf("%w for resource '%s/%s'", k8types.ErrorClusterUndefined, m.GetNamespace(), m.GetName())
}

//-----------------------------------------------------------------------------
// SetClusterRefByOwnerReference -- setup cluster Ref
// returns:
//    true if some of Refs was really modifyed
func SetClusterRefByOwnerReference(ctx context.Context, m k8types.K8sObject, ref *metav1.OwnerReference) (rv bool) {
	var (
		noChangesNeed bool
	)
	// by label
	if SetLabel(m, kiConfig.ClusterRefLabel, ref.Name) {
		rv = true
	}
	// // by meta/v1 API call
	// if m.GetClusterName() != ref.Name {
	// 	m.SetClusterName(ref.Name)
	// 	rv = true
	// }
	// by owner definition
	ows := m.GetOwnerReferences()
	n := -1
	for i := range ows {
		if ows[i].APIVersion == ClusterAPIgroupVersion && ows[i].Kind == ClusterAPIkind {
			if ows[i].Name == ref.Name && ows[i].UID == ref.UID {
				noChangesNeed = true
			} else {
				// need to change existing clusterRef
				n = i
			}
			break
		}
	}
	switch {
	case noChangesNeed:
		return rv
	case n >= 0:
		ows[n] = *ref // change existing clusterRef
		rv = true
	default:

		ows = append(ows, *ref) // OwnerReference for Cluster does not exists, should be append
		rv = true
	}
	m.SetOwnerReferences(ows)
	return rv
}

//-----------------------------------------------------------------------------
// SetClusterRefByClusterName -- setup cluster Ref
// returns:
//      true if some of Refs was really modifyed
//      error -- if happens
func SetClusterRefByClusterName(ctx context.Context, cl client.Client, m k8types.K8sObject, clusterName string) (bool, error) {
	clusterRef, err := getClusterOwnerReference(ctx, cl, m.GetNamespace(), clusterName)
	if err != nil {
		return false, err
	}
	return SetClusterRefByOwnerReference(ctx, m, clusterRef), nil
}

//-----------------------------------------------------------------------------

func GetClusterNameFromContext(ctx context.Context) (string, error) {
	clusterName, ok := ctx.Value(kiConfig.ClusterRefLabel).(string)
	switch {
	case !ok:
		return "", fmt.Errorf("%w of cluster name in the context, should be string, given: %#v", k8types.ErrorWrongFormat, ctx.Value(kiConfig.ClusterRefLabel))
	case clusterName == "":
		return "", fmt.Errorf("%w of cluster name in the context: empty string not allowed", k8types.ErrorWrongFormat)
	}
	return clusterName, nil
}
