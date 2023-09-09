package people_action

import (
	"context"
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
)

func (a *PeopleAction) EmployeeHandbook(
	ctx context.Context,
	opts *people_dto.EmployeesSearchOpts,
) (*EmployeesSearchResult, error) {
	var (
		ll = a.logger.Named("EmployeeHandbook")
		//bossIDs     = make([]uint32, 0)
		positionIDMap = make(map[uint16]struct{}, 0)
		deptIDMap     = make(map[uint16]struct{}, 0)
	)

	employees, total, err := a.peoplePvd.SearchEmployees(ctx, opts)
	if err != nil {
		ll.Named("peoplePvd.SearchEmployees").Nested(err)
		return nil, err
	}

	for _, em := range employees {
		if _, ok := positionIDMap[em.PositionID]; !ok {
			positionIDMap[em.PositionID] = struct{}{}
		}
		if _, ok := deptIDMap[em.DeptID]; !ok {
			deptIDMap[em.DeptID] = struct{}{}
		}
	}

	ret := &EmployeesSearchResult{
		Total:       total,
		Employees:   employees,
		PositionMap: nil,
		DeptMap:     nil,
		UserBossMap: nil,
	}

	return ret, nil
}
