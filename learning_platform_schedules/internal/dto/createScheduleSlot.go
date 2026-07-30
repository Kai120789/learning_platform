package dto

import (
	"time"
)

type CreateScheduleSlot struct {
	ScheduleID int64     `json:"schedule_id"`
	StartTime  time.Time `json:"start_time"`
	Duration   *int64    `json:"duration"`
}
