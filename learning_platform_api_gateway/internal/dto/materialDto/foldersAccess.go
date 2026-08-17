package materialDto

type FoldersAccess struct {
	FolderIDs []int64 `json:"folder_ids"`
	UserIDs   []int64 `json:"user_ids"`
}
