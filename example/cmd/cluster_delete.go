package cmd

import (
	"flag"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	mccclient "github.com/Mirantis/mcc-api/example/lib/mccclient"
)

type ClusterDeleteOptions struct {
	ClusterOptions
}

var cdo = &ClusterDeleteOptions{}

var clusterDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete managed cluster",
	Long:  `Delete managed cluster with one command`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := DeleteManagedCluster(); err != nil {
			klog.Exit(err)
		}
	},
}

func init() {
	clusterDeleteCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	cdo.RegisterOptions(clusterDeleteCmd)

	RootCmd.AddCommand(clusterDeleteCmd)
}

func DeleteManagedCluster() error {
	mccClient, err := mccclient.NewMccClient(cdo.KubeconfigPath)
	if err != nil {
		return err
	}

	return mccClient.Delete(cdo.ClusterName, cdo.Namespace)
}
