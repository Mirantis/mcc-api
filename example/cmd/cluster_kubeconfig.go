package cmd

import (
	"flag"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	mccclient "github.com/Mirantis/mcc-api/example/lib/mccclient"
)

type ClusterGetKubeconfigOptions struct {
	KubeconfigOutput string
	Realm            string
	Username         string
	Password         string
}

var cgko = &ClusterGetKubeconfigOptions{}

var clusterGetKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Get managed cluster kubeconfig",
	Long:  `Waits for managed cluster kubeconfig to be ready and writes it out to a file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetKubeconfig(); err != nil {
			klog.Exit(err)
		}
	},
}

func init() {
	clusterGetKubeconfigCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	clo.RegisterOptions(clusterGetKubeconfigCmd)
	clusterGetKubeconfigCmd.PersistentFlags().StringVar(&cgko.KubeconfigOutput, "kubeconfig-output", "kubeconfig", "Kubeconfig output file name")
	clusterGetKubeconfigCmd.PersistentFlags().StringVar(&cgko.Realm, "realm", "iam", "Keycloak realm for getting auth token")
	clusterGetKubeconfigCmd.PersistentFlags().StringVar(&cgko.Username, "username", "writer", "Username for getting auth token")
	clusterGetKubeconfigCmd.PersistentFlags().StringVar(&cgko.Password, "password", "password", "Password for getting auth token")

	RootCmd.AddCommand(clusterGetKubeconfigCmd)
}

func GetKubeconfig() error {
	mccClient, err := mccclient.NewMccClient(clo.KubeconfigPath)
	if err != nil {
		return err
	}

	return mccClient.GetKubeconfig(cgko.KubeconfigOutput, clo.ClusterName, clo.Namespace, cgko.Realm, cgko.Username, cgko.Password)
}
