package types

import (
	"sort"
	"strings"
)

//-----------------------------------------------------------------------------
// +gocode:public-api=true
type WarningsList []string

func (in WarningsList) Append(msg string) WarningsList {
	needUpdate := true
	msg = strings.TrimSpace(msg)
	for i := range in {
		if in[i] == msg {
			needUpdate = false
			break
		}
	}
	if needUpdate {
		in = append(in, msg)
	}
	return in.Sorted()
}
func (in WarningsList) Sorted() (rv WarningsList) {
	rv = append(WarningsList{}, in...)
	sort.Strings(rv)
	return rv
}
func (in WarningsList) String() string {
	return strings.Join(in.Sorted(), "\n")
}
