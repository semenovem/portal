package people_controller

import (
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
)

type Controller struct {
	logger    pkg.Logger
	fail      *fail.Service
	com       *controller.Common
	audit     *audit.AuditProvider
	peopleAct *people_action.PeopleAction
}

func New(
	arg *controller.CntArgs,
	peopleAct *people_action.PeopleAction,
) *Controller {
	return &Controller{
		logger:    arg.Logger.Named("people-cnt"),
		fail:      arg.FailureService,
		com:       arg.Common,
		audit:     arg.Audit,
		peopleAct: peopleAct,
	}
}
