package grpc

import (
	"context"
	authGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"learning-platform/api-gateway/internal/dto/authDto"
	"time"
)

type AuthClient struct {
	client authGRPC.AuthClient
}

/*
ChangePassword
ForceChangePassword
ChangeEmail
ForceChangeEmail
*/

func NewAuthGrpcConnection(authGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		authGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewAuthClient(connection *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		client: authGRPC.NewAuthClient(connection),
	}
}

func (a *AuthClient) Login(req authDto.LoginRequest, userId int64) (*authDto.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcBody := &authGRPC.LoginRequest{
		UserId:   userId,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := a.client.Login(ctx, grpcBody)
	if err != nil {
		return nil, err
	}

	return &authDto.LoginResponse{
		SessionID: res.GetSessionId(),
	}, nil
}

func (a *AuthClient) Register(req authDto.RegisterRequest, userId int64) (*authDto.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcBody := &authGRPC.RegisterRequest{
		UserId:   userId,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := a.client.Register(ctx, grpcBody)
	if err != nil {
		return nil, err
	}

	return &authDto.RegisterResponse{
		SessionID: res.GetSessionId(),
	}, nil
}

func (a *AuthClient) RefreshTokens(refreshToken string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	md := metadata.Pairs("authorization", "Bearer "+refreshToken)
	ctxWithCooke := metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err := a.client.RefreshTokens(ctxWithCooke, &authGRPC.RefreshTokensRequest{})
	if err != nil {
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
		return err
	}

	return nil
}

func (a *AuthClient) LogoutAll(userId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.client.LogoutAll(ctx, &authGRPC.LogoutAllRequest{UserId: userId})
	if err != nil {
		return err
	}

	return nil
}
