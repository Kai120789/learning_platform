package service

import (
	"fmt"
	"learning-platform/groups/internal/models"
)

type GroupUserService struct {
	storage GroupUserStorage
	group   GetGroupService
}

type GroupUserStorage interface {
	AddUsersToGroup(userIDs []int64, groupID int64) error
	RemoveUserFromGroup(userID int64, groupID int64) error
	GetUserGroups(userID int64) ([]models.Group, error)
	GetGroupsByTutorId(tutorId int64) ([]models.Group, error)
	GetGroupUsers(groupID int64) ([]int64, error)
}

type GetGroupService interface {
	GetGroupById(id int64) (*models.Group, error)
}

func NewGroupUserService(
	storage GroupUserStorage,
	group GetGroupService,
) *GroupUserService {
	return &GroupUserService{
		storage: storage,
		group:   group,
	}
}

func (g *GroupUserService) AddUsersToGroup(userIDs []int64, groupID int64) ([]int64, error) {
	_, err := g.group.GetGroupById(groupID)
	if err != nil {
		return nil, fmt.Errorf("add user to group (get group): %w", err)
	}

	err = g.storage.AddUsersToGroup(userIDs, groupID)
	if err != nil {
		return nil, fmt.Errorf("add user to group: %w", err)
	}

	return g.GetGroupUsers(groupID)
}

func (g *GroupUserService) RemoveUserFromGroup(userID int64, groupID int64) error {
	_, err := g.group.GetGroupById(groupID)
	if err != nil {
		return fmt.Errorf("remove user from group (get group): %w", err)
	}

	err = g.storage.RemoveUserFromGroup(userID, groupID)
	if err != nil {
		return fmt.Errorf("remove userfrom group: %w", err)
	}

	return nil
}

func (g *GroupUserService) GetUserGroups(userID int64) ([]models.Group, error) {
	res, err := g.storage.GetUserGroups(userID)
	if err != nil {
		return nil, fmt.Errorf("get user groups: %w", err)
	}

	return res, nil
}

func (g *GroupUserService) GetGroupsByTutorId(tutorId int64) ([]models.Group, error) {
	res, err := g.storage.GetGroupsByTutorId(tutorId)
	if err != nil {
		return nil, fmt.Errorf("get tutor groups: %w", err)
	}

	return res, nil
}

func (g *GroupUserService) GetGroupUsers(groupID int64) ([]int64, error) {
	_, err := g.group.GetGroupById(groupID)
	if err != nil {
		return nil, fmt.Errorf("get group users (get group): %w", err)
	}

	res, err := g.storage.GetGroupUsers(groupID)
	if err != nil {
		return nil, fmt.Errorf("get group users: %w", err)
	}

	return res, nil
}
