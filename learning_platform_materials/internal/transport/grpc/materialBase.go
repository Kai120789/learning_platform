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

type MaterialBaseService interface {
	GetStudentMaterials(studentID int64, folderID *int64) (*dto.FoldersAndMaterials, error)
	GetTutorMaterials(tutorID int64, folderID *int64) (*dto.FoldersAndMaterials, error)
	CreateMaterials(materials []dto.CreateMaterial) ([]models.Material, error)
	MoveMaterials(materialIDs []int64, folderID *int64) error
	RenameMaterial(materialID int64, newTitle string) error
	DeleteOneMaterial(materialID int64) error
	DeleteMaterials(materialIDs []int64) error
}

func (m *MaterialsGRPCServer) GetStudentMaterials(
	ctx context.Context,
	in *materialGRPC.GetStudentMaterialsRequest,
) (*materialGRPC.GetStudentMaterialsResponse, error) {
	res, err := m.service.MaterialBaseService.GetStudentMaterials(in.GetStudentId(), in.FolderId)
	if err != nil {
		m.logger.Error(
			"failed to get student materials",
			zap.Int64("studentID", in.GetStudentId()),
			zap.Int64p("folderID", in.FolderId),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get student materials")
	}

	var resMaterials []*materialGRPC.FolderMaterial
	for _, oneMaterial := range res.Materials {
		resMaterials = append(
			resMaterials,
			&materialGRPC.FolderMaterial{
				Id:            oneMaterial.ID,
				Title:         oneMaterial.Title,
				Size:          uint64(oneMaterial.Size),
				FolderId:      utils.DBInt8ToOptional(oneMaterial.FolderID),
				TutorId:       oneMaterial.TutorID,
				MimeType:      oneMaterial.MimeType,
				MediaObjectId: utils.DBUUIDToString(oneMaterial.MediaObjectID),
			},
		)
	}

	var resFolders []*materialGRPC.Folder
	for _, oneFolder := range res.Folders {
		resFolders = append(
			resFolders,
			&materialGRPC.Folder{
				Id:             oneFolder.ID,
				Title:          oneFolder.Title,
				ParentFolderId: utils.DBInt8ToOptional(oneFolder.ParentFolderID),
				TutorId:        oneFolder.TutorID,
			},
		)
	}

	return &materialGRPC.GetStudentMaterialsResponse{
		Materials: resMaterials,
		Folders:   resFolders,
	}, nil
}

func (m *MaterialsGRPCServer) GetTutorMaterials(
	ctx context.Context,
	in *materialGRPC.GetTutorMaterialsRequest,
) (*materialGRPC.GetTutorMaterialsResponse, error) {
	res, err := m.service.MaterialBaseService.GetTutorMaterials(in.GetTutorId(), in.FolderId)
	if err != nil {
		m.logger.Error(
			"failed to get tutor materials",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64p("folderID", in.FolderId),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get tutor materials")
	}

	var resMaterials []*materialGRPC.FolderMaterial
	for _, oneMaterial := range res.Materials {
		resMaterials = append(
			resMaterials,
			&materialGRPC.FolderMaterial{
				Id:            oneMaterial.ID,
				Title:         oneMaterial.Title,
				Size:          uint64(oneMaterial.Size),
				FolderId:      utils.DBInt8ToOptional(oneMaterial.FolderID),
				TutorId:       oneMaterial.TutorID,
				MimeType:      oneMaterial.MimeType,
				MediaObjectId: utils.DBUUIDToString(oneMaterial.MediaObjectID),
			},
		)
	}

	var resFolders []*materialGRPC.Folder
	for _, oneFolder := range res.Folders {
		resFolders = append(
			resFolders,
			&materialGRPC.Folder{
				Id:             oneFolder.ID,
				Title:          oneFolder.Title,
				ParentFolderId: utils.DBInt8ToOptional(oneFolder.ParentFolderID),
				TutorId:        oneFolder.TutorID,
			},
		)
	}

	return &materialGRPC.GetTutorMaterialsResponse{
		Materials: resMaterials,
		Folders:   resFolders,
	}, nil
}

func (m *MaterialsGRPCServer) CreateMaterials(
	ctx context.Context,
	in *materialGRPC.CreateMaterialsRequest,
) (*materialGRPC.CreateMaterialsResponse, error) {
	var materials []dto.CreateMaterial
	for _, oneMaterial := range in.GetMaterials() {
		materials = append(materials,
			dto.CreateMaterial{
				Title:         oneMaterial.GetTitle(),
				Size:          int64(oneMaterial.GetSize()),
				FolderID:      oneMaterial.FolderId,
				TutorID:       oneMaterial.GetTutorId(),
				MimeType:      oneMaterial.GetMimeType(),
				MediaObjectID: oneMaterial.GetMediaObjectId(),
			},
		)
	}

	res, err := m.service.MaterialBaseService.CreateMaterials(materials)
	if err != nil {
		m.logger.Error("failed to create materials", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to create materials")
	}

	var resMaterials []*materialGRPC.FolderMaterial
	for _, oneMaterial := range res {
		resMaterials = append(
			resMaterials,
			&materialGRPC.FolderMaterial{
				Id:            oneMaterial.ID,
				Title:         oneMaterial.Title,
				Size:          uint64(oneMaterial.Size),
				FolderId:      utils.DBInt8ToOptional(oneMaterial.FolderID),
				TutorId:       oneMaterial.TutorID,
				MimeType:      oneMaterial.MimeType,
				MediaObjectId: utils.DBUUIDToString(oneMaterial.MediaObjectID),
			},
		)
	}

	return &materialGRPC.CreateMaterialsResponse{
		Materials: resMaterials,
	}, nil
}

func (m *MaterialsGRPCServer) MoveMaterials(
	ctx context.Context,
	in *materialGRPC.MoveMaterialsRequest,
) (*materialGRPC.MoveMaterialsResponse, error) {
	err := m.service.MaterialBaseService.MoveMaterials(in.GetMaterialIds(), in.FolderId)
	if err != nil {
		m.logger.Error(
			"failed to move materials",
			zap.Int64s("materialIDs", in.GetMaterialIds()),
			zap.Int64p("folderID", in.FolderId),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to move materials")
	}
	return &materialGRPC.MoveMaterialsResponse{}, nil
}

func (m *MaterialsGRPCServer) RenameMaterial(
	ctx context.Context,
	in *materialGRPC.RenameMaterialRequest,
) (*materialGRPC.RenameMaterialResponse, error) {
	err := m.service.MaterialBaseService.RenameMaterial(in.GetMaterialId(), in.GetTitle())
	if err != nil {
		m.logger.Error(
			"failed to rename material",
			zap.Int64("materialID", in.GetMaterialId()),
			zap.String("newTitle", in.GetTitle()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to rename material")
	}

	return &materialGRPC.RenameMaterialResponse{}, nil
}

func (m *MaterialsGRPCServer) DeleteOneMaterial(
	ctx context.Context,
	in *materialGRPC.DeleteOneMaterialRequest,
) (*materialGRPC.DeleteOneMaterialResponse, error) {
	err := m.service.MaterialBaseService.DeleteOneMaterial(in.GetMaterialId())
	if err != nil {
		m.logger.Error(
			"failed to delete one material",
			zap.Int64("materialID", in.GetMaterialId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete one material")
	}

	return &materialGRPC.DeleteOneMaterialResponse{}, nil
}

func (m *MaterialsGRPCServer) DeleteMaterials(
	ctx context.Context,
	in *materialGRPC.DeleteMaterialsRequest,
) (*materialGRPC.DeleteMaterialsResponse, error) {
	err := m.service.MaterialBaseService.DeleteMaterials(in.GetMaterialIds())
	if err != nil {
		m.logger.Error(
			"failed to delete materials",
			zap.Int64s("materialIDs", in.GetMaterialIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete materials")
	}

	return &materialGRPC.DeleteMaterialsResponse{}, nil
}
