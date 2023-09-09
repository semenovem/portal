package people_dto

import (
	"github.com/semenovem/portal/internal/abc/people"
)

// EmployeesSearchOpts параметры поиска сотрудников
type EmployeesSearchOpts struct {
	Limit  uint32
	Offset uint32
	//Expired         bool       // Включая у кого истек срок действия
	//ExpiredAtAfter  *time.Time // Включая истекшие записи после
	//ExpiredAtBefore *time.Time // Включая у кого истек срок действия
	Fired bool // Включая уволенных
}

// EmployeesSearchResult результат поиска сотрудников
type EmployeesSearchResult struct {
	Total       uint32
	Employees   []*people.Employee
	PositionMap map[uint16]*people.UserPosition
	DeptMap     map[uint16]*people.UserDept
	UserBossMap map[uint16]*people.UserBoss
}
