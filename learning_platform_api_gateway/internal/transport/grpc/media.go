package grpc

import (
	"context"
	mediaGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/media"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/mediaDto"
	"time"
)

type MediaClient struct {
	client mediaGRPC.MediaClient
}

func NewMediaGrpcConnection(mediaGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		mediaGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewMediaClient(connection *grpc.ClientConn) *MediaClient {
	return &MediaClient{
		client: mediaGRPC.NewMediaClient(connection),
	}
}

func (m *MediaClient) CreateOne(file mediaDto.FileDataType) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.client.CreateOne(ctx, &mediaGRPC.CreateOneRequest{
		File: &mediaGRPC.FileDataType{
			FileMetadata: &mediaGRPC.FileMetadata{
				FileName:    file.FileMetadata.FileName,
				ContentType: file.FileMetadata.ContentType,
				Size:        file.FileMetadata.Size,
			},
			Data: file.Data,
		},
	})
	if err != nil {
		return "", err
	}

	return res.GetObjectId(), nil
}

func (m *MediaClient) CreateMany(files []mediaDto.FileDataType) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var requestFiles []*mediaGRPC.FileDataType

	for _, file := range files {
		requestFiles = append(requestFiles, &mediaGRPC.FileDataType{
			FileMetadata: &mediaGRPC.FileMetadata{
				FileName:    file.FileMetadata.FileName,
				ContentType: file.FileMetadata.ContentType,
				Size:        file.FileMetadata.Size,
			},
			Data: file.Data,
		})
	}

	res, err := m.client.CreateMany(ctx, &mediaGRPC.CreateManyRequest{
		Files: requestFiles,
	})
	if err != nil {
		return nil, err
	}

	return res.GetObjectIds(), nil
}

func (m *MediaClient) GetOne(objectID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.client.GetOne(ctx, &mediaGRPC.GetOneRequest{ObjectId: objectID})
	if err != nil {
		return "", err
	}

	return res.GetUrl(), nil
}

func (m *MediaClient) GetMany(objectIDs []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.client.GetMany(ctx, &mediaGRPC.GetManyRequest{ObjectIds: objectIDs})
	if err != nil {
		return nil, err
	}

	return res.GetUrls(), nil
}

func (m *MediaClient) DeleteOne(objectID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.DeleteOne(ctx, &mediaGRPC.DeleteOneRequest{ObjectId: objectID})
	if err != nil {
		return err
	}

	return nil
}

func (m *MediaClient) DeleteMany(objectIDs []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.client.DeleteMany(ctx, &mediaGRPC.DeleteManyRequest{ObjectIds: objectIDs})
	if err != nil {
		return err
	}

	return nil
}
