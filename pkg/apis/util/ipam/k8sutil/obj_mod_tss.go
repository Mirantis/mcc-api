package k8sutil

import (
	k8sutilTypes "github.com/Mirantis/mcc-api/v2/pkg/apis/util/ipam/k8sutil/types"
	"strings"
	"time"
)

// +gocode:public-api=true
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

// +gocode:public-api=true
func GetTimeFromObjUpdatedString(s string) (rv time.Time) {
	tmp := strings.Split(s, " ")
	if len(tmp) > 0 {
		if t, err := time.Parse(time.RFC3339Nano, tmp[0]); err == nil {
			rv = t.UTC()
		}
	}
	return rv
}
