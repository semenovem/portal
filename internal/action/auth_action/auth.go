package auth_action

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
)

// Logout разлогин авторизованной сессии пользователя
func (a *AuthAction) Logout(ctx context.Context, payload *jwtoken.RefreshPayload) (uint32, error) {
	ll := a.logger.Named("Logout").With("sessionID", payload.SessionID)

	session, err := a.getSessionByRefresh(ctx, payload)
	if err != nil {
		if IsAuthErr(err) {
			ll.Named("getSessionByRefresh").AuthTag().Info(err.Error())
			return 0, err
		}

		ll.Named("getSessionByRefresh").Nested(err.Error())
		return 0, err
	}

	if err = a.authPvd.LogoutSession(ctx, payload.SessionID); err != nil {
		ll.Named("LogoutSession").Nested(err.Error())
		return 0, err
	}

	ll.AuthTag().With("userID", session.UserID).Debug("success")

	return session.UserID, nil
}

// Refresh выпустить новый refresh токен и прекратить действие предыдущего
func (a *AuthAction) Refresh(
	ctx context.Context,
	payload *jwtoken.RefreshPayload,
) (*it.AuthSession, error) {
	ll := a.logger.Named("Refresh").With("sessionID", payload.SessionID)

	sessionOld, err := a.getSessionByRefresh(ctx, payload)
	if err != nil {
		ll.Named("getSessionByRefresh").Nested(err.Error())
		return nil, err
	}

	refreshID := uuid.New()

	// Проверить актуальность сотрудника
	userAuth, err := a.peoplePvd.GetUserAuth(ctx, sessionOld.UserID)
	if err != nil {
		ll.Named("GetUserAuth").Nested(err.Error())

		if provider.IsNoRows(err) {
			return nil, errUserNoFound
		}

		return nil, err
	}

	if err = a.canLogin(userAuth); err != nil {
		ll.Named("canLogin").Nested(err.Error())
		return nil, err
	}

	err = a.authPvd.UpdateRefreshSession(ctx, payload.SessionID, payload.RefreshID, refreshID)
	if err != nil {
		ll = ll.Named("UpdateRefreshSession")

		if provider.IsNoRows(err) {
			err = errors.New("authorized session could not be updated")

			ll.With("refreshID_old", payload.RefreshID).
				With("refreshID_new", refreshID).
				AuthTag().Info(err.Error())

			return nil, err
		}

		ll.Nested(err.Error())
		return nil, err
	}

	return sessionOld.Reissue(refreshID), nil
}
