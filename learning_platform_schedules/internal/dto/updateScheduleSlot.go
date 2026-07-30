package dto

import (
	"time"
)

type UpdateScheduleSlot struct {
	StartTime time.Time `json:"start_time"`
	Duration  *int64    `json:"duration"`
}
