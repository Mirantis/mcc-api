package v1alpha1

import (
	metallbv1beta1 "github.com/Mirantis/mcc-api/v2/pkg/apis/external/metallb/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MetalLBBGPAdvertisement struct {
	Labels map[string]string                   `json:"labels,omitempty"`
	Name   string                              `json:"name"`
	Spec   metallbv1beta1.BGPAdvertisementSpec `json:"spec,omitempty"`
}
type MetalLBAddressPool struct {
	Labels map[string]string              `json:"labels,omitempty"`
	Name   string                         `json:"name"`
	Spec   metallbv1beta1.AddressPoolSpec `json:"spec"`
}

// MetalLBConfigList contains a list of MetalLBConfig
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type MetalLBConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MetalLBConfig `json:"items"`
}
type MetalLBBGPPeer struct {
	Labels map[string]string          `json:"labels,omitempty"`
	Name   string                     `json:"name"`
	Spec   metallbv1beta1.BGPPeerSpec `json:"spec,omitempty"`
}
type MetalLBIPAddressPool struct {
	Labels map[string]string                `json:"labels,omitempty"`
	Name   string                           `json:"name"`
	Spec   metallbv1beta1.IPAddressPoolSpec `json:"spec"`
}
type MetalLBBFDProfile struct {
	Labels map[string]string             `json:"labels,omitempty"`
	Name   string                        `json:"name"`
	Spec   metallbv1beta1.BFDProfileSpec `json:"spec,omitempty"`
}

// MetalLBConfigSpec defines a list of MetalLB CR objects templates
type MetalLBConfigSpec struct {
	// +optional
	AllowToUpdateByIPAM bool `json:"allowToUpdateByIPAM"`
	// +optional
	AddressPools []MetalLBAddressPool `json:"addressPools,omitempty"`
	// +optional
	BFDProfiles []MetalLBBFDProfile `json:"bfdProfiles,omitempty"`
	// +optional
	BGPAdvertisements []MetalLBBGPAdvertisement `json:"bgpAdvertisements,omitempty"`
	// +optional
	BGPPeers []MetalLBBGPPeer `json:"bgpPeers,omitempty"`
	// +optional
	Communities []MetalLBCommunity `json:"communities,omitempty"`
	// +optional
	IPAddressPools []MetalLBIPAddressPool `json:"ipAddressPools,omitempty"`
	// +optional
	L2Advertisements []MetalLBL2Advertisement `json:"l2Advertisements,omitempty"`
}
type MetalLBCommunity struct {
	Labels map[string]string            `json:"labels,omitempty"`
	Name   string                       `json:"name"`
	Spec   metallbv1beta1.CommunitySpec `json:"spec,omitempty"`
}
type MetalLBL2Advertisement struct {
	Labels map[string]string                  `json:"labels,omitempty"`
	Name   string                             `json:"name"`
	Spec   metallbv1beta1.L2AdvertisementSpec `json:"spec,omitempty"`
}

// MetalLBConfig represents the MetalLB configuration for a particular cluster.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type MetalLBConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MetalLBConfigSpec `json:"spec"`
}

// +gocode:public-api=true
type MetallbAddressPool struct {
	Addresses         []string           `json:"addresses"`
	AutoAssign        *bool              `json:"auto-assign,omitempty"`
	AvoidBuggyIPs     bool               `json:"avoid-buggy-ips,omitempty"`
	BGPAdvertisements []BGPAdvertisement `json:"bgp-advertisements,omitempty"`
	Name              string             `json:"name"`
	Protocol          string             `json:"protocol"`
}

// +gocode:public-api=true
type BGPAdvertisement struct {
	AggregationLength *int     `json:"aggregation-length"`
	LocalPref         *uint32  `json:"localpref"`
	Communities       []string `json:"communities"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&MetalLBConfig{}, &MetalLBConfigList{})
}
