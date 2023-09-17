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
	ll := a.logger.Named("Logout").With("sessionID", payload.SessionID)

	session, err := a.getSessionByRefresh(ctx, payload)
	if err != nil {
		ll.Named("getSessionByRefresh").Nested(err)
		return 0, err
	}

	if err = a.authPvd.LogoutSession(ctx, payload.SessionID); err != nil {
		ll.Named("LogoutSession").Nested(err)
		return 0, err
	}

	ll.With("userID", session.UserID).AuthDebugStr("user is logouted")

	return session.UserID, nil
}

// Refresh выпустить новый refresh токен и прекратить действие предыдущего
func (a *AuthAction) Refresh(
	ctx context.Context,
	payload *jwtoken.RefreshPayload,
) (*auth.Session, error) {
	ll := a.logger.Named("Refresh").With("sessionID", payload.SessionID)

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
		ll = ll.Named("UpdateRefreshSession")

		if throw.IsNotFoundErr(err) {
			err = throw.NewAuthErr("no auth session with the specified refresh token - could not be updated")

			ll.With("refreshID_old", payload.RefreshID).
				With("refreshID_new", refreshID).
				Auth(err)

			return nil, err
		}

		ll.Nested(err)
		return nil, err
	}

	return sessionOld.Reissue(refreshID), nil
}
