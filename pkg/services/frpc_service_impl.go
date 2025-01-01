package services

import (
	"context"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/google/go-github/v68/github"
	"github.com/samber/lo"

	"github.com/woxQAQ/frp-webconsole/pkg/models"
	"github.com/woxQAQ/frp-webconsole/pkg/stores"
	"github.com/woxQAQ/frp-webconsole/pkg/types"
)

type FrpcServiceImpl struct {
	ghStore stores.GithubClient
}

func (s *FrpcServiceImpl) GetFrpcConfig(ctx context.Context) (string, error) {
	return "", nil
}

func (s *FrpcServiceImpl) InstallFrpc(ctx context.Context) error {
	_ = v1.ClientCommonConfig{}
	return nil
}

func (s *FrpcServiceImpl) ListFrpRelease(ctx context.Context, page, pageSize int, system models.System) ([]models.FrpRelease, error) {
	releases, err := s.ghStore.GetReleaseList(ctx, types.FrpOwner, types.FrpRepo)
	if err != nil {
		return nil, err
	}
	return lo.FilterMap(releases, func(release *github.RepositoryRelease, _ int) (models.FrpRelease, bool) {
		frpRelease, err := models.NewFrpRelease(release, system)
		if err != nil {
			return models.FrpRelease{}, false
		}
		return *frpRelease, true
	}), nil
}
