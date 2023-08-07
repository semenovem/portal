package view

type VehicleShort struct {
	ID   uint32 `json:"id"`
	Name uint32 `json:"name"`
}

type Vehicle struct {
	VehicleShort
}
