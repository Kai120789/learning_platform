package dto

type CreateFolder struct {
	Title          string `json:"title"`
	ParentFolderID *int64 `json:"parent_folder_id"`
	TutorID        int64  `json:"tutor_id"`
}
