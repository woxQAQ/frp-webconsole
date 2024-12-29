package stores

import (
	"context"

	"github.com/google/go-github/v68/github"
)

type GithubClient interface {
	GetReleaseList(ctx context.Context, owner, repo string) ([]*github.RepositoryRelease, error)
}

type githubClient struct {
}

func NewGithubClient() GithubClient {
	return &githubClient{}
}

func (c *githubClient) GetReleaseList(ctx context.Context, owner, repo string) ([]*github.RepositoryRelease, error) {
	cli := github.NewClient(nil)
	releases, _, err := cli.Repositories.ListReleases(ctx, owner, repo, &github.ListOptions{
		Page:    1,
		PerPage: 20,
	})
	if err != nil {
		return nil, err
	}
	return releases, nil
}
