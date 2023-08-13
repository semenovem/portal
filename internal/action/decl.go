package action

import (
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
)

type AuthAct struct {
	logger     pkg.Logger
	audit      *provider.Audit
	authPvd    *provider.AuthPvd
	peoplePvd  *provider.PeoplePvd
	passwdAuth it.UserPasswdAuthenticator
}

func NewAuth(
	logger pkg.Logger,
	passwdAuth it.UserPasswdAuthenticator,
	audit *provider.Audit,
	authPvd *provider.AuthPvd,
	peoplePvd *provider.PeoplePvd,
) *AuthAct {
	return &AuthAct{
		logger:     logger.Named("AuthAct"),
		passwdAuth: passwdAuth,
		audit:      audit,
		authPvd:    authPvd,
		peoplePvd:  peoplePvd,
	}
}
