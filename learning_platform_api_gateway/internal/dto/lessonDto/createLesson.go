package lessonDto

import (
	"time"
)

type CreateLesson struct {
	BoardID    *int64      `json:"board_id"`
	MeetLink   *string     `json:"meet_link"`
	StartTime  time.Time   `json:"start_time"`
	Duration   int64       `json:"duration"`
	TutorID    int64       `json:"tutor_id"`
	MediaItems []MediaItem `json:"media_items"`
	UserIDs    []int64     `json:"user_ids"`
}
