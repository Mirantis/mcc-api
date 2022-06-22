/*
Copyright 2022 The Kubernetes Authors.

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

// Package apis contains Kubernetes API groups.
package public

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	miracephv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/common/miraceph/v1alpha1"
	autoscalerv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/autoscaler/v1alpha1"
	azurev1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/azure/v1alpha1"
	bmv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/baremetal/v1alpha1"
	byov1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/byo/v1alpha1"
	dnsv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/dns/v1alpha1"
	equinixv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/equinix/v1alpha1"
	iamv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/iam/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/apis/public/openstackproviderconfig/v1alpha1"
	storagev1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/storage/v1alpha1"
	vspherev1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/vsphere/v1alpha1"
)

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes = runtime.SchemeBuilder{
	v1alpha1.SchemeBuilder.AddToScheme,
	kaasv1alpha1.SchemeBuilder.AddToScheme,
	miracephv1alpha1.SchemeBuilder.AddToScheme,
	bmv1alpha1.SchemeBuilder.AddToScheme,
	byov1alpha1.SchemeBuilder.AddToScheme,
	vspherev1alpha1.SchemeBuilder.AddToScheme,
	equinixv1alpha1.SchemeBuilder.AddToScheme,
	azurev1alpha1.SchemeBuilder.AddToScheme,
	storagev1alpha1.SchemeBuilder.AddToScheme,
	iamv1alpha1.SchemeBuilder.AddToScheme,
	autoscalerv1alpha1.SchemeBuilder.AddToScheme,
	dnsv1alpha1.SchemeBuilder.AddToScheme,
}

// AddToScheme adds all Resources to the Scheme
var AddToScheme = AddToSchemes.AddToScheme
var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)

func init() {
	err := AddToScheme(Scheme)
	if err != nil {
		panic(err)
	}
}
