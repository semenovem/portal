package auth_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg/txt"
	"net/http"

	_ "github.com/semenovem/portal/pkg/failing"
)

// Login docs
//
//	@Summary	Авторизация пользователя
//	@Description
//	@Produce	json
//	@Param		payload	body		loginForm	true	"Логин/пароль"
//	@Success	200		{object}	loginResponse
//	@Failure	400		{object}	failing.Response
//	@Router		/auth/login [POST]
//	@Tags		auth
//	@Security	ApiKeyAuth
func (cnt *Controller) Login(c echo.Context) error {
	var (
		ll   = cnt.logger.Named("Login")
		form = new(loginForm)
		ctx  = c.Request().Context()
	)

	if nested := cnt.com.ExtractForm(c, form); nested != nil {
		ll.Named("ExtractForm").Nestedf(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	ll = ll.With("login", form.Login)

	session, err := cnt.authAct.NewLogin(
		ctx,
		form.Login,
		form.Passwd,
		form.DeviceID,
	)
	if err != nil {
		ll.Named("NewLogin").Nested(err)

		if auth_action.IsAuthErr(err) {
			return cnt.failing.Send(c, "", http.StatusBadRequest, txt.AuthInvalidLogoPasswd, err)
		}

		return cnt.failing.SendInternalServerErr(c, "", err)
	}

	pair, nested := cnt.pairToken(session)
	if nested != nil {
		ll.Named("pairToken").Nestedf(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	cnt.audit.Auth(session.UserID, audit.UserLogin, audit.P{
		"user-agent": c.Request().UserAgent(),
		"sessionID":  session.ID,
	})

	for _, cookie := range cnt.refreshTokenCookies(pair.Refresh) {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, loginResponse{
		refreshResponse: refreshResponse{
			AccessToken:  pair.Access,
			RefreshToken: pair.Refresh,
		},
		UserID: session.UserID,
	})
}

// Logout docs
//
//	@Summary	Выход из авторизованной сессии пользователя
//	@Description
//	@Produce	json
//	@Param		refresh-token	header		string	true	"refresh токен"
//	@Success	200				{object}	loginResponse
//	@Failure	400				{object}	failing.Response
//	@Router		/auth/logout [POST]
//	@Tags		auth
//	@Security	ApiKeyAuth
func (cnt *Controller) Logout(c echo.Context) error {
	var (
		ll  = cnt.logger.Named("Logout")
		ctx = c.Request().Context()
	)

	for _, cookie := range cnt.refreshTokenCookies("") {
		c.SetCookie(cookie)
	}

	payload, nested := cnt.extractRefreshToken(c)
	if nested != nil {
		ll.Named("GetRefreshPayload").Nestedf(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	ll = ll.With("payload", payload)

	userID, err := cnt.authAct.Logout(ctx, payload)
	if err != nil {
		ll.Named("Logout").Nested(err)

		if auth_action.IsAuthErr(err) {
			return cnt.failing.Send(c, "", http.StatusUnauthorized, err)
		}

		return cnt.failing.SendInternalServerErr(c, "", err)
	}

	cnt.audit.Auth(userID, audit.UserLogout, audit.P{
		"user-agent": c.Request().UserAgent(),
		"sessionID":  payload.SessionID,
	})

	return c.NoContent(http.StatusNoContent)
}
