package router

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/semenovem/portal/internal/provider/auth_provider"
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
				return fail.Send(c, "", http.StatusUnauthorized, errors.New("empty header [Authorization] token"))
			}

			split := strings.Fields(tokenFromHeader)
			if len(split) != 2 {
				ll.AuthTag().Debug("invalid bearer token")
				return fail.Send(c, "", http.StatusUnauthorized, errors.New("invalid bearer token"))
			}

			payload, err := jwtService.GetAccessPayload(split[1])
			if err != nil {
				ll.AuthTag().Named("GetAccessPayload").Info(err.Error())
				return fail.Send(c, "", http.StatusUnauthorized, err)
			}

			if payload.IsExpired() {
				ll.AuthTag().Debug("access token expired")
				return fail.Send(c, "", http.StatusUnauthorized, errors.New("access token expired"))
			}

			// Проверить отозванные сессии
			isCancel, err := AuthPvd.IsSessionCanceled(c.Request().Context(), payload.SessionID)
			if err != nil {
				ll.Named("IsSessionCanceled").Nested(err.Error())
				return fail.SendInternalServerErr(c, "", err)
			}

			fmt.Printf("!!!!!!!!!!! = isCancel =  %+v\n", payload)
			fmt.Printf("!!!!!!!!!!! = isCancel =  %+v\n", isCancel)

			//if !token.Valid {
			//	ll.Named("AUTH").Warn("Token invalid")
			//	return fail.Send(c, "", http.StatusUnauthorized, errors.New("token invalid"))
			//}
			//
			//authToken, err := extractAuthTokenFromAccessToken(token, c)
			//if err != nil {
			//	return fail.Send(c, "", http.StatusUnauthorized, err)
			//}
			//
			//ctx := c.Request().Context()
			//
			//has, err := action.HasTokenRevokedInRedis(ctx, authToken.RefreshID)
			//if err != nil {
			//	ll.Named("REDIS").With("authToken", authToken).Error(err)
			//	return fail.SendInternalServerErr(c, "", err)
			//}
			//
			//if has {
			//	err = errors.New("refresh token has been revoked")
			//	ll.Named("AUTH").With("authToken", authToken).Debug(err)
			//
			//	return fail.Send(c, "", http.StatusUnauthorized, err)
			//}
			//
			//has, nested := hasUserAuth2factor(c.Request().Context(), authToken)
			//if nested != nil {
			//	ll.Named("NESTED").Debug(nested.Message())
			//	return fail.SendNested(c, "", nested)
			//}
			//
			//c.Set(common.AuthTokenKeyName, authToken)
			//c.Set(authUserHave2FA, has)

			return next(c)
		}
	}
}
