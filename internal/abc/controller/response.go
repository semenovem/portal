package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/throw"
	"github.com/semenovem/portal/pkg/txt"
	"net/http"
)

// Response ответ клиенту
func (a *Common) Response(c echo.Context, ll pkg.Logger, err error) error {
	var (
		retCode int
		text    string
	)

	switch t := err.(type) {

	case throw.AccessErr:
		retCode = http.StatusForbidden

	case throw.InvalidErr:
		retCode = http.StatusBadRequest

	case throw.BadRequestErr:
		ll.BadRequest(err)
		retCode = http.StatusBadRequest

		switch t.Target() {
		case throw.Err400DuplicateEmail:
			text = txt.RestrictDuplicateEmail
		case throw.Err400DuplicateLogin:
			text = txt.RestrictDuplicateLogin
		case throw.Err400FiredBehind:
			text = txt.RuleFiredBehindWorked
		}

	case throw.NotFoundErr:
		ll.NotFound(err)
		retCode = http.StatusNotFound

		switch t {
		case throw.Err404User:
			text = txt.NotFoundUser
		}

	case throw.AuthErr:
		ll.Auth(err)
		retCode = http.StatusUnauthorized

	default:
		ll.Error(err.Error())
		return a.fail.SendInternalServerErr(c, "", err)
	}

	if text != "" {
		return a.fail.Send(c, "", retCode, err, text)
	}

	return a.fail.Send(c, "", retCode, err)
}
