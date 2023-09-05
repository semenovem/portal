package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeopleProvider) SearchEmployees(
	ctx context.Context,
	opts *people.EmployeesSearchOpts,
) (*people.EmployeesSearchResult, error) {
	var (
		total uint32
		ls    = make([]*it.EmployeeProfile, 0)

		sq = `SELECT e.user_id
		FROM people.employees AS e
		WHERE (e.fired_at IS NULL OR e.fired_at > now())`
	)

	rows, err := p.db.Query(ctx, sq)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := it.EmployeeProfile{}

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

	return &people.EmployeesSearchResult{
		Total:     total,
		Employees: ls,
	}, err
}
