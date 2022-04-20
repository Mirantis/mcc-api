/*
Copyright © 2020 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

var (
	ProviderConfig       Config
	ConfigFilePath       string
	SingleRegionMode     bool
	DefaultNamespace     = "default"
	ReconcileTimeout     = DefaultReconcileTimeout
	NetconfigNetplanPath = "/etc/netplan/60-kaas-lcm-netplan.yaml"

	// L2TemplatePreInstallIgnoreNamespaces -- list of namespaces,
	// which will be ignored during copy of pre-installed L2Templates
	L2TemplatePreInstallIgnoreNamespaces = []string{
		"default",
		"ceph",
		"ceph-lcm-mirantis",
		"istio-system",
		"kaas",
		"kube-node-lease",
		"kube-public",
		"kube-system",
		"local-path-storage",
		"metallb-system",
		"node-feature-discovery",
		"openstack-ceph-shared",
		"rook-ceph",
		"stacklight",
	}

	// RequiredCRDs -- list of CRD resources, required for IPAM to work
	RequiredCRDs = []string{
		"baremetalhosts.metal3.io",
		"clusters.cluster.k8s.io",
		"machines.cluster.k8s.io",
		"ipamhosts.ipam.mirantis.com",
		"ipaddrs.ipam.mirantis.com",
		"subnets.ipam.mirantis.com",
		"subnetpools.ipam.mirantis.com",
		"l2templates.ipam.mirantis.com",
	}
)

// Shared config for all cloud providers
type Config struct {
	// Defines the region provider will be responsible for.
	// Only objects with matched "kaas.mirantis.com/region" annotation will be managed.
	Region string `json:"region,omitempty"`

	// Path to a regional kubeconfig. In-cluster is used if not set.
	RegionalKubeConfig string `json:"regionalKubeConfig,omitempty"`
	// Path to a management kubeconfig. In-cluster is used if not set.
	ManagementKubeConfig string `json:"managementKubeConfig,omitempty"`
}

func (c *Config) Parse(configFilePath string) error {
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("could not open config file %s: %w", configFilePath, err)
	}

	err = yaml.Unmarshal(configFile, &c)
	if err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", configFilePath, err)
	}
	return nil
}
