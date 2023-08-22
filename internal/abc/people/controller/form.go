package people_controller

import "time"

type UserForm struct {
	UserID uint32 `param:"user_id" validate:"required"`
}

type CreateUserForm struct {
	FirstName string
	Surname   string
	Note      string
	Position  string
	Status    string
	Roles     []string
	ExpiredAt time.Time `json:"expired_at" binding:"datetime=2006-01-02T15:04:05Z07:00" validate:"omitempty"`
	Login     string
	Passwd    string
}
