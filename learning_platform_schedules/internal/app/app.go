package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/schedules/internal/config"
	"learning-platform/schedules/internal/service"
	"learning-platform/schedules/internal/storage"
	"learning-platform/schedules/internal/transport/grpc"
	"learning-platform/schedules/pkg/logger"
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
			ScheduleStorage:                 storageLayer.ScheduleStorage,
			ScheduleSlotStorage:             storageLayer.ScheduleSlotsStorage,
			ScheduleSlotsForScheduleStorage: storageLayer.ScheduleSlotsStorage,
		},
	)

	grpcServer := grpc.New(
		log,
		cfg,
		&grpc.Service{
			ScheduleService:     serviceLayer.ScheduleService,
			ScheduleSlotService: serviceLayer.ScheduleSlotService,
		},
	)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
