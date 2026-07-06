package service

import (
	"fmt"
	"learning-platform/api-gateway/internal/dto/groupDto"
	"learning-platform/api-gateway/internal/dto/userDto"
)

type GroupService struct {
	client      GroupClient
	userService UserGroupService
}

type GroupClient interface {
	CreateGroup(group groupDto.CreateGroupRequest) (*groupDto.GroupResponse, error)
	UpdateGroup(groupId int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error)
	RemoveGroup(groupId int64) error
	GetGroupById(groupId int64) (*groupDto.GroupResponse, error)
	GetGroups() ([]groupDto.GroupResponse, error)
	AddUsersToGroup(groupId int64, userIds []int64) ([]int64, error)
	RemoveUserFromGroup(userId int64, groupId int64) error
	GetUserGroups(userId int64) ([]groupDto.GroupResponse, error)
	GetGroupsByTutorId(tutorId int64) ([]groupDto.GroupResponse, error)
	GetGroupUsers(groupId int64) ([]int64, error)
}

type UserGroupService interface {
	GetUserById(id int64) (*userDto.GetUser, error)
}

func NewGroupService(client GroupClient, userService UserGroupService) *GroupService {
	return &GroupService{
		client:      client,
		userService: userService,
	}
}

func (g *GroupService) CreateGroup(group groupDto.CreateGroupRequest) (*groupDto.GroupResponse, error) {
	res, err := g.client.CreateGroup(group)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) UpdateGroup(groupId int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error) {
	res, err := g.client.UpdateGroup(groupId, newGroup)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) RemoveGroup(groupId int64) error {
	err := g.client.RemoveGroup(groupId)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) GetGroupById(groupId int64) (*groupDto.GroupResponse, error) {
	res, err := g.client.GetGroupById(groupId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) GetGroups() ([]groupDto.GroupResponse, error) {
	res, err := g.client.GetGroups()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) AddUsersToGroup(groupId int64, userIds []int64) ([]int64, error) {
	res, err := g.client.AddUsersToGroup(groupId, userIds)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) RemoveUserFromGroup(userId int64, groupId int64) error {
	user, err := g.userService.GetUserById(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user does not exists, %w", err)
	}

	err = g.client.RemoveUserFromGroup(userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) GetUserGroups(userId int64) ([]groupDto.GroupResponse, error) {
	user, err := g.userService.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user does not exists, %w", err)
	}

	res, err := g.client.GetUserGroups(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) GetGroupsByTutorId(tutorId int64) ([]groupDto.GroupResponse, error) {
	user, err := g.userService.GetUserById(tutorId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user does not exists, %w", err)
	}

	res, err := g.client.GetGroupsByTutorId(tutorId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) GetGroupUsers(groupId int64) ([]int64, error) {
	res, err := g.client.GetGroupUsers(groupId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
