package auth

import (
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type Controller struct {
	logger  pkg.Logger
	failing *failing.Service
}

func New(arg *controller.CntArgs) *Controller {
	return &Controller{
		logger:  arg.Logger.Named("vehicle"),
		failing: arg.Failing,
	}
}
