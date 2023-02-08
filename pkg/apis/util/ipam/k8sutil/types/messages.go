package types

import (
	"regexp"
	"sort"
	"strings"
)

//-----------------------------------------------------------------------------
// +gocode:public-api=true
type Messages []string

func (in Messages) Append(msg string) Messages {
	needUpdate := true
	msg = strings.TrimSpace(msg)
	upperMsg := strings.ToUpper(msg)
	for i := range in {
		if strings.EqualFold(in[i], upperMsg) {
			needUpdate = false
			break
		}
	}
	if needUpdate {
		in = append(in, msg)
	}
	return in.Sorted()
}
func (in Messages) ReCheck(re *regexp.Regexp) bool {
	rv := false
	for i := range in {
		if re.MatchString(in[i]) {
			rv = true
			break
		}
	}
	return rv
}
func (in Messages) ReCheckMessage(expr string) (bool, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return false, err
	}
	return in.ReCheck(re), nil
}
func (in Messages) HasErrors() bool {
	re := regexp.MustCompile(`^ERR:`)
	return in.ReCheck(re)
}
func (in Messages) Sorted() Messages {
	sort.Strings(in)
	return in
}
func (in Messages) String() string {
	return strings.Join(in.Sorted(), "\n")
}
