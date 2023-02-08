/*
Copyright © 2020 Mirantis

Inspired by https://github.com/inwinstack/ipam/

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
	"net"
	"regexp"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/config"
)

var subnetStateRE = regexp.MustCompile(`^([[:upper:]]+)`)

func (in *Subnet) GetState() string {
	res := subnetStateRE.FindStringSubmatch(in.Status.StatusMessage)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}

func (in *Subnet) GetSubnetRef() string {
	return fmt.Sprintf("%s/%s/%s", in.Namespace, in.Name, in.UID)
}

func (in Subnet) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

// ----------------------------------------------------------------------------

func (in *Subnet) StatusToYAML() (rv []byte, err error) {
	return yaml.Marshal(&in.Status)
}

func (in *Subnet) YAMLtoStatus(b []byte) (err error) {
	err = yaml.Unmarshal(b, &in.Status)
	return err
}

func (in *Subnet) GetStatus() interface{} {
	return &in.Status
}
func (in *Subnet) GetSpec() interface{} {
	return &in.Spec
}
func (in *Subnet) GetMetadata() interface{} {
	return in.GetObjectMeta()
}

// GetPermanentID -- returns PermanentID
func (in *Subnet) GetPermanentID() string {
	uid := in.GetLabels()[config.PermanentIDlabel]
	if uid == "" {
		uid = string(in.GetUID())
	}
	return uid
}

// ----------------------------------------------------------------------------
func (in *Subnet) GetObjCreated() string {
	return in.Status.ObjCreated
}
func (in *Subnet) GetObjUpdated() string {
	return in.Status.ObjUpdated
}
func (in *Subnet) GetObjStatusUpdated() string {
	return in.Status.ObjStatusUpdated
}
func (in *Subnet) SetObjCreated(s string) (rv bool) {
	if in.Status.ObjCreated != s {
		in.Status.ObjCreated = s
		rv = true
	}
	return rv
}
func (in *Subnet) SetObjUpdated(s string) (rv bool) {
	if in.Status.ObjUpdated != s {
		in.Status.ObjUpdated = s
		rv = true
	}
	return rv
}
func (in *Subnet) SetObjStatusUpdated(s string) (rv bool) {
	if in.Status.ObjStatusUpdated != s {
		in.Status.ObjStatusUpdated = s
		rv = true
	}
	return rv
}

// ----------------------------------------------------------------------------

func (in SubnetStatus) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

//-----------------------------------------------------------------------------
func (r AllocatedIPs) getBySmth(n int, tst string) string {
	if tst == "" {
		return "" // input error
	}
	for i := range r {
		tmp := strings.Split(r[i], ":")
		if n > len(tmp)-1 {
			// out of range
			continue
		}
		if tmp[n] == tst {
			return r[i] // found
		}
	}
	return "" // not found
}

func (r AllocatedIPs) GetByIP(ipaddr string) string {
	return r.getBySmth(0, ipaddr)
}

func (r AllocatedIPs) GetByUID(uid string) string {
	return r.getBySmth(1, uid)
}

func (r AllocatedIPs) GetIPonlyList() []string {
	ipOnlyList := []string{}
	for i := range r {
		ip := strings.Split(r[i], ":")[0]
		if ip != "" && net.ParseIP(ip) != nil {
			ipOnlyList = append(ipOnlyList, ip)
		}
	}
	return ipOnlyList
}
