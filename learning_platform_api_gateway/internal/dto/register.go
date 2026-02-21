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
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
}
