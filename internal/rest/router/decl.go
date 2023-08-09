package router

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/rest/controller"
	authCnt "github.com/semenovem/portal/internal/rest/controller/auth"
	vehicleCnt "github.com/semenovem/portal/internal/rest/controller/vehicle"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/txt"
	"net/http"
)

type Config struct {
	Ctx    context.Context
	Logger pkg.Logger
	DB     *pgxpool.Pool
	Redis  *redis.Client
	Global *config.API
}

type Router struct {
	ctx                 context.Context
	logger              pkg.Logger
	server              *echo.Echo
	addr                string
	unauth, auth, admin *echo.Group
	vehicleCnt          *vehicleCnt.Controller
	authCnt             *authCnt.Controller
}

func New(cfg *Config) (*Router, error) {
	var (
		ll = cfg.Logger.Named("router")
		e  = echo.New()
	)

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, panicMessage{
			Code:    http.StatusNotFound,
			Message: "method didn't exists",
		})
	}

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
		middleware.Logger(),
		panicRecover(ll, cfg.Global.Base.CliMode),
		middleware.CORSWithConfig(corsConfig),
	)

	var err error

	if e.Validator, err = newValidation(); err != nil {
		ll.Named("newValidation").Error(err.Error())
		return nil, err
	}

	failure := failing.New(&failing.Config{
		IsDevMode:             cfg.Global.IsDev(),
		Logger:                cfg.Logger,
		Messages:              txt.GetMessages(),
		ValidationMessageMap:  validators,
		HTTPStatuses:          txt.GetHTTPStatuses(),
		UnknownMessage:        unknownFailing,
		InvalidRequestMessage: invalidFailing,
	})

	// REST контроллеры
	arg := controller.CntArgs{
		Logger:  cfg.Logger,
		Failing: failure,
		Act:     controller.NewAction(cfg.Logger, failure),
	}

	r := &Router{
		ctx:        cfg.Ctx,
		logger:     cfg.Logger.Named("router"),
		server:     e,
		addr:       cfg.Global.RestPort,
		vehicleCnt: vehicleCnt.New(&arg),
		authCnt:    authCnt.New(&arg),
	}

	g := e.Group("v1")

	r.unauth = g
	r.auth = g
	r.admin = g

	r.addRoutes()

	return r, nil
}
