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
	"net"
	"regexp"
	"strconv"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
	"github.com/Mirantis/mcc-api/pkg/apis/util/ipam/cidr32"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

var subnetPoolStateRE = regexp.MustCompile(`^([[:upper:]]+)`)

func (in *SubnetPool) GetState() string {
	res := subnetPoolStateRE.FindStringSubmatch(in.Status.StatusMessage)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}

func (in *SubnetPool) GetSubnetPoolRef() string {
	return fmt.Sprintf("%s/%s", in.Namespace, in.Name)
}

func (in SubnetPool) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

func (in SubnetPoolStatus) String() (rv string) {
	buff, err := yaml.Marshal(in)
	if err != nil {
		rv = fmt.Sprintf("---\nerror: %s\n", err)
	} else {
		rv = fmt.Sprintf("---\n%s\n", string(buff))
	}
	return rv
}

func (in *SubnetPool) GetBlockSize() int {
	if in.Status.BlockSize == "" {
		return 0
	}
	tmp := strings.TrimPrefix(in.Status.BlockSize, "/")
	bs, _ := strconv.Atoi(tmp)
	return bs
}

//-----------------------------------------------------------------------------
// Subnet allocation related methods

func (in *SubnetPool) GetAllocationReqs() (rv []string) {
	for k := range in.GetAnnotations() {
		if strings.HasPrefix(k, config.AllocationReqAnnotationPrefix) {
			rv = append(rv, strings.TrimPrefix(k, config.AllocationReqAnnotationPrefix))
		}
	}
	return rv
}

func (in *SubnetPool) AllocateSubnet(s *Subnet) (cidr, gw string, ns []string, err error) {
	var allocatdBlock *cidr32.Net32

	_, poolCIDR, err := net.ParseCIDR(in.Spec.CIDR)
	if err != nil {
		return "", "", []string{}, fmt.Errorf("wrong CIDR '%s': %w", in.Spec.CIDR, err)
	}
	if cached := in.Status.AllocatedSubnets.GetByUID(s.GetPermanentID()); cached != "" {
		tmp := strings.Split(cached, ":")
		allocatdBlock = cidr32.CidrToNet32(tmp[0])
		cidr = allocatdBlock.String()
	} else {
		poolNet32 := cidr32.IPnetToNet32(poolCIDR)
		allocatedList := cidr32.Net32List{}
		for i := range in.Status.AllocatedSubnets {
			tmp := strings.Split(in.Status.AllocatedSubnets[i], ":")
			if len(tmp) < 2 {
				continue
			}
			if net32 := cidr32.CidrToNet32(tmp[0]); net32 != nil {
				allocatedList = append(allocatedList, *net32)
			}
		}

		allocatdBlock = poolNet32.GetFreeBlock(uint8(in.GetBlockSize()), allocatedList)
		if allocatdBlock.Empty() {
			return "", "", []string{}, fmt.Errorf("Net from SubnetPool: %w ", k8types.ErrorUnableToAllocate) //nolint:stylecheck
		}

		cidr = allocatdBlock.String()
		in.Status.AllocatedSubnets = append(in.Status.AllocatedSubnets, fmt.Sprintf("%s:%s", cidr, s.GetPermanentID()))
		in.Status.Allocatable = in.Status.Capacity - len(in.Status.AllocatedSubnets)
	}

	switch strings.ToLower(in.Spec.GatewayPolicy) {
	case "first":
		gw = cidr32.Uint32toString(allocatdBlock.First32() + 1)
	case "last":
		gw = cidr32.Uint32toString(allocatdBlock.Last32() - 1)
	}

	return cidr, gw, in.Spec.Nameservers, nil
}

//-----------------------------------------------------------------------------
func (in AllocatedSubnets) getBySmth(n int, tst string) string {
	if tst == "" {
		return "" // input error
	}
	for i := range in {
		tmp := strings.Split(in[i], ":")
		if n > len(tmp)-1 {
			// out of range
			continue
		}
		if tmp[n] == tst {
			return in[i] // found
		}
	}
	return "" // not found
}

func (in AllocatedSubnets) GetByCIDR(cidr string) string {
	return in.getBySmth(0, cidr)
}

func (in AllocatedSubnets) GetByUID(uid string) string {
	return in.getBySmth(1, uid)
}

// GetPermanentID -- returns PermanentID
func (in *SubnetPool) GetPermanentID() string {
	uid := in.GetLabels()[config.PermanentIDlabel]
	if uid == "" {
		uid = string(in.GetUID())
	}
	return uid
}

// ----------------------------------------------------------------------------
func (in *SubnetPool) GetObjCreated() string {
	return in.Status.ObjCreated
}
func (in *SubnetPool) GetObjUpdated() string {
	return in.Status.ObjUpdated
}
func (in *SubnetPool) GetObjStatusUpdated() string {
	return in.Status.ObjStatusUpdated
}
func (in *SubnetPool) SetObjCreated(s string) (rv bool) {
	if in.Status.ObjCreated != s {
		in.Status.ObjCreated = s
		rv = true
	}
	return rv
}
func (in *SubnetPool) SetObjUpdated(s string) (rv bool) {
	if in.Status.ObjUpdated != s {
		in.Status.ObjUpdated = s
		rv = true
	}
	return rv
}
func (in *SubnetPool) SetObjStatusUpdated(s string) (rv bool) {
	if in.Status.ObjStatusUpdated != s {
		in.Status.ObjStatusUpdated = s
		rv = true
	}
	return rv
}

// ----------------------------------------------------------------------------

func (in *SubnetPool) StatusToYAML() (rv []byte, err error) {
	return yaml.Marshal(&in.Status)
}

func (in *SubnetPool) YAMLtoStatus(b []byte) (err error) {
	err = yaml.Unmarshal(b, &in.Status)
	return err
}

func (in *SubnetPool) GetStatus() interface{} {
	return &in.Status
}
func (in *SubnetPool) GetSpec() interface{} {
	return &in.Spec
}
func (in *SubnetPool) GetMetadata() interface{} {
	return in.GetObjectMeta()
}

// ----------------------------------------------------------------------------
