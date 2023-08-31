package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/throw"
)

// NewLogin авторизация пользователя по логопасу
func (a *AuthAction) NewLogin(
	ctx context.Context,
	login, passwd string,
	deviceID uuid.UUID,
) (*it.AuthSession, error) {
	ll := a.logger.Named("NewLogin").With("login", login).With("deviceID", deviceID)

	userAuth, err := a.peoplePvd.GetUserByLogin(ctx, login)
	if err != nil {
		ll = ll.Named("GetUserByLogin")

		if provider.IsNoRows(err) {
			ll.Auth(throw.Err404User)
			return nil, throw.Err404User
		}

		ll.Nested(err)
		return nil, err
	}

	ll.With("userID", userAuth.ID)

	if !a.passwdAuth.Matching(userAuth.PasswdHash, passwd) {
		ll.Named("PasswordMatching").AuthDebug(throw.ErrAuthPasswdIncorrect)
		return nil, throw.ErrAuthPasswdIncorrect
	}

	session, err := a.newSession(ctx, userAuth.ToUserAuth(), deviceID)
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
