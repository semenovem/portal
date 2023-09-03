package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/abc/auth/controller"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/abc/media/controller"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/abc/people/controller"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/abc/store/action"
	"github.com/semenovem/portal/internal/abc/store/controller"
	"github.com/semenovem/portal/internal/abc/vehicle/controller"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
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
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,

	authAct *auth_action.AuthAction,
	peopleAct *people_action.PeopleAction,
	storeAct *store_action.StoreAction,
	mediaAct *media_action.MediaAction,
) (*Router, error) {
	var (
		ll = logger.Named("router")
		e  = echo.New()
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

	var err error

	if e.Validator, err = newValidation(); err != nil {
		ll.Named("newValidation").Error(err.Error())
		return nil, err
	}

	failureService := fail.New(&fail.Config{
		IsDevMode:             config.IsDev(),
		Logger:                logger,
		Messages:              txt.GetMessages(),
		ValidationMessageMap:  validators,
		HTTPStatuses:          txt.GetHTTPStatuses(),
		UnknownMessage:        unknownFail,
		InvalidRequestMessage: invalidFail,
	})

	cntArg := &controller.CntArgs{
		Logger:         logger,
		FailureService: failureService,
		Audit:          auditService,
		Common: controller.NewAction(
			logger,
			failureService,
			authPvd,
			peoplePvd,
		),
	}

	r := &Router{
		ctx:    ctx,
		logger: logger.Named("router"),
		server: e,
		addr:   fmt.Sprintf(":%d", config.Rest.Port),

		vehicleCnt: vehicle_controller.New(cntArg),

		authCnt: auth_controller.New(
			cntArg,
			jwtService,
			authAct,
			strings.Split(config.JWT.ServedDomains, ","),
			time.Hour*24*time.Duration(config.JWT.RefreshTokenLifetimeDay),
			config.JWT.RefreshTokenCookieName,
		),

		peopleCnt: people_controller.New(cntArg, peopleAct),
		storeCnt:  store_controller.New(cntArg, storeAct),
		mediaCnt:  media_controller.New(cntArg, mediaAct),
	}

	r.unauth = e.Group("")
	r.auth = r.unauth.Group("", tokenMiddleware(
		logger,
		failureService,
		jwtService,
		authPvd,
	))

	r.addRoutes()

	return r, nil
}
