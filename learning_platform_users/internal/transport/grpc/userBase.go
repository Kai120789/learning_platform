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
	"time"
)

type UserBaseService interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUserById(userID int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ChangePassword(userID int64, newPasswordHash string) error
	ChangeEmail(userID int64, newEmail string) error
	GetUserData(userID int64) (*dto.UserData, error)
}

func (u *UserGRPCServer) CreateUser(
	ctx context.Context,
	in *userGRPC.CreateUserRequest,
) (*userGRPC.CreateUserResponse, error) {
	userDto := dto.CreateUser{
		Email:        in.GetEmail(),
		Name:         in.GetName(),
		Surname:      in.GetSurname(),
		Patronymic:   in.Patronymic,
		Role:         protoToEnumRole(in.GetRole()),
		Gender:       protoToEnumGender(in.GetGender()),
		Language:     protoToEnumLanguage(in.GetLanguage()),
		PasswordHash: in.GetPasswordHash(),
		BirthDate:    mapDate(in.BirthDate),
	}

	userID, err := u.UserBaseService.CreateUser(userDto)
	if err != nil {
		u.logger.Error(
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

func (u *UserGRPCServer) GetUserById(
	ctx context.Context,
	in *userGRPC.GetUserByIdRequest,
) (*userGRPC.GetUserByIdResponse, error) {
	res, err := u.UserBaseService.GetUserById(in.GetUserId())
	if err != nil {
		u.logger.Error(
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

func (u *UserGRPCServer) GetUserByEmail(
	ctx context.Context,
	in *userGRPC.GetUserByEmailRequest,
) (*userGRPC.GetUserByEmailResponse, error) {
	res, err := u.UserBaseService.GetUserByEmail(in.GetEmail())
	if err != nil {
		u.logger.Error(
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

func (u *UserGRPCServer) GetUserData(
	ctx context.Context,
	in *userGRPC.GetUserDataRequest,
) (*userGRPC.GetUserDataResponse, error) {
	res, err := u.UserBaseService.GetUserData(in.GetUserId())
	if err != nil {
		u.logger.Error(
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
			Name:       res.UserInfo.Name,
			Surname:    res.UserInfo.Surname,
			Patronymic: res.UserInfo.Patronymic,
			City:       res.UserInfo.City,
			About:      res.UserInfo.About,
			Avatar:     res.UserInfo.Avatar,
			Gender:     enumToProtoGender(res.UserInfo.Gender),
			BirthDate: &userGRPC.Date{
				Day:   int32(res.UserInfo.BirthDate.Day()),
				Month: int32(res.UserInfo.BirthDate.Month()),
				Year:  int32(res.UserInfo.BirthDate.Year()),
			},
		},
		UserSettings: &userGRPC.UpdateUserSettingsResponse{
			Is_2FaEnabled:          res.UserSettings.Is2FaEnabled,
			IsNotificationsEnabled: res.UserSettings.IsNotificationsEnabled,
			Language:               enumToProtoLanguage(res.UserSettings.Language),
			Theme:                  enumToProtoTheme(res.UserSettings.Theme),
		},
	}, nil
}

func (u *UserGRPCServer) ChangePassword(
	ctx context.Context,
	in *userGRPC.ChangePasswordRequest,
) (*userGRPC.ChangePasswordResponse, error) {
	err := u.UserBaseService.ChangePassword(in.GetUserId(), in.GetNewPassword())
	if err != nil {
		u.logger.Error(
			"failed to change password",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to change user password")
	}
	return &userGRPC.ChangePasswordResponse{}, nil
}

func (u *UserGRPCServer) ChangeEmail(
	ctx context.Context,
	in *userGRPC.ChangeEmailRequest,
) (*userGRPC.ChangeEmailResponse, error) {
	err := u.UserBaseService.ChangeEmail(in.GetUserId(), in.GetNewEmail())
	if err != nil {
		u.logger.Error(
			"failed to change email",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to change user email")
	}
	return &userGRPC.ChangeEmailResponse{}, nil
}

func protoToEnumRole(role userGRPC.UserRole) enum.UserRole {
	switch role {
	case userGRPC.UserRole_TUTOR:
		return enum.RoleTutor
	case userGRPC.UserRole_STUDENT:
		return enum.RoleStudent
	case userGRPC.UserRole_ADMIN:
		return enum.RoleAdmin
	default:
		return ""
	}
}

func protoToEnumGender(gender userGRPC.UserGender) enum.UserGender {
	switch gender {
	case userGRPC.UserGender_MALE:
		return enum.GenderMale
	case userGRPC.UserGender_FEMALE:
		return enum.GenderFemale
	default:
		return enum.GenderUnknown
	}
}

func protoToEnumLanguage(language userGRPC.UserLanguage) enum.UserLanguage {
	switch language {
	case userGRPC.UserLanguage_RU:
		return enum.LanguageRU
	case userGRPC.UserLanguage_EN:
		return enum.LanguageEN
	default:
		return ""
	}
}

func protoToEnumTheme(theme userGRPC.UserTheme) enum.UserTheme {
	switch theme {
	case userGRPC.UserTheme_LIGHT:
		return enum.ThemeLight
	case userGRPC.UserTheme_DARK:
		return enum.ThemeDark
	default:
		return ""
	}
}

func enumToProtoRole(role enum.UserRole) userGRPC.UserRole {
	switch role {
	case enum.RoleTutor:
		return userGRPC.UserRole_TUTOR
	case enum.RoleStudent:
		return userGRPC.UserRole_STUDENT
	case enum.RoleAdmin:
		return userGRPC.UserRole_ADMIN
	default:
		return userGRPC.UserRole_USER_ROLE_UNSPECIFIED
	}
}

func enumToProtoStatus(status enum.UserStatus) userGRPC.UserStatus {
	switch status {
	case enum.StatusActive:
		return userGRPC.UserStatus_ACTIVE
	case enum.StatusInactive:
		return userGRPC.UserStatus_INACTIVE
	case enum.StatusBanned:
		return userGRPC.UserStatus_BANNED
	default:
		return userGRPC.UserStatus_STATUS_UNSPECIFIED
	}
}

func enumToProtoGender(gender enum.UserGender) userGRPC.UserGender {
	switch gender {
	case enum.GenderMale:
		return userGRPC.UserGender_MALE
	case enum.GenderFemale:
		return userGRPC.UserGender_FEMALE
	default:
		return userGRPC.UserGender_ENUM_GENDER_UNSPECIFIED
	}
}

func enumToProtoTheme(theme enum.UserTheme) userGRPC.UserTheme {
	switch theme {
	case enum.ThemeLight:
		return userGRPC.UserTheme_LIGHT
	case enum.ThemeDark:
		return userGRPC.UserTheme_DARK
	default:
		return userGRPC.UserTheme_ENUM_THEME_UNSPECIFIED
	}
}

func enumToProtoLanguage(language enum.UserLanguage) userGRPC.UserLanguage {
	switch language {
	case enum.LanguageRU:
		return userGRPC.UserLanguage_RU
	case enum.LanguageEN:
		return userGRPC.UserLanguage_EN
	default:
		return userGRPC.UserLanguage_ENUM_LANGUAGE_UNSPECIFIED
	}
}

func mapDate(bd *userGRPC.Date) *time.Time {
	birthDate := time.Date(
		int(bd.GetYear()),
		time.Month(bd.GetMonth()),
		int(bd.GetDay()),
		0, 0, 0, 0,
		time.UTC,
	)

	return &birthDate
}
