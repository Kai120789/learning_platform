package lessonDto

import "learning-platform/api-gateway/internal/dto/enum"

type MediaItemResponse struct {
	ID        int64          `json:"id"`
	LessonID  int64          `json:"lesson_id"`
	S3Link    string         `json:"s3_link"`
	S3Preview string         `json:"s3_preview"`
	Type      enum.MediaType `json:"type"`
}
