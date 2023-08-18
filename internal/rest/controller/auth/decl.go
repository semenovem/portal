package auth

import (
	"github.com/semenovem/portal/internal/action/auth_action"
	"github.com/semenovem/portal/internal/rest/controller"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/jwtoken"
	"net/http"
	"time"
)

type Controller struct {
	logger                    pkg.Logger
	failing                   *failing.Service
	jwt                       *jwtoken.Service
	com                       *controller.Common
	authAct                   *auth_action.AuthAction
	jwtServedDomains          []string
	jwtRefreshTokenLife       time.Duration
	jwtRefreshTokenCookieName string
}

func New(
	arg *controller.CntArgs,
	jwt *jwtoken.Service,
	authAct *auth_action.AuthAction,
	jwtServedDomains []string,
	jwtRefreshTokenLife time.Duration,
	jwtRefreshTokenCookieName string,
) *Controller {
	return &Controller{
		logger:                    arg.Logger.Named("auth-cnt"),
		failing:                   arg.Failing,
		com:                       arg.Common,
		authAct:                   authAct,
		jwt:                       jwt,
		jwtServedDomains:          jwtServedDomains,
		jwtRefreshTokenLife:       jwtRefreshTokenLife,
		jwtRefreshTokenCookieName: jwtRefreshTokenCookieName,
	}
}

func (cnt *Controller) refreshTokenCookies(refreshToken string) []*http.Cookie {
	cookies := make([]*http.Cookie, 0, len(cnt.jwtServedDomains))

	for _, domain := range cnt.jwtServedDomains {
		cookie := &http.Cookie{
			Name:   cnt.jwtRefreshTokenCookieName,
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
			cookie.Expires = time.Now().Add(cnt.jwtRefreshTokenLife * time.Second)
		}

		cookies = append(cookies, cookie)
	}

	return cookies
}
