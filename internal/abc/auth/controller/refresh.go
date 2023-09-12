package auth_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/throw"
	"net/http"

	_ "github.com/semenovem/portal/pkg/fail"
)

// Refresh docs
//
//	@Summary	Обновление токена авторизации
//	@Description
//	@Produce	json
//	@Param		refresh-token	header		string	true	"refresh токен"
//	@Success	200				{object}	refreshResponse
//	@Failure	400				{object}	fail.Response
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
		ll.Named("ExtractRefreshToken").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
	}

	ll = ll.With("payload", payload)

	session, err := cnt.authAct.Refresh(ctx, payload)
	if err != nil {
		ll.Named("Refresh").Nested(err)

		if throw.IsNotFoundErr(err) {
			err = throw.NewAuthErr(throw.Err404AuthSession.Error())
		}
		return cnt.com.Response(c, ll, err)
	}

	pair, nested := cnt.pairToken(session)
	if nested != nil {
		ll.Named("pairToken").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
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
