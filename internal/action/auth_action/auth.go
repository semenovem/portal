package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg/audit"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/logger"
)

const (
	authErrNoFoundUserByLogin = "no found user by login"
	authErrPasswdIncorrect    = "password is incorrect"
	authErrUserNotWorks       = "user not works"
)

type AuthErr string

// NewLogin авторизация пользователя по логопасу
func (a *AuthAction) NewLogin(
	ctx context.Context,
	login, passwd, userAgent string,
	deviceID uuid.UUID,
) (*it.AuthSession, AuthErr, error) {
	var (
		ll           = a.logger.Named("NewLogin")
		auditPayload = audit.P{"login": login, "deviceID": deviceID, "userAgent": userAgent}
	)

	loggingUser, err := a.peoplePvd.GetUserByLogin(ctx, login)
	if err != nil {
		ll = ll.Named("GetUserByLogin")

		if provider.IsNoRows(err) {
			ll.Tags(logger.AuthTag, logger.ClientTag).Info(authErrNoFoundUserByLogin)
			a.audit.Refusal(audit.UserLogin, authErrNoFoundUserByLogin, auditPayload)

			return nil, authErrNoFoundUserByLogin, nil
		}

		ll.Nested(err.Error())
		return nil, "", err
	}

	auditPayload["userID"] = loggingUser

	if !a.passwdAuth.Matching(loggingUser.PasswdHash, passwd) {
		ll.Named("Matching").Tags(logger.AuthTag, logger.ClientTag).Debug(authErrPasswdIncorrect)
		a.audit.Refusal(audit.UserLogin, authErrPasswdIncorrect, auditPayload)

		return nil, authErrPasswdIncorrect, nil
	}

	session, authErr, err := a.newSession(ctx, loggingUser.ToUser(), deviceID)
	if err != nil || authErr != "" {
		ll = ll.Named("newSession")
		if err != nil {
			ll.Nested(err.Error())
		} else {
			ll.Nested(string(authErr))
		}

		a.audit.Refusal(audit.UserLogin, audit.Cause(authErr), auditPayload)

		return nil, authErr, err
	}

	auditPayload["sessionID"] = session.ID
	auditPayload["refreshID"] = session.RefreshID
	a.audit.Approved(audit.UserLogin, auditPayload)

	return session, "", nil
}

// NewSession создание новой авторизованной сессии
func (a *AuthAction) NewSession(
	ctx context.Context,
	user *it.User,
	userAgent string,
	deviceID uuid.UUID,
) (*it.AuthSession, AuthErr, error) {
	var (
		ll           = a.logger.Named("NewSession")
		auditPayload = audit.P{"userID": user.ID, "deviceID": deviceID, "userAgent": userAgent}
	)

	session, authErr, err := a.newSession(ctx, user, deviceID)
	if err != nil || authErr != "" {
		ll.Named("newSession").Nested(err.Error())
		return nil, authErr, err
	}

	auditPayload["sessionID"] = session.ID
	auditPayload["refreshID"] = session.RefreshID
	a.audit.Approved(audit.UserLogin, auditPayload)

	return session, "", nil
}

// Создание новой авторизованной сессии
func (a *AuthAction) newSession(
	ctx context.Context,
	user *it.User,
	deviceID uuid.UUID,
) (*it.AuthSession, AuthErr, error) {
	ll := a.logger.Named("newSession")

	if err := user.IsWorks(); err != nil {
		s := authErrUserNotWorks + "(" + err.Error() + ")"
		ll.AuthTag().Named("IsWorks").With("user", user).Debug(s)

		return nil, AuthErr(s), nil
	}

	session, err := a.authPvd.CreateSession(ctx, user.ID, deviceID)
	if err != nil {
		ll.Named("CreateSession").Nested(err.Error())
		return nil, "", err
	}

	return session, "", nil
}

// Logout разлогин Сессии
func (a *AuthAction) Logout(ctx context.Context, sessionID uint32) error {
	ll := a.logger.Named("Logout")

	if err := a.authPvd.LogoutSession(ctx, sessionID); err != nil {
		ll.Named("LogoutSession").Nested(err.Error())
		return err
	}

	// TODO сообщение в аудит безопасности

	return nil
}
