package service

import "learning-platform/api-gateway/internal/dto/mediaDto"

type MediaService struct {
	client MediaClient
}

type MediaClient interface {
	CreateOne(file mediaDto.FileDataType) (string, error)
	CreateMany(files []mediaDto.FileDataType) ([]string, error)
	GetOne(objectID string) (string, error)
	GetMany(objectIDs []string) ([]string, error)
	DeleteOne(objectID string) error
	DeleteMany(objectIDs []string) error
}

func NewMediaService(client MediaClient) *MediaService {
	return &MediaService{
		client: client,
	}
}

func (m *MediaService) CreateOne(file mediaDto.FileDataType) (string, error) {
	res, err := m.client.CreateOne(file)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m *MediaService) CreateMany(files []mediaDto.FileDataType) ([]string, error) {
	res, err := m.client.CreateMany(files)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MediaService) GetOne(objectID string) (string, error) {
	res, err := m.client.GetOne(objectID)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m *MediaService) GetMany(objectIDs []string) ([]string, error) {
	res, err := m.client.GetMany(objectIDs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MediaService) DeleteOne(objectID string) error {
	err := m.client.DeleteOne(objectID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MediaService) DeleteMany(objectIDs []string) error {
	err := m.client.DeleteMany(objectIDs)
	if err != nil {
		return err
	}

	return nil
}
