package materialDto

type MaterialFolder struct {
	ID             int64  `json:"id"`
	Title          string `json:"title"`
	ParentFolderID *int64 `json:"parent_folder_id"`
	TutorID        int64  `json:"tutor_id"`
}
