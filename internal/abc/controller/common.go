package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"net/http"
)

const (
	ThisUserID = "this_user_id"
)

type Common struct {
	logger    pkg.Logger
	failing   *failing.Service
	authPvd   *auth_provider.AuthProvider
	peoplePvd *people_provider.PeopleProvider
}

func NewAction(
	logger pkg.Logger,
	failure *failing.Service,
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,
) *Common {
	return &Common{
		logger:    logger.Named("Common"),
		failing:   failure,
		authPvd:   authPvd,
		peoplePvd: peoplePvd,
	}
}

// ExtractForm получить данные из запроса
func (a *Common) ExtractForm(c echo.Context, form interface{}) failing.Nested {
	if err := c.Bind(form); err != nil {
		a.logger.Named("ExtractForm.bind").ClientTag().Debug(err.Error())
		return failing.NewNested(http.StatusBadRequest, err)
	}

	if err := c.Validate(form); err != nil {
		a.logger.Named("ExtractForm.validate").ClientTag().With("form", form).Debug(err.Error())
		return failing.NewNestedValidation(err)
	}

	return nil
}

// ExtractThisUser получить данные авторизованного пользователя из запроса
func (a *Common) ExtractThisUser(c echo.Context) (uint32, failing.Nested) {
	userID, ok := c.Get(ThisUserID).(uint32)
	if !ok {
		a.logger.Named("ExtractThisUser").ClientTag().Debug("invalid format user_id")
		return 0, failing.NewNested(http.StatusBadRequest, errors.New("invalid format user_id"))
	}

	return userID, nil
}

// ExtractUserAndForm получить данные из запроса и авторизованного пользователя
func (a *Common) ExtractUserAndForm(c echo.Context, form interface{}) (uint32, failing.Nested) {
	if nested := a.ExtractForm(c, form); nested != nil {
		a.logger.Named("ExtractForm.bind").Nested(nested.Message())
		return 0, nested
	}

	userID, nested := a.ExtractThisUser(c)
	if nested != nil {
		a.logger.Named("ExtractThisUser").Nested(nested.Message())
		return 0, nested
	}

	return userID, nil
}
