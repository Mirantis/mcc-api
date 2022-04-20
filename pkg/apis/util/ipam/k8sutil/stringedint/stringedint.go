package stringedint

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/intstr"

	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

// StringedInt is a type that can hold an int32 or a string.
// see k8s.io/apimachinery/pkg/util/intstr
// When used in JSON or YAML marshalling and unmarshalling, it produces
// or consumes the  inner type. This allows you to have, for example,
// a JSON field that can accept a name or number.
// The difference with k8s.io/apimachinery/pkg/util/intstr is
// any stringed Int values (i.e. "42", "0", etc...) will me marhalled as Int.
//
// +protobuf=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:openapi-gen=true

type WrongStringedInt struct {
	intstr.IntOrString
}

// MarshalJSON implements the json.Marshaller interface.
// try to return Int if it possible
func (in WrongStringedInt) MarshalJSON() ([]byte, error) {
	switch in.Type {
	case intstr.String:
		rv, err := strconv.Atoi(strings.TrimSpace(in.StrVal))
		if err == nil {
			return json.Marshal(fmt.Sprint(rv)) // Pure int to string
		}
		return json.Marshal(in.StrVal)
	case intstr.Int:
		return json.Marshal(fmt.Sprint(in.IntValue()))
	default:
		return []byte{}, fmt.Errorf("%w: impossible IntOrString.Type: %d", k8types.ErrorWrongFormat, in.Type)
	}
}
