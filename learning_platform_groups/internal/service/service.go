package service

type Service struct {
	GroupBaseService *GroupBaseService
	GroupUserService *GroupUserService
}

type Storage struct {
	GroupBaseStorage GroupBaseStorage
	GroupUserStorage GroupUserStorage
}

func New(
	storage *Storage,
) *Service {
	groupBaseService := NewGroupBaseService(storage.GroupBaseStorage)

	return &Service{
		GroupBaseService: groupBaseService,
		GroupUserService: NewGroupUserService(storage.GroupUserStorage, groupBaseService),
	}
}
