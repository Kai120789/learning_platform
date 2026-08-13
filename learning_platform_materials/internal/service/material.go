package service

import (
	"learning-platform/materials/internal/dto"
	"learning-platform/materials/internal/models"
)

type MaterialService struct {
	storage MaterialStorage
	folder  GetMaterialFolderService
}

type MaterialStorage interface {
	GetStudentMaterials(studentID int64, folderID *int64) ([]models.Material, error)
	GetTutorMaterials(tutorID int64, folderID *int64) ([]models.Material, error)
	CreateMaterials(materials []dto.CreateMaterial) ([]models.Material, error)
	MoveMaterials(materialIDs []int64, folderID *int64) error
	RenameMaterial(materialID int64, newTitle string) error
	DeleteOneMaterial(materialID int64) error
	DeleteMaterials(materialIDs []int64) error
}

type GetMaterialFolderService interface {
	GetStudentFolders(studentID int64, folderID *int64) ([]models.MaterialFolder, error)
	GetTutorFolders(tutorID int64, folderID *int64) ([]models.MaterialFolder, error)
}

func NewMaterialService(storage MaterialStorage, folder GetMaterialFolderService) *MaterialService {
	return &MaterialService{
		storage: storage,
		folder:  folder,
	}
}

func (m *MaterialService) GetStudentMaterials(studentID int64, folderID *int64) (*dto.FoldersAndMaterials, error) {
	materials, err := m.storage.GetStudentMaterials(studentID, folderID)
	if err != nil {
		return nil, err
	}

	folders, err := m.folder.GetStudentFolders(studentID, folderID)
	if err != nil {
		return nil, err
	}

	return &dto.FoldersAndMaterials{
		Folders:   folders,
		Materials: materials,
	}, nil
}

func (m *MaterialService) GetTutorMaterials(tutorID int64, folderID *int64) (*dto.FoldersAndMaterials, error) {
	materials, err := m.storage.GetTutorMaterials(tutorID, folderID)
	if err != nil {
		return nil, err
	}

	folders, err := m.folder.GetTutorFolders(tutorID, folderID)
	if err != nil {
		return nil, err
	}

	return &dto.FoldersAndMaterials{
		Folders:   folders,
		Materials: materials,
	}, nil
}

func (m *MaterialService) CreateMaterials(materials []dto.CreateMaterial) ([]models.Material, error) {
	res, err := m.storage.CreateMaterials(materials)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MaterialService) MoveMaterials(materialIDs []int64, folderID *int64) error {
	err := m.storage.MoveMaterials(materialIDs, folderID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MaterialService) RenameMaterial(materialID int64, newTitle string) error {
	err := m.storage.RenameMaterial(materialID, newTitle)
	if err != nil {
		return err
	}
	return nil
}

func (m *MaterialService) DeleteOneMaterial(materialID int64) error {
	err := m.storage.DeleteOneMaterial(materialID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MaterialService) DeleteMaterials(materialIDs []int64) error {
	err := m.storage.DeleteMaterials(materialIDs)
	if err != nil {
		return err
	}
	return nil
}
