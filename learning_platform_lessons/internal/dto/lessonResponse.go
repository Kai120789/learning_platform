package dto

import (
	"learning-platform/lessons/internal/models/enum"
	"time"
)

type LessonResponse struct {
	ID         int64               `json:"id"`
	BoardID    *int64              `json:"board_id"`
	MeetLink   *string             `json:"meet_link"`
	StartTime  time.Time           `json:"start_time"`
	Duration   int64               `json:"duration"`
	TutorID    int64               `json:"tutor_id"`
	Status     enum.LessonStatus   `json:"status"`
	UserIDs    []int64             `json:"user_ids"`
	MediaItems []MediaItemResponse `json:"media_items"`
}
