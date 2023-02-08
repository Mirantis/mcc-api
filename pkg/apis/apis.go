package public

import (
	autoscalerv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/autoscaler/v1alpha1"
	azurev1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/azure/v1alpha1"
	bmv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/baremetal/v1alpha1"
	byov1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/byo/v1alpha1"
	dnsv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/dns/v1alpha1"
	equinixv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	iamv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/iam/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/openstackproviderconfig/v1alpha1"
	storagev1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/storage/v1alpha1"
	vspherev1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/vsphere/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	// +gocode:public-api=true
	Scheme = runtime.NewScheme()
	// +gocode:public-api=true
	Codecs = serializer.NewCodecFactory(Scheme)
)

// AddToSchemes may be used to add all resources defined in the project to a Scheme
// +gocode:public-api=true
var AddToSchemes = runtime.SchemeBuilder{
	v1alpha1.SchemeBuilder.AddToScheme,
	kaasv1alpha1.SchemeBuilder.AddToScheme,
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
// +gocode:public-api=true
var AddToScheme = AddToSchemes.AddToScheme

// +gocode:public-api=true
func init() {
	err := AddToScheme(Scheme)
	if err != nil {
		panic(err)
	}
}
