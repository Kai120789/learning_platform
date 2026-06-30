package service

type Service struct {
	ScheduleService     *ScheduleService
	ScheduleSlotService *ScheduleSlotService
}

type Storage struct {
	ScheduleStorage                 ScheduleStorage
	ScheduleSlotStorage             ScheduleSlotStorage
	ScheduleSlotsForScheduleStorage ScheduleSlotsForScheduleStorage
}

func New(storage *Storage) *Service {
	return &Service{
		ScheduleService:     NewScheduleService(storage.ScheduleStorage, storage.ScheduleSlotsForScheduleStorage),
		ScheduleSlotService: NewScheduleSlotService(storage.ScheduleSlotStorage),
	}
}
