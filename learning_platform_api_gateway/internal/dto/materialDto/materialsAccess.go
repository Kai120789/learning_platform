package materialDto

type MaterialsAccess struct {
	MaterialIDs []int64 `json:"material_ids"`
	UserIDs     []int64 `json:"user_ids"`
}
