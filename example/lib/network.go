package lib

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"

	"github.com/Mirantis/mcc-api/example/lib/openstack"
)

func GetExternalNetworkID(osConfigPath, osCloud, network string) (string, error) {
	if network == "" {
		return "", fmt.Errorf("network id/name is not defined")
	}
	osCloudConfig, err := GetOsCloudConfig(osConfigPath, osCloud)
	if err != nil {
		return "", err
	}
	client, err := openstack.NewOpenStackClientFromCloud(osCloudConfig, openstack.RequestTimeout)
	if err != nil {
		return "", err
	}

	netPagesByID, err := networks.List(client.Network, networks.ListOpts{ID: network}).AllPages()
	if err != nil {
		return "", err
	}
	networksByID, err := networks.ExtractNetworks(netPagesByID)
	if err != nil {
		return "", err
	}
	if len(networksByID) > 1 {
		return "", fmt.Errorf("too many networks found with id %s", network)
	}
	if len(networksByID) == 1 {
		return networksByID[0].ID, nil
	}

	netPagesByName, err := networks.List(client.Network, networks.ListOpts{Name: network}).AllPages()
	if err != nil {
		return "", err
	}
	networksByName, err := networks.ExtractNetworks(netPagesByName)
	if err != nil {
		return "", err
	}
	if len(networksByID) == 0 && len(networksByName) == 0 {
		return "", fmt.Errorf("network with id/name %s not found", network)
	}
	if len(networksByName) > 1 {
		return "", fmt.Errorf("too many networks found with name %s", network)
	}
	return networksByName[0].ID, nil
}
