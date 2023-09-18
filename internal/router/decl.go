package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/auth/controller"
	"github.com/semenovem/portal/internal/abc/media/controller"
	"github.com/semenovem/portal/internal/abc/people/controller"
	"github.com/semenovem/portal/internal/abc/store/controller"
	"github.com/semenovem/portal/internal/abc/vehicle/controller"
	"github.com/semenovem/portal/pkg"
	"net/http"
)

const (
	requestIDKeyName = "request_id"
)

type Router struct {
	ctx                 context.Context
	logger              pkg.Logger
	server              *echo.Echo
	addr                string
	unauth, auth, admin *echo.Group
	vehicleCnt          *vehicle_controller.Controller
	authCnt             *auth_controller.Controller
	peopleCnt           *people_controller.Controller
	storeCnt            *store_controller.Controller
	mediaCnt            *media_controller.Controller
}

func (r *Router) Start() {
	go func() {
		<-r.ctx.Done()
		if err := r.server.Close(); err != nil {
			r.logger.Named("Close").Error(err.Error())
		}
	}()

	r.logger.Infof("router start on %s", r.addr)

	r.server.HidePort = true
	r.server.HideBanner = true

	if err := r.server.Start(r.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		r.logger.Named("Start").Error(err.Error())
	}
}

func ExtractRequestID(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(requestIDKeyName))
}
