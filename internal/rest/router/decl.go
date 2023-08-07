package router

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/internal/rest/controller/vehicle"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
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
	ctx                  context.Context
	logger               pkg.Logger
	server               *echo.Echo
	addr                 string
	nonAuth, auth, admin *echo.Group
	vehicleCnt           *vehicle.Controller
}

func New(cfg *Config) *Router {
	var (
		ll = cfg.Logger.Named("router")
		e  = echo.New()
	)

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

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, panicMessage{
			Code:    http.StatusNotFound,
			Message: "method didn't exists",
		})
	}

	failure := failing.New(&failing.Config{
		IsDevMode:             false,
		Logger:                ll,
		TranslatorDefault:     nil,
		Translators:           nil,
		Messages:              nil,
		ValidationMessageMap:  nil,
		HTTPStatuses:          nil,
		UnknownMessage:        failing.Message{},
		InvalidRequestMessage: failing.Message{},
	})

	// REST контроллеры
	arg := controller.CntArgs{
		Logger:  cfg.Logger,
		Failing: failure,
	}

	r := &Router{
		ctx:        cfg.Ctx,
		logger:     cfg.Logger.Named("router"),
		server:     e,
		addr:       cfg.Global.RestPort,
		vehicleCnt: vehicle.New(&arg),
	}

	g := e.Group("v1")

	r.nonAuth = g
	r.auth = g
	r.admin = g

	r.addRoutes()

	return r
}
