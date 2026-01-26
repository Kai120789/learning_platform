package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/users/internal/config"
	"learning-platform/users/internal/service"
	"learning-platform/users/internal/storage"
	"learning-platform/users/internal/transport/grpc"
	"learning-platform/users/pkg/logger"
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

	storageLayer := storage.New(log, dbConn)

	serviceLayer := service.New(log, &service.Storage{
		UserStorage:         storageLayer.UserStorage,
		UserInfoStorage:     storageLayer.UserInfoStorage,
		UserSettingsStorage: storageLayer.UserSettingsStorage,
	})

	grpcServer := grpc.New(log, cfg, serviceLayer)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
