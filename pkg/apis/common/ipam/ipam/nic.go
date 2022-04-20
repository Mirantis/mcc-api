package ipam

import (
	"fmt"
	"net"
	"sort"
	"strings"
)

// NicMacMapItem is NIC identification structure
type NicMacMapItem struct {
	Name        string   `json:"name"`
	MAC         string   `json:"mac"`
	IP          string   `json:"ip,omitempty"`
	Gateway     string   `json:"gateway,omitempty"`
	Nameservers []string `json:"nameservers,omitempty"`
	Primary     bool     `json:"primary,omitempty"`
	Online      bool     `json:"online,omitempty"`
	IPref       string   `json:"ipRef,omitempty"` // reference to IP CR in format namespace/name/UID
}

// String -- implements Stringer interface
func (r *NicMacMapItem) String() string {
	return fmt.Sprintf("%s/%s", r.MAC, r.Name)
}

// DashedMAC -- returns MAC in dashed notation
func (r *NicMacMapItem) DashedMAC() string {
	return strings.ToUpper(strings.ReplaceAll(r.MAC, ":", "-"))
}

// GetRef -- return Reference to IP CR in struct form
func (r *NicMacMapItem) GetRef() Ref {
	rv := Ref{}
	rv.SetRef(r.IPref) //nolint:errcheck
	return rv
}

// SetRef -- setup short Reference to IP CR from struct form
func (r *NicMacMapItem) SetRef(rr Ref) {
	r.IPref = rr.GetRef()
}

// NicMacMap -- sorted by MAC NICname-MAC-IP slice of all NICs of host
type NicMacMap []NicMacMapItem

// String -- implements Stringer interface
func (r *NicMacMap) String() string {
	return fmt.Sprintf("[%s]", strings.Join(r.StringSlice(), ", "))
}

// StringSlice -- returns slice of string notations
func (r *NicMacMap) StringSlice() []string {
	rv := make([]string, len(*r))
	for i := range *r {
		rv[i] = (*r)[i].String()
	}
	return rv
}

// GetByMac -- returns pointer to NicMacMapItem, searched by MAC
func (r *NicMacMap) GetByMac(mac string) (*NicMacMapItem, bool) {
	for i := range *r {
		if strings.EqualFold((*r)[i].MAC, mac) {
			return &((*r)[i]), true
		}
	}
	return nil, false
}

// GetIndexByMac -- returns index to NicMacMapItem, searched by MAC
func (r *NicMacMap) GetIndexByMac(mac string) int {
	for i := range *r {
		if strings.EqualFold((*r)[i].MAC, mac) {
			return i
		}
	}
	return -1
}

// GetByName -- returns pointer to NicMacMapItem, searched by NIC name
func (r *NicMacMap) GetByName(name string) (*NicMacMapItem, bool) {
	for i := range *r {
		if strings.EqualFold((*r)[i].Name, name) {
			return &((*r)[i]), true
		}
	}
	return nil, false
}

// GetByIP -- returns pointer to NicMacMapItem, searched by NIC IP address
func (r *NicMacMap) GetByIP(ip string) (*NicMacMapItem, bool) {
	ipAddr := strings.Split(ip, "/")[0]
	for i := range *r {
		nicAddr := strings.Split((*r)[i].IP, "/")[0]
		if nicAddr == ipAddr {
			return &((*r)[i]), true
		}
	}
	return nil, false
}

// GetPrimary -- returns pointer to NicMacMapItem, Which has Primary flag
func (r *NicMacMap) GetPrimary() (*NicMacMapItem, bool) {
	for i := range *r {
		if (*r)[i].Primary {
			return &((*r)[i]), true
		}
	}
	return nil, false
}

// Append -- append NIC
// Only IPv4 addresses are allowed, but:
// * unaddressed NICs allowed
// * IPv6 only NICs added as unaddressed
// * unaddressed record will be replaced by addressed
func (r *NicMacMap) Append(n *NicMacMapItem) {
	var (
		replaceNo        int
		replaceToIP      string
		replaceToNIC     string
		replaceToPrimary bool
	)

	if n.MAC == "" {
		return
	}
	if n.IP != "" {
		ipString := strings.Split(n.IP, "%")[0]
		if ip := net.ParseIP(ipString); ip == nil || ip.To4() == nil {
			n.IP = ""
		}
	}
	replaceNo = -1 // planning addition
exLoop:
	for i := range *r {
		if strings.EqualFold((*r)[i].MAC, n.MAC) {
			switch {
			case (!(*r)[i].Primary && n.Primary && n.IP != ""):
				replaceNo = i
				replaceToIP = n.IP
				replaceToNIC = n.Name
				replaceToPrimary = n.Primary
				break exLoop
			case (*r)[i].IP != "":
				return
			case (*r)[i].IP == "" && n.IP != "":
				replaceNo = i
				replaceToIP = n.IP
				replaceToNIC = n.Name
				replaceToPrimary = n.Primary
				break exLoop
			}
		}
	}

	if replaceNo == -1 {
		rec := *n
		rec.MAC = strings.ToUpper(rec.MAC)
		(*r) = append((*r), rec)
	} else {
		(*r)[replaceNo].IP = replaceToIP
		(*r)[replaceNo].Name = replaceToNIC
		(*r)[replaceNo].Primary = replaceToPrimary
	}
	r.SortByMAC()
}

// Sort -- sort NicMacMap by MAC
func (r *NicMacMap) SortByMAC() {
	sort.Slice((*r), func(i, j int) bool { return (*r)[i].MAC < (*r)[j].MAC })
}

// Remove -- remove item from NicMacMapItem
func (r *NicMacMap) Remove(k string) {
	n := -1
	for i := range *r {
		if (*r)[i].IP == k || strings.EqualFold((*r)[i].Name, k) || strings.EqualFold((*r)[i].MAC, k) {
			n = i
			break
		}
	}
	switch {
	case n == 0:
		(*r) = (*r)[1:]
	case n == len((*r))-1:
		(*r) = (*r)[:n]
	case n > 0:
		(*r) = append((*r)[:n], (*r)[n+1:]...)
	}
}
