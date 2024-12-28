package services

type FrpcServiceImpl struct {
}

func (s *FrpcServiceImpl) GetFrpcConfig() (string, error) {
	return "", nil
}

func (s *FrpcServiceImpl) InstallFrpc() error {
	panic("not implemented")
}
