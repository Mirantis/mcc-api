package cmd

import (
	"github.com/spf13/cobra"
)

type ClusterOptions struct {
	ClusterName    string
	Namespace      string
	KubeconfigPath string
}

var clo = &ClusterOptions{}

func (clo *ClusterOptions) RegisterOptions(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&clo.ClusterName, "cluster-name", "kaas-child", "Managed cluster name")
	cmd.PersistentFlags().StringVar(&clo.Namespace, "namespace", "kaas-child", "Managed cluster namespace")
	cmd.PersistentFlags().StringVar(&clo.KubeconfigPath, "management-kubeconfig", "kubeconfig", "Path to management cluster kubeconfig")
}
