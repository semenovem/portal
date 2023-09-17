package people_controller

type userPathForm struct {
	UserID uint32 `json:"user_id" param:"user_id" validate:"required"`
}

type userCreateForm struct {
	Firstname  *string   `json:"firstname" validate:"omitempty,user-name-vld-tag"`
	Surname    *string   `json:"surname" validate:"omitempty,user-name-vld-tag"`
	Patronymic *string   `json:"patronymic" validate:"omitempty,user-name-vld-tag"`
	Note       *string   `json:"note"`
	Status     *string   `json:"status" validate:"omitempty,user-status-vld-tag"`
	Groups     *[]string `json:"groups" validate:"omitempty,user-groups-vld-tag"`
	AvatarID   *uint32   `json:"avatar_id"`
	Login      *string   `json:"login" validate:"omitempty,user-login-vld-tag"`
	Passwd     string    `json:"passwd" validate:"omitempty,user-password-vld-tag"`
	ExpiredAt  *string   `json:"expired_at" validate:"omitempty,time-vld-tag"`
}

type employeeCreateForm struct {
	userCreateForm
	PositionID *uint16 `json:"position_id" validate:"omitempty"`
	DeptID     *uint16 `json:"dept_id" validate:"omitempty"`
	WorkedAt   *string `json:"worked_at" validate:"omitempty,time-vld-tag"`
	FiredAt    *string `json:"fired_at" validate:"omitempty,time-vld-tag"`
}

type employeeUpdateForm struct {
	UserID uint32 `json:"user_id" param:"user_id" validate:"required" swaggerignore:"true"`
	employeeCreateForm
}

type freeLoginNameForm struct {
	LoginName string `param:"login_name" validate:"required"`
}

type handbookForm struct {
	DeptID     []uint16 `json:"dept_id[]" query:"dept_id[]" validate:"omitempty"`
	PositionID []uint16 `json:"position_id[]" query:"position_id[]" validate:"omitempty"`
	Order      []string `json:"order[]" query:"order[]" validate:"omitempty"`
}
