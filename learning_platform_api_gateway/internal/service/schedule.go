package service

import "learning-platform/api-gateway/internal/dto/scheduleDto"

type ScheduleService struct {
	client ScheduleClient
}

type ScheduleClient interface {
	GetAllSchedules() ([]scheduleDto.ScheduleResponse, error)
	GetScheduleByID(scheduleID int64) (*scheduleDto.ScheduleResponse, error)
	GetSchedulesByTutorID(tutorID int64) ([]scheduleDto.ScheduleResponse, error)
	CreateSchedule(schedule scheduleDto.CreateSchedule) (*scheduleDto.ScheduleResponse, error)
	UpdateSchedule(schedule scheduleDto.UpdateSchedule) (*scheduleDto.ScheduleResponse, error)
	DeleteSchedule(scheduleID int64) error
	UpdateScheduleSlot(scheduleSlotID int64, updatedSlot scheduleDto.CreateScheduleSlot) (*scheduleDto.ScheduleSlot, error)
	BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error
	DeleteLessonFromScheduleSlot(scheduleSlotID int64) error
}

func NewScheduleService(client ScheduleClient) *ScheduleService {
	return &ScheduleService{
		client: client,
	}
}

func (s *ScheduleService) GetAllSchedules() ([]scheduleDto.ScheduleResponse, error) {
	res, err := s.client.GetAllSchedules()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScheduleService) GetScheduleByID(scheduleID int64) (*scheduleDto.ScheduleResponse, error) {
	res, err := s.client.GetScheduleByID(scheduleID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScheduleService) GetSchedulesByTutorID(tutorID int64) ([]scheduleDto.ScheduleResponse, error) {
	res, err := s.client.GetSchedulesByTutorID(tutorID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScheduleService) CreateSchedule(schedule scheduleDto.CreateSchedule) (*scheduleDto.ScheduleResponse, error) {
	res, err := s.client.CreateSchedule(schedule)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScheduleService) UpdateSchedule(schedule scheduleDto.UpdateSchedule) (*scheduleDto.ScheduleResponse, error) {
	res, err := s.client.UpdateSchedule(schedule)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScheduleService) DeleteSchedule(scheduleID int64) error {
	err := s.client.DeleteSchedule(scheduleID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleService) UpdateScheduleSlot(
	scheduleSlotID int64,
	updatedSlot scheduleDto.CreateScheduleSlot,
) (*scheduleDto.ScheduleSlot, error) {
	res, err := s.client.UpdateScheduleSlot(scheduleSlotID, updatedSlot)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScheduleService) BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error {
	err := s.client.BindLessonToScheduleSlot(scheduleSlotID, lessonID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleService) DeleteLessonFromScheduleSlot(scheduleSlotID int64) error {
	err := s.client.DeleteLessonFromScheduleSlot(scheduleSlotID)
	if err != nil {
		return err
	}

	return nil
}
