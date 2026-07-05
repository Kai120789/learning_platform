package lessonDto

import (
	"time"
)

type UpdateLesson struct {
	ID              int64       `json:"id"`
	BoardID         *int64      `json:"board_id"`
	MeetLink        *string     `json:"meet_link"`
	StartTime       time.Time   `json:"start_time"`
	Duration        int64       `json:"duration"`
	MediaItems      []MediaItem `json:"media_items"`
	UserIDs         []int64     `json:"user_ids"`
	DeletedUserIDs  []int64     `json:"deleted_user_ids"`
	DeletedMediaIDs []int64     `json:"deleted_media_ids"`
}
