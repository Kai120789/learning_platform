package service

import (
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/materialDto"
	"learning-platform/api-gateway/internal/dto/mediaDto"
)

type MaterialService struct {
	client MaterialClient
	media  MaterialMediaService
}

type MaterialClient interface {
	CreateFolder(folder materialDto.CreateFolder) (*materialDto.MaterialFolder, error)
	MoveFolders(folderIDs []int64, parentFolderID *int64) error
	RenameFolder(folderID int64, newTitle string) error
	DeleteOneFolder(folderID int64) error
	DeleteFolders(folderIDs []int64) error
	GetStudentMaterials(studentID int64, folderID *int64) (*materialDto.FoldersAndMaterials, error)
	GetTutorMaterials(tutorID int64, folderID *int64) (*materialDto.FoldersAndMaterials, error)
	CreateMaterials(materials []materialDto.CreateMaterial) ([]materialDto.Material, error)
	MoveMaterials(materialIDs []int64, folderID *int64) error
	RenameMaterial(materialID int64, newTitle string) error
	DeleteOneMaterial(materialID int64) error
	DeleteMaterials(materialIDs []int64) error
	UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error
	UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error
}

type MaterialMediaService interface {
	CreateOne(file mediaDto.FileDataType) (string, error)
	CreateMany(files []mediaDto.FileDataType) ([]string, error)
	GetOne(objectID string) (string, error)
	GetMany(objectIDs []string) ([]string, error)
	DeleteOne(objectID string) error
	DeleteMany(objectIDs []string) error
}

func NewMaterialService(client MaterialClient, media MaterialMediaService) *MaterialService {
	return &MaterialService{
		client: client,
		media:  media,
	}
}

func (m *MaterialService) CreateFolder(folder materialDto.CreateFolder) (*materialDto.MaterialFolder, error) {
	res, err := m.client.CreateFolder(folder)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MaterialService) MoveFolders(folderIDs []int64, parentFolderID *int64) error {
	err := m.client.MoveFolders(folderIDs, parentFolderID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) RenameFolder(folderID int64, newTitle string) error {
	err := m.client.RenameFolder(folderID, newTitle)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) DeleteOneFolder(folderID int64) error {
	err := m.client.DeleteOneFolder(folderID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) DeleteFolders(folderIDs []int64) error {
	err := m.client.DeleteFolders(folderIDs)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) GetMaterials(
	userID int64,
	folderID *int64,
	role enum.UserRole,
) (*materialDto.FoldersAndMaterials, error) {
	var res *materialDto.FoldersAndMaterials
	var err error
	if role == enum.RoleStudent {
		res, err = m.client.GetStudentMaterials(userID, folderID)
	} else {
		res, err = m.client.GetTutorMaterials(userID, folderID)
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MaterialService) CreateMaterials(tutorID int64, folderID *int64, files []mediaDto.FileDataType) ([]materialDto.Material, error) {
	objectIDs, err := m.media.CreateMany(files)
	if err != nil {
		return nil, err
	}

	materials := make([]materialDto.CreateMaterial, len(files))

	for i, file := range files {
		materials[i] = materialDto.CreateMaterial{
			Title:         file.FileMetadata.FileName,
			Size:          int64(file.FileMetadata.Size),
			FolderID:      folderID,
			TutorID:       tutorID,
			MimeType:      file.FileMetadata.ContentType,
			MediaObjectID: objectIDs[i],
		}
	}

	res, err := m.client.CreateMaterials(materials)
	if err != nil {
		_ = m.media.DeleteMany(objectIDs)
		return nil, err
	}

	return res, nil
}

func (m *MaterialService) MoveMaterials(materialIDs []int64, folderID *int64) error {
	err := m.client.MoveMaterials(materialIDs, folderID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) RenameMaterial(materialID int64, newTitle string) error {
	err := m.client.RenameMaterial(materialID, newTitle)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) DeleteOneMaterial(materialID int64) error {
	err := m.client.DeleteOneMaterial(materialID)
	if err != nil {
		return err
	}

	//TODO: добавить удаление из сервиса медиа

	return nil
}

func (m *MaterialService) DeleteMaterials(materialIDs []int64) error {
	err := m.client.DeleteMaterials(materialIDs)
	if err != nil {
		return err
	}

	//TODO: добавить удаление из сервиса медиа

	return nil
}

func (m *MaterialService) UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error {
	err := m.client.UpdateUsersMaterialsAccess(userIDs, materialIDs)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialService) UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error {
	err := m.client.UpdateUsersFoldersAccess(userIDs, folderIDs)
	if err != nil {
		return err
	}

	return nil
}
