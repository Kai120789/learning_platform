package grpc

import (
	"context"
	scheduleGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/schedule"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models"
	"learning-platform/schedules/internal/utils"
)

type ScheduleSlotService interface {
	UpdateScheduleSlot(scheduleSlotID int64, updateSlot dto.CreateScheduleSlot) (*models.ScheduleSlot, error)
	BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error
	DeleteLessonFromScheduleSlot(scheduleSlotID int64) error
}

func (s *ScheduleGRPCServer) UpdateScheduleSlot(
	ctx context.Context,
	in *scheduleGRPC.UpdateScheduleSlotRequest,
) (*scheduleGRPC.UpdateScheduleSlotResponse, error) {
	updateSlot := dto.CreateScheduleSlot{
		StartTime: in.GetStartTime().AsTime(),
		Duration:  in.Duration,
		LessonID:  in.LessonId,
	}

	slot, err := s.service.ScheduleSlotService.UpdateScheduleSlot(in.GetId(), updateSlot)
	if err != nil {
		s.logger.Error(
			"failed to update schedule slot",
			zap.Int64("scheduleSlotID", in.GetId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update schedule slot")
	}

	return &scheduleGRPC.UpdateScheduleSlotResponse{
		Id:         slot.ID,
		ScheduleId: slot.ScheduleID,
		StartTime:  timestamppb.New(slot.StartTime.Time),
		Status:     enumToProtoStatus(slot.Status),
		Duration:   utils.DBInt8ToOptional(slot.Duration),
		LessonId:   utils.DBInt8ToOptional(slot.LessonID),
	}, nil
}

func (s *ScheduleGRPCServer) BindLessonToScheduleSlot(
	ctx context.Context,
	in *scheduleGRPC.BindLessonToScheduleSlotRequest,
) (*scheduleGRPC.BindLessonToScheduleSlotResponse, error) {
	err := s.service.ScheduleSlotService.BindLessonToScheduleSlot(in.GetScheduleSlotId(), in.GetLessonId())
	if err != nil {
		s.logger.Error(
			"failed to bind lesson to schedule slot",
			zap.Int64("lessonID", in.GetLessonId()),
			zap.Int64("scheduleSlotID", in.GetScheduleSlotId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to bind lesson to schedule slot")
	}

	return &scheduleGRPC.BindLessonToScheduleSlotResponse{}, nil
}

func (s *ScheduleGRPCServer) DeleteLessonFromScheduleSlot(
	ctx context.Context,
	in *scheduleGRPC.DeleteLessonFromScheduleSlotRequest,
) (*scheduleGRPC.DeleteLessonFromScheduleSlotResponse, error) {
	err := s.service.ScheduleSlotService.DeleteLessonFromScheduleSlot(in.GetScheduleSlotId())
	if err != nil {
		s.logger.Error(
			"failed to delete lesson from schedule slot",
			zap.Int64("scheduleSlotID", in.GetScheduleSlotId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete lesson from schedule slot")
	}
	return &scheduleGRPC.DeleteLessonFromScheduleSlotResponse{}, nil
}
