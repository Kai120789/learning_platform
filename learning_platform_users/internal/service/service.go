package service

import (
	"go.uber.org/zap"
	"learning-platform/users/internal/service/group"
	"learning-platform/users/internal/service/user"
)

type Service struct {
	UserService  *user.UserService
	GroupService *group.GroupService
}

func New(
	logger *zap.Logger,
	userStorage *user.UserStorage,
	groupStorage *group.GroupStorage,
) *Service {
	userService := user.NewUserService(logger, userStorage)
	return &Service{
		UserService:  userService,
		GroupService: group.NewGroupService(logger, groupStorage, userService.UserBaseService),
	}
}
