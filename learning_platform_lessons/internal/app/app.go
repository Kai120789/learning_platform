package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/lessons/internal/config"
	"learning-platform/lessons/internal/service"
	"learning-platform/lessons/internal/storage"
	"learning-platform/lessons/internal/transport/grpc"
	"learning-platform/lessons/pkg/logger"
)

func StartApp() {
	cfg := config.GetConfig()

	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	log := zapLog.ZapLogger

	dbConn, err := storage.Connection(cfg.DBDSN)
	if err != nil {
		log.Fatal("error connect to db: ", zap.Error(err))
	}
	defer dbConn.Close()

	storageLayer := storage.New(dbConn)

	serviceLayer := service.New(
		&service.Storage{
			LessonStorage:      storageLayer.LessonStorage,
			LessonMediaStorage: storageLayer.LessonMediaStorage,
			LessonUserStorage:  storageLayer.LessonUserStorage,
		},
	)

	grpcServer := grpc.New(log, cfg, serviceLayer)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
