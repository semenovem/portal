package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc"
	"github.com/semenovem/portal/internal/abc/auth/controller"
	"github.com/semenovem/portal/internal/abc/controller"
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
	"time"
)

func New(
	ctx context.Context,
	logger pkg.Logger,
	mainConfig *config.Main,
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

	e.Use(
		echoprometheus.NewMiddleware("myapp"),
		contextMiddleware(time.Millisecond*time.Duration(mainConfig.Controller.MinTimeContextMs)),
	)

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
		panicRecover(ll, mainConfig.Base.CliMode),
		middleware.CORSWithConfig(corsConfig),
	)

	if e.Validator, err = newValidation(); err != nil {
		ll.Named("newValidation").Error(err.Error())
		return nil, err
	}

	failService := fail.New(&fail.Config{
		IsDevMode:             mainConfig.IsDev(),
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
		MainConfig:     mainConfig,
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
		addr:   fmt.Sprintf(":%d", mainConfig.Rest.Port),

		vehicleCnt: vehicle_controller.New(controllerInitArgs),

		authCnt: auth_controller.New(
			controllerInitArgs,
			jwtService,
			loginPasswdAuth,
			actions.Auth,
		),

		peopleCnt: people_controller.New(controllerInitArgs, loginPasswdAuth, actions.People),
		storeCnt:  store_controller.New(controllerInitArgs, actions.Store),
		mediaCnt:  media_controller.New(controllerInitArgs, actions.Media),
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
