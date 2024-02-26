package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/auth"
	"github.com/semenovem/portal/pkg/jwtoken"
	"github.com/semenovem/portal/pkg/throw"
)

// Logout разлогин авторизованной сессии пользователя
func (a *AuthAction) Logout(ctx context.Context, payload *jwtoken.RefreshPayload) (uint32, error) {
	ll := a.logger.Func(ctx, "Logout").With("sessionID", payload.SessionID)

	session, err := a.getSessionByRefresh(ctx, payload)
	if err != nil {
		ll.Named("getSessionByRefresh").Nested(err)
		return 0, err
	}

	if err = a.authPvd.LogoutSession(ctx, payload.SessionID); err != nil {
		ll.Named("LogoutSession").Nested(err)
		return 0, err
	}

	return session.UserID, nil
}

// Refresh выпустить новый refresh токен и прекратить действие предыдущего
func (a *AuthAction) Refresh(
	ctx context.Context,
	payload *jwtoken.RefreshPayload,
) (*auth.Session, error) {
	const label = "AuthAction.Refresh"
	ll := a.logger.Func(ctx, "Refresh").With("sessionID", payload.SessionID)

	sessionOld, err := a.getSessionByRefresh(ctx, payload)
	if err != nil {
		ll.Named("getSessionByRefresh").Nested(err)
		return nil, err
	}

	refreshID := uuid.New()

	// Проверить актуальность сотрудника
	userAuth, err := a.peoplePvd.GetUserAuth(ctx, sessionOld.UserID)
	if err != nil {
		ll.Named("GetUserAuth").Nested(err)
		return nil, err
	}

	if err = a.canLogin(userAuth); err != nil {
		ll.Named("canLogin").Nested(err)
		return nil, err
	}

	err = a.authPvd.UpdateRefreshSession(ctx, payload.SessionID, payload.RefreshID, refreshID)
	if err != nil {
		if throw.IsNotFound(err) {
			err = throw.NewAuth("no auth session with the specified refresh token - could not be updated")
		}

		return nil, throw.Trace(err, label, map[string]any{
			"payload": payload,
		})
	}

	return sessionOld.Reissue(refreshID), nil
}
