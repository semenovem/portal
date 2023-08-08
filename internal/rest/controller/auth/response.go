package auth

import "github.com/semenovem/portal/internal/view"

type LoginResponse struct {
	AccessToken string               `json:"access_token"`
	Items       []*view.VehicleShort `json:"items"`
}
