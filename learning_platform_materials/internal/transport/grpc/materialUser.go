package grpc

import (
	"context"
	materialGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/material"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MaterialUserService interface {
	UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error
	UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error
}

func (m *MaterialsGRPCServer) UpdateUsersMaterialsAccess(
	ctx context.Context,
	in *materialGRPC.UpdateUsersMaterialsAccessRequest,
) (*materialGRPC.UpdateUsersMaterialsAccessResponse, error) {
	err := m.service.MaterialUserService.UpdateUsersMaterialsAccess(
		in.GetUserIds(),
		in.GetMaterialIds(),
	)
	if err != nil {
		m.logger.Error(
			"failed to update users materials access",
			zap.Int64s("userIDs", in.GetUserIds()),
			zap.Int64s("materialIDs", in.GetMaterialIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update users materials access")
	}

	return &materialGRPC.UpdateUsersMaterialsAccessResponse{}, nil
}

func (m *MaterialsGRPCServer) UpdateUsersFoldersAccess(
	ctx context.Context,
	in *materialGRPC.UpdateUsersFoldersAccessRequest,
) (*materialGRPC.UpdateUsersFoldersAccessResponse, error) {
	err := m.service.MaterialUserService.UpdateUsersFoldersAccess(
		in.GetUserIds(),
		in.GetFolderIds(),
	)
	if err != nil {
		m.logger.Error(
			"failed to update users folders access",
			zap.Int64s("userIDs", in.GetUserIds()),
			zap.Int64s("folderIDs", in.GetFolderIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update users folders access")
	}

	return &materialGRPC.UpdateUsersFoldersAccessResponse{}, nil
}
