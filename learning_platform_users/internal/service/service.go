package service

import (
	"go.uber.org/zap"
	"learning-platform/users/internal/service/group"
	"learning-platform/users/internal/service/user"
)

type Service struct {
	GroupService group.GroupService
	UserService  user.UserService
}

func New(
	logger *zap.Logger,
	userStorage *user.UserStorage,
	groupStorage *group.GroupStorage,
) *Service {
	return &Service{
		GroupService: *group.NewGroupService(logger, groupStorage),
		UserService:  *user.NewUserService(logger, userStorage),
	}
}
