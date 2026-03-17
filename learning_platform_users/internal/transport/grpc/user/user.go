package user

import (
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
)

type UserGRPCServer struct {
	userGRPC.UnimplementedUserServer
	UserBaseService     UserBaseService
	UserSettingsService UserSettingsService
	UserInfoService     UserInfoService
}

func NewUserGRPCServer(
	user UserBaseService,
	userSettings UserSettingsService,
	userInfo UserInfoService,
) userGRPC.UserServer {
	return &UserGRPCServer{
		UserBaseService:     user,
		UserSettingsService: userSettings,
		UserInfoService:     userInfo,
	}
}
