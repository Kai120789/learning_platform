package dto

type CreateJWT struct {
	UserId      int64
	UserEmail   string
	SignedKey   string
	SessionId   *string
	Issuer      string
	AccessTime  int64
	RefreshTime int64
}
