package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
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
			ll.AuthStr(errUserNoFound.msg)
			return nil, errUserNoFound
		}

		ll.Nested(err)
		return nil, err
	}

	ll.With("userID", userAuth.ID)

	if !a.passwdAuth.Matching(userAuth.PasswdHash, passwd) {
		ll.Named("PasswordMatching").AuthDebugStr(errPasswdIncorrect.msg)
		return nil, errPasswdIncorrect
	}

	session, err := a.newSession(ctx, userAuth.ToUserAuth(), deviceID)
	if err != nil {
		ll = ll.Named("newSession")

		if IsAuthErr(err) {
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
