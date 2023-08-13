package it

import (
	"github.com/google/uuid"
)

type AuthSession struct {
	ID        uint32
	UserID    uint32
	DeviceID  uuid.UUID
	RefreshID uuid.UUID
}
