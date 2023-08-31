package store_controller

import (
	controller2 "github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/store/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
)

type Controller struct {
	logger   pkg.Logger
	fail     *fail.Service
	com      *controller2.Common
	storeAct *store_action.StoreAction
	audit    *audit.AuditProvider
}

func New(
	arg *controller2.CntArgs,
	storeAct *store_action.StoreAction,
) *Controller {
	return &Controller{
		logger:   arg.Logger.Named("auth-cnt"),
		fail:     arg.FailureService,
		com:      arg.Common,
		storeAct: storeAct,
		audit:    arg.Audit,
	}
}
