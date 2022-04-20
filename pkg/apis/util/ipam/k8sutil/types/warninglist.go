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
	"sort"
	"strings"
)

//-----------------------------------------------------------------------------
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
	rv = append(WarningsList{}, in...) // deep copy
	sort.Strings(rv)
	return rv
}

func (in WarningsList) String() string {
	return strings.Join(in.Sorted(), "\n")
}
