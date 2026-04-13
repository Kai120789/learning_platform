package service

import (
	"learning-platform/users/internal/service/group"
	"learning-platform/users/internal/service/user"
)

type Service struct {
	UserService  *user.UserService
	GroupService *group.GroupService
}

func New(
	userStorage *user.UserStorage,
	groupStorage *group.GroupStorage,
) *Service {
	userService := user.NewUserService(userStorage)
	return &Service{
		UserService:  userService,
		GroupService: group.NewGroupService(groupStorage, userService.UserBaseService),
	}
}
