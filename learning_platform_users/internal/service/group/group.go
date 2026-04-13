package group

type GroupService struct {
	GroupBaseService *GroupBaseService
	GroupUserService *GroupUserService
}

type GroupStorage struct {
	GroupBaseStorage GroupBaseStorage
	GroupUserStorage GroupUserStorage
}

func NewGroupService(
	storage *GroupStorage,
	userService GetUserService,
) *GroupService {
	groupBaseService := NewGroupBaseService(storage.GroupBaseStorage)

	return &GroupService{
		GroupBaseService: groupBaseService,
		GroupUserService: NewGroupUserService(storage.GroupUserStorage, userService, groupBaseService),
	}
}
