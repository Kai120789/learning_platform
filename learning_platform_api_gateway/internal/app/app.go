package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/redis"
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

	// TODO: init handler

	// TODO: init router

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: http.NewServeMux(),
	}

	log.Info("http server started", zap.String("address", cfg.ServerAddress))
	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start http server", zap.Error(err))
	}
}
