package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"lastname"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	SessionId string `json:"session_id"`
	UserId    int64  `json:"user_id"`
}
