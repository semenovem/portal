package media_controller

import (
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/provider/audit_provider"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type Controller struct {
	logger   pkg.Logger
	failing  *failing.Service
	com      *controller.Common
	mediaAct *media_action.MediaAction
	audit    *audit_provider.AuditProvider
}

func New(
	arg *controller.CntArgs,
	mediaAct *media_action.MediaAction,
	audit *audit_provider.AuditProvider,
) *Controller {
	return &Controller{
		logger:   arg.Logger.Named("auth-cnt"),
		failing:  arg.Failing,
		com:      arg.Common,
		mediaAct: mediaAct,
		audit:    audit,
	}
}
