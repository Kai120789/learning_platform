package dto

type RegisterRequest struct {
	UserId   int64
	Email    string
	Name     string
	Surname  string
	LastName string
	Role     string
	Password string
}

type RegisterResponse struct {
	SessionId string
	UserId    int64
}
