package grpc

import (
	"context"
	authGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/auth"
	"learning-platform/auth/internal/dto"
)

type AuthService interface {
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	Register()
	RefreshTokens()
	Logout()
	LogoutAll()
	ChangePassword()
	ForceChangePassword()
	ChangeEmail()
	ForceChangeEmail()
}

type AuthGRPCServer struct {
	authGRPC.UnimplementedAuthServer
	service AuthService
}

func NewAuthGRPCServer(service AuthService) authGRPC.AuthServer {
	return &AuthGRPCServer{
		service: service,
	}
}

func (g *AuthGRPCServer) Login(
	ctx context.Context,
	in *authGRPC.LoginRequest,
) (*authGRPC.LoginResponse, error) {
	email := in.GetEmail()

	password := in.GetPassword()

	res, err := g.service.Login(dto.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &authGRPC.LoginResponse{
		AccessToken: res.AccessToken,
		UserId:      res.UserId,
	}, nil
}

func (g *AuthGRPCServer) Register(
	ctx context.Context,
	in *authGRPC.RegisterRequest,
) (*authGRPC.RegisterResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) RefreshTokens(
	ctx context.Context,
	in *authGRPC.RefreshTokensRequest,
) (*authGRPC.RefreshTokensResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) Logout(
	ctx context.Context,
	in *authGRPC.LogoutRequest,
) (*authGRPC.LogoutResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) LogoutAll(
	ctx context.Context,
	in *authGRPC.LogoutAllRequest,
) (*authGRPC.LogoutAllResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) ChangePassword(
	ctx context.Context,
	in *authGRPC.ChangePasswordRequest,
) (*authGRPC.ChangePasswordResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) ForceChangePassword(
	ctx context.Context,
	in *authGRPC.ForceChangePasswordRequest,
) (*authGRPC.ForceChangePasswordResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) ChangeEmail(
	ctx context.Context,
	in *authGRPC.ChangeEmailRequest,
) (*authGRPC.ChangeEmailResponse, error) {
	return nil, nil
}

func (g *AuthGRPCServer) ForceChangeEmail(
	ctx context.Context,
	in *authGRPC.ForceChangeEmailRequest,
) (*authGRPC.ForceChangeEmailResponse, error) {
	return nil, nil
}
