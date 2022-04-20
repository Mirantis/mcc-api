package helmutil

import (
	"reflect"
	"testing"
)

func TestOverrideREWithRE(t *testing.T) {
	dst := REFromValues(Values{
		"a": "a1",
		"b": "b1",
		"c": Values{
			"d": "d1",
			"e": "e1",
		},
	})
	err := MarshalRE(&dst)
	if err != nil {
		t.Fatalf("failed to marshal values: %v", err)
	}

	src := REFromValues(Values{
		"b": "b2",
		"c": Values{
			"d": "d2",
		},
	})

	err = OverrideREWithRE(&dst, &src)
	if err != nil {
		t.Fatalf("failed to override values: %v", err)
	}
	res := dst.Object.(*Unstructured).Object
	t.Logf("override result: %v", res)

	expected := map[string]interface{}{
		"a": "a1",
		"b": "b2",
		"c": map[string]interface{}{
			"d": "d2",
			"e": "e1",
		},
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("res != expected: %#v != %#v", res, expected)
	}
}
