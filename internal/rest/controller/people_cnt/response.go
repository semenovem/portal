package people_cnt

import (
	"github.com/semenovem/portal/internal/rest/view"
)

// ProfilePublic общедоступный профиль пользователя
type ProfilePublic struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

// ProfileFull полный профиль пользователя
type ProfileFull struct {
	ProfilePublic
}

type ListResponse struct {
	Total uint32               `json:"total"`
	Items []*view.VehicleShort `json:"items"`
}
