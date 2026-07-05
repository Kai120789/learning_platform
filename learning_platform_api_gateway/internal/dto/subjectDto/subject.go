package subjectDto

import (
	"learning-platform/api-gateway/internal/dto/enum"
)

type Subject struct {
	ID    int64     `json:"id"`
	Code  string    `json:"code"`
	Title string    `json:"title"`
	Type  enum.Type `json:"type"`
}
