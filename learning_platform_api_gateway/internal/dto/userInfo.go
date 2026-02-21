package dto

type UserInfo struct {
	UserId   int64
	Name     string
	Surname  string
	Lastname *string
	City     *string
	About    *string
	Role     string
	Status   string
	Class    *int64
}
