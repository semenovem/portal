package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/jwtoken"
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

	if nested := cnt.com.ExtractFormFromRequest(c, form); nested != nil {
		ll.Named("ExtractFormFromRequest").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	ll = ll.With("login", form.Login)

	authSession, authErr, err := cnt.authAct.NewLogin(
		ctx,
		form.Login,
		form.Passwd,
		c.Request().UserAgent(),
		form.DeviceID,
	)
	if err != nil {
		ll.Named("NewLogin").Nested(err.Error())
		return cnt.failing.SendInternalServerErr(c, "", err)
	}

	if authErr != "" {
		ll.AuthTag().Named("NewLogin").Nested(string(authErr))
		return cnt.failing.Send(c, "", http.StatusBadRequest, txt.AuthInvalidLogoPasswd, err)
	}

	pair, err := cnt.jwt.NewPairTokens(&jwtoken.TokenParams{
		SessionID: authSession.ID,
		UserID:    authSession.UserID,
		RefreshID: authSession.RefreshID,
	})

	if err != nil {
		ll.Named("NewPairTokens").Error(err.Error())
		return cnt.failing.SendInternalServerErr(c, "", err)
	}

	for _, cookie := range cnt.refreshTokenCookies(pair.Refresh) {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, loginResponse{
		AccessToken:  pair.Access,
		RefreshToken: pair.Refresh,
		UserID:       authSession.UserID,
	})
}

// Logout docs
//
//	@Summary	Выход из авторизованной сессии пользователя
//	@Description
//	@Produce	json
//	@Success	200		{object}	loginResponse
//	@Failure	400		{object}	failing.Response
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

	refreshCookie, err := c.Cookie(cnt.jwtRefreshTokenCookieName)
	if err != nil {
		ll.Named("Cookie").Error(err.Error())
		return cnt.failing.Send(c, "", http.StatusOK, err)
	}

	payload, err := cnt.jwt.GetRefreshPayload(refreshCookie.Value)
	if err != nil {
		ll.Named("GetRefreshPayload").Error(err.Error())
		return cnt.failing.Send(c, "", http.StatusOK, err)
	}

	ll = ll.With("payload", payload)

	if payload.IsExpired() {
		ll.Named("IsExpired").Error("refresh token is expired")
		// TODO сообщение в аудит безопасности
		return cnt.failing.Send(c, "", http.StatusOK, err)
	}

	if err = cnt.authAct.Logout(ctx, payload.SessionID); err != nil {
		ll.Named("Logout").Nested(err.Error())
		return cnt.failing.Send(c, "", http.StatusOK, err)
	}

	return c.NoContent(http.StatusNoContent)
}
