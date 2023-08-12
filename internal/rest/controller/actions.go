package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/provider"
	authprovider "github.com/semenovem/portal/internal/provider/auth"
	peopleprovider "github.com/semenovem/portal/internal/provider/people"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/entity"
	"github.com/semenovem/portal/pkg/failing"
	"net/http"
)

type ActionConfig struct {
	Logger         pkg.Logger
	Failing        *failing.Service
	PeopleProvider *peopleprovider.Provider
	AuthProvider   *authprovider.Provider
}

type Action struct {
	logger         pkg.Logger
	failing        *failing.Service
	peopleProvider *peopleprovider.Provider
	authProvider   *authprovider.Provider
}

func NewAction(c *ActionConfig) *Action {
	return &Action{
		logger:         c.Logger.Named("Action"),
		failing:        c.Failing,
		peopleProvider: c.PeopleProvider,
		authProvider:   c.AuthProvider,
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

// GetUserByLogin получить пользователя по email
func (a *Action) GetUserByLogin(ctx context.Context, login string) (*entity.LoggingUser, failing.Nested) {
	user, err := a.peopleProvider.GetUserByLogin(ctx, login)
	if err != nil {
		ll := a.logger.Named("GetUserByLogin").With("login", login)

		if provider.IsNoRows(err) {
			ll.Named("CLIENT").Info("no found user by login")
			return nil, failing.NewNested(http.StatusNotFound, err)
		}

		ll.Named("DATABASE").Error(err.Error())
		return nil, failing.NewNested(http.StatusInternalServerError, err)
	}

	return user, nil
}
