package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/failing"
	"net/http"
)

// Login docs
//
//	@Summary	Авторизация пользователя
//	@Description
//	@Produce	json
//	@Param		payload	body		LoginForm	true	"Логин/пароль"
//	@Success	200		{object}	LoginResponse
//	@Failure	400		{object}	failing.Response
//	@Router		/auth/login [POST]
//	@Tags		auth
//	@Security	ApiKeyAuth
func (cnt *Controller) Login(c echo.Context) error {
	var (
		ll   = cnt.logger.Named("Login")
		form = new(LoginForm)
	)

	if nested := cnt.act.ExtractFormFromRequest(c, form); nested != nil {
		ll.Named("ExtractFormFromRequest").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	f := failing.Response{
		Code:             0,
		Message:          "",
		ValidationErrors: nil,
		AdditionalFields: nil,
	}

	return c.JSON(http.StatusOK, f)
}
