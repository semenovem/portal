package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/auth"
	"github.com/semenovem/portal/pkg/throw"
)

// NewLogin авторизация пользователя по логопасу
func (a *AuthAction) NewLogin(
	ctx context.Context,
	login, passwdHash string,
	deviceID uuid.UUID,
) (*auth.Session, error) {
	ll := a.logger.Named("NewLogin").With("login", login).With("deviceID", deviceID)

	userAuth, err := a.peoplePvd.GetUserByLogin(ctx, login, passwdHash)
	if err != nil {
		ll.Named("GetUserByLogin").Nested(err)
		return nil, err
	}

	ll.With("userID", userAuth.ID)

	session, err := a.newSession(ctx, userAuth, deviceID)
	if err != nil {
		ll = ll.Named("newSession")

		if throw.IsAuthErr(err) {
			ll.Auth(err)
			return nil, err
		}

		ll.Nested(err)
		return nil, err
	}

	ll.With("sessionID", session.ID).With("refreshID", session.RefreshID).
		AuthDebugStr("user is logged")

	return session, nil
}
