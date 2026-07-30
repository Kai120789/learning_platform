package scheduleDto

import "time"

type UpdateSchedule struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
