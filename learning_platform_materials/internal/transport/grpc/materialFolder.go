package grpc

import (
	"context"
	materialGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/material"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/materials/internal/dto"
	"learning-platform/materials/internal/models"
	"learning-platform/materials/internal/utils"
)

type MaterialFolderService interface {
	CreateFolder(folder dto.CreateFolder) (*models.MaterialFolder, error)
	MoveFolders(folderIDs []int64, newParentFolderID *int64) error
	RenameFolder(folderID int64, newFolderTitle string) error
	DeleteOneFolder(folderID int64) error
	DeleteFolders(folderIDs []int64) error
}

func (m *MaterialsGRPCServer) CreateFolder(
	ctx context.Context,
	in *materialGRPC.CreateFolderRequest,
) (*materialGRPC.CreateFolderResponse, error) {
	folder := dto.CreateFolder{
		Title:          in.GetTitle(),
		ParentFolderID: in.ParentFolderId,
		TutorID:        in.GetTutorId(),
	}

	res, err := m.service.MaterialFolderService.CreateFolder(folder)
	if err != nil {
		m.logger.Error("failed to create folder")
		return nil, status.Error(codes.Internal, "failed to create folder")
	}

	return &materialGRPC.CreateFolderResponse{
		Folder: &materialGRPC.Folder{
			Id:             res.ID,
			Title:          res.Title,
			ParentFolderId: utils.DBInt8ToOptional(res.ParentFolderID),
			TutorId:        res.TutorID,
		},
	}, nil
}

func (m *MaterialsGRPCServer) MoveFolders(
	ctx context.Context,
	in *materialGRPC.MoveFoldersRequest,
) (*materialGRPC.MoveFoldersResponse, error) {
	err := m.service.MaterialFolderService.MoveFolders(in.GetFolderIds(), in.ParentFolderId)
	if err != nil {
		m.logger.Error(
			"failed to move folders",
			zap.Int64s("folderIDs", in.GetFolderIds()),
			zap.Int64p("parentFolderID", in.ParentFolderId),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to move folders")
	}

	return &materialGRPC.MoveFoldersResponse{}, nil
}

func (m *MaterialsGRPCServer) RenameFolder(
	ctx context.Context,
	in *materialGRPC.RenameFolderRequest,
) (*materialGRPC.RenameFolderResponse, error) {
	err := m.service.MaterialFolderService.RenameFolder(in.GetFolderId(), in.GetTitle())
	if err != nil {
		m.logger.Error(
			"failed to rename folder",
			zap.Int64("folderID", in.GetFolderId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to rename folder")
	}

	return &materialGRPC.RenameFolderResponse{}, nil
}

func (m *MaterialsGRPCServer) DeleteOneFolder(
	ctx context.Context,
	in *materialGRPC.DeleteOneFolderRequest,
) (*materialGRPC.DeleteOneFolderResponse, error) {
	err := m.service.MaterialFolderService.DeleteOneFolder(in.GetFolderId())
	if err != nil {
		m.logger.Error(
			"failed to delete one folder",
			zap.Int64("folderID", in.GetFolderId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete one folder")
	}
	return &materialGRPC.DeleteOneFolderResponse{}, nil
}

func (m *MaterialsGRPCServer) DeleteFolders(
	ctx context.Context,
	in *materialGRPC.DeleteFoldersRequest,
) (*materialGRPC.DeleteFoldersResponse, error) {
	err := m.service.MaterialFolderService.DeleteFolders(in.GetFolderIds())
	if err != nil {
		m.logger.Error(
			"failed to delete folders",
			zap.Int64s("folderIDs", in.GetFolderIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete folders")
	}
	return &materialGRPC.DeleteFoldersResponse{}, nil
}
