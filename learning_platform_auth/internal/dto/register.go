package dto

type RegisterRequest struct {
	UserId   int64   `json:"user_id"`
	Email    string  `json:"email"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	LastName *string `json:"last_name"`
	Role     string  `json:"role"`
	Password string  `json:"password"`
}

type RegisterResponse struct {
	SessionId string
	UserId    int64
}
