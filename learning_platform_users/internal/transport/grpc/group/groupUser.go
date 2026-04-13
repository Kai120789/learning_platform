package group

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
)

type GroupUserService interface {
	AddUsersToGroup(userIds []int64, groupId int64) ([]dto.ShortUserInfo, error)
	RemoveUserFromGroup(userId int64, groupId int64) error
	GetUserGroups(userId int64) ([]models.Group, error)
	GetGroupsByTutorId(tutorId int64) ([]models.Group, error)
	GetGroupUsers(groupId int64) ([]dto.ShortUserInfo, error)
}

func (g *GroupGRPCServer) AddUsersToGroup(
	ctx context.Context,
	in *groupGRPC.AddUsersToGroupRequest,
) (*groupGRPC.AddUsersToGroupResponse, error) {
	users, err := g.GroupUserService.AddUsersToGroup(
		in.GetUserIds(),
		in.GetGroupId(),
	)
	if err != nil {
		g.logger.Error(
			"failed add user to group",
			zap.Int64s("userIds", in.GetUserIds()),
			zap.Int64("groupId", in.GetGroupId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed add user to group")
	}

	usersShortInfo := make([]*groupGRPC.UserShortInfo, len(users))
	for ind, oneUser := range users {
		usersShortInfo[ind] = &groupGRPC.UserShortInfo{
			Id:    oneUser.Id,
			Name:  oneUser.Name,
			Email: oneUser.Email,
		}
	}

	return &groupGRPC.AddUsersToGroupResponse{
		GroupId: in.GetGroupId(),
		Users:   usersShortInfo,
	}, nil
}

func (g *GroupGRPCServer) RemoveUserFromGroup(
	ctx context.Context,
	in *groupGRPC.RemoveUserFromGroupRequest,
) (*groupGRPC.RemoveUserFromGroupResponse, error) {
	err := g.GroupUserService.RemoveUserFromGroup(
		in.GetUserId(),
		in.GetGroupId(),
	)
	if err != nil {
		g.logger.Error(
			"failed remove user from group",
			zap.Int64("userId", in.GetUserId()),
			zap.Int64("groupId", in.GetGroupId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed remove user from group")
	}

	return &groupGRPC.RemoveUserFromGroupResponse{}, nil
}

func (g *GroupGRPCServer) GetUserGroups(
	ctx context.Context,
	in *groupGRPC.GetUserGroupsRequest,
) (*groupGRPC.GetUserGroupsResponse, error) {
	groups, err := g.GroupUserService.GetUserGroups(in.GetUserId())
	if err != nil {
		g.logger.Error(
			"failed get user groups",
			zap.Int64("userId", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get user groups")
	}

	resGroups := make([]*groupGRPC.GetGroupByIdResponse, len(groups))
	for ind, group := range groups {
		resGroups[ind] = &groupGRPC.GetGroupByIdResponse{
			Id:          group.Id,
			Title:       group.Title,
			Description: group.Description,
			SubjectId:   group.SubjectId,
			TutorId:     group.TutorId,
			TgGroupLink: &group.TgGroupLink,
			TgChatId:    &group.TgChatId,
		}
	}

	return &groupGRPC.GetUserGroupsResponse{
		Groups: resGroups,
	}, nil
}

func (g *GroupGRPCServer) GetGroupsByTutorId(
	ctx context.Context,
	in *groupGRPC.GetGroupsByTutorIdRequest,
) (*groupGRPC.GetGroupsByTutorIdResponse, error) {
	groups, err := g.GroupUserService.GetUserGroups(in.GetTutorId())
	if err != nil {
		g.logger.Error(
			"failed get tutor groups",
			zap.Int64("tutorId", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get tutor groups")
	}

	resGroups := make([]*groupGRPC.GetGroupByIdResponse, len(groups))
	for ind, group := range groups {
		resGroups[ind] = &groupGRPC.GetGroupByIdResponse{
			Id:          group.Id,
			Title:       group.Title,
			Description: group.Description,
			SubjectId:   group.SubjectId,
			TutorId:     group.TutorId,
			TgGroupLink: &group.TgGroupLink,
			TgChatId:    &group.TgChatId,
		}
	}

	return &groupGRPC.GetGroupsByTutorIdResponse{
		Groups: resGroups,
	}, nil
}

func (g *GroupGRPCServer) GetGroupUsers(
	ctx context.Context,
	in *groupGRPC.GetGroupUsersRequest,
) (*groupGRPC.GetGroupUsersResponse, error) {
	groupUsers, err := g.GroupUserService.GetGroupUsers(in.GetGroupId())
	if err != nil {
		g.logger.Error(
			"failed get group users",
			zap.Int64("groupId", in.GetGroupId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get group users")
	}

	usersShortInfo := make([]*groupGRPC.UserShortInfo, len(groupUsers))
	for ind, oneUser := range groupUsers {
		usersShortInfo[ind] = &groupGRPC.UserShortInfo{
			Id:    oneUser.Id,
			Name:  oneUser.Name,
			Email: oneUser.Email,
		}
	}

	return &groupGRPC.GetGroupUsersResponse{
		Users: usersShortInfo,
	}, nil
}
