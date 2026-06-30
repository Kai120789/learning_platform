package grpc

import (
	scheduleGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/schedule"
	"go.uber.org/zap"
)

type ScheduleGRPCServer struct {
	scheduleGRPC.UnimplementedScheduleServer
	service *Service
	logger  *zap.Logger
}

type Service struct {
	ScheduleService     ScheduleService
	ScheduleSlotService ScheduleSlotService
}

func NewScheduleGRPCServer(
	logger *zap.Logger,
	service *Service,
) scheduleGRPC.ScheduleServer {
	return &ScheduleGRPCServer{
		logger:  logger,
		service: service,
	}
}
