package objects

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"

	exampleLib "github.com/Mirantis/mcc-api/example/lib"
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	openstackconfigv1 "github.com/Mirantis/mcc-api/pkg/apis/public/openstackproviderconfig/v1alpha1"
	util "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

const (
	RegionKey = "kaas.mirantis.com/region"
)

func GenerateCluster(path, clusterName, namespace, externalNetworkID, credentialsName, releaseName, keyName, region string) (*clusterv1.Cluster, error) {
	cluster, err := loadCluster(path)
	if err != nil {
		return nil, errors.Errorf("load cluster fails: %w", err)
	}
	cluster.Namespace = namespace
	// Hardcoded value
	exampleLib.SetProviderLabel(&cluster.ObjectMeta, "openstack")
	exampleLib.SetRegionLabel(&cluster.ObjectMeta, region)

	providerSpecObj, err := util.GetClusterSpecObj(cluster)
	if err != nil {
		return nil, errors.Errorf("get cluster spec object fails: %w", err)
	}
	switch providerSpec := providerSpecObj.(type) {
	case *openstackconfigv1.OpenstackClusterProviderSpec:
		if externalNetworkID != "" {
			providerSpec.ExternalNetworkID = externalNetworkID
		}
	}

	providerSpec, err := util.GetClusterSpec(cluster)
	if err != nil {
		return nil, errors.Errorf("get cluster spec fails: %w", err)
	}
	if credentialsName != "" {
		providerSpec.Credentials = credentialsName
	}
	if releaseName != "" {
		providerSpec.Release = releaseName
	}
	if clusterName != "" {
		cluster.Name = clusterName
	}
	if cluster.Namespace == "" {
		cluster.Namespace = corev1.NamespaceDefault
	}
	if keyName != "" {
		providerSpec.PublicKeys = []kaasv1alpha1.PublicKeyRef{{Name: keyName}}
	}

	providerSpec.KaaS.Management.Enabled = false
	providerSpec.KaaS.Management.HelmReleases = []kaasv1alpha1.HelmRelease{}

	return cluster, nil
}

func GenerateMachines(path, clusterName, machinePrefix, region string) ([]*clusterv1.Machine, error) {
	machines, err := LoadMachines(path)
	if err != nil {
		return machines, errors.Errorf("load machines fails: %w", err)
	}
	for i, machine := range machines {
		if machinePrefix != "" {
			machine.GenerateName = ""
			machine.Name = fmt.Sprintf("%s-%d", machinePrefix, i)
		}
		machine.ObjectMeta.Labels[clusterv1.MachineClusterLabelName] = clusterName
		machine.Labels[RegionKey] = region
	}
	return machines, nil
}
