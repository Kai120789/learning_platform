package dto

type UsersWithCount struct {
	Users []UserShortInfo `json:"users"`
	Count int64           `json:"count"`
}
