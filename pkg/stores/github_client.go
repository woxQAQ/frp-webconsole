package stores

import (
	"context"

	"github.com/google/go-github/v68/github"
)

type GithubClient interface {
	GetReleaseList(ctx context.Context, owner, repo string) ([]*github.RepositoryRelease, error)
}

type githubClient struct {
	ghClient *github.Client
}

func NewGithubClient(ghClient *github.Client) GithubClient {
	return &githubClient{
		ghClient: ghClient,
	}
}

func (c *githubClient) GetReleaseList(ctx context.Context, owner, repo string) ([]*github.RepositoryRelease, error) {
	releases, _, err := c.ghClient.Repositories.ListReleases(ctx, owner, repo, &github.ListOptions{
		Page:    1,
		PerPage: 20,
	})
	if err != nil {
		return nil, err
	}
	return releases, nil
}
