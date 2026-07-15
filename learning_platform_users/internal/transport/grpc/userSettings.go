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

type UserSettingsService interface {
	UpdateUserSettings(userSettings dto.UserSettingsRequest) (*models.UserSettings, error)
	UpdateUserTheme(userID int64, theme enum.UserTheme) error
}

func (u *UserGRPCServer) UpdateUserSettings(
	ctx context.Context,
	in *userGRPC.UpdateUserSettingsRequest,
) (*userGRPC.UpdateUserSettingsResponse, error) {
	userSettings := dto.UserSettingsRequest{
		UserID:                 in.GetUserId(),
		Is2FaEnabled:           in.GetIs_2FaEnabled(),
		IsNotificationsEnabled: in.GetIsNotificationsEnabled(),
		Language:               protoToEnumLanguage(in.GetLanguage()),
	}

	res, err := u.UserSettingsService.UpdateUserSettings(userSettings)
	if err != nil {
		u.logger.Error(
			"failed to update user settings",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user settings")
	}

	return &userGRPC.UpdateUserSettingsResponse{
		Is_2FaEnabled:          res.Is2FaEnabled,
		IsNotificationsEnabled: res.IsNotificationsEnabled,
		Language:               enumToProtoLanguage(res.Language),
		Theme:                  enumToProtoTheme(res.Theme),
	}, nil
}

func (u *UserGRPCServer) UpdateUserTheme(
	ctx context.Context,
	in *userGRPC.UpdateUserThemeRequest,
) (*userGRPC.UpdateUserThemeResponse, error) {
	err := u.UserSettingsService.UpdateUserTheme(
		in.GetUserId(),
		protoToEnumTheme(in.GetTheme()),
	)
	if err != nil {
		u.logger.Error(
			"failed to update user theme",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user theme")
	}

	return &userGRPC.UpdateUserThemeResponse{}, nil
}
