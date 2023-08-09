package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"net/http"
)

type Action struct {
	logger  pkg.Logger
	failing *failing.Service
}

func NewAction(logger pkg.Logger, failure *failing.Service) *Action {
	return &Action{
		logger:  logger,
		failing: failure,
	}
}

// ExtractFormFromRequest получить данные из запроса
func (a *Action) ExtractFormFromRequest(c echo.Context, form interface{}) failing.Nested {
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
