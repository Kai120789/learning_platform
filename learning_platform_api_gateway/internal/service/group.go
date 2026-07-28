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
	UpdateGroup(groupID int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error)
	RemoveGroup(groupID int64) error
	GetGroupByID(groupID int64) (*groupDto.GroupResponse, error)
	GetGroups() ([]groupDto.GroupResponse, error)
	AddUsersToGroup(groupID int64, userIDs []int64) ([]int64, error)
	RemoveUserFromGroup(userID int64, groupID int64) error
	GetGroupsByStudentID(userID int64) ([]groupDto.GroupResponse, error)
	GetGroupsByTutorID(tutorID int64) ([]groupDto.GroupResponse, error)
	GetGroupUsers(groupID int64) ([]int64, error)
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

func (g *GroupService) UpdateGroup(groupID int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error) {
	res, err := g.client.UpdateGroup(groupID, newGroup)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) RemoveGroup(groupID int64) error {
	err := g.client.RemoveGroup(groupID)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) GetGroupByID(groupID int64) (*groupDto.GroupResponse, error) {
	res, err := g.client.GetGroupByID(groupID)
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

func (g *GroupService) AddUsersToGroup(groupID int64, userIDs []int64) ([]int64, error) {
	res, err := g.client.AddUsersToGroup(groupID, userIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupService) RemoveUserFromGroup(userID int64, groupID int64) error {
	err := g.client.RemoveUserFromGroup(userID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) GetGroupsByStudentID(userID int64) ([]groupDto.GroupFullResponse, error) {
	res, err := g.client.GetGroupsByStudentID(userID)
	if err != nil {
		return nil, err
	}

	resGroups, err := g.mapGroupsWithSubjectAndUsersDTO(res)
	if err != nil {
		return nil, err
	}

	return resGroups, nil
}

func (g *GroupService) GetGroupsByTutorID(tutorID int64) ([]groupDto.GroupFullResponse, error) {
	res, err := g.client.GetGroupsByTutorID(tutorID)
	if err != nil {
		return nil, err
	}

	resGroups, err := g.mapGroupsWithSubjectAndUsersDTO(res)
	if err != nil {
		return nil, err
	}

	return resGroups, nil
}

func (g *GroupService) GetGroupUsers(groupID int64) ([]int64, error) {
	res, err := g.client.GetGroupUsers(groupID)
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
