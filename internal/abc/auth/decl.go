package auth

import "github.com/google/uuid"

type Session struct {
	ID        uint32
	UserID    uint32
	DeviceID  uuid.UUID
	RefreshID uuid.UUID
}

func (a *Session) Reissue(refreshID uuid.UUID) *Session {
	return &Session{
		ID:        a.ID,
		UserID:    a.UserID,
		DeviceID:  a.DeviceID,
		RefreshID: refreshID,
	}
}
