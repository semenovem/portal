package auth

import (
	"github.com/google/uuid"
)

type LoginForm struct {
	Login    string    `json:"login" validate:"required,UserLoginVLDTag"`
	Password string    `json:"password" validate:"required,UserPasswordVLDTag"`
	DeviceID uuid.UUID `json:"device_id" validate:"omitempty"`
}
