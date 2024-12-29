package services

import (
	"context"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/woxQAQ/frp-webconsole/pkg/models"
)

type FrpcServiceImpl struct {
}

func (s *FrpcServiceImpl) GetFrpcConfig(ctx context.Context) (string, error) {
	return "", nil
}

func (s *FrpcServiceImpl) InstallFrpc(ctx context.Context) error {
	_ = v1.ClientCommonConfig{}
	return nil
}

func (s *FrpcServiceImpl) ListFrpRelease(ctx context.Context) ([]models.FrpRelease, error) {
	return nil, nil
}
