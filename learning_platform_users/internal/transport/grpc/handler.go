package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/service"
)

type UserGRPCServer struct {
	userGRPC.UnimplementedUserServer
	service service.Service
}

func NewUserGRPCServer(service service.Service) userGRPC.UserServer {
	return &UserGRPCServer{
		service: service,
	}
}

func (g *UserGRPCServer) CreateUser(
	ctx context.Context,
	in *userGRPC.CreateUserRequest,
) (*userGRPC.CreateUserResponse, error) {
	userDto := dto.CreateUser{
		Email:        in.GetEmail(),
		Name:         in.GetName(),
		Surname:      in.GetSurname(),
		LastName:     in.GetLastName(),
		Role:         protoRoleToString(in.GetRole()),
		PasswordHash: in.GetPasswordHash(),
	}

	userId, err := g.service.UserService.CreateUser(userDto)
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
	res, err := g.service.UserService.GetUser(in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &userGRPC.GetUserResponse{
		UserId:       res.UserId,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
	}, nil
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
	err := g.service.UserService.ChangePassword(in.GetUserId(), in.GetNewPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to change user password")
	}
	return &userGRPC.ChangePasswordResponse{}, nil
}

func (g *UserGRPCServer) ChangeEmail(
	ctx context.Context,
	in *userGRPC.ChangeEmailRequest,
) (*userGRPC.ChangeEmailResponse, error) {
	err := g.service.UserService.ChangeEmail(in.GetUserId(), in.GetNewEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to change user email")
	}
	return &userGRPC.ChangeEmailResponse{}, nil
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
