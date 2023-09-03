package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/fail"
	"net/http"
)

// ExtractForm получить данные из запроса
func (a *Common) ExtractForm(c echo.Context, form interface{}) fail.Nested {
	if err := c.Bind(form); err != nil {
		a.logger.Named("ExtractForm.bind").BadRequest(err)
		return fail.NewNested(http.StatusBadRequest, err)
	}

	if err := c.Validate(form); err != nil {
		a.logger.Named("ExtractForm.validate").With("form", form).BadRequest(err)
		return fail.NewNestedValidation(err)
	}

	return nil
}

// ExtractThisUser получить данные авторизованного пользователя из запроса
func (a *Common) ExtractThisUser(c echo.Context) (uint32, fail.Nested) {
	userID, ok := c.Get(ThisUserID).(uint32)
	if !ok {
		err := a.logger.Named("ExtractThisUser").BadRequestStrRetErr("invalid format user_id")
		return 0, fail.NewNested(http.StatusBadRequest, err)
	}

	return userID, nil
}

// ExtractUserAndForm получить данные из запроса и авторизованного пользователя
func (a *Common) ExtractUserAndForm(c echo.Context, form interface{}) (uint32, fail.Nested) {
	if nested := a.ExtractForm(c, form); nested != nil {
		a.logger.Named("ExtractForm.bind").Nestedf(nested.Message())
		return 0, nested
	}

	userID, nested := a.ExtractThisUser(c)
	if nested != nil {
		a.logger.Named("ExtractThisUser").Nestedf(nested.Message())
		return 0, nested
	}

	return userID, nil
}
