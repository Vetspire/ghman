package runs

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
)

func Delete(ctx context.Context, client *github.Client, fileName string) error {
	runs, err := getAll(ctx, client, fileName)
	if err != nil {
		return err
	}
	for _, run := range runs {
		if _, err := client.Actions.DeleteWorkflowRun(ctx, "Vetspire", "vetspire", *run.ID); err != nil {
			return fmt.Errorf("error deleting workflow run %d: %v", *run.ID, err)
		}
	}
	return nil
}

func getAll(ctx context.Context, client *github.Client, fileName string) ([]*github.WorkflowRun, error) {
	var allRuns []*github.WorkflowRun
	opt := &github.ListWorkflowRunsOptions{}
	for {
		runs, resp, err := client.Actions.ListWorkflowRunsByFileName(ctx, "Vetspire", "vetspire", fileName, opt)
		if err != nil {
			return nil, fmt.Errorf("error listing runs for filename %s: %v", fileName, err)
		}
		allRuns = append(allRuns, runs.WorkflowRuns...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	
	return allRuns, nil
}
