package auth_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/auth"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
)

func (p *AuthProvider) GetSession(ctx context.Context, sessionID uint32) (*auth.Session, error) {
	var (
		s  = auth.Session{ID: sessionID}
		sq = `SELECT user_id, device_id, refresh_id
			FROM auth.sessions
			WHERE logouted = false AND id = $1`
	)

	err := p.db.QueryRow(ctx, sq, sessionID).Scan(&s.UserID, &s.DeviceID, &s.RefreshID)
	if err != nil {
		if provider.IsNoRow(err) {
			return nil, throw.Err404AuthSession
		}

		p.logger.Func(ctx, "GetSession").With("sessionID", sessionID).DB(err)
		return nil, err
	}

	return &s, nil
}
