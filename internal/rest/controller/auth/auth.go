package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/txt"
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

	fmt.Println("!!!!!!!!!!!! authSession = ", authSession)

	f := failing.Response{
		Code:             0,
		Message:          "",
		ValidationErrors: nil,
		AdditionalFields: nil,
	}

	return c.JSON(http.StatusOK, f)
}
