/*
Copyright © 2020 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"regexp"
	"sort"
	"strings"
)

//-----------------------------------------------------------------------------
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
