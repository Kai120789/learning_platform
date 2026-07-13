package dto

type TokenBundle struct {
	SessionId    string `json:"session_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
