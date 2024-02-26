package auth_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc/auth"
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
	"net/http"
	"time"
)

type Config struct {
	JWTServedDomains          []string
	JWTRefreshTokenLife       time.Duration
	JWTRefreshTokenCookieName string
}

type Controller struct {
	logger          pkg.Logger
	mainConfig      *config.Platform
	fail            *fail.Service
	jwt             *jwtoken.Service
	loginPasswdAuth it.LoginPasswdAuthenticator
	com             *controller.Common
	authAct         *auth_action.AuthAction
	audit           *audit.AuditProvider
}

func New(
	arg *controller.InitArgs,
	jwt *jwtoken.Service,
	loginPasswdAuth it.LoginPasswdAuthenticator,
	authAct *auth_action.AuthAction,
) *Controller {
	return &Controller{
		logger:          arg.Logger.Named("auth-cnt"),
		mainConfig:      arg.MainConfig,
		fail:            arg.FailureService,
		loginPasswdAuth: loginPasswdAuth,
		com:             arg.Common,
		authAct:         authAct,
		audit:           arg.Audit,
		jwt:             jwt,
	}
}

func (cnt *Controller) refreshTokenCookies(refreshToken string) []*http.Cookie {
	cookies := make([]*http.Cookie, 0, len(cnt.mainConfig.Auth.JWT.ServedDomains.Val))

	for _, domain := range cnt.mainConfig.Auth.JWT.ServedDomains.Val {
		cookie := &http.Cookie{
			Name:   cnt.mainConfig.Auth.JWT.RefreshTokenCookieName,
			Path:   "/",
			Domain: domain,
			//Secure:   true,
			//HttpOnly: true,
			//SameSite: http.SameSiteLaxMode,
		}

		if refreshToken == "" {
			cookie.Expires = time.Now().Add(-666 * time.Second)
		} else {
			cookie.Value = refreshToken
			cookie.Expires = time.Now().Add(cnt.mainConfig.Auth.JWT.RefreshTokenLifetime.Val)
		}

		cookies = append(cookies, cookie)
	}

	return cookies
}

// Получить токен и проверить срок его действия
func (cnt *Controller) extractRefreshToken(c echo.Context) (*jwtoken.RefreshPayload, fail.Nested) {
	var (
		ll = cnt.logger.Func(c.Request().Context(), "extractRefreshToken")
	)

	refreshCookie, err := c.Cookie(cnt.mainConfig.Auth.JWT.RefreshTokenCookieName)
	if err != nil {
		ll.Named("Cookie").Auth(err)
		return nil, fail.NewNested(http.StatusUnauthorized, err)
	}

	payload, err := cnt.jwt.GetRefreshPayload(refreshCookie.Value)
	if err != nil {
		ll.Named("GetRefreshPayload").Auth(err)
		return nil, fail.NewNested(http.StatusUnauthorized, err)
	}

	if payload.IsExpired() {
		ll.With("payload", payload).AuthStr("refresh token is expired")
		return nil, fail.NewNested(http.StatusUnauthorized, err)
	}

	return payload, nil
}

// ExtractRefreshToken выпустить новый токен
func (cnt *Controller) pairToken(session *auth.Session) (*jwtoken.PairTokens, fail.Nested) {
	pair, err := cnt.jwt.NewPairTokens(&jwtoken.TokenParams{
		SessionID: session.ID,
		UserID:    session.UserID,
		RefreshID: session.RefreshID,
	})
	if err != nil {
		cnt.logger.Named("pairToken").Auth(err)
		return nil, fail.NewNested(http.StatusInternalServerError, err)
	}

	return pair, nil
}
