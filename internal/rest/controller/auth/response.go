package auth

import "github.com/semenovem/portal/internal/view"

type ListResponse struct {
	Total uint32               `json:"total"`
	Items []*view.VehicleShort `json:"items"`
}
