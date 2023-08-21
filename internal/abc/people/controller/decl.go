package people_controller

import (
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type Controller struct {
	logger    pkg.Logger
	failing   *failing.Service
	act       *controller.Common
	peopleAct *people_action.PeopleAction
}

func New(
	arg *controller.CntArgs,
	peopleAct *people_action.PeopleAction,
) *Controller {
	return &Controller{
		logger:    arg.Logger.Named("people-cnt"),
		failing:   arg.Failing,
		act:       arg.Common,
		peopleAct: peopleAct,
	}
}
