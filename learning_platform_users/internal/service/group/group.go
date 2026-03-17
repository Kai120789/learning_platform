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
) *GroupService {
	return &GroupService{
		GroupBaseService: NewGroupBaseService(logger, storage.GroupBaseStorage),
		GroupUserService: NewGroupUserService(logger, storage.GroupUserStorage),
	}
}
