package provider

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/pkg/it"
)

func (p *AuthPvd) CreateSession(
	ctx context.Context,
	userID uint32,
	deviceID uuid.UUID,
) (*it.AuthSession, error) {
	var (
		ll        = p.logger.Named("CreateSession")
		refreshID = uuid.New()
		sessionID uint32

		sq = `INSERT INTO auth.sessions (user_id, device_id, refresh_id) VALUES ($1, $2, $3)
			   RETURNING id;`
	)

	if err := p.db.QueryRow(ctx, sq, userID, deviceID, refreshID).Scan(&sessionID); err != nil {
		ll.DBTag().Named("QueryRow").With("userID", userID).Error(err.Error())
		return nil, err
	}

	return &it.AuthSession{
		ID:        sessionID,
		UserID:    userID,
		DeviceID:  deviceID,
		RefreshID: refreshID,
	}, nil
}

//func (p *AuthPvd) CreateSession(
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
