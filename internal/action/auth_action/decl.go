package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/internal/provider/auth_provider"
	"github.com/semenovem/portal/internal/provider/people_provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
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

var (
	errUserNoFound          = newAuthErr("user no found")
	errPasswdIncorrect      = newAuthErr("password is incorrect")
	errUserNotWorks         = newAuthErr("user not works")
	errSessionNotFound      = newAuthErr("auth session not found")
	errRefreshUnknown       = newAuthErr("refresh token data does not match")
	errOnetimeEntryNotFound = newAuthErr("onetime entry not found")
)

type AuthErr struct {
	msg string
}

func (e AuthErr) Error() string {
	return e.msg
}

func IsAuthErr(err error) bool {
	_, ok := err.(*AuthErr)
	return ok
}

func newAuthErr(msg string) *AuthErr {
	return &AuthErr{
		msg: msg,
	}
}

func (a *AuthAction) canLogin(user *it.UserAuth) error {
	if err := user.CanLogging(); err != nil {
		s := errUserNotWorks.msg + "(" + err.Error() + ")"
		a.logger.Named("canLogin").With("user", user).Debug(s)

		return errUserNotWorks
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
			ll.Debug(errSessionNotFound.msg)
			return nil, errSessionNotFound
		}

		ll.Named("GetSession").Nested(err.Error())
		return nil, err
	}

	if session.RefreshID != payload.RefreshID {
		ll.Named("refreshToken").
			With("refreshID_from_user", payload.RefreshID).
			With("refreshID_from_DB", session.RefreshID).
			AuthTag().Info(errRefreshUnknown.msg)

		return nil, errRefreshUnknown
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
		ll.Named("canLogin").Nested(err.Error())
		return nil, err
	}

	session, err := a.authPvd.CreateSession(ctx, user.ID, deviceID)
	if err != nil {
		ll.Named("CreateSession").Nested(err.Error())
		return nil, err
	}

	return session, nil
}
