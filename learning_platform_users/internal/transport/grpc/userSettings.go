package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
)

type UserSettingsService interface {
	UpdateUserSettings(userSettings dto.UserSettings) (*models.UserSettings, error)
}

func (g *UserGRPCServer) UpdateUserSettings(
	ctx context.Context,
	in *userGRPC.UpdateUserSettingsRequest,
) (*userGRPC.UpdateUserSettingsResponse, error) {
	userSettings := dto.UserSettings{
		UserID:                 in.GetUserId(),
		Is2FaEnabled:           in.GetIs_2FaEnabled(),
		IsNotificationsEnabled: in.GetIsNotificationsEnabled(),
	}

	res, err := g.UserSettingsService.UpdateUserSettings(userSettings)
	if err != nil {
		g.logger.Error(
			"failed to update user settings",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user settings")
	}

	return &userGRPC.UpdateUserSettingsResponse{
		Is_2FaEnabled:          res.Is2FaEnabled,
		IsNotificationsEnabled: res.IsNotificationsEnabled,
	}, nil
}
