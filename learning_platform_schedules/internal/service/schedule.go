package service

import (
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models"
	"learning-platform/schedules/internal/utils"
)

type ScheduleService struct {
	scheduleStorage     ScheduleStorage
	scheduleSlotStorage ScheduleSlotsForScheduleStorage
}

type ScheduleStorage interface {
	GetScheduleByID(scheduleID int64) (*models.Schedule, error)
	GetAllSchedules() ([]models.Schedule, error)
	GetSchedulesByTutorID(tutorID int64) ([]models.Schedule, error)
	CreateSchedule(schedule dto.CreateSchedule) (*models.Schedule, error)
	UpdateSchedule(schedule dto.UpdateSchedule) (*models.Schedule, error)
	DeleteSchedule(scheduleID int64) error
}

type ScheduleSlotsForScheduleStorage interface {
	GetAllScheduleSlots(scheduleID int64) ([]models.ScheduleSlot, error)
	SetScheduleSlots(scheduleID int64, slots []dto.CreateScheduleSlot) ([]models.ScheduleSlot, error)
	DeleteScheduleSlots(scheduleSlotIDs []int64) error
}

func NewScheduleService(
	scheduleStorage ScheduleStorage,
	scheduleSlotStorage ScheduleSlotsForScheduleStorage,
) *ScheduleService {
	return &ScheduleService{
		scheduleStorage:     scheduleStorage,
		scheduleSlotStorage: scheduleSlotStorage,
	}
}

func (s *ScheduleService) GetAllSchedules() ([]dto.ScheduleResponse, error) {
	var resSchedules []dto.ScheduleResponse
	schedules, err := s.scheduleStorage.GetAllSchedules()
	if err != nil {
		return nil, err
	}

	for _, oneSchedule := range schedules {
		resOneSchedule, err := s.buildScheduleWithSlotsDTO(&oneSchedule, nil)
		if err != nil {
			return nil, err
		}

		resSchedules = append(resSchedules, *resOneSchedule)
	}
	return resSchedules, nil
}

func (s *ScheduleService) GetScheduleByID(scheduleID int64) (*dto.ScheduleResponse, error) {
	schedule, err := s.scheduleStorage.GetScheduleByID(scheduleID)
	if err != nil {
		return nil, err
	}

	resSchedule, err := s.buildScheduleWithSlotsDTO(schedule, nil)
	if err != nil {
		return nil, err
	}

	return resSchedule, nil
}

func (s *ScheduleService) GetSchedulesByTutorID(tutorID int64) ([]dto.ScheduleResponse, error) {
	var resSchedules []dto.ScheduleResponse
	schedules, err := s.scheduleStorage.GetSchedulesByTutorID(tutorID)
	if err != nil {
		return nil, err
	}

	for _, oneSchedule := range schedules {
		resOneSchedule, err := s.buildScheduleWithSlotsDTO(&oneSchedule, nil)
		if err != nil {
			return nil, err
		}

		resSchedules = append(resSchedules, *resOneSchedule)
	}
	return resSchedules, nil
}

func (s *ScheduleService) CreateSchedule(newSchedule dto.CreateSchedule) (*dto.ScheduleResponse, error) {
	schedule, err := s.scheduleStorage.CreateSchedule(newSchedule)
	if err != nil {
		return nil, err
	}

	scheduleSlots, err := s.scheduleSlotStorage.SetScheduleSlots(schedule.ID, newSchedule.Slots)
	if err != nil {
		return nil, err
	}

	resSchedule, err := s.buildScheduleWithSlotsDTO(schedule, scheduleSlots)
	if err != nil {
		return nil, err
	}

	return resSchedule, nil
}

func (s *ScheduleService) UpdateSchedule(newSchedule dto.UpdateSchedule) (*dto.ScheduleResponse, error) {
	schedule, err := s.scheduleStorage.UpdateSchedule(newSchedule)
	if err != nil {
		return nil, err
	}

	err = s.scheduleSlotStorage.DeleteScheduleSlots(newSchedule.DeleteScheduleSlotIDs)
	if err != nil {
		return nil, err
	}

	_, err = s.scheduleSlotStorage.SetScheduleSlots(schedule.ID, newSchedule.Slots)
	if err != nil {
		return nil, err
	}

	resSchedule, err := s.buildScheduleWithSlotsDTO(schedule, nil)
	if err != nil {
		return nil, err
	}

	return resSchedule, nil
}

func (s *ScheduleService) DeleteSchedule(scheduleID int64) error {
	err := s.scheduleStorage.DeleteSchedule(scheduleID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleService) buildScheduleWithSlotsDTO(
	schedule *models.Schedule,
	newScheduleSlots []models.ScheduleSlot,
) (*dto.ScheduleResponse, error) {
	var scheduleSlots []models.ScheduleSlot
	if newScheduleSlots != nil {
		scheduleSlots = newScheduleSlots
	} else {
		dbScheduleSlots, err := s.scheduleSlotStorage.GetAllScheduleSlots(schedule.ID)
		if err != nil {
			return nil, err
		}
		scheduleSlots = dbScheduleSlots
	}

	var resScheduleSlots []dto.ScheduleSlot
	for _, oneSlot := range scheduleSlots {
		resScheduleSlots = append(resScheduleSlots, dto.ScheduleSlot{
			ID:         oneSlot.ID,
			ScheduleID: oneSlot.ScheduleID,
			StartTime:  oneSlot.StartTime.Time,
			Status:     oneSlot.Status,
			Duration:   utils.DBInt8ToOptional(oneSlot.Duration),
			LessonID:   utils.DBInt8ToOptional(oneSlot.LessonID),
		})
	}

	return &dto.ScheduleResponse{
		ID:        schedule.ID,
		TutorID:   schedule.TutorID,
		StartTime: schedule.StartTime.Time,
		EndTime:   schedule.EndTime.Time,
		Slots:     resScheduleSlots,
	}, nil
}
