package materialDto

type MoveFolders struct {
	FolderIDs      []int64 `json:"folder_ids"`
	ParentFolderID *int64  `json:"parent_folder_id"`
}
