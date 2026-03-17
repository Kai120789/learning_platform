package group

import (
	"github.com/Kai120789/learning_platform_models/models"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type GroupBaseService struct {
	logger  *zap.Logger
	storage GroupBaseStorage
}

type GroupBaseStorage interface {
	CreateGroup(groupDto dto.CreateGroup) (*int64, error)
	UpdateGroup(id int64, groupDto dto.UpdateGroup) error
	RemoveGroup(id int64) error
	GetGroupById(id int64) (*models.Group, error)
	GetGroups() ([]models.Group, error)
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

func (g *GroupBaseService) CreateGroup(groupDto dto.CreateGroup) (*models.Group, error) {
	groupId, err := g.storage.CreateGroup(groupDto)
	if err != nil {
		g.logger.Error("create group error", zap.Error(err))
		return nil, err
	}

	return g.GetGroupById(*groupId)
}

func (g *GroupBaseService) UpdateGroup(id int64, groupDto dto.UpdateGroup) (*models.Group, error) {
	err := g.storage.UpdateGroup(id, groupDto)
	if err != nil {
		g.logger.Error("update group error", zap.Error(err))
		return nil, err
	}

	return g.GetGroupById(id)
}

func (g *GroupBaseService) RemoveGroup(id int64) error {
	err := g.storage.RemoveGroup(id)
	if err != nil {
		g.logger.Error("remove group error", zap.Error(err))
		return err
	}

	return nil
}

func (g *GroupBaseService) GetGroupById(id int64) (*models.Group, error) {
	return g.storage.GetGroupById(id)
}

func (g *GroupBaseService) GetGroups() ([]models.Group, error) {
	return g.storage.GetGroups()
}
