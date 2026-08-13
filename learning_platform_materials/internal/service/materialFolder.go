package service

import (
	"learning-platform/materials/internal/dto"
	"learning-platform/materials/internal/models"
)

type MaterialFolderService struct {
	storage MaterialFolderStorage
}

type MaterialFolderStorage interface {
	CreateFolder(folder dto.CreateFolder) (*models.MaterialFolder, error)
	MoveFolders(folderIDs []int64, newParentFolderID *int64) error
	RenameFolder(folderID int64, newFolderTitle string) error
	DeleteOneFolder(folderID int64) error
	DeleteFolders(folderIDs []int64) error
	GetStudentFolders(studentID int64, folderID *int64) ([]models.MaterialFolder, error)
	GetTutorFolders(tutorID int64, folderID *int64) ([]models.MaterialFolder, error)
}

func NewMaterialFolderService(storage MaterialFolderStorage) *MaterialFolderService {
	return &MaterialFolderService{
		storage: storage,
	}
}

func (mf *MaterialFolderService) CreateFolder(folder dto.CreateFolder) (*models.MaterialFolder, error) {
	res, err := mf.storage.CreateFolder(folder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (mf *MaterialFolderService) MoveFolders(folderIDs []int64, newParentFolderID *int64) error {
	err := mf.storage.MoveFolders(folderIDs, newParentFolderID)
	if err != nil {
		return err
	}
	return nil
}

func (mf *MaterialFolderService) RenameFolder(folderID int64, newFolderTitle string) error {
	err := mf.storage.RenameFolder(folderID, newFolderTitle)
	if err != nil {
		return err
	}
	return nil
}

func (mf *MaterialFolderService) DeleteOneFolder(folderID int64) error {
	err := mf.storage.DeleteOneFolder(folderID)
	if err != nil {
		return err
	}
	return nil
}

func (mf *MaterialFolderService) DeleteFolders(folderIDs []int64) error {
	err := mf.storage.DeleteFolders(folderIDs)
	if err != nil {
		return err
	}
	return nil
}

func (mf *MaterialFolderService) GetStudentFolders(studentID int64, folderID *int64) ([]models.MaterialFolder, error) {
	res, err := mf.storage.GetStudentFolders(studentID, folderID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (mf *MaterialFolderService) GetTutorFolders(tutorID int64, folderID *int64) ([]models.MaterialFolder, error) {
	res, err := mf.storage.GetTutorFolders(tutorID, folderID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
