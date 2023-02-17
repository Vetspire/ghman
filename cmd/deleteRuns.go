package cmd

import (
	"Vetspire/ghman/pkg/runs"
	"context"

	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var all bool

// deleteRunsCmd represents the deleteRuns command
var deleteRunsCmd = &cobra.Command{
	Use:   "delete [flags] file",
	Short: "Delete workflow runs for a specific workflow based on the workflow name.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if token != "" {
			ctx = context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client = github.NewClient(tc)
		} else {
			Logger.Fatal("no access token provided")
		}
		workflowFile := args[0]
		if workflowFile == "" {
			Logger.Error("no workflow file name specified")
		}
		if all {
			if err := runs.Delete(ctx, client, workflowFile); err != nil {
				Logger.Fatal("error deleting all runs from workflow", zap.Error(err))
			}
		}
	},
}

func init() {
	runsCmd.AddCommand(deleteRunsCmd)
	deleteRunsCmd.Flags().BoolVarP(&all, "all", "a", true, "Delete all runs for workflow")
}
