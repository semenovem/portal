package router

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/internal/provider/auth_provider"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/jwtoken"
	"net/http"
	"runtime"
	"strings"
)

type panicMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func panicRecover(logger pkg.Logger, cli bool) echo.MiddlewareFunc {
	ll := logger.Named("panicRecover")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					var (
						stack  = make([]byte, middleware.DefaultRecoverConfig.StackSize)
						length = runtime.Stack(stack, !middleware.DefaultRecoverConfig.DisableStackAll)
						msg    = fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					)

					if cli {
						fmt.Println(msg)
					} else {
						ll.Error(msg)
					}

					c.JSON(http.StatusInternalServerError, panicMessage{
						Code:    http.StatusInternalServerError,
						Message: "Internal Server Error",
					})
				}
			}()

			return next(c)
		}
	}
}

func tokenMiddleware(
	logger pkg.Logger,
	fail *failing.Service,
	jwtService *jwtoken.Service,
	AuthPvd *auth_provider.AuthProvider,
) echo.MiddlewareFunc {
	ll := logger.Named("tokenMiddleware")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenFromHeader := c.Request().Header.Get("Authorization")

			if tokenFromHeader == "" {
				ll.AuthTag().Debug("empty header [Authorization] token")
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					errors.New("empty header [Authorization] token"),
				)
			}

			split := strings.Fields(tokenFromHeader)
			if len(split) != 2 {
				ll.AuthTag().Debug("invalid bearer token")
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					errors.New("invalid bearer token"),
				)
			}

			payload, err := jwtService.GetAccessPayload(split[1])
			if err != nil {
				ll.AuthTag().Named("GetAccessPayload").Info(err.Error())
				return fail.Send(c, "", http.StatusUnauthorized, err)
			}

			if payload.IsExpired() {
				ll.AuthTag().Debug("access token expired")
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					errors.New("access token expired"),
				)
			}

			// Проверить отозванные сессии
			isCancel, err := AuthPvd.IsSessionCanceled(c.Request().Context(), payload.SessionID)
			if err != nil {
				ll.Named("IsSessionCanceled").Nested(err.Error())
				return fail.SendInternalServerErr(c, "", err)
			}

			if isCancel {
				ll.Debug("user is logouted")
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					errors.New("user is logouted"),
				)
			}

			c.Set(controller.ThisUserID, payload.UserID)

			return next(c)
		}
	}
}
