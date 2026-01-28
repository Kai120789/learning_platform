package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
)

type RegisterUseCase interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
}

type UserGRPCServer struct {
	userGRPC.UnimplementedUserServer
	register RegisterUseCase
}

func NewUserGRPCServer(register RegisterUseCase) userGRPC.UserServer {
	return &UserGRPCServer{
		register: register,
	}
}

func (g *UserGRPCServer) CreateUser(
	ctx context.Context,
	in *userGRPC.CreateUserRequest,
) (*userGRPC.CreateUserResponse, error) {
	userDto := dto.CreateUser{
		Email:        in.Email,
		Name:         in.Name,
		Surname:      in.Surname,
		LastName:     *in.LastName,
		Role:         protoRoleToString(in.Role),
		PasswordHash: in.PasswordHash,
	}

	userId, err := g.register.CreateUser(userDto)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userGRPC.CreateUserResponse{
		UserId: *userId,
	}, nil
}

func (g *UserGRPCServer) GetUser(
	ctx context.Context,
	in *userGRPC.GetUserRequest,
) (*userGRPC.GetUserResponse, error) {
	return nil, nil
}

func (g *UserGRPCServer) GetUserData(
	ctx context.Context,
	in *userGRPC.GetUserDataRequest,
) (*userGRPC.GetUserDataResponse, error) {
	return nil, nil
}

func (g *UserGRPCServer) ChangePassword(
	ctx context.Context,
	in *userGRPC.ChangePasswordRequest,
) (*userGRPC.ChangePasswordResponse, error) {
	return nil, nil
}

func (g *UserGRPCServer) ChangeEmail(
	ctx context.Context,
	in *userGRPC.ChangeEmailRequest,
) (*userGRPC.ChangeEmailResponse, error) {
	return nil, nil
}

func (g *UserGRPCServer) UpdateUserInfo(
	ctx context.Context,
	in *userGRPC.UpdateUserInfoRequest,
) (*userGRPC.UpdateUserInfoResponse, error) {
	return nil, nil
}

func (g *UserGRPCServer) UpdateUserSettings(
	ctx context.Context,
	in *userGRPC.UpdateUserSettingsRequest,
) (*userGRPC.UpdateUserSettingsResponse, error) {
	return nil, nil
}

func protoRoleToString(role userGRPC.UserRole) string {
	switch role {
	case userGRPC.UserRole_TUTOR:
		return "TUTOR"
	case userGRPC.UserRole_STUDENT:
		return "STUDENT"
	case userGRPC.UserRole_ADMIN:
		return "ADMIN"
	default:
		return "UNSPECIFIED"
	}
}
