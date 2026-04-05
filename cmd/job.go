package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/topcug/kubectl-triage/internal/kube"
	"github.com/topcug/kubectl-triage/internal/render"
	"github.com/topcug/kubectl-triage/internal/triage"
	"github.com/spf13/cobra"
)

var jobCmd = &cobra.Command{
	Use:   "job <n>",
	Short: "Triage a job",
	Long:  "Collect first-response context for a suspicious or misbehaving job.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cs, err := kube.NewClient(kubeconfig, kubecontext)
		if err != nil {
			return fmt.Errorf("create client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		report, err := triage.AssembleJob(ctx, cs, name, namespace)
		if err != nil {
			return err
		}

		switch outputFormat {
		case "json":
			return render.JSON(os.Stdout, report)
		case "markdown", "md":
			render.Markdown(os.Stdout, report)
		default:
			render.Table(os.Stdout, report, verbose)
		}
		return nil
	},
}
