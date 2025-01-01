package services

import (
	"context"

	"github.com/woxQAQ/frp-webconsole/pkg/models"
)

type FrpcService interface {
	GetFrpcConfig(ctx context.Context) (string, error)
	InstallFrpc(ctx context.Context) error
	ListFrpRelease(ctx context.Context, page, pageSize int, system models.System) ([]models.FrpRelease, error)
}
