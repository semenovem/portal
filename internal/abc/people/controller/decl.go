package people_controller

import (
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/it"
)

type Controller struct {
	logger         pkg.Logger
	mainConfig     *config.Main
	fail           *fail.Service
	com            *controller.Common
	audit          *audit.AuditProvider
	userPasswdAuth it.LoginPasswdAuthenticator
	peopleAct      *people_action.PeopleAction
}

func New(
	arg *controller.InitArgs,
	userPasswdAuth it.LoginPasswdAuthenticator,
	peopleAct *people_action.PeopleAction,
) *Controller {
	return &Controller{
		logger:         arg.Logger.Named("people-cnt"),
		mainConfig:     arg.MainConfig,
		fail:           arg.FailureService,
		com:            arg.Common,
		audit:          arg.Audit,
		userPasswdAuth: userPasswdAuth,
		peopleAct:      peopleAct,
	}
}
