package cloudinit

import (
	"fmt"
	"net"
	"strings"

	"sigs.k8s.io/yaml"

	kaasIpam "github.com/Mirantis/mcc-api/v2/pkg/apis/common/ipam/ipam"
)

// -----------------------------------------------------------------------------
// https://specs.openstack.org/openstack/nova-specs/specs/liberty/implemented/metadata-service-network-info.html

// OSmetadataLink --
type OSmetadataLink struct {
	Name string `json:"id"`                   // "interface0",
	Type string `json:"type"`                 // "phy",
	MAC  string `json:"ethernet_mac_address"` // "a0:36:9f:2c:e8:80",
	// MTU  int    `json:"mtu,omitempty"`        // 9000
}

// OSmetadataNetworkRoute --
type OSmetadataNetworkRoute struct {
	Net  string `json:"network"` // "0.0.0.0",
	Mask string `json:"netmask"` // "0.0.0.0",
	Gw   string `json:"gateway"` // "23.253.157.1"
}

// OSmetadataNetwork --
type OSmetadataNetwork struct {
	Type   string                   `json:"type"`       // "ipv4",
	Link   string                   `json:"link"`       // "vlan0",
	IP     string                   `json:"ip_address"` // "23.253.157.244",
	Mask   string                   `json:"netmask"`    // "255.255.255.0",
	Routes []OSmetadataNetworkRoute `json:"routes,omitempty"`
}

// OSmetadataService --
type OSmetadataService struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

// OSmetadataNetworkConfig -- Ntconfig struct
type OSmetadataNetworkConfig struct {
	Links    []OSmetadataLink    `json:"links"`
	Networks []OSmetadataNetwork `json:"networks,omitempty"`
	Services []OSmetadataService `json:"services,omitempty"`
}

// AddIface returns pointer to freshly added Iface
func (in *OSmetadataNetworkConfig) AddIface(mac, name string) *OSmetadataLink {
	rv := OSmetadataLink{
		Name: name,
		Type: "phy",
		MAC:  strings.ToLower(mac),
	}
	in.Links = append(in.Links, rv)
	return &in.Links[len(in.Links)-1]
}

// GetServiceIndex returns index of service, or -1 if not found
func (in *OSmetadataNetworkConfig) GetServiceIndex(ip, srvtype string) int {
	rv := -1
	for i := range in.Services {
		if in.Services[i].Address == ip && in.Services[i].Type == srvtype {
			rv = i
			break
		}
	}
	return rv
}

// AddService -- add service
func (in *OSmetadataNetworkConfig) AddService(ip, srvtype string) {
	in.Services = append(in.Services, OSmetadataService{
		Type:    srvtype,
		Address: ip,
	})
}

// AddIPAddr -- Add IP address and optional gateway
func (in *OSmetadataLink) AddIPAddr(base *OSmetadataNetworkConfig, ip, gw string) {
	if ip == "" {
		return
	}
	ipv4Addr, ipv4Net, err := net.ParseCIDR(ip)
	if err != nil {
		return
	}
	// convrt CIDR netmask to non-CIDR notation
	maskList := make([]string, 4)
	for i := 0; i < len(ipv4Net.Mask); i++ {
		maskList[i] = fmt.Sprint(ipv4Net.Mask[i])
	}
	mask := strings.Join(maskList, ".")
	//---------------------------------
	rv := OSmetadataNetwork{
		Link: in.Name,
		Type: "ipv4",
		IP:   ipv4Addr.String(),
		Mask: mask,
	}
	if gw != "" {
		rv.Routes = []OSmetadataNetworkRoute{{
			Net:  "0.0.0.0",
			Mask: "0.0.0.0",
			Gw:   gw,
		}}
	}
	base.Networks = append(base.Networks, rv)
}

// AddNameservers -- Add one or more nameserver's IP addresses into config
func (in *OSmetadataLink) AddNameservers(base *OSmetadataNetworkConfig, addrs []string) {
	for _, nsAddr := range addrs {
		if base.GetServiceIndex(nsAddr, "dns") < 0 {
			base.AddService(nsAddr, "dns")
		}
	}
}

// -----------------------------------------------------------------------------

// Generate -- process incoming data and generates NetworkConfig
func (in *OSmetadataNetworkConfig) Generate(nics kaasIpam.NicMacMap) {
	in.Links = []OSmetadataLink{}
	in.Networks = []OSmetadataNetwork{}
	in.Services = []OSmetadataService{}

	for i := range nics {
		if !nics[i].Primary || nics[i].Name == "" || nics[i].MAC == "" || nics[i].IP == "" {
			// skip unnamd interface without MAC
			continue
		}

		iface := in.AddIface(nics[i].MAC, nics[i].Name)
		iface.AddIPAddr(in, nics[i].IP, nics[i].Gateway)
		iface.AddNameservers(in, nics[i].Nameservers)
	}
}

// String
func (in *OSmetadataNetworkConfig) String() string {
	rv, _ := yaml.Marshal(in)
	return string(rv)
}

// -----------------------------------------------------------------------------

// NewOSmetadataNetworkConfig -- returns pointer to OSmetadataNetworkConfig instance
func NewOSmetadataNetworkConfig() *OSmetadataNetworkConfig {
	rv := &OSmetadataNetworkConfig{
		Links:    []OSmetadataLink{},
		Networks: []OSmetadataNetwork{},
		Services: []OSmetadataService{},
	}
	return rv
}
