package v1alpha1

const (
	DevicesPrivateNetwork = "10.0.0.0/8"

	LocalBGPRouter1Address = "169.254.255.1"
	LocalBGPRouter2Address = "169.254.255.2"
	LocalBGPMyASN          = 65000
	LocalBGPPeerASN        = 65530

	BIRDDirectInterface     = "bond0"
	KeepalivedVRRPInterface = "bond0"
	DevicePrivateNetworkIf  = "bond0:0"

	// These may depend on MCC release, handle with care
	ServiceLbIPcountManagement = 11
	ServiceLbIPcountRegional   = 8
	ServiceLbIPcountManaged    = 6
)

var (
	// We have 2 BGP speakers in our deployment: bird and metallb. To avoid conflicts
	// (when more than one speaker on the same node tries to setup BGP session with the same peer),
	// we use different peers for bird and metallb. Equinix local BGP setup already provides
	// two BGP routers so we can choose one for bird and another one for metallb.

	RegionalBIRDPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter1Address},
		},
	}
	RegionalMetalLBPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter2Address},
		},
	}

	// For managed clusters we can use both BGP routers for both BGP speakers (bird and metallb)
	// as bird and metallb are not deployed on the same nodes in managed clusters.

	ManagedBIRDPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter1Address, LocalBGPRouter2Address},
		},
	}
	ManagedMetalLBPeers = []BGPPeer{
		{
			PeerAs:  LocalBGPPeerASN,
			PeerIPs: []string{LocalBGPRouter1Address, LocalBGPRouter2Address},
		},
	}
)
