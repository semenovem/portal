package people_dto

import "github.com/semenovem/portal/internal/abc/people"

// EmployeesSearchOpts параметры поиска сотрудников
type EmployeesSearchOpts struct {
	Limit  uint32
	Offset uint32
	//Expired         bool       // Включая у кого истек срок действия
	//ExpiredAtAfter  *time.Time // Включая истекшие записи после
	//ExpiredAtBefore *time.Time // Включая у кого истек срок действия
	Fired    bool // Включая уволенных
	Statuses []people.UserStatus
}
