package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/media/internal/config"
	"learning-platform/media/internal/s3"
	"learning-platform/media/internal/service"
	"learning-platform/media/internal/transport/grpc"
	"learning-platform/media/pkg/logger"
)

func StartApp() {
	cfg := config.GetConfig()

	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	log := zapLog.ZapLogger

	minio, err := s3.New(cfg.MinioUrl, cfg.MinioRootUser, cfg.MinioRootPassword, cfg.MinioBucket)
	if err != nil {
		log.Fatal("error init minio: ", zap.Error(err))
	}

	err = minio.CreateBucket(cfg.MinioBucket)
	if err != nil {
		log.Fatal("error minio create bucket: ", zap.Error(err))
	}

	serviceLayer := service.New(minio)

	grpcServer := grpc.New(log, cfg, serviceLayer)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
