package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/redis"
	"learning-platform/api-gateway/internal/service"
	"learning-platform/api-gateway/internal/transport/grpc"
	"learning-platform/api-gateway/internal/transport/http/handler"
	"learning-platform/api-gateway/internal/transport/http/router"
	"learning-platform/api-gateway/pkg/logger"
	"net/http"
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

	client, err := grpc.NewClient(
		cfg.UserServiceUrl,
		cfg.AuthServiceUrl,
		log,
	)
	if err != nil {
		log.Fatal("init grpc client error", zap.Error(err))
	}

	serviceLayer := service.New(&service.Client{
		UserClient: client.UserClient,
		AuthClient: client.AuthClient,
	}, log)

	_ = serviceLayer

	handlerLayer := handler.New(&handler.Service{
		AuthService: serviceLayer.AuthService,
		UserService: serviceLayer.UserService,
	}, log)

	r := router.New(&router.Handler{
		AuthHandler: handlerLayer.AuthHandler,
		UserHandler: handlerLayer.UserHandler,
	})

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	log.Info("http server started", zap.String("address", cfg.ServerAddress))
	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start http server", zap.Error(err))
	}
}
