package grpc

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
)

type UserSettingsService interface {
	UpdateUserSettings(userSettings dto.UserSettings) (*models.UserSettings, error)
}

func (g *UserGRPCServer) UpdateUserSettings(
	ctx context.Context,
	in *userGRPC.UpdateUserSettingsRequest,
) (*userGRPC.UpdateUserSettingsResponse, error) {
	userSettings := dto.UserSettings{
		UserId:                 in.GetUserId(),
		Is2FaEnabled:           in.GetIs_2FaEnabled(),
		IsNotificationsEnabled: in.GetIsNotificationsEnabled(),
	}

	res, err := g.UserSettingsService.UpdateUserSettings(userSettings)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update user settings")
	}

	return &userGRPC.UpdateUserSettingsResponse{
		Is_2FaEnabled:          res.Is2FaEnabled,
		IsNotificationsEnabled: res.IsNotificationsEnabled,
	}, nil
}
