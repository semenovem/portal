package auth_action

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
	"github.com/semenovem/portal/pkg/logger"
)

// NewLogin авторизация пользователя по логопасу
func (a *AuthAction) NewLogin(
	ctx context.Context,
	login, passwd string,
	deviceID uuid.UUID,
) (*it.AuthSession, error) {
	var (
		ll = a.logger.Named("NewLogin").
			With("login", login).
			With("deviceID", deviceID)
	)

	loggingUser, err := a.peoplePvd.GetUserByLogin(ctx, login)
	if err != nil {
		ll = ll.Named("GetUserByLogin")

		if provider.IsNoRows(err) {
			ll.Tags(logger.AuthTag, logger.ClientTag).Info(errUserNoFound.msg)
			return nil, errUserNoFound
		}

		ll.Nested(err.Error())
		return nil, err
	}

	ll.With("userID", loggingUser.ID)

	if !a.passwdAuth.Matching(loggingUser.PasswdHash, passwd) {
		ll.Named("Matching").AuthTag().ClientTag().Debug(errPasswdIncorrect.msg)
		return nil, errPasswdIncorrect
	}

	session, err := a.newSession(ctx, loggingUser.ToUser(), deviceID)
	if err != nil {
		ll = ll.Named("newSession")

		if IsAuthErr(err) {
			ll.AuthTag().Info(err.Error())
			return nil, err
		}

		ll.Nested(err.Error())
		return nil, err
	}

	ll.AuthTag().With("sessionID", session.ID).With("refreshID", session.RefreshID).
		Debug("success")

	return session, nil
}

// NewSession Создание новой авторизованной сессии
//func (a *AuthAction) NewSession(
//	ctx context.Context,
//	userID uint32,
//	deviceID uuid.UUID,
//) (*it.AuthSession, error) {
//	ll := a.logger.Named("NewSession")
//
//	user, err := a.peoplePvd.GetUser(ctx, userID)
//	if err != nil {
//		ll.Named("GetUser").Nested(err.Error())
//
//		if provider.IsNoRows(err) {
//			return nil, errUserNoFound
//		}
//
//		return nil, err
//	}
//
//	session, err := a.newSession(ctx, user, deviceID)
//	if err != nil {
//		ll.Named("newSession").Nested(err.Error())
//		return nil, err
//	}
//
//	return session, nil
//}

func (a *AuthAction) canAuth(user *it.User) error {
	if err := user.IsWorks(); err != nil {
		s := errUserNotWorks.msg + "(" + err.Error() + ")"
		a.logger.Named("canAuth").With("user", user).Debug(s)

		return errUserNotWorks
	}

	return nil
}

// Создание новой авторизованной сессии
func (a *AuthAction) newSession(
	ctx context.Context,
	user *it.User,
	deviceID uuid.UUID,
) (*it.AuthSession, error) {
	ll := a.logger.Named("newSession")

	if err := a.canAuth(user); err != nil {
		ll.Named("canAuth").Nested(err.Error())
		return nil, err
	}

	session, err := a.authPvd.CreateSession(ctx, user.ID, deviceID)
	if err != nil {
		ll.Named("CreateSession").Nested(err.Error())
		return nil, err
	}

	return session, nil
}

// checkRefresh Проверить актуальность refresh токена
func (a *AuthAction) checkRefresh(
	ctx context.Context,
	payload *jwtoken.RefreshPayload,
) (*it.AuthSession, error) {
	ll := a.logger.Named("checkRefresh").With("sessionID", payload.SessionID)

	session, err := a.authPvd.GetSession(ctx, payload.SessionID)
	if err != nil {
		ll = ll.Named("GetSession")

		if provider.IsNoRows(err) {
			ll.Debug(sessionNotFoundErrMsg.msg)
			return nil, sessionNotFoundErrMsg
		}

		ll.Named("GetSession").Nested(err.Error())
		return nil, err
	}

	if session.RefreshID != payload.RefreshID {
		ll.Named("refreshToken").
			With("refreshID_from_user", payload.RefreshID).
			With("refreshID_from_DB", session.RefreshID).
			AuthTag().Info(refreshUnknown.msg)

		return nil, refreshUnknown
	}

	return session, nil
}

// Logout разлогин авторизованной сессии пользователя
func (a *AuthAction) Logout(ctx context.Context, payload *jwtoken.RefreshPayload) (uint32, error) {
	ll := a.logger.Named("Logout").With("sessionID", payload.SessionID)

	session, err := a.checkRefresh(ctx, payload)
	if err != nil {
		if IsAuthErr(err) {
			ll.Named("checkRefresh").AuthTag().Info(err.Error())
			return 0, err
		}

		ll.Named("checkRefresh").Nested(err.Error())
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

	sessionOld, err := a.checkRefresh(ctx, payload)
	if err != nil {
		ll.Named("checkRefresh").Nested(err.Error())
		return nil, err
	}

	refreshID := uuid.New()

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
