package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/materials/internal/config"
	"learning-platform/materials/internal/service"
	"learning-platform/materials/internal/storage"
	"learning-platform/materials/internal/transport/grpc"
	"learning-platform/materials/pkg/logger"
)

func StartApp() {
	cfg := config.GetConfig()

	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	log := zapLog.ZapLogger

	conn, err := storage.Connection(cfg.DBDSN)
	if err != nil {
		log.Fatal("error connect to db: ", zap.Error(err))
	}
	defer conn.Close()

	storageLayer := storage.New(conn)

	serviceLayer := service.New(&service.Storage{
		MaterialStorage:       storageLayer.MaterialStorage,
		MaterialFolderStorage: storageLayer.MaterialFolderStorage,
		MaterialUserStorage:   storageLayer.MaterialUserStorage,
	})

	grpcServer := grpc.New(cfg, log, &grpc.MaterialService{
		MaterialBaseService:   serviceLayer.MaterialService,
		MaterialFolderService: serviceLayer.MaterialFolderService,
		MaterialUserService:   serviceLayer.MaterialUserService,
	})

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
