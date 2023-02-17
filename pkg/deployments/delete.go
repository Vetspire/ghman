package deployments

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
)

func Delete(ctx context.Context, client *github.Client, env, sha string) error {
	opt := &github.DeploymentsListOptions{
		SHA:         sha,
		Environment: env,
	}
	deployments, _, err := client.Repositories.ListDeployments(ctx, "Vetspire", "vetspire", opt)
	if err != nil {
		return fmt.Errorf("error listing deployments: %v", err)
	}
	if len(deployments) != 0 {
		return fmt.Errorf("more than 1 deployment found for env %s and sha %s", env, sha)
	}
	if _, err := client.Repositories.DeleteDeployment(ctx, "Vetspire", "vetspire", *deployments[0].ID); err != nil {
		return fmt.Errorf("error deleting deployment: %v", err)
	}
	return nil
}
