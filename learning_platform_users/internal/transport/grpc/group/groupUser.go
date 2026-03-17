package group

import (
	"context"
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
)

type GroupUserService interface {
	AddUsersToGroup() error
	RemoveUserFromGroup() error
	GetUserGroups() error
	GetGroupsByTutorId() error
	GetGroupUsers() error
}

func AddUsersToGroup(
	ctx context.Context,
	in *groupGRPC.AddUsersToGroupRequest,
) (*groupGRPC.AddUsersToGroupResponse, error) {
	return nil, nil
}

func RemoveUserFromGroup(
	ctx context.Context,
	in *groupGRPC.RemoveUserFromGroupRequest,
) (*groupGRPC.RemoveUserFromGroupResponse, error) {
	return nil, nil
}

func GetUserGroups(
	ctx context.Context,
	in *groupGRPC.GetUserGroupsRequest,
) (*groupGRPC.GetUserGroupsResponse, error) {
	return nil, nil
}

func GetGroupsByTutorId(
	ctx context.Context,
	in *groupGRPC.GetGroupsByTutorIdRequest,
) (*groupGRPC.GetGroupsByTutorIdResponse, error) {
	return nil, nil
}

func GetGroupUsers(
	ctx context.Context,
	in *groupGRPC.GetGroupUsersRequest,
) (*groupGRPC.GetGroupUsersResponse, error) {
	return nil, nil
}
