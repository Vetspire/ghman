package cmd

import (
	"Vetspire/ghman/pkg/deployments"
	"context"

	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// deleteDeploymentCmd represents the deleteDeployment command
var deleteDeploymentCmd = &cobra.Command{
	Use:   "delete [flags] env commit-sha",
	Short: "Delete a deployment based on the commit SHA of the deployment and the environment.",
	Long: `This command deletes a deployment based on the commit SHA of the deployment and
the deployment environment passed in as arguments.

The token used for this will need to have access to read and write deployments in the repository.`,
	Args: cobra.ExactArgs(2),
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
		if err := deployments.Delete(ctx, client, args[0], args[1]); err != nil {
			Logger.Fatal("deployments delete error", zap.Error(err))
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(deleteDeploymentCmd)
}
