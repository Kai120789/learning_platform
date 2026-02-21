package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto"
	"time"
)

type UserClient struct {
	client userGRPC.UserClient
	logger *zap.Logger
}

/*
GetAllUsersWithData
ChangePassword
ChangeEmail
UpdateUserInfo
UpdateUserSettings
*/

func NewUserGrpcConnection(userGrpcUrl string, logger *zap.Logger) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		userGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("failed to create user grpc client", zap.Error(err))
		return nil, err
	}

	return conn, nil
}

func NewUserClient(connection *grpc.ClientConn, logger *zap.Logger) *UserClient {
	return &UserClient{
		client: userGRPC.NewUserClient(connection),
		logger: logger,
	}
}

func (u *UserClient) GetUserByEmail(email string) (*dto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserByEmail(ctx, &userGRPC.GetUserByEmailRequest{Email: email})
	if err != nil {
		u.logger.Error("failed to send get user by email grpc query", zap.Error(err))
		return nil, err
	}

	return &dto.GetUser{
		UserId:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}

func (u *UserClient) GetUserById(id int64) (*dto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserById(ctx, &userGRPC.GetUserByIdRequest{UserId: id})
	if err != nil {
		u.logger.Error("failed to send get user by id grpc query", zap.Error(err))
		return nil, err
	}

	return &dto.GetUser{
		UserId:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}

func (u *UserClient) GetUserData(id int64) (*dto.UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserData(ctx, &userGRPC.GetUserDataRequest{UserId: id})
	if err != nil {
		u.logger.Error("failed to send get user data request", zap.Error(err))
		return nil, err
	}

	return &dto.UserData{
		UserId: res.GetUserId(),
		Email:  res.GetEmail(),
		UserInfo: dto.UserInfo{
			UserId:   res.GetUserId(),
			Name:     res.GetUserInfo().GetName(),
			Surname:  res.GetUserInfo().GetSurname(),
			Lastname: getOptionalFieldString(res.GetUserInfo().GetLastname()),
			City:     getOptionalFieldString(res.GetUserInfo().GetCity()),
			About:    getOptionalFieldString(res.GetUserInfo().GetAbout()),
			Role:     protoRoleToString(res.GetUserInfo().GetRole()),
			Status:   protoStatusToString(res.GetUserInfo().GetStatus()),
			Class:    getOptionalFieldInt(res.GetUserInfo().GetClass()),
		},
		UserSettings: dto.UserSettings{},
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
		Role:         stringToProtoUserRole(newUser.Role),
		PasswordHash: newUser.Password,
	})
	if err != nil {
		u.logger.Error("failed to send create user grpc query", zap.Error(err))
		return nil, err
	}

	return &res.UserId, nil
}

func stringToProtoUserRole(role string) userGRPC.UserRole {
	switch role {
	case "TUTOR":
		return userGRPC.UserRole_TUTOR
	case "STUDENT":
		return userGRPC.UserRole_STUDENT
	case "ADMIN":
		return userGRPC.UserRole_ADMIN
	default:
		return userGRPC.UserRole_USER_ROLE_UNSPECIFIED
	}
}

func protoRoleToString(role userGRPC.UserRole) string {
	switch role {
	case userGRPC.UserRole_TUTOR:
		return "TUTOR"
	case userGRPC.UserRole_STUDENT:
		return "STUDENT"
	case userGRPC.UserRole_ADMIN:
		return "ADMIN"
	default:
		return "UNSPECIFIED"
	}
}

func stringToProtoStatus(status string) userGRPC.Status {
	switch status {
	case "ACTIVE":
		return userGRPC.Status_ACTIVE
	case "INACTIVE":
		return userGRPC.Status_INACTIVE
	case "BANNED":
		return userGRPC.Status_BANNED
	default:
		return userGRPC.Status_STATUS_UNSPECIFIED
	}
}

func protoStatusToString(status userGRPC.Status) string {
	switch status {
	case userGRPC.Status_ACTIVE:
		return "ACTIVE"
	case userGRPC.Status_INACTIVE:
		return "INACTIVE"
	case userGRPC.Status_BANNED:
		return "BANNED"
	default:
		return "UNSPECIFIED"
	}
}

func getOptionalFieldString(val string) *string { return &val }

func getOptionalFieldInt(val int64) *int64 { return &val }
