package ipam

import (
	"fmt"
	"strings"

	"github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/v2/pkg/apis/util/ipam/k8sutil/types"
)

const (
	MACshouldBeGenerated = "VI:GENERATED"
)

// L2TemplateSelector describes a criteria to select L2Template for IpamHost
type L2TemplateSelector struct {
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
}

// Ref -- is a reference to another CRD
type Ref struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

// GetRef -- returns a Ref in the short string form (without UID)
func (r *Ref) GetRef() string {
	namespace := r.Namespace
	if namespace == "" {
		namespace = config.DefaultNamespace
	}

	return fmt.Sprintf("%s/%s", namespace, r.Name)
}

// SetRef -- setup a Ref from a short string form (without UID)
func (r *Ref) SetRef(ss string) error {
	aaa := strings.Split(ss, "/")
	if len(aaa) < 2 {
		return fmt.Errorf("%w of short string reference: '%s'", k8types.ErrorWrongFormat, ss)
	}
	if aaa[0] == "" {
		r.Namespace = config.DefaultNamespace
	} else {
		r.Namespace = aaa[0]
	}
	r.Name = aaa[1]
	return nil
}

func (r Ref) String() string {
	return r.GetRef()
}
