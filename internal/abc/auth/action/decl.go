package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/auth"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/jwtoken"
	"github.com/semenovem/portal/pkg/throw"
)

type AuthAction struct {
	logger    pkg.Logger
	authPvd   *auth_provider.AuthProvider
	peoplePvd *people_provider.PeopleProvider
}

func New(
	logger pkg.Logger,
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,
) *AuthAction {
	return &AuthAction{
		logger:    logger.Named("AuthAction"),
		authPvd:   authPvd,
		peoplePvd: peoplePvd,
	}
}

func (a *AuthAction) canLogin(user *people.UserAuth) error {
	if err := user.CanLogging(); err != nil {
		err = throw.NewWithTargetErr(throw.ErrAuthUserCannotLogin, err)
		a.logger.Named("canLogin").With("user", user).Auth(err)
		return err
	}

	return nil
}

// getSessionByRefresh Получить авторизованную сессию по данным из refresh токена
func (a *AuthAction) getSessionByRefresh(
	ctx context.Context,
	payload *jwtoken.RefreshPayload,
) (*auth.Session, error) {
	ll := a.logger.Named("getSessionByRefresh").With("sessionID", payload.SessionID)

	session, err := a.authPvd.GetSession(ctx, payload.SessionID)
	if err != nil {
		ll.Named("GetSession").Nested(err)
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
	user *people.UserAuth,
	deviceID uuid.UUID,
) (*auth.Session, error) {
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
