package types

import (
	"fmt"
	"strings"

	k8types "github.com/Mirantis/mcc-api/v2/pkg/apis/util/ipam/k8sutil/types"
)

const (
	SubnetScopeGlobal    = "global"
	SubnetScopeNamespace = "namespace"
	SubnetScopeCluster   = "cluster"
)

// LabelSelector -- is a altrnative for network searchning
// instead just fetch by name, network will be searched by labels (name and value).
//
// l3Layout:
// - scope: namespace
//   subnetName: kaas-mgmt
//   labelSelector:
//     SVC/mcc-lcm:    '1'
// 	   another/label:  another_value # optional
type LabelSelector map[string]string

type L3SubnetSrc struct {
	SubnetName     string        `json:"subnetName"`
	SubnetPool     string        `json:"subnetPool,omitempty"`
	Scope          string        `json:"scope"`
	LabelSelector  LabelSelector `json:"labelSelector,omitempty"`
	CacheSubnetKey string        `json:"-"` // exportd because k8s-apimachinery unaccept unexported fields, non-marshalled to prevent be API field
}

type L3Layout []L3SubnetSrc

func (r L3Layout) IsEmpty() bool {
	return len(r) == 0
}

func (r L3Layout) GetScope(subnetName string) string {
	for i := range r {
		if subnetName == r[i].SubnetName {
			return r[i].Scope
		}
	}
	return ""
}

func (r L3Layout) GetSubnetPoolName(subnetName string) string {
	for i := range r {
		if subnetName == r[i].SubnetName {
			return r[i].SubnetPool
		}
	}
	return ""
}

func (r L3Layout) GetSubnetKeyFromCache(subnetName string) string {
	for i := range r {
		if subnetName == r[i].SubnetName {
			return r[i].CacheSubnetKey
		}
	}
	return ""
}

func (r L3Layout) SetSubnetKeyCache(subnetName, key string) bool {
	for i := range r {
		if subnetName == r[i].SubnetName {
			r[i].CacheSubnetKey = key
			return true
		}
	}
	return false
}

// ----------------------------------------------------------------------------

type NpTemplate string
type IfMapping []string

// NP is a incoming data set to render NetPlan template
type NP struct {
	l3Layout L3Layout
	Nics     IfMapping
	Macs     IfMapping
}

func (r *NP) SetL3Layout(l L3Layout) {
	r.l3Layout = l
}
func (r *NP) GetL3Layout() L3Layout {
	return r.l3Layout
}

func (r *NP) SetIfMapping(m, macs IfMapping) {
	r.Nics = m
	r.Macs = macs
}
func (r *NP) GetNic(i int) string {
	return r.Nics[i]
}
func (r *NP) GetMac(i int) string {
	return r.Macs[i]
}

type Netplan interface {
	GetNic(int) string
	GetMac(int) string
	GetIP(string) string
	GetGateway(string) string
	GetNameservers(string) string
	GetCIDR(string) string
}

// ----------------------------------------------------------------------------

type Np4validation struct {
	NP
}

func (r *Np4validation) GetIP(key string) string {
	tmp := strings.Split(key, ":")
	if len(tmp) < 1 {
		panic(fmt.Errorf("%w of argument:'%s'", k8types.ErrorWrongFormat, key))
	}
	// IP addresses may be requested by following form:
	//   {{ip "N:subnet-name"}}
	//   {{ip "vifName:subnet-name"}}
	// in the first form N is a interface number from the IfMapping
	return "1.1.1.1" // address is not important for simple validation
}
func (r *Np4validation) GetGateway(subnet string) string {
	return "192.168.0.1" // address network is not important for simple validation
}
func (r *Np4validation) GetNameservers(subnet string) string {
	return "[\"8.8.8.8\", \"8.8.4.4\"]"
}
func (r *Np4validation) GetCIDR(subnet string) string {
	return "192.168.0.0/24" // is not important for simple validation
}

// ----------------------------------------------------------------------------
