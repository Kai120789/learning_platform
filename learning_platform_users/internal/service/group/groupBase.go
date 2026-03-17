package group

import "go.uber.org/zap"

type GroupBaseService struct {
	logger  *zap.Logger
	storage GroupUserStorage
}

type GroupBaseStorage interface {
}

func NewGroupBaseService(
	logger *zap.Logger,
	storage GroupBaseStorage,
) *GroupBaseService {
	return &GroupBaseService{
		logger:  logger,
		storage: storage,
	}
}
