package grpc

import (
	"context"
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/group"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/groups/internal/models"
)

type GroupUserService interface {
	RemoveUserFromGroup(userId int64, groupId int64) error
	GetUserGroups(userId int64) ([]models.Group, error)
	GetGroupsByTutorId(tutorId int64) ([]models.Group, error)
	AddUsersToGroup(userIDs []int64, groupID int64) ([]int64, error)
	GetGroupUsers(groupID int64) ([]int64, error)
}

func (g *GroupGRPCServer) AddUsersToGroup(
	ctx context.Context,
	in *groupGRPC.AddUsersToGroupRequest,
) (*groupGRPC.AddUsersToGroupResponse, error) {
	userIDs, err := g.service.GroupUserService.AddUsersToGroup(
		in.GetUserIds(),
		in.GetGroupId(),
	)
	if err != nil {
		g.logger.Error(
			"failed add user to group",
			zap.Int64s("userIDs", in.GetUserIds()),
			zap.Int64("groupID", in.GetGroupId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed add user to group")
	}

	return &groupGRPC.AddUsersToGroupResponse{
		GroupId: in.GetGroupId(),
		UserIds: userIDs,
	}, nil
}

func (g *GroupGRPCServer) RemoveUserFromGroup(
	ctx context.Context,
	in *groupGRPC.RemoveUserFromGroupRequest,
) (*groupGRPC.RemoveUserFromGroupResponse, error) {
	err := g.service.GroupUserService.RemoveUserFromGroup(
		in.GetUserId(),
		in.GetGroupId(),
	)
	if err != nil {
		g.logger.Error(
			"failed remove user from group",
			zap.Int64("userID", in.GetUserId()),
			zap.Int64("groupID", in.GetGroupId()),
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
	groups, err := g.service.GroupUserService.GetUserGroups(in.GetUserId())
	if err != nil {
		g.logger.Error(
			"failed get user groups",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get user groups")
	}

	var resGroups []*groupGRPC.GetGroupByIdResponse
	for _, oneGroup := range groups {
		resGroups = append(resGroups, mapGroupToGrpc(&oneGroup))
	}

	return &groupGRPC.GetUserGroupsResponse{
		Groups: resGroups,
	}, nil
}

func (g *GroupGRPCServer) GetGroupsByTutorId(
	ctx context.Context,
	in *groupGRPC.GetGroupsByTutorIdRequest,
) (*groupGRPC.GetGroupsByTutorIdResponse, error) {
	groups, err := g.service.GroupUserService.GetGroupsByTutorId(in.GetTutorId())
	if err != nil {
		g.logger.Error(
			"failed get tutor groups",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get tutor groups")
	}

	var resGroups []*groupGRPC.GetGroupByIdResponse
	for _, oneGroup := range groups {
		resGroups = append(resGroups, mapGroupToGrpc(&oneGroup))
	}

	return &groupGRPC.GetGroupsByTutorIdResponse{
		Groups: resGroups,
	}, nil
}

func (g *GroupGRPCServer) GetGroupUsers(
	ctx context.Context,
	in *groupGRPC.GetGroupUsersRequest,
) (*groupGRPC.GetGroupUsersResponse, error) {
	userIDs, err := g.service.GroupUserService.GetGroupUsers(in.GetGroupId())
	if err != nil {
		g.logger.Error(
			"failed get group users",
			zap.Int64("groupID", in.GetGroupId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get group users")
	}

	return &groupGRPC.GetGroupUsersResponse{
		UserIds: userIDs,
	}, nil
}
