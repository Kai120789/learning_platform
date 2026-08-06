package grpc

import (
	"context"
	mediaGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/media"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/media/internal/dto"
)

type MediaGRPCServer struct {
	mediaGRPC.UnimplementedMediaServer
	service Service
	logger  *zap.Logger
}

type Service interface {
	CreateOne(ctx context.Context, file dto.FileDataType) (string, error)
	CreateMany(ctx context.Context, files []dto.FileDataType) ([]string, error)
	//CreateOneStream(ctx context.Context) (string, error)
	GetOne(ctx context.Context, objectID string) (string, error)
	GetMany(ctx context.Context, objectIDs []string) ([]string, error)
	DeleteOne(ctx context.Context, objectID string) error
	DeleteMany(ctx context.Context, objectIDs []string) error
}

func NewMediaGRPCServer(
	logger *zap.Logger,
	service Service,
) mediaGRPC.MediaServer {
	return &MediaGRPCServer{
		logger:  logger,
		service: service,
	}
}

func (m *MediaGRPCServer) CreateOne(
	ctx context.Context,
	in *mediaGRPC.CreateOneRequest,
) (*mediaGRPC.CreateOneResponse, error) {
	res, err := m.service.CreateOne(ctx, dto.FileDataType{
		Data: in.GetFile().GetData(),
		FileMetadata: dto.FileMetadata{
			FileName:    in.GetFile().GetFileMetadata().GetFileName(),
			ContentType: in.GetFile().GetFileMetadata().GetContentType(),
			Size:        in.GetFile().GetFileMetadata().GetSize(),
		},
	})
	if err != nil {
		m.logger.Error("failed to upload one file", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to upload one file")
	}

	return &mediaGRPC.CreateOneResponse{
		ObjectId: res,
	}, nil
}

func (m *MediaGRPCServer) CreateMany(
	ctx context.Context,
	in *mediaGRPC.CreateManyRequest,
) (*mediaGRPC.CreateManyResponse, error) {
	var requestFiles []dto.FileDataType

	for _, file := range in.GetFiles() {
		requestFiles = append(requestFiles, dto.FileDataType{
			Data: file.GetData(),
			FileMetadata: dto.FileMetadata{
				FileName:    file.GetFileMetadata().GetFileName(),
				ContentType: file.GetFileMetadata().GetContentType(),
				Size:        file.GetFileMetadata().GetSize(),
			},
		})
	}

	res, err := m.service.CreateMany(ctx, requestFiles)
	if err != nil {
		m.logger.Error("failed to upload many files", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to upload many files")
	}

	return &mediaGRPC.CreateManyResponse{
		ObjectIds: res,
	}, nil
}

func (m *MediaGRPCServer) GetOne(
	ctx context.Context,
	in *mediaGRPC.GetOneRequest,
) (*mediaGRPC.GetOneResponse, error) {
	res, err := m.service.GetOne(ctx, in.GetObjectId())
	if err != nil {
		m.logger.Error(
			"failed to get one file",
			zap.String("objectID", in.GetObjectId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get one file")
	}

	return &mediaGRPC.GetOneResponse{
		Url: res,
	}, nil
}

func (m *MediaGRPCServer) GetMany(
	ctx context.Context,
	in *mediaGRPC.GetManyRequest,
) (*mediaGRPC.GetManyResponse, error) {
	res, err := m.service.GetMany(ctx, in.GetObjectIds())
	if err != nil {
		m.logger.Error(
			"failed to get many files",
			zap.Strings("objectIDs", in.GetObjectIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get many files")
	}

	return &mediaGRPC.GetManyResponse{
		Urls: res,
	}, nil
}

func (m *MediaGRPCServer) DeleteOne(
	ctx context.Context,
	in *mediaGRPC.DeleteOneRequest,
) (*mediaGRPC.DeleteOneResponse, error) {
	err := m.service.DeleteOne(ctx, in.GetObjectId())
	if err != nil {
		m.logger.Error(
			"failed to delete one file",
			zap.String("objectID", in.GetObjectId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete one file")
	}

	return &mediaGRPC.DeleteOneResponse{}, nil
}

func (m *MediaGRPCServer) DeleteMany(
	ctx context.Context,
	in *mediaGRPC.DeleteManyRequest,
) (*mediaGRPC.DeleteManyResponse, error) {
	err := m.service.DeleteMany(ctx, in.GetObjectIds())
	if err != nil {
		m.logger.Error(
			"failed to delete many files",
			zap.Strings("objectIDs", in.GetObjectIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete many files")
	}

	return &mediaGRPC.DeleteManyResponse{}, nil
}
