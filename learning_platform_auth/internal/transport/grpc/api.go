package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"learning-platform/auth/internal/dto"
	"time"
)

type UserApi struct {
	client userGRPC.UserClient
	logger *zap.Logger
}

func NewUserApi(client userGRPC.UserClient, logger *zap.Logger) *UserApi {
	return &UserApi{
		client: client,
		logger: logger,
	}
}

func (a *UserApi) GetUserByEmail(email string) (*dto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := a.client.GetUserByEmail(ctx, &userGRPC.GetUserByEmailRequest{Email: email})
	if err != nil {
		a.logger.Error("failed to send get user grpc query", zap.Error(err))
		return nil, err
	}

	return &dto.GetUser{
		UserId:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}
