package cloudinit

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"sigs.k8s.io/yaml"

	kaasIpam "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam"
	kiTypes "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/ipam/types"
	"github.com/Mirantis/mcc-api/pkg/apis/util/ipam/cidr32"
	"github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/stringedint"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

// -----------------------------------------------------------------------------
// StateString -- status in format "OK/ERR: YYYYMMDDHHMMSS sha256 checksum"
type StateString string

func (in StateString) splitted() (rv []string) {
	rv = regexp.MustCompile(`:?\s+`).Split(string(in), -1)
	if len(rv) > 3 {
		rv = rv[:3]
	}
	return rv
}

func (in StateString) GetChecksum() (rv string) {
	fields := in.splitted()
	if len(fields) > 2 {
		rv = fields[2]
	}
	return rv
}

func (in StateString) GetTsString() (rv string) {
	fields := in.splitted()
	if len(fields) > 1 {
		rv = fields[1]
	}
	return rv
}

func (in StateString) GetTs() (rv time.Time) {
	rv, _ = time.Parse(time.RFC3339Nano, in.GetTsString())
	return rv
}

func (in StateString) IsOK() (rv bool) {
	fields := in.splitted()
	if len(fields) > 0 && fields[0] == "OK" {
		rv = true
	}
	return rv
}

// -----------------------------------------------------------------------------

// V2Nameservers -- section to define Nameserver addresses
type V2Nameservers struct {
	Search    []string `json:"search,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
}

// V2Route -- section to define Route
type V2Route struct {
	From                    string                        `json:"from,omitempty"`
	To                      string                        `json:"to,omitempty"`
	Via                     string                        `json:"via,omitempty"`
	OnLink                  string                        `json:"on-link,omitempty"`
	Metric                  *stringedint.WrongStringedInt `json:"metric,omitempty"`
	Type                    string                        `json:"type,omitempty"`
	Scope                   string                        `json:"scope,omitempty"`
	Table                   *stringedint.WrongStringedInt `json:"table,omitempty"`
	Mtu                     *stringedint.WrongStringedInt `json:"mtu,omitempty"`
	CongestionWindow        *stringedint.WrongStringedInt `json:"congestion-window,omitempty"`
	AdvertisedReceiveWindow *stringedint.WrongStringedInt `json:"advertised-receive-window,omitempty"`
}

// V2RoutingPolicy -- section to define routing policy
type V2RoutingPolicy struct {
	From     string                        `json:"from,omitempty"`
	To       string                        `json:"to,omitempty"`
	Table    *stringedint.WrongStringedInt `json:"table,omitempty"`
	Priority *stringedint.WrongStringedInt `json:"priority,omitempty"`
	Mark     *stringedint.WrongStringedInt `json:"mark,omitempty"`
	ToS      *stringedint.WrongStringedInt `json:"type-of-service,omitempty"`
}

// -----------------------------------------------------------------------------

type OVSconfigSSL struct {
	CaCert      string `json:"ca-cert,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	PrivateKey  string `json:"private-key,omitempty"`
}

type OVSconfigController struct {
	Addresses      []string `json:"addresses,omitempty"`
	ConnectionMode string   `json:"connection-mode,omitempty"`
}

type OVSconfig struct {
	ExternalIDs   map[string]string    `json:"external-ids,omitempty"`
	OtherConfig   map[string]string    `json:"other-config,omitempty"`
	LACP          string               `json:"lacp,omitempty"`
	FailMode      string               `json:"fail-mode,omitempty"`
	McastSnooping bool                 `json:"mcast-snooping,omitempty"`
	Protocols     []string             `json:"protocols,omitempty"`
	RSTP          bool                 `json:"rstp,omitempty"`
	Ports         [][]string           `json:"ports,omitempty"`
	Controller    *OVSconfigController `json:"controller,omitempty"`
	Ssl           *OVSconfigSSL        `json:"ssl,omitempty"`
}

// SCBase -- base for all SavdConfig (SC*) drivatives
type SCBase struct {
	Name                       string            `json:"set-name,omitempty"`
	Addresses                  []string          `json:"addresses,omitempty"`
	Gateway                    string            `json:"gateway4,omitempty"`
	Nameservers                *V2Nameservers    `json:"nameservers,omitempty"`
	Dhcp4                      bool              `json:"dhcp4"`
	Dhcp6                      bool              `json:"dhcp6"`
	Optional                   bool              `json:"optional,omitempty"`
	Critical                   bool              `json:"critical,omitempty"`
	ActivationMode             string            `json:"activation-mode,omitempty"`
	Macaddress                 string            `json:"macaddress,omitempty"`
	Mtu                        int               `json:"mtu,omitempty"`
	Routes                     []V2Route         `json:"routes,omitempty"`
	RoutingPolicy              []V2RoutingPolicy `json:"routing-policy,omitempty"`
	AcceptRA                   bool              `json:"accept-ra,omitempty"`
	EmitLLDP                   bool              `json:"emit-lldp,omitempty"`
	ReceiveChecksumOffload     bool              `json:"receive-checksum-offload,omitempty"`
	TransmitChecksumOffload    bool              `json:"transmit-checksum-offload,omitempty"`
	TcpSegmentationOffload     bool              `json:"tcp-segmentation-offload,omitempty"`
	Tcp6SegmentationOffload    bool              `json:"tcp6-segmentation-offload,omitempty"`
	GenericSegmentationOffload bool              `json:"generic-segmentation-offload,omitempty"`
	GenericReceiveOffload      bool              `json:"generic-receive-offload,omitempty"`
	LargeReceiveOffload        bool              `json:"large-receive-offload,omitempty"`
	Openvswitch                *OVSconfig        `json:"openvswitch,omitempty"`
}

// AddIPAddr -- Add one or more IP addresses into config
func (in *SCBase) AddIPAddr(addr string) {
	if addr == "" {
		return
	}
	in.Addresses = append(in.Addresses, addr)
}

// SetGateway -- Add one or more IP addresses into config
func (in *SCBase) SetGateway(addr string) {
	if addr == "" {
		return
	}
	in.Gateway = addr
}

// SetNameservers -- Add one or more nameserver's IP addresses into config
func (in *SCBase) SetNameservers(addrs []string) {
	if len(addrs) < 1 {
		return
	}
	if in.Nameservers == nil {
		in.Nameservers = &V2Nameservers{}
	}
	in.Nameservers.Addresses = addrs
}

// -----------------------------------------------------------------------------

// SCVlan -- struct, describes one vlan item
type SCVlan struct {
	SCBase `json:",inline"`
	ID     *stringedint.WrongStringedInt `json:"id"`
	Link   string                        `json:"link"`
}

// SCVlans -- map of vlans
type SCVlans map[string]*SCVlan

func (in SCVlans) GetIfAddressPlan() kiTypes.IfAddressPlan {
	rv := kiTypes.IfAddressPlan{}
	for k := range in {
		name := k
		if in[k].Name != "" {
			name = in[k].Name
		}
		addrs := processAdrrList(in[k].Addresses)
		if len(addrs) > 0 {
			rv = append(rv, kiTypes.IfAddressPlanItem{
				IfName:    name,
				Addresses: addrs,
			})
		}
	}
	rv.Sort()
	return rv
}

func (in SCVlans) GetNetplanSectionByIfname(ifname string) (*SCBase, error) {
	if ifname != "" {
		for n := range in {
			if n == ifname || in[n].Name == ifname {
				return &in[n].SCBase, nil
			}
		}
	}
	return nil, k8types.ErrorNotFound
}

type BridgeParametrs struct {
	AgeingTime   string                        `json:"ageing-time,omitempty"`
	ForwardDelay string                        `json:"forward-delay,omitempty"`
	HelloTime    string                        `json:"hello-time,omitempty"`
	MaxAge       string                        `json:"max-age,omitempty"`
	PathCost     string                        `json:"path-cost,omitempty"`
	Stp          bool                          `json:"stp,omitempty"`
	Priority     *stringedint.WrongStringedInt `json:"priority,omitempty"`
	PortPriority *stringedint.WrongStringedInt `json:"port-priority,omitempty"`
}

// SCBridge -- struct, describes one ethernet item
type SCBridge struct {
	SCBase     `json:",inline"`
	Interfaces []string         `json:"interfaces,omitempty"`
	Parameters *BridgeParametrs `json:"parameters,omitempty"`
}

// SCBridges -- map of ethernets
type SCBridges map[string]*SCBridge

func (in SCBridges) GetIfAddressPlan() kiTypes.IfAddressPlan {
	rv := kiTypes.IfAddressPlan{}
	for k := range in {
		name := k
		if in[k].Name != "" {
			name = in[k].Name
		}
		addrs := processAdrrList(in[k].Addresses)
		if len(addrs) > 0 {
			rv = append(rv, kiTypes.IfAddressPlanItem{
				IfName:    name,
				Addresses: addrs,
			})
		}
	}
	rv.Sort()
	return rv
}

func (in SCBridges) GetNetplanSectionByIfname(ifname string) (*SCBase, error) {
	if ifname != "" {
		for n := range in {
			if n == ifname || in[n].Name == ifname {
				return &in[n].SCBase, nil
			}
		}
	}
	return nil, k8types.ErrorNotFound
}

// SCBond -- struct, describes one ethernet item
type SCBond struct {
	SCBase     `json:",inline"`
	Interfaces []string                                 `json:"interfaces,omitempty"`
	Parameters map[string]*stringedint.WrongStringedInt `json:"parameters,omitempty"`
}

// SCBonds -- map of ethernets
type SCBonds map[string]*SCBond

func (in SCBonds) GetIfAddressPlan() kiTypes.IfAddressPlan {
	rv := kiTypes.IfAddressPlan{}
	for k := range in {
		name := k
		if in[k].Name != "" {
			name = in[k].Name
		}
		addrs := processAdrrList(in[k].Addresses)
		if len(addrs) > 0 {
			rv = append(rv, kiTypes.IfAddressPlanItem{
				IfName:    name,
				Addresses: addrs,
			})
		}
	}
	rv.Sort()
	return rv
}

func (in SCBonds) GetNetplanSectionByIfname(ifname string) (*SCBase, error) {
	if ifname != "" {
		for n := range in {
			if n == ifname || in[n].Name == ifname {
				return &in[n].SCBase, nil
			}
		}
	}
	return nil, k8types.ErrorNotFound
}

// SCEthernet -- struct, describes one ethernet item
type SCEthernet struct {
	SCBase                      `json:",inline"`
	Match                       map[string]string             `json:"match,omitempty"`
	Link                        string                        `json:"link,omitempty"`
	VirtualFunctionCount        *stringedint.WrongStringedInt `json:"virtual-function-count,omitempty"`
	EmbeddedSwitchMode          string                        `json:"embedded-switch-mode,omitempty"`
	DelayVirtualFunctionsRebind bool                          `json:"delay-virtual-functions-rebind,omitempty"`
}

// AddMacMatching -- add MAC address matching for the NIC
func (in *SCEthernet) AddMacMatching(mac string) {
	if mac == "" {
		return
	}
	if in.Match == nil {
		in.Match = map[string]string{}
	}
	in.Match["macaddress"] = strings.ToLower(mac)
}

// SCEthernets -- map of ethernets
type SCEthernets map[string]*SCEthernet

func (in SCEthernets) GetIfAddressPlan() kiTypes.IfAddressPlan {
	rv := kiTypes.IfAddressPlan{}
	for k := range in {
		name := k
		if in[k].Name != "" {
			name = in[k].Name
		}
		addrs := processAdrrList(in[k].Addresses)
		if len(addrs) > 0 {
			rv = append(rv, kiTypes.IfAddressPlanItem{
				IfName:    name,
				Addresses: addrs,
			})
		}
	}
	rv.Sort()
	return rv
}

func (in SCEthernets) GetNetplanSectionByIfname(ifname string) (*SCBase, error) {
	if ifname != "" {
		for n := range in {
			if n == ifname || in[n].Name == ifname {
				return &in[n].SCBase, nil
			}
		}
	}
	return nil, k8types.ErrorNotFound
}

// UserDataNetworkV2 -- set of data, required to generate Cloud-init v2 Ntconfig
type UserDataNetworkV2 struct {
	Version   int         `json:"version"`
	Renderer  string      `json:"renderer,omitempty"`
	Ethernets SCEthernets `json:"ethernets,omitempty"`
	Bridges   SCBridges   `json:"bridges,omitempty"`
	Bonds     SCBonds     `json:"bonds,omitempty"`
	Vlans     SCVlans     `json:"vlans,omitempty"`
}

// Generate -- process incoming data and generates NetworkConfig
func (in *UserDataNetworkV2) Generate(nics kaasIpam.NicMacMap) {
	for i := range nics {
		if nics[i].Name == "" || nics[i].MAC == "" {
			// skip NICs with unknown Name or MAC
			continue
		}

		if _, ok := in.Ethernets[nics[i].Name]; !ok {
			in.Ethernets[nics[i].Name] = &SCEthernet{}
		}

		if !strings.HasPrefix(nics[i].MAC, "VI:") {
			in.Ethernets[nics[i].Name].Name = nics[i].Name
			in.Ethernets[nics[i].Name].AddMacMatching(nics[i].MAC)
		}
		in.Ethernets[nics[i].Name].AddIPAddr(nics[i].IP)
		in.Ethernets[nics[i].Name].SetGateway(nics[i].Gateway)
		in.Ethernets[nics[i].Name].SetNameservers(nics[i].Nameservers)
	}
}

// String
func (in *UserDataNetworkV2) String() string {
	rv, _ := yaml.Marshal(in)
	return string(rv)
}

func (in *UserDataNetworkV2) CheckSum() (rv string) {
	if yml, err := yaml.Marshal(in); err != nil {
		rv = ""
	} else {
		rv = fmt.Sprintf("%x", sha256.Sum256(yml))
	}
	return rv
}

// -----------------------------------------------------------------------------

// NewUserDataNetworkV2 -- returns pointer to UserDataNetworkV2 instance
func NewUserDataNetworkV2() *UserDataNetworkV2 {
	rv := &UserDataNetworkV2{
		Version:   2,
		Ethernets: make(SCEthernets),
	}

	return rv
}

// -----------------------------------------------------------------------------
// NpConfigWithAllElements -- fulfilled Netplan config
const NpConfigWithAllElements = `---
  version: 2
  renderer: networkd
  ethernets:
    {{nic 0}}:                  # for direct usage
      dhcp4: false
      dhcp6: false
      addresses:
        - {{ip "0:subnet-0"}}
      gateway4: {{gateway_from_subnet "subnet-0"}}
      nameservers:
        addresses: {{nameservers_from_subnet "subnet-0"}}
      match:
        macaddress: {{mac 0}}
      set-name: {{nic 0}}
    {{nic 1}}:                  # for vlans
      match:
        macaddress: {{mac 1}}
      set-name: {{nic 1}}
    {{nic 2}}:                 # for bond
      match:
        macaddress: {{mac 2}}
      set-name: {{nic 2}}
    {{nic 3}}:                 # for bond
      match:
        macaddress: {{mac 3}}
      set-name: {{nic 3}}
  bonds:
    bond0:
      interfaces:
        - {{nic 2}}
        - {{nic 3}}
  vlans:
    vlan1:
      id: 101
      link: {{nic 1}}
      addresses:
        - {{ip "vlan1:subnet-1"}}
    vlan2:
      id: 102
      link: bond0
      addresses:
        - {{ip "vlan2:subnet-2"}}
    vlan3:                 # for bridge
      id: 103
      link: {{nic 1}}
  bridges:
    br1:
      interfaces:
        - vlan3
      addresses:
        - {{ip "br1:subnet-3"}}
`

type NetplanIfaceSectioner interface {
	GetNetplanSectionByIfname(string) (*SCBase, error)
}

// ----------------------------------------------------------------------------

func processAdrrList(in []string) []string {
	rv := []string{}
	for i := range in {
		ipAddr := cidr32.CleanIPv4(in[i])
		if ipAddr != "" {
			rv = append(rv, ipAddr)
		}
	}
	sort.Strings(rv)
	return rv
}
