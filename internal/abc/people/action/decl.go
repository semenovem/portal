package people_action

import (
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
)

type PeopleAction struct {
	logger         pkg.Logger
	userPasswdAuth it.UserPasswdAuthenticator
	peoplePvd      *people_provider.PeopleProvider
}

func New(
	logger pkg.Logger,
	userPasswdAuth it.UserPasswdAuthenticator,
	peoplePvd *people_provider.PeopleProvider,
) *PeopleAction {
	return &PeopleAction{
		logger:         logger.Named("PeopleAction"),
		userPasswdAuth: userPasswdAuth,
		peoplePvd:      peoplePvd,
	}
}

// EmployeesSearchResult результат поиска сотрудников
type EmployeesSearchResult struct {
	Total       uint32
	Employees   []*people.EmployeeSlim
	PositionMap map[uint16]*people.UserPosition
	DeptMap     map[uint16]*people.UserDept
	UserBossMap map[uint16]*people.UserBoss
}
