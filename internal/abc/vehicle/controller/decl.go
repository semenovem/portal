package vehicle_controller

import (
	"github.com/semenovem/portal/config"
	controller2 "github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
)

type Controller struct {
	logger     pkg.Logger
	mainConfig *config.Platform
	fail       *fail.Service
	act        *controller2.Common
}

func New(arg *controller2.InitArgs) *Controller {
	return &Controller{
		logger:     arg.Logger.Named("vehicle-cnt"),
		mainConfig: arg.MainConfig,
		fail:       arg.FailureService,
		act:        arg.Common,
	}
}
