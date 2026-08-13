package dto

type CreateMaterial struct {
	Title         string `json:"title"`
	Size          int64  `json:"size"`
	FolderID      *int64 `json:"folder_id"`
	TutorID       int64  `json:"tutor_id"`
	MimeType      string `json:"mime_type"`
	MediaObjectID string `json:"media_object_id"`
}
