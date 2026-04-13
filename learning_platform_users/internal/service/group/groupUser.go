package group

import (
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"learning-platform/users/internal/dto"
)

type GroupUserService struct {
	storage GroupUserStorage
	user    GetUserService
	group   GetGroupService
}

type GroupUserStorage interface {
	AddUsersToGroup(userIds []int64, groupId int64) error
	RemoveUserFromGroup(userId int64, groupId int64) error
	GetUserGroups(userId int64) ([]models.Group, error)
	GetGroupsByTutorId(tutorId int64) ([]models.Group, error)
	GetGroupUsers(groupId int64) ([]dto.ShortUserInfo, error)
}

type GetGroupService interface {
	GetGroupById(id int64) (*models.Group, error)
}

type GetUserService interface {
	GetUserById(userId int64) (*models.User, error)
}

func NewGroupUserService(
	storage GroupUserStorage,
	user GetUserService,
	group GetGroupService,
) *GroupUserService {
	return &GroupUserService{
		storage: storage,
		user:    user,
		group:   group,
	}
}

func (g *GroupUserService) AddUsersToGroup(userIds []int64, groupId int64) ([]dto.ShortUserInfo, error) {
	for _, userId := range userIds {
		user, err := g.user.GetUserById(userId)
		if user == nil {
			return nil, fmt.Errorf("add user to group (user not found): %w", err)
		} else if err != nil {
			return nil, fmt.Errorf("add user to group (get user): %w", err)
		}
	}

	group, err := g.group.GetGroupById(groupId)
	if group == nil {
		return nil, fmt.Errorf("add user to group (group not found): %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("add user to group (get group): %w", err)
	}

	err = g.storage.AddUsersToGroup(userIds, groupId)
	if err != nil {
		return nil, fmt.Errorf("add user to group: %w", err)
	}

	return g.storage.GetGroupUsers(groupId)
}

func (g *GroupUserService) RemoveUserFromGroup(userId int64, groupId int64) error {
	user, err := g.user.GetUserById(userId)
	if user == nil {
		return fmt.Errorf("remove user from group (user not found): %w", err)
	} else if err != nil {
		return fmt.Errorf("remove user from group (get user): %w", err)
	}

	group, err := g.group.GetGroupById(groupId)
	if group == nil {
		return fmt.Errorf("remove user from group (group not found): %w", err)
	} else if err != nil {
		return fmt.Errorf("remove user from group (get group): %w", err)
	}

	err = g.storage.RemoveUserFromGroup(userId, groupId)
	if err != nil {
		return fmt.Errorf("remove userfrom group: %w", err)
	}

	return nil
}

func (g *GroupUserService) GetUserGroups(userId int64) ([]models.Group, error) {
	user, err := g.user.GetUserById(userId)
	if user == nil {
		return nil, fmt.Errorf("get user groups (user not found): %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("get user groups (get user): %w", err)
	}

	res, err := g.storage.GetUserGroups(userId)
	if err != nil {
		return nil, fmt.Errorf("get user groups: %w", err)
	}

	return res, nil
}

func (g *GroupUserService) GetGroupsByTutorId(tutorId int64) ([]models.Group, error) {
	user, err := g.user.GetUserById(tutorId)
	if user == nil {
		return nil, fmt.Errorf("get tutor groups (user not found): %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("get tutor groups (get user): %w", err)
	}

	res, err := g.storage.GetGroupsByTutorId(tutorId)
	if err != nil {
		return nil, fmt.Errorf("get tutor groups: %w", err)
	}

	return res, nil
}

func (g *GroupUserService) GetGroupUsers(groupId int64) ([]dto.ShortUserInfo, error) {
	group, err := g.group.GetGroupById(groupId)
	if group == nil {
		return nil, fmt.Errorf("get group users (group not found): %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("get group users (get group): %w", err)
	}

	res, err := g.storage.GetGroupUsers(groupId)
	if err != nil {
		return nil, fmt.Errorf("get group users: %w", err)
	}

	return res, nil
}
