package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
	"github.com/semenovem/portal/pkg/throw"
)

type AuthAction struct {
	logger     pkg.Logger
	authPvd    *auth_provider.AuthProvider
	peoplePvd  *people_provider.PeopleProvider
	passwdAuth it.UserPasswdAuthenticator
}

func New(
	logger pkg.Logger,
	passwdAuth it.UserPasswdAuthenticator,
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,
) *AuthAction {
	return &AuthAction{
		logger:     logger.Named("AuthAction"),
		passwdAuth: passwdAuth,
		authPvd:    authPvd,
		peoplePvd:  peoplePvd,
	}
}

func (a *AuthAction) canLogin(user *it.UserAuth) error {
	if err := user.CanLogging(); err != nil {
		err = throw.NewAuthErr(throw.MsgUserCantLogin, err.Error())
		a.logger.Named("canLogin").With("user", user).Auth(err)

		return err
	}

	return nil
}

// getSessionByRefresh Получить авторизованную сессию по данным из refresh токена
func (a *AuthAction) getSessionByRefresh(
	ctx context.Context,
	payload *jwtoken.RefreshPayload,
) (*it.AuthSession, error) {
	ll := a.logger.Named("getSessionByRefresh").With("sessionID", payload.SessionID)

	session, err := a.authPvd.GetSession(ctx, payload.SessionID)
	if err != nil {
		ll = ll.Named("GetSession")

		if provider.IsNoRows(err) {
			ll.Auth(throw.Err404AuthSession)
			return nil, throw.Err404AuthSession
		}

		ll.Nested(err)
		return nil, err
	}

	if session.RefreshID != payload.RefreshID {
		ll.Named("refreshToken").
			With("refreshID_from_user", payload.RefreshID).
			With("refreshID_from_DB", session.RefreshID).
			Auth(throw.ErrAuthRefreshUnknown)

		return nil, throw.ErrAuthRefreshUnknown
	}

	return session, nil
}

// Создание новой авторизованной сессии
func (a *AuthAction) newSession(
	ctx context.Context,
	user *it.UserAuth,
	deviceID uuid.UUID,
) (*it.AuthSession, error) {
	ll := a.logger.Named("newSession")

	if err := a.canLogin(user); err != nil {
		ll.Named("canLogin").Nested(err)
		return nil, err
	}

	session, err := a.authPvd.CreateSession(ctx, user.ID, deviceID)
	if err != nil {
		ll.Named("CreateSession").Nested(err)
		return nil, err
	}

	return session, nil
}
