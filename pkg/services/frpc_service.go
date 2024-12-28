package services

type FrpcService interface {
	GetFrpcConfig() (string, error)
	InstallFrpc() error
}
