package grpc

import (
	authGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client authGRPC.AuthClient
	logger *zap.Logger
}

func NewAuthGrpcConnection(authGrpcUrl string, logger *zap.Logger) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		authGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("failed to create auth grpc client", zap.Error(err))
		return nil, err
	}

	return conn, nil
}

func NewAuthClient(connection *grpc.ClientConn, logger *zap.Logger) *AuthClient {
	return &AuthClient{
		client: authGRPC.NewAuthClient(connection),
		logger: logger,
	}
}
