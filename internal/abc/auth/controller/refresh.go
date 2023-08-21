package auth_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"

	_ "github.com/semenovem/portal/pkg/failing"
)

// Refresh docs
//
//	@Summary	Обновление токена авторизации
//	@Description
//	@Produce	json
//	@Param		refresh-token	header		string	true	"refresh токен"
//	@Success	200				{object}	refreshResponse
//	@Failure	400				{object}	failing.Response
//	@Router		/auth/refresh [POST]
//	@Tags		auth
//	@Security	ApiKeyAuth
func (cnt *Controller) Refresh(c echo.Context) error {
	var (
		ll  = cnt.logger.Named("Refresh")
		ctx = c.Request().Context()
	)

	payload, nested := cnt.extractRefreshToken(c)
	if nested != nil {
		ll.Named("ExtractRefreshToken").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	ll = ll.With("payload", payload)

	session, err := cnt.authAct.Refresh(ctx, payload)
	if err != nil {
		ll.Named("Refresh").Nested(err.Error())
		return cnt.failing.Send(c, "", http.StatusOK, err)
	}

	pair, nested := cnt.pairToken(session)
	if nested != nil {
		ll.Named("pairToken").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	for _, cookie := range cnt.refreshTokenCookies(pair.Refresh) {
		c.SetCookie(cookie)
	}

	ll.With("sessionIDNew", session.ID).Debug("success")

	return c.JSON(http.StatusOK, refreshResponse{
		AccessToken:  pair.Access,
		RefreshToken: pair.Refresh,
	})
}
