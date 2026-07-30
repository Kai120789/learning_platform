package scheduleDto

import (
	"time"
)

type CreateSchedule struct {
	Title     string    `json:"title"`
	TutorID   int64     `json:"tutor_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
