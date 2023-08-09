package auth

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint32 `json:"user_id"`
}
