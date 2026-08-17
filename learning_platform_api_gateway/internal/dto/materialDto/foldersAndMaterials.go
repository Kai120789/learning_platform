package materialDto

type FoldersAndMaterials struct {
	Folders   []MaterialFolder `json:"folders"`
	Materials []Material       `json:"materials"`
}
