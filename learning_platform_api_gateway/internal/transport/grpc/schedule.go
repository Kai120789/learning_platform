package grpc

import (
	"context"
	scheduleGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/schedule"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/scheduleDto"
	"time"
)

type ScheduleClient struct {
	client scheduleGRPC.ScheduleClient
}

func NewScheduleGrpcConnection(scheduleGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		scheduleGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewScheduleClient(connection *grpc.ClientConn) *ScheduleClient {
	return &ScheduleClient{
		client: scheduleGRPC.NewScheduleClient(connection),
	}
}

func (s *ScheduleClient) GetAllSchedules() ([]scheduleDto.ScheduleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetAllSchedules(ctx, &scheduleGRPC.GetAllSchedulesRequest{})
	if err != nil {
		return nil, err
	}

	var resSchedules []scheduleDto.ScheduleResponse
	for _, oneSchedule := range res.GetSchedules() {
		resSchedules = append(resSchedules, *grpcScheduleToDTO(oneSchedule))
	}

	return resSchedules, nil
}

func (s *ScheduleClient) GetScheduleByID(scheduleID int64) (*scheduleDto.ScheduleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetScheduleByID(ctx, &scheduleGRPC.GetScheduleByIDRequest{ScheduleId: scheduleID})
	if err != nil {
		return nil, err
	}

	return grpcScheduleToDTO(res), nil
}

func (s *ScheduleClient) GetSchedulesByTutorID(tutorID int64) ([]scheduleDto.ScheduleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetSchedulesByTutorID(ctx, &scheduleGRPC.GetSchedulesByTutorIDRequest{TutorId: tutorID})
	if err != nil {
		return nil, err
	}

	var resSchedules []scheduleDto.ScheduleResponse
	for _, oneSchedule := range res.GetSchedules() {
		resSchedules = append(resSchedules, *grpcScheduleToDTO(oneSchedule))
	}

	return resSchedules, nil
}

func (s *ScheduleClient) CreateSchedule(schedule scheduleDto.CreateSchedule) (*scheduleDto.ScheduleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var slots []*scheduleGRPC.CreateScheduleSlot

	for _, oneSlot := range schedule.Slots {
		slots = append(slots, &scheduleGRPC.CreateScheduleSlot{
			StartTime: timestamppb.New(oneSlot.StartTime),
			Duration:  oneSlot.Duration,
			LessonId:  oneSlot.LessonID,
		})
	}

	newSchedule := &scheduleGRPC.CreateScheduleRequest{
		TutorId:   schedule.TutorID,
		StartTime: timestamppb.New(schedule.StartTime),
		EndTime:   timestamppb.New(schedule.EndTime),
		Slots:     slots,
	}

	res, err := s.client.CreateSchedule(ctx, newSchedule)
	if err != nil {
		return nil, err
	}

	return grpcScheduleToDTO(res.GetSchedule()), nil
}

func (s *ScheduleClient) UpdateSchedule(schedule scheduleDto.UpdateSchedule) (*scheduleDto.ScheduleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newSlots []*scheduleGRPC.CreateScheduleSlot

	for _, oneSlot := range schedule.Slots {
		newSlots = append(newSlots, &scheduleGRPC.CreateScheduleSlot{
			StartTime: timestamppb.New(oneSlot.StartTime),
			Duration:  oneSlot.Duration,
			LessonId:  oneSlot.LessonID,
		})
	}

	newSchedule := &scheduleGRPC.UpdateScheduleRequest{
		Id:                     schedule.ID,
		StartTime:              timestamppb.New(schedule.StartTime),
		EndTime:                timestamppb.New(schedule.EndTime),
		Slots:                  newSlots,
		DeletedScheduleSlotIds: schedule.DeleteScheduleSlotIDs,
	}

	res, err := s.client.UpdateSchedule(ctx, newSchedule)
	if err != nil {
		return nil, err
	}

	return grpcScheduleToDTO(res.GetSchedule()), nil
}

func (s *ScheduleClient) DeleteSchedule(scheduleID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.DeleteSchedule(ctx, &scheduleGRPC.DeleteScheduleRequest{ScheduleId: scheduleID})
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleClient) UpdateScheduleSlot(scheduleSlotID int64, updatedSlot scheduleDto.CreateScheduleSlot) (*scheduleDto.ScheduleSlot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpdateScheduleSlot(ctx, &scheduleGRPC.UpdateScheduleSlotRequest{
		Id:        scheduleSlotID,
		StartTime: timestamppb.New(updatedSlot.StartTime),
		Duration:  updatedSlot.Duration,
		LessonId:  updatedSlot.LessonID,
	})
	if err != nil {
		return nil, err
	}

	return &scheduleDto.ScheduleSlot{
		ID:         res.GetId(),
		ScheduleID: res.GetScheduleId(),
		StartTime:  res.GetStartTime().AsTime(),
		Status:     protoToEnumStatus(res.GetStatus()),
		Duration:   res.Duration,
		LessonID:   res.LessonId,
	}, nil
}

func (s *ScheduleClient) BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.BindLessonToScheduleSlot(ctx, &scheduleGRPC.BindLessonToScheduleSlotRequest{
		ScheduleSlotId: scheduleSlotID,
		LessonId:       lessonID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleClient) DeleteLessonFromScheduleSlot(scheduleSlotID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.DeleteLessonFromScheduleSlot(ctx, &scheduleGRPC.DeleteLessonFromScheduleSlotRequest{
		ScheduleSlotId: scheduleSlotID,
	})
	if err != nil {
		return err
	}

	return nil
}

func grpcScheduleToDTO(schedule *scheduleGRPC.GetScheduleByIDResponse) *scheduleDto.ScheduleResponse {
	var resSlots []scheduleDto.ScheduleSlot

	for _, oneSlot := range schedule.GetSlots() {
		resSlots = append(resSlots, scheduleDto.ScheduleSlot{
			ID:         oneSlot.GetId(),
			ScheduleID: oneSlot.GetScheduleId(),
			StartTime:  oneSlot.GetStartTime().AsTime(),
			Status:     protoToEnumStatus(oneSlot.GetStatus()),
			Duration:   oneSlot.Duration,
			LessonID:   oneSlot.LessonId,
		})
	}

	return &scheduleDto.ScheduleResponse{
		ID:        schedule.GetId(),
		TutorID:   schedule.GetTutorId(),
		StartTime: schedule.GetStartTime().AsTime(),
		EndTime:   schedule.GetEndTime().AsTime(),
		Slots:     resSlots,
	}
}

func protoToEnumStatus(status scheduleGRPC.Status) enum.ScheduleSlotStatus {
	switch status {
	case scheduleGRPC.Status_FREE:
		return enum.StatusFree
	case scheduleGRPC.Status_BOOKED:
		return enum.StatusBooked
	default:
		return ""
	}
}
