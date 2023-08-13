package auth

import (
	"github.com/semenovem/portal/internal/action"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/jwtoken"
)

type Controller struct {
	logger    pkg.Logger
	failing   *failing.Service
	jwt       *jwtoken.Service
	com       *controller.Common
	peoplePvd *provider.PeoplePvd
	authAct   *action.AuthAct
}

func New(
	arg *controller.CntArgs,
	jwt *jwtoken.Service,
	peoplePvd *provider.PeoplePvd,
	authAct *action.AuthAct,
) *Controller {
	return &Controller{
		logger:    arg.Logger.Named("auth-cnt"),
		failing:   arg.Failing,
		com:       arg.Common,
		peoplePvd: peoplePvd,
		authAct:   authAct,
	}
}
