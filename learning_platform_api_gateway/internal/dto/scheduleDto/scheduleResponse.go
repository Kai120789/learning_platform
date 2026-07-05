package scheduleDto

import "time"

type ScheduleResponse struct {
	ID        int64          `json:"id"`
	TutorID   int64          `json:"tutor_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	Slots     []ScheduleSlot `json:"slots"`
}
