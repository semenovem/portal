package auth_action

import (
	"github.com/semenovem/portal/internal/provider/audit_provider"
	"github.com/semenovem/portal/internal/provider/auth_provider"
	"github.com/semenovem/portal/internal/provider/people_provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
)

type AuthAction struct {
	logger     pkg.Logger
	audit      *audit_provider.AuditProvider
	authPvd    *auth_provider.AuthProvider
	peoplePvd  *people_provider.PeopleProvider
	passwdAuth it.UserPasswdAuthenticator
}

func NewAuth(
	logger pkg.Logger,
	passwdAuth it.UserPasswdAuthenticator,
	audit *audit_provider.AuditProvider,
	authPvd *auth_provider.AuthProvider,
	peoplePvd *people_provider.PeopleProvider,
) *AuthAction {
	return &AuthAction{
		logger:     logger.Named("AuthAction"),
		passwdAuth: passwdAuth,
		audit:      audit,
		authPvd:    authPvd,
		peoplePvd:  peoplePvd,
	}
}
