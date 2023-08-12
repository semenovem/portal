package auth

import (
	"github.com/google/uuid"
)

type LoginForm struct {
	Login    string    `json:"login" validate:"required,user-login-vld-tag"`
	Password string    `json:"password" validate:"required,user-password-vld-tag"`
	DeviceID uuid.UUID `json:"device_id" validate:"omitempty"`
}
