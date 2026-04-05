package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	kubeconfig   string
	kubecontext  string
	outputFormat string
	namespace    string
	verbose      bool
)

var rootCmd = &cobra.Command{
	Use:   "kubectl-triage",
	Short: "First-response context for suspicious Kubernetes workloads",
	Long: `kubectl-triage is a read-only kubectl plugin that collects the most useful
first checks for a pod, deployment, or job — without jumping between ten commands.

It is safe to run in production clusters. It never modifies cluster state.`,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "path to kubeconfig file (default: $KUBECONFIG or ~/.kube/config)")
	rootCmd.PersistentFlags().StringVar(&kubecontext, "context", "", "kubernetes context to use")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "namespace of the target resource")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format: table | json | markdown")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "show full event list and owner chain")

	rootCmd.AddCommand(podCmd)
	rootCmd.AddCommand(deploymentCmd)
	rootCmd.AddCommand(jobCmd)
}
