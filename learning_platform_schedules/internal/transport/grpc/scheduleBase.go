package grpc

import (
	"context"
	scheduleGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/schedule"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models/enum"
)

type ScheduleService interface {
	GetAllSchedules() ([]dto.ScheduleResponse, error)
	GetScheduleByID(scheduleID int64) (*dto.ScheduleResponse, error)
	GetSchedulesByTutorID(tutorID int64) ([]dto.ScheduleResponse, error)
	CreateSchedule(newSchedule dto.CreateSchedule) (*dto.ScheduleResponse, error)
	UpdateSchedule(newSchedule dto.UpdateSchedule) (*dto.ScheduleResponse, error)
	DeleteSchedule(scheduleID int64) error
}

func (s *ScheduleGRPCServer) GetAllSchedules(
	ctx context.Context,
	in *scheduleGRPC.GetAllSchedulesRequest,
) (*scheduleGRPC.GetAllSchedulesResponse, error) {
	schedules, err := s.service.ScheduleService.GetAllSchedules()
	if err != nil {
		s.logger.Error("failed to get all schedules", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to get all schedules")
	}

	var resSchedules []*scheduleGRPC.GetScheduleByIDResponse
	for _, oneSchedule := range schedules {
		resSchedules = append(resSchedules, scheduleResponseDTOToProto(oneSchedule))
	}

	return &scheduleGRPC.GetAllSchedulesResponse{
		Schedules: resSchedules,
	}, nil
}

func (s *ScheduleGRPCServer) GetScheduleByID(
	ctx context.Context,
	in *scheduleGRPC.GetScheduleByIDRequest,
) (*scheduleGRPC.GetScheduleByIDResponse, error) {
	schedule, err := s.service.ScheduleService.GetScheduleByID(in.GetScheduleId())
	if err != nil {
		s.logger.Error(
			"failed to get schedule by id",
			zap.Int64("scheduleID", in.GetScheduleId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get schedule by id")
	}

	resSchedule := scheduleResponseDTOToProto(*schedule)

	return &scheduleGRPC.GetScheduleByIDResponse{
		Id:        resSchedule.GetId(),
		TutorId:   resSchedule.GetTutorId(),
		StartTime: resSchedule.GetStartTime(),
		EndTime:   resSchedule.GetEndTime(),
		Slots:     resSchedule.GetSlots(),
	}, nil
}

func (s *ScheduleGRPCServer) GetSchedulesByTutorID(
	ctx context.Context,
	in *scheduleGRPC.GetSchedulesByTutorIDRequest,
) (*scheduleGRPC.GetSchedulesByTutorIDResponse, error) {
	schedules, err := s.service.ScheduleService.GetSchedulesByTutorID(in.GetTutorId())
	if err != nil {
		s.logger.Error(
			"failed to get schedules by tutorID",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get schedules by tutorID")
	}

	var resSchedules []*scheduleGRPC.GetScheduleByIDResponse
	for _, oneSchedule := range schedules {
		resSchedules = append(resSchedules, scheduleResponseDTOToProto(oneSchedule))
	}

	return &scheduleGRPC.GetSchedulesByTutorIDResponse{
		Schedules: resSchedules,
	}, nil
}

func (s *ScheduleGRPCServer) CreateSchedule(
	ctx context.Context,
	in *scheduleGRPC.CreateScheduleRequest,
) (*scheduleGRPC.CreateScheduleResponse, error) {
	var createSlots []dto.CreateScheduleSlot

	for _, slot := range in.GetSlots() {
		createSlots = append(createSlots, dto.CreateScheduleSlot{
			StartTime: slot.StartTime.AsTime(),
			Duration:  slot.Duration,
			LessonID:  slot.LessonId,
		})
	}

	createDto := dto.CreateSchedule{
		TutorID:   in.GetTutorId(),
		StartTime: in.GetStartTime().AsTime(),
		EndTime:   in.GetEndTime().AsTime(),
		Slots:     createSlots,
	}

	schedule, err := s.service.ScheduleService.CreateSchedule(createDto)
	if err != nil {
		s.logger.Error(
			"failed to create schedule",
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to create schedule")
	}

	resSchedule := scheduleResponseDTOToProto(*schedule)

	return &scheduleGRPC.CreateScheduleResponse{
		Schedule: resSchedule,
	}, nil
}

func (s *ScheduleGRPCServer) UpdateSchedule(
	ctx context.Context,
	in *scheduleGRPC.UpdateScheduleRequest,
) (*scheduleGRPC.UpdateScheduleResponse, error) {
	var createSlots []dto.CreateScheduleSlot

	for _, slot := range in.GetSlots() {
		createSlots = append(createSlots, dto.CreateScheduleSlot{
			StartTime: slot.StartTime.AsTime(),
			Duration:  slot.Duration,
			LessonID:  slot.LessonId,
		})
	}

	createDto := dto.UpdateSchedule{
		ID:                    in.GetId(),
		StartTime:             in.GetStartTime().AsTime(),
		EndTime:               in.GetEndTime().AsTime(),
		Slots:                 createSlots,
		DeleteScheduleSlotIDs: in.GetDeletedScheduleSlotIds(),
	}

	schedule, err := s.service.ScheduleService.UpdateSchedule(createDto)
	if err != nil {
		s.logger.Error(
			"failed to update schedule",
			zap.Int64("scheduleID", in.GetId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update schedule")
	}

	resSchedule := scheduleResponseDTOToProto(*schedule)

	return &scheduleGRPC.UpdateScheduleResponse{
		Schedule: resSchedule,
	}, nil
}

func (s *ScheduleGRPCServer) DeleteSchedule(
	ctx context.Context,
	in *scheduleGRPC.DeleteScheduleRequest,
) (*scheduleGRPC.DeleteScheduleResponse, error) {
	err := s.service.ScheduleService.DeleteSchedule(in.GetScheduleId())
	if err != nil {
		s.logger.Error(
			"failed to delete schedule",
			zap.Int64("scheduleID", in.GetScheduleId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete schedule")
	}

	return &scheduleGRPC.DeleteScheduleResponse{}, nil
}

func scheduleResponseDTOToProto(schedule dto.ScheduleResponse) *scheduleGRPC.GetScheduleByIDResponse {
	var resScheduleSlots []*scheduleGRPC.ScheduleSlot
	for _, oneSlot := range schedule.Slots {
		resScheduleSlots = append(resScheduleSlots, &scheduleGRPC.ScheduleSlot{
			Id:         oneSlot.ID,
			ScheduleId: oneSlot.ScheduleID,
			StartTime:  timestamppb.New(oneSlot.StartTime),
			Status:     enumToProtoStatus(oneSlot.Status),
			Duration:   oneSlot.Duration,
			LessonId:   oneSlot.LessonID,
		})
	}

	return &scheduleGRPC.GetScheduleByIDResponse{
		Id:        schedule.ID,
		TutorId:   schedule.TutorID,
		StartTime: timestamppb.New(schedule.StartTime),
		EndTime:   timestamppb.New(schedule.EndTime),
		Slots:     resScheduleSlots,
	}
}

func enumToProtoStatus(status enum.ScheduleSlotStatus) scheduleGRPC.Status {
	switch status {
	case enum.StatusFree:
		return scheduleGRPC.Status_FREE
	case enum.StatusBooked:
		return scheduleGRPC.Status_BOOKED
	default:
		return scheduleGRPC.Status_SCHEDULE_STATUS_UNSPECIFIED
	}
}
