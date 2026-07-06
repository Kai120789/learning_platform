package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/userDto"
	"time"
)

type UserClient struct {
	client userGRPC.UserClient
}

/*
GetAllUsersWithData
ChangePassword
ChangeEmail
UpdateUserInfo
UpdateUserSettings
*/

func NewUserGrpcConnection(userGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		userGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewUserClient(connection *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: userGRPC.NewUserClient(connection),
	}
}

func (u *UserClient) GetUserByEmail(email string) (*userDto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserByEmail(ctx, &userGRPC.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	return &userDto.GetUser{
		UserID:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}

func (u *UserClient) GetUserById(id int64) (*userDto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserById(ctx, &userGRPC.GetUserByIdRequest{UserId: id})
	if err != nil {
		return nil, err
	}

	return &userDto.GetUser{
		UserID:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}

func (u *UserClient) GetUserData(id int64) (*userDto.UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserData(ctx, &userGRPC.GetUserDataRequest{UserId: id})
	if err != nil {
		return nil, err
	}

	return &userDto.UserData{
		UserID: res.GetUserId(),
		Email:  res.GetEmail(),
		Role:   protoRoleToEnum(res.GetRole()),
		Status: protoStatusToEnum(res.GetStatus()),
		UserInfo: userDto.UserInfo{
			UserID:   res.GetUserId(),
			Name:     res.GetUserInfo().GetName(),
			Surname:  res.GetUserInfo().GetSurname(),
			Lastname: res.GetUserInfo().Lastname,
			City:     res.GetUserInfo().City,
			About:    res.GetUserInfo().About,
		},
		UserSettings: userDto.UserSettings{
			UserID:                 res.GetUserId(),
			Is2FaEnabled:           res.GetUserSettings().GetIs_2FaEnabled(),
			IsNotificationsEnabled: res.GetUserSettings().GetIsNotificationsEnabled(),
		},
	}, nil
}

func (u *UserClient) CreateUser(newUser dto.RegisterRequest) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.CreateUser(ctx, &userGRPC.CreateUserRequest{
		Email:        newUser.Email,
		Name:         newUser.Name,
		Surname:      newUser.Surname,
		LastName:     &newUser.LastName,
		Role:         enumToProtoUserRole(newUser.Role),
		PasswordHash: newUser.Password,
	})
	if err != nil {
		return nil, err
	}

	resUserId := res.GetUserId()

	return &resUserId, nil
}

func enumToProtoUserRole(role enum.UserRole) userGRPC.UserRole {
	switch role {
	case enum.RoleTutor:
		return userGRPC.UserRole_TUTOR
	case enum.RoleStudent:
		return userGRPC.UserRole_STUDENT
	case enum.RoleAdmin:
		return userGRPC.UserRole_ADMIN
	default:
		return userGRPC.UserRole_USER_ROLE_UNSPECIFIED
	}
}

func protoRoleToEnum(role userGRPC.UserRole) enum.UserRole {
	switch role {
	case userGRPC.UserRole_TUTOR:
		return enum.RoleTutor
	case userGRPC.UserRole_STUDENT:
		return enum.RoleStudent
	case userGRPC.UserRole_ADMIN:
		return enum.RoleAdmin
	default:
		return ""
	}
}

func enumToProtoStatus(status enum.UserStatus) userGRPC.Status {
	switch status {
	case enum.StatusActive:
		return userGRPC.Status_ACTIVE
	case enum.StatusInactive:
		return userGRPC.Status_INACTIVE
	case enum.StatusBanned:
		return userGRPC.Status_BANNED
	default:
		return userGRPC.Status_STATUS_UNSPECIFIED
	}
}

func protoStatusToEnum(status userGRPC.Status) enum.UserStatus {
	switch status {
	case userGRPC.Status_ACTIVE:
		return enum.StatusActive
	case userGRPC.Status_INACTIVE:
		return enum.StatusInactive
	case userGRPC.Status_BANNED:
		return enum.StatusBanned
	default:
		return ""
	}
}
