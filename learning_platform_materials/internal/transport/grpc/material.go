package grpc

import (
	materialGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/material"
	"go.uber.org/zap"
)

type MaterialsGRPCServer struct {
	materialGRPC.UnimplementedMaterialServer
	service *MaterialService
	logger  *zap.Logger
}

type MaterialService struct {
	MaterialBaseService   MaterialBaseService
	MaterialFolderService MaterialFolderService
	MaterialUserService   MaterialUserService
}

func NewMaterialsGRPCServer(
	service *MaterialService,
	logger *zap.Logger,
) *MaterialsGRPCServer {
	return &MaterialsGRPCServer{
		service: service,
		logger:  logger,
	}
}
