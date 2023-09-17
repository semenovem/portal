package people_action

import (
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
)

type PeopleAction struct {
	logger         pkg.Logger
	userPasswdAuth it.LoginPasswdAuthenticator
	peoplePvd      *people_provider.PeopleProvider
}

func New(
	logger pkg.Logger,
	userPasswdAuth it.LoginPasswdAuthenticator,
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
	Employees   []*people_provider.EmployeeModel
	PositionMap map[uint16]*people_provider.PositionModel
	DeptMap     map[uint16]*people_provider.DeptModel
	UserBossMap map[uint32]*people_provider.EmployeeModel
}
