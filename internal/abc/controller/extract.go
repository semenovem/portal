package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg"
	"net/http"
)

// ExtractForm New получить данные формы
func (a *Common) ExtractForm(c echo.Context, ll pkg.Logger, form interface{}) error {
	if err := c.Bind(form); err != nil {
		ll.Named("ExtractForm.bind").BadRequest(err)
		return a.fail.Send(c, "", http.StatusBadRequest, err)
	}

	if err := c.Validate(form); err != nil {
		ll.Named("ExtractForm.validate").With("form", form).BadRequest(err)
		return a.fail.SendValidationErr(c, "", err)
	}

	return nil
}

func ExtractThis(c echo.Context) This {
	userID, ok := c.Get(ThisUserID).(uint32)
	if ok {
		return This{
			UserID: userID,
		}
	}

	panic(fmt.Sprintf(
		"there is no user in the request context: path=%s, key=%s",
		c.Path(),
		ThisUserID,
	))
}

func ExtractThisUserID(c echo.Context) uint32 {
	userID, ok := c.Get(ThisUserID).(uint32)
	if ok {
		return userID
	}

	panic(fmt.Sprintf(
		"there is no user in the request context: path=%s, key=%s",
		c.Path(),
		ThisUserID,
	))
}
