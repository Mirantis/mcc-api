package public

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/baremetal/v1alpha1"
	bmv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/external/metal3-io/v1alpha1"
	metal3v1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/metal3/v1alpha1"
)

// +gocode:public-api=true
func init() {

	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
	AddToSchemes = append(AddToSchemes, bmv1alpha1.SchemeBuilder.AddToScheme)
	AddToSchemes = append(AddToSchemes, metal3v1alpha1.SchemeBuilder.AddToScheme)
}
