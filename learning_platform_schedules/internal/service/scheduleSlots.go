package service

import (
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models"
)

type ScheduleSlotService struct {
	storage ScheduleSlotStorage
}

type ScheduleSlotStorage interface {
	UpdateScheduleSlot(scheduleSlotID int64, scheduleSlot dto.UpdateScheduleSlot) (*models.ScheduleSlot, error)
	BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error
	DeleteLessonFromScheduleSlot(scheduleSlotID int64) error
	CreateScheduleSlot(slot dto.CreateScheduleSlot) (*models.ScheduleSlot, error)
	DeleteOneScheduleSlot(scheduleSlotID int64) error
}

func NewScheduleSlotService(storage ScheduleSlotStorage) *ScheduleSlotService {
	return &ScheduleSlotService{
		storage: storage,
	}
}

func (ss *ScheduleSlotService) UpdateScheduleSlot(
	scheduleSlotID int64,
	updateSlot dto.UpdateScheduleSlot,
) (*models.ScheduleSlot, error) {
	scheduleSlot, err := ss.storage.UpdateScheduleSlot(scheduleSlotID, updateSlot)
	if err != nil {
		return nil, err
	}

	return scheduleSlot, nil
}

func (ss *ScheduleSlotService) BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error {
	err := ss.storage.BindLessonToScheduleSlot(scheduleSlotID, lessonID)
	if err != nil {
		return err
	}

	return nil
}

func (ss *ScheduleSlotService) DeleteLessonFromScheduleSlot(scheduleSlotID int64) error {
	err := ss.storage.DeleteLessonFromScheduleSlot(scheduleSlotID)
	if err != nil {
		return err
	}

	return nil
}

func (ss *ScheduleSlotService) CreateScheduleSlot(
	slot dto.CreateScheduleSlot,
) (*models.ScheduleSlot, error) {
	res, err := ss.storage.CreateScheduleSlot(slot)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *ScheduleSlotService) DeleteOneScheduleSlot(scheduleSlotID int64) error {
	err := ss.storage.DeleteOneScheduleSlot(scheduleSlotID)
	if err != nil {
		return err
	}

	return nil
}
