package enum

type UserRole string

const (
	RoleAdmin   UserRole = "ADMIN"
	RoleTutor   UserRole = "TUTOR"
	RoleStudent UserRole = "STUDENT"
)
