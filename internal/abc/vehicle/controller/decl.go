package vehicle_controller

import (
	controller2 "github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
)

type Controller struct {
	logger pkg.Logger
	fail   *fail.Service
	act    *controller2.Common
}

func New(arg *controller2.CntArgs) *Controller {
	return &Controller{
		logger: arg.Logger.Named("vehicle-cnt"),
		fail:   arg.FailureService,
		act:    arg.Common,
	}
}
