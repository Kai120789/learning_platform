package authDto

type CreateJWT struct {
	UserID      int64  `json:"user_id"`
	UserEmail   string `json:"user_email"`
	SignedKey   string `json:"signed_key"`
	Issuer      string `json:"issuer"`
	AccessTime  int64  `json:"access_time"`
	RefreshTime int64  `json:"refresh_time"`
}
