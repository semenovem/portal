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

func (a *AuthSession) Reissue(refreshID uuid.UUID) *AuthSession {
	return &AuthSession{
		ID:        a.ID,
		UserID:    a.UserID,
		DeviceID:  a.DeviceID,
		RefreshID: refreshID,
	}
}
