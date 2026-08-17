package grpc

import (
	"context"
	materialGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/material"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/materialDto"
	"time"
)

type MaterialClient struct {
	client materialGRPC.MaterialClient
}

func NewMaterialGrpcConnection(materialGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		materialGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewMaterialClient(connection *grpc.ClientConn) *MaterialClient {
	return &MaterialClient{
		client: materialGRPC.NewMaterialClient(connection),
	}
}

func (m *MaterialClient) CreateFolder(folder materialDto.CreateFolder) (*materialDto.MaterialFolder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.client.CreateFolder(ctx, &materialGRPC.CreateFolderRequest{
		Title:          folder.Title,
		ParentFolderId: folder.ParentFolderID,
		TutorId:        folder.TutorID,
	})
	if err != nil {
		return nil, err
	}

	return &materialDto.MaterialFolder{
		ID:             res.GetFolder().GetId(),
		Title:          res.GetFolder().GetTitle(),
		ParentFolderID: res.GetFolder().ParentFolderId,
		TutorID:        res.GetFolder().GetTutorId(),
	}, nil
}

func (m *MaterialClient) MoveFolders(folderIDs []int64, parentFolderID *int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.MoveFolders(ctx, &materialGRPC.MoveFoldersRequest{
		FolderIds:      folderIDs,
		ParentFolderId: parentFolderID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) RenameFolder(folderID int64, newTitle string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.RenameFolder(ctx, &materialGRPC.RenameFolderRequest{
		FolderId: folderID,
		Title:    newTitle,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) DeleteOneFolder(folderID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.DeleteOneFolder(ctx, &materialGRPC.DeleteOneFolderRequest{
		FolderId: folderID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) DeleteFolders(folderIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.DeleteFolders(ctx, &materialGRPC.DeleteFoldersRequest{
		FolderIds: folderIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) GetStudentMaterials(studentID int64, folderID *int64) (*materialDto.FoldersAndMaterials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.client.GetStudentMaterials(ctx, &materialGRPC.GetStudentMaterialsRequest{
		StudentId: studentID,
		FolderId:  folderID,
	})
	if err != nil {
		return nil, err
	}

	var resFolders []materialDto.MaterialFolder
	for _, folder := range res.GetFolders() {
		resFolders = append(resFolders, materialDto.MaterialFolder{
			ID:             folder.GetId(),
			Title:          folder.GetTitle(),
			ParentFolderID: folder.ParentFolderId,
			TutorID:        folder.GetTutorId(),
		})
	}

	var resMaterials []materialDto.Material
	for _, material := range res.GetMaterials() {
		resMaterials = append(resMaterials, materialDto.Material{
			ID:            material.GetId(),
			Title:         material.GetTitle(),
			Size:          int64(material.GetSize()),
			FolderID:      material.FolderId,
			TutorID:       material.GetTutorId(),
			MimeType:      material.GetMimeType(),
			MediaObjectID: material.GetMediaObjectId(),
		})
	}

	return &materialDto.FoldersAndMaterials{
		Materials: resMaterials,
		Folders:   resFolders,
	}, nil
}

func (m *MaterialClient) GetTutorMaterials(tutorID int64, folderID *int64) (*materialDto.FoldersAndMaterials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.client.GetTutorMaterials(ctx, &materialGRPC.GetTutorMaterialsRequest{
		TutorId:  tutorID,
		FolderId: folderID,
	})
	if err != nil {
		return nil, err
	}

	var resFolders []materialDto.MaterialFolder
	for _, folder := range res.GetFolders() {
		resFolders = append(resFolders, materialDto.MaterialFolder{
			ID:             folder.GetId(),
			Title:          folder.GetTitle(),
			ParentFolderID: folder.ParentFolderId,
			TutorID:        folder.GetTutorId(),
		})
	}

	var resMaterials []materialDto.Material
	for _, material := range res.GetMaterials() {
		resMaterials = append(resMaterials, materialDto.Material{
			ID:            material.GetId(),
			Title:         material.GetTitle(),
			Size:          int64(material.GetSize()),
			FolderID:      material.FolderId,
			TutorID:       material.GetTutorId(),
			MimeType:      material.GetMimeType(),
			MediaObjectID: material.GetMediaObjectId(),
		})
	}

	return &materialDto.FoldersAndMaterials{
		Materials: resMaterials,
		Folders:   resFolders,
	}, nil
}

func (m *MaterialClient) CreateMaterials(materials []materialDto.CreateMaterial) ([]materialDto.Material, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var requestMaterials []*materialGRPC.NewFolderMaterial
	for _, material := range materials {
		requestMaterials = append(requestMaterials, &materialGRPC.NewFolderMaterial{
			Title:         material.Title,
			Size:          uint64(material.Size),
			FolderId:      material.FolderID,
			TutorId:       material.TutorID,
			MimeType:      material.MimeType,
			MediaObjectId: material.MediaObjectID,
		})
	}

	res, err := m.client.CreateMaterials(ctx, &materialGRPC.CreateMaterialsRequest{
		Materials: requestMaterials,
	})
	if err != nil {
		return nil, err
	}

	var resMaterials []materialDto.Material
	for _, material := range res.GetMaterials() {
		resMaterials = append(resMaterials, materialDto.Material{
			ID:            material.GetId(),
			Title:         material.GetTitle(),
			Size:          int64(material.GetSize()),
			FolderID:      material.FolderId,
			TutorID:       material.GetTutorId(),
			MimeType:      material.GetMimeType(),
			MediaObjectID: material.GetMediaObjectId(),
		})
	}

	return resMaterials, nil
}

func (m *MaterialClient) MoveMaterials(materialIDs []int64, folderID *int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.MoveMaterials(ctx, &materialGRPC.MoveMaterialsRequest{
		FolderId:    folderID,
		MaterialIds: materialIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) RenameMaterial(materialID int64, newTitle string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.RenameMaterial(ctx, &materialGRPC.RenameMaterialRequest{
		MaterialId: materialID,
		Title:      newTitle,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) DeleteOneMaterial(materialID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.DeleteOneMaterial(ctx, &materialGRPC.DeleteOneMaterialRequest{
		MaterialId: materialID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) DeleteMaterials(materialIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.DeleteMaterials(ctx, &materialGRPC.DeleteMaterialsRequest{
		MaterialIds: materialIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.UpdateUsersMaterialsAccess(ctx, &materialGRPC.UpdateUsersMaterialsAccessRequest{
		MaterialIds: materialIDs,
		UserIds:     userIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialClient) UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.UpdateUsersFoldersAccess(ctx, &materialGRPC.UpdateUsersFoldersAccessRequest{
		FolderIds: folderIDs,
		UserIds:   userIDs,
	})
	if err != nil {
		return err
	}

	return nil
}
