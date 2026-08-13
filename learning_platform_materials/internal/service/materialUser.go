package service

type MaterialUserService struct {
	storage MaterialUserStorage
}

type MaterialUserStorage interface {
	UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error
	UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error
}

func NewMaterialUserService(storage MaterialUserStorage) *MaterialUserService {
	return &MaterialUserService{
		storage: storage,
	}
}

func (mu *MaterialUserService) UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error {
	err := mu.storage.UpdateUsersMaterialsAccess(userIDs, materialIDs)
	if err != nil {
		return err
	}
	return nil
}

func (mu *MaterialUserService) UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error {
	err := mu.storage.UpdateUsersFoldersAccess(userIDs, folderIDs)
	if err != nil {
		return err
	}
	return nil
}
