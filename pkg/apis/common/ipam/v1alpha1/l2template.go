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

package v1alpha1

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/thoas/go-funk"
	"sigs.k8s.io/yaml"

	"github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/config"
	ipam "github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/ipam"
	l2tmplTypes "github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/ipam/l2template/types"
)

type prioOrderItem struct {
	ItemIndex int
	Score     int64
}
type prioOrder []*prioOrderItem

func (s prioOrder) Len() int           { return len(s) }
func (s prioOrder) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s prioOrder) Less(i, j int) bool { return s[i].Score < s[j].Score }

func (in *L2Template) GenerateIfMappingByPrio(nmm ipam.NicMacMap) (rv l2tmplTypes.IfMapping) {
	provision := false // flag whether we  should honor to provisioning NIC place
	for i := range in.Spec.AutoIfMappingPrio {
		if config.ProvisionInterfaceKey == in.Spec.AutoIfMappingPrio[i] {
			provision = true
		}
	}
	newOrder := prioOrder{}
	rg := regexp.MustCompile(`([[:alpha:]]+)(\d+)([[:alpha:]]*)(\d*)`)
	for n := range nmm {
		res := rg.FindStringSubmatch(nmm[n].Name)
		if provision && nmm[n].Primary {
			// provision NIC
			k := funk.IndexOfString(in.Spec.AutoIfMappingPrio, config.ProvisionInterfaceKey) + 1
			if k > 0 {
				newOrder = append(newOrder, &prioOrderItem{Score: int64(1000000 * k), ItemIndex: n})
			}
		} else if len(res) > 1 {
			// another NICs
			k := int64(1000000 * (funk.IndexOfString(in.Spec.AutoIfMappingPrio, res[1]) + 1))
			if k > 0 {
				bus, _ := strconv.ParseInt(res[2], 10, 64)
				bus *= 1000
				if res[4] != "" {
					slot, _ := strconv.ParseInt(res[4], 10, 64)
					bus += slot
				}
				newOrder = append(newOrder, &prioOrderItem{Score: k + bus, ItemIndex: n})
			}
		}
	}
	sort.Sort(newOrder)
	for _, nord := range newOrder {
		rv = append(rv, nmm[nord.ItemIndex].Name)
	}

	return rv
}

func (in *L2Template) GetNNGenVer() string {
	return fmt.Sprintf("%s/%s/%d/%s/%s",
		in.GetNamespace(),
		in.GetName(),
		in.GetGeneration(),
		in.GetResourceVersion(),
		in.GetUID(),
	)
}

func (in L2Template) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

func (in L2TemplateStatus) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

// GetPermanentID -- returns PermanentID
func (in *L2Template) GetPermanentID() string {
	uid := in.GetLabels()[config.PermanentIDlabel]
	if uid == "" {
		uid = string(in.GetUID())
	}
	return uid
}

// ----------------------------------------------------------------------------
func (in *L2Template) GetObjCreated() string {
	return in.Status.ObjCreated
}
func (in *L2Template) GetObjUpdated() string {
	return in.Status.ObjUpdated
}
func (in *L2Template) GetObjStatusUpdated() string {
	return in.Status.ObjStatusUpdated
}
func (in *L2Template) SetObjCreated(s string) (rv bool) {
	if in.Status.ObjCreated != s {
		in.Status.ObjCreated = s
		rv = true
	}
	return rv
}
func (in *L2Template) SetObjUpdated(s string) (rv bool) {
	if in.Status.ObjUpdated != s {
		in.Status.ObjUpdated = s
		rv = true
	}
	return rv
}
func (in *L2Template) SetObjStatusUpdated(s string) (rv bool) {
	if in.Status.ObjStatusUpdated != s {
		in.Status.ObjStatusUpdated = s
		rv = true
	}
	return rv
}

// ----------------------------------------------------------------------------

func (in *L2Template) StatusToYAML() (rv []byte, err error) {
	return yaml.Marshal(&in.Status)
}

func (in *L2Template) YAMLtoStatus(b []byte) (err error) {
	err = yaml.Unmarshal(b, &in.Status)
	return err
}

func (in *L2Template) GetStatus() interface{} {
	return &in.Status
}
func (in *L2Template) GetSpec() interface{} {
	return &in.Spec
}
func (in *L2Template) GetMetadata() interface{} {
	return in.GetObjectMeta()
}

// ----------------------------------------------------------------------------
