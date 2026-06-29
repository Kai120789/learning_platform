package dto

import "learning-platform/lessons/internal/models/enum"

type MediaItem struct {
	S3Link    string         `json:"s3_link"`
	S3Preview string         `json:"s3_preview"`
	Type      enum.MediaType `json:"type"`
}
