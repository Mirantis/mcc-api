package v1alpha1

const (
	// +gocode:public-api=true
	DevicesPrivateNetwork = "10.0.0.0/8"
	// +gocode:public-api=true
	LocalBGPRouter1Address = "169.254.255.1"
	// +gocode:public-api=true
	LocalBGPRouter2Address = "169.254.255.2"
	// +gocode:public-api=true
	LocalBGPMyASN = 65000
	// +gocode:public-api=true
	LocalBGPPeerASN = 65530
	// +gocode:public-api=true
	BIRDDirectInterface = "bond0"
	// +gocode:public-api=true
	KeepalivedVRRPInterface = "bond0"
	// +gocode:public-api=true
	DevicePrivateNetworkIf = "bond0:0"

	// These may depend on MCC release, handle with care
	// +gocode:public-api=true
	ServiceLbIPcountManagement = 11
	// +gocode:public-api=true
	ServiceLbIPcountRegional = 8
	// +gocode:public-api=true
	ServiceLbIPcountManaged = 6
)

var (
	// We have 2 BGP speakers in our deployment: bird and MetalLB. To avoid conflicts
	// (when more than one speaker on the same node tries to setup BGP session with the same peer),
	// we use different peers for bird and MetalLB. Equinix local BGP setup already provides
	// two BGP routers so we can choose one for bird and another one for MetalLB.
	// +gocode:public-api=true
	RegionalBIRDPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter1Address},
		},
	}
	// +gocode:public-api=true
	RegionalMetalLBPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter2Address},
		},
	}

	// For managed clusters we can use both BGP routers for both BGP speakers (bird and MetalLB)
	// as bird and MetalLB are not deployed on the same nodes in managed clusters.
	// +gocode:public-api=true
	ManagedBIRDPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter1Address, LocalBGPRouter2Address},
		},
	}
	// +gocode:public-api=true
	ManagedMetalLBPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter1Address, LocalBGPRouter2Address},
		},
	}
)
