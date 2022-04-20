package helmutil

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/helm/pkg/chartutil"
)

type Values = chartutil.Values

func ToValues(v interface{}) (Values, bool) {
	switch c := v.(type) {
	case Values:
		return c, true
	case map[string]interface{}:
		return Values(c), true
	default:
		return nil, false
	}
}

// DeepCopy based on k8s.io/apimachinery/pkg/runtime.DeepCopyJSON with added Values type handling
func DeepCopy(x Values) Values {
	return DeepCopyValue(x).(Values)
}

func DeepCopyValue(x interface{}) interface{} {
	switch x := x.(type) {
	case Values:
		if x == nil {
			// Typed nil - an interface{} that contains a type map[string]interface{} with a value of nil
			return x
		}
		clone := make(Values, len(x))
		for k, v := range x {
			clone[k] = DeepCopyValue(v)
		}
		return clone
	case map[string]interface{}:
		if x == nil {
			// Typed nil - an interface{} that contains a type map[string]interface{} with a value of nil
			return x
		}
		clone := make(Values, len(x))
		for k, v := range x {
			clone[k] = DeepCopyValue(v)
		}
		return clone
	case []interface{}:
		if x == nil {
			// Typed nil - an interface{} that contains a type []interface{} with a value of nil
			return x
		}
		clone := make([]interface{}, len(x))
		for i, v := range x {
			clone[i] = DeepCopyValue(v)
		}
		return clone
	case string, int64, bool, float64, nil, json.Number:
		return x
	default:
		panic(fmt.Errorf("cannot deep copy %T", x))
	}
}

//go:generate go run ../../vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go -O zz_generated.deepcopy -i ./... -h ../../hack/boilerplate.go.txt
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:deepcopy-gen=true
type Unstructured struct {
	unstructured.Unstructured `json:",inline"`
}

func (in *Unstructured) DeepCopy() *Unstructured {
	if in == nil {
		return nil
	}
	out := new(Unstructured)
	*out = *in
	out.Object = DeepCopy(in.Object)
	return out
}

func REFromValues(v Values) runtime.RawExtension {
	return runtime.RawExtension{
		Object: &Unstructured{
			unstructured.Unstructured{
				Object: v,
			},
		},
	}
}

func ParseRE(re *runtime.RawExtension) (Values, error) {
	if re.Raw != nil {
		v, err := chartutil.ReadValues(re.Raw)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse raw value from extension")
		}
		*re = REFromValues(v)
		return v, nil
	}
	if re.Object == nil {
		v := map[string]interface{}{}
		*re = REFromValues(v)
		return v, nil
	}
	uns, ok := re.Object.(*Unstructured)
	if !ok {
		return nil, errors.New("object in extension is not unstructured object")
	}
	return uns.Object, nil
}

func MarshalRE(re *runtime.RawExtension) error {
	if re.Raw != nil {
		return nil
	}
	uns, ok := re.Object.(*Unstructured)
	if !ok {
		return errors.New("object in extension is not unstructured object")
	}
	val, err := json.Marshal(uns.Object)
	if err != nil {
		return errors.Wrap(err, "coudn't marshal values")
	}
	re.Raw = val
	re.Object = nil
	return nil
}

func OverrideValuesWithValues(dst Values, src Values) {
	for k, sv := range src {
		dv, ok := dst[k]
		if ok {
			dvc, dvisv := ToValues(dv)
			svc, svisv := ToValues(sv)
			if dvisv && svisv {
				OverrideValuesWithValues(dvc, svc)
				continue
			}
		}
		dst[k] = sv
	}
}

func OverrideREWithValues(dst *runtime.RawExtension, src Values) error {
	dstv, err := ParseRE(dst)
	if err != nil {
		return err
	}
	OverrideValuesWithValues(dstv, src)
	return nil
}

func OverrideREWithRE(dst, src *runtime.RawExtension) error {
	srcv, err := ParseRE(src)
	if err != nil {
		return err
	}
	return OverrideREWithValues(dst, srcv)
}

func jsonPath(fields []string) string {
	return "." + strings.Join(fields, ".")
}

func NestedObject(v Values, fields ...string) (interface{}, bool, error) {
	var val interface{} = v

	for i, field := range fields {
		m, ok := ToValues(val)
		if !ok {
			return "", false, fmt.Errorf("%v accessor error: %v is of the type %T, expected map[string]interface{} or Values", jsonPath(fields[:i+1]), val, val)
		}

		val, ok = m[field]
		if !ok {
			return "", false, nil
		}
	}

	return val, true, nil
}

func NestedString(v Values, fields ...string) (string, bool, error) {
	val, exists, err := NestedObject(v, fields...)
	if !exists {
		return "", exists, err
	}
	res, ok := val.(string)
	if !ok {
		return "", false, fmt.Errorf("%v accessor error: %v is of the type %T, expected string", jsonPath(fields), val, val)
	}
	return res, true, nil
}

func NestedBool(v Values, fields ...string) (res bool, exists bool, err error) {
	var (
		val interface{}
		ok  bool
	)
	val, exists, err = NestedObject(v, fields...)
	if !exists {
		return false, exists, err
	}
	res, ok = val.(bool)
	if !ok {
		return false, false, fmt.Errorf("%v accessor error: %v is of the type %T, expected bool", jsonPath(fields), val, val)
	}
	return res, true, nil
}

func IsNestedBoolValueSet(v Values, key string) bool {
	val, ok, err := NestedBool(v, key)
	if !ok || err != nil {
		return false
	}

	return val
}

func NestedSlice(v Values, fields ...string) ([]interface{}, bool, error) {
	val, exists, err := NestedObject(v, fields...)
	if !exists {
		return []interface{}{}, exists, err
	}
	res, ok := val.([]interface{})
	if !ok {
		return []interface{}{}, false, fmt.Errorf("%v accessor error: %v is of the type %T, expected []Values", jsonPath(fields), val, val)
	}
	return res, true, nil
}

func NestedValues(v Values, fields ...string) (Values, bool, error) {
	val, exists, err := NestedObject(v, fields...)
	if !exists {
		return Values{}, exists, err
	}
	res, ok := ToValues(val)
	if !ok {
		return Values{}, false, fmt.Errorf("%v accessor error: %v is of the type %T, expected Values", jsonPath(fields), val, val)
	}
	return res, true, nil
}

func StringsToInterfaces(list []string) []interface{} {
	res := make([]interface{}, len(list))
	for i, v := range list {
		res[i] = interface{}(v)
	}
	return res
}
