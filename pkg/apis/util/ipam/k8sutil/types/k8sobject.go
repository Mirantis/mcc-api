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

package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// K8sObject -- Minimalistic interface, described k8s object
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
