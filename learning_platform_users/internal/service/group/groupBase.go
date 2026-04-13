package group

import (
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"learning-platform/users/internal/dto"
)

type GroupBaseService struct {
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
	storage GroupBaseStorage,
) *GroupBaseService {
	return &GroupBaseService{
		storage: storage,
	}
}

func (g *GroupBaseService) CreateGroup(groupDto dto.CreateGroup) (*models.Group, error) {
	groupId, err := g.storage.CreateGroup(groupDto)
	if err != nil {
		return nil, fmt.Errorf("create group: %w", err)
	}

	return g.GetGroupById(*groupId)
}

func (g *GroupBaseService) UpdateGroup(id int64, groupDto dto.UpdateGroup) (*models.Group, error) {
	err := g.storage.UpdateGroup(id, groupDto)
	if err != nil {
		return nil, fmt.Errorf("update group: %w", err)
	}

	return g.GetGroupById(id)
}

func (g *GroupBaseService) RemoveGroup(id int64) error {
	err := g.storage.RemoveGroup(id)
	if err != nil {
		return fmt.Errorf("remove group: %w", err)
	}

	return nil
}

func (g *GroupBaseService) GetGroupById(id int64) (*models.Group, error) {
	group, err := g.storage.GetGroupById(id)
	if err != nil {
		return nil, fmt.Errorf("get group: %w", err)
	}

	return group, nil
}

func (g *GroupBaseService) GetGroups() ([]models.Group, error) {
	groups, err := g.storage.GetGroups()
	if err != nil {
		return nil, fmt.Errorf("get groups: %w", err)
	}
	return groups, nil
}
