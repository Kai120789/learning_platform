package dto

type FileMetadata struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	Size        uint64 `json:"size"`
}
