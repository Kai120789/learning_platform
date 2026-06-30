package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/schedules/internal/models/enum"
)

type ScheduleSlot struct {
	ID         int64               `json:"id"`
	ScheduleID int64               `json:"schedule_id"`
	StartTime  pgtype.Timestamptz  `json:"start_time"`
	Status     enum.ScheduleStatus `json:"status"`
	Duration   pgtype.Int8         `json:"duration"`
	LessonID   pgtype.Int8         `json:"lesson_id"`
	CreatedAt  pgtype.Timestamptz  `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz  `json:"updated_at"`
}
