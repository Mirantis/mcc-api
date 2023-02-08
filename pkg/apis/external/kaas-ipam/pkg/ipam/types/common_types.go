package types

// L2TemplateSelector describes a criteria to select L2Template for IpamHost
// +gocode:public-api=true
type L2TemplateSelector struct {
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
}
