package grpc

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
)

type UserInfoService interface {
	UpdateUserInfo(userInfo dto.UserInfo) (*models.UserInfo, error)
}

func (g *UserGRPCServer) UpdateUserInfo(
	ctx context.Context,
	in *userGRPC.UpdateUserInfoRequest,
) (*userGRPC.UpdateUserInfoResponse, error) {
	userInfo := dto.UserInfo{
		UserId:   in.GetUserId(),
		Name:     in.GetName(),
		Surname:  in.GetSurname(),
		Lastname: stringToOptionalString(in.GetLastname()),
		City:     stringToOptionalString(in.GetCity()),
		About:    stringToOptionalString(in.GetAbout()),
		Class:    intToOptionalInt(in.GetClass()),
	}

	res, err := g.UserInfoService.UpdateUserInfo(userInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update user info")
	}
	return &userGRPC.UpdateUserInfoResponse{
		Name:     res.Name,
		Surname:  res.Name,
		Lastname: &res.Lastname.String,
		City:     &res.City.String,
		About:    &res.About.String,
		Role:     stringToProtoRole(res.Role),
		Status:   stringToProtoStatus(res.Status),
		Class:    &res.Class.Int64,
	}, nil
}

func stringToOptionalString(value string) *string {
	return &value
}

func intToOptionalInt(value int64) *int64 {
	return &value
}
