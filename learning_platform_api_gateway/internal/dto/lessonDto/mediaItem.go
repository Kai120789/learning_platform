package lessonDto

import "learning-platform/api-gateway/internal/dto/enum"

type MediaItem struct {
	S3Link    string         `json:"s3_link"`
	S3Preview string         `json:"s3_preview"`
	Type      enum.MediaType `json:"type"`
}
