package auth

type refreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"` // TODO для разработки
}

type loginResponse struct {
	refreshResponse
	UserID uint32 `json:"user_id"`
}

type onetimeAuthResponse struct {
	URI     string `json:"uri"`
	EntryID string `json:"entry_id"`
}
