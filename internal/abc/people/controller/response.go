package people_controller

type userCreateResponse struct {
	UserID uint32 `json:"user_id"`
}

type freeLoginNameResponse struct {
	Free        bool   `json:"free"`
	ValidateErr string `json:"validate_err,omitempty"` // Что не так с введенным логином
}

type employeeHandbookResponse struct {
	Total     uint32                 `json:"total"`
	Employees []*employeeProfileView `json:"employees"`
}
