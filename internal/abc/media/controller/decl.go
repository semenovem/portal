package media_controller

import (
	controller2 "github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type Controller struct {
	logger   pkg.Logger
	failing  *failing.Service
	com      *controller2.Common
	mediaAct *media_action.MediaAction
	audit    *audit.AuditProvider
}

func New(
	arg *controller2.CntArgs,
	mediaAct *media_action.MediaAction,
) *Controller {
	return &Controller{
		logger:   arg.Logger.Named("auth-cnt"),
		failing:  arg.FailureService,
		audit:    arg.Audit,
		com:      arg.Common,
		mediaAct: mediaAct,
	}
}