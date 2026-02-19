package grpc

import (
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userGRPC.UserClient
	logger *zap.Logger
}

func NewUserGrpcConnection(userGrpcUrl string, logger *zap.Logger) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		userGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("failed to create user grpc client", zap.Error(err))
		return nil, err
	}

	return conn, nil
}

func NewUserClient(connection *grpc.ClientConn, logger *zap.Logger) *UserClient {
	return &UserClient{
		client: userGRPC.NewUserClient(connection),
		logger: logger,
	}
}
