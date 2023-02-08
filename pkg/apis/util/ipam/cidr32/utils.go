package cidr32

import (
	"net"
	"regexp"
)

// +gocode:public-api=true
var reIPaddr = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)

// CleanIpV4 -- got a string, search first substring looks like IPv4 address,
// validete it and return. Returns empty string if error
// +gocode:public-api=true
func CleanIPv4(addr string) string {
	addr = string(reIPaddr.Find([]byte(addr)))
	if addr == "" {
		return ""
	}
	ip := net.ParseIP(addr)
	if ip == nil {
		return ""
	}
	return ip.String()
}
