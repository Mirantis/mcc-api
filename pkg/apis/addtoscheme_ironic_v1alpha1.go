package public

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/ironic/v1alpha1"
)

// +gocode:public-api=true
func init() {

	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
