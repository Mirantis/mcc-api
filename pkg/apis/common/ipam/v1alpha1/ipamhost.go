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
	"strings"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"

	"github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/v2/pkg/apis/util/ipam/k8sutil/types"
)

//-----------------------------------------------------------------------------
func (in IpamHostStatus) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

func (in *IpamHostStatus) AddMessage(msg string) {
	in.Messages = in.Messages.Append(msg)
}

//-----------------------------------------------------------------------------

func (in *IpamHost) getRefFromLabel(kind string) (rv *types.NamespacedName, err error) {
	labels := in.GetLabels()
	machineShortRef, ok := labels[fmt.Sprintf("ipam/%s", kind)]
	fields := strings.SplitN(machineShortRef, ".", 2)
	if !ok || len(fields) != 2 {
		return nil, fmt.Errorf("IpamHost '%s/%s': reference to object '%s' %w", in.Namespace, in.Name, kind, k8types.ErrorNotFound)
	}
	rv = &types.NamespacedName{
		Namespace: fields[0],
		Name:      fields[1],
	}
	return rv, nil
}

// GetMachineRef -- returns Ref to Machine
func (in *IpamHost) GetMachineRef() (*types.NamespacedName, error) {
	return in.getRefFromLabel("MachineRef")
}

// GetBMHostRef -- returns Ref to BMHost
func (in *IpamHost) GetBMHostRef() (*types.NamespacedName, error) {
	return in.getRefFromLabel("BMHostRef")
}

// GetPermanentID -- returns PermanentID
func (in *IpamHost) GetPermanentID() string {
	uid := in.GetLabels()[config.PermanentIDlabel]
	if uid == "" {
		uid = string(in.GetUID())
	}
	return uid
}

// ----------------------------------------------------------------------------
func (in *IpamHost) GetObjCreated() string {
	return in.Status.ObjCreated
}
func (in *IpamHost) GetObjUpdated() string {
	return in.Status.ObjUpdated
}
func (in *IpamHost) GetObjStatusUpdated() string {
	return in.Status.ObjStatusUpdated
}
func (in *IpamHost) SetObjCreated(s string) (rv bool) {
	if in.Status.ObjCreated != s {
		in.Status.ObjCreated = s
		rv = true
	}
	return rv
}
func (in *IpamHost) SetObjUpdated(s string) (rv bool) {
	if in.Status.ObjUpdated != s {
		in.Status.ObjUpdated = s
		rv = true
	}
	return rv
}
func (in *IpamHost) SetObjStatusUpdated(s string) (rv bool) {
	if in.Status.ObjStatusUpdated != s {
		in.Status.ObjStatusUpdated = s
		rv = true
	}
	return rv
}

// ----------------------------------------------------------------------------

func (in *IpamHost) StatusToYAML() (rv []byte, err error) {
	return yaml.Marshal(&in.Status)
}

func (in *IpamHost) YAMLtoStatus(b []byte) (err error) {
	err = yaml.Unmarshal(b, &in.Status)
	return err
}

func (in *IpamHost) GetStatus() interface{} {
	return &in.Status
}
func (in *IpamHost) GetSpec() interface{} {
	return &in.Spec
}
func (in *IpamHost) GetMetadata() interface{} {
	return in.GetObjectMeta()
}

// ----------------------------------------------------------------------------
