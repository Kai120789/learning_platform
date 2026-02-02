package dto

type RegisterRequest struct {
	Email    string
	Name     string
	Surname  string
	LastName string
	Role     string
	Password string
}

type RegisterResponse struct {
	AccessToken string
	UserId      int64
}
