package group

import (
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type GroupUserService struct {
	logger  *zap.Logger
	storage GroupUserStorage
	user    GetUserService
	group   GetGroupService
}

type GroupUserStorage interface {
	AddUsersToGroup(userId int64, groupId int64) error
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
	logger *zap.Logger,
	storage GroupUserStorage,
	user GetUserService,
	group GetGroupService,
) *GroupUserService {
	return &GroupUserService{
		logger:  logger,
		storage: storage,
		user:    user,
		group:   group,
	}
}

func (g *GroupUserService) AddUsersToGroup(userId int64, groupId int64) ([]dto.ShortUserInfo, error) {
	user, err := g.user.GetUserById(userId)
	if user == nil {
		g.logger.Error(fmt.Sprintf("user with id %d not found", userId), zap.Error(err))
		return nil, err
	} else if err != nil {
		g.logger.Error("failed get user by id", zap.Error(err))
		return nil, err
	}

	group, err := g.group.GetGroupById(groupId)
	if group == nil {
		g.logger.Error(fmt.Sprintf("group with id %d not found", groupId), zap.Error(err))
		return nil, err
	} else if err != nil {
		g.logger.Error("failed get group by id", zap.Error(err))
		return nil, err
	}

	err = g.storage.AddUsersToGroup(userId, groupId)
	if err != nil {
		return nil, err
	}

	return g.storage.GetGroupUsers(groupId)
}

func (g *GroupUserService) RemoveUserFromGroup(userId int64, groupId int64) error {
	user, err := g.user.GetUserById(userId)
	if user == nil {
		g.logger.Error(fmt.Sprintf("user with id %d not found", userId), zap.Error(err))
		return err
	} else if err != nil {
		g.logger.Error("failed get user by id", zap.Error(err))
		return err
	}

	group, err := g.group.GetGroupById(groupId)
	if group == nil {
		g.logger.Error(fmt.Sprintf("group with id %d not found", groupId), zap.Error(err))
		return err
	} else if err != nil {
		g.logger.Error("failed get group by id", zap.Error(err))
		return err
	}

	err = g.storage.RemoveUserFromGroup(userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupUserService) GetUserGroups(userId int64) ([]models.Group, error) {
	user, err := g.user.GetUserById(userId)
	if user == nil {
		g.logger.Error(fmt.Sprintf("user with id %d not found", userId), zap.Error(err))
		return nil, err
	} else if err != nil {
		g.logger.Error("failed get user by id", zap.Error(err))
		return nil, err
	}

	res, err := g.storage.GetUserGroups(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupUserService) GetGroupsByTutorId(tutorId int64) ([]models.Group, error) {
	user, err := g.user.GetUserById(tutorId)
	if user == nil {
		g.logger.Error(fmt.Sprintf("tutor with id %d not found", tutorId), zap.Error(err))
		return nil, err
	} else if err != nil {
		g.logger.Error("failed get user by id", zap.Error(err))
		return nil, err
	}

	res, err := g.storage.GetGroupsByTutorId(tutorId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupUserService) GetGroupUsers(groupId int64) ([]dto.ShortUserInfo, error) {
	group, err := g.group.GetGroupById(groupId)
	if group == nil {
		g.logger.Error(fmt.Sprintf("group with id %d not found", groupId), zap.Error(err))
		return nil, err
	} else if err != nil {
		g.logger.Error("failed get group by id", zap.Error(err))
		return nil, err
	}

	res, err := g.storage.GetGroupUsers(groupId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
