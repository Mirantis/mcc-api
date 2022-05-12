package cmd

import (
	"flag"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	mccclient "github.com/Mirantis/mcc-api/example/lib/mccclient"
)

type ClusterUpgradeOptions struct {
	ClusterKubeconfig string
	ReleaseName       string
	Latest            bool
}

var cuo = &ClusterUpgradeOptions{}

var clusterUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade managed cluster release version",
	Long:  `Upgrade managed cluster release version with one command`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := UpgradeManagedCluster(); err != nil {
			klog.Exit(err)
		}
	},
}

func init() {
	clusterUpgradeCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	clo.RegisterOptions(clusterUpgradeCmd)
	clusterUpgradeCmd.PersistentFlags().StringVar(&cuo.ReleaseName, "release-name", "", "The name of cluster release to upgrade managed cluster")

	RootCmd.AddCommand(clusterUpgradeCmd)
}

func UpgradeManagedCluster() error {
	mccClient, err := mccclient.NewMccClient(clo.KubeconfigPath)
	if err != nil {
		return err
	}

	return mccClient.Upgrade(clo.ClusterName, clo.Namespace, cuo.ReleaseName)
}
