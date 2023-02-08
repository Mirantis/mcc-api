package testutil

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
)

// GetVanillaCluster return a bare minimum functional cluster resource object
// +gocode:public-api=true
func GetVanillaCluster() v1alpha1.Cluster {
	return v1alpha1.Cluster{
		Spec: v1alpha1.ClusterSpec{
			ClusterNetwork: &v1alpha1.ClusterNetworkingConfig{
				Services: v1alpha1.NetworkRanges{
					CIDRBlocks: []string{"10.96.0.0/12"},
				},
				Pods: v1alpha1.NetworkRanges{
					CIDRBlocks: []string{"192.168.0.0/16"},
				},
			},
		},
	}
}
