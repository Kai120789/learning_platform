package scheduleDto

import "time"

type UpdateSchedule struct {
	ID                    int64                `json:"id"`
	StartTime             time.Time            `json:"start_time"`
	EndTime               time.Time            `json:"end_time"`
	Slots                 []CreateScheduleSlot `json:"slots"`
	DeleteScheduleSlotIDs []int64              `json:"delete_schedule_slot_ids"`
}
