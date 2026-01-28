package grpc

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
)

type UserService interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUser(userId int64) (*models.User, error)
	ChangePassword(userId int64, newPasswordHash string) error
	ChangeEmail(userId int64, newEmail string) error
	GetUserData(userId int64) (*dto.UserData, error)
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

	userId, err := g.UserService.CreateUser(userDto)
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
	res, err := g.UserService.GetUser(in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &userGRPC.GetUserResponse{
		UserId:       res.Id,
		Email:        res.Email,
		PasswordHash: res.Password,
	}, nil
}

func (g *UserGRPCServer) GetUserData(
	ctx context.Context,
	in *userGRPC.GetUserDataRequest,
) (*userGRPC.GetUserDataResponse, error) {
	res, err := g.UserService.GetUserData(in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user data")
	}

	return &userGRPC.GetUserDataResponse{
		UserId: res.UserId,
		Email:  res.Email,
		UserInfo: &userGRPC.UpdateUserInfoResponse{
			Name:     res.UserInfo.Name,
			Surname:  res.UserInfo.Surname,
			Lastname: res.UserInfo.Lastname,
			City:     res.UserInfo.City,
			About:    res.UserInfo.About,
			Role:     stringToProtoRole(res.UserInfo.Role),
			Status:   stringToProtoStatus(res.UserInfo.Status),
			Class:    res.UserInfo.Class,
		},
		UserSettings: &userGRPC.UpdateUserSettingsResponse{
			Is_2FaEnabled:          res.UserSettings.Is2FaEnabled,
			IsNotificationsEnabled: res.UserSettings.IsNotificationsEnabled,
		},
	}, nil
}

func (g *UserGRPCServer) ChangePassword(
	ctx context.Context,
	in *userGRPC.ChangePasswordRequest,
) (*userGRPC.ChangePasswordResponse, error) {
	err := g.UserService.ChangePassword(in.GetUserId(), in.GetNewPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to change user password")
	}
	return &userGRPC.ChangePasswordResponse{}, nil
}

func (g *UserGRPCServer) ChangeEmail(
	ctx context.Context,
	in *userGRPC.ChangeEmailRequest,
) (*userGRPC.ChangeEmailResponse, error) {
	err := g.UserService.ChangeEmail(in.GetUserId(), in.GetNewEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to change user email")
	}
	return &userGRPC.ChangeEmailResponse{}, nil
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

func stringToProtoRole(role string) userGRPC.UserRole {
	switch role {
	case "TUTOR":
		return userGRPC.UserRole_TUTOR
	case "STUDENT":
		return userGRPC.UserRole_STUDENT
	case "ADMIN":
		return userGRPC.UserRole_ADMIN
	default:
		return userGRPC.UserRole_USER_ROLE_UNSPECIFIED
	}
}

func stringToProtoStatus(status string) userGRPC.Status {
	switch status {
	case "ACTIVE":
		return userGRPC.Status_ACTIVE
	case "INACTIVE":
		return userGRPC.Status_INACTIVE
	case "BANNED":
		return userGRPC.Status_BANNED
	default:
		return userGRPC.Status_STATUS_UNSPECIFIED
	}
}
