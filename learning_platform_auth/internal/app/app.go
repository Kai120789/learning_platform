package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
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

	// TODO: connect postgres

	// TODO: connect redis

	service := service.New(cfg, log, nil)

	gRPCServer := grpc.New(cfg, log, service)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := gRPCServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
