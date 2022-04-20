package cidr32

import (
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"unsafe"
)

var CidrMaskExists = regexp.MustCompile(`/\d{1,2}$`)

// ----------------------------------------------------------------------------

// CompareIP -- positive result if address `b` more address `a`
// or ngative if less
// returns 0 if the both addresses are equal
func CompareIP(a, b net.IP) int {
	return Compare32(IPtoUint32(a), IPtoUint32(b))
}

// Compare32 -- positive result if address `b` more address `a`
// or ngative if less
// returns 0 if the both addresses are equal
func Compare32(a, b uint32) int {
	return int(b) - int(a)
}

// ----------------------------------------------------------------------------

// IPtoUint32 -- convert net.IP to Uint32
func IPtoUint32(ip net.IP) (rv uint32) {
	addr := ip.To4()
	for i, k := 0, 24; i < 4; i, k = i+1, k-8 {
		rv += uint32(addr[i]) << k
	}
	return rv
}

// CidrToUint32 -- convert string IP to Uint32
func CidrToUint32(s string) uint32 {
	if !CidrMaskExists.MatchString(s) {
		s += "/32"
	}
	ip, _, err := net.ParseCIDR(s)
	if err != nil {
		return 0
	}
	return IPtoUint32(ip)
}

// Uint32toIP -- convert Uint32 to net.IP
func Uint32toIP(ip uint32) (rv net.IP) {
	tmp := make([]byte, 4)
	copy(tmp, (*[4]byte)(unsafe.Pointer(&ip))[:])
	for f, l := 0, len(tmp)-1; f < len(tmp)/2; f, l = f+1, l-1 {
		tmp[f], tmp[l] = tmp[l], tmp[f]
	}
	rv = net.IP(tmp)
	return rv
}

// Uint32toString -- convert Uint32 to string IP representation
func Uint32toString(ip uint32) (rv string) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, ip)
	return fmt.Sprintf("%d.%d.%d.%d", bs[0], bs[1], bs[2], bs[3])
}

// NextIP -- returns next IP address
func NextIP(ip net.IP) net.IP {
	return Uint32toIP(IPtoUint32(ip) + 1)
}

// PrevIP -- returns next IP address
func PrevIP(ip net.IP) net.IP {
	return Uint32toIP(IPtoUint32(ip) - 1)
}
