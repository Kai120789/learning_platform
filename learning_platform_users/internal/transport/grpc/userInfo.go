package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/utils"
)

type UserInfoService interface {
	UpdateUserInfo(userInfo dto.UserInfoRequest) (*models.UserInfo, error)
	UpdateUserAvatar(userID int64, avatar string) error
	GetUsersShortInfo(userIDs []int64) ([]dto.UserShortInfo, error)
	UpdateUserTgUsername(userID int64, tgUsername string) error
}

func (u *UserGRPCServer) UpdateUserInfo(
	ctx context.Context,
	in *userGRPC.UpdateUserInfoRequest,
) (*userGRPC.UpdateUserInfoResponse, error) {
	userInfo := dto.UserInfoRequest{
		UserID:     in.GetUserId(),
		Name:       in.GetName(),
		Surname:    in.GetSurname(),
		Patronymic: in.Patronymic,
		City:       in.City,
		About:      in.About,
		Gender:     protoToEnumGender(in.GetGender()),
		BirthDate:  mapDate(in.BirthDate),
	}

	res, err := u.UserInfoService.UpdateUserInfo(userInfo)
	if err != nil {
		u.logger.Error(
			"failed to update user info",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user info")
	}
	return &userGRPC.UpdateUserInfoResponse{
		Name:       res.Name,
		Surname:    res.Name,
		Patronymic: utils.DBStringToOptional(res.Patronymic),
		TgUsername: utils.DBStringToOptional(res.TgUsername),
		City:       utils.DBStringToOptional(res.City),
		About:      utils.DBStringToOptional(res.About),
		Avatar:     utils.DBStringToOptional(res.Avatar),
		Gender:     enumToProtoGender(res.Gender),
		BirthDate: &userGRPC.Date{
			Year:  int32(res.BirthDate.Time.Year()),
			Month: int32(res.BirthDate.Time.Month()),
			Day:   int32(res.BirthDate.Time.Day()),
		},
	}, nil
}

func (u *UserGRPCServer) UpdateUserAvatar(
	ctx context.Context,
	in *userGRPC.UpdateUserAvatarRequest,
) (*userGRPC.UpdateUserAvatarResponse, error) {
	err := u.UserInfoService.UpdateUserAvatar(in.GetUserId(), in.GetAvatarUrl())
	if err != nil {
		u.logger.Error(
			"failed to update user avatar",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user avatar")
	}
	return nil, nil
}

func (u *UserGRPCServer) GetUsersShortInfo(
	ctx context.Context,
	in *userGRPC.GetUsersShortInfoRequest,
) (*userGRPC.GetUsersShortInfoResponse, error) {
	resUsers, err := u.UserInfoService.GetUsersShortInfo(in.GetUserIds())
	if err != nil {
		u.logger.Error(
			"failed to get users short info",
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get users short info")
	}

	var users []*userGRPC.UserShortInfo
	for _, oneUser := range resUsers {
		users = append(users, &userGRPC.UserShortInfo{
			Id:         oneUser.ID,
			Name:       oneUser.Name,
			Surname:    oneUser.Surname,
			Patronymic: oneUser.Patronymic,
			TgUsername: oneUser.TgUsername,
		})
	}

	return &userGRPC.GetUsersShortInfoResponse{
		Users: users,
	}, nil
}

func (u *UserGRPCServer) UpdateUserTgUsername(
	ctx context.Context,
	in *userGRPC.UpdateUserTgUsernameRequest,
) (*userGRPC.UpdateUserTgUsernameResponse, error) {
	err := u.UserInfoService.UpdateUserTgUsername(in.GetUserId(), in.GetTgUsername())
	if err != nil {
		u.logger.Error(
			"failed to update user tg username",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update user tg username")
	}
	return nil, nil
}
