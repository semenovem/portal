package vehicle_controller

import (
	controller2 "github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type Controller struct {
	logger  pkg.Logger
	failing *failing.Service
	act     *controller2.Common
}

func New(arg *controller2.CntArgs) *Controller {
	return &Controller{
		logger:  arg.Logger.Named("vehicle-cnt"),
		failing: arg.FailureService,
		act:     arg.Common,
	}
}
