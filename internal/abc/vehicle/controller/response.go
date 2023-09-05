package vehicle_controller

type ListResponse struct {
	Total uint32          `json:"total"`
	Items []*VehicleShort `json:"items"`
}

type vehicleView struct {
}

type VehicleShort struct {
	ID   uint32 `json:"id"`
	Name uint32 `json:"name"`
}
