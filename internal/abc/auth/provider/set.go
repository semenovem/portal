package auth_provider

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/auth"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
)

func (p *AuthProvider) CreateSession(
	ctx context.Context,
	userID uint32,
	deviceID uuid.UUID,
) (*auth.Session, error) {
	var (
		refreshID = uuid.New()
		sessionID uint32

		sq = `INSERT INTO auth.sessions (user_id, device_id, refresh_id) VALUES ($1, $2, $3)
			   RETURNING id;`
	)

	if err := p.db.QueryRow(ctx, sq, userID, deviceID, refreshID).Scan(&sessionID); err != nil {
		p.logger.Func(ctx, "CreateSession.QueryRow").With("userID", userID).DB(err)
		return nil, err
	}

	return &auth.Session{
		ID:        sessionID,
		UserID:    userID,
		DeviceID:  deviceID,
		RefreshID: refreshID,
	}, nil
}

func (p *AuthProvider) LogoutSession(ctx context.Context, sessionID uint32) error {
	var (
		ll = p.logger.Func(ctx, "LogoutSession").With("sessionID", sessionID)
		sq = `UPDATE auth.sessions SET logouted = true WHERE logouted = false AND id = $1`
	)

	if result, err := p.db.Exec(ctx, sq, sessionID); err != nil {
		ll.Named("Exec").DB(err)
		return err
	} else if result.RowsAffected() == 0 {
		ll.AuthStr("auth session id not found (possible already logouted)")
		return throw.Err404AuthSession
	}

	if err := p.sessionCanceled(ctx, sessionID); err != nil {
		ll.Named("sessionCanceled").Nested(err)
		return err
	}

	return nil
}

func (p *AuthProvider) UpdateRefreshSession(
	ctx context.Context,
	sessionID uint32,
	refreshOldID uuid.UUID,
	refreshNewID uuid.UUID,
) error {
	sq := `UPDATE auth.sessions SET refresh_id = $3
	WHERE logouted = false AND refresh_id = $2 AND id = $1`

	result, err := p.db.Exec(ctx, sq, sessionID, refreshOldID, refreshNewID)
	if err != nil {
		p.logger.Func(ctx, "UpdateRefreshSession").
			With("sessionID", sessionID).
			With("refreshOldID", refreshOldID).
			With("refreshNewID", refreshNewID).DB(err)

		return err
	}

	if result.RowsAffected() == 0 {
		return throw.Err404AuthSession
	}

	return nil
}

func (p *AuthProvider) NewOnetimeEntry(
	ctx context.Context,
	userID uint32,
) (entryID uuid.UUID, err error) {
	entryID = uuid.New()

	err = p.redis.Set(ctx, getOnetimeEntryKeyName(entryID), userID, p.config.OnetimeEntryLifetime).Err()
	if err != nil {
		p.logger.Func(ctx, "NewOnetimeEntry").With("onetimeEntryID", entryID).Error(err.Error())
	}

	return entryID, err
}

func (p *AuthProvider) GetDelOnetimeEntry(
	ctx context.Context,
	entryID uuid.UUID,
) (userID uint32, err error) {
	v, err := p.redis.GetDel(ctx, getOnetimeEntryKeyName(entryID)).Uint64()
	if err != nil {
		ll := p.logger.Func(ctx, "GetDelOnetimeEntry").With("onetimeEntryID", entryID)

		if provider.IsNoRec(err) {
			ll.Debug(provider.MsgErrNoRecordRedis)
		} else {
			ll.Error(err.Error())
		}
	}

	return uint32(v), err
}
