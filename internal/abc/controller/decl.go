package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/throw"
	"net/http"
)

const (
	ThisUserID = "this_user_id"
)

type Common struct {
	logger    pkg.Logger
	fail      *fail.Service
	authPvd   *auth_provider.AuthProvider
	peoplePvd *people_provider.PeopleProvider
}

func NewAction(
	logger pkg.Logger,
	fail *fail.Service,
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,
) *Common {
	return &Common{
		logger:    logger.Named("Common"),
		fail:      fail,
		authPvd:   authPvd,
		peoplePvd: peoplePvd,
	}
}

// Response ответ клиенту
func (a *Common) Response(c echo.Context, err error) error {
	switch err.(type) {

	case *throw.AccessErr:
		fmt.Printf(">>>>>>>>>>  %T \n", err)

		return a.fail.Send(c, "", http.StatusForbidden, err)
	case *throw.BadRequestErr:
		return a.fail.Send(c, "", http.StatusBadRequest, err)
	case throw.NotFoundErr:
		return a.fail.Send(c, "", http.StatusNotFound, err)
	}

	return a.fail.SendInternalServerErr(c, "", err)
}
