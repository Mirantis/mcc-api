package v1alpha1

type MetallbAddressPool struct {
	Addresses         []string           `json:"addresses"`
	AutoAssign        *bool              `json:"auto-assign,omitempty"`
	AvoidBuggyIPs     bool               `json:"avoid-buggy-ips,omitempty"`
	BGPAdvertisements []BGPAdvertisement `json:"bgp-advertisements,omitempty"`
	Name              string             `json:"name"`
	Protocol          string             `json:"protocol"`
}

type BGPAdvertisement struct {
	AggregationLength *int     `json:"aggregation-length"`
	LocalPref         *uint32  `json:"localpref"`
	Communities       []string `json:"communities"`
}
