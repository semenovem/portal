package auth_provider

import (
	"context"
	"github.com/semenovem/portal/pkg/it"
)

func (p *AuthProvider) GetSession(ctx context.Context, sessionID uint32) (*it.AuthSession, error) {
	var (
		s  = it.AuthSession{ID: sessionID}
		sq = `SELECT user_id, device_id, refresh_id
			FROM auth.sessions
			WHERE logouted = false AND id = $1`
	)

	err := p.db.QueryRow(ctx, sq, sessionID).Scan(&s.UserID, &s.DeviceID, &s.RefreshID)
	if err != nil {
		p.logger.Named("GetSession").With("sessionID", sessionID).DB(err)
		return nil, err
	}

	return &s, nil
}
