package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/service"
	"learning-platform/auth/internal/storage/postgres"
	"learning-platform/auth/internal/storage/redis"
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

	postgresConn, err := postgres.Connection(cfg.DBDSN, log)
	if err != nil {
		log.Fatal("error connect to postgres", zap.Error(err))
	}
	defer postgresConn.Close()

	postgresStorage := postgres.New(postgresConn, log)

	redisConn, err := redis.Connection(cfg.RedisUrl)
	if err != nil {
		log.Fatal("error connect to redis", zap.Error(err))
	}
	defer redisConn.Close()

	redisStorage := redis.New(redisConn, log)

	service := service.New(cfg, log, postgresStorage, redisStorage)

	gRPCServer := grpc.New(cfg, log, service)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := gRPCServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
