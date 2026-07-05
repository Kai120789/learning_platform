package scheduleDto

import (
	"learning-platform/api-gateway/internal/dto/enum"
	"time"
)

type ScheduleSlot struct {
	ID         int64               `json:"id"`
	ScheduleID int64               `json:"schedule_id"`
	StartTime  time.Time           `json:"start_time"`
	Status     enum.ScheduleStatus `json:"status"`
	Duration   *int64              `json:"duration"`
	LessonID   *int64              `json:"lesson_id"`
}
