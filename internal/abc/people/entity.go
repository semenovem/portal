package people

import "github.com/semenovem/portal/pkg/it"

// EmployeesSearchOpts параметры поиска сотрудников
type EmployeesSearchOpts struct {
	Limit  uint32
	Offset uint32
}

// EmployeesSearchResult результат поиска сотрудников
type EmployeesSearchResult struct {
	Total     uint32
	Employees []*it.EmployeeProfile
}
