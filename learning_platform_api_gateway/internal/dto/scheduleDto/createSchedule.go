package scheduleDto

import (
	"time"
)

type CreateSchedule struct {
	TutorID   int64                `json:"tutor_id"`
	StartTime time.Time            `json:"start_time"`
	EndTime   time.Time            `json:"end_time"`
	Slots     []CreateScheduleSlot `json:"slots"`
}
