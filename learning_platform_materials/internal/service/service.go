package service

type Service struct {
	MaterialService       *MaterialService
	MaterialFolderService *MaterialFolderService
	MaterialUserService   *MaterialUserService
}

type Storage struct {
	MaterialStorage       MaterialStorage
	MaterialFolderStorage MaterialFolderStorage
	MaterialUserStorage   MaterialUserStorage
}

func New(storage *Storage) *Service {
	folderService := NewMaterialFolderService(storage.MaterialFolderStorage)

	return &Service{
		MaterialService:       NewMaterialService(storage.MaterialStorage, folderService),
		MaterialFolderService: folderService,
		MaterialUserService:   NewMaterialUserService(storage.MaterialUserStorage),
	}
}
