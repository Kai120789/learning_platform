package dto

import (
	"time"
)

type CreateScheduleSlot struct {
	StartTime time.Time `json:"start_time"`
	Duration  *int64    `json:"duration"`
	LessonID  *int64    `json:"lesson_id"`
}
