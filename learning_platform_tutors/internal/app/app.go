package app

import (
	"fmt"
	"go.uber.org/zap"
	"learning-platform/tutors/internal/config"
	"learning-platform/tutors/internal/service"
	"learning-platform/tutors/internal/storage"
	"learning-platform/tutors/internal/transport/grpc"
	"learning-platform/tutors/pkg/logger"
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
			TutorReviewStorage:  storageLayer.TutorReviewStorage,
			TutorOfferStorage:   storageLayer.TutorOfferStorage,
			TutorStudentStorage: storageLayer.TutorStudentStorage,
			TutorSubjectStorage: storageLayer.TutorSubjectStorage,
		},
	)

	grpcServer := grpc.New(
		cfg,
		log,
		serviceLayer.TutorService,
		serviceLayer.TutorReviewService,
		serviceLayer.TutorOfferService,
		serviceLayer.TutorStudentService,
		serviceLayer.TutorSubjectService,
	)

	log.Info("grpc server started", zap.String("address", cfg.GRPCServerAddress))
	if err := grpcServer.Run(); err != nil {
		log.Error("failed to start gRPC server", zap.Error(err))
	}
}
