package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/action/auth_action"
	"github.com/semenovem/portal/internal/provider/audit_provider"
	"github.com/semenovem/portal/internal/provider/auth_provider"
	"github.com/semenovem/portal/internal/provider/people_provider"
	"github.com/semenovem/portal/internal/rest/controller"
	authController "github.com/semenovem/portal/internal/rest/controller/auth"
	"github.com/semenovem/portal/internal/rest/controller/people_cnt"
	vehicleController "github.com/semenovem/portal/internal/rest/controller/vehicle"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/jwtoken"
	"github.com/semenovem/portal/pkg/txt"
	"net/http"
	"strings"
	"time"
)

type Router struct {
	ctx                 context.Context
	logger              pkg.Logger
	server              *echo.Echo
	addr                string
	unauth, auth, admin *echo.Group
	vehicleCnt          *vehicleController.Controller
	authCnt             *authController.Controller
	peopleCnt           *people_cnt.Controller
}

func New(
	ctx context.Context,
	logger pkg.Logger,
	config config.API,
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,
	authAct *auth_action.AuthAction,
	audit *audit_provider.AuditProvider,
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
		panicRecover(ll, config.Base.CliMode),
		middleware.CORSWithConfig(corsConfig),
	)

	var err error

	if e.Validator, err = newValidation(); err != nil {
		ll.Named("newValidation").Error(err.Error())
		return nil, err
	}

	failure := failing.New(&failing.Config{
		IsDevMode:             config.IsDev(),
		Logger:                logger,
		Messages:              txt.GetMessages(),
		ValidationMessageMap:  validators,
		HTTPStatuses:          txt.GetHTTPStatuses(),
		UnknownMessage:        unknownFailing,
		InvalidRequestMessage: invalidFailing,
	})

	jwtService := jwtoken.New(&jwtoken.Config{
		AccessTokenSecret:    config.JWT.AccessTokenSecret,
		RefreshTokenSecret:   config.JWT.RefreshTokenSecret,
		AccessTokenLifetime:  time.Minute * time.Duration(config.JWT.AccessTokenLifetimeMin),
		RefreshTokenLifetime: time.Hour * 24 * time.Duration(config.JWT.RefreshTokenLifetimeDay),
	})

	// контроллеры
	common := controller.NewAction(
		logger,
		failure,
		authPvd,
		peoplePvd,
	)

	cntArg := &controller.CntArgs{
		Logger:  logger,
		Failing: failure,
		Common:  common,
	}

	var (
		authCnt = authController.New(
			cntArg,
			jwtService,
			authAct,
			audit,
			strings.Split(config.JWT.ServedDomains, ","),
			time.Hour*24*time.Duration(config.JWT.RefreshTokenLifetimeDay),
			config.JWT.RefreshTokenCookieName,
		)

		vehicleCnt = vehicleController.New(cntArg)
	)

	peopleCnt := people_cnt.New(cntArg)

	r := &Router{
		ctx:        ctx,
		logger:     logger.Named("router"),
		server:     e,
		addr:       fmt.Sprintf(":%d", config.Rest.Port),
		vehicleCnt: vehicleCnt,
		authCnt:    authCnt,
		peopleCnt:  peopleCnt,
	}

	r.unauth = e.Group("")
	r.auth = r.unauth.Group("", tokenMiddleware(logger, failure, jwtService, authPvd))

	r.addRoutes()

	return r, nil
}
