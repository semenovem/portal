package auth

import (
	"github.com/google/uuid"
)

type LoginForm struct {
	Login    string    `json:"login" validate:"required,user-login-vld-tag"`
	Passwd   string    `json:"password" validate:"required"`
	DeviceID uuid.UUID `json:"device_id" validate:"omitempty"`
}
