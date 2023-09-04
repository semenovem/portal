package vehicle_controller

import (
	"github.com/semenovem/portal/internal/rest/view"
)

type ListResponse struct {
	Total uint32               `json:"total"`
	Items []*view.VehicleShort `json:"items"`
}

type vehicleView struct {
}
