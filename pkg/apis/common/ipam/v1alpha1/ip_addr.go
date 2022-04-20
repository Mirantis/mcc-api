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
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
)

func (in *IPaddr) IsForVIF() bool {
	if tmp := strings.Split(in.Status.MAC, ":"); len(tmp) > 1 && tmp[0] == "VI" {
		return true
	}
	return false
}

func (in *IPaddr) GetIPref() string {
	return fmt.Sprintf("%s/%s", in.Namespace, in.Name)
}

func (in IPaddrStatus) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

// GetPermanentID -- returns PermanentID
func (in *IPaddr) GetPermanentID() string {
	uid := in.GetLabels()[config.PermanentIDlabel]
	if uid == "" {
		uid = string(in.GetUID())
	}
	return uid
}

// ----------------------------------------------------------------------------
func (in *IPaddr) GetObjCreated() string {
	return in.Status.ObjCreated
}
func (in *IPaddr) GetObjUpdated() string {
	return in.Status.ObjUpdated
}
func (in *IPaddr) GetObjStatusUpdated() string {
	return in.Status.ObjStatusUpdated
}
func (in *IPaddr) SetObjCreated(s string) (rv bool) {
	if in.Status.ObjCreated != s {
		in.Status.ObjCreated = s
		rv = true
	}
	return rv
}
func (in *IPaddr) SetObjUpdated(s string) (rv bool) {
	if in.Status.ObjUpdated != s {
		in.Status.ObjUpdated = s
		rv = true
	}
	return rv
}
func (in *IPaddr) SetObjStatusUpdated(s string) (rv bool) {
	if in.Status.ObjStatusUpdated != s {
		in.Status.ObjStatusUpdated = s
		rv = true
	}
	return rv
}

// ----------------------------------------------------------------------------

func (in *IPaddr) StatusToYAML() (rv []byte, err error) {
	return yaml.Marshal(&in.Status)
}

func (in *IPaddr) YAMLtoStatus(b []byte) (err error) {
	err = yaml.Unmarshal(b, &in.Status)
	return err
}

func (in *IPaddr) GetStatus() interface{} {
	return &in.Status
}
func (in *IPaddr) GetSpec() interface{} {
	return &in.Spec
}
func (in *IPaddr) GetMetadata() interface{} {
	return in.GetObjectMeta()
}

// ----------------------------------------------------------------------------
