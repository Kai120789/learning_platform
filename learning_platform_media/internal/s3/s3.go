package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"learning-platform/media/internal/dto"
	"sync"
	"time"
)

type Minio struct {
	client *minio.Client
	bucket string
}

func New(s3Url, rootUser, rootPassword, bucketName string) (*Minio, error) {
	client, err := minio.New(s3Url, &minio.Options{
		Creds: credentials.NewStaticV4(rootUser, rootPassword, ""),
	})
	if err != nil {
		return nil, err
	}

	return &Minio{
		client: client,
		bucket: bucketName,
	}, nil
}

func (m *Minio) CreateBucket(bucketName string) error {
	exists, err := m.client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err := m.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Minio) CreateOne(ctx context.Context, file dto.FileDataType) (string, error) {
	objectID := uuid.New().String()

	reader := bytes.NewReader(file.Data)

	_, err := m.client.PutObject(
		ctx,
		m.bucket,
		objectID,
		reader,
		int64(len(file.Data)),
		minio.PutObjectOptions{
			ContentType: file.FileMetadata.ContentType,
		},
	)
	if err != nil {
		return "", fmt.Errorf("create object: %w", err)
	}

	return objectID, nil
}

func (m *Minio) CreateMany(ctx context.Context, files []dto.FileDataType) ([]string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errChan := make(chan error, 1)
	resObjectIDs := make([]string, len(files))

	var wg sync.WaitGroup

	for ind, file := range files {
		wg.Add(1)
		go func(file dto.FileDataType, ind int) {
			defer wg.Done()
			objectID, err := m.CreateOne(ctx, file)
			if err != nil {
				cancel()
				select {
				case errChan <- err:
				default:
				}
				return
			}
			resObjectIDs[ind] = objectID
		}(file, ind)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		idsToDelete := make([]string, 0, len(resObjectIDs))
		for _, id := range resObjectIDs {
			if id != "" {
				idsToDelete = append(idsToDelete, id)
			}
		}
		_ = m.DeleteMany(context.Background(), idsToDelete)
		return nil, fmt.Errorf("create many objects: %w", err)
	default:
	}

	return resObjectIDs, nil
}

func (m *Minio) GetOne(ctx context.Context, objectID string) (string, error) {
	url, err := m.client.PresignedGetObject(
		ctx,
		m.bucket,
		objectID,
		time.Second*24*60*60,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("get object: %w", err)
	}

	return url.String(), nil
}

func (m *Minio) GetMany(ctx context.Context, objectIDs []string) ([]string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	resURLs := make([]string, len(objectIDs))
	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	for ind, objectID := range objectIDs {
		wg.Add(1)
		go func(objectID string, ind int) {
			defer wg.Done()
			url, err := m.GetOne(ctx, objectID)
			if err != nil {
				cancel()
				select {
				case errChan <- err:
				default:
				}
				return
			}

			resURLs[ind] = url
		}(objectID, ind)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return nil, fmt.Errorf("get many objects: %w", err)
	default:
	}

	return resURLs, nil
}

func (m *Minio) DeleteOne(ctx context.Context, objectID string) error {
	err := m.client.RemoveObject(
		ctx,
		m.bucket,
		objectID,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("delete object: %w", err)
	}

	return nil
}

func (m *Minio) DeleteMany(ctx context.Context, objectIDs []string) error {
	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	for _, objectID := range objectIDs {
		wg.Add(1)
		go func(objectID string) {
			defer wg.Done()
			err := m.DeleteOne(ctx, objectID)
			if err != nil {
				select {
				case errChan <- err:
				default:
				}
			}

			return
		}(objectID)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return fmt.Errorf("delete many: %w", err)
	default:
	}

	return nil
}
