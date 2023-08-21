package auth_controller

import (
	"github.com/google/uuid"
)

type loginForm struct {
	Login    string    `json:"login" validate:"required,user-login-vld-tag"`
	Passwd   string    `json:"password" validate:"required"`
	DeviceID uuid.UUID `json:"device_id" validate:"omitempty"`
}

type entryPointForm struct {
	EntryID uuid.UUID `param:"entry_id" validate:"required"`
}

type onetimeAuthForm struct {
	UserID uint32 `json:"user_id" validate:"required"`
}
