package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/groups/internal/config"
	"learning-platform/groups/internal/service"
	"learning-platform/groups/internal/storage"
	"learning-platform/groups/internal/transport/grpc"
	"learning-platform/groups/pkg/logger"
)

func StartApp() {
	cfg := config.Getconfig()

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
		GroupUserStorage: storageLayer.GroupUserStorage,
		GroupBaseStorage: storageLayer.GroupBaseStorage,
	})

	grpcServer := grpc.New(cfg, log, &grpc.GroupService{
		GroupBaseService: serviceLayer.GroupBaseService,
		GroupUserService: serviceLayer.GroupUserService,
	})

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
