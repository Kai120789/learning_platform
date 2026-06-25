package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/models/enum"
)

type UserBaseService interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUserById(userID int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ChangePassword(userID int64, newPasswordHash string) error
	ChangeEmail(userID int64, newEmail string) error
	GetUserData(userID int64) (*dto.UserData, error)
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
		Role:         protoRoleToEnum(in.GetRole()),
		PasswordHash: in.GetPasswordHash(),
	}

	userID, err := g.UserBaseService.CreateUser(userDto)
	if err != nil {
		g.logger.Error(
			"failed to create user",
			zap.String("email", in.GetEmail()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userGRPC.CreateUserResponse{
		UserId: *userID,
	}, nil
}

func (g *UserGRPCServer) GetUserById(
	ctx context.Context,
	in *userGRPC.GetUserByIdRequest,
) (*userGRPC.GetUserByIdResponse, error) {
	res, err := g.UserBaseService.GetUserById(in.GetUserId())
	if err != nil {
		g.logger.Error(
			"failed to get user",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &userGRPC.GetUserByIdResponse{
		UserId:       res.ID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Role:         enumToProtoRole(res.Role),
		Status:       enumToProtoStatus(res.Status),
	}, nil
}

func (g *UserGRPCServer) GetUserByEmail(
	ctx context.Context,
	in *userGRPC.GetUserByEmailRequest,
) (*userGRPC.GetUserByEmailResponse, error) {
	res, err := g.UserBaseService.GetUserByEmail(in.GetEmail())
	if err != nil {
		g.logger.Error(
			"failed to get user by email",
			zap.String("userEmail", in.GetEmail()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &userGRPC.GetUserByEmailResponse{
		UserId:       res.ID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Role:         enumToProtoRole(res.Role),
		Status:       enumToProtoStatus(res.Status),
	}, nil
}

func (g *UserGRPCServer) GetUserData(
	ctx context.Context,
	in *userGRPC.GetUserDataRequest,
) (*userGRPC.GetUserDataResponse, error) {
	res, err := g.UserBaseService.GetUserData(in.GetUserId())
	if err != nil {
		g.logger.Error(
			"failed to get user data",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get user data")
	}

	return &userGRPC.GetUserDataResponse{
		UserId: res.UserID,
		Email:  res.Email,
		Role:   enumToProtoRole(res.Role),
		Status: enumToProtoStatus(res.Status),
		UserInfo: &userGRPC.UpdateUserInfoResponse{
			Name:     res.UserInfo.Name,
			Surname:  res.UserInfo.Surname,
			Lastname: res.UserInfo.Lastname,
			City:     res.UserInfo.City,
			About:    res.UserInfo.About,
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
	err := g.UserBaseService.ChangePassword(in.GetUserId(), in.GetNewPassword())
	if err != nil {
		g.logger.Error(
			"failed to change password",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to change user password")
	}
	return &userGRPC.ChangePasswordResponse{}, nil
}

func (g *UserGRPCServer) ChangeEmail(
	ctx context.Context,
	in *userGRPC.ChangeEmailRequest,
) (*userGRPC.ChangeEmailResponse, error) {
	err := g.UserBaseService.ChangeEmail(in.GetUserId(), in.GetNewEmail())
	if err != nil {
		g.logger.Error(
			"failed to change email",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to change user email")
	}
	return &userGRPC.ChangeEmailResponse{}, nil
}

func protoRoleToEnum(role userGRPC.UserRole) enum.UserRole {
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

func enumToProtoRole(role enum.UserRole) userGRPC.UserRole {
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

func enumToProtoStatus(status enum.UserStatus) userGRPC.Status {
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
