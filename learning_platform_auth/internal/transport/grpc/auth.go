package grpc

import (
	"context"
	authGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/auth/internal/dto"
)

type AuthService interface {
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	Register(request dto.RegisterRequest) (*dto.RegisterResponse, error)
	RefreshTokens(accessToken string) (*string, error)
	CheckPassword(password string, passwordHash string) (bool, error)
	GeneratePasswordHash(password string) (*string, error)
	Logout(accessToken string) error
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
	res, err := g.service.Login(dto.LoginRequest{
		UserId:   in.GetUserId(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to login user")
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
	request := dto.RegisterRequest{
		UserId:   in.GetUserId(),
		Email:    in.GetEmail(),
		Name:     in.GetName(),
		Surname:  in.GetSurname(),
		LastName: in.GetLastName(),
		Role:     protoAuthRoleToString(in.GetRole()),
		Password: in.GetPassword(),
	}

	res, err := g.service.Register(request)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &authGRPC.RegisterResponse{
		UserId:      res.UserId,
		AccessToken: res.AccessToken,
	}, nil
}

func (g *AuthGRPCServer) RefreshTokens(
	ctx context.Context,
	in *authGRPC.RefreshTokensRequest,
) (*authGRPC.RefreshTokensResponse, error) {
	newAccessToken, err := g.service.RefreshTokens(in.GetAccessToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to refresh tokens")
	}

	return &authGRPC.RefreshTokensResponse{
		AccessToken: *newAccessToken,
	}, nil
}

func (g *AuthGRPCServer) CheckPassword(
	ctx context.Context,
	in *authGRPC.CheckPasswordRequest,
) (*authGRPC.CheckPasswordResponse, error) {
	isValid, err := g.service.CheckPassword(in.GetPassword(), in.GetPasswordHash())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed valid password")
	}

	return &authGRPC.CheckPasswordResponse{IsValid: isValid}, nil
}

func (g *AuthGRPCServer) GeneratePasswordHash(
	ctx context.Context,
	in *authGRPC.GeneratePasswordHashRequest,
) (*authGRPC.GeneratePasswordHashResponse, error) {
	passwordHash, err := g.service.GeneratePasswordHash(in.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed generate hash for password")
	}

	return &authGRPC.GeneratePasswordHashResponse{PasswordHash: *passwordHash}, nil
}

func (g *AuthGRPCServer) Logout(
	ctx context.Context,
	in *authGRPC.LogoutRequest,
) (*authGRPC.LogoutResponse, error) {
	err := g.service.Logout(in.AccessToken)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed logout user")
	}

	return &authGRPC.LogoutResponse{}, nil
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

func protoAuthRoleToString(role authGRPC.UserRole) string {
	switch role {
	case authGRPC.UserRole_TUTOR:
		return "TUTOR"
	case authGRPC.UserRole_STUDENT:
		return "STUDENT"
	default:
		return "UNSPECIFIED"
	}
}
