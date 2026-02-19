package grpc

import (
	"go.uber.org/zap"
)

type Client struct {
	UserClient *UserClient
	AuthClient *AuthClient
}

func NewClient(
	userGrpcUrl string,
	authGrpcUrl string,
	logger *zap.Logger,
) (*Client, error) {
	userGrpcConnection, err := NewUserGrpcConnection(userGrpcUrl, logger)
	if err != nil {
		logger.Error("user grpc connection error", zap.Error(err))
		return nil, err
	}

	authGrpcConnection, err := NewAuthGrpcConnection(authGrpcUrl, logger)
	if err != nil {
		logger.Error("auth grpc connection error", zap.Error(err))
		return nil, err
	}

	return &Client{
		UserClient: NewUserClient(userGrpcConnection, logger),
		AuthClient: NewAuthClient(authGrpcConnection, logger),
	}, nil
}
