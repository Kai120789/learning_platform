package materialDto

type MoveMaterials struct {
	MaterialIDs []int64 `json:"material_ids"`
	FolderID    *int64  `json:"folder_id"`
}
