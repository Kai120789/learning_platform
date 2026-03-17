package group

import (
	"context"
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
)

type GroupBaseService interface {
	CreateGroup(groupDto dto.CreateGroup) (*models.Group, error)
	UpdateGroup(id int64, groupDto dto.UpdateGroup) (*models.Group, error)
	RemoveGroup(id int64) error
	GetGroupById(id int64) (*models.Group, error)
	GetGroups() ([]models.Group, error)
}

func (g *GroupGRPCServer) CreateGroup(
	ctx context.Context,
	in *groupGRPC.CreateGroupRequest,
) (*groupGRPC.CreateGroupResponse, error) {
	tgGroupLink := in.GetTgGroupLink()
	tgChatId := in.GetTgChatId()

	createGroupDto := dto.CreateGroup{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		SubjectId:   in.GetSubjectId(),
		TutorId:     in.GetTutorId(),
		TgGroupLink: &tgGroupLink,
		TgChatId:    &tgChatId,
	}

	group, err := g.GroupBaseService.CreateGroup(createGroupDto)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed create group")
	}

	return &groupGRPC.CreateGroupResponse{
		Id:          group.Id,
		Title:       group.Title,
		Description: group.Description,
		SubjectId:   group.SubjectId,
		TutorId:     group.TutorId,
		TgGroupLink: &group.TgGroupLink,
		TgChatId:    &group.TgChatId,
	}, nil
}

func (g *GroupGRPCServer) UpdateGroup(
	ctx context.Context,
	in *groupGRPC.UpdateGroupRequest,
) (*groupGRPC.UpdateGroupResponse, error) {
	tgGroupLink := in.GetTgGroupLink()
	tgChatId := in.GetTgChatId()

	updateGroupDto := dto.UpdateGroup{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		TgGroupLink: &tgGroupLink,
		TgChatId:    &tgChatId,
	}

	group, err := g.GroupBaseService.UpdateGroup(in.GetId(), updateGroupDto)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed update group")
	}

	return &groupGRPC.UpdateGroupResponse{
		Id:          group.Id,
		Title:       group.Title,
		Description: group.Description,
		SubjectId:   group.SubjectId,
		TutorId:     group.TutorId,
		TgGroupLink: &group.TgGroupLink,
		TgChatId:    &group.TgChatId,
	}, nil

}

func (g *GroupGRPCServer) RemoveGroup(
	ctx context.Context,
	in *groupGRPC.RemoveGroupRequest,
) (*groupGRPC.RemoveGroupResponse, error) {
	err := g.GroupBaseService.RemoveGroup(in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed delete group")
	}

	return &groupGRPC.RemoveGroupResponse{}, nil
}

func (g *GroupGRPCServer) GetGroupById(
	ctx context.Context,
	in *groupGRPC.GetGroupByIdRequest,
) (*groupGRPC.GetGroupByIdResponse, error) {
	group, err := g.GroupBaseService.GetGroupById(in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed get group by id")
	}

	if group == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("group with id %d not found", in.GetId()))
	}

	return &groupGRPC.GetGroupByIdResponse{
		Id:          group.Id,
		Title:       group.Title,
		Description: group.Description,
		SubjectId:   group.SubjectId,
		TutorId:     group.TutorId,
		TgGroupLink: &group.TgGroupLink,
		TgChatId:    &group.TgChatId,
	}, nil
}

func (g *GroupGRPCServer) GetGroups(
	ctx context.Context,
	in *groupGRPC.GetGroupsRequest,
) (*groupGRPC.GetGroupsResponse, error) {
	groups, err := g.GroupBaseService.GetGroups()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed get groups")
	}

	var resGroups []*groupGRPC.GetGroupByIdResponse
	for _, group := range groups {
		resGroups = append(resGroups, &groupGRPC.GetGroupByIdResponse{
			Id:          group.Id,
			Title:       group.Title,
			Description: group.Description,
			SubjectId:   group.SubjectId,
			TutorId:     group.TutorId,
			TgGroupLink: &group.TgGroupLink,
			TgChatId:    &group.TgChatId,
		})
	}
	return &groupGRPC.GetGroupsResponse{Groups: resGroups}, nil
}
