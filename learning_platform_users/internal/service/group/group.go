package group

import "go.uber.org/zap"

type GroupService struct {
	GroupBaseService *GroupBaseService
	GroupUserService *GroupUserService
}

type GroupStorage struct {
	GroupBaseStorage GroupBaseStorage
	GroupUserStorage GroupUserStorage
}

func NewGroupService(
	logger *zap.Logger,
	storage *GroupStorage,
	userService GetUserService,
) *GroupService {
	groupBaseService := NewGroupBaseService(logger, storage.GroupBaseStorage)

	return &GroupService{
		GroupBaseService: groupBaseService,
		GroupUserService: NewGroupUserService(logger, storage.GroupUserStorage, userService, groupBaseService),
	}
}
