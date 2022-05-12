package cmd

import (
	"flag"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	exampleLib "github.com/Mirantis/mcc-api/example/lib"
	"github.com/Mirantis/mcc-api/example/lib/mccclient"
	"github.com/Mirantis/mcc-api/example/lib/objects"
	sshLib "github.com/Mirantis/mcc-api/example/lib/ssh"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

type CreateOptions struct {
	ClusterOptions

	Region          string
	OsConfigPath    string
	OsCloud         string
	CredentialsName string
	KeyName         string
	PrivateKeyPath  string
	MachinePrefix   string
	ExternalNetwork string
	ReleaseName     string

	Cluster  string
	Machines string
}

var copts = &CreateOptions{}

var clusterCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create managed cluster",
	Long:  `Create managed cluster with one command`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreateManagedCluster(); err != nil {
			klog.Exit(err)
		}
	},
}

func init() {
	clusterCreateCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	clo.RegisterOptions(clusterCreateCmd)
	clusterCreateCmd.PersistentFlags().StringVar(&copts.Region, "region", "region-one", "Cluster region")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.OsConfigPath, "os-config-path", "", "Path to openstack cloud config file (clouds.yaml)")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.OsCloud, "os-cloud", "", "Cloud name in clouds.yaml")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.CredentialsName, "credentials-name", "cloud-config", "The name of the OpenStackCredential object with os-cloud creds")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.KeyName, "keyname", "cluster-key", "SSH Key Name")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.PrivateKeyPath, "private-key-path", "ssh_key", "Desired path to ssh private key")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.MachinePrefix, "machine-prefix", "", "Machines name prefix. If set, machine name will be in format \"machinePrefix-i\", generateName will be ignored")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.ExternalNetwork, "external-network", "public", "Openstack External Network ID/Name")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.ReleaseName, "release-name", "", "The name of managed cluster k8s release")

	clusterCreateCmd.PersistentFlags().StringVar(&copts.Cluster, "cluster", "", "Path to cluster.yaml")
	clusterCreateCmd.PersistentFlags().StringVar(&copts.Machines, "machines", "", "Path to machines.yaml")

	RootCmd.AddCommand(clusterCreateCmd)
}

func CreateManagedCluster() error {
	mccClient, err := mccclient.NewMccClient(clo.KubeconfigPath)
	if err != nil {
		return err
	}

	// Generate prerequisites for the cluster deployment
	kaasPublicKey, externalNetworkID, err := preparePrerequisites(mccClient)
	if err != nil {
		return errors.Wrap(err, "prepare prerequisties step fails")
	}

	// Prepare cluster object
	cluster, err := objects.GenerateCluster(
		copts.Cluster,
		clo.ClusterName,
		clo.Namespace,
		externalNetworkID,
		copts.CredentialsName,
		copts.ReleaseName,
		copts.KeyName,
		copts.Region,
	)
	if err != nil {
		return err
	}

	// Prepare Machines objects
	machines, err := objects.GenerateMachines(
		copts.Machines,
		clo.ClusterName,
		copts.MachinePrefix,
		copts.Region,
	)
	if err != nil {
		return err
	}

	err = mccClient.CreateManagedCluster(cluster, machines, kaasPublicKey)
	if err != nil {
		return err
	}

	// New managed cluster kubeconfig preparing
	err = mccClient.GetKubeconfig(fmt.Sprintf("%s-kubeconfig", clo.ClusterName), clo.ClusterName, clo.Namespace, cgko.Realm, cgko.Username, cgko.Password)
	if err != nil {
		return err
	}
	klog.Infof("Managed cluster %s successfully created", clo.ClusterName)
	return nil
}

func preparePrerequisites(mccClient *mccclient.MccClient) (_ *kaasv1alpha1.PublicKey, externalNetworkID string, _ error) {
	// 1 Create namespace
	if err := mccClient.ManagementCluster.EnsureNamespace(clo.Namespace); err != nil {
		return nil, "", errors.Wrapf(err, "error ensuring namespace %s", clo.Namespace)
	}
	klog.Infof("Namspace %s is present", clo.Namespace)
	// 2. Create OS credentials
	externalNetworkID = copts.ExternalNetwork

	secretExists, err := mccClient.CloudConfigExists(copts.CredentialsName, clo.Namespace)
	if err != nil {
		return nil, "", errors.Wrap(err, "error checking if cloud config exists")
	}

	if !secretExists {
		klog.Infof("Credential %v does not exist in a namespace %v, creating...", copts.CredentialsName, clo.Namespace)
		cloudConfig, err := exampleLib.GetOsCloudConfig(copts.OsConfigPath, copts.OsCloud)
		if err != nil {
			return nil, "", errors.Wrap(err, "get cloud config fails")
		}
		err = mccClient.CreateCloudConfig(copts.Region, copts.CredentialsName, clo.Namespace, cloudConfig)
		if err != nil {
			return nil, "", errors.Wrap(err, "unable to create cloud config")
		}

		externalNetworkID, err = exampleLib.GetExternalNetworkID(copts.OsConfigPath, copts.OsCloud, copts.ExternalNetwork)
		if err != nil {
			return nil, "", errors.Wrap(err, "get external network ID fails")
		}
	}
	// 3. SSH Public key
	_, publicKey, err := sshLib.GetSSHKey(copts.PrivateKeyPath)
	if err != nil {
		return nil, "", err
	}

	kaasPublicKey := &kaasv1alpha1.PublicKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      copts.KeyName,
			Namespace: clo.Namespace,
		},
		Spec: kaasv1alpha1.PublicKeySpec{
			PublicKey: string(ssh.MarshalAuthorizedKey(publicKey)),
		},
	}
	return kaasPublicKey, externalNetworkID, nil
}
