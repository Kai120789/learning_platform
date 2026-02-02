package app

import (
	"fmt"
	userGrpc "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/redis"
	"learning-platform/auth/internal/service"
	"learning-platform/auth/internal/transport/grpc"
	"learning-platform/auth/pkg/logger"
)

func Start() {
	cfg := config.GetConfig()

	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	log := zapLog.ZapLogger

	redisConn, err := redis.Connection(cfg.RedisUrl)
	if err != nil {
		log.Fatal("error connect to redis", zap.Error(err))
	}
	defer redisConn.Close()

	redisStorage := redis.New(redisConn, log, cfg)

	userGrpcConn, err := grpc.NewUserGrpcClient(cfg.UserServiceUrl, log)
	if err != nil {
		log.Error("user grpc connect error", zap.Error(err))
	}
	defer userGrpcConn.Close()

	userClient := userGrpc.NewUserClient(userGrpcConn)

	userApi := grpc.NewUserApi(userClient, log)

	serviceLayer := service.New(cfg, log, redisStorage, userApi)

	gRPCServer := grpc.New(cfg, log, serviceLayer)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := gRPCServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
