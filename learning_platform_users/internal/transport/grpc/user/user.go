package user

import (
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
)

type UserGRPCServer struct {
	userGRPC.UnimplementedUserServer
	UserBaseService     UserBaseService
	UserSettingsService UserSettingsService
	UserInfoService     UserInfoService
	logger              *zap.Logger
}

func NewUserGRPCServer(
	user UserBaseService,
	userSettings UserSettingsService,
	userInfo UserInfoService,
	logger *zap.Logger,
) userGRPC.UserServer {
	return &UserGRPCServer{
		UserBaseService:     user,
		UserSettingsService: userSettings,
		UserInfoService:     userInfo,
		logger:              logger,
	}
}
