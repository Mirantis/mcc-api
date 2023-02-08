package public

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/equinix/v1alpha2"
)

// +gocode:public-api=true
func init() {

	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme, v1alpha2.SchemeBuilder.AddToScheme)
}
