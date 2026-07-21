package grpc

import (
	"context"
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/groupDto"
	"time"
)

type GroupClient struct {
	client groupGRPC.GroupClient
}

func NewGroupGrpcConnection(groupGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		groupGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewGroupClient(connection *grpc.ClientConn) *GroupClient {
	return &GroupClient{
		client: groupGRPC.NewGroupClient(connection),
	}
}

func (g *GroupClient) CreateGroup(group groupDto.CreateGroupRequest, tutorID int64) (*groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reqBody := &groupGRPC.CreateGroupRequest{
		Title:       group.Title,
		Description: group.Description,
		SubjectId:   group.SubjectID,
		TutorId:     tutorID,
		TgGroupLink: group.TgGroupLink,
		TgChatId:    group.TgChatID,
	}

	resGroup, err := g.client.CreateGroup(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	return mapGroupGrpcToDTO(resGroup.GetGroup()), nil
}

func (g *GroupClient) UpdateGroup(groupId int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reqBody := &groupGRPC.UpdateGroupRequest{
		Id:          groupId,
		Title:       newGroup.Title,
		Description: newGroup.Description,
		TgGroupLink: newGroup.TgGroupLink,
		TgChatId:    newGroup.TgChatID,
	}

	resGroup, err := g.client.UpdateGroup(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	return mapGroupGrpcToDTO(resGroup.GetGroup()), nil
}

func (g *GroupClient) RemoveGroup(groupId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := g.client.RemoveGroup(ctx, &groupGRPC.RemoveGroupRequest{Id: groupId})
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupClient) GetGroupById(groupId int64) (*groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resGroup, err := g.client.GetGroupById(ctx, &groupGRPC.GetGroupByIdRequest{Id: groupId})
	if err != nil {
		return nil, err
	}

	return mapGroupGrpcToDTO(resGroup), nil
}

func (g *GroupClient) GetGroups() ([]groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resGroups, err := g.client.GetGroups(ctx, &groupGRPC.GetGroupsRequest{})
	if err != nil {
		return nil, err
	}

	var groups []groupDto.GroupResponse
	for _, oneGroup := range resGroups.GetGroups() {
		groups = append(groups, *mapGroupGrpcToDTO(oneGroup))
	}

	return groups, nil
}

func (g *GroupClient) AddUsersToGroup(groupId int64, userIds []int64) ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resUsers, err := g.client.AddUsersToGroup(ctx, &groupGRPC.AddUsersToGroupRequest{
		GroupId: groupId,
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}

	return resUsers.GetUserIds(), nil
}

func (g *GroupClient) RemoveUserFromGroup(userId int64, groupId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := g.client.RemoveUserFromGroup(ctx, &groupGRPC.RemoveUserFromGroupRequest{
		UserId:  userId,
		GroupId: groupId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupClient) GetUserGroups(userId int64) ([]groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resGroups, err := g.client.GetUserGroups(ctx, &groupGRPC.GetUserGroupsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	var groups []groupDto.GroupResponse
	for _, oneGroup := range resGroups.GetGroups() {
		groups = append(groups, *mapGroupGrpcToDTO(oneGroup))
	}

	return groups, nil
}

func (g *GroupClient) GetGroupsByTutorId(tutorId int64) ([]groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resGroups, err := g.client.GetGroupsByTutorId(ctx, &groupGRPC.GetGroupsByTutorIdRequest{TutorId: tutorId})
	if err != nil {
		return nil, err
	}

	var groups []groupDto.GroupResponse
	for _, oneGroup := range resGroups.GetGroups() {
		groups = append(groups, *mapGroupGrpcToDTO(oneGroup))
	}

	return groups, nil
}

func (g *GroupClient) GetGroupUsers(groupId int64) ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resUsers, err := g.client.GetGroupUsers(ctx, &groupGRPC.GetGroupUsersRequest{GroupId: groupId})
	if err != nil {
		return nil, err
	}

	return resUsers.GetUserIds(), nil
}

func mapGroupGrpcToDTO(group *groupGRPC.GetGroupByIdResponse) *groupDto.GroupResponse {
	return &groupDto.GroupResponse{
		ID:          group.GetId(),
		Title:       group.GetTitle(),
		Description: group.GetDescription(),
		SubjectID:   group.GetSubjectId(),
		TutorID:     group.GetTutorId(),
		TgGroupLink: group.TgGroupLink,
		TgChatID:    group.TgChatId,
	}
}
