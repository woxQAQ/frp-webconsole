package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v68/github"
	"github.com/samber/lo"
	"github.com/woxQAQ/frp-webconsole/pkg/gen/frpc"
	"github.com/woxQAQ/frp-webconsole/pkg/stores"
	"github.com/woxQAQ/frp-webconsole/pkg/types"
)

type FrpcServiceImpl struct {
	ghStore stores.GithubClient
}

// ListFrpRelease implements ListFrpRelease.
func (f *FrpcServiceImpl) ListFrpRelease(ctx context.Context, payload *frpc.ListFrpReleasePayload) (res []*frpc.FrpRelease, err error) {
	releases, err := f.ghStore.GetReleaseList(ctx, types.FrpOwner, types.FrpRepo)
	if err != nil {
		return nil, err
	}
	return lo.Map(releases, func(r *github.RepositoryRelease, _ int) *frpc.FrpRelease {
		assets := lo.Filter(r.Assets, func(a *github.ReleaseAsset, _ int) bool {
			return strings.Contains(*a.Name, fmt.Sprintf("%s_%s", *payload.Os, *payload.Arch))
		})
		return &frpc.FrpRelease{
			TagName: r.TagName,
			Assets: lo.Map(assets, func(a *github.ReleaseAsset, _ int) *frpc.FrpAsset {
				return &frpc.FrpAsset{
					Name:        a.Name,
					Size:        a.Size,
					Downloads:   a.DownloadCount,
					DownloadURL: a.BrowserDownloadURL,
				}
			}),
			CreatedAt: lo.ToPtr(r.CreatedAt.String()),
		}
	}), nil
}

func NewFrpcService(ghStore stores.GithubClient) frpc.Service {
	return &FrpcServiceImpl{
		ghStore: ghStore,
	}
}
