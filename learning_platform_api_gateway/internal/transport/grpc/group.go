package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/groupDto"
	"time"
)

type GroupClient struct {
	client userGRPC.GroupClient
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
		client: userGRPC.NewGroupClient(connection),
	}
}

func (g *GroupClient) CreateGroup(group groupDto.CreateGroupRequest) (*groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reqBody := &userGRPC.CreateGroupRequest{
		Title:       group.Title,
		Description: group.Description,
		SubjectId:   group.SubjectId,
		TutorId:     group.TutorId,
		TgGroupLink: group.TgGroupLink,
		TgChatId:    group.TgChatId,
	}

	resGroup, err := g.client.CreateGroup(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	return &groupDto.GroupResponse{
		Id:          resGroup.GetId(),
		Title:       resGroup.GetTitle(),
		Description: resGroup.GetDescription(),
		SubjectId:   resGroup.GetSubjectId(),
		TutorId:     resGroup.GetTutorId(),
		TgGroupLink: getOptionalFieldString(resGroup.GetTgGroupLink()),
		TgChatId:    getOptionalFieldInt(resGroup.GetTgChatId()),
	}, nil
}

func (g *GroupClient) UpdateGroup(groupId int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reqBody := &userGRPC.UpdateGroupRequest{
		Id:          groupId,
		Title:       newGroup.Title,
		Description: newGroup.Description,
		TgGroupLink: newGroup.TgGroupLink,
		TgChatId:    newGroup.TgChatId,
	}

	resGroup, err := g.client.UpdateGroup(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	return &groupDto.GroupResponse{
		Id:          resGroup.GetId(),
		Title:       resGroup.GetTitle(),
		Description: resGroup.GetDescription(),
		SubjectId:   resGroup.GetSubjectId(),
		TutorId:     resGroup.GetTutorId(),
		TgGroupLink: getOptionalFieldString(resGroup.GetTgGroupLink()),
		TgChatId:    getOptionalFieldInt(resGroup.GetTgChatId()),
	}, nil
}

func (g *GroupClient) RemoveGroup(groupId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := g.client.RemoveGroup(ctx, &userGRPC.RemoveGroupRequest{Id: groupId})
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupClient) GetGroupById(groupId int64) (*groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resGroup, err := g.client.GetGroupById(ctx, &userGRPC.GetGroupByIdRequest{Id: groupId})
	if err != nil {
		return nil, err
	}

	return &groupDto.GroupResponse{
		Id:          resGroup.GetId(),
		Title:       resGroup.GetTitle(),
		Description: resGroup.GetDescription(),
		SubjectId:   resGroup.GetSubjectId(),
		TutorId:     resGroup.GetTutorId(),
		TgGroupLink: getOptionalFieldString(resGroup.GetTgGroupLink()),
		TgChatId:    getOptionalFieldInt(resGroup.GetTgChatId()),
	}, nil
}

func (g *GroupClient) GetGroups() ([]groupDto.GroupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resGroups, err := g.client.GetGroups(ctx, &userGRPC.GetGroupsRequest{})
	if err != nil {
		return nil, err
	}

	resGroupsAsDto := make([]groupDto.GroupResponse, len(resGroups.GetGroups()))

	for ind, resGroup := range resGroups.GetGroups() {
		resGroupsAsDto[ind] = groupDto.GroupResponse{
			Id:          resGroup.GetId(),
			Title:       resGroup.GetTitle(),
			Description: resGroup.GetDescription(),
			SubjectId:   resGroup.GetSubjectId(),
			TutorId:     resGroup.GetTutorId(),
			TgGroupLink: getOptionalFieldString(resGroup.GetTgGroupLink()),
			TgChatId:    getOptionalFieldInt(resGroup.GetTgChatId()),
		}
	}

	return resGroupsAsDto, nil
}

func (g *GroupClient) AddUsersToGroup(groupId int64, userIds []int64) ([]groupDto.ShortUserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resAddUsers, err := g.client.AddUsersToGroup(ctx, &userGRPC.AddUsersToGroupRequest{
		GroupId: groupId,
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}

	resUsers := make([]groupDto.ShortUserInfo, len(resAddUsers.GetUsers()))

	for ind, oneAddedUser := range resAddUsers.Users {
		resUsers[ind] = groupDto.ShortUserInfo{
			Id:    oneAddedUser.GetId(),
			Name:  oneAddedUser.GetName(),
			Email: oneAddedUser.GetEmail(),
		}
	}

	return resUsers, nil
}

func (g *GroupClient) RemoveUserFromGroup(userId int64, groupId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := g.client.RemoveUserFromGroup(ctx, &userGRPC.RemoveUserFromGroupRequest{
		UserId:  userId,
		GroupId: groupId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupClient) GetUserGroups(userId int64) ([]groupDto.GroupResponse, error) {
	ctx, cacnel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cacnel()

	resGroups, err := g.client.GetUserGroups(ctx, &userGRPC.GetUserGroupsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	resGroupsAsDto := make([]groupDto.GroupResponse, len(resGroups.GetGroups()))

	for ind, resGroup := range resGroups.GetGroups() {
		resGroupsAsDto[ind] = groupDto.GroupResponse{
			Id:          resGroup.GetId(),
			Title:       resGroup.GetTitle(),
			Description: resGroup.GetDescription(),
			SubjectId:   resGroup.GetSubjectId(),
			TutorId:     resGroup.GetTutorId(),
			TgGroupLink: getOptionalFieldString(resGroup.GetTgGroupLink()),
			TgChatId:    getOptionalFieldInt(resGroup.GetTgChatId()),
		}
	}

	return resGroupsAsDto, nil
}

func (g *GroupClient) GetGroupsByTutorId(tutorId int64) ([]groupDto.GroupResponse, error) {
	ctx, cacnel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cacnel()

	resGroups, err := g.client.GetGroupsByTutorId(ctx, &userGRPC.GetGroupsByTutorIdRequest{TutorId: tutorId})
	if err != nil {
		return nil, err
	}

	resGroupsAsDto := make([]groupDto.GroupResponse, len(resGroups.GetGroups()))

	for ind, resGroup := range resGroups.GetGroups() {
		resGroupsAsDto[ind] = groupDto.GroupResponse{
			Id:          resGroup.GetId(),
			Title:       resGroup.GetTitle(),
			Description: resGroup.GetDescription(),
			SubjectId:   resGroup.GetSubjectId(),
			TutorId:     resGroup.GetTutorId(),
			TgGroupLink: getOptionalFieldString(resGroup.GetTgGroupLink()),
			TgChatId:    getOptionalFieldInt(resGroup.GetTgChatId()),
		}
	}

	return resGroupsAsDto, nil
}

func (g *GroupClient) GetGroupUsers(groupId int64) ([]groupDto.ShortUserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resUsers, err := g.client.GetGroupUsers(ctx, &userGRPC.GetGroupUsersRequest{GroupId: groupId})
	if err != nil {
		return nil, err
	}

	resUsersAsDto := make([]groupDto.ShortUserInfo, len(resUsers.GetUsers()))
	for ind, user := range resUsers.GetUsers() {
		resUsersAsDto[ind] = groupDto.ShortUserInfo{
			Id:    user.GetId(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		}
	}

	return resUsersAsDto, nil
}
