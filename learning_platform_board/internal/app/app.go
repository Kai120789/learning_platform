package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/board/internal/config"
	"learning-platform/board/internal/service"
	"learning-platform/board/internal/storage"
	"learning-platform/board/internal/transport/grpc"
	"learning-platform/board/internal/transport/ws"
	"learning-platform/board/pkg/logger"
	"net/http"
)

func StartApp() {
	cfg := config.GetConfig()

	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	log := zapLog.ZapLogger

	dbConn, err := storage.PostgresConnection(cfg.DBDSN)
	if err != nil {
		log.Fatal("error connect to db: ", zap.Error(err))
	}
	defer dbConn.Close()

	mongoClient, err := storage.MongoConnection(cfg.MongoUri)
	if err != nil {
		log.Fatal("error connect to mongo: ", zap.Error(err))
	}

	storageLayer := storage.New(dbConn, mongoClient)

	serviceLayer := service.New(
		&service.Storage{
			PostgresStorage: storageLayer.PostgresStorage,
			MongoStorage:    storageLayer.MongoStorage,
		},
	)

	grpcServer := grpc.New(
		cfg,
		log,
		serviceLayer.BoardService,
	)

	hub := ws.NewHub(serviceLayer.LiveBoardService)

	r := ws.NewWSRouter(hub)

	go func() {
		server := &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: r,
		}

		log.Info("http server started", zap.String("address", cfg.ServerAddress))
		if err := server.ListenAndServe(); err != nil {
			log.Error("failed to start http server", zap.Error(err))
		}
	}()

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
