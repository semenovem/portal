package store_controller

import (
	"github.com/semenovem/portal/internal/abc/store/action"
	"github.com/semenovem/portal/internal/provider/audit_provider"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type Controller struct {
	logger   pkg.Logger
	failing  *failing.Service
	com      *controller.Common
	storeAct *store_action.StoreAction
	audit    *audit_provider.AuditProvider
}

func New(
	arg *controller.CntArgs,
	storeAct *store_action.StoreAction,
	audit *audit_provider.AuditProvider,
) *Controller {
	return &Controller{
		logger:   arg.Logger.Named("auth-cnt"),
		failing:  arg.Failing,
		com:      arg.Common,
		storeAct: storeAct,
		audit:    audit,
	}
}
