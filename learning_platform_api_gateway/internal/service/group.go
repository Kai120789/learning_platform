package service

import (
	"learning-platform/api-gateway/internal/dto/groupDto"
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"learning-platform/api-gateway/internal/dto/userDto"
)

type GroupService struct {
	client         GroupClient
	userService    GroupUserService
	subjectService GroupSubjectService
}

type GroupClient interface {
	CreateGroup(group groupDto.CreateGroupRequest, tutorID int64) (*groupDto.GroupResponse, error)
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

type GroupUserService interface {
	GetUsersShortInfo(userIDs []int64) ([]userDto.UserShortInfo, error)
}

type GroupSubjectService interface {
	GetOneSubject(subjectID int64) (*subjectDto.Subject, error)
	GetAllSubjects() ([]subjectDto.Subject, error)
}

func NewGroupService(
	client GroupClient,
	userService GroupUserService,
	subjectService GroupSubjectService,
) *GroupService {
	return &GroupService{
		client:         client,
		userService:    userService,
		subjectService: subjectService,
	}
}

func (g *GroupService) CreateGroup(group groupDto.CreateGroupRequest, tutorID int64) (*groupDto.GroupFullResponse, error) {
	res, err := g.client.CreateGroup(group, tutorID)
	if err != nil {
		return nil, err
	}

	subject, err := g.subjectService.GetOneSubject(group.SubjectID)
	if err != nil {
		return nil, err
	}

	resGroup := &groupDto.GroupFullResponse{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Subject:     *subject,
		Users:       nil,
		TutorID:     res.TutorID,
		TgGroupLink: res.TgGroupLink,
		TgChatID:    res.TgChatID,
	}

	return resGroup, nil
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
	err := g.client.RemoveUserFromGroup(userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) GetUserGroups(userId int64) ([]groupDto.GroupFullResponse, error) {
	var resGroups []groupDto.GroupFullResponse
	res, err := g.client.GetUserGroups(userId)
	if err != nil {
		return nil, err
	}

	subjects, err := g.subjectService.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	subjectByID := make(map[int64]subjectDto.Subject)

	for _, oneSubject := range subjects {
		subjectByID[oneSubject.ID] = oneSubject
	}

	for _, oneGroup := range res {
		resGroups = append(resGroups, groupDto.GroupFullResponse{
			ID:          oneGroup.ID,
			Title:       oneGroup.Title,
			Description: oneGroup.Description,
			Subject:     subjectByID[oneGroup.SubjectID],
			TutorID:     oneGroup.TutorID,
			TgGroupLink: oneGroup.TgGroupLink,
			TgChatID:    oneGroup.TgChatID,
		})
	}

	return resGroups, nil
}

func (g *GroupService) GetGroupsByTutorId(tutorId int64) ([]groupDto.GroupFullResponse, error) {
	res, err := g.client.GetGroupsByTutorId(tutorId)
	if err != nil {
		return nil, err
	}

	resGroups, err := g.mapGroupsWithSubjectAndUsersDTO(res)
	if err != nil {
		return nil, err
	}

	return resGroups, nil
}

func (g *GroupService) GetGroupUsers(groupId int64) ([]int64, error) {
	res, err := g.client.GetGroupUsers(groupId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) mapGroupsWithSubjectAndUsersDTO(groups []groupDto.GroupResponse) ([]groupDto.GroupFullResponse, error) {
	var resGroups []groupDto.GroupFullResponse

	subjects, err := g.subjectService.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	subjectByID := make(map[int64]subjectDto.Subject)

	for _, oneSubject := range subjects {
		subjectByID[oneSubject.ID] = oneSubject
	}

	for _, oneGroup := range groups {
		userIDs, err := g.client.GetGroupUsers(oneGroup.ID)
		if err != nil {
			return nil, err
		}

		users, err := g.userService.GetUsersShortInfo(userIDs)
		if err != nil {
			return nil, err
		}

		resGroups = append(resGroups, groupDto.GroupFullResponse{
			ID:          oneGroup.ID,
			Title:       oneGroup.Title,
			Description: oneGroup.Description,
			Subject:     subjectByID[oneGroup.SubjectID],
			TutorID:     oneGroup.TutorID,
			Users:       users,
			TgGroupLink: oneGroup.TgGroupLink,
			TgChatID:    oneGroup.TgChatID,
		})
	}

	return resGroups, nil
}
