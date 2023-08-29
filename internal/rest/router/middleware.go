package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/it"
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
	authPvd *auth_provider.AuthProvider,
) echo.MiddlewareFunc {
	ll := logger.Named("tokenMiddleware")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenFromHeader := c.Request().Header.Get("Authorization")

			if tokenFromHeader == "" {
				ll.Auth(it.ErrAuthCookieEmpty)
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					it.ErrAuthCookieEmpty,
				)
			}

			split := strings.Fields(tokenFromHeader)
			if len(split) != 2 {
				ll.Auth(it.ErrInvalidBearer)
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					it.ErrInvalidBearer,
				)
			}

			payload, err := jwtService.GetAccessPayload(split[1])
			if err != nil {
				ll.Named("GetAccessPayload").Auth(err)
				return fail.Send(c, "", http.StatusUnauthorized, err)
			}

			if payload.IsExpired() {
				ll.AuthDebug(it.ErrAccessTokenExp)
				return fail.Send(
					c,
					"",
					http.StatusUnauthorized,
					it.ErrAccessTokenExp,
				)
			}

			// Проверить отозванные сессии
			isCancel, err := authPvd.IsSessionCanceled(c.Request().Context(), payload.SessionID)
			if err != nil {
				ll.Named("IsSessionCanceled").Nested(err)
				return fail.SendInternalServerErr(c, "", err)
			}

			if isCancel {
				ll.AuthDebug(it.ErrUserLogouted)
				return fail.Send(c, "", http.StatusUnauthorized, it.ErrUserLogouted)
			}

			c.Set(controller.ThisUserID, payload.UserID)

			return next(c)
		}
	}
}
