package people_action

import (
	"context"
	people_provider "github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/util"
)

func (a *PeopleAction) EmployeeHandbook(
	ctx context.Context,
	opts *people_provider.EmployeesSearchOpts,
) (*EmployeesSearchResult, error) {
	var (
		ll = a.logger.Func(ctx, "EmployeeHandbook")
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
		var p, d = em.PositionID(), em.DeptID()

		if _, ok := positionIDMap[p]; !ok {
			positionIDMap[p] = struct{}{}
		}
		if _, ok := deptIDMap[d]; !ok {
			deptIDMap[d] = struct{}{}
		}
	}

	ret := &EmployeesSearchResult{
		Total:       total,
		Employees:   employees,
		PositionMap: nil,
		DeptMap:     nil,
		UserBossMap: nil,
	}

	if ret.PositionMap, err = a.peoplePvd.GetPositionMap(ctx, util.NumMapToArr(positionIDMap)); err != nil {
		ll.Named("peoplePvd.GetPositionMap").Nested(err)
		return nil, err
	}
	if ret.DeptMap, err = a.peoplePvd.GetDeptMap(ctx, util.NumMapToArr(deptIDMap)); err != nil {
		ll.Named("peoplePvd.GetDeptMap").Nested(err)
		return nil, err
	}

	return ret, nil
}
