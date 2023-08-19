package auth_provider

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/pkg/it"
)

func (p *AuthProvider) CreateSession(
	ctx context.Context,
	userID uint32,
	deviceID uuid.UUID,
) (*it.AuthSession, error) {
	var (
		refreshID = uuid.New()
		sessionID uint32

		sq = `INSERT INTO auth.sessions (user_id, device_id, refresh_id) VALUES ($1, $2, $3)
			   RETURNING id;`
	)

	if err := p.db.QueryRow(ctx, sq, userID, deviceID, refreshID).Scan(&sessionID); err != nil {
		p.logger.DBTag().Named("CreateSession.QueryRow").
			With("userID", userID).Error(err.Error())
		return nil, err
	}

	return &it.AuthSession{
		ID:        sessionID,
		UserID:    userID,
		DeviceID:  deviceID,
		RefreshID: refreshID,
	}, nil
}

func (p *AuthProvider) LogoutSession(ctx context.Context, sessionID uint32) error {
	var (
		ll = p.logger.Named("LogoutSession").With("sessionID", sessionID)
		sq = `UPDATE auth.sessions SET logouted = true WHERE logouted = false AND id = $1`
	)

	if result, err := p.db.Exec(ctx, sq, sessionID); err != nil {
		ll.Named("Exec").DBTag().Error(err.Error())
		return err
	} else if result.RowsAffected() == 0 {
		ll.Named("Exec").AuthTag().Info("auth session id (not logouted) not found")
	}

	if err := p.sessionCanceled(ctx, sessionID); err != nil {
		ll.Named("sessionCanceled").Nested(err.Error())
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
	var (
		ll = p.logger.Named("UpdateRefreshSession").
			With("sessionID", sessionID).
			With("refreshOldID", refreshOldID).
			With("refreshNewID", refreshNewID)

		sq = `UPDATE auth.sessions SET refresh_id = $3
				WHERE logouted = false AND refresh_id = $2 AND id = $1`
	)

	result, err := p.db.Exec(ctx, sq, sessionID, refreshOldID, refreshNewID)
	if err != nil {
		ll.Named("Exec").DBTag().Error(err.Error())
		return err
	}

	if result.RowsAffected() == 0 {
		ll.Named("RowsAffected").Debug(err.Error())
		return pgx.ErrNoRows
	}

	return nil
}

//func (p *AuthProvider) CreateSession(
//	ctx context.Context,
//	userID uint32,
//	deviceID uuid.UUID,
//) (*it.AuthSession, error) {
//	var (
//		ll        = p.logger.Named("CreateSession")
//		refreshID = uuid.New()
//		sessionID uint32
//
//		sq = `INSERT INTO auth.sessions (user_id, device_id, refresh_id) VALUES ($1, $2, $3)
//			   RETURNING id;`
//	)
//
//	tx, err := p.db.Begin(ctx)
//	if err != nil {
//		ll.DBTag().Named("Begin").Error(err.Error())
//		return nil, err
//	}
//
//	defer func() {
//		if err1 := tx.Rollback(ctx); err1 != nil && err1 != pgx.ErrTxClosed {
//			ll.DBTag().Named("Rollback").Error(err1.Error())
//		}
//	}()
//
//	if err = tx.QueryRow(ctx, sq, userID, deviceID, refreshID).Scan(&sessionID); err != nil {
//		ll.DBTag().Named("QueryRow").With("userID", userID).Error(err.Error())
//		return nil, err
//	}
//
//	if _, err = tx.Exec(ctx, sq2, sessionID, userAgent); err != nil {
//		ll.DBTag().Named("Exec").Error(err.Error())
//		return nil, err
//	}
//
//	if err = tx.Commit(ctx); err != nil {
//		ll.DBTag().Named("Commit").Error(err.Error())
//		return nil, err
//	}
//
//	return &it.AuthSession{
//		ID:        sessionID,
//		UserID:    userID,
//		DeviceID:  deviceID,
//		RefreshID: refreshID,
//	}, nil
//}
