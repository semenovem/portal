package people_controller

type userCreateResponse struct {
	UserID uint32 `json:"user_id"`
}

type freeLoginNameResponse struct {
	Free        bool   `json:"free"`
	ValidateErr string `json:"validate_err,omitempty"`
}

type publicHandbookResponse struct {
	Employees []*publicEmployeeView `json:"employees"`
}
