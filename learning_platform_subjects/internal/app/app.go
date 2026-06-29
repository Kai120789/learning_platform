package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/subjects/internal/config"
	"learning-platform/subjects/internal/service"
	"learning-platform/subjects/internal/storage"
	"learning-platform/subjects/internal/transport/grpc"
	"learning-platform/subjects/pkg/logger"
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
			SubjectStorage:     storageLayer.SubjectStorage,
			UserSubjectStorage: storageLayer.UserSubjectStorage,
		},
	)

	grpcServer := grpc.New(
		log,
		cfg,
		serviceLayer.SubjectService,
		serviceLayer.UserSubjectService,
	)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
