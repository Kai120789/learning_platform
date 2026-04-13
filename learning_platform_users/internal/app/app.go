package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/users/internal/config"
	"learning-platform/users/internal/service"
	groupService "learning-platform/users/internal/service/group"
	userService "learning-platform/users/internal/service/user"
	"learning-platform/users/internal/storage"
	"learning-platform/users/internal/transport/grpc"
	"learning-platform/users/internal/transport/grpc/group"
	"learning-platform/users/internal/transport/grpc/user"
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

	storageLayer := storage.New(dbConn)

	serviceLayer := service.New(&userService.UserStorage{
		UserBaseStorage:     storageLayer.UserStorage.UserBaseStorage,
		UserInfoStorage:     storageLayer.UserStorage.UserInfoStorage,
		UserSettingsStorage: storageLayer.UserStorage.UserSettingsStorage,
	}, &groupService.GroupStorage{
		GroupBaseStorage: storageLayer.GroupStorage.GroupBaseStorage,
		GroupUserStorage: storageLayer.GroupStorage.GroupUserStorage,
	})

	grpcServer := grpc.New(log, cfg, &user.UserGRPCServer{
		UserBaseService:     serviceLayer.UserService.UserBaseService,
		UserInfoService:     serviceLayer.UserService.UserInfoService,
		UserSettingsService: serviceLayer.UserService.UserSettingsService,
	}, &group.GroupGRPCServer{
		GroupBaseService: serviceLayer.GroupService.GroupBaseService,
		GroupUserService: nil,
	})

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
