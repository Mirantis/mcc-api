package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
)

var RootCmd = &cobra.Command{
	Use:   "example",
	Short: "Run example",
	Long:  `Run example of using MCC public API`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Flags().Set("logtostderr", "false")
		if err != nil {
			return err
		}
		return cmd.Flags().Set("alsologtostderr", "true")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Do Stuff Here
		return cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) //revive:disable:deep-exit // cobra exit status control
	}
}

func init() {
	klog.InitFlags(flag.CommandLine)
	RootCmd.SetGlobalNormalizationFunc(cliflag.WordSepNormalizeFunc)
	RootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
