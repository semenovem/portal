package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/provider/auth_provider"
	"github.com/semenovem/portal/internal/provider/people_provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"net/http"
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

// ExtractFormFromRequest получить данные из запроса
func (a *Common) ExtractFormFromRequest(c echo.Context, form interface{}) failing.Nested {
	if err := c.Bind(form); err != nil {
		a.logger.Named("GetFormFromRequest.bind.CLIENT").Debug(err.Error())
		return failing.NewNested(http.StatusBadRequest, err)
	}

	if err := c.Validate(form); err != nil {
		a.logger.Named("GetFormFromRequest.validate.CLIENT").With("form", form).Debug(err.Error())
		return failing.NewNestedValidation(err)
	}

	return nil
}
