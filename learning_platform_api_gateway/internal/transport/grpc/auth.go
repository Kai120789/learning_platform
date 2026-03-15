package grpc

import (
	"context"
	authGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"learning-platform/api-gateway/internal/dto"
	"time"
)

type AuthClient struct {
	client authGRPC.AuthClient
	logger *zap.Logger
}

/*
LogoutAll
ChangePassword
ForceChangePassword
ChangeEmail
ForceChangeEmail
*/

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

func (a *AuthClient) Login(req dto.LoginRequest, userId int64) (*dto.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcBody := &authGRPC.LoginRequest{
		UserId:   userId,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := a.client.Login(ctx, grpcBody)
	if err != nil {
		a.logger.Error("failed to send login grpc request", zap.Error(err))
		return nil, err
	}

	return &dto.LoginResponse{
		SessionId: res.GetSessionId(),
		UserId:    res.GetUserId(),
	}, nil
}

func (a *AuthClient) Register(req dto.RegisterRequest, userId int64) (*dto.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcBody := &authGRPC.RegisterRequest{
		UserId:   userId,
		Email:    req.Email,
		Name:     req.Name,
		Surname:  req.Surname,
		LastName: &req.LastName,
		Role:     stringAuthToProtoUserRole(req.Role),
		Password: req.Password,
	}

	res, err := a.client.Register(ctx, grpcBody)
	if err != nil {
		a.logger.Error("failed to send register grpc request", zap.Error(err))
		return nil, err
	}

	return &dto.RegisterResponse{
		UserId:    res.GetUserId(),
		SessionId: res.GetSessionId(),
	}, nil
}

func (a *AuthClient) RefreshTokens(refreshToken string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	md := metadata.Pairs("authorization", "Bearer "+refreshToken)
	ctxWithCooke := metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err := a.client.RefreshTokens(ctxWithCooke, &authGRPC.RefreshTokensRequest{})
	if err != nil {
		a.logger.Error("failed to send refresh tokens grpc request", zap.Error(err))
		return nil, err
	}

	resAccessToken := res.AccessToken

	return &resAccessToken, nil
}

func (a *AuthClient) CheckPassword(password string, passwordHash string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := a.client.CheckPassword(ctx, &authGRPC.CheckPasswordRequest{
		Password:     password,
		PasswordHash: passwordHash,
	})
	if err != nil {
		a.logger.Error("failed to send check password grpc request", zap.Error(err))
		return false, err
	}

	return res.GetIsValid(), nil
}

func (a *AuthClient) GeneratePasswordHash(password string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := a.client.GeneratePasswordHash(ctx, &authGRPC.GeneratePasswordHashRequest{
		Password: password,
	})
	if err != nil {
		a.logger.Error("failed to send check password grpc request", zap.Error(err))
		return nil, err
	}

	passwordHash := res.GetPasswordHash()

	return &passwordHash, nil
}

func (a *AuthClient) Logout(accessToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	md := metadata.Pairs("authorization", "Bearer "+accessToken)
	ctxWithCooke := metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	_, err := a.client.Logout(ctxWithCooke, &authGRPC.LogoutRequest{})
	if err != nil {
		a.logger.Error("failed to send logout grpc request", zap.Error(err))
		return err
	}

	return nil
}

func (a *AuthClient) LogoutAll(userId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.client.LogoutAll(ctx, &authGRPC.LogoutAllRequest{UserId: userId})
	if err != nil {
		a.logger.Error("failed to send logout grpc request", zap.Error(err))
		return err
	}

	return nil
}

func stringAuthToProtoUserRole(role string) authGRPC.UserRole {
	switch role {
	case "TUTOR":
		return authGRPC.UserRole_TUTOR
	case "STUDENT":
		return authGRPC.UserRole_STUDENT
	default:
		return authGRPC.UserRole_ENUM_NAME_UNSPECIFIED
	}
}
