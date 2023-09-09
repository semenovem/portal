package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people"
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
	"time"
)

func (p *PeopleProvider) SearchEmployees(
	ctx context.Context,
	opts *people_dto.EmployeesSearchOpts,
) (*people_dto.EmployeesSearchResult, error) {
	var (
		total uint32
		//bossIDs     = make([]uint32, 0)
		//positionIDs = make([]uint32, 0)
		//deptIDs     = make([]uint32, 0)
		ls = make([]*people.Employee, 0)

		sq = `SELECT e.user_id
		FROM people.employees AS e
		WHERE (e.fired_at IS NULL OR e.fired_at > now())`
	)

	rows, err := p.db.Query(ctx, sq)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := people.Employee{
			ID:          0,
			Status:      "",
			Roles:       nil,
			AvatarID:    0,
			FirstName:   "",
			Surname:     "",
			Note:        "",
			ExpiredAt:   nil,
			PositionID:  0,
			DeptID:      0,
			StartWorkAt: time.Time{},
			FiredAt:     nil,
		}

		if err = rows.Scan(
			&o.ID,
		); err != nil {
			return nil, err
		}

		ls = append(ls, &o)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return &people_dto.EmployeesSearchResult{
		Total:     total,
		Employees: ls,
	}, err
}
