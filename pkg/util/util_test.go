package util

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	for _, tc := range []struct {
		Name   string
		Target []string
		Source []string
		Want   []string
	}{
		{
			Name: "empty",
			Want: []string{},
		},
		{
			Name:   "empty source list",
			Target: []string{"a"},
			Want:   []string{},
		},
		{
			Name:   "empty target list",
			Source: []string{"a"},
			Want:   []string{"a"},
		},
		{
			Name:   "source and target is equal",
			Target: []string{"a", "c", "b"},
			Source: []string{"a", "b", "c"},
			Want:   []string{},
		},
		{
			Name:   "add new items",
			Target: []string{"a", "b", "c"},
			Source: []string{"f", "g", "e"},
			Want:   []string{"f", "g", "e"},
		},
		{
			Name:   "exclude duplicates",
			Target: []string{"a", "b", "c"},
			Source: []string{"a", "c", "e", "e", "f", "f", "g", "g"},
			Want:   []string{"e", "f", "g"},
		},
		{
			Name:   "empty result",
			Target: []string{"a", "b", "c"},
			Source: []string{"b", "a"},
			Want:   []string{},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			got := Diff(tc.Target, tc.Source)
			if !reflect.DeepEqual(tc.Want, got) {
				t.Errorf("want: %q, got: %q", tc.Want, got)
			}
		})

		if t.Failed() {
			return
		}
	}
}
