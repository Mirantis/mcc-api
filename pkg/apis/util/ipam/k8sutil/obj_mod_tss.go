/*
Copyright © 2021 Mirantis

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

package k8sutil

import (
	"strings"
	"time"

	k8sutilTypes "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

func GetIpamK8sObjRealCreationTime(obj k8sutilTypes.ObjModTss) (rv time.Time) {
	tss := k8sutilTypes.TimeSlice{}
	if t := obj.GetCreationTimestamp().Time; !t.IsZero() {
		tss = append(tss, t.UTC())
	}
	if created := obj.GetObjCreated(); created != "" {
		if tmp := strings.Split(created, " "); len(tmp) > 0 {
			if t, err := time.Parse(time.RFC3339Nano, tmp[0]); err == nil {
				tss = append(tss, t.UTC())
			}
		}
	}
	if len(tss) > 0 {
		tss.Sort()
		rv = tss[0].UTC()
	}
	return rv
}

func GetTimeFromObjUpdatedString(s string) (rv time.Time) {
	tmp := strings.Split(s, " ")
	if len(tmp) > 0 {
		if t, err := time.Parse(time.RFC3339Nano, tmp[0]); err == nil {
			rv = t.UTC()
		}
	}
	return rv
}
