package service

import (
	"context"
	"learning-platform/media/internal/dto"
)

type Service struct {
	minio MinioClient
}

type MinioClient interface {
	CreateOne(ctx context.Context, file dto.FileDataType) (string, error)
	CreateMany(ctx context.Context, files []dto.FileDataType) ([]string, error)
	GetOne(ctx context.Context, objectID string) (string, error)
	GetMany(ctx context.Context, objectIDs []string) ([]string, error)
	DeleteOne(ctx context.Context, objectID string) error
	DeleteMany(ctx context.Context, objectIDs []string) error
}

func New(minio MinioClient) *Service {
	return &Service{
		minio: minio,
	}
}

func (s *Service) CreateOne(ctx context.Context, file dto.FileDataType) (string, error) {
	res, err := s.minio.CreateOne(ctx, file)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *Service) CreateMany(ctx context.Context, files []dto.FileDataType) ([]string, error) {
	res, err := s.minio.CreateMany(ctx, files)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GetOne(ctx context.Context, objectID string) (string, error) {
	res, err := s.minio.GetOne(ctx, objectID)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *Service) GetMany(ctx context.Context, objectIDs []string) ([]string, error) {
	res, err := s.minio.GetMany(ctx, objectIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) DeleteOne(ctx context.Context, objectID string) error {
	err := s.minio.DeleteOne(ctx, objectID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteMany(ctx context.Context, objectIDs []string) error {
	err := s.minio.DeleteMany(ctx, objectIDs)
	if err != nil {
		return err
	}

	return nil
}
