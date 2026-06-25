package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
)

type UserInfoService interface {
	UpdateUserInfo(userInfo dto.UserInfo) (*models.UserInfo, error)
}

func (g *UserGRPCServer) UpdateUserInfo(
	ctx context.Context,
	in *userGRPC.UpdateUserInfoRequest,
) (*userGRPC.UpdateUserInfoResponse, error) {
	userInfo := dto.UserInfo{
		UserID:   in.GetUserId(),
		Name:     in.GetName(),
		Surname:  in.GetSurname(),
		Lastname: stringToOptionalString(in.GetLastname()),
		City:     stringToOptionalString(in.GetCity()),
		About:    stringToOptionalString(in.GetAbout()),
	}

	res, err := g.UserInfoService.UpdateUserInfo(userInfo)
	if err != nil {
		g.logger.Error(
			"failed to update user info",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user info")
	}
	return &userGRPC.UpdateUserInfoResponse{
		Name:     res.Name,
		Surname:  res.Name,
		Lastname: &res.Lastname.String,
		City:     &res.City.String,
		About:    &res.About.String,
	}, nil
}

func stringToOptionalString(value string) *string {
	return &value
}
