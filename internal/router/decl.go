package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc"
	"github.com/semenovem/portal/internal/abc/auth/controller"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/media"
	"github.com/semenovem/portal/internal/abc/media/controller"
	"github.com/semenovem/portal/internal/abc/people/controller"
	"github.com/semenovem/portal/internal/abc/store/controller"
	"github.com/semenovem/portal/internal/abc/vehicle/controller"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
	"github.com/semenovem/portal/pkg/txt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
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

func New(
	ctx context.Context,
	logger pkg.Logger,
	config *config.API,
	auditService *audit.AuditProvider,
	jwtService *jwtoken.Service,
	loginPasswdAuth it.LoginPasswdAuthenticator,

	providers *abc.Providers,
	actions *abc.Actions,
) (*Router, error) {
	var (
		ll  = logger.Named("router")
		e   = echo.New()
		err error
	)

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, panicMessage{
			Code:    http.StatusNotFound,
			Message: "method didn't exists",
		})
	}

	e.Use(echoprometheus.NewMiddleware("myapp"))
	e.GET("/metrics", echoprometheus.NewHandler())

	corsConfig := middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXFrameOptions,
			echo.HeaderXContentTypeOptions,
			echo.HeaderContentSecurityPolicy,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
		MaxAge: 60,
	}

	e.Use(
		//middleware.Logger(),
		panicRecover(ll, config.Base.CliMode),
		middleware.CORSWithConfig(corsConfig),
	)

	if e.Validator, err = newValidation(); err != nil {
		ll.Named("newValidation").Error(err.Error())
		return nil, err
	}

	failService := fail.New(&fail.Config{
		IsDevMode:             config.IsDev(),
		Logger:                logger,
		Messages:              txt.GetMessages(),
		ValidationMessageMap:  validators,
		HTTPStatuses:          txt.GetHTTPStatuses(),
		UnknownMessage:        unknownFail,
		InvalidRequestMessage: invalidFail,
	})

	controllerInitArgs := &controller.InitArgs{
		Logger:         logger,
		FailureService: failService,
		Audit:          auditService,
		Common: controller.NewAction(
			logger,
			failService,
			providers.Auth,
			providers.People,
		),
	}

	r := &Router{
		ctx:    ctx,
		logger: logger.Named("router"),
		server: e,
		addr:   fmt.Sprintf(":%d", config.Rest.Port),

		vehicleCnt: vehicle_controller.New(controllerInitArgs),

		authCnt: auth_controller.New(
			controllerInitArgs,
			jwtService,
			loginPasswdAuth,
			actions.Auth,
			strings.Split(config.JWT.ServedDomains, ","),
			time.Hour*24*time.Duration(config.JWT.RefreshTokenLifetimeDay),
			config.JWT.RefreshTokenCookieName,
		),

		peopleCnt: people_controller.New(controllerInitArgs, loginPasswdAuth, actions.People),
		storeCnt:  store_controller.New(controllerInitArgs, actions.Store),
		mediaCnt: media_controller.New(controllerInitArgs, &media.ConfigMedia{
			AvatarMaxBytes: uint32(config.Upload.AvatarMaxMB) * 1024 * 1024,
			ImageMaxBytes:  uint32(config.Upload.ImageMaxMB) * 1024 * 1024,
			VideoMaxBytes:  uint32(config.Upload.VideoMaxMB) * 1024 * 1024,
			DocMaxBytes:    uint32(config.Upload.DocMaxMB) * 1024 * 1024,
		}, actions.Media),
	}

	r.unauth = e.Group("")
	r.auth = r.unauth.Group("", tokenMiddleware(
		logger,
		failService,
		jwtService,
		providers.Auth,
	))

	r.addRoutes()

	return r, nil
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
