package services

import (
	"context"

	"github.com/woxQAQ/frp-webconsole/pkg/gen/frpc"
	"github.com/woxQAQ/frp-webconsole/pkg/stores"
)

type FrpcServiceImpl struct {
	ghStore stores.GithubClient
}

// ListFrpRelease implements ListFrpRelease.
func (f *FrpcServiceImpl) ListFrpRelease(ctx context.Context, payload *frpc.ListFrpReleasePayload) (res *frpc.FrpRelease, err error) {
	panic("not implemented") // TODO: Implement
}

func NewFrpcService(ghStore stores.GithubClient) frpc.Service {
	return &FrpcServiceImpl{
		ghStore: ghStore,
	}
}
